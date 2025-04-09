package main

import (
	"fmt"
	"slices"
)

func main() {
	// Dynamic arrays
	// Uninitialized slice is nil
	var nums []int
	fmt.Println(nums) // []	

	var nums2 = make([]int, 2) // size 2
	fmt.Println(nums2) // [0 0]
	fmt.Println(len(nums2)) // 2
	fmt.Println(cap(nums2)) // 2 (capacity - max num of elements in the slice (this too increases dynamically))

	var nums3 = make([]int, 2, 5) // initial size 2 (generally, we keep this 0), capacity 5
	fmt.Println(nums3) // [0 0]
	nums3 = append(nums3, 1, 2, 3, 4) // append elements
	fmt.Println(nums3) // [0 0 1 2 3 4]
	fmt.Println(len(nums3)) // 6
	fmt.Println(cap(nums3)) // 10 (capacity - max num of elements in the slice (this too increases dynamically))

	nums4 := []int{}
	nums4 = append(nums4, 1, 2, 3, 4) // append elements
	fmt.Println(nums4) // [1 2 3 4]

	// copy
	nums5 := make([]int, len(nums4)) // create a slice of same length as nums4
	copy(nums5, nums4) // copy nums4 to nums5
	fmt.Println(nums4, nums5)

	// slice operator
	var nums6 = []int{1, 2, 3, 4, 5}
	fmt.Println(nums6[1:3]) // [2 3] (start index inclusive, end index exclusive)

	// slices package
	fmt.Println(slices.Equal(nums4, nums5)) // true

	// 2D slices
	var nums7 = [][]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(nums7) // [[1 2] [3 4] [5 6]]
}