package object

// NativeFunc implements ObjectImpl interfaces
type NativeFunc struct {
	ObjectImpl
	name string
	code func(args ...Object) (Object, error)
}

// NewNativeFunc creates new NativeFunc objects with
// name of function and native Golang function
func NewNativeFunc(n string, c func(args ...Object) (Object, error)) *NativeFunc {
	return &NativeFunc{
		name: n,
		code: c,
	}
}

// TypeName returns type name of native function and its name
func (o *NativeFunc) TypeName() string {
	return "native-func: " + o.name
}

// String returns string representation of native function object
func (o *NativeFunc) String() string {
	return "<native-func>"
}

// CanCall indicates if object is callable
func (o *NativeFunc) CanCall() bool {
	return true
}

// Call iinvoke native function using underlied Golang realization
func (o *NativeFunc) Call(args ...Object) (Object, error) {
	return o.code(args...)
}
