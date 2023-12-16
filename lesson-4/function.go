package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

//main function, entry point for all go programs
func main() {
	// var name string

	// //call greetUser function
 	// greetUser()

	// fmt.Println("Enter your name:")
	// fmt.Scan(&name)
	// greetUserWithName(name)

	var books = []string{}
	books = append(books, "Things fall apart", "Game of thrones", "Willy will")

	firstTitle := []string{}
	//for Each loop
	for _, book := range books {
		var main = strings.Fields(book)
		fmt.Println(main[0])
		firstTitle = append(firstTitle, main[0])
	}

	fmt.Println(firstTitle)
	fmt.Println(reflect.TypeOf(firstTitle))


	//A func that takes in a list of numbers and returns a new list with eash number ++
	numsArr  := []int {1,2,0,3,4,5}
	fmt.Println(increaseByOne(numsArr))

	//*** loops ***

	//infinite loop
	// var count int
	// for {
	// 	fmt.Scan(&count)
	// 	fmt.Printf("Loop %v", count + 1)
	// }

	//for loop
	for i:= 1; i <= 10 ; i++ {
		fmt.Println(i)
	} 

	name_1, name_2, name_3 := returnNamesVariables()
	fmt.Println(name_1, name_2, name_3)

	var day string
	
	log.Println("Enter a day")
	fmt.Scan(&day)
	fmt.Println(getDay(day))

}

//function returning multiple variables
func returnNamesVariables()(string, string, string) {
	name_1 := "smith"
	name_2 := "peter"
	name_3 := "chike"
	
	//return multiple variables at once
	return name_1, name_2, name_3
}

//create greetUser function
func greetUser() {
	fmt.Println("Hello World!")
}

func greetUserWithName(name string) {
	fmt.Printf("Hello %v", name)
}

//a function that take in an slice and returns a new slice
func increaseByOne (arr []int) []int {
	var newArr = []int{}
	for i := 0; i < len(arr); i++ {
		if arr[i] < 1 {
			continue
		}
		newArr = append(newArr, arr[i] + 1)
	}
	return newArr
}

//check var type func
func getDay( val  string) uint{

 switch val {
 	case "sunday":
		return 1
	case "monday":
		return 2
	case "tuesday":
		return 3
	case "wednesday":
		return 4
	case "thursday":
		return 5
	case "friday":
		return 6
	case "saturday":
		return 7
	default:
		return 419
 }
}