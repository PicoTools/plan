package builtin

import (
	"testing"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/stretchr/testify/require"
)

func TestOrd(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		v, err := Ord(object.NewBool(true))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		v, err := Ord(object.NewInt(0x33))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		v, err := Ord(object.NewInt(0x3d))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()
		v, err := Ord(object.NewFloat(12.11))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("null", func(t *testing.T) {
		t.Parallel()
		v, err := Ord(object.NewNull())
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("list", func(t *testing.T) {
		t.Parallel()
		v, err := Ord(object.NewList([]object.Object{}))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("dict", func(t *testing.T) {
		t.Parallel()
		v, err := Ord(object.NewDict(make(map[string]object.Object)))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Ord(object.NewStr("A"))
		require.NoError(t, err)
		require.Equal(t, object.NewInt(0x41), v)
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Ord(object.NewStr("Ф"))
		require.NoError(t, err)
		require.Equal(t, object.NewInt(1060), v)
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Ord(object.NewStr("的"))
		require.NoError(t, err)
		require.Equal(t, object.NewInt(0x7684), v)
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Ord(object.NewStr("AAAA"))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("native func", func(t *testing.T) {
		t.Parallel()
		v, err := Ord(object.NewNativeFunc("a", nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("runtime func", func(t *testing.T) {
		t.Parallel()
		v, err := Ord(object.NewRuntimeFunc(nil, nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("more args", func(t *testing.T) {
		t.Parallel()
		v, err := Ord(object.NewBool(true), object.NewFloat(2.0))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})
}
