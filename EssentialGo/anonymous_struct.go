package main

import "fmt"

func main() {

	data := struct {
		Number int
		Text   string
	}{
		12,
		"Hello",
	}

	fmt.Printf("Data: %+v\n", data)
}
