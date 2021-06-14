package main

import "fmt"

func test2(myFunc func(int) int) {
	fmt.Println(myFunc(10))
}

func greet_test(name string) {
	fmt.Printf("Welcome %s\n", name)
}

func main() {
	test := func(x int) int {
		return x * 10
	}

	test_2 := func(x int) int {
		return x * 15
	}

	greet_test("Pablo")

	//Equals
	test2(test)
	fmt.Println(test(12))

	//Equals
	test2(test_2)
	fmt.Println(test_2(123))

}
