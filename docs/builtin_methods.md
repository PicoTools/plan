# Builtin methods

## dict.len

Signature: `dict.len()`

Arguments: <none>

Get number of keys in `dict`.
```
a = {"a": 0, "b": 1};
println(a.len()); // 2
```

## dict.pop

Signature: `dict.pop(arg0)`

Arguments:
- `arg0`: `str`

Pop value from `dict` by key with `str` type.
```
a = {"a": 0, "b": 1};
println(a); // {a: 0, b: 1}
a.pop("a");
println(a); //  {b: 1}
```

## list.len

Signature: `list.len()`

Arguments: <none>

Get number of objects in `list`.
```
a = [1, 2, 3, 4, 5];
println(a.len()); // 5
```

## list.reverse

Signature: `list.reverse()`

Arguments: <none>

Reverse order of objects in `list`.
```
a = [1, 2, 3, 4, 5];
a.reverse();
println(a); // [5, 4, 3, 2, 1]
```

## list.pop

Signature: `list.pop(arg0)`

Arguments:
- `arg0`: `int`

Pop value from `list` by index with `int` type.
```
a = [1, 2, 3, 4, 5];
a.pop(0);
a.pop(1);
println(a); // [2, 4, 5]
```

## str.len

Signature: `str.len()`

Arguments: <none>

Get number of characters (Golang runes) in `str`.
```
a = "привет";
b = "hello";
println(a.len()); // 6
println(b.len()); // 5
```

## str.reverse

Signature: `list.reverse()`

Arguments: <none>

Reverse order of characters (Golang runes) in `str`.
```
a = "привет";
b = "hello";
a.reverse();
b.reverse();
println(a); // тевирп
println(b); // olleh
```

## str.split

Signature: `str.split(arg0)`

Arguments:
- `arg0`: `str`

Split `str` by delimeter (underhood will split by rune). Result will be `list`.
```
a = "hello world";
b = a.split(" ");
println(b); // [hello, world]

a = "привет мир";
b = a.split("и");
println(b); // [пр, вет м, р]
```