package main

import "fmt"

func main() {
	var foo string                        //""                     // Declaring a variable with initial value
	var bar = 1                           // Type inference
	var i, j int = 1, 2                   // Multiple variable declaration
	c, python, java := true, false, "no!" // Short hand declaration

	fmt.Println(foo, bar, i, j, c, python, java)
}
