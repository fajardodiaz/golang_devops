package main

import "fmt"

type ResourceFormatter struct{}

func (r *ResourceFormatter) FormatHTTP(resource string) string {
	return fmt.Sprintf("http://%s", resource)
}

func (r *ResourceFormatter) FormatHTTPS(resource string) string {
	return fmt.Sprintf("https://%s", resource)
}

type Request struct {
	Resource string
}

type AuthenticatedRequest struct {
	Request,
	Username, Password string
	ResourceFormatter
}

func main() {
	//Composition provides an alternative to inheritance
	ar := new(AuthenticatedRequest)
	ar.Request = "example.com/request"
	ar.Username = "pfajardo"
	ar.Password = "abc123"
	fmt.Printf("%#v\n", ar)

	fmt.Println(ar.FormatHTTP(ar.Request))
	fmt.Println(ar.FormatHTTPS(ar.Request))
	fmt.Println(ar.FormatHTTPS("www.google.com"))

}
