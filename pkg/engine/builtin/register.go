package builtin

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode/utf8"

	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/PicoTools/plan/pkg/engine/storage"
	"github.com/PicoTools/plan/pkg/engine/utils"
)

// Register builtin functions
func Register() {
	// assert: false -> throw exception and stop execution
	registerBuiltin("assert", Assert)
	// print: print string without new line
	registerBuiltin("print", Print)
	// println: print string with new line
	registerBuiltin("println", Println)
	// is_bool: is object of type 'bool'
	registerBuiltin("is_bool", IsBool)
	// is_dict: is object of type 'dict'
	registerBuiltin("is_dict", IsDict)
	// is_float: is object of type 'float'
	registerBuiltin("is_float", IsFloat)
	// is_int: is object of type 'int'
	registerBuiltin("is_int", IsInt)
	// is_list: is object of type 'list'
	registerBuiltin("is_list", IsList)
	// is_null: is object of type 'null'
	registerBuiltin("is_null", IsNull)
	// is_str: is object of type 'str'
	registerBuiltin("is_str", IsStr)
	// bool: cast to 'bool'
	registerBuiltin("bool", Bool)
	// float: cast to 'float'
	registerBuiltin("float", Float)
	// int: cast to 'int'
	registerBuiltin("int", Int)
	// str: cast to 'str'
	registerBuiltin("str", Str)
	// chr: get character based on int code
	registerBuiltin("chr", Chr)
	// ord: get int code of character
	registerBuiltin("ord", Ord)
	// hex: converts string to its hex representation
	registerBuiltin("hex", Hex)
	// unhex: unhexify string
	registerBuiltin("unhex", Unhex)
	// base64_enc: encode string in base64
	registerBuiltin("base64_enc", Base64Enc)
	// base64_dec: decode string from base64
	registerBuiltin("base64_dec", Base64Dec)
	// base32_enc: encode string in base32
	registerBuiltin("base32_enc", Base32Enc)
	// base32_dec: decode string from base32
	registerBuiltin("base32_dec", Base32Dec)
	// md5: get md5 hash of string
	registerBuiltin("md5", Md5)
	// sha1: get sha1 hash of string
	registerBuiltin("sha1", Sha1)
	// sha256: get sha256 hash of string
	registerBuiltin("sha256", Sha256)
	// gzip: gzip data
	registerBuiltin("gzip", Gzip)
	// gunzip: unpack gzipped data
	registerBuiltin("gunzip", Gunzip)
	// fread: read file from FS
	registerBuiltin("fread", Fread)
	// fwrite: write file on FS
	registerBuiltin("fwrite", Fwrite)
}

// registerBuiltin registers builtin function to reduce boilerplate
func registerBuiltin(name string, fn func(args ...object.Object) (object.Object, error)) {
	storage.BuiltinFunctions[name] = object.NewNativeFunc(name, fn)
}

func Assert(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	// get lhs type
	lhs, ok := args[0].(*object.Bool)
	if !ok {
		return nil, fmt.Errorf("expecting bool, got %s", args[0].TypeName())
	}
	if !lhs.Value() {
		return nil, fmt.Errorf("assertion occured")
	}
	return object.NewNull(), nil
}

func Print(args ...object.Object) (object.Object, error) {
	for _, arg := range args {
		fmt.Print(arg.String())
	}
	return object.NewNull(), nil
}

func Println(args ...object.Object) (object.Object, error) {
	for _, arg := range args {
		fmt.Print(arg.String())
	}
	fmt.Println()
	return object.NewNull(), nil
}

func IsBool(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.Bool); ok {
		return object.NewBool(true), nil
	}
	return object.NewBool(false), nil
}

func IsDict(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.Dict); ok {
		return object.NewBool(true), nil
	}
	return object.NewBool(false), nil
}

func IsFloat(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.Float); ok {
		return object.NewBool(true), nil
	}
	return object.NewBool(false), nil
}

func IsInt(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.Int); ok {
		return object.NewBool(true), nil
	}
	return object.NewBool(false), nil
}

func IsList(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.List); ok {
		return object.NewBool(true), nil
	}
	return object.NewBool(false), nil
}

func IsNull(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.Null); ok {
		return object.NewBool(true), nil
	}
	return object.NewBool(false), nil
}

func IsStr(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.Str); ok {
		return object.NewBool(true), nil
	}
	return object.NewBool(false), nil
}

func Bool(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch obj := args[0].(type) {
	case *object.Bool:
		return obj, nil
	case *object.Dict:
		if len(obj.Value()) == 0 {
			return object.NewBool(false), nil
		}
		return object.NewBool(true), nil
	case *object.Float:
		if obj.Value() == 0.0 {
			return object.NewBool(false), nil
		}
		return object.NewBool(true), nil
	case *object.Int:
		if obj.Value() == 0 {
			return object.NewBool(false), nil
		}
		return object.NewBool(true), nil
	case *object.List:
		if len(obj.Value()) == 0 {
			return object.NewBool(false), nil
		}
		return object.NewBool(true), nil
	case *object.Null:
		return object.NewBool(false), nil
	case *object.Str:
		if len(obj.Value()) == 0 {
			return object.NewBool(false), nil
		}
		return object.NewBool(true), nil
	}
	return nil, fmt.Errorf("unknown type '%s'", args[0].TypeName())
}

