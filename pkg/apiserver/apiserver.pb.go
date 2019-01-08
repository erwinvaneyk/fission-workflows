// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/apiserver/apiserver.proto

/*
Package apiserver is a generated protocol buffer package.

It is generated from these files:
	pkg/apiserver/apiserver.proto

It has these top-level messages:
	WorkflowList
	AddTaskRequest
	InvocationListQuery
	WorkflowInvocationList
	ObjectEvents
	Health
*/
package apiserver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import fission_workflows_types1 "github.com/fission/fission-workflows/pkg/types"
import fission_workflows_version "github.com/fission/fission-workflows/pkg/version"
import fission_workflows_eventstore "github.com/fission/fission-workflows/pkg/fes"
import google_protobuf2 "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WorkflowList struct {
	Workflows []string `protobuf:"bytes,1,rep,name=workflows" json:"workflows,omitempty"`
}

func (m *WorkflowList) Reset()                    { *m = WorkflowList{} }
func (m *WorkflowList) String() string            { return proto.CompactTextString(m) }
func (*WorkflowList) ProtoMessage()               {}
func (*WorkflowList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WorkflowList) GetWorkflows() []string {
	if m != nil {
		return m.Workflows
	}
	return nil
}

type AddTaskRequest struct {
	InvocationID string                         `protobuf:"bytes,1,opt,name=invocationID" json:"invocationID,omitempty"`
	Task         *fission_workflows_types1.Task `protobuf:"bytes,2,opt,name=task" json:"task,omitempty"`
}

func (m *AddTaskRequest) Reset()                    { *m = AddTaskRequest{} }
func (m *AddTaskRequest) String() string            { return proto.CompactTextString(m) }
func (*AddTaskRequest) ProtoMessage()               {}
func (*AddTaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddTaskRequest) GetInvocationID() string {
	if m != nil {
		return m.InvocationID
	}
	return ""
}

func (m *AddTaskRequest) GetTask() *fission_workflows_types1.Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type InvocationListQuery struct {
	Workflows []string `protobuf:"bytes,1,rep,name=workflows" json:"workflows,omitempty"`
}

func (m *InvocationListQuery) Reset()                    { *m = InvocationListQuery{} }
func (m *InvocationListQuery) String() string            { return proto.CompactTextString(m) }
func (*InvocationListQuery) ProtoMessage()               {}
func (*InvocationListQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *InvocationListQuery) GetWorkflows() []string {
	if m != nil {
		return m.Workflows
	}
	return nil
}

type WorkflowInvocationList struct {
	Invocations []string `protobuf:"bytes,1,rep,name=invocations" json:"invocations,omitempty"`
}

func (m *WorkflowInvocationList) Reset()                    { *m = WorkflowInvocationList{} }
func (m *WorkflowInvocationList) String() string            { return proto.CompactTextString(m) }
func (*WorkflowInvocationList) ProtoMessage()               {}
func (*WorkflowInvocationList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *WorkflowInvocationList) GetInvocations() []string {
	if m != nil {
		return m.Invocations
	}
	return nil
}

type ObjectEvents struct {
	Metadata *fission_workflows_types1.ObjectMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Events   []*fission_workflows_eventstore.Event    `protobuf:"bytes,2,rep,name=events" json:"events,omitempty"`
}

func (m *ObjectEvents) Reset()                    { *m = ObjectEvents{} }
func (m *ObjectEvents) String() string            { return proto.CompactTextString(m) }
func (*ObjectEvents) ProtoMessage()               {}
func (*ObjectEvents) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ObjectEvents) GetMetadata() *fission_workflows_types1.ObjectMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ObjectEvents) GetEvents() []*fission_workflows_eventstore.Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type Health struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *Health) Reset()                    { *m = Health{} }
func (m *Health) String() string            { return proto.CompactTextString(m) }
func (*Health) ProtoMessage()               {}
func (*Health) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Health) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*WorkflowList)(nil), "fission.workflows.apiserver.WorkflowList")
	proto.RegisterType((*AddTaskRequest)(nil), "fission.workflows.apiserver.AddTaskRequest")
	proto.RegisterType((*InvocationListQuery)(nil), "fission.workflows.apiserver.InvocationListQuery")
	proto.RegisterType((*WorkflowInvocationList)(nil), "fission.workflows.apiserver.WorkflowInvocationList")
	proto.RegisterType((*ObjectEvents)(nil), "fission.workflows.apiserver.ObjectEvents")
	proto.RegisterType((*Health)(nil), "fission.workflows.apiserver.Health")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for WorkflowAPI service

