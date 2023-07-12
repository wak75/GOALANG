package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main(){
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Favourite Number:")
	number,_ := reader.ReadString('\n')

	enteredNum, err := strconv.ParseInt(strings.TrimSpace(number), 10, 64)

	if err!=nil{
		fmt.Println(err)
	} else {
		fmt.Println("the numner you have entered is", number)
		newNumber:= enteredNum + 59;
		fmt.Println("The updated number is ", newNumber)
	}



}