package main

import (
	"fmt"
	"time"
)

func main(){

	// Create two channels:
	// one for messages and one for signals
	messages := make(chan string)
	signals := make(chan string)

	// Goroutine to send a message after 1 second
	go func() {
		time.Sleep(4 * time.Second)
		messages <- "message from goroutine"
	}()

	// Goroutine to send a signal after 2 seconds
	go func() {
		time.Sleep(4 * time.Second)
		signals <- "signal from goroutine"
	}()

	// The 'select' statement listens to both channels
	// and executes the case for the channel that is ready first.
	select {
	case msg := <-messages:
		fmt.Println("Received message: ", msg)
	case sig := <-signals:
		fmt.Println("Received signal: ", sig)
	case <- time.After(3 * time.Second):
		fmt.Println("No activity. Timeout!")
	}


}