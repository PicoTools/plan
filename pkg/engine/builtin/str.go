package builtin

import (
	"fmt"

	"github.com/PicoTools/plan/pkg/engine/object"
)

func Str(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch obj := args[0].(type) {
	case *object.Bool:
		return object.NewStr(obj.String()), nil
	case *object.Dict:
		return object.NewStr(obj.String()), nil
	case *object.Float:
		return object.NewStr(obj.String()), nil
	case *object.Int:
		return object.NewStr(obj.String()), nil
	case *object.List:
		return object.NewStr(obj.String()), nil
	case *object.Null:
		return object.NewStr(obj.String()), nil
	case *object.Str:
		return obj, nil
	case *object.NativeFunc:
		return object.NewStr(obj.String()), nil
	case *object.RuntimeFunc:
		return object.NewStr(obj.String()), nil
	}
	return nil, fmt.Errorf("unsupported conversation from '%s'", args[0].TypeName())
}
