package builtin

import (
	"testing"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/PicoTools/plan/pkg/parser"
	"github.com/stretchr/testify/require"
)

func TestStr(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		v, err := Str(object.NewBool(true))
		require.NoError(t, err)
		require.Equal(t, "true", v.String())
	})

	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		v, err := Str(object.NewBool(false))
		require.NoError(t, err)
		require.Equal(t, "false", v.String())
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		v, err := Str(object.NewInt(0))
		require.NoError(t, err)
		require.Equal(t, "0", v.String())
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		v, err := Str(object.NewInt(989))
		require.NoError(t, err)
		require.Equal(t, "989", v.String())
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()
		v, err := Str(object.NewFloat(0.0))
		require.NoError(t, err)
		require.Equal(t, "0", v.String())
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()
		v, err := Str(object.NewFloat(0.0000000001))
		require.NoError(t, err)
		require.Equal(t, "0.0000000001", v.String())
	})

	t.Run("null", func(t *testing.T) {
		t.Parallel()
		v, err := Str(object.NewNull())
		require.NoError(t, err)
		require.Equal(t, "<null>", v.String())
	})

	t.Run("list", func(t *testing.T) {
		t.Parallel()
		v, err := Str(object.NewList([]object.Object{
			object.NewBool(true),
			object.NewInt(123),
			object.NewStr("whoami"),
		}))
		require.NoError(t, err)
		require.Equal(t, "[true, 123, whoami]", v.String())
	})

	t.Run("dict", func(t *testing.T) {
		t.Parallel()
		_, err := Str(object.NewDict(map[string]object.Object{
			"a": object.NewDict(map[string]object.Object{}),
			"b": object.NewStr("AAAA"),
		}))
		require.NoError(t, err)
		// TODO: ???
		///require.Equal(t, "{a: {}, b: AAAA}", v.GetValue())
	})

	t.Run("str", func(t *testing.T) {
		t.Parallel()
		v, err := Str(object.NewStr("hello world"))
		require.NoError(t, err)
		require.Equal(t, "hello world", v.String())
	})

	t.Run("native func", func(t *testing.T) {
		t.Parallel()
		v, err := Str(object.NewNativeFunc("test", func(args ...object.Object) (object.Object, error) { return nil, nil }))
		require.NoError(t, err)
		require.Equal(t, "<native-func: test>", v.String())
	})

	t.Run("runtime func", func(t *testing.T) {
		t.Parallel()
		v, err := Str(object.NewRuntimeFunc([]string{}, []parser.IStmtContext{}))
		require.NoError(t, err)
		require.Equal(t, "<runtime-func: <unknown>>", v.String())
	})

	t.Run("more args", func(t *testing.T) {
		t.Parallel()
		v, err := Str(object.NewBool(true), object.NewFloat(2.0))
		require.Error(t, err)
		require.Equal(t, nil, v)
	})
}
