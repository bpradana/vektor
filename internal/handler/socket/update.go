package socket

import (
	"fmt"
	"net"
	"time"
	"vektor/internal/common"
	"vektor/internal/usecase"
)

type UpdateHandler struct {
	usecase usecase.UsecaseContract
}

func NewUpdateHandler(usecase usecase.UsecaseContract) HandlerContract {
	return &UpdateHandler{
		usecase: usecase,
	}
}

func (h *UpdateHandler) Handle(conn net.Conn, query common.Query) error {
	start_time := time.Now()

	data, err := h.usecase.Update(query.Key, query.Vectors)
	if err != nil {
		conn.Write([]byte(err.Error()))
		return err
	}

	run_time := time.Since(start_time).Milliseconds()
	conn.Write([]byte(fmt.Sprintf("%s %s success, took %d ms", query.Action, data.Key, run_time)))

	return nil
}
