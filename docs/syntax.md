# PLAN syntax

We can charactarized PLAN as follows: syntax more like `C` way, typification and logic of processing like `Python` way.

## Values

PLAN is dynamically typification, so variable can be set by all of data types:
```
a = false; // bool
a = {"a": "b"}; // dict
a = 1.0; // float
a = 23; // int
a = [1, 2, 3]; // list
a = null; // null
a = "hello"; // str
a = fn () { return "test"; }; // closure
```

## Indexer

Indexer `[]` can be used to access values/objects from `str`/`dict`/`list` data types. 

**In case of `str` logic of indexer use runes instead of raw bytes.**
```
list_value = [1, 2, 3, 4];
println(list_value[1]); // 2

dict_value = {"a": 123, "b": list_value};
println(dict_value[1]); // [1, 2, 3, 4]

str_value = "привет";
println(str_value[1]); // и
```


Also supports multidemension indexers:
```
test = [[[1], 2], 3];
println(test[0][0][0]); // 1
```

## Funtions

Functions, defined in PLAN context, stored in runtime using [RuntimeFunc](../pkg/engine/object/runtime_func.go) type.

Definition of functions looks like that:
```
fn test01(a, b, c) {
    return a + b + c;
}

fn test02() {
    a = 1;
}
```

In case of obvious `return` statement, function will return object with specific data type. Otherwise `null` will be returned.
```
println(test01(1, 2, 3)); // 6
println(test02()); // <null>
```

## Closures

Closures can be used to store functions in variables and call it later.
```
my_closure = fn(a, b, c) {
    return a * b + c;
};

result = @my_closure(2, 2, 2);
println(result); // 6
```

As closure stores in variable it can be reassigned:
```
a = fn(a) {
    return a;
};
b = a;
println(@b(1)); // 1
```

## `if` statement

`if` statement looks like that:
```
a = 3;

if a == 0 {
    println("A");
} elif a == 2 {
    println("B");
} elif a == 3 {
    println("C"); // will be printed
} else {
    println("D");
}
```

`if`/`elif` block expect `bool` condition.

## `for` statement

`for` statement looks like that:
```
for i = 0; i < 5; i += 1 {
    println(i);
}
```

In case of `for` loop, `continue` and `break` can be used to change control flow.
```
a = 3;
for i = 0; i < 10; i += 1 {
    if i == a {
        break;
    }
}
```

## `while` statement

`while` statement looks like that:
```
a = 10;
while a > 0 {
    println(a);
    a -= 1;
}
```

In case of `while` loop, `continue` and `break` can be used to change control flow.
```
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
```
fn add(a, b) {
    return a + b;
}
```

In `main.pico` include script and run function `add()`.
```
include("add.pico");

println(add(1, 2)); // 3
```

## Code comments

PLAN support comments in code using next notations:
```
// this is my awesome comment

# this is my awesome comment too

/*
    this is
    my myltiline
    awesome
    comment
*/
```