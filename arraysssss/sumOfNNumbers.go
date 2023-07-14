package arrayws

import "fmt"

func SumOfNNumbers(n int) int {
	var sum int
	// fmt.Println("enter the n number to print sum of it")
	// fmt.Scan(&n)

	for i := 0; i < n; i++ {
		sum = sum + i

	}
	fmt.Println("the sum of numbers is", sum)
	return sum
}
