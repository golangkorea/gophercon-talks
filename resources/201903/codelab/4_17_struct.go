package main

import "fmt"

type Person struct {
	_         struct{}
	FirstName *string
	LastName  *string
	Age       *int
}

func main() {
	var h1 Person = Person{FirstName: "Hyejong", LastName: "Hong", Age: 29}
	var h2 Person = Person{"Kwon", "Minjae", 25}

	fmt.Println(h1, h2)
}
