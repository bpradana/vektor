package repository

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"vektor/internal/constants"
)

func (r *Repository) Files() ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(r.config.DataPath)
	if err != nil {
		log.Printf("[vektor] [internal] [repository] [files] [Files] error: %s", err)
		return nil, fmt.Errorf(constants.ERROR_READING_DB)
	}

	return files, nil
}
