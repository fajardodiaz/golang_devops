// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"log"
)

func main() {
	str2 := "60"
	var intstr2 int
	_, err := fmt.Sscanf(str2, "%d", &intstr2)
	if err != nil {
		log.Fatalf("fmt.Sscanf failed with '%s'\n", err)
	}
	fmt.Printf("intstr2: %d\n", intstr2)
}
