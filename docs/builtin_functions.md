# Builtin functions

## assert

Signature: `assert(arg0)`

Arguments:
- `arg0`: `bool`

Throw error if argument if `false`:
```
assert(1 == 1); # no error
assert(1 == 2); # error occured
```

## print

Signature: `print(arg0, arg1, ..., argN)`

Arguments:
- `arg0`: `*`
- `arg1`: `*`
- `argN`: `*`

Print string line in default STDOUT without '\n'. Print will automatically convert each argument in string, using object's interface function `String()`.
```
print("hello", " world"); # hello world
print(1, " ", true); # 1 true
print([0, 1, 2], {"a": "b"}); # [0, 1, 2]{a: b}
```

## println

Signature: `println(arg0, arg1, ..., argN)`

Arguments:
- `arg0`: `*`
- `arg1`: `*`
- `argN`: `*`

Print string line in default STDOUT with '\n'. Print will automatically convert each argument in string, using object's interface function `String()`.
```
println("hello", " world"); # hello world
println(1, " ", true); # 1 true
println([0, 1, 2], {"a": "b"}); # [0, 1, 2]{a: b}
```

## is_bool

Signature: `is_bool(arg0)`

Arguments:
- `arg0`: `*`

Return `true`, if data type of `arg0` is `bool`.
```
println(is_bool("a")); # false
println(is_bool(false)); # true
a = true;
println(is_bool(a)); # true
```

## is_dict

Signature: `is_dict(arg0)`

Arguments:
- `arg0`: `*`

Return `true`, if data type of `arg0` is `dict`.
```
println(is_dict("a")); # false
println(is_dict({"a": true})); # true
```

## is_float

Signature: `is_float(arg0)`

Arguments:
- `arg0`: `*`

Return `true`, if data type of `arg0` is `float`.
```
println(is_float(1)); # false
println(is_float(float(1))); # true
println(is_float(1.0)); # true
```

## is_int

Signature: `is_int(arg0)`

Arguments:
- `arg0`: `*`

Return `true`, if data type of `arg0` is `int`.
```
println(is_int(1)); # true
println(is_int(1.0)); # false
println(is_int(true)); # false
```

## is_list

Signature: `is_list(arg0)`

Arguments:
- `arg0`: `*`

Return `true`, if data type of `arg0` is `list`.
```
println(is_list(null)); # false
println(is_list(["a", 0", true])); # true
```

## is_null

Signature: `is_null(arg0)`

Arguments:
- `arg0`: `*`

Return `true`, if data type of `arg0` is `null`.
```
a = null;
println(is_null(a)); # true
println(is_null([])); # false
```

## is_str

Signature: `is_str(arg0)`

Arguments:
- `arg0`: `*`

Return `true`, if data type of `arg0` is `str`.
```
println(is_str("hello world")); # true
println(is_str(["A"])); # false
```

## bool

Signature: `bool(arg0)`

Arguments:
- `arg0`: `bool`/`dict`/`float`/`int`/`list`/`null`/`str`

Cast `arg0` to `bool`.
```
println(bool(true)); # true
println(bool(false)); # false
println(bool({})); # false
println(bool({"a":"b"})); # true
println(bool(0.0)); # false
println(bool(2.3)); # true
println(bool(0)); # false
println(bool(1)); # true
println(bool([])); # false
println(bool(["a", "b"])); # true
println(bool(null)); # false
println(bool("")); # false
println(bool("hello")); # true
```

## float

Signature: `float(arg0)`

Arguments:
- `arg0`: `bool`/`float`/`int`

Cast `arg0` to `float`.
```
println(float(true)); # 1
println(float(false)); # 0
println(float(1.0)); # 1
println(float(23.91)); # 23.91
println(float(0)); # 0
println(float(29)); # 29
```

## int

Signature: `int(arg0)`

Arguments:
- `arg0`: `bool`/`float`/`int`/`str`

Cast `arg0` to `int`.
```
println(int(true)); # 1
println(int(false)); # 0
println(int(1.5)); # 1
println(int(1.4)); # 1
println(int(123)); # 123
println(int("2024")); # 2024
```

## str

Signature: `str(arg0)`

Arguments:
- `arg0`: `*`

Cast `arg0` to `str`. Will use underly `String()` implementation of each type.

```
println(true); # true
println([1]); # [1]
println(fn () {}); # <native-func>
```

