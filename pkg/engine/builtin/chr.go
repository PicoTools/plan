package builtin

import (
	"fmt"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/PicoTools/plan/pkg/engine/utils"
)

func Chr(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch obj := args[0].(type) {
	case *object.Bool:
		return object.NewStr(string(rune(utils.BoolToInt(obj.Value())))), nil
	case *object.Int:
		return object.NewStr(string(rune(obj.Value()))), nil
	}
	return nil, fmt.Errorf("unsupported type '%s'", args[0].TypeName())
}
