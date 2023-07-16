package main

import(
	"fmt"
)

func main()  {
	//same as in other languages but no need to add break 

	number := 55
	var result string
	switch number {
	case 10:
		result="number is 10"
		// fallthrough
	case 20 :
		result = "number is 20"
		//fallthrough
	case 55:
		result = "number is 55"
		//fallthrough
	default:
		result = "number is not defined"
	
	}

	fmt.Println(result)

}