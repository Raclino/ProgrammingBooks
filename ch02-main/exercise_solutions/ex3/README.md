# Exercise 3

## Question
Write a program with three variables, one named `b` of type `byte`, one named `smallI` of type `int32`, and one named `bigI` of type `uint64`. Assign each variable the maximum legal value for its type, then add `1` to each variable. Print out their values.

## Solution

The maximum values for each type is specified in the table in Chapter 2. If you are industrious, you can also find the constants declared in the standard library in the `math` package.

Adding 1 to each variable causes an overflow, not an error. You get the minimum value for each one. 