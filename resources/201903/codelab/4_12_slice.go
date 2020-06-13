package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c", "d"}
	var bytes = make([]byte, 5, 10)
	b := make([]int, 3, 3)

	fmt.Println(letters, bytes, b)
}
