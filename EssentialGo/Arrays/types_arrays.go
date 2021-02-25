package main

import (
	"fmt"
)

type Array struct {
	Number int
	Text   string
}

func main() {
	arrayStruct := [8]Array{}
	fmt.Println(arrayStruct)
	fmt.Println(arrayStruct)

}
