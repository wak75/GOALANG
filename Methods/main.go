package main

import ("fmt")

func main(){

	emp := Employee{"king", 10, 999999.99}
	emp.printData()

}


type Employee struct {
	name string
	years int
	salary float32
}

func (e Employee ) printData()  {
	fmt.Println("Name of the employee", e.name)
	fmt.Println("Years of experience : ", e.years)
	fmt.Printf("The salary is %v\n", e.salary)
}