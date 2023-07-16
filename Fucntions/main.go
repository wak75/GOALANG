package main

import ("fmt")

func main(){

	function1()
	function2("king")
	resultFnc3 :=function3(1,2,3,4,5)
	fmt.Println(resultFnc3)

	resultFnc4, lenresultFnc4 := function4(1,2,3,4,5,6,7)
	fmt.Println(resultFnc4, lenresultFnc4)

}

//no arguments/return type
func function1()  {
	fmt.Println("This is from function one")
}

//with argument no return type
func function2(name string){
	fmt.Println("Hello", name, "how are you")
}

//with argument and one return type
func function3(value ... int) int{
	result:=0
	for i:=0;i<len(value);i++{
		result+=value[i]
	}
	return result
}

//with argument with 2 return type
func function4(value ... int) (int, int){
	result:=0
	for i:=0;i<len(value);i++{
		result+=value[i]
	}
	return result, len(value)
}