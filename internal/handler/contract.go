package handler

import "net"

type HandlerContract interface {
	Handle(conn net.Conn) error
}
