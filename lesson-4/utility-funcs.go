package main

import "fmt"

func greetUserWithName(name string) {
	fmt.Printf("Hello %v\n", name)
}

func showAuthorName() {
	fmt.Println("This is a constant value declared in function file within this package:", authorName)
}