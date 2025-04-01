# PLAN syntax

PLAN syntax resembles `C` in structure (brackets and semicolons), while its typing system and execution logic are more similar to `Python`.

## Values

PLAN uses dynamic typification, so variable can be set by all of data types:

```python
a = false; # bool
a = {"a": "b"}; # dict
a = 1.0; # float
a = 23; # int
a = [1, 2, 3]; # list
a = null; # null
a = "hello"; # str
a = fn () { return "test"; }; # closure
```

## Indexer

Indexer `[]` can be used to access values/objects from `str`/`dict`/`list` data types. 

**In case of `str` logic of indexer use runes instead of raw bytes.**

```python
list_value = [1, 2, 3, 4];
println(list_value[1]); # 2

dict_value = {"a": 123, "b": list_value};
println(dict_value[1]); # [1, 2, 3, 4]

str_value = "привет";
println(str_value[1]); # и
```

Also supports multidemension indexers:
```python
test = [[[1], 2], 3];
println(test[0][0][0]); # 1
```

## Funtions

Functions, defined in PLAN context, stored in runtime using [RuntimeFunc](../pkg/engine/object/runtime_func.go) type.

Definition of functions looks like that:
```python
fn test01(a, b, c) {
    return a + b + c;
}

fn test02() {
    a = 1;
}
```

In case of obvious `return` statement, function will return object with specific data type. Otherwise `null` will be returned.
```python
println(test01(1, 2, 3)); # 6
println(test02()); # <null>
```

You can create nested functions that will be available inside of the nested scope. For example:
```python
# function in scope #0
fn hello() {
    # function in scope #1
    fn hello() {
        return "hello";
    }
    # function in scope #1
    fn world() {
        return "world";
    }

    return hello() + " " + world();
}
hello(); // hello world
```

## Closures

Closures can be used to store functions in variables and call it later.
```python
my_closure = fn(a, b, c) {
    return a * b + c;
};

result = @my_closure(2, 2, 2);
println(result); # 6
```

As closure stores in variable it can be reassigned:
```python
a = fn(a) {
    return a;
};
b = a;
println(@b(1)); # 1
```

## `if` statement

`if` statement looks like that:
```python
a = 3;

if a == 0 {
    println("A");
} elif a == 2 {
    println("B");
} elif a == 3 {
    println("C"); # will be printed
} else {
    println("D");
}
```

> :warning: `if`/`elif` conditions will try to automatically convert data type to bool.
```python
if "hello" {
    println("A"); # will be printed
}

if 23 {
    println("B"); # will be printed
}

if [] {
    println("C"); # will NOT be printed
}
```

## `for` statement

`for` statement looks like that:
```python
for i = 0; i < 5; i += 1 {
    println(i);
}
```

In case of `for` loop, `continue` and `break` can be used to change control flow.
```python
a = 3;
for i = 0; i < 10; i += 1 {
    if i == a {
        break;
    }
}
```

## `while` statement

`while` statement looks like that:
```python
a = 10;
while a > 0 {
    println(a);
    a -= 1;
}
```

In case of `while` loop, `continue` and `break` can be used to change control flow.
```python
a = 5;
while a > 0 {
    a -= 1;
    if a == 4 {
        continue;
    }
    println(a);
}
```

## `include` statement

PLAN offered with `include` statement, which support including of another scripts.

Place in `add.pico` next code:
```python
fn add(a, b) {
    return a + b;
}
```

In `main.pico` include script and run function `add()`.
```python
include("add.pico");

println(add(1, 2)); # 3
```

## Code comments

PLAN supports comments in code using next notations:
```c
// this is my awesome comment

# this is my awesome comment too

/*
    this is
    my myltiline
    awesome
    comment
*/
```