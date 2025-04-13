package main

import "fmt"

func main() {
	// Iterating slices
	nums := []int{4, 5, 6}
	for idx, num := range nums {
		fmt.Println(idx, num)
	}

	// Iterating maps
	m := map[string]int{"a": 1, "b": 2}
	for key, value := range m {
		fmt.Println(key, value)
	}

	// Iterating strings
	for idx, char := range "hello" {
		fmt.Println(idx, string(char)) // char is a unicode character
	}	
}