package builtin

import (
	"fmt"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/PicoTools/plan/pkg/engine/utils"
)

func Float(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch obj := args[0].(type) {
	case *object.Bool:
		return object.NewFloat(utils.BoolToFloat(obj.Value())), nil
	case *object.Float:
		return obj, nil
	case *object.Int:
		return object.NewFloat(utils.IntToFloat(obj.Value())), nil
	}
	return nil, fmt.Errorf("unsupported conversation from '%s'", args[0].TypeName())
}
