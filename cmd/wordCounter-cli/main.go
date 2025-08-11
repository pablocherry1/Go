package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("not-a-real-file.txt")
	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
	fmt.Println("File opened successfully!")
	file.Close()
}
