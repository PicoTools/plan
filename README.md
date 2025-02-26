# PLAN

PLAN (Pico LANguage) is a scripting language (interpreter-based) for embedding in Golang applications, based on ANTLR4 grammar.

## Why yet another scripting language?

The development of this language was part of the creation of a C2 framework — [Pico](https://github.com/PicoTools/pico), which we position as an interface between the operator and the beacon.  

The PLAN runtime is used in the [operator's CLI](https://github.com/PicoTools/pico-cli) to extend the built-in functionality of interactions with the server and beacons.  

Performance was not the primary goal of this language, but it can be easily improved by mapping functions written in native Golang.

## Syntax

PLAN syntax resembles C in structure (brackets and semicolons), while its typing system and execution logic are more similar to Python.

A C-like example (explicit brackets and semicolons):
```c
fn fib(n) {
  if n < 2 {
    return n;
  } else {
    return fib(n-2) + fib(n-1);
  }
}

println(fib(35));
```

A Python-like example (dynamic typing of variable a):
```python
a = 1;
a = [a, 2, 3, 4];
a = {"test": a};
a = "hello";
a = a + " world";
println(a);
```

For more information and examples, see [syntax](docs/syntax.md).

## Data Types

All data types sit on top of an abstraction called object.

PLAN comes with the following data types, which implement the object interface:
- `bool` (bool in Golang)
- `dict` (map in Golang, where the key is a string)
- `float` (float64 in Golang)
- `int` (int64 in Golang)
- `list` (a list of objects)
- `null` (an empty object)
- `str` (a string in Golang, operating with runes)
- `closures` (a special type of object that can be called inline)

For more information and examples, see data [data types](docs/data_types.md).

## References

Here you can find additional information:
- [Syntax](docs/syntax.md)
- [Data types](docs/data_types.md)
- [Builtin functions](docs/builtin_functions.md)
- [Builtin methods](docs/builtin_methods.md)

## TODO

- [ ] Improve documentation quality
- [x] Built-in methods for data types
- [ ] Built-in functions for handling regex
- [ ] Built-in functions for network interactions
- [ ] Extend built-in functions for filesystem interactions
- [ ] Extend built-in functions for string processing
- [ ] Testing of critical components separately
- [ ] Expand the collection of examples with different algorithm implementations
