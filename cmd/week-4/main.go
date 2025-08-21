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

}