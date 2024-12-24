package object

import (
	"math"
	"strings"

	planerrors "github.com/PicoTools/plan/pkg/engine/errors"
	"github.com/PicoTools/plan/pkg/engine/utils"
	"github.com/PicoTools/plan/pkg/parser"
)

// Bool implements ObjectImpl interfaces
type Bool struct {
	ObjectImpl
	value bool
}

// NewBool create new Bool object containing bool value
func NewBool(v bool) *Bool {
	return &Bool{value: v}
}

// TypeName returns name of Bool type
func (o *Bool) TypeName() string {
	return "bool"
}

// Strings returns string representation of Bool object
func (o *Bool) String() string {
	if o.value {
		return "true"
	} else {
		return "false"
	}
}

// GetValue returns underly boolean value
func (o *Bool) GetValue() any {
	return o.value
}

// BinaryOp provides logic of binary operations between 2 objects
func (o *Bool) BinaryOp(op int, rhs Object) (Object, error) {
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
func (o *Bool) UnaryOp(op int) (Object, error) {
	switch op {
	case parser.PLANLexerNot:
		return o.LogicalNot()
	case parser.PLANLexerSubtract:
		return o.Negative()
	}
	return nil, planerrors.ErrInvalidOp
}

// LogicalNot implements logical NOT on Bool object
func (o *Bool) LogicalNot() (Object, error) {
	o.value = !o.value
	return o, nil
}

// Negative implements unary negation of Bool object
func (o *Bool) Negative() (Object, error) {
	return NewInt(-utils.BoolToInt(o.value)), nil
}

// LogicalOr implements logical OR between Bool object and other types of objects
func (o *Bool) LogicalOr(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(o.value || rs.value), nil
	case *Null:
		return o, nil
	}
	return nil, planerrors.ErrInvalidOp
}

