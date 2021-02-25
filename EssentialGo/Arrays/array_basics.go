package main

import "fmt"

func main() {
	a := [2]int{1, 2} //Array of 2 items
	fmt.Println(a[0])
	fmt.Println(a[1])

	//Modify the array
	a[0] = 10
	fmt.Println(a[0])
	fmt.Println(a[1])

	//Print the size of an array
	fmt.Println("Size of the array")
	fmt.Println(len(a))
	fmt.Println("Capacity of the array")
	fmt.Println(cap(a))

	//[...] array
	b := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Size of the array")
	fmt.Println(len(b))
	fmt.Println("Capacity of the array")
	fmt.Println(cap(b))

	//Array of pointers
	ptrs := [8]*int{}
	fmt.Println(ptrs)
}
