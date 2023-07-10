package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DataPath string `json:"data_path"`
}

func NewConfig() Config {
	config_file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer config_file.Close()

	config := Config{}
	decoder := json.NewDecoder(config_file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
