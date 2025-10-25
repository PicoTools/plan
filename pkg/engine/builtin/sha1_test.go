package builtin

import (
	"testing"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/stretchr/testify/require"
)

func TestSha1(t *testing.T) {
	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Sha1(object.NewStr(""))
		require.NoError(t, err)
		require.Equal(t, "\xda9\xa3\xee^kK\r2U\xbf\xef\x95`\x18\x90\xaf\xd8\a\t", v.String())
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Sha1(object.NewStr("hello"))
		require.NoError(t, err)
		require.Equal(t, "\xaa\xf4\xc6\x1d\xdc\xc5\xe8\xa2ھ\xde\x0f;H,ٮ\xa9CM", v.String())
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Sha1(object.NewStr("привет"))
		require.NoError(t, err)
		require.Equal(t, "\xe2E\x05\xf9M\xb2\xb5\xdfL|%\x96\xb0x\x8er\x0e\a0!", v.String())
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Sha1(object.NewStr("你好"))
		require.NoError(t, err)
		require.Equal(t, "D\x0e\xe0\x85:\xd1韖+c\xe4Y\xef\x99-|!\x17\"", v.String())
	})

	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		v, err := Sha1(object.NewBool(true))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		v, err := Sha1(object.NewInt(1))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()
		v, err := Sha1(object.NewFloat(0.012))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("null", func(t *testing.T) {
		t.Parallel()
		v, err := Sha1(object.NewNull())
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("list", func(t *testing.T) {
		t.Parallel()
		v, err := Sha1(object.NewList(nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("dict", func(t *testing.T) {
		t.Parallel()
		v, err := Sha1(object.NewDict(nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("native func", func(t *testing.T) {
		t.Parallel()
		v, err := Sha1(object.NewNativeFunc("a", nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("runtime func", func(t *testing.T) {
		t.Parallel()
		v, err := Sha1(object.NewRuntimeFunc(nil, nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("more args", func(t *testing.T) {
		t.Parallel()
		v, err := Sha1(object.NewBool(true), object.NewFloat(2.0))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})
}
