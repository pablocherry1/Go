package main

import (
	"fmt"
	"sync"
	"time"
)

// The sayHello function now accepts a pointer to a WaitGroup.
// The asterisk (*) in *sync.WaitGroup means this parameter is a pointer.
// This allows the function to modify the original WaitGroup in the main function.


func sayHello(wg *sync.WaitGroup, c chan<- time.Duration) {

	defer wg.Done() //This tells the WaitGroup this goroutine is done

	start := time.Now()
	time.Sleep(2 * time.Second) //Wait for 2 seconds
	
	fmt.Println("Hello from the goroutine!")
	elapsed := time.Since(start)
	c <- elapsed // Send the elapsed time through the channel
}

func main () {
	var wg sync.WaitGroup

	//Create a channel that will send a time.Duration value.
	c := make(chan time.Duration)

	wg.Add(1) // Tell the WaitGroup we have 1 goroutine to wait for

	// The 'go' keyword starts a new goroutine.
	// The ampersand (&) in &wg gets the memory address of the wg variable.
	// We pass this address (the pointer) to the sayHello function.
	// The c is to pass the channel to the goroutine
	go sayHello(&wg, c)
	fmt.Println("Hello from the main function!")

	//Receive the value from the channel. 
	//This will block until a value is sent.
	duration := <-c

	wg.Wait() // Wait until the WaitGroup's counter is back to zero

	fmt.Printf("The goroutine took %v to complete. \n", duration)
}

