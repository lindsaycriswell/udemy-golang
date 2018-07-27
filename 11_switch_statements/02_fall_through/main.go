package main

import "fmt"

func main() {
	switch "LC" {
	case "Ian":
		fmt.Println("This is")
	case "Peter":
		fmt.Println("a")
		fallthrough
	case "Sean":
		fmt.Println("contrived")
	case "LC":
		fmt.Println("example")
		fallthrough
	default:
		fmt.Println("EBUIS")
	}
}
