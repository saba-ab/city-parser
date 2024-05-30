package parser

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Street struct {
	Name string `json:"name"`
}
type City struct {
	Name    string   `json:"name"`
	Streets []Street `json:"streets"`
}

func ParseFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		return fmt.Errorf("could not parse file: %v", err)
	}

	var streets []Street
	doc.Find(".sc-1499352d-34.jlQhwo").Each(func(i int, s *goquery.Selection) {
		street := s.Text()
		streets = append(streets, Street{
			Name: street,
		})
	})
	cityName := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	city := City{
		Name:    cityName,
		Streets: streets,
	}
	jsonFileName := fmt.Sprintf("%s.json", cityName)
	jsonFile, err := os.Create(jsonFileName)
	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}
	defer jsonFile.Close()
	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(city)
	if err != nil {
		return fmt.Errorf("could not write json: %v", err)
	}
	log.Printf("Successfully parsed %s and created file: %s\n", filePath, jsonFileName)
	return nil
}
