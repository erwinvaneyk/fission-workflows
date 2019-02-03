// Bundle package contains integration tests that run using the bundle
package bundle

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/fission/fission-workflows/pkg/api"
	"github.com/fission/fission-workflows/pkg/apiserver"
	"github.com/fission/fission-workflows/pkg/fnenv/native/builtin"
	"github.com/fission/fission-workflows/pkg/types"
	"github.com/fission/fission-workflows/pkg/types/typedvalues"
	"github.com/fission/fission-workflows/pkg/util"
	"github.com/fission/fission-workflows/test/integration"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

const (
	TestSuiteTimeout = 10 * time.Minute
	TestTimeout      = time.Minute
	gRPCAddress      = ":5555"
)

func TestMain(m *testing.M) {
	if testing.Short() {
		log.Info("Short test; skipping bundle integration tests.")
		return
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), TestSuiteTimeout)
	integration.SetupBundle(ctx)

	time.Sleep(time.Duration(4) * time.Second)

	exitCode := m.Run()
	defer os.Exit(exitCode)
	// Teardown
	cancelFn()
	<-time.After(time.Duration(5) * time.Second) // Needed in order to let context cancel propagate
}

// Tests the submission of a workflow
func TestWorkflowCreate(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), TestTimeout)
	defer cancelFn()
	client := setup(ctx)

	// Test workflow creation
	spec := &types.WorkflowSpec{
		ApiVersion: types.WorkflowAPIVersion,
		OutputTask: "fakeFinalTask",
		Tasks: map[string]*types.TaskSpec{
			"fakeFinalTask": {
				FunctionRef: "noop",
			},
		},
	}
	wfId, err := client.Workflow.Create(ctx, spec)
	defer client.Workflow.Delete(ctx, wfId)
	assert.NoError(t, err)
	assert.NotNil(t, wfId)
	assert.NotEmpty(t, wfId.GetId())

	time.Sleep(time.Duration(2) * time.Second)
	// Test workflow list
	l, err := client.Workflow.List(ctx, &empty.Empty{})
	assert.NoError(t, err)
	if len(l.Workflows) != 1 || l.Workflows[0] != wfId.Id {
		t.Errorf("Listed workflows '%v' did not match expected workflow '%s'", l.Workflows, wfId.Id)
	}

	// Test workflow get
	wf, err := client.Workflow.Get(ctx, wfId)
	assert.NoError(t, err)
	assert.Equal(t, wf.Spec, spec)
	assert.Equal(t, wf.Status.Status, types.WorkflowStatus_READY)
}

func TestWorkflowInvocation(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), TestTimeout)
	defer cancelFn()
	client := setup(ctx)

	// Test workflow creation
	wfSpec := &types.WorkflowSpec{
		ApiVersion: types.WorkflowAPIVersion,
		OutputTask: "fakeFinalTask",
		Tasks: map[string]*types.TaskSpec{
			"fakeFinalTask": {
				FunctionRef: "noop",
				Inputs: map[string]*typedvalues.TypedValue{
					types.InputMain: typedvalues.MustWrap("{$.Tasks.FirstTask.Output}"),
				},
				Requires: map[string]*types.TaskDependencyParameters{
					"FirstTask": {},
				},
			},
			"FirstTask": {
				FunctionRef: "noop",
				Inputs: map[string]*typedvalues.TypedValue{
					types.InputMain: typedvalues.MustWrap("{$.Invocation.Inputs.default.toUpperCase()}"),
				},
			},
		},
	}
	wfm, err := client.Workflow.Create(ctx, wfSpec)
	assert.NoError(t, err)
	if wfm == nil || len(wfm.GetId()) == 0 {
		t.Errorf("Invalid id returned '%v'", wfm)
	}
	defer client.Workflow.Delete(ctx, wfm)

	// Create invocation
	expectedOutput := "Hello world!"
	tv, err := typedvalues.Wrap(expectedOutput)
	etv, err := typedvalues.Wrap(strings.ToUpper(expectedOutput))
	assert.NoError(t, err)

	wiSpec := &types.WorkflowInvocationSpec{
		WorkflowId: wfm.Id,
		Inputs: map[string]*typedvalues.TypedValue{
			types.InputMain: tv,
		},
	}
	result, err := client.Invocation.InvokeSync(ctx, wiSpec)
	assert.NoError(t, err)
	wiId := result.Metadata.Id

	// Test invocation list
	l, err := client.Invocation.List(ctx, &apiserver.InvocationListQuery{})
	assert.NoError(t, err)
	if len(l.Invocations) != 1 || l.Invocations[0] != wiId {
		t.Errorf("Listed invocations '%v' did not match expected invocation '%s'", l.Invocations, wiId)
	}

	// Test invocation get, give some slack to actually invoke it
	var invocation *types.WorkflowInvocation
	deadline := time.Now().Add(time.Duration(10) * time.Second)
	tick := time.NewTicker(time.Duration(100) * time.Millisecond)
	for ti := range tick.C {
		invoc, err := client.Invocation.Get(ctx, &types.ObjectMetadata{Id: wiId})
		assert.NoError(t, err)
		if invoc.Status.Finished() || ti.After(deadline) {
			invocation = invoc
			tick.Stop()
			break
		}
	}
	util.AssertProtoEqual(t, wiSpec, invocation.Spec)
	assert.Equal(t, etv.Value, invocation.Status.Output.Value)
	assert.True(t, invocation.Status.Successful())
}

