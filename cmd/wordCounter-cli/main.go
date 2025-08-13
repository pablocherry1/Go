package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {

	fileName := "not-a-real-file.txt"

	file, err := os.Open(fileName)
	if err != nil {
		myError := fmt.Errorf("Could not open file '%s': %w", fileName, err)
		fmt.Println(myError)
		return
	} else {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		wordCount := 0

		for scanner.Scan() {
			wordCount++
		}

		fmt.Printf("The file has %d words. \n", wordCount)
	}
}
