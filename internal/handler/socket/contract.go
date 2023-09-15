package socket

import (
	"net"
	"vektor/internal/common"
)

type HandlerContract interface {
	Handle(conn net.Conn, query common.Query) error
}
