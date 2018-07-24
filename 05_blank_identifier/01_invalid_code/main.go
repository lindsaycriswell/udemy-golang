package main

import "fmt"

func main() {
	a := "stored in a"
	b := "stored in b"

	fmt.Println("a - ", a)
	// b is invalid code bc it is not being used
}

// # command-line-arguments
// 05_blank_identifier/01_invalid_code/main.go:7:2: b declared and not used
