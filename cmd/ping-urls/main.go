package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done() //This tells the WaitGroup this goroutine is done
	time.Sleep(1 * time.Second) //Wait for 2 seconds
	fmt.Println("Hello from the goroutine!")
}

func main () {
	var wg sync.WaitGroup
	wg.Add(1) // Tell the WaitGroup we have 1 goroutine to wait for


	go sayHello(&wg)
	fmt.Println("Hello from the main function!")

	wg.Wait() // Wait for the goroutine to finish
}