// LogicalAnd implements logical AND between Bool object and other types of objects
func (o *Bool) LogicalAnd(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(o.value && rs.value), nil
	case *Null:
		return rs, nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Equal implements checking of equation between Bool object and other types of objects
func (o *Bool) Equal(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(o.value == rs.value), nil
	case *Dict:
		return NewBool(false), nil
	case *Float:
		return NewBool(utils.BoolToFloat(o.value) == rs.value), nil
	case *Int:
		return NewBool(utils.BoolToInt(o.value) == rs.value), nil
	case *List:
		return NewBool(false), nil
	case *Null:
		return NewBool(false), nil
	case *Str:
		return NewBool(false), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// NotEqual implements checking of unequation between Bool object and other types of objects
func (o *Bool) NotEqual(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(o.value != rs.value), nil
	case *Dict:
		return NewBool(true), nil
	case *Float:
		return NewBool(utils.BoolToFloat(o.value) != rs.value), nil
	case *Int:
		return NewBool(utils.BoolToInt(o.value) != rs.value), nil
	case *List:
		return NewBool(true), nil
	case *Null:
		return NewBool(true), nil
	case *Str:
		return NewBool(true), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// GtEq implements 'great than or equal' checking between Bool object and other types of objects
func (o *Bool) GtEq(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(utils.BoolToInt(o.value) >= utils.BoolToInt(rs.value)), nil
	case *Float:
		return NewBool(utils.BoolToFloat(o.value) >= rs.value), nil
	case *Int:
		return NewBool(utils.BoolToInt(o.value) >= rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Gt implements 'great than' checking between Bool object and other types of objects
func (o *Bool) Gt(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(utils.BoolToInt(o.value) > utils.BoolToInt(rs.value)), nil
	case *Float:
		return NewBool(utils.BoolToFloat(o.value) > rs.value), nil
	case *Int:
		return NewBool(utils.BoolToInt(o.value) > rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// LtEq implements 'less than or equal' checking between Bool object and other types of objects
func (o *Bool) LtEq(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(utils.BoolToInt(o.value) <= utils.BoolToInt(rs.value)), nil
	case *Float:
		return NewBool(utils.BoolToFloat(o.value) <= rs.value), nil
	case *Int:
		return NewBool(utils.BoolToInt(o.value) <= rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Lt implements 'less than' checking between Bool object and other types of objects
func (o *Bool) Lt(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(utils.BoolToInt(o.value) < utils.BoolToInt(rs.value)), nil
	case *Float:
		return NewBool(utils.BoolToFloat(o.value) < rs.value), nil
	case *Int:
		return NewBool(utils.BoolToInt(o.value) < rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Add implements summation of Bool object and other types of objects
func (o *Bool) Add(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewInt(utils.BoolToInt(o.value) + utils.BoolToInt(rs.value)), nil
	case *Float:
		return NewFloat(utils.BoolToFloat(o.value) + rs.value), nil
	case *Int:
		return NewInt(utils.BoolToInt(o.value) + rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Sub implements subtraction of Bool object and other types of objects
func (o *Bool) Sub(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewInt(utils.BoolToInt(o.value) - utils.BoolToInt(rs.value)), nil
	case *Float:
		return NewFloat(utils.BoolToFloat(o.value) - rs.value), nil
	case *Int:
		return NewInt(utils.BoolToInt(o.value) - rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Pow implements exponention of Bool object by other types of objects
func (o *Bool) Pow(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewInt(int64(math.Pow(utils.BoolToFloat(o.value), utils.BoolToFloat(rs.value)))), nil
	case *Float:
		return NewFloat(math.Pow(utils.BoolToFloat(o.value), rs.value)), nil
	case *Int:
		return NewInt(int64(math.Pow(utils.BoolToFloat(o.value), utils.IntToFloat(rs.value)))), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Mul implements multiplication of Bool object and other types of objects
func (o *Bool) Mul(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewInt(utils.BoolToInt(o.value) * utils.BoolToInt(rs.value)), nil
	case *Float:
		return NewFloat(utils.BoolToFloat(o.value) * rs.value), nil
	case *Int:
		return NewInt(utils.BoolToInt(o.value) * rs.value), nil
	case *List:
		var list []Object
		var i int64
		for i = 0; i < utils.BoolToInt(o.value); i++ {
			list = append(list, rs.value...)
		}
		return NewList(list), nil
	case *Str:
		return NewStr(strings.Repeat(rs.value, int(utils.BoolToInt(o.value)))), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Div implements division of Bool object and other types of objects
func (o *Bool) Div(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		if !rs.value {
			return nil, planerrors.ErrDivByZero
		}
		return NewFloat(utils.BoolToFloat(o.value) / utils.BoolToFloat(rs.value)), nil
	case *Float:
		if rs.value == 0.0 {
			return nil, planerrors.ErrDivByZero
		}
		return NewFloat(utils.BoolToFloat(o.value) / rs.value), nil
	case *Int:
		if rs.value == 0 {
			return nil, planerrors.ErrDivByZero
		}
		return NewFloat(utils.BoolToFloat(o.value) / utils.IntToFloat(rs.value)), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Mod implements obtaining of reminder of Bool object and other types of objects
func (o *Bool) Mod(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		if !rs.value {
			return nil, planerrors.ErrModByZero
		}
		return NewInt(utils.BoolToInt(o.value) % utils.BoolToInt(rs.value)), nil
	case *Float:
		if rs.value == 0.0 {
			return nil, planerrors.ErrModByZero
		}
		return NewFloat(math.Mod(utils.BoolToFloat(o.value), rs.value)), nil
	case *Int:
		if rs.value == 0 {
			return nil, planerrors.ErrModByZero
		}
		return NewInt(utils.BoolToInt(o.value) % rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Xor implements XORing between Bool object and other types of objects
func (o *Bool) Xor(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewBool(utils.IntToBool(utils.BoolToInt(o.value) ^ utils.BoolToInt(rs.value))), nil
	case *Int:
		return NewInt(utils.BoolToInt(o.value) ^ rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}
