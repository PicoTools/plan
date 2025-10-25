package builtin

import (
	"testing"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/stretchr/testify/require"
)

func TestIsInt(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		v, err := IsInt(object.NewBool(true))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(false), v)
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		v, err := IsInt(object.NewInt(-1))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(true), v)
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()
		v, err := IsInt(object.NewFloat(1.23))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(false), v)
	})

	t.Run("list", func(t *testing.T) {
		t.Parallel()
		v, err := IsInt(object.NewList([]object.Object{}))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(false), v)
	})

	t.Run("dict", func(t *testing.T) {
		t.Parallel()
		v, err := IsInt(object.NewDict(map[string]object.Object{}))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(false), v)
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		v, err := IsInt(object.NewStr("abc"))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(false), v)
	})

	t.Run("null", func(t *testing.T) {
		t.Parallel()
		v, err := IsInt(object.NewNull())
		require.NoError(t, err)
		require.Equal(t, object.NewBool(false), v)
	})

	t.Run("more args", func(t *testing.T) {
		t.Parallel()
		v, err := IsInt(object.NewBool(true), object.NewFloat(2.0))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})
}
