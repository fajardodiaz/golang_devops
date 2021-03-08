package main

import "fmt"

func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}

func main() {
	area, _ := rectProps(11, 5.6)
	fmt.Printf("Area: %0.2f\n", area)
}
