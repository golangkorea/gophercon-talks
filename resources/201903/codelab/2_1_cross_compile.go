package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

// GOOS=linux   GOARCH=amd64 go build 4_1_cross_compile.go
// GOOS=windows GOARCH=amd64 go build 4_1_cross_compile.go
