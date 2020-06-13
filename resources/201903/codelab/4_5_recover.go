package main

import "fmt"

func Something() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	panic("Error!!")
}

func main() {
	Something()

	fmt.Println("Hello World.")
}