type WorkflowAPIClient interface {
	Create(ctx context.Context, in *fission_workflows_types1.WorkflowSpec, opts ...grpc.CallOption) (*fission_workflows_types1.ObjectMetadata, error)
	List(ctx context.Context, in *google_protobuf2.Empty, opts ...grpc.CallOption) (*WorkflowList, error)
	Get(ctx context.Context, in *fission_workflows_types1.ObjectMetadata, opts ...grpc.CallOption) (*fission_workflows_types1.Workflow, error)
	Delete(ctx context.Context, in *fission_workflows_types1.ObjectMetadata, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	Validate(ctx context.Context, in *fission_workflows_types1.WorkflowSpec, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	Events(ctx context.Context, in *fission_workflows_types1.ObjectMetadata, opts ...grpc.CallOption) (*ObjectEvents, error)
}

type workflowAPIClient struct {
	cc *grpc.ClientConn
}

func NewWorkflowAPIClient(cc *grpc.ClientConn) WorkflowAPIClient {
	return &workflowAPIClient{cc}
}

func (c *workflowAPIClient) Create(ctx context.Context, in *fission_workflows_types1.WorkflowSpec, opts ...grpc.CallOption) (*fission_workflows_types1.ObjectMetadata, error) {
	out := new(fission_workflows_types1.ObjectMetadata)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowAPI/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowAPIClient) List(ctx context.Context, in *google_protobuf2.Empty, opts ...grpc.CallOption) (*WorkflowList, error) {
	out := new(WorkflowList)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowAPI/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowAPIClient) Get(ctx context.Context, in *fission_workflows_types1.ObjectMetadata, opts ...grpc.CallOption) (*fission_workflows_types1.Workflow, error) {
	out := new(fission_workflows_types1.Workflow)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowAPI/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowAPIClient) Delete(ctx context.Context, in *fission_workflows_types1.ObjectMetadata, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowAPI/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowAPIClient) Validate(ctx context.Context, in *fission_workflows_types1.WorkflowSpec, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowAPI/Validate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowAPIClient) Events(ctx context.Context, in *fission_workflows_types1.ObjectMetadata, opts ...grpc.CallOption) (*ObjectEvents, error) {
	out := new(ObjectEvents)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowAPI/Events", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WorkflowAPI service

type WorkflowAPIServer interface {
	Create(context.Context, *fission_workflows_types1.WorkflowSpec) (*fission_workflows_types1.ObjectMetadata, error)
	List(context.Context, *google_protobuf2.Empty) (*WorkflowList, error)
	Get(context.Context, *fission_workflows_types1.ObjectMetadata) (*fission_workflows_types1.Workflow, error)
	Delete(context.Context, *fission_workflows_types1.ObjectMetadata) (*google_protobuf2.Empty, error)
	Validate(context.Context, *fission_workflows_types1.WorkflowSpec) (*google_protobuf2.Empty, error)
	Events(context.Context, *fission_workflows_types1.ObjectMetadata) (*ObjectEvents, error)
}

func RegisterWorkflowAPIServer(s *grpc.Server, srv WorkflowAPIServer) {
	s.RegisterService(&_WorkflowAPI_serviceDesc, srv)
}

func _WorkflowAPI_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fission_workflows_types1.WorkflowSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowAPIServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowAPI/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowAPIServer).Create(ctx, req.(*fission_workflows_types1.WorkflowSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowAPI_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf2.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowAPIServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowAPI/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowAPIServer).List(ctx, req.(*google_protobuf2.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fission_workflows_types1.ObjectMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowAPIServer).Get(ctx, req.(*fission_workflows_types1.ObjectMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowAPI_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fission_workflows_types1.ObjectMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowAPIServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowAPI/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowAPIServer).Delete(ctx, req.(*fission_workflows_types1.ObjectMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowAPI_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fission_workflows_types1.WorkflowSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowAPIServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowAPI/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowAPIServer).Validate(ctx, req.(*fission_workflows_types1.WorkflowSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowAPI_Events_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fission_workflows_types1.ObjectMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowAPIServer).Events(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowAPI/Events",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowAPIServer).Events(ctx, req.(*fission_workflows_types1.ObjectMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkflowAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fission.workflows.apiserver.WorkflowAPI",
	HandlerType: (*WorkflowAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _WorkflowAPI_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _WorkflowAPI_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _WorkflowAPI_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _WorkflowAPI_Delete_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _WorkflowAPI_Validate_Handler,
		},
		{
			MethodName: "Events",
			Handler:    _WorkflowAPI_Events_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/apiserver/apiserver.proto",
}

// Client API for WorkflowInvocationAPI service

type WorkflowInvocationAPIClient interface {
	// Create a new workflow invocation
	//
	// In case the invocation specification is missing fields or contains invalid fields, a HTTP 400 is returned.
	Invoke(ctx context.Context, in *fission_workflows_types1.WorkflowInvocationSpec, opts ...grpc.CallOption) (*fission_workflows_types1.ObjectMetadata, error)
	InvokeSync(ctx context.Context, in *fission_workflows_types1.WorkflowInvocationSpec, opts ...grpc.CallOption) (*fission_workflows_types1.WorkflowInvocation, error)
	AddTask(ctx context.Context, in *AddTaskRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	// Cancel a workflow invocation
	//
	// This action is irreverisble. A canceled invocation cannot be resumed or restarted.
	// In case that an invocation already is canceled, has failed or has completed, nothing happens.
	// In case that an invocation does not exist a HTTP 404 error status is returned.
	Cancel(ctx context.Context, in *fission_workflows_types1.ObjectMetadata, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	List(ctx context.Context, in *InvocationListQuery, opts ...grpc.CallOption) (*WorkflowInvocationList, error)
	// Get the specification and status of a workflow invocation
	//
	// Get returns three different aspects of the workflow invocation, namely the spec (specification), status and logs.
	// To lighten the request load, consider using a more specific request.
	Get(ctx context.Context, in *fission_workflows_types1.ObjectMetadata, opts ...grpc.CallOption) (*fission_workflows_types1.WorkflowInvocation, error)
	Events(ctx context.Context, in *fission_workflows_types1.ObjectMetadata, opts ...grpc.CallOption) (*ObjectEvents, error)
	Validate(ctx context.Context, in *fission_workflows_types1.WorkflowInvocationSpec, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
}

type workflowInvocationAPIClient struct {
	cc *grpc.ClientConn
}

func NewWorkflowInvocationAPIClient(cc *grpc.ClientConn) WorkflowInvocationAPIClient {
	return &workflowInvocationAPIClient{cc}
}

func (c *workflowInvocationAPIClient) Invoke(ctx context.Context, in *fission_workflows_types1.WorkflowInvocationSpec, opts ...grpc.CallOption) (*fission_workflows_types1.ObjectMetadata, error) {
	out := new(fission_workflows_types1.ObjectMetadata)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowInvocationAPI/Invoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowInvocationAPIClient) InvokeSync(ctx context.Context, in *fission_workflows_types1.WorkflowInvocationSpec, opts ...grpc.CallOption) (*fission_workflows_types1.WorkflowInvocation, error) {
	out := new(fission_workflows_types1.WorkflowInvocation)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowInvocationAPI/InvokeSync", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowInvocationAPIClient) AddTask(ctx context.Context, in *AddTaskRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowInvocationAPI/AddTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowInvocationAPIClient) Cancel(ctx context.Context, in *fission_workflows_types1.ObjectMetadata, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowInvocationAPI/Cancel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowInvocationAPIClient) List(ctx context.Context, in *InvocationListQuery, opts ...grpc.CallOption) (*WorkflowInvocationList, error) {
	out := new(WorkflowInvocationList)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowInvocationAPI/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowInvocationAPIClient) Get(ctx context.Context, in *fission_workflows_types1.ObjectMetadata, opts ...grpc.CallOption) (*fission_workflows_types1.WorkflowInvocation, error) {
	out := new(fission_workflows_types1.WorkflowInvocation)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowInvocationAPI/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowInvocationAPIClient) Events(ctx context.Context, in *fission_workflows_types1.ObjectMetadata, opts ...grpc.CallOption) (*ObjectEvents, error) {
	out := new(ObjectEvents)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowInvocationAPI/Events", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowInvocationAPIClient) Validate(ctx context.Context, in *fission_workflows_types1.WorkflowInvocationSpec, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowInvocationAPI/Validate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WorkflowInvocationAPI service

type WorkflowInvocationAPIServer interface {
	// Create a new workflow invocation
	//
	// In case the invocation specification is missing fields or contains invalid fields, a HTTP 400 is returned.
	Invoke(context.Context, *fission_workflows_types1.WorkflowInvocationSpec) (*fission_workflows_types1.ObjectMetadata, error)
	InvokeSync(context.Context, *fission_workflows_types1.WorkflowInvocationSpec) (*fission_workflows_types1.WorkflowInvocation, error)
	AddTask(context.Context, *AddTaskRequest) (*google_protobuf2.Empty, error)
	// Cancel a workflow invocation
	//
	// This action is irreverisble. A canceled invocation cannot be resumed or restarted.
	// In case that an invocation already is canceled, has failed or has completed, nothing happens.
	// In case that an invocation does not exist a HTTP 404 error status is returned.
	Cancel(context.Context, *fission_workflows_types1.ObjectMetadata) (*google_protobuf2.Empty, error)
	List(context.Context, *InvocationListQuery) (*WorkflowInvocationList, error)
	// Get the specification and status of a workflow invocation
	//
	// Get returns three different aspects of the workflow invocation, namely the spec (specification), status and logs.
	// To lighten the request load, consider using a more specific request.
	Get(context.Context, *fission_workflows_types1.ObjectMetadata) (*fission_workflows_types1.WorkflowInvocation, error)
	Events(context.Context, *fission_workflows_types1.ObjectMetadata) (*ObjectEvents, error)
	Validate(context.Context, *fission_workflows_types1.WorkflowInvocationSpec) (*google_protobuf2.Empty, error)
}

func RegisterWorkflowInvocationAPIServer(s *grpc.Server, srv WorkflowInvocationAPIServer) {
	s.RegisterService(&_WorkflowInvocationAPI_serviceDesc, srv)
}

func _WorkflowInvocationAPI_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fission_workflows_types1.WorkflowInvocationSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowInvocationAPI/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).Invoke(ctx, req.(*fission_workflows_types1.WorkflowInvocationSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowInvocationAPI_InvokeSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fission_workflows_types1.WorkflowInvocationSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).InvokeSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowInvocationAPI/InvokeSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).InvokeSync(ctx, req.(*fission_workflows_types1.WorkflowInvocationSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowInvocationAPI_AddTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).AddTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowInvocationAPI/AddTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).AddTask(ctx, req.(*AddTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowInvocationAPI_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fission_workflows_types1.ObjectMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowInvocationAPI/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).Cancel(ctx, req.(*fission_workflows_types1.ObjectMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowInvocationAPI_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvocationListQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowInvocationAPI/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).List(ctx, req.(*InvocationListQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowInvocationAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fission_workflows_types1.ObjectMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowInvocationAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).Get(ctx, req.(*fission_workflows_types1.ObjectMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowInvocationAPI_Events_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fission_workflows_types1.ObjectMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).Events(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowInvocationAPI/Events",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).Events(ctx, req.(*fission_workflows_types1.ObjectMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowInvocationAPI_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fission_workflows_types1.WorkflowInvocationSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowInvocationAPI/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).Validate(ctx, req.(*fission_workflows_types1.WorkflowInvocationSpec))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkflowInvocationAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fission.workflows.apiserver.WorkflowInvocationAPI",
	HandlerType: (*WorkflowInvocationAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _WorkflowInvocationAPI_Invoke_Handler,
		},
		{
			MethodName: "InvokeSync",
			Handler:    _WorkflowInvocationAPI_InvokeSync_Handler,
		},
		{
			MethodName: "AddTask",
			Handler:    _WorkflowInvocationAPI_AddTask_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _WorkflowInvocationAPI_Cancel_Handler,
		},
		{
			MethodName: "List",
			Handler:    _WorkflowInvocationAPI_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _WorkflowInvocationAPI_Get_Handler,
		},
		{
			MethodName: "Events",
			Handler:    _WorkflowInvocationAPI_Events_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _WorkflowInvocationAPI_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/apiserver/apiserver.proto",
}

// Client API for AdminAPI service

type AdminAPIClient interface {
	Status(ctx context.Context, in *google_protobuf2.Empty, opts ...grpc.CallOption) (*Health, error)
	Version(ctx context.Context, in *google_protobuf2.Empty, opts ...grpc.CallOption) (*fission_workflows_version.Info, error)
}

type adminAPIClient struct {
	cc *grpc.ClientConn
}

func NewAdminAPIClient(cc *grpc.ClientConn) AdminAPIClient {
	return &adminAPIClient{cc}
}

func (c *adminAPIClient) Status(ctx context.Context, in *google_protobuf2.Empty, opts ...grpc.CallOption) (*Health, error) {
	out := new(Health)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.AdminAPI/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminAPIClient) Version(ctx context.Context, in *google_protobuf2.Empty, opts ...grpc.CallOption) (*fission_workflows_version.Info, error) {
	out := new(fission_workflows_version.Info)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.AdminAPI/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AdminAPI service

type AdminAPIServer interface {
	Status(context.Context, *google_protobuf2.Empty) (*Health, error)
	Version(context.Context, *google_protobuf2.Empty) (*fission_workflows_version.Info, error)
}

func RegisterAdminAPIServer(s *grpc.Server, srv AdminAPIServer) {
	s.RegisterService(&_AdminAPI_serviceDesc, srv)
}

func _AdminAPI_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf2.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAPIServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.AdminAPI/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAPIServer).Status(ctx, req.(*google_protobuf2.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminAPI_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf2.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAPIServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.AdminAPI/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAPIServer).Version(ctx, req.(*google_protobuf2.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdminAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fission.workflows.apiserver.AdminAPI",
	HandlerType: (*AdminAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _AdminAPI_Status_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _AdminAPI_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/apiserver/apiserver.proto",
}

func init() { proto.RegisterFile("pkg/apiserver/apiserver.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 818 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x4e, 0xdb, 0x58,
	0x10, 0x96, 0x81, 0x35, 0xc9, 0x98, 0x45, 0xd9, 0x01, 0x42, 0x36, 0x80, 0xc8, 0x1e, 0xb4, 0xda,
	0x10, 0x76, 0xed, 0xdd, 0x20, 0xed, 0x45, 0x56, 0x5a, 0x89, 0x05, 0xb4, 0x8d, 0xd4, 0x8a, 0x36,
	0x20, 0x90, 0x50, 0x6f, 0x1c, 0xe7, 0x24, 0x71, 0x93, 0xd8, 0xa9, 0x7d, 0x12, 0x14, 0x50, 0x6e,
	0x78, 0x02, 0xa4, 0x5e, 0xf6, 0xa2, 0xd7, 0x7d, 0x84, 0x3e, 0x47, 0x5f, 0xa1, 0x0f, 0x52, 0xf9,
	0xf8, 0x38, 0x71, 0x08, 0xce, 0x8f, 0xaa, 0x5e, 0x40, 0x92, 0xf1, 0x7c, 0xf3, 0x7d, 0x33, 0x67,
	0x66, 0x8e, 0x61, 0xa7, 0xdd, 0xa8, 0x69, 0x7a, 0xdb, 0x74, 0xa9, 0xd3, 0xa5, 0xce, 0xf0, 0x9b,
	0xda, 0x76, 0x6c, 0x66, 0xe3, 0x56, 0xd5, 0x74, 0x5d, 0xd3, 0xb6, 0xd4, 0x1b, 0xdb, 0x69, 0x54,
	0x9b, 0xf6, 0x8d, 0xab, 0x0e, 0x5c, 0xd2, 0x85, 0x9a, 0xc9, 0xea, 0x9d, 0xb2, 0x6a, 0xd8, 0x2d,
	0x4d, 0xf8, 0x05, 0x9f, 0x7f, 0x0c, 0xfc, 0x35, 0x8f, 0x80, 0xf5, 0xda, 0xd4, 0xf5, 0xff, 0xfb,
	0x81, 0xd3, 0xff, 0xce, 0x8c, 0xed, 0x52, 0x87, 0x3f, 0x15, 0x9f, 0x02, 0xff, 0xf7, 0xcc, 0xf8,
	0x2a, 0x75, 0xbd, 0x3f, 0x81, 0xdb, 0xaa, 0xd9, 0x76, 0xad, 0x49, 0x35, 0xfe, 0xab, 0xdc, 0xa9,
	0x6a, 0xb4, 0xd5, 0x66, 0x3d, 0xf1, 0x70, 0x5b, 0x3c, 0xd4, 0xdb, 0xa6, 0xa6, 0x5b, 0x96, 0xcd,
	0x74, 0x66, 0xda, 0x96, 0x80, 0x92, 0xdf, 0x61, 0xe5, 0x4a, 0x44, 0x7e, 0x6e, 0xba, 0x0c, 0xb7,
	0x21, 0x3e, 0x60, 0x4a, 0x49, 0x99, 0xc5, 0x6c, 0xbc, 0x34, 0x34, 0x90, 0x1a, 0xac, 0x1e, 0x55,
	0x2a, 0x17, 0xba, 0xdb, 0x28, 0xd1, 0xb7, 0x1d, 0xea, 0x32, 0x24, 0xb0, 0x62, 0x5a, 0x5d, 0xdb,
	0xe0, 0x41, 0x8b, 0x27, 0x29, 0x29, 0x23, 0x65, 0xe3, 0xa5, 0x11, 0x1b, 0xfe, 0x05, 0x4b, 0x4c,
	0x77, 0x1b, 0xa9, 0x85, 0x8c, 0x94, 0x55, 0xf2, 0x3b, 0xea, 0x78, 0xf9, 0xfd, 0x22, 0xf2, 0xb8,
	0xdc, 0x95, 0x1c, 0xc2, 0x5a, 0x71, 0x10, 0xc2, 0x13, 0xf6, 0xaa, 0x43, 0x9d, 0xde, 0x14, 0x75,
	0x05, 0x48, 0x06, 0xb9, 0x8c, 0x82, 0x31, 0x03, 0xca, 0x50, 0x51, 0x80, 0x0c, 0x9b, 0xc8, 0x83,
	0x04, 0x2b, 0x67, 0xe5, 0x37, 0xd4, 0x60, 0xa7, 0x5d, 0x6a, 0x31, 0x17, 0x8f, 0x21, 0xd6, 0xa2,
	0x4c, 0xaf, 0xe8, 0x4c, 0xe7, 0x49, 0x29, 0xf9, 0xdf, 0x22, 0x85, 0xfb, 0xc0, 0x17, 0xc2, 0xbd,
	0x34, 0x00, 0xe2, 0x3f, 0x20, 0x53, 0x1e, 0x2e, 0xb5, 0x90, 0x59, 0xcc, 0x2a, 0xf9, 0xbd, 0x27,
	0x42, 0xf8, 0x0e, 0xcc, 0x76, 0xa8, 0xca, 0xa9, 0x4b, 0x02, 0x42, 0x32, 0x20, 0x3f, 0xa3, 0x7a,
	0x93, 0xd5, 0x31, 0x09, 0xb2, 0xcb, 0x74, 0xd6, 0x71, 0x45, 0x79, 0xc5, 0xaf, 0xfc, 0xc3, 0x0f,
	0xa0, 0x04, 0x19, 0x1f, 0xbd, 0x2c, 0xa2, 0x05, 0xf2, 0xb1, 0x43, 0x75, 0x46, 0xf1, 0xd7, 0x48,
	0xad, 0x81, 0xff, 0x79, 0x9b, 0x1a, 0xe9, 0x59, 0x53, 0x22, 0xeb, 0xf7, 0x9f, 0xbf, 0xbc, 0x5b,
	0x58, 0x25, 0x71, 0x2d, 0x70, 0x2c, 0x48, 0x39, 0x7c, 0x0d, 0x4b, 0xbc, 0xbc, 0x49, 0xd5, 0xef,
	0x31, 0x35, 0x68, 0x40, 0xf5, 0xd4, 0x6b, 0xc0, 0xf4, 0xbe, 0x3a, 0x61, 0xd2, 0xd4, 0x70, 0xdf,
	0x91, 0x9f, 0x38, 0x81, 0x82, 0x43, 0x02, 0x34, 0x61, 0xf1, 0x7f, 0xca, 0x70, 0x56, 0x8d, 0xe9,
	0x5f, 0xa6, 0xe6, 0x4c, 0x92, 0x9c, 0x25, 0x81, 0xab, 0x03, 0x16, 0xed, 0xce, 0xac, 0xf4, 0x51,
	0x07, 0xf9, 0x84, 0x36, 0x29, 0xa3, 0xb3, 0xb3, 0x45, 0xe4, 0x1c, 0x50, 0xe4, 0x1e, 0x53, 0xd4,
	0x21, 0x76, 0xa9, 0x37, 0xcd, 0xca, 0x1c, 0xa7, 0x13, 0x45, 0xb1, 0xc3, 0x29, 0x36, 0x09, 0x0e,
	0x29, 0xba, 0x22, 0xb4, 0x77, 0x2a, 0x77, 0x20, 0x8b, 0x1e, 0x9e, 0x39, 0x99, 0xc9, 0x07, 0x15,
	0x9e, 0x8b, 0x80, 0x1c, 0x37, 0x46, 0xf3, 0xd3, 0xfc, 0xa6, 0xcd, 0x7f, 0x8c, 0xc1, 0xc6, 0xf8,
	0x10, 0x7a, 0xcd, 0x79, 0x0b, 0xb2, 0x67, 0x68, 0x50, 0xd4, 0xa6, 0xa6, 0x3f, 0x44, 0xce, 0xd7,
	0xa6, 0xa2, 0xf8, 0x44, 0xd1, 0x86, 0xb3, 0xed, 0x95, 0xe4, 0xbd, 0x04, 0xe0, 0x93, 0x9f, 0xf7,
	0x2c, 0x63, 0x7e, 0x01, 0x07, 0x73, 0x00, 0x88, 0xc6, 0x45, 0xec, 0x93, 0x44, 0x48, 0x84, 0xe6,
	0xf6, 0x2c, 0xa3, 0x20, 0xe5, 0xae, 0x11, 0xc7, 0xcc, 0xf8, 0x41, 0x82, 0x65, 0xb1, 0x56, 0xf1,
	0x60, 0xe2, 0x49, 0x8c, 0x2e, 0xdf, 0xc8, 0x06, 0x39, 0xe3, 0x0a, 0x8a, 0x24, 0x13, 0xa6, 0xba,
	0x0b, 0xef, 0xe4, 0xbe, 0xe6, 0xad, 0x59, 0xd7, 0x53, 0x44, 0xd2, 0x53, 0xdd, 0xd0, 0x00, 0xf9,
	0x58, 0xb7, 0x0c, 0xda, 0xfc, 0xf6, 0xf9, 0x48, 0x71, 0x6d, 0x98, 0x4b, 0x8c, 0x92, 0x56, 0xfa,
	0x78, 0x2f, 0x89, 0x75, 0xf2, 0xe7, 0xc4, 0x1a, 0x3c, 0x71, 0x2f, 0xa4, 0x0f, 0x67, 0x5a, 0x34,
	0xa3, 0x48, 0xb2, 0xc6, 0x95, 0xfc, 0x88, 0xe1, 0x66, 0xc1, 0xce, 0x9c, 0x4b, 0x67, 0xae, 0xce,
	0x10, 0xb9, 0xe3, 0x78, 0xee, 0xfd, 0xef, 0x3a, 0xb3, 0xbb, 0x9c, 0xf7, 0x67, 0xdc, 0x7c, 0xcc,
	0x2b, 0xa6, 0x16, 0x59, 0x68, 0x39, 0xcd, 0x3d, 0x1c, 0x51, 0x27, 0x2d, 0x58, 0xc9, 0x7a, 0x98,
	0x35, 0xb4, 0xa8, 0xf2, 0x9f, 0x24, 0x88, 0x1d, 0x55, 0x5a, 0x26, 0x5f, 0x0f, 0x57, 0x20, 0x9f,
	0xf3, 0x5b, 0x2d, 0xf2, 0x36, 0xd9, 0x9b, 0x98, 0xb0, 0x7f, 0x55, 0x92, 0x04, 0x27, 0x05, 0x8c,
	0x69, 0x75, 0x6e, 0xb8, 0xc5, 0x0b, 0x58, 0xbe, 0xf4, 0xdf, 0xb2, 0x22, 0x23, 0xef, 0x3e, 0x11,
	0x39, 0x78, 0x33, 0x2b, 0x5a, 0x55, 0x3b, 0x14, 0x55, 0x98, 0xff, 0x53, 0xae, 0xe3, 0x03, 0xee,
	0xb2, 0xcc, 0xe3, 0x1d, 0x7e, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xe5, 0x32, 0xc9, 0x78, 0x0a,
	0x00, 0x00,
}
