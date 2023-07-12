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
}
