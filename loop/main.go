package main

import (
	"fmt"
)

func main() {

	sum := 0

	for i := 0; i < 10; i++ {
		sum += 5
		fmt.Println("the value of sum is", sum)
	}

	var slice = []string{"toca toca", "bang bang", "rise", "good bye", "Dont let me down"}

	for i := range slice {
		fmt.Println(slice[i])
	}

	//for as while loop

	var x int = 0

	for x != 10 {
		fmt.Println("the X is :", x)
		x++
	}
}
