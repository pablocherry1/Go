package main

import(
	"fmt"
	"os"
	"gopkg.in/yaml.v2" //go get gopkg.in/yaml.v2 on bash
)

// Define structs that match the YAML structure
type PrometheusConfig struct {
	Global	GlobalConfig	`yaml:"global"`
	ScrapeConfigs []ScrapeConfig	`yaml:"scrape_configs"`
}

type GlobalConfig struct {
	ScrapeInterval string `yaml:"scrape_interval"`
}

type ScrapeConfig struct {
	JobName	string	`yaml:"job_name"`
	ScrapeInterval	string	`yaml:"scrape_interval"`
	StaticConfigs []StaticConfig	`yaml:"static_configs"`
}

type StaticConfig struct {
	Targets []string `yaml:"targets"`
}

func main() {
	// 1. Read the YAML file
	fileContent, err := os.ReadFile("prometheus.yml")

	if err != nil {
		fmt.Printf("Error reading file: %v \n", err)
		return
	}

	// 2. Unmarshal the YAML into our struct
	var config PrometheusConfig
	err = yaml.Unmarshal(fileContent, &config)
	if err != nil {
		fmt.Printf("Error unmarshaling YAML: %v \n", err)
		return
	}

	// 3. Print a specific value from the parsed data
	fmt.Println("Scrape interval for 'prometheus' job:", config.ScrapeConfigs[0].ScrapeInterval)
	fmt.Println(config)

}
