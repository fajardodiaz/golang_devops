package main

import "fmt"

func main() {
	//Slices is a growable sequence of values of the same type

	slice := make([]int, 0, 12)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = append(slice, 10)
	slice = append(slice, 11)
	slice = append(slice, 12)
	slice = append(slice, 13)
	slice = append(slice, 14)
	slice = append(slice, 15)
	slice = append(slice, 16)
	slice = append(slice, 17)
	slice = append(slice, 18)
	slice = append(slice, 19)
	slice = append(slice, 20)
	slice = append(slice, 21)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
