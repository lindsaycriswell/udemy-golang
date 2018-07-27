package main

import "fmt"

func main() {
	switch "LC" {
	case "Ian":
		fmt.Println("This is")
	case "Peter":
		fmt.Println("a")
	case "Sean":
		fmt.Println("contrived")
	case "LC":
		fmt.Println("example")
	default:
		fmt.Println("EBUIS")
	}
}
