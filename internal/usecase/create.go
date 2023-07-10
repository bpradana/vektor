package usecase

import (
	"fmt"
	"vektor/internal/constants"
	"vektor/internal/entity"
)

func (u *Usecase) Create(key string, vectors [][]float64) (entity.Data, error) {
	_, err := u.repository.Read(key)
	if err == nil {
		return entity.Data{}, fmt.Errorf(constants.ERROR_KEY_ALREADY_EXIST, key)
	}

	data, err := u.repository.Create(key, vectors)
	if err != nil {
		return entity.Data{}, err
	}

	return data, nil
}
