package builtin

import (
	"crypto/sha256"
	"fmt"

	"github.com/PicoTools/plan/pkg/engine/object"
)

func Sha256(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting 'str' as 1st argument, got '%s'", args[0].TypeName())
	}
	sha256sum := sha256.Sum256([]byte(str.Value()))
	return object.NewStr(string(sha256sum[:])), nil
}
