package main

import (
	"fmt"
	"log"
)

func main() {
	//Convert string to float with Sscanf
	s3 := "1.2345"
	var f float64
	_, err := fmt.Sscanf(s3, "%f", &f)
	if err != nil {
		log.Fatalf("fmt.Sscanf failed with '%s'\n", f)
	}
	fmt.Printf("f: %f\n", f)
	fmt.Printf("The value of f is %T\n", f)
}
