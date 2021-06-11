package main

import "fmt"

type Request struct {
	Resource string
}

type AuthenticatedRequest struct {
	Request,
	Username, Password string
}

func main() {
	//Composition provides an alternative to inheritance
	ar := new(AuthenticatedRequest)
	ar.Request = "example.com/request"
	ar.Username = "pfajardo"
	ar.Password = "abc123"
	fmt.Printf("%#v", ar)

}
