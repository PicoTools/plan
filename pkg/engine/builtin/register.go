package builtin

import (
	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/PicoTools/plan/pkg/engine/storage"
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
