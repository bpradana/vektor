package server

import (
	"fmt"
	"log"
	"net"
	"vektor/config"
	"vektor/internal/router"
)

type SocketServer struct {
	config config.Config
	router router.SocketRouter
}

func NewSocketServer(config config.Config, router *router.SocketRouter) ServerContract {
	return &SocketServer{
		config: config,
		router: *router,
	}
}

func (s *SocketServer) Start() error {
	address := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Printf("[server] [socket] [Start] error: %s", err)
		return err
	}

	log.Printf("[server] [socket] [Start] listening on %s", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("[server] [socket] [Start] error: %s", err)
			return err
		}

		go s.router.Route(conn)
	}
}
