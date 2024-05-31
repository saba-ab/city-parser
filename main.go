package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/saba-ab/city-parser/parser"
)

func main() {
	directories := []string{"raw_data_ss", "raw_data_myhome"}
	var allCities []parser.City
	totalStreetCount := 0

	for _, dir := range directories {
		cities, err := parser.ParseDirectory(dir)
		if err != nil {
			log.Fatalf("Error parsing directory %s: %v", dir, err)
		}
		allCities = append(allCities, cities...)
	}

	// Calculate the total street count
	for _, city := range allCities {
		totalStreetCount += city.StreetCount
	}

	outputData := struct {
		TotalStreetCount int           `json:"total_street_count"`
		Cities           []parser.City `json:"cities"`
	}{
		TotalStreetCount: totalStreetCount,
		Cities:           allCities,
	}

	outputFile := filepath.Join("parsed_data", "streets.json")
	if err := os.MkdirAll(filepath.Dir(outputFile), 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	jsonFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Failed to create JSON file: %v", err)
	}
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(outputData)
	if err != nil {
		log.Fatalf("Failed to write to JSON file: %v", err)
	}

	log.Printf("Parsed data and created %s\n", outputFile)
}
