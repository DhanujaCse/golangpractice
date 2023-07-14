package arrayws

import "fmt"

func EvenOrOdd(num int) {
	if num%2 == 0 {
		fmt.Printf("%v is even number", num)
	} else {
		fmt.Printf("%v is odd number", num)

	}

}
