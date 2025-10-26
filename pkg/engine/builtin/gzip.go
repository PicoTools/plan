package builtin

import (
	"bytes"
	"compress/gzip"
	"fmt"

	"github.com/PicoTools/plan/pkg/engine/object"
)

func Gzip(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting 'str' as 1st argument, got '%s'", args[0].TypeName())
	}
	b := &bytes.Buffer{}
	w := gzip.NewWriter(b)
	_, err := w.Write([]byte(str.Value()))
	if err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return object.NewStr(b.String()), nil
}
