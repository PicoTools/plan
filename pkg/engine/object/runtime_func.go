package object

import (
	"github.com/PicoTools/plan/pkg/parser"
)

// RuntimeFunc implements ObjectImpl interfaces
type RuntimeFunc struct {
	ObjectImpl
	name string
	args []string
	code []parser.IStmtContext
}

// NewRuntimeFunc creates new RuntimeFunc objects
// using arguments (args) and runtime statements array
func NewRuntimeFunc(args []string, code []parser.IStmtContext) *RuntimeFunc {
	return &RuntimeFunc{
		name: "<unknown>",
		args: args,
		code: code,
	}
}

// SetName set name of RuntimeFunc object
func (o *RuntimeFunc) SetName(name string) {
	o.name = name
}

// TypeName returns type name of runtime function and its name
func (o *RuntimeFunc) TypeName() string {
	return "native-func: " + o.name
}

// GetName returns name of runtime function
func (o *RuntimeFunc) GetName() string {
	return o.name
}

// String returns string representation of runtime function object
func (o *RuntimeFunc) String() string {
	return "<native-func>"
}

// CanCall indicates if object is callable
func (o *RuntimeFunc) CanCall() bool {
	return true
}

// GetCode returns array of runtime PLAN statements
func (o *RuntimeFunc) GetCode() []parser.IStmtContext {
	return o.code
}

// GetArgsLen returns length of arguments
func (o *RuntimeFunc) GetArgsLen() int {
	return len(o.args)
}

// GetArgs return list of arguments for RuntimeFunc object
func (o *RuntimeFunc) GetArgs() []string {
	return o.args
}
