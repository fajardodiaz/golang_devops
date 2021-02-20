package main

import "fmt"

func main() {
	fmt.Println("zero values for basic types:")

	var zeroBool bool
	fmt.Printf("bool:       %v\n", zeroBool)

	var zeroInt int
	fmt.Printf("int:        %v\n", zeroInt)

	var zeroF32 float32
	fmt.Printf("float32:    %v\n", zeroF32)

	var zeroF64 float64
	fmt.Printf("float64:    %v\n", zeroF64)

	var zeroStr string
	fmt.Printf("string:     %#v\n", zeroStr)

	var zeroPtr *int
	fmt.Printf("pointer:    %v\n", zeroPtr)

	var zeroSlice []uint32
	fmt.Printf("slice:      %v\n", zeroSlice)

	var zeroMap map[string]int
	fmt.Printf("map:        %#v\n", zeroMap)

	var zeroInterface interface{}
	fmt.Printf("interface:  %v\n", zeroInterface)

	var zeroChan chan bool
	fmt.Printf("channel:    %v\n", zeroChan)

	var zeroArray [5]int
	fmt.Printf("array:      %v\n", zeroArray)

	type struc struct {
		a int
		b string
	}
	var zeroStruct struc
	fmt.Printf("struct:     %#v\n", zeroStruct)

	var zeroFunc func(bool)
	fmt.Printf("function:   %v\n", zeroFunc)

}
