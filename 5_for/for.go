package main

import "fmt"

func main() {
	// Go only supports the for loop

	// while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// infinite loop
	// for {
	// 	fmt.Println("Hello")
	// }

	// classic for loop
	for i:=0; i<3; i++ {
		fmt.Println(i)
		if i==1 {
			break
		}
	}

	// for range loop
	for i := range 3{
		fmt.Println(i) // 0, 1, 2
	}
}