package builtin

import (
	"testing"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/stretchr/testify/require"
)

func TestMd5(t *testing.T) {
	t.Run("str", func(t *testing.T) {
		v, err := Md5(object.NewStr(""))
		require.NoError(t, err)
		require.Equal(t, "\xd4\x1d\x8cُ\x00\xb2\x04\xe9\x80\t\x98\xec\xf8B~", v.String())
	})

	t.Run("str", func(t *testing.T) {
		v, err := Md5(object.NewStr("hello"))
		require.NoError(t, err)
		require.Equal(t, "]A@*\xbcK*v\xb9q\x9d\x91\x10\x17Œ", v.String())
	})

	t.Run("str", func(t *testing.T) {
		v, err := Md5(object.NewStr("привет"))
		require.NoError(t, err)
		require.Equal(t, "`\x833\xad\xc7/TPx\xed\xe3\xaa\xd7\x1b\xfet", v.String())
	})

	t.Run("str", func(t *testing.T) {
		v, err := Md5(object.NewStr("你好"))
		require.NoError(t, err)
		require.Equal(t, "~\xcah\x9f\r3\x89\xd9ަj\xe1\x12\xe5\xcf\xd7", v.String())
	})

	t.Run("bool", func(t *testing.T) {
		v, err := Md5(object.NewBool(true))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("int", func(t *testing.T) {
		v, err := Md5(object.NewInt(1))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("float", func(t *testing.T) {
		v, err := Md5(object.NewFloat(0.012))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("null", func(t *testing.T) {
		v, err := Md5(object.NewNull())
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("list", func(t *testing.T) {
		v, err := Md5(object.NewList(nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("dict", func(t *testing.T) {
		v, err := Md5(object.NewDict(nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("native func", func(t *testing.T) {
		v, err := Md5(object.NewNativeFunc("a", nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("runtime func", func(t *testing.T) {
		v, err := Md5(object.NewRuntimeFunc(nil, nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("more args", func(t *testing.T) {
		t.Parallel()
		v, err := Md5(object.NewBool(true), object.NewFloat(2.0))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})
}
