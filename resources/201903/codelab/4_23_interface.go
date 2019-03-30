package main

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}

type Stringer2 interface {
	String() string
}

type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(uint64(i), 2)
}

func main() {
	b := Binary(200)

	s := Stringer(b)
	fmt.Println(s.String())

	b2, ok := s.(Binary) // interface value -> concrete value
	if ok != false {
		fmt.Println(b2, ok)

	}
}
