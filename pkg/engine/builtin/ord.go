package builtin

import (
	"fmt"
	"unicode/utf8"

	"github.com/PicoTools/plan/pkg/engine/object"
)

func Ord(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch obj := args[0].(type) {
	case *object.Str:
		if utf8.RuneCountInString(obj.Value()) != 1 {
			return nil, fmt.Errorf("'str' must have only one char")
		}
		r, _ := utf8.DecodeRuneInString(obj.Value())
		return object.NewInt(int64(r)), nil
	}
	return nil, fmt.Errorf("unsupported type '%s'", args[0].TypeName())
}
