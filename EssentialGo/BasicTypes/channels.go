package main

import "fmt"

func main() {
	//Create unbuffered channel of int values with capacity of 1
	ch := make(chan int)
	//Starta a new go routine that sends value 3 over a channel
	go func() { ch <- 3 }()
	//Read the value from a channel
	//it waits until go routine above sends a value
	n := <-ch
	fmt.Printf("n: %d\n", n)
}
