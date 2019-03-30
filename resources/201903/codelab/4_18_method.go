package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) Major() bool { return p.Age >= 18 }
func (p *Person) Birthday() { p.Age++ }

func main(){
	var h1 Person = Person{FirstName:"Hyejong", LastName:"Hong",Age:17}
	h1.Birthday()
	fmt.Println(h1.Major())
}
