package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		// anonymous function - assigned to variable
		// func increment() would not work
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
