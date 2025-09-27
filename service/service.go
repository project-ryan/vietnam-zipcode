package service

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

// Location represents a geographic location with province and ward
type Location struct {
	Province string `json:"province"`
	Ward     string `json:"ward"`
}

// ZipcodeService defines the interface for zipcode operations
type ZipcodeService interface {
	GetLocation(zipcode string) (*Location, error)
	LoadData(filename string) error
}

// MemoryZipcodeService implements ZipcodeService using in-memory storage
type MemoryZipcodeService struct {
	zipToLocation map[string]Location
	zipcodeRegex  *regexp.Regexp
}

// NewMemoryZipcodeService creates a new instance of MemoryZipcodeService
func NewMemoryZipcodeService() *MemoryZipcodeService {
	return &MemoryZipcodeService{
		zipToLocation: make(map[string]Location),
		zipcodeRegex:  regexp.MustCompile(`^\d{5}$`), // Vietnam zipcodes are 5 digits
	}
}

// LoadData loads zipcode data from a JSON file
func (s *MemoryZipcodeService) LoadData(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read data file %s: %w", filename, err)
	}

	if err := json.Unmarshal(data, &s.zipToLocation); err != nil {
		return fmt.Errorf("failed to parse JSON data: %w", err)
	}

	fmt.Println("Database is ready")

	return nil
}

// GetLocation retrieves location information for a given zipcode
func (s *MemoryZipcodeService) GetLocation(zipcode string) (*Location, error) {
	if !s.zipcodeRegex.MatchString(zipcode) {
		return nil, fmt.Errorf("invalid zipcode format: %s", zipcode)
	}

	location, exists := s.zipToLocation[zipcode]
	if !exists {
		return nil, fmt.Errorf("zipcode not found: %s", zipcode)
	}

	return &location, nil
}

// GetDataSize returns the number of zipcodes loaded
func (s *MemoryZipcodeService) GetDataSize() int {
	return len(s.zipToLocation)
}
