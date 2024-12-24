package object

import (
	"fmt"
	"strings"
	"unicode/utf8"

	planerrors "github.com/PicoTools/plan/pkg/engine/errors"
	"github.com/PicoTools/plan/pkg/engine/utils"
	"github.com/PicoTools/plan/pkg/parser"
)

// Str implements ObjectImpl interface
type Str struct {
	ObjectImpl
	value   string
	methods map[string]*NativeFunc
}

// NewBool create new Str object containing string value
func NewStr(v string) *Str {
	s := &Str{value: v}
	s.fillMethods()
	return s
}

// fillMethods fills Str object's methods in storage
func (o *Str) fillMethods() {
	o.methods = make(map[string]*NativeFunc)
	o.methods["len"] = NewNativeFunc("len", o.MethodLen)
	o.methods["reverse"] = NewNativeFunc("reverse", o.MethodReverse)
	o.methods["split"] = NewNativeFunc("split", o.MethodSplit)
}

// TypeName returns type name
func (o *Str) TypeName() string {
	return "str"
}

// Strings returns string representation of Str object (itself)
func (o *Str) String() string {
	return o.value
}

// GetValue returns underly string value
func (o *Str) GetValue() any {
	return o.value
}

// IndexGet returns value from Str by index. Operates on rune level
func (o *Str) IndexGet(rs Object) (Object, error) {
	idx, ok := rs.(*Int)
	if !ok {
		return nil, planerrors.ErrInvalidIndexType
	}
	if idx.value < 0 || idx.value >= int64(utf8.RuneCountInString(o.value)) {
		return nil, planerrors.ErrIndexOutOfRange
	}
	// rune's width (in bytes)
	w := 0
	// runes counter
	c := 0
	for i := 0; i < len(o.value); i += w {
		var r rune
		r, w = utf8.DecodeRuneInString(o.value[i:])
		if c == int(idx.value) {
			return NewStr(string(r)), nil
		}
		c++
	}
	return nil, planerrors.ErrNotImplemented
}

// IndexGet set value in Str by index. Operates on rune level
func (o *Str) IndexSet(idx Object, rs Object) error {
	idxInt, ok := idx.(*Int)
	if !ok {
		return planerrors.ErrInvalidIndexType
	}
	rsStr, ok := rs.(*Str)
	if !ok {
		return fmt.Errorf("invalid type '%s' of assignment to str", rs.TypeName())
	}
	if idxInt.value < 0 || idxInt.value >= int64(utf8.RuneCountInString(o.value)) {
		return planerrors.ErrIndexOutOfRange
	}
	if utf8.RuneCountInString(rsStr.value) != 1 {
		return fmt.Errorf("expected str with length of 1, got %d", len(rsStr.value))
	}
	res := ""
	w := 0
	c := 0
	for i := 0; i < len(o.value); i += w {
		var r rune
		r, w = utf8.DecodeRuneInString(o.value[i:])
		if c == int(idxInt.value) {
			res += rsStr.value
		} else {
			res += string(r)
		}
		c++
	}
	o.value = res
	return nil
}

// BinaryOp provides logic of binary operations between 2 objects
func (o *Str) BinaryOp(op int, rhs Object) (Object, error) {
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
	case parser.PLANLexerAssMul:
		return o.Mul(rhs)
	case parser.PLANLexerMultiply:
		return o.Mul(rhs)
	}
	return nil, planerrors.ErrInvalidOp
}

// UnaryOp provides logic of unary operations on object
func (o *Str) UnaryOp(op int) (Object, error) {
	switch op {
	case parser.PLANLexerNot:
		return o.LogicalNot()
	}
	return nil, planerrors.ErrInvalidOp
}

// LogicalNot implements logical NOT on Str object
func (o *Str) LogicalNot() (Object, error) {
	if len(o.value) == 0 {
		return NewBool(true), nil
	}
	return NewBool(false), nil
}

// LogicalOr implements logical OR between Str object and other types of objects
func (o *Str) LogicalOr(rs Object) (Object, error) {
	switch rs.(type) {
	case *Null:
		return o, nil
	}
	return nil, planerrors.ErrInvalidOp
}

// LogicalAnd implements logical AND between Str object and other types of objects
func (o *Str) LogicalAnd(rs Object) (Object, error) {
	switch rs.(type) {
	case *Null:
		return rs, nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Equal implements checking of equation between Str object and other types of objects
func (o *Str) Equal(rs Object) (Object, error) {
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
		return NewBool(false), nil
	case *Null:
		return NewBool(false), nil
	case *Str:
		return NewBool(o.value == rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// NotEqual implements checking of unequation between Str object and other types of objects
func (o *Str) NotEqual(rs Object) (Object, error) {
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
		return NewBool(true), nil
	case *Null:
		return NewBool(true), nil
	case *Str:
		return NewBool(o.value != rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// GtEq implements 'great than or equal' checking between Str object and other types of objects
func (o *Str) GtEq(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Str:
		return NewBool(o.value >= rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Gt implements 'great than' checking between Str object and other types of objects
func (o *Str) Gt(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Str:
		return NewBool(o.value > rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// LtEq implements 'less than or equal' checking between Str object and other types of objects
func (o *Str) LtEq(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Str:
		return NewBool(o.value <= rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Lt implements 'less than' checking between Str object and other types of objects
func (o *Str) Lt(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Str:
		return NewBool(o.value < rs.value), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Add implements summation of Str object and other types of objects
func (o *Str) Add(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Str:
		o.value += rs.value
		return o, nil
	}
	return nil, planerrors.ErrInvalidOp
}

// Mul implements multiplication of Str object and other types of objects
func (o *Str) Mul(rs Object) (Object, error) {
	switch rs := rs.(type) {
	case *Bool:
		return NewStr(strings.Repeat(o.value, int(utils.BoolToInt(rs.value)))), nil
	case *Int:
		// in case of negative value
		count := 0
		if int(rs.value) > 0 {
			count = int(rs.value)
		}
		return NewStr(strings.Repeat(o.value, count)), nil
	}
	return nil, planerrors.ErrInvalidOp
}

// MethodCall dispatch method from storage for Str object and executes it
func (o *Str) MethodCall(name string, args ...Object) (Object, error) {
	m := o.methods[name]
	if m == nil {
		return nil, planerrors.ErrUnknownMethod
	}
	return m.Call(args...)
}

// MethodLen returns number of runes in Str object via method
func (o *Str) MethodLen(args ...Object) (Object, error) {
	if len(args) > 0 {
		return nil, fmt.Errorf("expecting 0 arguments, got %d", len(args))
	}
	return NewInt(int64(utf8.RuneCountInString(o.value))), nil
}

// MethodReverse reverse runes in Str objects via method
func (o *Str) MethodReverse(args ...Object) (Object, error) {
	if len(args) > 0 {
		return nil, fmt.Errorf("expecting 0 arguments, got %d", len(args))
	}
	size := len(o.value)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(o.value[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	o.value = string(buf)
	return NewNull(), nil
}

// MethodSplit splits Str object on list of subtrings by delimeter via method
func (o *Str) MethodSplit(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 arguments, got %d", len(args))
	}
	key, ok := args[0].(*Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[1].TypeName())
	}
	temp := strings.Split(o.value, key.value)
	result := make([]Object, 0)
	for _, v := range temp {
		result = append(result, NewStr(v))
	}
	return NewList(result), nil
}
