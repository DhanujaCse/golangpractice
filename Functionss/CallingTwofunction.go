package main

import (
	"fmt"
)

func callAnyOneFunc(f func(int, int) int, a, b int) {
	fmt.Println(f(a, b))
}

func area(a, b int) int {
	return a * b
}
func Sum(a, b int) int {
	return a - b
}
