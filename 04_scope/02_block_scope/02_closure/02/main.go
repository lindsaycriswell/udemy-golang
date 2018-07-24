package main

import "fmt"

var x = 0

func increment() int {
	x++
	return x
	// closure makes new value of x available below
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}
