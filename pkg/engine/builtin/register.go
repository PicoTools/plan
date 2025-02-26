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
	storage.BuiltinFunctions["assert"] = object.NewNativeFunc("assert", Assert)
	// print: print string without new line
	storage.BuiltinFunctions["print"] = object.NewNativeFunc("print", Print)
	// println: print string with new line
	storage.BuiltinFunctions["println"] = object.NewNativeFunc("println", Println)
	// is_bool: is object of type 'bool'
	storage.BuiltinFunctions["is_bool"] = object.NewNativeFunc("is_bool", IsBool)
	// is_dict: is object of type 'dict'
	storage.BuiltinFunctions["is_dict"] = object.NewNativeFunc("is_dict", IsDict)
	// is_float: is object of type 'float'
	storage.BuiltinFunctions["is_float"] = object.NewNativeFunc("is_float", IsFloat)
	// is_int: is object of type 'int'
	storage.BuiltinFunctions["is_int"] = object.NewNativeFunc("is_int", IsInt)
	// is_list: is object of type 'list'
	storage.BuiltinFunctions["is_list"] = object.NewNativeFunc("is_list", IsList)
	// is_null: is object of type 'null'
	storage.BuiltinFunctions["is_null"] = object.NewNativeFunc("is_null", IsNull)
	// is_str: is object of type 'str'
	storage.BuiltinFunctions["is_str"] = object.NewNativeFunc("is_str", IsStr)
	// bool: cast to 'bool'
	storage.BuiltinFunctions["bool"] = object.NewNativeFunc("bool", Bool)
	// float: cast to 'float'
	storage.BuiltinFunctions["float"] = object.NewNativeFunc("float", Float)
	// int: cast to 'int'
	storage.BuiltinFunctions["int"] = object.NewNativeFunc("int", Int)
	// str: cast to 'str'
	storage.BuiltinFunctions["str"] = object.NewNativeFunc("str", Str)
	// chr: get character based on int code
	storage.BuiltinFunctions["chr"] = object.NewNativeFunc("chr", Chr)
	// ord: get int code of character
	storage.BuiltinFunctions["ord"] = object.NewNativeFunc("ord", Ord)
	// hex: converts string to its hex representation
	storage.BuiltinFunctions["hex"] = object.NewNativeFunc("hex", Hex)
	// unhex: unhexify string
	storage.BuiltinFunctions["unhex"] = object.NewNativeFunc("unhex", Unhex)
	// base64_enc: encode string in base64
	storage.BuiltinFunctions["base64_enc"] = object.NewNativeFunc("base64_enc", Base64Enc)
	// base64_dec: decode string from base64
	storage.BuiltinFunctions["base64_dec"] = object.NewNativeFunc("base64_dec", Base64Dec)
	// base32_enc: encode string in base32
	storage.BuiltinFunctions["base32_enc"] = object.NewNativeFunc("base32_enc", Base32Enc)
	// base32_dec: decode string from base32
	storage.BuiltinFunctions["base32_dec"] = object.NewNativeFunc("base32_dec", Base32Dec)
	// md5: get md5 hash of string
	storage.BuiltinFunctions["md5"] = object.NewNativeFunc("md5", Md5)
	// sha1: get sha1 hash of string
	storage.BuiltinFunctions["sha1"] = object.NewNativeFunc("sha1", Sha1)
	// sha256: get sha256 hash of string
	storage.BuiltinFunctions["sha256"] = object.NewNativeFunc("sha256", Sha256)
	// gzip: gzip data
	storage.BuiltinFunctions["gzip"] = object.NewNativeFunc("gzip", Gzip)
	// gunzip: unpack gzipped data
	storage.BuiltinFunctions["gunzip"] = object.NewNativeFunc("gunzip", Gunzip)
	// fread: read file from FS
	storage.BuiltinFunctions["fread"] = object.NewNativeFunc("fread", Fread)
	// fwrite: write file on FS
	storage.BuiltinFunctions["fwrite"] = object.NewNativeFunc("fwrite", Fwrite)
}

