package main

import "golang.org/x/exp/constraints"

type Number interface {
    constraints.Integer | constraints.Float
}

// Add adds two input integers and returns the sum of them
//
// For more information, refer to [https://www.mathsisfun.com/numbers/addition.html]
func Add[T Number](a, b T) T {
    return a + b
}

