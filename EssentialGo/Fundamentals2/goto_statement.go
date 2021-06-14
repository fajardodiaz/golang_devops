package main

import "fmt"

func printIsOdd(num int) {
	if num%2 == 1 {
		goto isOdd
	} else {
		goto isEven
	}

isOdd:
	fmt.Printf("%d is odd\n", num)
	return

isEven:
	fmt.Printf("%d is odd\n", num)
	return

}

func main() {
	printIsOdd(12)
	printIsOdd(15)
}
