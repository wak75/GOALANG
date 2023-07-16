package main

import (
	"fmt"
	"time"
)

func main(){

	MT := Bike{20.0,"MT Rage", time.Date(2022, time.April, 10,10,10,10, 10 ,time.UTC) }

	fmt.Printf("%v+\n", MT)
	fmt.Println("Name of the bike", MT.Model)
	fmt.Println("HorsePower is ", MT.HorsePower)
	fmt.Println("Manfacturing Year is",MT.Manufacturing_Year )

	MT.Model = "MT HyperSports"
	fmt.Printf("%v+\n", MT)

}


type Bike struct {
	HorsePower float32;
	Model string
	Manufacturing_Year time.Time
}