package main

import "fmt"

func main() {
	x := 5
	y := x
	x = 10

	//A and B point to same slice in the memory, use the same address
	var a []int = []int{3, 4, 5}
	//b is not a copy
	b := a

	//If you modify a or b, changes the slice
	b[0] = 100
	a[1] = 200

	fmt.Println(x, y)
	fmt.Println(a, b)
}
