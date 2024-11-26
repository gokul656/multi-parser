package models

import (
	"log"
	"sort"
)

type SortOrder int

const (
	Ascending SortOrder = iota
	Descending
)

type SortRule struct {
	Column string
	Order  SortOrder
}

type AdapterMetadata struct {
	Filename  string     `json:"filename,omitempty"`
	Columns   []string   `json:"columns,omitempty"`
	Limit     int        `json:"limit,omitempty"`
	Offset    int        `json:"offset,omitempty"`
	SortRules []SortRule `json:"sort_rules,omitempty"`
	SortedBy  []string   `json:"sorted_by,omitempty"`
}

func (m *AdapterMetadata) SortCSV(data [][]string, rules []SortRule) ([][]string, error) {
	log.Println("sorting csv")
	// Create a mapping of column names to indexes
	columnIndex := make(map[string]int)
	for i, col := range data[0] {
		columnIndex[col] = i
	}

	// Sort data rows (excluding the header row)
	sort.SliceStable(data[1:], func(i, j int) bool {
		for _, rule := range rules {
			colIdx, ok := columnIndex[rule.Column]
			if !ok {
				// Skip unknown columns
				continue
			}

			valI := data[1+i][colIdx]
			valJ := data[1+j][colIdx]

			// Compare values based on sort order
			if rule.Order == Ascending {
				if valI != valJ {
					return valI < valJ
				}
			} else {
				if valI != valJ {
					return valI > valJ
				}
			}
		}
		return false // If all columns are equal, maintain current order
	})

	return data, nil
}
