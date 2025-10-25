package builtin

import (
	"testing"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/stretchr/testify/require"
)

func TestBase32Enc(t *testing.T) {
	t.Run("str", func(t *testing.T) {
		v, err := Base32Enc(object.NewStr(""))
		require.NoError(t, err)
		require.Equal(t, "", v.String())
	})

	t.Run("str", func(t *testing.T) {
		v, err := Base32Enc(object.NewStr("hello"))
		require.NoError(t, err)
		require.Equal(t, "NBSWY3DP", v.String())
	})

	t.Run("str", func(t *testing.T) {
		v, err := Base32Enc(object.NewStr("привет"))
		require.NoError(t, err)
		require.Equal(t, "2C75DAGQXDILFUFV2GBA====", v.String())
	})

	t.Run("str", func(t *testing.T) {
		v, err := Base32Enc(object.NewStr("你好"))
		require.NoError(t, err)
		require.Equal(t, "4S62BZNFXU======", v.String())
	})

	t.Run("bool", func(t *testing.T) {
		v, err := Base32Enc(object.NewBool(true))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("int", func(t *testing.T) {
		v, err := Base32Enc(object.NewInt(1))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("float", func(t *testing.T) {
		v, err := Base32Enc(object.NewFloat(0.012))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("null", func(t *testing.T) {
		v, err := Base32Enc(object.NewNull())
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("list", func(t *testing.T) {
		v, err := Base32Enc(object.NewList(nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("dict", func(t *testing.T) {
		v, err := Base32Enc(object.NewDict(nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("native func", func(t *testing.T) {
		v, err := Base32Enc(object.NewNativeFunc("a", nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("runtime func", func(t *testing.T) {
		v, err := Base32Enc(object.NewRuntimeFunc(nil, nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("more args", func(t *testing.T) {
		t.Parallel()
		v, err := Base32Enc(object.NewBool(true), object.NewFloat(2.0))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})
}