func TestDynamicWorkflowInvocation(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), TestTimeout)
	defer cancelFn()
	client := setup(ctx)

	wfSpec := &types.WorkflowSpec{
		ApiVersion: types.WorkflowAPIVersion,
		OutputTask: "fakeFinalTask",
		Tasks: map[string]*types.TaskSpec{
			"fakeFinalTask": {
				FunctionRef: "noop",
				Inputs: map[string]*typedvalues.TypedValue{
					types.InputMain: typedvalues.MustWrap("{$.Tasks.someConditionalTask.Output}"),
				},
				Requires: map[string]*types.TaskDependencyParameters{
					"FirstTask":           {},
					"someConditionalTask": {},
				},
			},
			"FirstTask": {
				FunctionRef: "noop",
				Inputs: map[string]*typedvalues.TypedValue{
					types.InputMain: typedvalues.MustWrap("{$.Invocation.Inputs.default.toUpperCase()}"),
				},
			},
			"someConditionalTask": {
				FunctionRef: "if",
				Inputs: map[string]*typedvalues.TypedValue{
					builtin.IfInputCondition: typedvalues.MustWrap("{$.Invocation.Inputs.default == 'FOO'}"),
					builtin.IfInputThen: typedvalues.MustWrap(&types.TaskSpec{
						FunctionRef: "noop",
						Inputs: map[string]*typedvalues.TypedValue{
							types.InputMain: typedvalues.MustWrap("{'consequent: ' + $.Tasks.FirstTask.Output}"),
						},
					}),
					builtin.IfInputElse: typedvalues.MustWrap(&types.TaskSpec{
						FunctionRef: "noop",
						Inputs: map[string]*typedvalues.TypedValue{
							types.InputMain: typedvalues.MustWrap("{'alternative: ' + $.Tasks.FirstTask.Output}"),
						},
					}),
				},
				Requires: map[string]*types.TaskDependencyParameters{
					"FirstTask": {},
				},
			},
		},
	}
	wfResp, err := client.Workflow.Create(ctx, wfSpec)
	defer client.Workflow.Delete(ctx, wfResp)
	assert.NoError(t, err)
	assert.NotNil(t, wfResp)
	assert.NotEmpty(t, wfResp.Id)

	// Test with main input
	wiSpec := &types.WorkflowInvocationSpec{
		WorkflowId: wfResp.Id,
		Inputs: map[string]*typedvalues.TypedValue{
			types.InputMain: typedvalues.MustWrap("foo"),
		},
	}
	wfi, err := client.Invocation.InvokeSync(ctx, wiSpec)
	assert.NoError(t, err)
	assert.NotEmpty(t, wfi.Status.DynamicTasks)
	assert.True(t, wfi.Status.Finished())
	assert.True(t, wfi.Status.Successful())
	assert.Equal(t, 4, len(wfi.Status.Tasks))
	output := typedvalues.MustUnwrap(wfi.Status.Output)
	assert.Equal(t, "alternative: FOO", output)

	// Test with body input
	wiSpec = &types.WorkflowInvocationSpec{
		WorkflowId: wfResp.Id,
		Inputs: map[string]*typedvalues.TypedValue{
			types.InputBody: typedvalues.MustWrap("foo"),
		},
	}
	wfi, err = client.Invocation.InvokeSync(ctx, wiSpec)
	assert.NoError(t, err)
	assert.NotEmpty(t, wfi.Status.DynamicTasks)
	assert.True(t, wfi.Status.Finished())
	assert.True(t, wfi.Status.Successful())
	assert.Equal(t, 4, len(wfi.Status.Tasks))
	output = typedvalues.MustUnwrap(wfi.Status.Output)
	assert.Equal(t, "alternative: FOO", output)
}

