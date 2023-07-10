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

func (r *Repository) Read(key string) (entity.Data, error) {
	filePath := filepath.Join(r.config.DataPath, key)
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		log.Printf("[vektor] [internal] [repository] [read] [Read] error: %s", err)
		return entity.Data{}, fmt.Errorf(constants.ERROR_KEY_DOES_NOT_EXIST, key)
	}
	defer file.Close()

	var data entity.Data
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		log.Printf("[vektor] [internal] [repository] [delete] [Delete] error: %s", err)
		return entity.Data{}, fmt.Errorf(constants.ERROR_DECODING_DATA, data)
	}

	return data, nil
}
