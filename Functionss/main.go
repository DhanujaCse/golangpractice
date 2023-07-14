package main

import (
	"fmt"
)

func CallAnyOneFunc1(f func(int, int) int, a, b int) {
	fmt.Println(f(a, b))
}

func Area1(a, b int) int {
	return a * b
}
func Sum1(a, b int) int {
	return a - b
}

func main() {

	// clode := ClosureExample()
	// clode(6)
	// fmt.Println(clode(5))

	CallAnyOneFunc1(Area1, 8, 8)
	callAnyOneFunc(area, 6, 6)
}
