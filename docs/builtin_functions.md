# Builtin functions

## assert

Signature: `assert(arg0)`

Arguments:
- `arg0`: `bool`

Throw error if argument if `false`:
```
assert(1 == 1); // no error
assert(1 == 2); // error occured
```

## print

Signature: `print(arg0, arg1, ..., argN)`

Arguments:
- `arg0`: `*`
- `arg1`: `*`
- `argN`: `*`

Print string line in default STDOUT without '\n'. Print will automatically convert each argument in string, using object's interface function `String()`.
```
print("hello", " world"); // hello world
print(1, " ", true); // 1 true
print([0, 1, 2], {"a": "b"}); // [0, 1, 2]{a: b}
```

## println

Signature: `println(arg0, arg1, ..., argN)`

Arguments:
- `arg0`: `*`
- `arg1`: `*`
- `argN`: `*`

Print string line in default STDOUT with '\n'. Print will automatically convert each argument in string, using object's interface function `String()`.
```
println("hello", " world"); // hello world
println(1, " ", true); // 1 true
println([0, 1, 2], {"a": "b"}); // [0, 1, 2]{a: b}
```

## is_bool

Signature: `is_bool(arg0)`

Arguments:
- `arg0`: `*`

Return `true`, if data type of `arg0` is `bool`.
```
println(is_bool("a")); // false
println(is_bool(false)); // true
a = true;
println(is_bool(a)); // true
```

## is_dict

Signature: `is_dict(arg0)`

Arguments:
- `arg0`: `*`

Return `true`, if data type of `arg0` is `dict`.
```
println(is_dict("a")); // false
println(is_dict({"a": true})); // true
```

## is_float

Signature: `is_float(arg0)`

Arguments:
- `arg0`: `*`

Return `true`, if data type of `arg0` is `float`.
```
println(is_float(1)); // false
println(is_float(float(1))); // true
println(is_float(1.0)); // true
```

## is_int

Signature: `is_int(arg0)`

Arguments:
- `arg0`: `*`

Return `true`, if data type of `arg0` is `int`.
```
println(is_int(1)); // true
println(is_int(1.0)); // false
println(is_int(true)); // false
```

## is_list

Signature: `is_list(arg0)`

Arguments:
- `arg0`: `*`

Return `true`, if data type of `arg0` is `list`.
```
println(is_list(null)); // false
println(is_list(["a", 0", true])); // true
```

## is_null

Signature: `is_null(arg0)`

Arguments:
- `arg0`: `*`

Return `true`, if data type of `arg0` is `null`.
```
a = null;
println(is_null(a)); // true
println(is_null([])); // false
```

## is_str

Signature: `is_str(arg0)`

Arguments:
- `arg0`: `*`

Return `true`, if data type of `arg0` is `str`.
```
println(is_str("hello world")); // true
println(is_str(["A"])); // false
```

## bool

Signature: `bool(arg0)`

Arguments:
- `arg0`: `bool`/`dict`/`float`/`int`/`list`/`null`/`str`

Cast `arg0` to `bool`.
```
println(bool(true)); // true
println(bool(false)); // false
println(bool({})); // false
println(bool({"a":"b"})); // true
println(bool(0.0)); // false
println(bool(2.3)); // true
println(bool(0)); // false
println(bool(1)); // true
println(bool([])); // false
println(bool(["a", "b"])); // true
println(bool(null)); // false
println(bool("")); // false
println(bool("hello")); // true
```

## float

Signature: `float(arg0)`

Arguments:
- `arg0`: `bool`/`float`/`int`

Cast `arg0` to `float`.
```
println(float(true)); // 1
println(float(false)); // 0
println(float(1.0)); // 1
println(float(23.91)); // 23.91
println(float(0)); // 0
println(float(29)); // 29
```

## int

Signature: `int(arg0)`

Arguments:
- `arg0`: `bool`/`float`/`int`/`str`

Cast `arg0` to `int`.
```
println(int(true)); // 1
println(int(false)); // 0
println(int(1.5)); // 1
println(int(1.4)); // 1
println(int(123)); // 123
println(int("2024")); // 2024
```

## str

Signature: `str(arg0)`

Arguments:
- `arg0`: `*`

Cast `arg0` to `str`. Will use underly `String()` implementation of each type.

```
println(true); // true
println([1]); // [1]
println(fn () {}); // <native-func>
```

## chr