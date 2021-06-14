package main

import "fmt"

type User struct {
	name string
}

//A getter
func (u *User) Name() string {
	return u.name
}

//A setter
func (u *User) SetName(newName string) {
	u.name = newName
}

func main() {
	//Struct methods are functions attached to structs
	pablo := new(User)
	pablo.SetName("Andres")
	fmt.Println(pablo.Name())
}
