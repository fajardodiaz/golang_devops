package main

import "fmt"

func main() {
	var s string
	s1 := "string\nliteral\nwith\tescape characters\n"
	s2 := `raw string literal
	which doesnt't recgonize escape characters like \n
	`
	fmt.Printf("sum of the strings\n'%s'\n", s+s1+s2)

}
