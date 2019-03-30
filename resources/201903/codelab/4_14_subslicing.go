package main

import "fmt"

func main(){
	s := []int{0, 1, 2, 3, 4}
	s = s[2:4]
	fmt.Println(s, len(s), cap(s))

	s = s[1:]
	fmt.Println(s, len(s), cap(s))
}
