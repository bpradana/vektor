package server

import (
	"net"
	"vektor/config"
	"vektor/internal/handler"
)

type Server struct {
	config  config.Config
	handler handler.HandlerContract
}

func NewServer(config config.Config, handler handler.HandlerContract) ServerContract {
	return &Server{
		config:  config,
		handler: handler,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.config.Port)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go s.handler.Handle(conn)
	}
}
