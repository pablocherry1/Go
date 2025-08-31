package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type Config struct {
	APIKey string `json:"api_key"`
	Port int `json:"port"`
	Enabled bool `json:"enabled"`
}

func main() {
	
	// 1. Read the JSON file into a byte slice
	fileContent, err := os.ReadFile("config.json")

	if err != nil {
		fmt.Printf("Error reading file: %v \n", err)
		return
	}

	// 2. Create a variable to hold the parsed data
	var config Config

	// 3. Unmarshal (parse) the JSON data into the struct
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		fmt.Printf("Error Unmarshaling JSON: %v \n", err)
		return
	}

	// 4. Print the parsed data
	fmt.Printf("Parsed Data:\n %+v \n", config)
    fmt.Println("API Key:", config.APIKey)
    fmt.Println("Port:", config.Port)
    fmt.Println("Enabled:", config.Enabled)

}