package main

import (
	"fmt",
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done() // defer makes this run once the function completes executing (like a cleanup function in useEffect of react)
	// wg.Done() decrements by 1
	fmt.Println("Doing task", id)
}

func main() {
	var wg sync.WaitGroup

	for i:=0; i<=10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait() // waits here until wg value goes back to default (0), i.e all wg above are done
}