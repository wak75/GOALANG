package main

import (
	"bufio"
	"fmt"
	"os"
)

const aConst int = 99;

func main() {
	fmt.Println("the constsant number is", aConst)
	// fmt.Print("Enter your name: ")
	// reader := bufio.NewReader(os.Stdin)

	// input, _ := reader.ReadString('\n')
	// fmt.Println("The name you have entered is ", input)
	// fmt.Printf("the type of input is %T" , input)


	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Where do you live: ")
	input,_ := reader.ReadString('\n')
	fmt.Println("You Live in", input)

}
