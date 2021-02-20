package main

import (
	"fmt"
	"unsafe"
)// :show start
s := "348"
var ia int
_, err := fmt.Sscanf(s, "%d", &ia)
if err != nil {
	log.Fatalf("fmt.Sscanf failed with '%s'\n", err)
}
fmt.Printf("i1: %d\n", ia)

func main() {
	var b bool = true
	fmt.Printf("b is: %v\n", b)
	fmt.Printf("b datatype is: %T\n", b)

	//Zero value of type bool is false
	var b2 bool
	fmt.Printf("zero value of bool is: %v\n", b2)

	//Size of bool variable is 1 byte
	b3 := true
	fmt.Printf("Size of bool is: %d\n", unsafe.Sizeof(b3))
}
