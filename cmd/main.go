package main

import (
	"log"

	"github.com/gokul656/multi-parser/domain/models"
	"github.com/gokul656/multi-parser/infra/csv"
)

func main() {
	filedetails := csv.FileDetails{
		Name: "allocate_bonds.csv",
		Path: "allocate_bonds.csv",
	}

	metadata := &models.AdapterMetadata{
		Filename: filedetails.Path,
		Limit:    1,
		Offset:   0,
		Columns:  []string{"User Id", "Requested By"},
		SortRules: []models.SortRule{
			{
				Column: "Requested By",
				Order:  models.Descending,
			},
		},
	}

	adapter := csv.NewAdapter(&filedetails)
	chunks, _ := adapter.Read(metadata)
	log.Println(string(chunks))
}
