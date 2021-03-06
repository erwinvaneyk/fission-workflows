package expr

import (
	"errors"
	"fmt"
	"time"

	"github.com/fatih/structs"
	"github.com/fission/fission-workflows/pkg/types/typedvalues"
	"github.com/fission/fission-workflows/pkg/util"
	"github.com/robertkrimen/otto"
	"github.com/sirupsen/logrus"

	// Import the underscore library for the Otto JavaScript engine.
	_ "github.com/robertkrimen/otto/underscore"
)

const (
	varScope         = "$"
	varCurrentTask   = "taskId"
	ResolvingTimeout = time.Duration(100) * time.Millisecond
)

var (
	ErrTimeOut      = errors.New("expression resolver timed out")
	DefaultResolver = NewJavascriptExpressionParser()
)

func Resolve(rootScope interface{}, currentTask string, expr *typedvalues.TypedValue) (*typedvalues.TypedValue, error) {
	return DefaultResolver.Resolve(rootScope, currentTask, expr)
}

// resolver resolves an expression within a given context/scope.
type Resolver interface {
	Resolve(rootScope interface{}, currentTask string, expr *typedvalues.TypedValue) (*typedvalues.TypedValue, error)
}

// Function is an interface for providing functions that are able to be injected into the Otto runtime.
type Function interface {
	Apply(vm *otto.Otto, call otto.FunctionCall) otto.Value
}

type JavascriptExpressionParser struct {
	vm *otto.Otto
}

func NewJavascriptExpressionParser() *JavascriptExpressionParser {
	vm := otto.New()

	// Load expression functions into Otto
	return &JavascriptExpressionParser{
		vm: vm,
	}
}

func (oe *JavascriptExpressionParser) Resolve(rootScope interface{}, currentTask string,
	expr *typedvalues.TypedValue) (*typedvalues.TypedValue, error) {

	switch expr.ValueType() {
	case typedvalues.TypeList:
		return oe.resolveList(rootScope, currentTask, expr)
	case typedvalues.TypeMap:
		return oe.resolveMap(rootScope, currentTask, expr)
	case typedvalues.TypeExpression:
		return oe.resolveExpr(rootScope, currentTask, expr)
	default:
		return expr, nil
	}
}

func (oe *JavascriptExpressionParser) resolveExpr(rootScope interface{}, currentTask string,
	expr *typedvalues.TypedValue) (*typedvalues.TypedValue, error) {

	if expr.ValueType() != typedvalues.TypeExpression {
		return nil, errors.New("expected expression to resolve")
	}

	defer func() {
		if caught := recover(); caught != nil {
			if ErrTimeOut != caught {
				panic(caught)
			}
		}
	}()

	// Setup the JavaScript interpreter
	scoped := oe.vm.Copy()
	injectFunctions(scoped, BuiltinFunctions)
	err := scoped.Set(varScope, rootScope)
	if err != nil {
		return nil, err
	}
	err = scoped.Set(varCurrentTask, currentTask)
	if err != nil {
		return nil, err
	}

	go func() {
		<-time.After(ResolvingTimeout)
		select {
		case scoped.Interrupt <- func() {
			panic(ErrTimeOut)
		}:
		default:
			// evaluation has already been interrupted / quit
		}
	}()

	e, err := typedvalues.UnwrapExpression(expr)
	if err != nil {
		return nil, fmt.Errorf("failed to format expression for resolving (%v)", err)
	}
	cleanExpr := typedvalues.RemoveExpressionDelimiters(e)
	jsResult, err := scoped.Run(cleanExpr)
	if err != nil {
		return nil, err
	}

	i, _ := jsResult.Export() // Err is always nil
	if structs.IsStruct(i) {
		mp, err := util.ConvertStructsToMap(i)
		if err != nil {
			return nil, err
		}
		i = mp
	}

	result, err := typedvalues.Wrap(i)
	if err != nil {
		return nil, err
	}
	result.SetMetadata("src", e)
	return result, nil
}

func (oe *JavascriptExpressionParser) resolveMap(rootScope interface{}, currentTask string,
	expr *typedvalues.TypedValue) (*typedvalues.TypedValue, error) {

	if expr.ValueType() != typedvalues.TypeMap {
		return nil, errors.New("expected map to resolve")
	}

	logrus.WithField("expr", expr).Debug("Resolving map")
	i, err := typedvalues.Unwrap(expr)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}
	obj := i.(map[string]interface{})
	for k, v := range obj { // TODO add priority here
		field, err := typedvalues.Wrap(v)
		if err != nil {
			return nil, err
		}

		resolved, err := oe.Resolve(rootScope, currentTask, field)
		if err != nil {
			return nil, err
		}

		actualVal, err := typedvalues.Unwrap(resolved)
		if err != nil {
			return nil, err
		}
		result[k] = actualVal
	}
	return typedvalues.Wrap(result)
}

func (oe *JavascriptExpressionParser) resolveList(rootScope interface{}, currentTask string,
	expr *typedvalues.TypedValue) (*typedvalues.TypedValue, error) {

	if expr.ValueType() != typedvalues.TypeList {
		return nil, errors.New("expected list to resolve")
	}

	logrus.WithField("expr", expr).Debug("Resolving list")
	i, err := typedvalues.Unwrap(expr)
	if err != nil {
		return nil, err
	}

	result := []interface{}{}
	obj := i.([]interface{})
	for _, v := range obj { // TODO add priority here
		field, err := typedvalues.Wrap(v)
		if err != nil {
			return nil, err
		}

		resolved, err := oe.Resolve(rootScope, currentTask, field)
		if err != nil {
			return nil, err
		}

		actualVal, err := typedvalues.Unwrap(resolved)
		if err != nil {
			return nil, err
		}
		result = append(result, actualVal)
	}
	return typedvalues.Wrap(result)
}

func injectFunctions(vm *otto.Otto, fns map[string]Function) {
	for varName := range fns {
		func(fnName string) {
			err := vm.Set(fnName, func(call otto.FunctionCall) otto.Value {
				return fns[fnName].Apply(vm, call)
			})
			if err != nil {
				panic(err)
			}
		}(varName)
	}
}
