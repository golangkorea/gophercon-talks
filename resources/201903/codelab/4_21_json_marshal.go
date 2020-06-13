package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
	F1 int `json:"f1"`
	F2 int `json:"F2"`
	F3 int `json:"-"`
}

func main() {
	t := T{1, 0, 3}
	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b) // {"f1":1,"F2":0}
}
