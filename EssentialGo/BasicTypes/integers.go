package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	//Zero value of integers is 0
	fmt.Println("Zero value of integers is 0")
	var i int
	fmt.Printf("i value is %b \n\n", i)

	//Convert int to string with strconv.Itoa
	var i2 = 45
	fmt.Printf("i2 value is a %T\n", i2)
	fmt.Printf("i2 value is %s\n", strconv.Itoa(i2))
	fmt.Printf("i2 datatype with strconv.Itoa is %T\n\n", strconv.Itoa(i2))

	//Convert int to string with fmt.Sprintf
	var i3 int = 12
	stri3 := fmt.Sprintf("%d\n", i3)
	fmt.Printf("i3 value is: %s", stri3)
	fmt.Printf("i3 datatype is: %T\n", i3)
	fmt.Printf("i3 value with fmt.Sprintf is %T\n\n", stri3)

	//Convert string to inr with strvonc.Atoi
	str1 := "48"
	il, err := strconv.Atoi(str1)
	if err != nil {
		log.Fatalf("strconv.Atoi() failed with %s\n", err)
	}
	fmt.Printf("il is: %s and data type is %T\n", str1, str1)
	fmt.Printf("il value with strconv.Atoi() is: %d\n", il)
	fmt.Printf("il data type with strconv.Atoi() is: %T\n\n", il)

}
