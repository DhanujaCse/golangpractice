package stringsfunct

import (
	"fmt"
)

func MultilineStrings() {

	string1 := `this is 
	string written 
	in a new line`
	fmt.Printf("this is string written in multi line thatis \n %v", string1)
}
