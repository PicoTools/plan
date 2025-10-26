package builtin

import (
	"testing"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/stretchr/testify/require"
)

func TestUnhex(t *testing.T) {
	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Unhex(object.NewStr(""))
		require.NoError(t, err)
		require.Equal(t, "", v.String())
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Unhex(object.NewStr("01020304"))
		require.NoError(t, err)
		require.Equal(t, "\x01\x02\x03\x04", v.String())
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Unhex(object.NewStr("68656c6c6f"))
		require.NoError(t, err)
		require.Equal(t, "hello", v.String())
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Unhex(object.NewStr("d0bfd180d0b8d0b2d0b5d182"))
		require.NoError(t, err)
		require.Equal(t, "привет", v.String())
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Unhex(object.NewStr("e4bda0e5a5bd"))
		require.NoError(t, err)
		require.Equal(t, "你好", v.String())
	})

	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		v, err := Unhex(object.NewBool(true))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		v, err := Unhex(object.NewInt(1))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()
		v, err := Unhex(object.NewFloat(0.012))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("null", func(t *testing.T) {
		t.Parallel()
		v, err := Unhex(object.NewNull())
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("list", func(t *testing.T) {
		t.Parallel()
		v, err := Unhex(object.NewList(nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("dict", func(t *testing.T) {
		t.Parallel()
		v, err := Unhex(object.NewDict(nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("native func", func(t *testing.T) {
		t.Parallel()
		v, err := Unhex(object.NewNativeFunc("a", nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("runtime func", func(t *testing.T) {
		t.Parallel()
		v, err := Unhex(object.NewRuntimeFunc(nil, nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("more args", func(t *testing.T) {
		t.Parallel()
		v, err := Unhex(object.NewBool(true), object.NewFloat(2.0))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})
}
