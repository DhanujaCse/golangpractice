package stringsfunct

import (
	"fmt"
	"strconv"
)

func StringFunc() {

	num := "233"
	val, err := strconv.Atoi(num)
	if err != nil {
		fmt.Printf("the value specified as String is integer is is %v", val)
	} else {
		fmt.Printf("Supplied value %s is a number with value %d\n", num, val)
	}

}
