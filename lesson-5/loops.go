package main

import "fmt"

func main() {
	//for loop
	for count := 1; count <= 10; count++ {
		fmt.Println(count)
	}

	//while loop
	count := 0
	for count < 5 {
		fmt.Printf("Whil loop: %v\n", count)
		count++
	}

	//forEach loop, loop through an iterable
	
	numsArr := []int{1,2,3,4,5}
	for _, num := range numsArr {
		var newNum int
		newNum = num + 1

		// fmt.Printf("Index: %v, Value: %v\n", index, num)
		fmt.Printf("Value: %v\n", newNum)

	}

}