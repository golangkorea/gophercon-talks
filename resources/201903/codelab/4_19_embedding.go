package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	string
}

func main() {
	var h1 Person = Person{FirstName: "Hyejong", LastName: "Hong", string: "2323"}
	h1.string = "14141"
	fmt.Println(h1)
}
