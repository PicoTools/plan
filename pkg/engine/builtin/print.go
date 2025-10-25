package builtin

import (
	"fmt"

	"github.com/PicoTools/plan/pkg/engine/object"
)

func Print(args ...object.Object) (object.Object, error) {
	for _, arg := range args {
		fmt.Print(arg.String())
	}
	return object.NewNull(), nil
}
