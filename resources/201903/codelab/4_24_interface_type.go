package main

import "fmt"

func main() {
	var i interface{}
	fmt.Println(i)

	i = 5
	fmt.Println(i)
	i = 1.6
	fmt.Println(i)
	i = "interface"
	fmt.Println(i)
}
