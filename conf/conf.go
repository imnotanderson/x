package conf

import (
	"net"
)

const (
	Game_port  = ":7777"
	Game_addr  = "127.0.0.1:7777"
	Gate_addr  = ":6666"
	Agent_addr = ":9999"
)

var (
	Auth func(conn net.Conn, data []byte) (uuid string, err error) = func(conn net.Conn, data []byte) (uuid string, err error) {
		uuid = "1"
		err = nil
		return
	}
)
