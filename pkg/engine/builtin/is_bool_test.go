package builtin

import (
	"testing"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/stretchr/testify/require"
)

func TestIsBool(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		v, err := IsBool(object.NewBool(true))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(true), v)
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		v, err := IsBool(object.NewInt(-1))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(false), v)
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()
		v, err := IsBool(object.NewFloat(1.23))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(false), v)
	})

	t.Run("list", func(t *testing.T) {
		t.Parallel()
		v, err := IsBool(object.NewList([]object.Object{}))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(false), v)
	})

	t.Run("dict", func(t *testing.T) {
		t.Parallel()
		v, err := IsBool(object.NewDict(map[string]object.Object{}))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(false), v)
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		v, err := IsBool(object.NewStr("abc"))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(false), v)
	})

	t.Run("null", func(t *testing.T) {
		t.Parallel()
		v, err := IsBool(object.NewNull())
		require.NoError(t, err)
		require.Equal(t, object.NewBool(false), v)
	})

	t.Run("more args", func(t *testing.T) {
		t.Parallel()
		v, err := IsBool(object.NewBool(true), object.NewFloat(2.0))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})
}
