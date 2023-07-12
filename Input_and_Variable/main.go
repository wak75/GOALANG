package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("The name you have entered is ", input)
	fmt.Printf("the type of input is %T" , input)
}
