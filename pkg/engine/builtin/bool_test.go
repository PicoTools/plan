package builtin

import (
	"testing"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/stretchr/testify/require"
)

func TestBool(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		v, err := Bool(object.NewBool(true))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(true), v)
	})

	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		v, err := Bool(object.NewBool(false))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(false), v)
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		v, err := Bool(object.NewInt(0))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(false), v)
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		v, err := Bool(object.NewInt(2))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(true), v)
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()
		v, err := Bool(object.NewFloat(0))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(false), v)
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()
		v, err := Bool(object.NewFloat(0.1))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(true), v)
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()
		v, err := Bool(object.NewFloat(1.23))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(true), v)
	})

	t.Run("null", func(t *testing.T) {
		t.Parallel()
		v, err := Bool(object.NewNull())
		require.NoError(t, err)
		require.Equal(t, object.NewBool(false), v)
	})

	t.Run("dict", func(t *testing.T) {
		t.Parallel()
		v, err := Bool(object.NewDict(map[string]object.Object{}))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(false), v)
	})

	t.Run("dict", func(t *testing.T) {
		t.Parallel()
		v, err := Bool(object.NewDict(map[string]object.Object{
			"a": object.NewBool(false),
			"b": object.NewInt(0),
		}))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(true), v)
	})

	t.Run("list", func(t *testing.T) {
		t.Parallel()
		v, err := Bool(object.NewList([]object.Object{}))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(false), v)
	})

	t.Run("list", func(t *testing.T) {
		t.Parallel()
		v, err := Bool(object.NewList([]object.Object{
			object.NewNull(),
			object.NewStr("hello"),
		}))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(true), v)
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		v, err := Bool(object.NewStr(""))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(false), v)
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		v, err := Bool(object.NewStr("abc"))
		require.NoError(t, err)
		require.Equal(t, object.NewBool(true), v)
	})

	t.Run("more args", func(t *testing.T) {
		t.Parallel()
		v, err := Bool(object.NewBool(true), object.NewFloat(2.0))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("native func", func(t *testing.T) {
		t.Parallel()
		v, err := Bool(object.NewNativeFunc("test", func(args ...object.Object) (object.Object, error) { return nil, nil }))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})
}
