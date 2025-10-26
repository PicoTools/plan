package builtin

import (
	"testing"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/stretchr/testify/require"
)

func TestFloat(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		v, err := Float(object.NewBool(true))
		require.NoError(t, err)
		require.Equal(t, object.NewFloat(1.0), v)
	})

	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		v, err := Float(object.NewBool(false))
		require.NoError(t, err)
		require.Equal(t, object.NewFloat(0.0), v)
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()
		v, err := Float(object.NewFloat(0.0))
		require.NoError(t, err)
		require.Equal(t, object.NewFloat(0.0), v)
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()
		v, err := Float(object.NewFloat(0.1223))
		require.NoError(t, err)
		require.Equal(t, object.NewFloat(0.1223), v)
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()
		v, err := Float(object.NewFloat(10.09))
		require.NoError(t, err)
		require.Equal(t, object.NewFloat(10.09), v)
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		v, err := Float(object.NewInt(0))
		require.NoError(t, err)
		require.Equal(t, object.NewFloat(0.0), v)
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		v, err := Float(object.NewInt(2))
		require.NoError(t, err)
		require.Equal(t, object.NewFloat(2.0), v)
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		v, err := Float(object.NewStr("abc"))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("list", func(t *testing.T) {
		t.Parallel()
		v, err := Float(object.NewList([]object.Object{}))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("dict", func(t *testing.T) {
		t.Parallel()
		v, err := Float(object.NewDict(map[string]object.Object{
			"a": object.NewNull(),
		}))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("null", func(t *testing.T) {
		t.Parallel()
		v, err := Float(object.NewNull())
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("native func", func(t *testing.T) {
		t.Parallel()
		v, err := Float(object.NewNativeFunc("a", nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("runtime func", func(t *testing.T) {
		t.Parallel()
		v, err := Float(object.NewRuntimeFunc(nil, nil))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("more args", func(t *testing.T) {
		t.Parallel()
		v, err := Float(object.NewBool(true), object.NewFloat(2.0))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})
}
