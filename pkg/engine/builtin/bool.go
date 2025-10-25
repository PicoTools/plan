package builtin

import (
	"fmt"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/PicoTools/plan/pkg/engine/utils"
)

func Bool(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch obj := args[0].(type) {
	case *object.Bool:
		return obj, nil
	case *object.Dict:
		if len(obj.Value()) == 0 {
			return object.NewBool(false), nil
		}
		return object.NewBool(true), nil
	case *object.Float:
		return object.NewBool(utils.FloatToBool(obj.Value())), nil
	case *object.Int:
		return object.NewBool(utils.IntToBool(obj.Value())), nil
	case *object.List:
		if len(obj.Value()) == 0 {
			return object.NewBool(false), nil
		}
		return object.NewBool(true), nil
	case *object.Null:
		return object.NewBool(false), nil
	case *object.Str:
		if len(obj.Value()) == 0 {
			return object.NewBool(false), nil
		}
		return object.NewBool(true), nil
	}
	return nil, fmt.Errorf("unsupported conversation from '%s'", args[0].TypeName())
}
