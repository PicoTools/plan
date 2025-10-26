package builtin

import (
	"os"
	"testing"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/stretchr/testify/require"
)

func TestFwrite(t *testing.T) {
	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Fwrite(object.NewStr("/unknown/path/unknown/file.txt"))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Fwrite(object.NewStr("/tmp/absolutly-random-path.txt"), object.NewStr("hello world"))
		require.NoError(t, err)
		data, err := os.ReadFile("/tmp/absolutly-random-path.txt")
		require.NoError(t, err)
		require.Equal(t, "hello world", string(data))
		require.Equal(t, object.NewNull(), v)
		require.NoError(t, os.Remove("/tmp/absolutly-random-path.txt"))
	})

	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		v, err := Fwrite(object.NewBool(true), object.NewBool(true))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		v, err := Fwrite(object.NewInt(1), object.NewInt(1))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()
		v, err := Fwrite(object.NewFloat(0.012), object.NewFloat(0.012))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("null", func(t *testing.T) {
		t.Parallel()
		v, err := Fwrite(object.NewNull(), object.NewNull())
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("list", func(t *testing.T) {
		t.Parallel()
		v, err := Fwrite(object.NewList(nil), object.NewList(nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("dict", func(t *testing.T) {
		t.Parallel()
		v, err := Fwrite(object.NewDict(nil), object.NewDict(nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("native func", func(t *testing.T) {
		t.Parallel()
		v, err := Fwrite(object.NewNativeFunc("a", nil), object.NewNativeFunc("a", nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("runtime func", func(t *testing.T) {
		t.Parallel()
		v, err := Fwrite(object.NewRuntimeFunc(nil, nil), object.NewRuntimeFunc(nil, nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("more args", func(t *testing.T) {
		t.Parallel()
		v, err := Fwrite(object.NewBool(true), object.NewFloat(2.0), object.NewBool(true))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})
}
