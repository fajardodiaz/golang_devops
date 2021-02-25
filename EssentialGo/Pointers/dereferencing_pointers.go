package main

import "fmt"

type Person struct {
	Name     string
	Lastname string
	Age      int
}

func main() {
	//A magic code
	pablo := new(Person)
	pablo.Name = "Pablo"
	pablo.Lastname = "Fajardo"
	pablo.Age = 21
	// fmt.Println(pablo)
	andres := pablo
	andres.Name = "Andres"
	andres.Lastname = "Diaz"
	// fmt.Println(pablo)
	// fmt.Println(andres)
	jose := *pablo
	jose.Name = "Jose"
	jose.Lastname = "Valle"
	jose.Age = 58
	// fmt.Println(jose)

	// fmt.Println("-----------------------------")
	fmt.Println(pablo)
	fmt.Println(andres)
	fmt.Println(jose)

	// ---------------------------------------
	p_pablo := &pablo
	p_andres := &andres
	p_jose := &jose

	fmt.Printf("Pablo: %p\nAndres: %p\nJose: %p\n", p_pablo, p_andres, p_jose)

}
