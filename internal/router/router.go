package router

import (
	"net"
	"vektor/config"
	"vektor/internal/common"
	"vektor/internal/constants"
	"vektor/internal/handler"
	"vektor/internal/repository"
	"vektor/internal/usecase"
)

type Router struct {
	cfg config.Config
}

func NewRouter(cfg config.Config) *Router {
	return &Router{
		cfg: cfg,
	}
}

func (r *Router) Route(conn net.Conn) error {
	repository := repository.NewRepository(r.cfg)

	usecase := usecase.NewUsecase(repository)

	createHandler := handler.NewCreateHandler(usecase)
	readHandler := handler.NewReadHandler(usecase)
	updateHandler := handler.NewUpdateHandler(usecase)
	deleteHandler := handler.NewDeleteHandler(usecase)
	searchHandler := handler.NewSearchHandler(usecase)

	buffer := make([]byte, 4096)

	for {
		numBytes, err := conn.Read(buffer)
		if err != nil {
			conn.Close()
			return err
		}

		query, err := common.ParseQuery(string(buffer[:numBytes]))
		if err != nil {
			conn.Close()
			return err
		}

		switch query.Action {
		case constants.ACTION_CREATE:
			err := createHandler.Handle(conn, *query)
			if err != nil {
				return err
			}
		case constants.ACTION_READ:
			err := readHandler.Handle(conn, *query)
			if err != nil {
				return err
			}
		case constants.ACTION_UPDATE:
			err := updateHandler.Handle(conn, *query)
			if err != nil {
				return err
			}
		case constants.ACTION_DELETE:
			err := deleteHandler.Handle(conn, *query)
			if err != nil {
				return err
			}
		case constants.ACTION_SEARCH:
			err := searchHandler.Handle(conn, *query)
			if err != nil {
				return err
			}
		}
	}
}
