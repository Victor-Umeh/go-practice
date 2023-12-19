package main

import "fmt"

type Others struct {
	Stack string
	Age int
	Tools []string
	Name string
	Self_Learner bool	
}

type person struct {
	Age int
	Name string
	Role string
	Certified bool
}

func main() {
	workTools := map[string]Others{
		"Victor": Others{
			Name: "Victor Onyeka",
			Age: 16,		
			Stack: "frontend",
			Self_Learner: true,
		},
	}

	mapMap := map[string]person{
		"Frank":{
			Name: "frank",
			Age: 28,
			Certified: true,
			Role: "ceo",
		},
		"Pete":{
			Name: "peter",
			Age: 32,
			Certified: true,
			Role: "coo",
		},
		"Okon":{
			Name: "okon",
			Age: 26,
			Certified: false,
			Role: "IT support",
		},
		"Bulus":{
			Name: "bulus",
			Age: 23,
			Certified: false,
			Role: "secretary",
		},
	}

for mapKey, mapValue := range mapMap {
	fmt.Println(mapKey, mapValue)
	fmt.Println(mapValue.Name)
	fmt.Println(mapValue.Role)
	fmt.Println(mapValue.Certified)
}


	// if workTools["Victor"].Self_Learner && workTools["Victor"].Stack == "frontend"{
	// 	fmt.Println("You are doing well son.")
	// }else {
	// 	fmt.Println("You are not worthy.")
	// }
	//var age int
	//fmt.Scan(&age)

	switch workTools["Victor"].Stack {
	case "frontend":
		fmt.Println("An eligible stack.")
	case "backend":
		fmt.Println("Not an eligible stack.")
	default:
		fmt.Println("How did you get here")
	} 

	// fmt.Println(workTools, age)
	// fmt.Println(workTools["Victor"].Age)
}