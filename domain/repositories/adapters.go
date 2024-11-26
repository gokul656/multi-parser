package repositories

import "github.com/gokul656/multi-parser/domain/models"

type Adapter[T any] interface {
	Read(metadata *models.AdapterMetadata) (T, error)
	ReadAll(metadata *models.AdapterMetadata) ([]T, error)
	ReadChunked(metadata models.AdapterMetadata) ([]T, error)
}
