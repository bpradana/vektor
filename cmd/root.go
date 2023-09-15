package cmd

import (
	"fmt"
	"vektor/config"
	"vektor/internal/router"
	"vektor/internal/server"
)

func Start() {
	config := config.NewConfig()
	router := router.NewSocketRouter(config)
	server := server.NewSocketServer(config, router)
	err := server.Start()
	if err != nil {
		fmt.Println(err)
	}
}