func Float(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch obj := args[0].(type) {
	case *object.Bool:
		return object.NewFloat(utils.BoolToFloat(obj.Value())), nil
	case *object.Float:
		return obj, nil
	case *object.Int:
		return object.NewFloat(utils.IntToFloat(obj.Value())), nil
	}
	return nil, fmt.Errorf("unknown type '%s'", args[0].TypeName())
}

func Int(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch obj := args[0].(type) {
	case *object.Bool:
		return object.NewInt(utils.BoolToInt(obj.Value())), nil
	case *object.Float:
		return object.NewInt(utils.FloatToInt(obj.Value())), nil
	case *object.Int:
		return obj, nil
	case *object.Str:
		val, err := strconv.Atoi(obj.Value())
		if err != nil {
			return nil, fmt.Errorf("unable convert str to int: %v", err)
		}
		return object.NewInt(int64(val)), nil
	}
	return nil, fmt.Errorf("unknown type '%s'", args[0].TypeName())
}

func Str(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch obj := args[0].(type) {
	case *object.Bool:
		return object.NewStr(obj.String()), nil
	case *object.Dict:
		return object.NewStr(obj.String()), nil
	case *object.Float:
		return object.NewStr(obj.String()), nil
	case *object.Int:
		return object.NewStr(obj.String()), nil
	case *object.List:
		return object.NewStr(obj.String()), nil
	case *object.Null:
		return object.NewStr(obj.String()), nil
	case *object.Str:
		return obj, nil
	}
	return nil, fmt.Errorf("unknown type '%s'", args[0].TypeName())
}

func Chr(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch obj := args[0].(type) {
	case *object.Bool:
		return object.NewStr(string(rune(utils.BoolToInt(obj.Value())))), nil
	case *object.Int:
		return object.NewStr(string(rune(obj.Value()))), nil
	}
	return nil, fmt.Errorf("unknown type '%s'", args[0].TypeName())
}

func Ord(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch obj := args[0].(type) {
	case *object.Str:
		if utf8.RuneCountInString(obj.Value()) != 1 {
			return nil, fmt.Errorf("str must have only one char")
		}
		r, _ := utf8.DecodeRuneInString(obj.Value())
		return object.NewInt(int64(r)), nil
	}
	return nil, fmt.Errorf("unknown type '%s'", args[0].TypeName())
}

func Base64Enc(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	return object.NewStr(base64.StdEncoding.EncodeToString([]byte(str.Value()))), nil
}

func Base64Dec(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	val, err := base64.StdEncoding.DecodeString(str.Value())
	if err != nil {
		return nil, err
	}
	return object.NewStr(string(val)), nil
}

func Base32Enc(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	return object.NewStr(base32.StdEncoding.EncodeToString([]byte(str.Value()))), nil
}

func Base32Dec(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	val, err := base32.StdEncoding.DecodeString(str.Value())
	if err != nil {
		return nil, err
	}
	return object.NewStr(string(val)), nil
}

func Md5(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	md5sum := md5.Sum([]byte(str.Value()))
	return object.NewStr(string(md5sum[:])), nil
}

func Sha1(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	sha1sum := sha1.Sum([]byte(str.Value()))
	return object.NewStr(string(sha1sum[:])), nil
}

func Sha256(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	sha256sum := sha256.Sum256([]byte(str.Value()))
	return object.NewStr(string(sha256sum[:])), nil
}

func Gzip(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	b := &bytes.Buffer{}
	w := gzip.NewWriter(b)
	_, err := w.Write([]byte(str.Value()))
	if err != nil {
		return nil, err
	}
	_ = w.Close()
	return object.NewStr(b.String()), nil
}

func Gunzip(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	r, err := gzip.NewReader(bytes.NewBufferString(str.Value()))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = r.Close()
	}()
	res, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return object.NewStr(string(res)), nil
}

func Fread(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	data, err := os.ReadFile(str.Value())
	if err != nil {
		return nil, err
	}
	return object.NewStr(string(data)), nil
}

func Fwrite(args ...object.Object) (object.Object, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("expecting 2 arguments, got %d", len(args))
	}
	path, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	data, ok := args[1].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 2nd argument, got '%s'", args[1].TypeName())
	}
	if err := os.WriteFile(path.Value(), []byte(data.Value()), 0640); err != nil {
		return nil, err
	}
	return object.NewNull(), nil
}

func Hex(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	return object.NewStr(hex.EncodeToString([]byte(str.Value()))), nil
}

func Unhex(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	v, err := hex.DecodeString(str.Value())
	if err != nil {
		return nil, err
	}
	return object.NewStr(string(v)), nil
}
