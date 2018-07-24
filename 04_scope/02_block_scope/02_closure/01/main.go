package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		// x is in scope here as a closure
		y := "stupid string example"
		fmt.Println(y)
	}
	// fmt.Println(y) // out of scope of y
}
