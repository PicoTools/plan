package object

import (
	"math"
	"strconv"

	planerrors "github.com/PicoTools/plan/pkg/engine/errors"
	"github.com/PicoTools/plan/pkg/engine/utils"
	"github.com/PicoTools/plan/pkg/parser"
)

// Float implements ObjectImpl interfaces
type Float struct {
	ObjectImpl
	value float64
}

// NewFloat create new Float object containing float64 value
func NewFloat(value float64) *Float {
	return &Float{value: value}
}

// TypeName returns name of Float type
func (o *Float) TypeName() string {
	return "float"
}

// Strings returns string representation of Float object
func (o *Float) String() string {
	return strconv.FormatFloat(o.value, 'f', -1, 64)
}

// GetValue returns underly float64 value
func (o *Float) GetValue() any {
	return o.value
}

// BinaryOp provides logic of binary operations between 2 objects
func (o *Float) BinaryOp(op int, rhs Object) (Object, error) {
	switch op {
	case parser.PLANLexerOr:
		return o.LogicalOr(rhs)
	case parser.PLANLexerAnd:
		return o.LogicalAnd(rhs)
	case parser.PLANLexerEq:
		return o.Equal(rhs)
	case parser.PLANLexerNeq:
		return o.NotEqual(rhs)
	case parser.PLANLexerGtEq:
		return o.GtEq(rhs)
	case parser.PLANLexerGt:
		return o.Gt(rhs)
	case parser.PLANLexerLtEq:
		return o.LtEq(rhs)
	case parser.PLANLexerLt:
		return o.Lt(rhs)
	case parser.PLANLexerAdd:
		return o.Add(rhs)
	case parser.PLANLexerAssSum:
		return o.Add(rhs)
	case parser.PLANLexerAssSub:
		return o.Sub(rhs)
	case parser.PLANLexerAssMul:
		return o.Mul(rhs)
	case parser.PLANLexerAssDiv:
		return o.Div(rhs)
	case parser.PLANLexerAssMod:
		return o.Mod(rhs)
	case parser.PLANLexerAssPow:
		return o.Pow(rhs)
	case parser.PLANLexerSubtract:
		return o.Sub(rhs)
	case parser.PLANLexerPow:
		return o.Pow(rhs)
	case parser.PLANLexerMultiply:
		return o.Mul(rhs)
	case parser.PLANLexerDivision:
		return o.Div(rhs)
	case parser.PLANLexerModulus:
		return o.Mod(rhs)
	}
	return nil, planerrors.ErrInvalidOp
}

// UnaryOp provides logic of unary operations on object
func (o *Float) UnaryOp(op int) (Object, error) {
	switch op {
	case parser.PLANLexerNot:
		return o.LogicalNot()
	case parser.PLANLexerSubtract:
		return o.Negative()
	}
	return nil, planerrors.ErrInvalidOp
}

// LogicalNot implements logical NOT on Float object
func (o *Float) LogicalNot() (Object, error) {
	return NewBool(!utils.FloatToBool(o.value)), nil
}

// Negative implements unary negation of Float object
func (o *Float) Negative() (Object, error) {
	o.value = -o.value
	return o, nil
}

// LogicalOr implements logical OR between Float object and other types of objects
func (o *Float) LogicalOr(rs Object) (Object, error) {
	switch rs.(type) {
	case *Null:
		return o, nil
	}
	return nil, planerrors.ErrInvalidOp
}

