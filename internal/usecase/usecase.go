package usecase

import "vektor/internal/repository"

type Usecase struct {
	repository repository.RepositoryContract
}

func NewUsecase(repository repository.RepositoryContract) UsecaseContract {
	return &Usecase{
		repository: repository,
	}
}
