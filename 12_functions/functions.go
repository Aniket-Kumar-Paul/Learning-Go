package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func getLanguages() (string, string, string) { // returning multiple values
	return "golang", "python", "java"
}

// This function is taking  a function as an argument(that takes an int and returns an int)
func processIt(myFunc func(a int) int) {
	myFunc(5)
}

// This function retuns a function
func returnFunc() func(a int) int {
	return func(a int) int {
		return a * 2
	}
}

func main() {
	result := add(5, 3)
	println("The sum is:", result)

	lang1, lang2, lang3 := getLanguages()
	println("Languages:", lang1, lang2, lang3)

	fn := func(a int) int {
		return a * 2
	}
	processIt(fn)

	returnFunc := returnFunc()
	fmt.Println("Returned function result:", returnFunc(5))
}