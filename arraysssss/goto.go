package arrayws

import "fmt"

func Goto() {
	var input int

Loop:
	fmt.Println("enter a age")
	fmt.Scan(&input)
	if input < 18 {

		fmt.Println("you are not eligible for voting")
		goto Loop
	} else {
		fmt.Println("you are eligible for voting")
	}
}
