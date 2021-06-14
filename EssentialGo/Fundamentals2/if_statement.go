package main

import "fmt"

func returnValues(ok bool) (int, bool) {
	return 5, ok
}

func main() {
	a := 5
	b := 10

	if a != b {
		fmt.Printf("%d is different than %d\n", a, b)
	}

	if n, ok := returnValues(false); ok {
		fmt.Printf("ok is true, n: %d\n", n)
	} else {
		fmt.Printf("ok is false, n: %d\n", n)
	}
}
