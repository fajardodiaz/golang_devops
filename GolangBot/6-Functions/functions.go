package main

import "fmt"

/*
Function syntax:


func function_name(parameter_name type) return_type{
	Code...
}

*/

func calculateBill(price float64, no int) float64 {
	totalPrice := price * float64(no)
	return totalPrice

}

func main() {
	fmt.Println("Welcome to my app")
	fmt.Println("Calculate the Bill")
	fmt.Printf("Bill: $%0.2f\n", calculateBill(10.50, 7))
}
