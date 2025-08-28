package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {

	// 1. Make an HTTP GET request
	resp, err := http.Get("http://httpbin.org/get")

	// 2. Handle network-level errors
	if err != nil {
		fmt.Printf("Error making request: %v \n", err)
		return
	}

	// 'defer' ensures the response body is closed when the function returns,
	// which is a crucial best practice for freeing up network resources.
	defer resp.Body.Close()

	// 3. Check the HTTP status code for errors
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Received non-200 status code: %d \n", resp.StatusCode)
		return
	}

	// 4. Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body %v \n", err)
		return
	}

	// 5. Print the body content
	fmt.Println("Response Body:")
	fmt.Println(string(body))


}