func TestInlineWorkflowInvocation(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), TestTimeout)
	defer cancelFn()
	client := setup(ctx)

	wfSpec := &types.WorkflowSpec{
		ApiVersion: types.WorkflowAPIVersion,
		OutputTask: "finalTask",
		Tasks: map[string]*types.TaskSpec{
			"nestedTask": {
				FunctionRef: "noop",
				Inputs: map[string]*typedvalues.TypedValue{
					builtin.NoopInput: typedvalues.MustWrap(&types.WorkflowSpec{
						OutputTask: "b",
						Tasks: map[string]*types.TaskSpec{
							"a": {
								FunctionRef: "noop",
								Inputs: map[string]*typedvalues.TypedValue{
									types.InputMain: typedvalues.MustWrap("inner1"),
								},
							},
							"b": {
								FunctionRef: "noop",
								Inputs: map[string]*typedvalues.TypedValue{
									types.InputMain: typedvalues.MustWrap("{output('a')}"),
								},
								Requires: map[string]*types.TaskDependencyParameters{
									"a": nil,
								},
							},
						},
					}),
				},
			},
			"finalTask": {
				FunctionRef: "noop",
				Inputs: map[string]*typedvalues.TypedValue{
					types.InputMain: typedvalues.MustWrap("output('nestedTask')"),
				},
				Requires: map[string]*types.TaskDependencyParameters{
					"nestedTask": {},
				},
			},
		},
	}
	wfResp, err := client.Workflow.Create(ctx, wfSpec)
	defer client.Workflow.Delete(ctx, wfResp)
	assert.NoError(t, err)
	assert.NotNil(t, wfResp)
	assert.NotEmpty(t, wfResp.Id)

	wiSpec := &types.WorkflowInvocationSpec{
		WorkflowId: wfResp.Id,
	}
	wfi, err := client.Invocation.InvokeSync(ctx, wiSpec)
	assert.NoError(t, err)
	assert.NotEmpty(t, wfi.Status.DynamicTasks)
	assert.True(t, wfi.Status.Finished())
	assert.True(t, wfi.Status.Successful())
	assert.Equal(t, 3, len(wfi.Status.Tasks))

	_, err = typedvalues.Unwrap(wfi.Status.Output)
	assert.NoError(t, err)
}

