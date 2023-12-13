package main

import "fmt"

func main() {

	//Exercise 1: UserInput Prompt
	//Basic usage of pointers in GO

	var firstName string
	var lastName string
	var age int
	var email string

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your age:")
	fmt.Scan(&age)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Printf("%v %v have successfully subscribed to HYDRA project, a cofrimation message will be sent to your email @%v if your age %v is eligeble", firstName, lastName, email, age)
}