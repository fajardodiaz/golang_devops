package main

import "fmt"

type T struct {
	I int
	S string
}

func main() {
	//Initialize a struct
	t := T{1, "one"}

	//Make a copy
	u := t

	if t == u {
		fmt.Println("t is equals u")
	}

	t.I = 100
	fmt.Printf("t.I = %d, u.I = %d\n", t.I, u.I) // t.I = 100, u.I = 1

}
