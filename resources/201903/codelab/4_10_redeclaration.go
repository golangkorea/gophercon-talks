package main

import "fmt"

//func main() {
//	foo := 1
//	foo := 2
//
//	fmt.Println(foo)
//}

//func main() {
//	foo := 1
//	{
//		foo := 2
//		fmt.Println(foo)
//	}
//	fmt.Println(foo)
//}

func main() {
	//foo, bar := 1, 2
	foo, err := someFunc()
	waldo, err := someFunc()

	fmt.Println(foo, waldo, err)
}
func someFunc() (int, error) {
	return 1, nil
}
