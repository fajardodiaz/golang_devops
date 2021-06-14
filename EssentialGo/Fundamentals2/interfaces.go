package main

import (
	"fmt"
	"strconv"
)

//Interfaces defines a set of methods on a type
//Interfaces are used to abstract behavior

type Stringer interface {
	String() string
}

type User struct {
	Name string
}

func (u *User) String() string {
	return u.Name
}

//Any tipe can implement an interface.
type MyInt int

func (mi MyInt) String() string {
	return strconv.Itoa(int(mi))
}

func printTypeAndString(s Stringer) {
	fmt.Printf("%T: %s\n", s, s)
}

func main() {
	u := &User{"Pablo"}
	printTypeAndString(u)

	n := MyInt(5)
	printTypeAndString(n)

}
