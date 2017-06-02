package agent

import (
	"github.com/imnotanderson/X/log"
	"github.com/imnotanderson/X/pb"
)

type Service struct {
	id     string
	chSend chan []byte
}

func NewService(id string) *Service {
	return &Service{
		id:     id,
		chSend: make(chan []byte, 128),
	}
}

func (s *Service) start_recv(conn pb.Connector_AcceptServer, die <-chan struct{}) <-chan *pb.Request {
	ch := make(chan *pb.Request, 1)
	go func() {
		defer close(ch)
		for {
			req, err := conn.Recv()
			if err != nil {
				log.Infof("recv err:", err)
				return
			}
			select {
			case ch <- req:
			case <-die:
				return
			}
		}
	}()
	return ch
}
