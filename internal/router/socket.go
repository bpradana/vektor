package router

import (
	"log"
	"net"
	"strings"
	"vektor/config"
	"vektor/internal/common"
	"vektor/internal/constants"
	"vektor/internal/handler/socket"
	"vektor/internal/repository"
	"vektor/internal/usecase"
)

type SocketRouter struct {
	cfg config.Config
}

func NewSocketRouter(cfg config.Config) *SocketRouter {
	return &SocketRouter{
		cfg: cfg,
	}
}

func (r *SocketRouter) Route(conn net.Conn) error {
	repository := repository.NewRepository(r.cfg)

	usecase := usecase.NewUsecase(repository)

	createHandler := socket.NewCreateHandler(usecase)
	readHandler := socket.NewReadHandler(usecase)
	updateHandler := socket.NewUpdateHandler(usecase)
	deleteHandler := socket.NewDeleteHandler(usecase)
	searchHandler := socket.NewSearchHandler(usecase)

	bufferString := ""

	for {
		buffer := make([]byte, r.cfg.BufferSize)
		numBytes, err := conn.Read(buffer)
		if err != nil {
			conn.Close()
			return err
		}

		bufferString += string(buffer[:numBytes])

		queries := strings.Split(bufferString, ";")

		for _, query := range queries[:len(queries)-1] {
			query, err := common.ParseQuery(query)
			if err != nil {
				conn.Write([]byte(err.Error()))
				continue
			}

			switch query.Action {
			case constants.ACTION_CREATE:
				err := createHandler.Handle(conn, *query)
				if err != nil {
					log.Println(err)
					continue
				}
			case constants.ACTION_READ:
				err := readHandler.Handle(conn, *query)
				if err != nil {
					continue
				}
			case constants.ACTION_UPDATE:
				err := updateHandler.Handle(conn, *query)
				if err != nil {
					log.Println(err)
					continue
				}
			case constants.ACTION_DELETE:
				err := deleteHandler.Handle(conn, *query)
				if err != nil {
					log.Println(err)
					continue
				}
			case constants.ACTION_SEARCH:
				err := searchHandler.Handle(conn, *query)
				if err != nil {
					log.Println(err)
					continue
				}
			default:
				conn.Write([]byte(constants.ERROR_INVALID_ACTION))
				continue
			}
		}

		bufferString = queries[len(queries)-1]
	}
}
