package stringsfunct

import (
	"fmt"
	"strings"
)

func CheckStringIsPresentOrNot() {
	result := strings.Contains("abc", "ab")
	fmt.Println("checking whether the string contains the given substring or not     ")
	fmt.Println(result)
	result1 := strings.Contains("abc", "vfb")

	fmt.Println(result1)
}
