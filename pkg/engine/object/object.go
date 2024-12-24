package object

import planerrors "github.com/PicoTools/plan/pkg/engine/errors"

type Object interface {
	// type name
	TypeName() string
	// string representation of object
	String() string
	BinaryOp(int, Object) (Object, error)
	UnaryOp(int) (Object, error)
	// is object callable
	CanCall() bool
	// call object with arguments
	Call(...Object) (Object, error)
	// call method of object
	MethodCall(string, ...Object) (Object, error)
	// get value
	GetValue() any
	// get value by index
	IndexGet(Object) (Object, error)
	// set value by index
	IndexSet(Object, Object) error
	// logical NOT
	LogicalNot() (Object, error)
	// logical OR
	LogicalOr(Object) (Object, error)
	// logical AND
	LogicalAnd(Object) (Object, error)
	// objects'es equality
	Equal(Object) (Object, error)
	// objects'es unequality
	NotEqual(Object) (Object, error)
	// unary negation
	Negative() (Object, error)
	// great than or equal
	GtEq(Object) (Object, error)
	// great than
	Gt(Object) (Object, error)
	// less than or equal
	LtEq(Object) (Object, error)
	// less than
	Lt(Object) (Object, error)
	// sum
	Add(Object) (Object, error)
	// subtractr
	Sub(Object) (Object, error)
	// pow
	Pow(Object) (Object, error)
	// multiplication
	Mul(Object) (Object, error)
	// division
	Div(Object) (Object, error)
	// obtain integer reminder
	Mod(Object) (Object, error)
	// xor
	Xor(Object) (Object, error)
}

type ObjectImpl struct{}

func (o *ObjectImpl) TypeName() string {
	panic(planerrors.ErrNotImplemented)
}

func (o *ObjectImpl) String() string {
	panic(planerrors.ErrNotImplemented)
}

func (o *ObjectImpl) BinaryOp(_ int, _ Object) (Object, error) {
	return nil, planerrors.ErrNotImplemented
}

func (o *ObjectImpl) UnaryOp(_ int) (Object, error) {
	return nil, planerrors.ErrNotImplemented
}

func (o *ObjectImpl) CanCall() bool {
	return false
}

func (o *ObjectImpl) Call(_ ...Object) (Object, error) {
	return nil, planerrors.ErrNotImplemented
}

func (o *ObjectImpl) GetValue() any {
	return nil
}

func (o *ObjectImpl) IndexGet(Object) (Object, error) {
	return nil, planerrors.ErrInvalidOp
}

func (o *ObjectImpl) IndexSet(Object, Object) error {
	return planerrors.ErrInvalidOp
}

func (o *ObjectImpl) MethodCall(string, ...Object) (Object, error) {
	return nil, planerrors.ErrUnknownMethod
}

func (a *ObjectImpl) LogicalNot() (Object, error) {
	return nil, planerrors.ErrNotImplemented
}

func (a *ObjectImpl) LogicalAnd(Object) (Object, error) {
	return nil, planerrors.ErrNotImplemented
}

func (a *ObjectImpl) LogicalOr(Object) (Object, error) {
	return nil, planerrors.ErrNotImplemented
}

func (a *ObjectImpl) Equal(Object) (Object, error) {
	return nil, planerrors.ErrNotImplemented
}

func (a *ObjectImpl) NotEqual(Object) (Object, error) {
	return nil, planerrors.ErrNotImplemented
}

func (a *ObjectImpl) Negative() (Object, error) {
	return nil, planerrors.ErrNotImplemented
}

func (a *ObjectImpl) GtEq(Object) (Object, error) {
	return nil, planerrors.ErrNotImplemented
}

func (a *ObjectImpl) Gt(Object) (Object, error) {
	return nil, planerrors.ErrNotImplemented
}

func (a *ObjectImpl) LtEq(Object) (Object, error) {
	return nil, planerrors.ErrNotImplemented
}

func (a *ObjectImpl) Lt(Object) (Object, error) {
	return nil, planerrors.ErrNotImplemented
}

func (a *ObjectImpl) Add(Object) (Object, error) {
	return nil, planerrors.ErrNotImplemented
}

func (a *ObjectImpl) Sub(Object) (Object, error) {
	return nil, planerrors.ErrNotImplemented
}

func (a *ObjectImpl) Pow(Object) (Object, error) {
	return nil, planerrors.ErrNotImplemented
}

func (a *ObjectImpl) Mul(Object) (Object, error) {
	return nil, planerrors.ErrNotImplemented
}

func (a *ObjectImpl) Div(Object) (Object, error) {
	return nil, planerrors.ErrNotImplemented
}

func (a *ObjectImpl) Mod(Object) (Object, error) {
	return nil, planerrors.ErrNotImplemented
}

func (a *ObjectImpl) Xor(Object) (Object, error) {
	return nil, planerrors.ErrNotImplemented
}
