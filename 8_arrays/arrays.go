package main

import "fmt"

func main() {
	// arrays -> fixed size, constant time access

	var nums [4]int // declare an array of 4 integers

	nums[0] = 1
	fmt.Println(len(nums)) // print the length of the array (4)
	fmt.Println(nums)

	myArr := [3]int{1, 2, 3} // declare and initialize an array of 3 integers
	fmt.Println(myArr)           // print the array

	// 2D arrays
	twoDarr := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(twoDarr) // print the 2D array
}
