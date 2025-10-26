package builtin

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"

	"github.com/PicoTools/plan/pkg/engine/object"
)

func Gunzip(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting 'str' as 1st argument, got '%s'", args[0].TypeName())
	}
	r, err := gzip.NewReader(bytes.NewBufferString(str.Value()))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = r.Close()
	}()
	res, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return object.NewStr(string(res)), nil
}
