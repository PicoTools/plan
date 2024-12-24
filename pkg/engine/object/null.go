package object

import (
	planerrors "github.com/PicoTools/plan/pkg/engine/errors"
	"github.com/PicoTools/plan/pkg/parser"
)

// Null implements ObjectImpl interface
type Null struct {
	ObjectImpl
}

// NewNull creates new Null object
func NewNull() *Null {
	return &Null{}
}

// TypeName returns type name
func (o *Null) TypeName() string {
	return "null"
}

// String returnes string representation of Null object
func (o *Null) String() string {
	return "<null>"
}

// BinaryOp provides logic of binary operations between 2 objects
func (o *Null) BinaryOp(op int, rhs Object) (Object, error) {
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

// UnaryOp provides logic of unary operations on object
func (o *Null) UnaryOp(op int) (Object, error) {
	switch op {
	case parser.PLANLexerNot:
		return o.LogicalNot()
	}
	return nil, planerrors.ErrInvalidOp
}

// LogicalNot implements logical NOT on Null object
func (o *Null) LogicalNot() (Object, error) {
	return NewBool(true), nil
}

// LogicalOr implements logical OR between Null object and other types of objects
func (o *Null) LogicalOr(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return rs, nil
	case *Dict:
		return rs, nil
	case *Float:
		return rs, nil
	case *Int:
		return rs, nil
	case *List:
		return rs, nil
	case *Null:
		return rs, nil
	case *Str:
		return rs, nil
	}
	return nil, planerrors.ErrInvalidOp
}

// LogicalAnd implements logical AND between Null object and other types of objects
func (o *Null) LogicalAnd(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return o, nil
	case *Dict:
		return o, nil
	case *Float:
		return o, nil
	case *Int:
		return o, nil
	case *List:
		return o, nil
	case *Null:
		return o, nil
	case *Str:
		return o, nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Equal implements checking of equation between Null object and other types of objects
func (o *Null) Equal(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(false), nil
	case *Dict:
		return NewBool(false), nil
	case *Float:
		return NewBool(false), nil
	case *Int:
		return NewBool(false), nil
	case *List:
		return NewBool(false), nil
	case *Null:
		return NewBool(true), nil
	case *Str:
		return NewBool(false), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// NotEqual implements checking of unequation between Null object and other types of objects
func (o *Null) NotEqual(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(true), nil
	case *Dict:
		return NewBool(true), nil
	case *Float:
		return NewBool(true), nil
	case *Int:
		return NewBool(true), nil
	case *List:
		return NewBool(true), nil
	case *Null:
		return NewBool(false), nil
	case *Str:
		return NewBool(true), nil
	}
	return nil, planerrors.ErrInvalidOp
}
