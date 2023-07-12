//you can not add different data types, you have to implesitly wrap them before before performning some operations

package main

import (
	"fmt"
	"math"
)

func main() {

	inta := 45
	floata := 44.45

	//sum := inta + floata //this will give you error
	sum := float64(inta) + floata //this will not give you error
	fmt.Println(sum)

	f1,f2,f3 := 78.6 , 448.9 , 87.662
	floatSum := f1+f2+f3
	floatSum = math.Round(floatSum*100)/100
	fmt.Println(floatSum)

	radius := 55.55
	circumfrence := 2 * math.Pi * radius

	fmt.Printf("The Circumfrence is: %.2f\n", circumfrence)

}
