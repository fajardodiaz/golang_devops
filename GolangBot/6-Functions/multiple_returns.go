package main

import "fmt"

func rectProps(length, width float64) (float64, float64) {
	area := length * width
	perimeter := (length + width) * 2
	return area, perimeter
}

func main() {
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area: %0.2f Perimeter: %0.2f\n", area, perimeter)
}
