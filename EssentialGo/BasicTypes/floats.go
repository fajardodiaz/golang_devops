package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	//Convert float to string wiht FormatFloat
	var f32 float32 = 1.3
	bitSize := 32
	s1 := strconv.FormatFloat(float64(f32), 'E', -1, bitSize)
	fmt.Printf("f32: %s\n", s1)
	fmt.Printf("datatype of f32: %T\n\n", s1)

	var f64 float64 = 8545.14554545
	bitSize = 64
	s2 := strconv.FormatFloat(f64, 'e', -1, bitSize)
	fmt.Printf("f64: %s\n", s2)
	fmt.Printf("datatype of f64: %T\n\n", s2)

	//Convert float to string with Sprintf
	var ff64 float64 = 1.3
	sf2 := strconv.FormatFloat(ff64, 'e', -1, bitSize)
	fmt.Printf("f64: %s\n", sf2)
	fmt.Printf("datatype of f64: %T\n\n", sf2)

	//Convert string to float with ParseFloat
	s := "1.2341"
	f64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatalf("strconv.ParseFloat() failed with '%s' \n", err)
	}
	fmt.Printf("f64: %f\n", f64)
	fmt.Printf("datatype of f64: %T\n\n", f64)

}
