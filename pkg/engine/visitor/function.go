package visitor

import (
	"fmt"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/PicoTools/plan/pkg/engine/scope"
	"github.com/PicoTools/plan/pkg/engine/storage"
	"github.com/PicoTools/plan/pkg/engine/types"
)

// invokeFunc invokes function from storage
func (v *Visitor) invokeFunc(name string, params ...object.Object) any {
	// get function object from storage
	val := storage.GetFunction(name)
	if val == nil {
		v.SetError(fmt.Errorf("undefined function '%s", name))
		return types.Failure
	}

	// check if object callable
	if !val.CanCall() {
		v.SetError(fmt.Errorf("uncallable function '%s'", name))
		return types.Failure
	}

	// create new context to execute function in
	scope.CurrentScope = scope.NewScope(
		scope.CurrentScope,
		scope.CurrentScope.Depth()+1,
		true,
		false,
		make(map[string]object.Object),
	)
	defer func() {
		scope.CurrentScope = scope.CurrentScope.Parent()
	}()

	// execute runtime function in special way
	if fn, ok := val.(*object.RuntimeFunc); ok {
		return v.invokeRuntimeFunc(fn, false, params...)
	}

	// execute builtin/user function
	res, err := val.Call(params...)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	return res
}

// invokeClosure invokes closure function
func (v *Visitor) invokeClosure(name string, params ...object.Object) any {
	// get function name from scope
	fn := scope.CurrentScope.Get(name, true)
	if fn == nil {
		v.SetError(fmt.Errorf("undefined closure '%s'", name))
		return types.Failure
	}

	// check if object is callable
	if !fn.CanCall() {
		v.SetError(fmt.Errorf("uncallable closure '%s'", name))
		return types.Failure
	}

	// create new context to execute function in
	scope.CurrentScope = scope.NewScope(
		scope.CurrentScope,
		scope.CurrentScope.Depth()+1,
		true,
		false,
		make(map[string]object.Object),
	)
	defer func() {
		scope.CurrentScope = scope.CurrentScope.Parent()
	}()

	// execute runtime function in special way
	if fn, ok := fn.(*object.RuntimeFunc); ok {
		return v.invokeRuntimeFunc(fn, true, params...)
	}

	v.SetError(fmt.Errorf("unexpected type '%s' of function for closure '%s'", fn.TypeName(), name))
	return types.Failure
}

// invokeRuntimeFunc invokes runtime function
func (v *Visitor) invokeRuntimeFunc(fn *object.RuntimeFunc, isClosure bool, params ...object.Object) any {
	// check if mislead number of arguments
	if fn.GetArgsLen() != len(params) {
		t := "function"
		if isClosure {
			t = "closure"
		}
		v.SetError(fmt.Errorf("%s '%s' expected %d arguments, got %d", t, fn.GetName(), fn.GetArgsLen(), len(params)))
		return types.Failure
	}
	// save arguments in scope
	args := fn.GetArgs()
	for i, arg := range params {
		scope.CurrentScope.Put(args[i], arg)
	}
	// execute statements in function body
	for _, item := range fn.GetCode() {
		if ok := v.Visit(item).(types.VisitResultType); !ok {
			return types.Failure
		}
		if retValue != nil {
			x := retValue
			ClearRet()
			return x
		}
	}
	return types.Success
}

// InvokeRuntimeFunc invokes directly runtime function with context
func (v *Visitor) InvokeRuntimeFunc(fn *object.RuntimeFunc, params ...object.Object) any {
	// create new context to execute function in
	scope.CurrentScope = scope.NewScope(
		scope.CurrentScope,
		scope.CurrentScope.Depth()+1,
		true,
		false,
		make(map[string]object.Object),
	)
	defer func() {
		scope.CurrentScope = scope.CurrentScope.Parent()
	}()

	return v.invokeRuntimeFunc(fn, false, params...)
}
