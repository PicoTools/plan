package builtin

import (
	"os"
	"testing"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/stretchr/testify/require"
)

func TestFread(t *testing.T) {
	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Fread(object.NewStr("/unknown/path/unknown/file.txt"))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		// Create temp file
		f, err := os.CreateTemp("", "tmpfile-")
		require.NoError(t, err)
		// Write data
		_, err = f.Write([]byte("hello world"))
		require.NoError(t, err)
		// Read data
		v, err := Fread(object.NewStr(f.Name()))
		require.NoError(t, err)
		require.Equal(t, "hello world", v.String())
		// Close and delete temp file
		require.NoError(t, f.Close())
		require.NoError(t, os.Remove(f.Name()))
	})

	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		v, err := Fread(object.NewBool(true))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		v, err := Fread(object.NewInt(1))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()
		v, err := Fread(object.NewFloat(0.012))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("null", func(t *testing.T) {
		t.Parallel()
		v, err := Fread(object.NewNull())
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("list", func(t *testing.T) {
		t.Parallel()
		v, err := Fread(object.NewList(nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("dict", func(t *testing.T) {
		t.Parallel()
		v, err := Fread(object.NewDict(nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("native func", func(t *testing.T) {
		t.Parallel()
		v, err := Fread(object.NewNativeFunc("a", nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("runtime func", func(t *testing.T) {
		t.Parallel()
		v, err := Fread(object.NewRuntimeFunc(nil, nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("more args", func(t *testing.T) {
		t.Parallel()
		v, err := Fread(object.NewBool(true), object.NewFloat(2.0))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})
}
