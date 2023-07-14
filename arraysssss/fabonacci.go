package arrayws

import "fmt"

func FabonacciSeries() {
	num1, num2 := 0, 1
	var next, n int

	fmt.Println("enter the number of terms ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		if i == 0 {
			fmt.Print(" ", num1)
		}
		if i == 1 {
			fmt.Print(" ", num2)

		} else {
			next = num1 + num2
			num1 = num2
			num2 = next
			fmt.Print(" ", next)
		}
	}
}
