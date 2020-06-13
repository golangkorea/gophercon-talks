package main

import "fmt"

func main() {
	var a [4]int
	b := [2]string{"Penn", "Teller"}
	c := [...]bool{true, false, false}

	fmt.Println(a, b, c)
}
