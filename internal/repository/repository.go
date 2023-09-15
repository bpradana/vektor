package repository

import "vektor/config"

type Repository struct {
	config config.Config
}

func NewRepository(config config.Config) RepositoryContract {
	return &Repository{
		config: config,
	}
}
