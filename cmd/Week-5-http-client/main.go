package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	// Make an HTTP GET request to a public service
	resp, err := http.Get("http://httpbin.org/get")

	// Handle network-level errors
	if err != nil {
		fmt.Printf("Error making request: %v \n", err)
		return
	}

	// Ensure the response body is closed to free up resources
	defer resp.Body.Close()

	// Check the HTTP status code for errors
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Received non-200 status code: %d \n", resp.StatusCode)
		return
	}

	// Read the response body from the server
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v \n", err)
		return
	}

	// Print the body content
	fmt.Println("Response Body:")
	fmt.Println(string(body))
	fmt.Print(resp)
}