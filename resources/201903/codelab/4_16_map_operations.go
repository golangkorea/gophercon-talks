package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["foo"] = 5 // write

	_ = m["foo"] // read

	foo, exist := m["foo"] // second read
	fmt.Println(foo, exist)

	//delete(m, "foo") // delete

	fmt.Println(m)

}
