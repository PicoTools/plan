package builtin

import (
	"testing"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/stretchr/testify/require"
)

func TestChr(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		v, err := Chr(object.NewBool(true))
		require.NoError(t, err)
		require.Equal(t, "\x01", v.String())
	})

	t.Run("bool", func(t *testing.T) {
		v, err := Chr(object.NewBool(false))
		require.NoError(t, err)
		require.Equal(t, "\x00", v.String())
	})

	t.Run("int", func(t *testing.T) {
		v, err := Chr(object.NewInt(0x33))
		require.NoError(t, err)
		require.Equal(t, "3", v.String())
	})

	t.Run("int", func(t *testing.T) {
		v, err := Chr(object.NewInt(0x3d))
		require.NoError(t, err)
		require.Equal(t, "=", v.String())
	})

	t.Run("int", func(t *testing.T) {
		v, err := Chr(object.NewInt(1060))
		require.NoError(t, err)
		require.Equal(t, "Ф", v.String())
	})

	t.Run("int", func(t *testing.T) {
		v, err := Chr(object.NewInt(0x7684))
		require.NoError(t, err)
		require.Equal(t, "的", v.String())
	})

	t.Run("float", func(t *testing.T) {
		v, err := Chr(object.NewFloat(12.11))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("null", func(t *testing.T) {
		v, err := Chr(object.NewNull())
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("list", func(t *testing.T) {
		v, err := Chr(object.NewList([]object.Object{}))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("dict", func(t *testing.T) {
		v, err := Chr(object.NewDict(make(map[string]object.Object)))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("str", func(t *testing.T) {
		v, err := Chr(object.NewStr("AAAA"))
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
		v, err := Bool(object.NewBool(true), object.NewFloat(2.0))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})
}
