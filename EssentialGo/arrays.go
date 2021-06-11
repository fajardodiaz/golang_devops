package main

import "fmt"

func main() {
	//Arrays in go have a fixed size, They can't grow
	a := [2]int{4, 5}
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	a3 := [...]int{1, 2, 3, 4, 5, 6, 67}
	fmt.Println(a3)
	fmt.Println(len(a3))
	fmt.Println(cap(a3))
}
