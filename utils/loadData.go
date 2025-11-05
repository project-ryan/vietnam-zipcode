package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/quanghia24/vietnam-zipcode/db"
)

func LoadData(filename string) (db.Data, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read data file %s: %w", filename, err)
	}

	var data db.Data

	if err := json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON data: %w", err)
	}

	return data, nil
}
