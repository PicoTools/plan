package builtin

import (
	"testing"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/stretchr/testify/require"
)

func TestAsert(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		t.Parallel()
		v, err := Assert(object.NewBool(true))
		require.NoError(t, err)
		require.Equal(t, object.NewNull(), v)
	})

	t.Run("false", func(t *testing.T) {
		t.Parallel()
		v, err := Assert(object.NewBool(false))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("more args", func(t *testing.T) {
		t.Parallel()
		v, err := Assert(object.NewBool(true), object.NewBool(false))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})

	t.Run("invalid arg type", func(t *testing.T) {
		t.Parallel()
		v, err := Assert(object.NewInt(123))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})
}
