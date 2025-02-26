package object

import (
	"fmt"
	"strings"

	planerrors "github.com/PicoTools/plan/pkg/engine/errors"
	"github.com/PicoTools/plan/pkg/engine/utils"
	"github.com/PicoTools/plan/pkg/parser"
)

// List implements ObjectImpl interfaces
type List struct {
	ObjectImpl
	value   []Object
	methods map[string]*NativeFunc
}

// NewList create new List object containing list of objects
func NewList(v []Object) *List {
	l := &List{value: v}
	l.fillMethods()
	return l
}

// fillMethods fills List object's methods in storage
func (o *List) fillMethods() {
	o.methods = make(map[string]*NativeFunc)
	o.methods["len"] = NewNativeFunc("len", o.MethodLen)
	o.methods["reverse"] = NewNativeFunc("reverse", o.MethodReverse)
	o.methods["pop"] = NewNativeFunc("pop", o.MethodPop)
	o.methods["append"] = NewNativeFunc("append", o.MethodAppend)
}

// TypeName returns name of List type
func (o *List) TypeName() string {
	return "list"
}

// Strings returns string representation of List object
func (o *List) String() string {
	var items []string
	for _, item := range o.value {
		items = append(items, item.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(items, ", "))
}

// GetValue returns underly value as list of objects
func (o *List) GetValue() any {
	return o.value
}

// BinaryOp provides logic of binary operations between 2 objects
func (o *List) BinaryOp(op int, rhs Object) (Object, error) {
	switch op {
	case parser.PLANLexerOr:
		return o.LogicalOr(rhs)
	case parser.PLANLexerAnd:
		return o.LogicalAnd(rhs)
	case parser.PLANLexerEq:
		return o.Equal(rhs)
	case parser.PLANLexerNeq:
		return o.NotEqual(rhs)
	case parser.PLANLexerAdd:
		return o.Add(rhs)
	case parser.PLANLexerAssSum:
		return o.Add(rhs)
	case parser.PLANLexerAssMul:
		return o.Mul(rhs)
	case parser.PLANLexerMultiply:
		return o.Mul(rhs)
	}
	return nil, planerrors.ErrInvalidOp
}

// UnaryOp provides logic of unary operations on object
func (o *List) UnaryOp(op int) (Object, error) {
	switch op {
	case parser.PLANLexerNot:
		return o.LogicalNot()
	}
	return nil, planerrors.ErrInvalidOp
}

// IndexGet returns value from List by index
func (o *List) IndexGet(rs Object) (Object, error) {
	idx, ok := rs.(*Int)
	if !ok {
		return nil, planerrors.ErrInvalidIndexType
	}
	if idx.value < 0 || int(idx.value) >= len(o.value) {
		return nil, planerrors.ErrIndexOutOfRange
	}
	return o.value[idx.value], nil
}

// IndexSet sets value in List by index
func (o *List) IndexSet(idx Object, rs Object) error {
	idxInt, ok := idx.(*Int)
	if !ok {
		return planerrors.ErrInvalidIndexType
	}
	if idxInt.value < 0 || int(idxInt.value) >= len(o.value) {
		return planerrors.ErrIndexOutOfRange
	}
	o.value[idxInt.value] = rs
	return nil
}

// LogicalNot implements logical NOT on List object
func (o *List) LogicalNot() (Object, error) {
	if len(o.value) == 0 {
		return NewBool(true), nil
	}
	return NewBool(false), nil
}

// LogicalOr implements logical OR between List object and other types of objects
func (o *List) LogicalOr(rs Object) (Object, error) {
	switch rs.(type) {
	case *Null:
		return o, nil
	}
	return nil, planerrors.ErrInvalidOp
}

// LogicalAnd implements logical AND between List object and other types of objects
func (o *List) LogicalAnd(rs Object) (Object, error) {
	switch rs.(type) {
	case *Null:
		return rs, nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Equal implements checking of equation between List object and other types of objects
func (o *List) Equal(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(false), nil
	case *Dict:
		return NewBool(false), nil
	case *Float:
		return NewBool(false), nil
	case *Int:
		return NewBool(false), nil
	case *List:
		if len(o.value) != len(rs.value) {
			return NewBool(false), nil
		}
		for i, v := range o.value {
			val, err := v.Equal(rs.value[i])
			if err != nil {
				return nil, err
			}
			if !val.GetValue().(bool) {
				return NewBool(false), nil
			}
		}
		return NewBool(true), nil
	case *Null:
		return NewBool(false), nil
	case *Str:
		return NewBool(false), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// NotEqual implements checking of unequation between List object and other types of objects
func (o *List) NotEqual(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(true), nil
	case *Dict:
		return NewBool(true), nil
	case *Float:
		return NewBool(true), nil
	case *Int:
		return NewBool(true), nil
	case *List:
		if len(o.value) != len(rs.value) {
			return NewBool(true), nil
		}
		for i, v := range o.value {
			val, err := v.Equal(rs.value[i])
			if err != nil {
				return nil, err
			}
			if !val.GetValue().(bool) {
				return NewBool(true), nil
			}
		}
		return NewBool(false), nil
	case *Null:
		return NewBool(true), nil
	case *Str:
		return NewBool(true), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Add implements summation of List object and other types of objects
func (o *List) Add(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		o.value = append(o.value, rs)
		return o, nil
	case *Dict:
		o.value = append(o.value, rs)
		return o, nil
	case *Float:
		o.value = append(o.value, rs)
		return o, nil
	case *Int:
		o.value = append(o.value, rs)
		return o, nil
	case *List:
		o.value = append(o.value, rs)
		return o, nil
	case *Null:
		o.value = append(o.value, rs)
		return o, nil
	case *Str:
		o.value = append(o.value, rs)
		return o, nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Mul implements multiplication of List object and other types of objects
func (o *List) Mul(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		var list []Object
		var i int64
		for i = 0; i < utils.BoolToInt(rs.value); i++ {
			list = append(list, o.value...)
		}
		return NewList(list), nil
	case *Int:
		var list []Object
		var i int64
		for i = 0; i < rs.value; i++ {
			list = append(list, o.value...)
		}
		return NewList(list), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// MethodCall dispatch method from storage for List object and executes it
func (o *List) MethodCall(name string, args ...Object) (Object, error) {
	m := o.methods[name]
	if m == nil {
		return nil, planerrors.ErrUnknownMethod
	}
	return m.Call(args...)
}

// MethodLen returns len of List object via method
func (o *List) MethodLen(args ...Object) (Object, error) {
	if len(args) > 0 {
		return nil, fmt.Errorf("expecting 0 arguments, got %d", len(args))
	}
	return NewInt(int64(len(o.value))), nil
}

// MethodReverse reverses list objects via method
func (o *List) MethodReverse(args ...Object) (Object, error) {
	if len(args) > 0 {
		return nil, fmt.Errorf("expecting 0 arguments, got %d", len(args))
	}
	var temp []Object
	for i := len(o.value) - 1; i >= 0; i-- {
		temp = append(temp, o.value[i])
	}
	o.value = temp
	return NewNull(), nil
}

// MethodPop remove value from List object via method
func (o *List) MethodPop(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 arguments, got %d", len(args))
	}
	idx, ok := args[0].(*Int)
	if !ok {
		return nil, fmt.Errorf("expecting int as 1st argument, got '%s'", args[1].TypeName())
	}
	if idx.GetValue().(int64) < 0 || int(idx.GetValue().(int64)) >= len(o.value) {
		return nil, planerrors.ErrIndexOutOfRange
	}
	o.value = append(o.value[:idx.GetValue().(int64)], o.value[idx.GetValue().(int64)+1:]...)
	return NewNull(), nil
}

// MethodPush append objects at the end of the list
func (o *List) MethodAppend(args ...Object) (Object, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("expecting 1 or more arguments, got %d", len(args))
	}
	o.value = append(o.value, args...)
	return NewNull(), nil
}
