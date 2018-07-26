package main

import "fmt"

func main() {
	for i := 0; i < 150; i++ {
		fmt.Println(i, "-", string(i), "-", []byte(string(i)))
	}
	foo := 'a'
	// single quotes make characters into runes!
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
	fmt.Printf("%T \n", rune(foo))
}