func TestParallelInvocation(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.Dial(gRPCAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	cl := apiserver.NewWorkflowAPIClient(conn)
	wi := apiserver.NewWorkflowInvocationAPIClient(conn)

	wfSpec := types.NewWorkflowSpec()

	taskSpec := &types.TaskSpec{
		FunctionRef: builtin.Sleep,
		Inputs:      types.Input("25ms"),
	}

	wfSpec.AddTask("p1", taskSpec)
	wfSpec.AddTask("p2", taskSpec)
	wfSpec.AddTask("p3", taskSpec)
	wfSpec.AddTask("p4", taskSpec)
	wfSpec.AddTask("p5", taskSpec)
	wfSpec.AddTask("await", &types.TaskSpec{
		FunctionRef: builtin.Sleep,
		Inputs:      types.Input("10ms"),
		Requires:    types.Require("p1", "p2", "p3", "p4", "p5"),
	})
	wfSpec.SetOutput("await")

	wfResp, err := cl.Create(ctx, wfSpec)
	defer cl.Delete(ctx, wfResp)
	assert.NoError(t, err, err)
	assert.NotNil(t, wfResp)
	assert.NotEmpty(t, wfResp.Id)

	wiSpec := types.NewWorkflowInvocationSpec(wfResp.Id)
	wfi, err := wi.InvokeSync(ctx, wiSpec)
	assert.NoError(t, err)
	assert.Empty(t, wfi.Status.DynamicTasks)
	assert.True(t, wfi.Status.Finished())
	assert.True(t, wfi.Status.Successful())
	assert.Equal(t, len(wfSpec.Tasks), len(wfi.Status.Tasks))

	// Check if pN tasks were run in parallel
	var minStartTime, maxStartTime time.Time
	for _, task := range wfi.Status.Tasks {
		if strings.HasPrefix(task.Spec.TaskId, "p") {
			tt, err := ptypes.Timestamp(task.Metadata.CreatedAt)
			assert.NoError(t, err)
			if minStartTime == (time.Time{}) || tt.Before(minStartTime) {
				minStartTime = tt
			}
			if maxStartTime == (time.Time{}) || tt.After(maxStartTime) {
				maxStartTime = tt
			}
		}
	}
	assert.InDelta(t, 0, maxStartTime.Sub(minStartTime).Nanoseconds(), float64(time.Second.Nanoseconds()))
}

func TestLongRunningWorkflowInvocation(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), TestTimeout)
	defer cancelFn()
	client := setup(ctx)

	wfSpec := &types.WorkflowSpec{
		ApiVersion: types.WorkflowAPIVersion,
		OutputTask: "final",
		Tasks: types.Tasks{
			"longSleep": {
				FunctionRef: builtin.Sleep,
				Inputs:      types.Input("5s"),
			},
			"afterSleep": {
				FunctionRef: builtin.Noop,
				Inputs:      types.Input("{ '4' }"),
				Requires:    types.Require("longSleep"),
			},
			"parallel1": {
				FunctionRef: builtin.Noop,
				Inputs:      types.Input("{ '1' }"),
				Requires:    types.Require("longSleep"),
			},
			"parallel2": {
				FunctionRef: builtin.Noop,
				Inputs:      types.Input("{ output('parallel1') + '2' }"),
				Requires:    types.Require("parallel1"),
			},
			"parallel3": {
				FunctionRef: builtin.Noop,
				Inputs:      types.Input("{ output('parallel2') + '3' }"),
				Requires:    types.Require("parallel2"),
			},
			"merge": {
				FunctionRef: builtin.Noop,
				Inputs:      types.Input("{ output('parallel3') + output('afterSleep') }"),
				Requires:    types.Require("parallel3", "afterSleep"),
			},
			"final": {
				FunctionRef: builtin.Noop,
				Inputs:      types.Input("{ output('merge') }"),
				Requires:    types.Require("merge"),
			},
		},
	}
	wfResp, err := client.Workflow.Create(ctx, wfSpec)
	defer client.Workflow.Delete(ctx, wfResp)
	assert.NoError(t, err, err)
	assert.NotNil(t, wfResp)
	assert.NotEmpty(t, wfResp.Id)

	wiSpec := types.NewWorkflowInvocationSpec(wfResp.Id)
	wfi, err := client.Invocation.InvokeSync(ctx, wiSpec)
	assert.NoError(t, err)
	assert.Empty(t, wfi.Status.DynamicTasks)
	assert.True(t, wfi.Status.Finished())
	assert.True(t, wfi.Status.Successful())
	assert.Equal(t, len(wfSpec.Tasks), len(wfi.Status.Tasks))

	output := typedvalues.MustUnwrap(wfi.Status.Output)
	assert.Equal(t, "1234", output)
}

func TestWorkflowCancellation(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), TestTimeout)
	defer cancelFn()
	client := setup(ctx)
	wfSpec := &types.WorkflowSpec{
		ApiVersion: types.WorkflowAPIVersion,
		OutputTask: "longSleep2",
		Tasks: types.Tasks{
			"longSleep": {
				FunctionRef: builtin.Sleep,
				Inputs:      types.Input("250ms"),
			},
			"longSleep2": {
				FunctionRef: builtin.Sleep,
				Inputs:      types.Input("5s"),
				Requires:    types.Require("longSleep"),
			},
		},
	}

	wfResp, err := client.Workflow.Create(ctx, wfSpec)
	defer client.Workflow.Delete(ctx, wfResp)
	assert.NoError(t, err)
	assert.NotNil(t, wfResp)
	assert.NotEmpty(t, wfResp.Id)

	wiSpec := types.NewWorkflowInvocationSpec(wfResp.Id)

	// Invoke and cancel the invocation
	cancelCtx, cancelFn := context.WithCancel(ctx)
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancelFn()
	}()
	resp, err := client.Invocation.InvokeSync(cancelCtx, wiSpec)
	assert.Error(t, err)
	assert.Empty(t, resp)
	time.Sleep(500 * time.Millisecond)

	wfis, err := client.Invocation.List(ctx, &apiserver.InvocationListQuery{
		Workflows: []string{wfResp.GetId()},
	})
	assert.NoError(t, err)
	wfiID := wfis.Invocations[0]
	wfi, err := client.Invocation.Get(ctx, &types.ObjectMetadata{Id: wfiID})
	assert.NoError(t, err)
	assert.False(t, wfi.GetStatus().Successful())
	assert.True(t, wfi.GetStatus().Finished())
	assert.Equal(t, api.ErrInvocationCanceled, wfi.GetStatus().GetError().Error())
}

