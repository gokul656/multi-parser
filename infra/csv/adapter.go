package csv

import "errors"

type FileDetails struct {
	Name string
	Size float32
	Path string
}

type CSVAdapter struct {
	Files map[string]*FileDetails
}

func (a *CSVAdapter) Get(name string) (*FileDetails, error) {
	details := a.Files[name]
	if details == nil {
		return nil, errors.New("file not found")
	}

	return details, nil
}

func NewAdapter(filedetails *FileDetails) *CSVAdapter {
	files := make(map[string]*FileDetails)
	files[filedetails.Name] = filedetails

	return &CSVAdapter{
		Files: files,
	}
}
