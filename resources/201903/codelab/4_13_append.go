package main

import "fmt"

func main(){
	s := make([]int, 4, 4)
	s = append(s, 1)

	fmt.Println(s,len(s),cap(s))
}
