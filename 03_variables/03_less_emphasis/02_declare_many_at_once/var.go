package main

import "fmt"

func main() {
	var message string
	var a, b, c int

	a = 1
	message = "Hello World"

	fmt.Println(message, a, b, c)
}

// b and c default to 0 (zero values) bc they are not assigned
