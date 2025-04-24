package main

import "fmt"

// by value
func changeNum(num int) {
	num = 5
	fmt.Println("Inside changeNum:", num)
}

// by reference
func changeNumRef(num *int) {
	*num = 5
	fmt.Println("Inside changeNumRef:", *num)
}

func main() {
	num := 1
	changeNum(num)
	fmt.Println("In main:", num)
	changeNumRef(&num)
	fmt.Println("In main:", num)
}