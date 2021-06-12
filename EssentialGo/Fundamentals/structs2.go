package main

import "fmt"

type User struct {
	FirstName, LastName string
	Email               string
	Age                 int
	userId              int
}

//Returns the full name
func (u *User) FullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func main() {
	var u User
	//Zero value of a struct
	fmt.Printf("u: %#v\n", u)

	//Pointer to user Struct
	pu := new(User)
	pu.FirstName = "Pablo"
	fmt.Printf("u: %#v\n", pu)

	//&User{} is the same as new(User)
	pa := &User{}
	pa.FirstName = "Andres"
	pa.LastName = "Fajardo"
	pa.Age = 21
	fmt.Printf("u: %#v\n", pa)
	fmt.Println("Function Fullname(): ", pa.FullName())

}
