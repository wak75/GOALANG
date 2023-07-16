package main

import (
	"fmt"
	"sort"
)

func main(){
	var slice = []string{"lambo", "mercedese", "Audi", "JLR" }
	fmt.Println(slice)

	//appending into a slice
	slice = append(slice, "KarlMann King")
	fmt.Println(slice)

	//removing from a slice
	//removing the first element
	slice = append(slice[1:len(slice)])
	fmt.Println(slice)

	//removing the last element
	slice = append(slice[:len(slice)-1])
	fmt.Println(slice)


	slice = append(slice[2:len(slice)])
	fmt.Println(slice)

	//removing from the middle part is still remaining


	//ways to create a slice

	numbers:=make([]int,10)
	numbers[0]=10
	numbers[1]=20
	numbers[2]=90
	numbers[3]=3
	numbers[4]=02
	numbers[5]=33
	numbers[6]=55
	numbers[7]=267

	fmt.Println(numbers)
	numbers = append(numbers, 786)
	
	
	fmt.Println(numbers)
	sort.Ints(numbers)

	fmt.Println(numbers)


}