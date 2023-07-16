package main

import (
	"fmt"
)

func main()  {
	var colors [4]string
	colors[1] = "red"
	colors[3] = "green"

	fmt.Println(colors)

	cars := [...]string{"Lambo","Buggati","Rolls Royce"}
	fmt.Println(cars)
}