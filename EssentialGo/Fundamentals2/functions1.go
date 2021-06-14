package main

import "fmt"

func test(x int) {
	fmt.Println(x)
}

func add_less(x, y int) (int, int) {
	return x + y, y - x
}

func main() {
	test(2)
	test(3)
	test(4)
	suma, resta := add_less(12, 67)
	fmt.Printf("Suma: %d Resta %d\n", suma, resta)
}
