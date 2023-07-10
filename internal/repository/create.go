package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"vektor/internal/constants"
	"vektor/internal/entity"
)

func (r *Repository) Create(key string, vectors [][]float64) (entity.Data, error) {
	data := entity.Data{
		Key:     key,
		Vectors: vectors,
	}

	filePath := filepath.Join(r.config.DataPath, key)
	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("[vektor] [internal] [repository] [create] [Create] error: %s", err)
		return entity.Data{}, fmt.Errorf(constants.ERROR_CREATING_KEY, key)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		log.Printf("[vektor] [internal] [repository] [create] [Create] error: %s", err)
		return entity.Data{}, fmt.Errorf(constants.ERROR_ENCODING_DATA, data)
	}

	return data, nil
}
