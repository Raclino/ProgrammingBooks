# Exercise 1

## Question
Write a program where you declare an integer variable called `i` with the value `20`. Assign `i` to a floating point variable named `f`. Print out `i` and `f`.

## Solution

The key here is that you need to use a type conversion to convert `i` to a `float64` so it can
be assigned. Go doesn't have automatic type promotion. Notice in our code that you don't need to 
specify the type for `i` or `f`; it can be inferred from the right hand side. The default type
for an integer literal is `int`, so `i`'s type can be inferred.  The `float64` type conversion on the right hand side of `f`'s declaration specifies `f`'s type.