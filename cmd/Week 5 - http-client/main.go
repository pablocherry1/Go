package main

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

// Define a struct to hold the data from the JSON response
type Post struct {
	UserID		int		`json:"userID"`
	ID			int		`json:"id"`
	Title		string		`json:"title"`
	Body		string		`json:"body"`
}

func main() {

	// 1. Make an HTTP GET request
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

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

	// Create a variable to hold the parsed data
	var post Post

	// Unmarshal the JSON data from the body into the struct
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v \n", err)
		return
	}
	
	// Print the parsed data from the struct
	fmt.Printf("Parsed Data: \n %+v \n ", post)



}