func Assert(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(*object.Bool); !ok {
		return nil, fmt.Errorf("expecting bool, got %s", args[0].TypeName())
	}
	if !args[0].(*object.Bool).GetValue().(bool) {
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
	switch args[0].(type) {
	case *object.Bool:
		return args[0], nil
	case *object.Dict:
		if len(args[0].(*object.Dict).GetValue().(map[string]object.Object)) == 0 {
			return object.NewBool(false), nil
		}
		return object.NewBool(true), nil
	case *object.Float:
		if args[0].(*object.Float).GetValue().(float64) == 0.0 {
			return object.NewBool(false), nil
		}
		return object.NewBool(true), nil
	case *object.Int:
		if args[0].(*object.Int).GetValue().(int64) == 0 {
			return object.NewBool(false), nil
		}
		return object.NewBool(true), nil
	case *object.List:
		if len(args[0].(*object.List).GetValue().([]object.Object)) == 0 {
			return object.NewBool(false), nil
		}
		return object.NewBool(true), nil
	case *object.Null:
		return object.NewBool(false), nil
	case *object.Str:
		if len(args[0].(*object.Str).GetValue().(string)) == 0 {
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
	switch args[0].(type) {
	case *object.Bool:
		return object.NewFloat(utils.BoolToFloat(args[0].(*object.Bool).GetValue().(bool))), nil
	case *object.Float:
		return args[0], nil
	case *object.Int:
		return object.NewFloat(utils.IntToFloat(args[0].(*object.Int).GetValue().(int64))), nil
	}
	return nil, fmt.Errorf("unknown type '%s'", args[0].TypeName())
}

func Int(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch args[0].(type) {
	case *object.Bool:
		return object.NewInt(utils.BoolToInt(args[0].(*object.Bool).GetValue().(bool))), nil
	case *object.Float:
		return object.NewInt(utils.FloatToInt(args[0].(*object.Float).GetValue().(float64))), nil
	case *object.Int:
		return args[0], nil
	case *object.Str:
		val, err := strconv.Atoi(args[0].(*object.Str).GetValue().(string))
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
	switch args[0].(type) {
	case *object.Bool:
		return object.NewStr(args[0].(*object.Bool).String()), nil
	case *object.Dict:
		return object.NewStr(args[0].(*object.Dict).String()), nil
	case *object.Float:
		return object.NewStr(args[0].(*object.Float).String()), nil
	case *object.Int:
		return object.NewStr(args[0].(*object.Int).String()), nil
	case *object.List:
		return object.NewStr(args[0].(*object.List).String()), nil
	case *object.Null:
		return object.NewStr(args[0].(*object.Null).String()), nil
	case *object.Str:
		return args[0], nil
	}
	return nil, fmt.Errorf("unknown type '%s'", args[0].TypeName())
}

func Chr(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch args[0].(type) {
	case *object.Bool:
		return object.NewStr(string(rune(utils.BoolToInt(args[0].(*object.Bool).GetValue().(bool))))), nil
	case *object.Int:
		return object.NewStr(string(rune(args[0].(*object.Int).GetValue().(int64)))), nil
	}
	return nil, fmt.Errorf("unknown type '%s'", args[0].TypeName())
}

func Ord(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	switch args[0].(type) {
	case *object.Str:
		if utf8.RuneCountInString(args[0].(*object.Str).GetValue().(string)) != 1 {
			return nil, fmt.Errorf("str must have only one char")
		}
		r, _ := utf8.DecodeRuneInString(args[0].(*object.Str).GetValue().(string))
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
	return object.NewStr(base64.StdEncoding.EncodeToString([]byte(str.GetValue().(string)))), nil
}

func Base64Dec(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	val, err := base64.StdEncoding.DecodeString(str.GetValue().(string))
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
	return object.NewStr(base32.StdEncoding.EncodeToString([]byte(str.GetValue().(string)))), nil
}

func Base32Dec(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	val, err := base32.StdEncoding.DecodeString(str.GetValue().(string))
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
	md5sum := md5.Sum([]byte(str.GetValue().(string)))
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
	sha1sum := sha1.Sum([]byte(str.GetValue().(string)))
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
	sha256sum := sha256.Sum256([]byte(str.GetValue().(string)))
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
	_, err := w.Write([]byte(str.GetValue().(string)))
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
	r, err := gzip.NewReader(bytes.NewBufferString(str.GetValue().(string)))
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
	data, err := os.ReadFile(str.GetValue().(string))
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
	if err := os.WriteFile(path.GetValue().(string), []byte(data.GetValue().(string)), 0640); err != nil {
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
	return object.NewStr(hex.EncodeToString([]byte(str.GetValue().(string)))), nil
}

func Unhex(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	str, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	v, err := hex.DecodeString(str.GetValue().(string))
	if err != nil {
		return nil, err
	}
	return object.NewStr(string(v)), nil
}
