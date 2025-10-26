package builtin

import (
	"encoding/base64"
	"fmt"

	"github.com/PicoTools/plan/pkg/engine/object"
)

func Base64Enc(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting 'str' as 1st argument, got '%s'", args[0].TypeName())
	}
	return object.NewStr(base64.StdEncoding.EncodeToString([]byte(str.Value()))), nil
}
