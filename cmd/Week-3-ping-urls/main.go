
package main

import (
	"fmt"
	"sync"
	"time"
)

// The function now accepts a pointer to a WaitGroup.
// The asterisk (*) in *sync.WaitGroup means this parameter is a pointer.
// This allows the function to modify the original WaitGroup in the main function.
// It will simulate pinging a URL and send the result back.

func pingURL(url string, wg *sync.WaitGroup, c chan <- string) {

	defer wg.Done() 
	//This tells the WaitGroup this goroutine is done 
	// even if there's an error

	start := time.Now()
	// Simulate network latency (e.g., 500ms to 1.5s)
	// In a real app, you'd make an actual HTTP request here.
	time.Sleep(time.Duration(500+time.Now().UnixNano()%1000) * time.Millisecond)
	elapsed := time.Since(start)

	// Format the result message
	result := fmt.Sprintf("Pinged %s: %v", url, elapsed)

	c <- result // Send result through the channel
}

func main () {
	var wg sync.WaitGroup

	//Create a channel that will send a time.Duration value.
	// We'll use a buffered channel here (size 3) to avoid blocking
	// if main isn't ready to receive immediately.
	results := make(chan string, 3)

	urls := []string{
		"google.com",
		"github.com",
		"golang.org",
	}

	// Loop through the URLs and launch a goroutine for each.
	for _, url := range urls {
		wg.Add(1)
		go pingURL(url, &wg, results)
		//Pass the URL, WaitGroup pointer, and channel
		// In Golang, the ampersand symbol & 
		// primarily serves as the address-of operator. 
	}

	// This goroutine will collect results from the channel.
	// We run this concurrently so main doesn't block waiting for all results
	// before the WaitGroup is done.
	go func() {
		for i := 0; i < len(urls); i++ {
			result := <- results
			// Receive a result from the channel
			fmt.Println(result)
		}
		close(results)
		// Close the channgel when all the results are received
	}()

	wg.Wait()
	//Block here until all goroutines have called wg.Done()

	// After wg.Wait() unblocks, ensure the results collection goroutine has finished.
	// This is important because the results channel might still be processing.
	// A simple way for now is to just wait for the channel to close.
	<-results // This will block until the channel is closed.
}
