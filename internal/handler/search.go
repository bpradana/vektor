package handler

import (
	"fmt"
	"net"
	"time"
	"vektor/internal/common"
	"vektor/internal/usecase"
)

type SearchHandler struct {
	usecase usecase.UsecaseContract
}

func NewSearchHandler(usecase usecase.UsecaseContract) HandlerContract {
	return &SearchHandler{
		usecase: usecase,
	}
}

func (h *SearchHandler) Handle(conn net.Conn, query common.Query) error {
	start_time := time.Now()

	data, err := h.usecase.Search(query.Vectors[0], query.Option)
	if err != nil {
		conn.Write([]byte(err.Error()))
		conn.Close()
		return err
	}

	run_time := time.Since(start_time).Milliseconds()
	conn.Write([]byte(fmt.Sprintf("%s %s success, took %d ms", query.Action, data.Key, run_time)))
	conn.Close()

	return nil
}
