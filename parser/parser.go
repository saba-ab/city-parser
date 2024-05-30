package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Street represents a street in the city.
type Street struct {
	Name string `json:"name"`
}

// City represents a city with its streets.
type City struct {
	Name        string   `json:"name"`
	Streets     []Street `json:"streets"`
	StreetCount int      `json:"street_count"`
}

// ParseFile parses a single HTML file and returns a City struct.
func ParseFile(filePath string) (City, error) {
	// Open the HTML file
	file, err := os.Open(filePath)
	if err != nil {
		return City{}, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Parse the HTML file
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		return City{}, fmt.Errorf("failed to parse HTML: %w", err)
	}
	cityName := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))

	var streets []Street
	// Determine which parsing method to use based on the structure of the HTML
	if doc.Find(".sc-1499352d-34.jlQhwo").Length() > 0 {
		streets = parseStreetsType1(doc, cityName)
	} else if doc.Find("div.flex.items-center label").Length() > 0 {
		streets = parseStreetsType2(doc, cityName)
	} else if doc.Find(".sc-111c43f2-34.jMjmXo").Length() > 0 { // Add condition for new HTML structure
		streets = parseStreetsType3(doc, cityName)
	} else {
		return City{}, fmt.Errorf("unsupported HTML structure")
	}

	city := City{
		Name:        cityName,
		Streets:     streets,
		StreetCount: len(streets),
	}

	return city, nil
}

// parseStreetsType1 parses the first type of HTML document and extracts the street names.
func parseStreetsType1(doc *goquery.Document, cityName string) []Street {
	var streets []Street
	doc.Find(".sc-1499352d-34.jlQhwo").Each(func(i int, s *goquery.Selection) {
		street := s.Text()
		if street != " " && cityName != street {
			streets = append(streets, Street{Name: street})
		}
	})
	return streets
}

// parseStreetsType2 parses the second type of HTML document and extracts the street names.
func parseStreetsType2(doc *goquery.Document, cityName string) []Street {
	var streets []Street
	doc.Find("div.flex.items-center label").Each(func(i int, s *goquery.Selection) {
		street := s.Text()
		if street != " " && cityName != street {
			streets = append(streets, Street{Name: street})
		}
	})
	return streets
}

// parseStreetsType3 parses the third type of HTML document and extracts the street names.

func parseStreetsType3(doc *goquery.Document, cityName string) []Street {
	var streets []Street
	doc.Find(".sc-111c43f2-34.jMjmXo").Each(func(i int, s *goquery.Selection) {
		street := s.Text()
		if street != " " && cityName != street {
			streets = append(streets, Street{Name: street})
		}
	})
	return streets
}

// ParseDirectory parses all HTML files in a directory and returns a list of cities.
func ParseDirectory(dir string) ([]City, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var cities []City
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".html" {
			filePath := filepath.Join(dir, file.Name())
			city, err := ParseFile(filePath)
			if err != nil {
				log.Printf("Error parsing file %s: %v\n", filePath, err)
				continue
			}
			cities = append(cities, city)
		}
	}
	return cities, nil
}
