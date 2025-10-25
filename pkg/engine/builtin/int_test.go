package builtin

import (
	"testing"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/stretchr/testify/require"
)

func TestInt(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		v, err := Int(object.NewBool(true))
		require.NoError(t, err)
		require.Equal(t, object.NewInt(1), v)
	})

	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		v, err := Int(object.NewBool(false))
		require.NoError(t, err)
		require.Equal(t, object.NewInt(0), v)
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()
		v, err := Int(object.NewFloat(0.1))
		require.NoError(t, err)
		require.Equal(t, object.NewInt(0), v)
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()
		v, err := Int(object.NewFloat(1.63))
		require.NoError(t, err)
		require.Equal(t, object.NewInt(1), v)
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()
		v, err := Int(object.NewFloat(1.0))
		require.NoError(t, err)
		require.Equal(t, object.NewInt(1), v)
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		v, err := Int(object.NewInt(0))
		require.NoError(t, err)
		require.Equal(t, object.NewInt(0), v)
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		v, err := Int(object.NewInt(2))
		require.NoError(t, err)
		require.Equal(t, object.NewInt(2), v)
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		v, err := Int(object.NewStr("0123"))
		require.NoError(t, err)
		require.Equal(t, object.NewInt(123), v)
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		v, err := Int(object.NewStr("abc"))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("list", func(t *testing.T) {
		t.Parallel()
		v, err := Int(object.NewList([]object.Object{}))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("dict", func(t *testing.T) {
		t.Parallel()
		v, err := Int(object.NewDict(map[string]object.Object{
			"a": object.NewNull(),
		}))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("null", func(t *testing.T) {
		t.Parallel()
		v, err := Int(object.NewNull())
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("native func", func(t *testing.T) {
		v, err := Chr(object.NewNativeFunc("a", nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("runtime func", func(t *testing.T) {
		v, err := Chr(object.NewRuntimeFunc(nil, nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("more args", func(t *testing.T) {
		t.Parallel()
		v, err := Int(object.NewBool(true), object.NewFloat(2.0))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})
}
