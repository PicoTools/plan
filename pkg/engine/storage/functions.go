package storage

import "github.com/PicoTools/plan/pkg/engine/object"

// builtin functions (0 priority)
var BuiltinFunctions = make(map[string]*object.NativeFunc)

// user functions implemented in Golang (1 priority)
var UserFunctions = make(map[string]*object.NativeFunc)

// runtime native PLAN functions (2 priority)
var NativeFunctions = make(map[string]*object.RuntimeFunc)

// GetFunction searches function in storage
func GetFunction(name string) object.Object {
	if val, ok := BuiltinFunctions[name]; ok {
		return val
	}

	if val, ok := UserFunctions[name]; ok {
		return val
	}

	if val, ok := NativeFunctions[name]; ok {
		return val
	}

	return nil
}
