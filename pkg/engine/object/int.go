package object

import (
	"math"
	"strconv"
	"strings"

	planerrors "github.com/PicoTools/plan/pkg/engine/errors"
	"github.com/PicoTools/plan/pkg/engine/utils"
	"github.com/PicoTools/plan/pkg/parser"
)

// Int implements ObjectImpl interfaces
type Int struct {
	ObjectImpl
	value int64
}

// NewInt create new Int object containing int64 value
func NewInt(value int64) *Int {
	return &Int{
		value: value,
	}
}

// TypeName returns name of Int type
func (o *Int) TypeName() string {
	return "int"
}

// Strings returns string representation of Int object
func (o *Int) String() string {
	return strconv.FormatInt(o.value, 10)
}

// GetValue returns underly int64 value
func (o *Int) GetValue() any {
	return o.value
}

// BinaryOp provides logic of binary operations between 2 objects
func (o *Int) BinaryOp(op int, rhs Object) (Object, error) {
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
	case parser.PLANLexerXor:
		return o.Xor(rhs)
	}
	return nil, planerrors.ErrInvalidOp
}

// UnaryOp provides logic of unary operations on object
func (o *Int) UnaryOp(op int) (Object, error) {
	switch op {
	case parser.PLANLexerNot:
		return o.LogicalNot()
	case parser.PLANLexerSubtract:
		return o.Negative()
	}
	return nil, planerrors.ErrInvalidOp
}

// LogicalNot implements logical NOT on Int object
func (o *Int) LogicalNot() (Object, error) {
	return NewBool(!utils.IntToBool(o.value)), nil
}

// Negative implements unary negation of Int object
func (o *Int) Negative() (Object, error) {
	o.value = -o.value
	return o, nil
}

// LogicalOr implements logical OR between Int object and other types of objects
func (o *Int) LogicalOr(rs Object) (Object, error) {
	switch rs.(type) {
	case *Null:
		return o, nil
	}
	return nil, planerrors.ErrInvalidOp
}

// LogicalAnd implements logical AND between Int object and other types of objects
func (o *Int) LogicalAnd(rs Object) (Object, error) {
	switch rs.(type) {
	case *Null:
		return rs, nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Equal implements checking of equation between Int object and other types of objects
func (o *Int) Equal(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(o.value == utils.BoolToInt(rs.value)), nil
	case *Dict:
		return NewBool(false), nil
	case *Float:
		return NewBool(utils.IntToFloat(o.value) == rs.value), nil
	case *Int:
		return NewBool(o.value == rs.value), nil
	case *List:
		return NewBool(false), nil
	case *Null:
		return NewBool(false), nil
	case *Str:
		return NewBool(false), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// NotEqual implements checking of unequation between Int object and other types of objects
func (o *Int) NotEqual(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(o.value != utils.BoolToInt(rs.value)), nil
	case *Dict:
		return NewBool(true), nil
	case *Float:
		return NewBool(utils.IntToFloat(o.value) != rs.value), nil
	case *Int:
		return NewBool(o.value != rs.value), nil
	case *List:
		return NewBool(true), nil
	case *Null:
		return NewBool(true), nil
	case *Str:
		return NewBool(true), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// GtEq implements 'great than or equal' checking between Int object and other types of objects
func (o *Int) GtEq(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(o.value >= utils.BoolToInt(rs.value)), nil
	case *Float:
		return NewBool(utils.IntToFloat(o.value) >= rs.value), nil
	case *Int:
		return NewBool(o.value >= rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Gt implements 'great than' checking between Int object and other types of objects
func (o *Int) Gt(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(o.value > utils.BoolToInt(rs.value)), nil
	case *Float:
		return NewBool(utils.IntToFloat(o.value) > rs.value), nil
	case *Int:
		return NewBool(o.value > rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// LtEq implements 'less than or equal' checking between Int object and other types of objects
func (o *Int) LtEq(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(o.value <= utils.BoolToInt(rs.value)), nil
	case *Float:
		return NewBool(utils.IntToFloat(o.value) <= rs.value), nil
	case *Int:
		return NewBool(o.value <= rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Lt implements 'less than' checking between Int object and other types of objects
func (o *Int) Lt(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(o.value < utils.BoolToInt(rs.value)), nil
	case *Float:
		return NewBool(utils.IntToFloat(o.value) < rs.value), nil
	case *Int:
		return NewBool(o.value < rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Add implements summation of Int object and other types of objects
func (o *Int) Add(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewInt(o.value + utils.BoolToInt(rs.value)), nil
	case *Float:
		return NewFloat(utils.IntToFloat(o.value) + rs.value), nil
	case *Int:
		return NewInt(o.value + rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Sub implements subtraction of Int object and other types of objects
func (o *Int) Sub(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewInt(o.value - utils.BoolToInt(rs.value)), nil
	case *Float:
		return NewFloat(utils.IntToFloat(o.value) - rs.value), nil
	case *Int:
		return NewInt(o.value - rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Pow implements exponention of Int object by other types of objects
func (o *Int) Pow(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewInt(int64(math.Pow(utils.IntToFloat(o.value), utils.BoolToFloat(rs.value)))), nil
	case *Float:
		return NewFloat(math.Pow(utils.IntToFloat(o.value), rs.value)), nil
	case *Int:
		return NewInt(int64(math.Pow(utils.IntToFloat(o.value), utils.IntToFloat(rs.value)))), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Mul implements multiplication of Int object and other types of objects
func (o *Int) Mul(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewInt(o.value * utils.BoolToInt(rs.value)), nil
	case *Float:
		return NewFloat(utils.IntToFloat(o.value) * rs.value), nil
	case *Int:
		return NewInt(o.value * rs.value), nil
	case *List:
		var list []Object
		var i int64
		for i = 0; i < o.value; i++ {
			list = append(list, rs.value...)
		}
		return NewList(list), nil
	case *Str:
		return NewStr(strings.Repeat(rs.value, int(o.value))), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Div implements division of Int object and other types of objects
func (o *Int) Div(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		if !rs.value {
			return nil, planerrors.ErrDivByZero
		}
		return NewFloat(utils.IntToFloat(o.value) / utils.BoolToFloat(rs.value)), nil
	case *Float:
		if rs.value == 0.0 {
			return nil, planerrors.ErrDivByZero
		}
		return NewFloat(utils.IntToFloat(o.value) / rs.value), nil
	case *Int:
		if rs.value == 0 {
			return nil, planerrors.ErrDivByZero
		}
		return NewInt(o.value / rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Mod implements obtaining of reminder of Int object and other types of objects
func (o *Int) Mod(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		if !rs.value {
			return nil, planerrors.ErrModByZero
		}
		return NewInt(o.value % utils.BoolToInt(rs.value)), nil
	case *Float:
		if rs.value == 0.0 {
			return nil, planerrors.ErrModByZero
		}
		return NewFloat(math.Mod(utils.IntToFloat(o.value), rs.value)), nil
	case *Int:
		if rs.value == 0 {
			return nil, planerrors.ErrModByZero
		}
		return NewInt(o.value % rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Xor implements XORing between Int object and other types of objects
func (o *Int) Xor(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewInt(o.value ^ utils.BoolToInt(rs.value)), nil
	case *Int:
		return NewInt(o.value ^ rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}
