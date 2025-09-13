package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	// Let's get a fake post from a public API
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	
	if err != nil {
		fmt.Printf("Error making request: %v \n:", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Received non-200 status code: %d \n", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v \n", err)
		return
	}

	fmt.Println("Response Body:")
	fmt.Println(string(body))

}