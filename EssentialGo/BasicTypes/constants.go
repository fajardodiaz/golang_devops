package main

import "fmt"

func main() {
	const (
		i  int = 32
		s      = "string"
		i2     = 33
	)

	var (
		b = []byte{3, 4} //This could not be a constant
	)

	fmt.Println(b)
}
