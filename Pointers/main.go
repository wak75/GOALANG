//pointers in go is similar to as C and C++

package main

import(
	"fmt"
)

func main(){

	var pointerp *int
	fmt.Printf("The value of uninitilized pointer is %v\n", pointerp)

	var interger int = 55;
	pointerp = &interger

	fmt.Printf("the value of pointer is %v *pointer will give the actual value of the pointer\n", *pointerp)
	fmt.Printf("the value of pointer is %v this will actuall show the address of the integer\n", pointerp)
	fmt.Printf("the address of the pointerp is %v\n", &pointerp)

	fmt.Println("Changing the actual data using pointers")

	*pointerp += 100;

	fmt.Println("The data of pointerp after updating is", *pointerp)
	fmt.Println("the data of the integer after updating the pointer is ", interger)
}