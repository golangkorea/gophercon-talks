package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	close(ch)

	ch <- 3
	fmt.Println("Hello")
}
