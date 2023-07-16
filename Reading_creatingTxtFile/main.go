package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	myString := "The String entered from the code"

	file, err := os.Create("./.txt")
	checkErr(err)

	length, err := io.WriteString(file, myString)
	checkErr(err)

	fmt.Println("the file has been written with", length, "of chatacters")

	defer file.Close() //defer is wait until every operation is done and then execute the command
	defer readAFile("./TextFile.txt")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readAFile(fileName string)  {
	data, err := ioutil.ReadFile(fileName)
	checkErr(err)

	fmt.Println(string(data))
}
