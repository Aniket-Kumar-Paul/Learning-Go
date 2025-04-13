package main

import "fmt"

func sum(numbers ...int) int { // variadic function that takes a variable number of int arguments
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
// NOTE: To take any type as input instead of just int, use interface{} as the type of the variadic parameter.

func main() {
	// Example usage of the variadic function
	fmt.Println(sum(1, 2, 3, 4, 5))

	nums := []int{1, 2, 3, 4, 5}
	result := sum(nums...) // using the variadic function with a slice
	fmt.Println(result)
}