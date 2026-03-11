// Package ch10_solution contains the solution to the exercises in chapter 10 of Learning Go, 2nd edition.
package ch10_solution

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add is a function that takes in two numbers values and returns their sum.
// It follows the rules specified by https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}

