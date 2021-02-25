package main

import "fmt"

func main() {
	v := 15

	pointer_v := &v
	fmt.Printf("V: %d  - Pointer: %p\n", v, pointer_v)

	//Change the value of v via the pointer with *
	*pointer_v = 4
	fmt.Printf("V: %d  - Pointer: %p\n", v, pointer_v)

	//Two pointer with the same value have the same address
	pointerV2 := &v
	fmt.Printf("V: %p  - Pointer: %p\n", pointer_v, pointerV2)

}
