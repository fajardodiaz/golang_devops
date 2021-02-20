package main

import "fmt"

func main() {
	var a1 = [2]byte{3, 8}
	//when using [...] size will be deduced from {...}
	a2 := [...]byte{10, 20, 30, 40}

	fmt.Println("Array ", a1)
	fmt.Println("Array length:", len(a1))

	fmt.Println("Array ", a2)
	fmt.Println("Array length:", len(a2))
}
