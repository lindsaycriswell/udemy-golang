package main

import "fmt"

func main() {
	fmt.Println(x)
	fmt.Println(y)
	x := 42
	// doesn't work bc x is declared after being called
}

var y = 42

// this works bc it's delcared at the package level so it's available in the block level
