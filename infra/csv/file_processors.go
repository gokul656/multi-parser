package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func (a *CSVAdapter) open(filePath string, columns []string, chunkSize *int, offset *int) ([][]string, error) {
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

	// Data collection
	data := [][]string{}
	data = append(data, columns)

	if chunkSize == nil {
		return a.read(data, columnIndices, *reader)
	}

	return a.readChunked(*chunkSize, data, columnIndices, *reader, *offset)
}

func (a *CSVAdapter) read(data [][]string, columnIndices []int, reader csv.Reader) ([][]string, error) {

	for {
		record, err := reader.Read()
		if err == io.EOF {
			// Process the remaining chunk
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading file: %v", err)
		}

		line := make([]string, 0)
		for _, i := range columnIndices {
			line = append(line, record[i])
		}

		// Add record to the chunk
		data = append(data, line)
	}

	return data, nil
}

func (a *CSVAdapter) readChunked(chunkSize int, data [][]string, columnIndices []int, reader csv.Reader, offset int) ([][]string, error) {
	if chunkSize <= 0 {
		return nil, fmt.Errorf("chunkSize must be greater than 0")
	}

	skippedRows := 0

	// Skip rows based on the offset
	for skippedRows < offset {
		_, err := reader.Read()
		if err == io.EOF {
			// If EOF is reached before the offset, return current data
			return data, nil
		}
		if err != nil {
			return nil, fmt.Errorf("error skipping rows: %v", err)
		}
		skippedRows++
	}

	rowCount := 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading file: %v", err)
		}

		// Add the record to the result set
		line := make([]string, 0)
		for _, i := range columnIndices {
			line = append(line, record[i])
		}

		data = append(data, line)
		rowCount++

		// Stop reading after reaching chunk size
		if rowCount >= chunkSize {
			break
		}
	}

	return data, nil
}

func (a *CSVAdapter) indexOf(columnName string, headers []string) int {
	for i, header := range headers {
		if header == columnName {
			return i
		}
	}
	return -1
}