func TestInvocationInvalid(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), TestTimeout)
	defer cancelFn()
	client := setup(ctx)

	wfSpec := &types.WorkflowSpec{
		ApiVersion: types.WorkflowAPIVersion,
		OutputTask: "nonexistentTask",
		Tasks: types.Tasks{
			"task1": {
				FunctionRef: builtin.Noop,
			},
		},
	}
	_, err := client.Workflow.Create(ctx, wfSpec)
	assert.Error(t, err)
}

func TestInvocationFailed(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), TestTimeout)
	defer cancelFn()
	client := setup(ctx)

	msg := "expected error"
	wfSpec := &types.WorkflowSpec{
		ApiVersion: types.WorkflowAPIVersion,
		OutputTask: "task1",
		Tasks: types.Tasks{
			"task1": {
				FunctionRef: builtin.Fail,
				Inputs:      types.Input(msg),
			},
		},
	}
	wfResp, err := client.Workflow.Create(ctx, wfSpec)
	defer client.Workflow.Delete(ctx, wfResp)
	assert.NoError(t, err, err)
	assert.NotNil(t, wfResp)
	assert.NotEmpty(t, wfResp.Id)

	wiSpec := types.NewWorkflowInvocationSpec(wfResp.Id)
	wfi, err := client.Invocation.InvokeSync(ctx, wiSpec)
	assert.NoError(t, err)
	assert.Empty(t, wfi.Status.DynamicTasks)
	assert.True(t, wfi.Status.Finished())
	assert.False(t, wfi.Status.Successful())
	assert.Equal(t, len(wfSpec.Tasks), len(wfi.Status.Tasks))
	// assert.Equal(t, msg, wfi.GetStatus().GetError().GetMessage())
	// TODO generate consistent error report!
}

func TestInvocationWithForcedOutputs(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), TestTimeout)
	defer cancelFn()
	client := setup(ctx)

	// Test workflow creation
	output := typedvalues.MustWrap("overrided output")
	wfSpec := &types.WorkflowSpec{
		ApiVersion: types.WorkflowAPIVersion,
		OutputTask: "t3",
		Tasks: map[string]*types.TaskSpec{
			"t1": {
				FunctionRef: "noop",
				// Output with a literal value
				Output: output,
			},
			"t2": {
				FunctionRef: "noop",
				Inputs: map[string]*typedvalues.TypedValue{
					types.InputMain: typedvalues.MustWrap("{$.Tasks.t1.Output}"),
				},
				Requires: map[string]*types.TaskDependencyParameters{
					"t1": {},
				},
				// Self-referencing output
				Output: typedvalues.MustWrap("{$.Tasks.t2.Output}"),
			},
			"t3": {
				FunctionRef: "noop",
				Inputs: map[string]*typedvalues.TypedValue{
					types.InputMain: typedvalues.MustWrap("initial output 2"),
				},
				Requires: map[string]*types.TaskDependencyParameters{
					"t2": {},
				},
				// Referencing output of another task
				Output: typedvalues.MustWrap("{$.Tasks.t2.Output}"),
			},
		},
	}
	wfID, err := client.Workflow.Create(ctx, wfSpec)
	assert.NoError(t, err)
	wfi, err := client.Invocation.InvokeSync(ctx, &types.WorkflowInvocationSpec{
		WorkflowId: wfID.GetId(),
	})
	assert.NoError(t, err)
	util.AssertProtoEqual(t, output.GetValue(), wfi.GetStatus().GetTasks()["t1"].GetStatus().GetOutput().GetValue())
	util.AssertProtoEqual(t, output.GetValue(), wfi.GetStatus().GetTasks()["t2"].GetStatus().GetOutput().GetValue())
	util.AssertProtoEqual(t, output.GetValue(), wfi.GetStatus().GetOutput().GetValue())
}

