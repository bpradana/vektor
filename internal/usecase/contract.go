package usecase

import (
	"vektor/internal/entity"
)

type UsecaseContract interface {
	Create(key string, vectors [][]float64) (entity.Data, error)
	Read(key string) (entity.Data, error)
	Update(key string, vectors [][]float64) (entity.Data, error)
	Delete(key string) error
	Search(vector []float64, method string, threshold float64) (entity.Data, error)
}
