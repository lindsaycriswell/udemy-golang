package main

import "fmt"

func wrapper() func() int {
	// func() int {...} is the return value
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
