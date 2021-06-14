package main

import "fmt"

func change_value(str *string) {
	*str = "Changed"
}

func change_value2(str string) {
	str = "Changed"
}

func main() {
	my_str := "Hello!"

	fmt.Println("With Pointers")
	//The string change
	fmt.Println(my_str)
	change_value(&my_str)
	fmt.Println(my_str)

	fmt.Println("Without Pointers")
	//The string doesn't change
	fmt.Println(my_str)
	change_value2(my_str)
	fmt.Println(my_str)
}
