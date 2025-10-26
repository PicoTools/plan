package builtin

import (
	"fmt"
	"strconv"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/PicoTools/plan/pkg/engine/utils"
)

func Int(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch obj := args[0].(type) {
	case *object.Bool:
		return object.NewInt(utils.BoolToInt(obj.Value())), nil
	case *object.Float:
		return object.NewInt(utils.FloatToInt(obj.Value())), nil
	case *object.Int:
		return obj, nil
	case *object.Str:
		val, err := strconv.Atoi(obj.Value())
		if err != nil {
			return nil, fmt.Errorf("unable convert 'str' to 'int': %v", err)
		}
		return object.NewInt(int64(val)), nil
	}
	return nil, fmt.Errorf("unsupported conversation from '%s'", args[0].TypeName())
}
