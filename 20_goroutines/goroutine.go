package main

import (
	"fmt",
	"time"
)

func task(id int) {
	fmt.Println("Doing task", id)
}

// Goroutines are functions or methods that run concurrently with other functions in Go. You launch a goroutine by prepending the go keyword to a function call
// They are lightweight threads because of :
// 1. Lower Memory Footprint
	// OS threads typically start with a stack size of 1MB.
	// Goroutines start with only 2 KB of stack memory, and they grow/shrink dynamically.
// 2. Managed by Go Runtime (Green Threads)
	// Go uses its own scheduler, called an M:N scheduler:
		// M goroutines multiplexed onto N OS threads.
	// This means hundreds of thousands of goroutines can be handled efficiently with a few OS threads.
// 3.  Cheap to Create and Destroy
	// You can spawn millions of goroutines in a Go program with little overhead.
	// Creating an OS thread is expensive (needs system calls, memory allocation, etc.). Goroutines are just function calls scheduled by the runtime.
func main() {
	for i:=0; i<=10; i++ {
		// go task(i)

		// directly run anonymous function
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second * 2)
}