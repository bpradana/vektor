package handler

import (
	"fmt"
	"net"
	"time"
	"vektor/internal/common"
	"vektor/internal/constants"
	"vektor/internal/entity"
	"vektor/internal/usecase"
)

type Handler struct {
	usecase usecase.UsecaseContract
}

func NewHandler(usecase usecase.UsecaseContract) HandlerContract {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) Handle(conn net.Conn) error {
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

		var data entity.Data
		var run_time int64
		switch query.Action {
		case constants.ACTION_CREATE:
			start_time := time.Now()
			data, err = h.usecase.Create(query.Key, query.Value)
			if err != nil {
				conn.Write([]byte(err.Error()))
				return err
			}
			run_time = time.Since(start_time).Milliseconds()
		case constants.ACTION_READ:
			start_time := time.Now()
			data, err = h.usecase.Read(query.Key)
			if err != nil {
				conn.Write([]byte(err.Error()))
				return err
			}
			run_time = time.Since(start_time).Milliseconds()
		case constants.ACTION_UPDATE:
			start_time := time.Now()
			data, err = h.usecase.Update(query.Key, query.Value)
			if err != nil {
				conn.Write([]byte(err.Error()))
				return err
			}
			run_time = time.Since(start_time).Milliseconds()
		case constants.ACTION_DELETE:
			start_time := time.Now()
			err = h.usecase.Delete(query.Key)
			if err != nil {
				conn.Write([]byte(err.Error()))
				return err
			}
			run_time = time.Since(start_time).Milliseconds()
		case constants.ACTION_SEARCH:
			start_time := time.Now()
			data, err = h.usecase.Search(query.Value[0], query.Key)
			if err != nil {
				conn.Write([]byte(err.Error()))
				return err
			}
			run_time = time.Since(start_time).Milliseconds()
		}

		response := fmt.Sprintf("%s %s success, took %d ms", query.Action, data.Key, run_time)
		_, err = conn.Write([]byte(response))
		if err != nil {
			conn.Close()
			return err
		}
	}
}
