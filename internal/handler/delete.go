package handler

import (
	"fmt"
	"net"
	"time"
	"vektor/internal/common"
	"vektor/internal/usecase"
)

type DeleteHandler struct {
	usecase usecase.UsecaseContract
}

func NewDeleteHandler(usecase usecase.UsecaseContract) HandlerContract {
	return &DeleteHandler{
		usecase: usecase,
	}
}

func (h *DeleteHandler) Handle(conn net.Conn, query common.Query) error {
	start_time := time.Now()

	err := h.usecase.Delete(query.Key)
	if err != nil {
		conn.Write([]byte(err.Error()))
		conn.Close()
		return err
	}

	run_time := time.Since(start_time).Milliseconds()
	conn.Write([]byte(fmt.Sprintf("%s %s success, took %d ms", query.Action, query.Key, run_time)))
	conn.Close()

	return nil
}
