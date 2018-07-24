package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
}

// this code makes b a pointer to the meory address where an int is stored
// b is of type "int pointer"
// *int -- the * is part of the type -- b is of type *int
// ('pointer to an int' - points to a memory address where an int is stored)
