package builtin

import (
	"testing"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/stretchr/testify/require"
)

func TestSha256(t *testing.T) {
	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Sha256(object.NewStr(""))
		require.NoError(t, err)
		require.Equal(t, "\xe3\xb0\xc4B\x98\xfc\x1c\x14\x9a\xfb\xf4șo\xb9$'\xaeA\xe4d\x9b\x93L\xa4\x95\x99\x1bxR\xb8U", v.String())
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Sha256(object.NewStr("hello"))
		require.NoError(t, err)
		require.Equal(t, ",\xf2M\xba_\xb0\xa3\x0e&\xe8;*Ź\xe2\x9e\x1b\x16\x1e\\\x1f\xa7B^s\x043b\x93\x8b\x98$", v.String())
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Sha256(object.NewStr("привет"))
		require.NoError(t, err)
		require.Equal(t, "\xe5\x8f\x1e\x8cU\xfa\x10[\xdd?@\xe5\x03~\xb0\xb09\xb5\x99\x8dR\xc0^lوx\xdd-\xa5ʲ", v.String())
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Sha256(object.NewStr("你好"))
		require.NoError(t, err)
		require.Equal(t, "g\r\x97CT,\xae>\xa7\xeb\xe3j\xf5k\xd56H\xb0\xa1\x12ab獁\xa3)4\xa7\x110.", v.String())
	})

	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		v, err := Sha256(object.NewBool(true))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		v, err := Sha256(object.NewInt(1))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()
		v, err := Sha256(object.NewFloat(0.012))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("null", func(t *testing.T) {
		t.Parallel()
		v, err := Sha256(object.NewNull())
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("list", func(t *testing.T) {
		t.Parallel()
		v, err := Sha256(object.NewList(nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("dict", func(t *testing.T) {
		t.Parallel()
		v, err := Sha256(object.NewDict(nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("native func", func(t *testing.T) {
		t.Parallel()
		v, err := Sha256(object.NewNativeFunc("a", nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("runtime func", func(t *testing.T) {
		t.Parallel()
		v, err := Sha256(object.NewRuntimeFunc(nil, nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("more args", func(t *testing.T) {
		t.Parallel()
		v, err := Sha256(object.NewBool(true), object.NewFloat(2.0))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})
}
