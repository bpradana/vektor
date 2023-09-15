package repository

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"vektor/internal/constants"
)

func (r *Repository) Delete(key string) error {
	filePath := filepath.Join(r.config.DataPath, key)
	err := os.Remove(filePath)
	if err != nil {
		log.Printf("[vektor] [internal] [repository] [delete] [Delete] error: %s", err)
		return fmt.Errorf(constants.ERROR_KEY_DOES_NOT_EXIST, key)
	}

	return nil
}
