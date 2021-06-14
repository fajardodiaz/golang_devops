package main

import "fmt"

func return_func(x string) func() {
	return func() { fmt.Println(x) }
}

func main() {
	return_func("Pablo")()
}
