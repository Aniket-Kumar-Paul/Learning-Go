package main

import "fmt"

func printAnySlice[T any](items []T) { // instead of any, we can also use interface{}
	for _, item := range items {
		fmt.Println(item)
	}
}

// Generics in function
// Instead of passing in multiple types one by one, we can also directly use the `comparable` type
// comparable -> an interface implemented by all comparable types (bool, nos., strings, pointers, channels, arrays of comparable types, structs whose fields are all comparable types)
func printSlice[T comparable int|string|bool](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// Generics in struct
type stack[T any] struct {
	elements []T
}

func main() {
	nums := []int{1,2,3}
	printSlice(nums)
	
	myStack := stack[string] {
		elements: []string{"golang", "c++"}
	}
}