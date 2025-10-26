package builtin

import (
	"encoding/hex"
	"fmt"

	"github.com/PicoTools/plan/pkg/engine/object"
)

func Unhex(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting 'str' as 1st argument, got '%s'", args[0].TypeName())
	}
	v, err := hex.DecodeString(str.Value())
	if err != nil {
		return nil, err
	}
	return object.NewStr(string(v)), nil
}
