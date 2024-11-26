package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func (a *CSVAdapter) open(filePath string, columns []string) ([][]string, error) {
	log.Printf("opening file: %v\n", filePath)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read the header row
	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read headers: %v", err)
	}

	// Map column names to indices
	columnIndices := make([]int, len(columns))
	for i, column := range columns {
		index := a.indexOf(column, headers)
		if index == -1 {
			return nil, fmt.Errorf("column '%v' not found in headers", column)
		}
		columnIndices[i] = index
	}

	// Collect all rows
	data := [][]string{}
	data = append(data, columns) // Include header

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading file: %v", err)
		}

		line := make([]string, 0)
		for _, i := range columnIndices {
			line = append(line, record[i])
		}

		data = append(data, line)
	}

	return data, nil
}

func (a *CSVAdapter) applyLimitAndOffset(data [][]string, limit, offset int) [][]string {
	// If Offset is beyond data length, return only the header row
	if offset >= len(data)-1 { // Exclude the header row from offset check
		return data[:1]
	}

	start := offset + 1 // Offset is 0-based; skip the header row

	// Calculate the end index based on Limit
	end := len(data)
	if limit > 0 && start+limit < len(data) {
		end = start + limit
	}

	// Include header row and apply offset/limit
	return append(data[:1], data[start:end]...)
}

func (a *CSVAdapter) indexOf(columnName string, headers []string) int {
	for i, header := range headers {
		if header == columnName {
			return i
		}
	}
	return -1
}
