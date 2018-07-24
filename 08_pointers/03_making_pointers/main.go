package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0x......

	var b *int = &a
	fmt.Println(b)  // 0x......
	fmt.Println(*b) // 43

	*b = 42
	fmt.Println(a) // 42
}
