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


	fmt.Println("Your name is", firstName)

	//Arrays
	var names [10]string
	names[0] = "frank"
	names[1] = "vick"

	// fmt.Printf("Array: %v \n", names)
	// fmt.Printf("Array type: %T \n", names)
	// fmt.Printf("First subscriber: %v \n", names[0])
	// fmt.Printf("First type: %T \n", names[0])
	// fmt.Println(len(names))
	// fmt.Printf("Array length %T", len(names))

   	//SLice
	var subscribersId []int

	subscribersId = append(subscribersId, 12, 22, 45)
	fmt.Println(subscribersId)
	fmt.Println(subscribersId[0])
	fmt.Println(subscribersId[1])
	fmt.Println(subscribersId[2])


	
	/*
	 Data types in golang:
	  1) int: Integers
	  2) float: Decimals
	  3) string: Strings😂
	  3) bool: true or false
	  4) Array
	  5)Slice
	  6) Other types in golang:
	     i) int8, int16, int32, int64
		 ii) uint, uint8, uint16, uint32, uint64, uintptr
		 iii) float32, float64
		 iv) complex64, complex128
	*/
}