func TestDeepRecursion(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), TestTimeout)
	defer cancelFn()
	client := setup(ctx)

	// Test workflow creation
	wfSpec := &types.WorkflowSpec{
		ApiVersion: types.WorkflowAPIVersion,
		OutputTask: "mainTask",
		Tasks: map[string]*types.TaskSpec{
			"mainTask": { // layer 1
				FunctionRef: builtin.Noop,
				Inputs: typedvalues.MustWrapMapTypedValue(map[string]interface{}{
					types.InputMain: &types.TaskSpec{ // layer 2
						FunctionRef: builtin.Noop,
						Inputs: typedvalues.MustWrapMapTypedValue(map[string]interface{}{
							types.InputMain: &types.TaskSpec{ // layer 3
								FunctionRef: builtin.Noop,
								Inputs: typedvalues.MustWrapMapTypedValue(map[string]interface{}{
									types.InputMain: &types.TaskSpec{ // layer 4
										FunctionRef: builtin.Noop,
										Inputs: typedvalues.MustWrapMapTypedValue(map[string]interface{}{
											types.InputMain: "foo",
										}),
									},
								}),
							},
						}),
					},
				}),
			},
		},
	}

	wfResp, err := client.Workflow.Create(ctx, wfSpec)
	defer client.Workflow.Delete(ctx, wfResp)

	assert.NoError(t, err)
	assert.NotNil(t, wfResp)
	assert.NotEmpty(t, wfResp.Id)

	wiSpec := &types.WorkflowInvocationSpec{
		WorkflowId: wfResp.Id,
		Inputs: typedvalues.MustWrapMapTypedValue(map[string]interface{}{
			types.InputMain: "foo",
		}),
	}
	wfi, err := client.Invocation.InvokeSync(ctx, wiSpec)
	assert.NoError(t, err)
	assert.NotEmpty(t, wfi.Status.DynamicTasks)
	assert.True(t, wfi.Status.Finished())
	assert.True(t, wfi.Status.Successful())

	output := typedvalues.MustUnwrap(wfi.Status.Output)
	d, _ := json.Marshal(output)
	fmt.Println(string(d))
	assert.Equal(t, typedvalues.MustUnwrap(wiSpec.Inputs[types.InputMain]), output)
}

func TestDeeplyNestedInvocation(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), TestTimeout)
	defer cancelFn()
	client := setup(ctx)

	// Test workflow creation
	wfSpec := &types.WorkflowSpec{
		ApiVersion: types.WorkflowAPIVersion,
		OutputTask: "CountUntil",
		Tasks: map[string]*types.TaskSpec{
			"CountUntil": {
				FunctionRef: builtin.While,
				Inputs: types.Inputs{
					builtin.WhileInputExpr:  typedvalues.MustWrap("{ !task().Inputs._prev || task().Inputs._prev < 5 }"),
					builtin.WhileInputLimit: typedvalues.MustWrap(10),
					builtin.WhileInputAction: typedvalues.MustWrap(&types.TaskSpec{
						FunctionRef: builtin.Noop,
						Inputs: types.Inputs{
							builtin.NoopInput: typedvalues.MustWrap("{ (task().Inputs._prev || 0) + 1 }"),
						},
					}),
				},
			},
		},
	}
	wfResp, err := client.Workflow.Create(ctx, wfSpec)
	defer client.Workflow.Delete(ctx, wfResp)

	assert.NoError(t, err)
	assert.NotNil(t, wfResp)
	assert.NotEmpty(t, wfResp.Id)

	wiSpec := &types.WorkflowInvocationSpec{
		WorkflowId: wfResp.Id,
	}
	wfi, err := client.Invocation.InvokeSync(ctx, wiSpec)
	assert.NoError(t, err)
	assert.NotEmpty(t, wfi.Status.DynamicTasks)
	assert.True(t, wfi.Status.Finished())
	assert.True(t, wfi.Status.Successful())

	output := typedvalues.MustUnwrap(wfi.Status.Output)
	assert.Equal(t, float64(5), output)
}

func setup(ctx context.Context) *apiserver.Client {
	conn, err := grpc.Dial(gRPCAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	c := apiserver.NewClient(conn)
	err = c.Await(ctx)
	if err != nil {
		panic(err)
	}
	return c
}
