package agent

import (
	"github.com/imnotanderson/X/log"
	"github.com/imnotanderson/X/pb"
)

type service struct {
	id     uint32
	chSend chan []byte
}

func NewService(id uint32) *service {
	return &service{
		id:     id,
		chSend: make(chan []byte, 128),
	}
}

func (s *service) start_recv(conn pb.Connector_AcceptServer, die <-chan struct{}) (ch_recv <-chan *pb.Request) {
	ch_recv = make(chan []byte, 1)
	defer close(ch_recv)
	for {
		req, err := conn.Recv()
		if err != nil {
			log.Infof("recv err:", err)
			return
		}
		select {
		case ch_recv <- req:
		case <-die:
			return
		}
	}
}
