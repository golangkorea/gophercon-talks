package main

import "fmt"

func main() {
	var m map[string]int
	m = make(map[string]int)

	names := map[string]int{
		"rsc": 3711,
	}

	tables := make(map[string]int)

	fmt.Println(m, names, tables)
}
