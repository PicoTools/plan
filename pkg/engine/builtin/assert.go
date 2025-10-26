package builtin

import (
	"fmt"

	"github.com/PicoTools/plan/pkg/engine/object"
)

func Assert(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	val, ok := args[0].(*object.Bool)
	if !ok {
		return nil, fmt.Errorf("expecting 'bool', got '%s'", args[0].TypeName())
	}
	if !val.Value() {
		return nil, fmt.Errorf("assertion occured")
	}
	return object.NewNull(), nil
}
