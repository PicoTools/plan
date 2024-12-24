package object

import (
	"fmt"
	"strings"

	planerrors "github.com/PicoTools/plan/pkg/engine/errors"
	"github.com/PicoTools/plan/pkg/parser"
)

// Dict implements ObjectImpl interfaces
type Dict struct {
	ObjectImpl
	value   map[string]Object
	methods map[string]*NativeFunc
}

// NewBool create new Dict object containing map of values
func NewDict(v map[string]Object) *Dict {
	d := &Dict{value: v}
	d.fillMethods()
	return d
}

// fillMethods fills Dict object's methods in storage
func (o *Dict) fillMethods() {
	o.methods = make(map[string]*NativeFunc)
	o.methods["len"] = NewNativeFunc("len", o.MethodLen)
	o.methods["pop"] = NewNativeFunc("pop", o.MethodPop)
}

// TypeName returns name of Dict type
func (o *Dict) TypeName() string {
	return "dict"
}

// Strings returns string representation of Dict object
func (o *Dict) String() string {
	var items []string
	for k, v := range o.value {
		items = append(items, fmt.Sprintf("%s: %s", k, v.String()))
	}
	return fmt.Sprintf("{%s}", strings.Join(items, ", "))
}

// GetValue returns underly values of Dict object
func (o *Dict) GetValue() any {
	return o.value
}

// BinaryOp provides logic of binary operations between 2 objects
func (o *Dict) BinaryOp(op int, rhs Object) (Object, error) {
	switch op {
	case parser.PLANLexerOr:
		return o.LogicalOr(rhs)
	case parser.PLANLexerAnd:
		return o.LogicalAnd(rhs)
	case parser.PLANLexerEq:
		return o.Equal(rhs)
	case parser.PLANLexerNeq:
		return o.NotEqual(rhs)
	}
	return nil, planerrors.ErrInvalidOp
}

// IndexGet returns value from Dict by index
func (o *Dict) IndexGet(rs Object) (Object, error) {
	idx, ok := rs.(*Str)
	if !ok {
		return nil, planerrors.ErrInvalidIndexType
	}
	res, ok := o.value[idx.value]
	if !ok {
		return NewNull(), nil
	} else {
		return res, nil
	}
}

// IndexSet sets value in Dict by index
func (o *Dict) IndexSet(idx Object, rs Object) error {
	idxStr, ok := rs.(*Str)
	if !ok {
		return planerrors.ErrInvalidIndexType
	}
	o.value[idxStr.value] = rs
	return nil
}

// UnaryOp provides logic of unary operations on object
func (o *Dict) UnaryOp(op int) (Object, error) {
	switch op {
	case parser.PLANLexerNot:
		return o.LogicalNot()
	}
	return nil, planerrors.ErrInvalidOp
}

// LogicalNot implements logical NOT on Dict object
func (o *Dict) LogicalNot() (Object, error) {
	if len(o.value) == 0 {
		return NewBool(true), nil
	}
	return NewBool(false), nil
}

// LogicalOr implements logical OR between Dict object and other types of objects
func (o *Dict) LogicalOr(rs Object) (Object, error) {
	switch rs.(type) {
	case *Null:
		return o, nil
	}
	return nil, planerrors.ErrInvalidOp
}

// LogicalAnd implements logical AND between Dict object and other types of objects
func (o *Dict) LogicalAnd(rs Object) (Object, error) {
	switch rs.(type) {
	case *Null:
		return rs, nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Equal implements checking of equation between Dict object and other types of objects
func (o *Dict) Equal(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(false), nil
	case *Dict:
		if len(o.value) != len(rs.(*Dict).value) {
			return NewBool(false), nil
		}
		for k, v := range o.value {
			if (rs.(*Dict).value)[k] == nil {
				return NewBool(false), nil
			}
			val, err := v.Equal(rs.(*Dict).value[k])
			if err != nil {
				return nil, err
			}
			if !val.(*Bool).value {
				return NewBool(false), nil
			}
		}
		return NewBool(true), nil
	case *Float:
		return NewBool(false), nil
	case *Int:
		return NewBool(false), nil
	case *List:
		return NewBool(false), nil
	case *Null:
		return NewBool(false), nil
	case *Str:
		return NewBool(false), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// NotEqual implements checking of unequation between Dict object and other types of objects
func (o *Dict) NotEqual(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(true), nil
	case *Dict:
		if len(o.value) != len(rs.(*Dict).value) {
			return NewBool(true), nil
		}
		for k, v := range o.value {
			if (rs.(*Dict).value)[k] == nil {
				return NewBool(true), nil
			}
			val, err := v.Equal(rs.(*Dict).value[k])
			if err != nil {
				return nil, err
			}
			if !val.(*Bool).value {
				return NewBool(true), nil
			}
		}
		return NewBool(false), nil
	case *Float:
		return NewBool(true), nil
	case *Int:
		return NewBool(true), nil
	case *List:
		return NewBool(true), nil
	case *Null:
		return NewBool(true), nil
	case *Str:
		return NewBool(true), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// MethodCall dispatch method from storage for Dict object and executes it
func (o *Dict) MethodCall(name string, args ...Object) (Object, error) {
	m := o.methods[name]
	if m == nil {
		return nil, planerrors.ErrUnknownMethod
	}
	return m.Call(args...)
}

// MethodLen returns len of Dict object via method
func (o *Dict) MethodLen(args ...Object) (Object, error) {
	if len(args) > 0 {
		return nil, fmt.Errorf("expecting 0 arguments, got %d", len(args))
	}
	return NewInt(int64(len(o.value))), nil
}

// MethodPop remove key from Dict object via method
func (o *Dict) MethodPop(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 arguments, got %d", len(args))
	}
	key, ok := args[0].(*Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[1].TypeName())
	}
	delete(o.value, key.GetValue().(string))
	return NewNull(), nil
}
