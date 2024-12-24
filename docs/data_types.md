# PLAN data types

All data types of language sit on abstraction (interface) called [`object`](../pkg/engine/object/object.go). We can divide all types in 3 groups:
- `essential`
- `composite`
- `functions`

Moreover data types has builtin methods. More information you can find [here](./builtin_methods.md).

## Essential data types

- [bool](../pkg/engine/object/bool.go) (`bool` type in Golang)
- [float](../pkg/engine/object/float.go) (`float64` type in Golang)
- [int](../pkg/engine/object/int.go) (`int64` type in Golang)
- [null](../pkg/engine/object/null.go) (empty object in Golang)
- [str](../pkg/engine/object/str.go) (`string` in Golang, process with runes)

## Composite data types

- [dict](../pkg/engine/object/dict.go) (`map` in Golang, where key is always `string`)
- [list](../pkg/engine/object/list.go) (`list` in Golang, can hold all object's types)

## Functions

There are 2 underly realization for different purposes:
- [`NativeFunc`](../pkg/engine/object/native_func.go): used to hold Golang defined functions (builtins and users).
- [`RuntimeFunc`](../pkg/engine/object/runtime_func.go): used to hold PLAN defined functions/closures in runtime. Map with this type of functions filled in process of visiting AST.

## Casting rules

In process of PLAN development the most part of casting rules has been taken from `Python`:

|       | bool  | dict | float | int   | list | null | str  |
|-------|-------|------|-------|-------|------|------|------|
| bool  | -     | bool | bool  | bool  | bool | bool | bool |
| dict  | x     | -    | x     | x     | x    | x    | x    |
| float | float | x    | -     | float | x    | x    | x    |
| int   | int   | x    | int   | -     | x    | x    | int* |
| list  | x     | x    | x     | x     | -    | x    | x    |
| null  | x     | x    | x     | x     | x    | -    | x    |
| str   | str   | str  | str   | str   | str  | str  | -    |

`-`: will return same object, as was passed for casting
`x`: will be thrown error and execution stopped
`int*`: `strconv.Atoi()` under the hood

Casting realized with using of next [builtin](../pkg/engine/builtin/register.go) functions:
- `bool(v)`: return `bool` if casting appliable
- `float(v)`: return `float` if casting appliable
- `int(v)`: return `int` if casting appliable
- `str(v)`: return `str` if casting appliable

To avoid runtime errors you can use next [builtin](../pkg/engine/builtin/register.go) function to determine data type:
- `is_bool(v)`: return `true` if `bool`
- `is_dict(v)`: return `true` if `dict`
- `is_float(v)`: return `true` if `float`
- `is_int(v)`: return `true` if `int`
- `is_list(v)`: return `true` if `list`
- `is_null(v)`: return `true` if `null`
- `is_str(v)`: return `true` if `str`

Full list of current builtin functions you can found [here](./builtin_functions.md).
