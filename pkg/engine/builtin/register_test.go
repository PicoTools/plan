package builtin

import (
	"testing"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/PicoTools/plan/pkg/engine/storage"
	"github.com/stretchr/testify/require"
)

func TestRegisterBuiltin(t *testing.T) {
	t.Run("reg fn", func(t *testing.T) {
		registerBuiltin("test01", func(args ...object.Object) (object.Object, error) {
			return nil, nil
		})
		_, ok := storage.BuiltinFunctions["test01"]
		require.Equal(t, true, ok)
	})

	t.Run("reg fn", func(t *testing.T) {
		registerBuiltin("test02", func(args ...object.Object) (object.Object, error) {
			return nil, nil
		})
		_, ok := storage.BuiltinFunctions["test02"]
		require.Equal(t, true, ok)
	})

	t.Run("reg fn", func(t *testing.T) {
		registerBuiltin("test03", func(args ...object.Object) (object.Object, error) {
			return nil, nil
		})
		_, ok := storage.BuiltinFunctions["test03"]
		require.Equal(t, true, ok)
	})
}
