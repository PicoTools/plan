package builtin

import (
	"fmt"

	"github.com/PicoTools/plan/pkg/engine/object"
)

func IsFloat(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.Float); ok {
		return object.NewBool(true), nil
	}
	return object.NewBool(false), nil
}
