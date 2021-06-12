package main

import "fmt"

type T struct {
	I int
	S string
}

func main() {
	//Initialize a struct
	t := T{1, "one"}

	// make struct copy
	u := t // u has its field values equal to t

	if u == t {
		fmt.Println("u and t are equal")
	}
}
