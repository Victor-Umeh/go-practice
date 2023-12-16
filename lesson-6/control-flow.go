package main

import "fmt"

var names = []string{"pollie", "sa", "g", "leo", "saba", "m", "moss", "dappa"}
var ifValidNames = []int{}
var switchValidNames = []int{}

func main() {
	//if, else-if, else statements

	//A simple programm that filters out names whose length are greater than 3
	
	

	for _, name := range names {
		if len(name) == 1 {
			ifValidNames = append(ifValidNames, 1)
		} else if len(name) == 2 {
			ifValidNames = append(ifValidNames, 2)
		} else if len(name) == 3 {
			ifValidNames = append(ifValidNames, 3)
		} else {
			continue
		}
		
		// if len(name) ==1 || len(name) ==2 || len(name) ==3 {
		// 	if len(name) == 1 {
		// 	ifValidNames = append(ifValidNames, 1)
		// 	}
		// 	if len(name) == 2 {
		// 		ifValidNames = append(ifValidNames, 2)
		// 		}
		// 	if len(name) == 3 {
		// 		ifValidNames = append(ifValidNames, 3)
		// 		}
		// }
	}

		//switch statements

		for _, name := range names {
			switch len(name) {
			case 1:
				 switchValidNames = append(switchValidNames, 1)
			case 2:
				switchValidNames = append(switchValidNames, 2)
			case 3:
				switchValidNames = append(switchValidNames, 3)
			default: 
			switchValidNames = append(switchValidNames, -1)
				
			}
		}

	fmt.Println("Slices:",switchValidNames, "Length:", len(switchValidNames))
}