package stringsfunct

import (
	"fmt"
	"strings"
)

func CompareStrings() {
	s := strings.Compare("abc", "Abc")
	fmt.Println()
	fmt.Printf("after comparing the strings %v", s)
}
