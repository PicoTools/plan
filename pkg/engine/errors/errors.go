package planerrors

import (
	"errors"
)

var (
	ErrInvalidOp        = errors.New("invalid operator")
	ErrNotImplemented   = errors.New("not implemented")
	ErrDivByZero        = errors.New("division by zero")
	ErrModByZero        = errors.New("modulo by zero")
	ErrInvalidIndexType = errors.New("invalid index type")
	ErrIndexOutOfRange  = errors.New("index out of range")
	ErrUnknownMethod    = errors.New("unknown method")
)