## chr

Signature: `chr(arg0)`

Arguments:
- `arg0`: `bool`/`int`

Returns character with type `str` associated with value of `arg0`.

```
println(chr(81)); # Q
println(chr(181)); # µ
println(chr(1181)); # ҝ
```

## ord

Signature: `ord(arg0)`

Arguments:
- `arg0`: `str`

Returns code with type `int` associated with `arg0` character. Length of `arg0` must be "1 rune".

```
println(ord("A")); # 65
println(ord("П")); # 1055
```

## hex

Signature: `hex(arg0)`

Arguments:
- `arg0`: `str`

Hexify `arg0`.

```
println(hex("hello")); # 68656c6c6f
println(hex("\x01\x02\x03")); # 010203
```

## unhex

Signature: `unhex(arg0)`

Arguments:
- `arg0`: `str`

Unhexify `arg0`.

```
println(unhex("68656c6c6f")); # hello
```

# base64_enc

Signature: `base64_enc(arg0)`

Arguments:
- `arg0`: `str`

Encodes `arg0` to base64 `str`.

```
println(base64_enc("hello")); # aGVsbG8=
println(base64_enc("привет")); # 0L/RgNC40LLQtdGC
```

# base64_dec

Signature: `base64_dec(arg0)`

Arguments:
- `arg0`: `str`

Decodes `arg0` from base64 to raw `str`.

```
println(base64_dec("aGVsbG8=")); # hello
println(base64_dec("0L/RgNC40LLQtdGC")); # привет
```

# base32_enc

Signature: `base32_enc(arg0)`

Arguments:
- `arg0`: `str`

Encodes `arg0` to base64 `str`.

```
println(base32_enc("hello")); # NBSWY3DP
println(base32_enc("привет")); # 2C75DAGQXDILFUFV2GBA====
```

# base32_dec

Signature: `base32_dec(arg0)`

Arguments:
- `arg0`: `str`

Decodes `arg0` from base64 to raw `str`.

```
println(base32_dec("NBSWY3DP")); # hello
println(base32_dec("2C75DAGQXDILFUFV2GBA====")); # привет
```

# md5

Signature: `md5(arg0)`

Arguments:
- `arg0`: `str`

Returns hex representation (bytes in `str`) of MD5 checksum algorithm on `arg0`.

```
println(hex(md5("hello"))); # 5d41402abc4b2a76b9719d911017c592
println(hex(md5("привет"))); # 608333adc72f545078ede3aad71bfe74
```

# sha1

Signature: `sha1(arg0)`

Arguments:
- `arg0`: `str`

Returns hex representation (bytes in `str`) of SHA1 checksum algorithm on `arg0`.

```
println(hex(sha1("hello"))); # aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d
println(hex(sha1("привет"))); # e24505f94db2b5df4c7c2596b0788e720e073021
```

# sha256

Signature: `sha1(arg0)`

Arguments:
- `arg0`: `str`

Returns hex representation (bytes in `str`) of SHA256 checksum algorithm on `arg0`.

```
println(hex(sha256("hello"))); # 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
println(hex(sha256("привет"))); # e58f1e8c55fa105bdd3f40e5037eb0b039b5998d52c05e6cd98878dd2da5cab2
```

# gzip

Signature: `gzip(arg0)`

Arguments:
- `arg0`: `str`

Returns gzipped value of `arg0` as bytes in `str`.

```
println(hex(gzip("hello"))); # 1f8b08000000000000ffca48cdc9c907040000ffff86a6103605000000
```

# gunzip

Signature: `gunzip(arg0)`

Arguments:
- `arg0`: `str`

Retuns gunzipped value of `arg0` as bytes in `str`.

```
println(gunzip(unhex("1f8b08000000000000ffca48cdc9c907040000ffff86a6103605000000"))); # hello
```

# fread

Signature: `fread(arg0)`

Arguments:
- `arg0`: `str`

Read bytes in `str` from file specified by `arg0`.

```
println(fread("/etc/hostname")); # plan.example.com
```

# fwrite

Signature: `fwrite(arg0, arg1)`

Arguments:
- `arg0`: `str`
- `arg1`: `str`

Write bytes from `arg1` to file specified by `arg0`.

```python
fwrite("/tmp/test", "hello\n");
println(fread("/tmp/test")); # hello
```