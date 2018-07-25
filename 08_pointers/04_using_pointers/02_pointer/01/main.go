package main

import "fmt"

func zero(z *int) {
	*z = 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0 bc it is dereferenced (value is changed) when func zero is called
}
