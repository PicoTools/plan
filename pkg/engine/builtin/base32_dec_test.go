package builtin

import (
	"testing"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/stretchr/testify/require"
)

func TestBase32Dec(t *testing.T) {
	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Base32Dec(object.NewStr(""))
		require.NoError(t, err)
		require.Equal(t, "", v.String())
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Base32Dec(object.NewStr("NBSWY3DP"))
		require.NoError(t, err)
		require.Equal(t, "hello", v.String())
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Base32Dec(object.NewStr("2C75DAGQXDILFUFV2GBA===="))
		require.NoError(t, err)
		require.Equal(t, "привет", v.String())
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Base32Dec(object.NewStr("4S62BZNFXU======"))
		require.NoError(t, err)
		require.Equal(t, "你好", v.String())
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Base32Dec(object.NewStr("aaaaaa"))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		v, err := Base32Dec(object.NewBool(true))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		v, err := Base32Dec(object.NewInt(1))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()
		v, err := Base32Dec(object.NewFloat(0.012))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("null", func(t *testing.T) {
		t.Parallel()
		v, err := Base32Dec(object.NewNull())
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("list", func(t *testing.T) {
		t.Parallel()
		v, err := Base32Dec(object.NewList(nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("dict", func(t *testing.T) {
		t.Parallel()
		v, err := Base32Dec(object.NewDict(nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("native func", func(t *testing.T) {
		t.Parallel()
		v, err := Base32Dec(object.NewNativeFunc("a", nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("runtime func", func(t *testing.T) {
		t.Parallel()
		v, err := Base32Dec(object.NewRuntimeFunc(nil, nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("more args", func(t *testing.T) {
		t.Parallel()
		v, err := Base32Dec(object.NewBool(true), object.NewFloat(2.0))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})
}
