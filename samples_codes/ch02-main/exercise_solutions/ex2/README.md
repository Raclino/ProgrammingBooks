# Exercise 2

## Question
Write a program where you declare a constant called `value` that can be assigned to both an integer and a floating point variable. Assign it to an integer called `i` and a floating point variable called `f`. Print out `i` and `f`.

## Solution

To solve this, make `value` an untyped constant that's assigned an integer literal value. You can then assign it to both an `int` and a `float64`. Since the value of `value` is an integer literal (which has a default type of `int`), you do not need to specify the type for `i` and can use the short declaration format. However, to make `f` a `float64`, you need to specify its type.

