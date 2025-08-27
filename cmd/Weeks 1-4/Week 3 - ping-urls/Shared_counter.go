package main

import (
	"fmt"
	"sync"
)

// A shared counter that all goroutines will try to modify
var counter int

// A mutex to protect the shared counter
var mu sync.Mutex

func add(wg *sync.WaitGroup) {
	defer wg.Done()
	
	for i := 0; i < 1000; i++ {
		// We lock the mutex to protect the critical section
		mu.Lock()
		
		counter++ // This is now safe
		
		// We unlock the mutex to allow other goroutines to access it
		mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	
	// Launch 100 goroutines, each adding to the counter
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add(&wg)
	}
	
	wg.Wait()
	
	fmt.Println("Final Counter:", counter)
}