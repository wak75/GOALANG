package main

import (
	"fmt"
)

func main() {
	//if else is same as in other languages, but with a slight variation
	//1. No need to add a bracket for the condition
	//2. else/ else if should start at the same line where the previous bracket ends
	//3. in the if condition you can initilize a variable

	var result string
	if age := 25; age < 20 {
		result = "Nope not allowed"
	} else if age > 20 && age < 30 {
		result = "you are allowed"
	} else {
		result = "you are too old to be allowed"
	}

	fmt.Println(result)
}
