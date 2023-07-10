package cmd

import (
	"fmt"
	"vektor/config"
	"vektor/internal/router"
	"vektor/internal/server"
)

func Start() {
	config := config.NewConfig()
	router := router.NewRouter(config)
	server := server.NewServer(config, router)
	err := server.Start()
	if err != nil {
		fmt.Println(err)
	}
}
