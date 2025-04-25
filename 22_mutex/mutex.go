package main

import (
	"fmt"
	"sync"
)

// Mutex -> to prevent race conditions
type post struct {
	views int
	mu sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	// Multiple go routines can't modify now at the same time until it's unlocked
	p.mu.Lock()
	p.views += 1
}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}
	for i:=0;i<100;i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait()

	fmt.Println(myPost.views)
}