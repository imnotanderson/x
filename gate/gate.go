//client session manager
package gate

import (
	"github.com/chrislonng/starx/log"
	"github.com/imnotanderson/X/agent"
	"github.com/imnotanderson/X/conf"
	"io"
	"net"
)

type Gate struct {
	addr string
}

var Module *Gate = &Gate{
	addr: conf.Gate_addr,
}

func (g *Gate) Init() {

}

func (g *Gate) Run(closeSign <-chan struct{}) {
	go g.waitForClient()
	<-closeSign
}

func (g *Gate) waitForClient() {
	lsn, err := net.Listen("tcp", g.addr)
	if checkErr(err) {
		return
	}
	for {
		conn, err := lsn.Accept()
		if checkErr(err) {
			return
		}
		go g.handleConn(conn)
	}
}

func (g *Gate) handleConn(conn net.Conn) {
	chRecv := make(chan []byte)
	header := make([]byte, 2)
	go func() {
		defer close(chRecv)
		for {
			size, err := io.ReadFull(conn, header)
			if checkErr(err) {
				return
			}
			payload := make([]byte, size)
			_, err = io.ReadFull(conn, payload)
			if checkErr(err) {
				return
			}
			chRecv <- payload
		}
	}()
	authData := <-chRecv
	if authData == nil {
		return
	}
	if g.auth(authData) == false {
		return
	}
	for {
		data := <-chRecv
		if data == nil {
			return
		}
		//todo:route to gamesvr

	}
}

func (g *Gate) auth(data []byte) bool {
	return true
}

func checkErr(err error) bool {
	if err != nil {
		log.Errorf("err %v", err)
		return true
	}
	return false
}
