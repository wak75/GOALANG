package main

import ("fmt")

func main()  {
	//map is onorded collection of key value pairs, same as dictionaries in python and hashtable in java

	states := make(map[string]string)

	fmt.Println(states)
	states["WB"] = "West Bengal"
	states["UP"] = "Utter Pardesh"
	states ["JNK"] = "Jammu and Kashmir"

	fmt.Println(states)

	//deleting a key
	delete(states, "UP")
	fmt.Println(states)

	//adding a key
	states["AS"] = "Assam"
	states["AP"] = "Andra Pardes"
	states["ND"] = "New Delhi"
	
	
	// for i:=0;i<len(states);i++ {
	// 	fmt.Println(states[i])
	// }

	for i, j := range states{
		//fmt.Println(states[i], states[j])
		fmt.Printf("%v: %v \n", i , j)
	}

}