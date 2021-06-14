package main

import "fmt"

func main() {

	//Basic if
	if 5 < 20 {
		fmt.Println("Welcome!")
	} else {
		fmt.Println("Byeeee!")
	}

	//Basic switch
	stmt := "switch"

	switch stmt {
	case "if", "for":
		fmt.Println("Statement: For")
	case "switch":
		fmt.Println("Statement: Switch")
	default:
		fmt.Println("I don't know")
	}
}
