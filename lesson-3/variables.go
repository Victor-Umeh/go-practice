package main

import "fmt"


func main() {
	// constants variables, cannot be reassigned a value
	const PI = 3.1428

	//type inferrence
	name := "pascal"
    var age = 35

    //unused varaibles
	var userName string
	var userId uint

	//assigns value to unused variables
	userName = "Leo"
	userId = 107

	fmt.Printf("Name: %v Id: %v \n", userName, userId)
	fmt.Printf("%T %T %v", PI, name, age)


	var firstName string

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)


	fmt.Print("Your name is ", firstName)
	/*
	 Data types in golang:
	  1) int: Integers
	  2) float: Decimals
	  3) string: Strings😂
	  3) bool: true or false
	  4) Other types in golang:
	     i) int8, int16, int32, int64
		 ii) uint, uint8, uint16, uint32, uint64, uintptr
		 iii) float32, float64
		 iv) complex64, complex128
	  
	*/
}

