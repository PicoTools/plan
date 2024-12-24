# PLAN

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

PLAN (Pico LANguage) is a scripting language (interpreter based) for embedding in Golang applications based on ANTLR4 grammar.

## Why yet another scripting language?

Development of language was part of creation C2 framework - [pico](https://github.com/PicoTools/piko), which we position as interface between operator and beacon. PLAN runtime used in [operator's CLI](https://github.com/PicoTools/c2mcli) to give possibility of extending stocked functionality of interaction with server and beacons, which we provided.

Speed was not a goal of this language, but it can be easily reached by mapping functions written on native Golang.

## What about syntax?

We can charactarized PLAN as follows: syntax more like `C` way, typification and logic of processing like `Python` way. 

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

## What about data types?

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

## What's next?

- [ ] Improve quality of documentation
- [x] Builtin methods for data types
- [ ] Builtin functions to handle regexp
- [ ] Builtin functions to handle network interaction
- [ ] Extend builtin functions to handle FS interaction
- [ ] Extend builtin functions for string processing
- [ ] Tests of seprated critical components
- [ ] Extend collection of samples with realization of different algos