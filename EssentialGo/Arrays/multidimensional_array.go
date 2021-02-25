package main

import "fmt"

func main() {
	multiDimArray := [2][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}
	multiDimArray2 := [2][5]int{[5]int{1, 2, 3, 4, 5}, [5]int{6, 7, 8, 9, 10}}
	fmt.Println(multiDimArray)
	fmt.Println(multiDimArray2)
	fmt.Println(multiDimArray2[0][0])
	fmt.Println(multiDimArray2[1][0])
}
