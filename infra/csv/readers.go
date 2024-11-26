package csv

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gokul656/multi-parser/domain/models"
)

func (a *CSVAdapter) Read(metadata *models.AdapterMetadata) ([]byte, error) {
	data, err := a.open(metadata.Filename, metadata.Columns)
	if err != nil {
		log.Fatalln(err)
	}

	if metadata.Limit == 0 {
		metadata.Limit = 100
	}

	// Sort the entire dataset based on the sort rules
	sortedData, err := metadata.SortCSV(data, metadata.SortRules)
	if err != nil {
		return nil, fmt.Errorf("unable to process csv: %v", err)
	}

	// Apply Limit and Offset to the sorted data
	finalData := a.applyLimitAndOffset(sortedData, metadata.Limit, metadata.Offset)

	// Convert the filtered data to JSON
	return a.ConvertToJSON(finalData), nil
}

func (a *CSVAdapter) ReadAll(metadata *models.AdapterMetadata) ([]interface{}, error) {
	return make([]interface{}, 0), nil
}

func (a *CSVAdapter) ConvertToJSON(data [][]string) []byte {
	log.Println("conveting csv to json")
	header := data[0]
	var result []map[string]string

	// Process each row (skip the header row)
	for _, chunk := range data[1:] {
		jsonData := map[string]string{}
		for i, h := range header {
			jsonData[h] = chunk[i]
		}

		result = append(result, jsonData)
	}

	// Marshal the slice to JSON
	output, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling to JSON: %v\n", err)
	}

	return output
}
