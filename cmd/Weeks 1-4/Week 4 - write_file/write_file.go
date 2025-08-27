package main

import (
	"fmt"
	"os"
)

func main() {

	message := []byte("Hello, this is a message from Go.")

	err := os.WriteFile("output.txt", message, 0644)

	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Println("File 'output.txt' written successfully!")

	// os.ReadFile reads the entire file into memory.
	// It returns the file's content as a byte slice and an error.
	content, err := os.ReadFile("not-a-real-file.txt")
	
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	
	// Convert the byte slice to a string and print it.
	fmt.Println("File content:\n" + string(content))

}