package builtin

import (
	"fmt"

	"github.com/PicoTools/plan/pkg/engine/object"
)

func Println(args ...object.Object) (object.Object, error) {
	for _, arg := range args {
		fmt.Print(arg.String())
	}
	fmt.Println()
	return object.NewNull(), nil
}
