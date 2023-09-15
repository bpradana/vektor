package usecase

import (
	"vektor/internal/entity"
)

func (u *Usecase) Update(key string, vectors [][]float64) (entity.Data, error) {
	data, err := u.repository.Update(key, vectors)
	if err != nil {
		return entity.Data{}, err
	}

	return data, nil
}
