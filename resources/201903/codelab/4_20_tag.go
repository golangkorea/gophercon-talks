package main

import (
	"fmt"
	"reflect"
)

type T struct {
	f1     string "f one"
	f2, f3 int64  `f four and five`
	f4     string `one:"1"`
}

func main() {
	t := reflect.TypeOf(T{})
	f1, _ := t.FieldByName("f1")
	fmt.Println(f1.Tag) // f one

	f2, _ := t.FieldByName("f2")
	f3, _ := t.FieldByName("f3")
	fmt.Println(f2.Tag, f3.Tag) // f four and five

	f4, _ := t.FieldByName("f4")
	fmt.Println(f4.Tag) // one:"1"

	v, ok := f4.Tag.Lookup("one")
	fmt.Printf("%s, %t\n", v, ok) // 1, true
}
