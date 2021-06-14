package main

import "fmt"

func main() {
	/**
	 * & = Pointer
	 * * = Dereference
	 */

	x := 7
	y := &x
	fmt.Println(x, y)

	*y = 1
	fmt.Println(x, y)

}
