package main

import (
	"fmt"
	"myapp/packageone"
)

var one = "One" // Global variable

func main() {
	var one = "this is a block level variable"
	fmt.Println(one)
	myFunc()	
	newString := packageone.PublicVar
	fmt.Println(newString)
	
	packageone.Exported() // This will work because Exported is exported

// 	secondString := packageone.PrivateVar
// 	fmt.Println(secondString) // This will not work because PrivateVar is not exported
}


func myFunc(){

	fmt.Println(one)
}