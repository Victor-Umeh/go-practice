package main

import (
	"fmt"
	"lesson-8/helper"
)

var userPronoun string

func main() {
	var userName string
	var userGender string

	fmt.Println("Input your name:")
	fmt.Scan(&userName)

	fmt.Println("Input your gender:")
	fmt.Scan(&userGender)

	userPronoun = helper.GenUserPronoun(userGender)
	
	fmt.Printf("Hello, %+v %+v \n", userPronoun, userName)
	fmt.Println(helper.Company)
}

