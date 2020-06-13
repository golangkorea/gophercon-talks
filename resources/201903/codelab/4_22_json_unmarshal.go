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
	b := []byte(`{"f1":1,"F2":2, "F3":3}`)
	t := T{}
	if err := json.Unmarshal(b, &t); err != nil {
		panic(err)
	}
	fmt.Println(t) // {1 2 0}
}
