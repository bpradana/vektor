package usecase

import (
	"vektor/internal/entity"
)

func (u *Usecase) Read(key string) (entity.Data, error) {
	data, err := u.repository.Read(key)
	if err != nil {
		return entity.Data{}, err
	}

	return data, nil
}
