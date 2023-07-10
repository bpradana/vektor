package server

import (
	"fmt"
	"net"
	"vektor/config"
	"vektor/internal/router"
)

type Server struct {
	config config.Config
	router router.Router
}

func NewServer(config config.Config, router *router.Router) ServerContract {
	return &Server{
		config: config,
		router: *router,
	}
}

func (s *Server) Start() error {
	address := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go s.router.Route(conn)
	}
}
