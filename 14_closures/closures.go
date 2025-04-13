package main

import "fmt"

// Closures in Go are anonymous functions that capture and use variables from their surrounding scope. They're used for maintaining state across function calls or for creating function factories.
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	increment := counter()
	fmt.Println(increment()) // Output: 1
	fmt.Println(increment()) // Output: 2
}