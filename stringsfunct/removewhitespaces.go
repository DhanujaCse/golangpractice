package stringsfunct

import (
	"fmt"
	"strings"
)

func ReplaceWhiteSpaces() {
	string1 := "this is a happy day"
	string2 := strings.ReplaceAll(string1, " ", "")
	fmt.Println(string2)

}
