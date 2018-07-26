package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue // moves to next iteration of the loop (only prints odd numbers in this example)
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
}
