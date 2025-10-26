package builtin

import (
	"testing"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/stretchr/testify/require"
)

func TestGunzip(t *testing.T) {
	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Gunzip(object.NewStr("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00"))
		require.NoError(t, err)
		require.Equal(t, "", v.String())
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Gunzip(object.NewStr("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xffbdbf\x01\x04\x00\x00\xff\xff\xcd\xfb<\xb6\x04\x00\x00\x00"))
		require.NoError(t, err)
		require.Equal(t, "\x01\x02\x03\x04", v.String())
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Gunzip(object.NewStr("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xcaH\xcd\xc9\xc9\a\x04\x00\x00\xff\xff\x86\xa6\x106\x05\x00\x00\x00"))
		require.NoError(t, err)
		require.Equal(t, "hello", v.String())
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Gunzip(object.NewStr("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xba\xb0\xffbÅ\x1d\x176]\xd8z\xb1\t\x10\x00\x00\xff\xffN>v.\f\x00\x00\x00"))
		require.NoError(t, err)
		require.Equal(t, "привет", v.String())
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Gunzip(object.NewStr("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xffz\xb2w\xc1ӥ{\x01\x01\x00\x00\xff\xffA\xb8\xa2P\x06\x00\x00\x00"))
		require.NoError(t, err)
		require.Equal(t, "你好", v.String())
	})

	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		v, err := Gunzip(object.NewBool(true))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		v, err := Gunzip(object.NewInt(1))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()
		v, err := Gunzip(object.NewFloat(0.012))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("null", func(t *testing.T) {
		t.Parallel()
		v, err := Gunzip(object.NewNull())
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("list", func(t *testing.T) {
		t.Parallel()
		v, err := Gunzip(object.NewList(nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("dict", func(t *testing.T) {
		t.Parallel()
		v, err := Gunzip(object.NewDict(nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("native func", func(t *testing.T) {
		t.Parallel()
		v, err := Gunzip(object.NewNativeFunc("a", nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("runtime func", func(t *testing.T) {
		t.Parallel()
		v, err := Gunzip(object.NewRuntimeFunc(nil, nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("more args", func(t *testing.T) {
		t.Parallel()
		v, err := Gunzip(object.NewBool(true), object.NewFloat(2.0))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})
}
