# PLAN

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

PLAN (Pico LANguage) is an interpreter-based scripting language for embedding in Golang applications created on the ANTLR4 grammar.

## Why yet another scripting language?

The development of the language was part of the creation of the C2 framework - [pico](https://github.com/PicoTools/pico), which we position as an interface between the operator and the agent. PLAN is used in the [operator CLI](https://github.com/PicoTools/pico-cli) to provide the ability to expand the existing functionality of interaction with the server and agents that we have provided.

Speed was not the goal of this language, but it can be easily achieved with display features written in native Golang.

## Syntax

We can charactarized PLAN as follows: syntax more like `C`, typification and logic of processing like `Python`. 

This example is `C` like (a lot of brakets and semicolons):
```
fn fib(n) {
  if n < 2 {
    return n;
  } else {
    return fib(n-2) + fib(n-1);
  }
}

println(fib(35));
```

This example is `Python` like (dynimac typification of variable `a`):
```
a = 1;
a = [a, 2, 3, 4];
a = {"kek": a};
a = "hello";
a = a + " world";
println(a);
```

More info about [syntax](docs/syntax.md) with examples.

## Data types

All data types sit on top of abstration as `object`.

PLAN offered with next data types, which implements object's interface:
- `bool` (bool in Golang)
- `dict` (map in Golang, where key is string)
- `float` (float64 in Golang)
- `int` (int64 in Golang)
- `list` (list of objects)
- `null` (just empty object)
- `str` (string in Golang, but operates using runes)
- `closures` (special type of object, that can be called inline)

More info about [data types](docs/data_types.md) with examples.

## References

Here you can find some more information:
- [Syntax](docs/syntax.md)
- [Data types](docs/data_types.md)
- [Builtin functions](docs/builtin_functions.md)
- [Builtin methods](docs/builtin_methods.md)
