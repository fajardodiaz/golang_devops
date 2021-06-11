package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[1] = "Pablo"
	m[2] = "Andres"
	m[3] = "Jose"

	fmt.Println(m[1])

}
