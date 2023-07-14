package main

import (
	"fmt"
)

func ClosureExample() func(int) int {
	return func(i int) int {
		fmt.Println("printed")
		return i * i

	}
}
