package builtin

import (
	"fmt"
	"os"

	"github.com/PicoTools/plan/pkg/engine/object"
)

func Fwrite(args ...object.Object) (object.Object, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("expecting 2 arguments, got %d", len(args))
	}
	path, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting 'str' as 1st argument, got '%s'", args[0].TypeName())
	}
	data, ok := args[1].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting 'str' as 2nd argument, got '%s'", args[1].TypeName())
	}
	if err := os.WriteFile(path.Value(), []byte(data.Value()), 0640); err != nil {
		return nil, err
	}
	return object.NewNull(), nil
}
