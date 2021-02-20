package main

import "fmt"

func main() {
	//Iterate over a string using bytes
	s := "string to count"
	for i := 0; i < len(s); i++ {
		c := s[i]
		fmt.Printf("Byte at index %d is '%c' (0x%x)\n", i, c, c)
	}

	//Iterate over a string using runes
	s2 := "string to count"
	for j, runeChar := range s2 {
		fmt.Printf("Rune at byte position %d is %#U\n", j, runeChar)
	}
}