// LogicalAnd implements logical AND between Float object and other types of objects
func (o *Float) LogicalAnd(rs Object) (Object, error) {
	switch rs.(type) {
	case *Null:
		return rs, nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Equal implements checking of equation between Float object and other types of objects
func (o *Float) Equal(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(o.value == utils.BoolToFloat(rs.value)), nil
	case *Dict:
		return NewBool(false), nil
	case *Float:
		return NewBool(o.value == rs.value), nil
	case *Int:
		return NewBool(o.value == utils.IntToFloat(rs.value)), nil
	case *List:
		return NewBool(false), nil
	case *Null:
		return NewBool(false), nil
	case *Str:
		return NewBool(false), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// NotEqual implements checking of unequation between Float object and other types of objects
func (o *Float) NotEqual(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(o.value != utils.BoolToFloat(rs.value)), nil
	case *Dict:
		return NewBool(true), nil
	case *Float:
		return NewBool(o.value != rs.value), nil
	case *Int:
		return NewBool(o.value != utils.IntToFloat(rs.value)), nil
	case *List:
		return NewBool(true), nil
	case *Null:
		return NewBool(true), nil
	case *Str:
		return NewBool(true), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// GtEq implements 'great than or equal' checking between Float object and other types of objects
func (o *Float) GtEq(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(o.value >= utils.BoolToFloat(rs.value)), nil
	case *Float:
		return NewBool(o.value >= rs.value), nil
	case *Int:
		return NewBool(o.value >= utils.IntToFloat(rs.value)), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Gt implements 'great than' checking between Float object and other types of objects
func (o *Float) Gt(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(o.value > utils.BoolToFloat(rs.value)), nil
	case *Float:
		return NewBool(o.value > rs.value), nil
	case *Int:
		return NewBool(o.value > utils.IntToFloat(rs.value)), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// LtEq implements 'less than or equal' checking between Float object and other types of objects
func (o *Float) LtEq(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(o.value <= utils.BoolToFloat(rs.value)), nil
	case *Float:
		return NewBool(o.value <= rs.value), nil
	case *Int:
		return NewBool(o.value <= utils.IntToFloat(rs.value)), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Lt implements 'less than' checking between Float object and other types of objects
func (o *Float) Lt(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(o.value < utils.BoolToFloat(rs.value)), nil
	case *Float:
		return NewBool(o.value < rs.value), nil
	case *Int:
		return NewBool(o.value < utils.IntToFloat(rs.value)), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Add implements summation of Float object and other types of objects
func (o *Float) Add(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewFloat(o.value + utils.BoolToFloat(rs.value)), nil
	case *Float:
		return NewFloat(o.value + rs.value), nil
	case *Int:
		return NewFloat(o.value + utils.IntToFloat(rs.value)), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Sub implements subtraction of Float object and other types of objects
func (o *Float) Sub(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewFloat(o.value - utils.BoolToFloat(rs.value)), nil
	case *Float:
		return NewFloat(o.value - rs.value), nil
	case *Int:
		return NewFloat(o.value - utils.IntToFloat(rs.value)), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Pow implements exponention of Float object by other types of objects
func (o *Float) Pow(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewFloat(math.Pow(o.value, utils.BoolToFloat(rs.value))), nil
	case *Float:
		return NewFloat(math.Pow(o.value, rs.value)), nil
	case *Int:
		return NewFloat(math.Pow(o.value, utils.IntToFloat(rs.value))), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Mul implements multiplication of Float object and other types of objects
func (o *Float) Mul(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewFloat(o.value * utils.BoolToFloat(rs.value)), nil
	case *Float:
		return NewFloat(o.value * rs.value), nil
	case *Int:
		return NewFloat(o.value * utils.IntToFloat(rs.value)), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Div implements division of Float object and other types of objects
func (o *Float) Div(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		if !rs.value {
			return nil, planerrors.ErrDivByZero
		}
		return NewFloat(o.value / utils.BoolToFloat(rs.value)), nil
	case *Float:
		if rs.value == 0.0 {
			return nil, planerrors.ErrDivByZero
		}
		return NewFloat(o.value / rs.value), nil
	case *Int:
		if rs.value == 0 {
			return nil, planerrors.ErrDivByZero
		}
		return NewFloat(o.value / utils.IntToFloat(rs.value)), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Mod implements obtaining of reminder of Float object and other types of objects
func (o *Float) Mod(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		if !rs.value {
			return nil, planerrors.ErrModByZero
		}
		return NewFloat(math.Mod(o.value, utils.BoolToFloat(rs.value))), nil
	case *Float:
		if rs.value == 0.0 {
			return nil, planerrors.ErrModByZero
		}
		return NewFloat(math.Mod(o.value, rs.value)), nil
	case *Int:
		if rs.value == 0 {
			return nil, planerrors.ErrModByZero
		}
		return NewFloat(math.Mod(o.value, utils.IntToFloat(rs.value))), nil
	}
	return nil, planerrors.ErrInvalidOp
}
