package repository

import (
	"io/fs"
	"vektor/internal/entity"
)

type RepositoryContract interface {
	Create(key string, vectors [][]float64) (entity.Data, error)
	Read(key string) (entity.Data, error)
	Update(key string, vectors [][]float64) (entity.Data, error)
	Delete(key string) error
	Files() ([]fs.FileInfo, error)
}
