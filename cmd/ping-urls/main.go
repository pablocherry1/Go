package main

import (
	"fmt"
	"sync"
	"time"
)

// The sayHello function now accepts a pointer to a WaitGroup.
// The asterisk (*) in *sync.WaitGroup means this parameter is a pointer.
// This allows the function to modify the original WaitGroup in the main function.
func sayHello(wg *sync.WaitGroup) {
	defer wg.Done() //This tells the WaitGroup this goroutine is done
	time.Sleep(1 * time.Second) //Wait for 2 seconds
	fmt.Println("Hello from the goroutine!")
}

func main () {
	var wg sync.WaitGroup
	wg.Add(1) // Tell the WaitGroup we have 1 goroutine to wait for

	// The 'go' keyword starts a new goroutine.
	// The ampersand (&) in &wg gets the memory address of the wg variable.
	// We pass this address (the pointer) to the sayHello function.
	go sayHello(&wg) //
	fmt.Println("Hello from the main function!")

	wg.Wait() // Wait until the WaitGroup's counter is back to zero
}