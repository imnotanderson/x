package types

import (
	"../log"
	"../pb"
	"context"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

type Stream struct {
	addr      string
	id        uint32
	chSend    chan *pb.Request
	chRecv    chan []byte
	die       chan struct{}
	connector pb.Connector_AcceptClient
}

func NewStream(addr string, serviceId uint32) {
	return &Stream{
		addr:   addr,
		id:     serviceId,
		chSend: make(chan *pb.Request, 128),
		chRecv: make(chan []byte, 128),
		die:    make(chan struct{}),
	}
}

func (s *Stream) Conn() {
	conn, err := grpc.Dial(s.addr, grpc.WithInsecure())
	if err != nil {
		log.Errorf("dial %v err: %v", s.addr, err)
		return
	}
	defer conn.Close()
	c := pb.NewConnectorClient(conn)
	connector, err := c.Accept(context.Background())
	if err != nil {
		log.Errorf("accecp err %v", err)
		return
	}
	//reg id
	regRequest := pb.ServiceRegRequest{
		ServiceId: s.id,
	}
	data, err := proto.Marshal(regRequest)
	if err != nil {
		log.Errorf("marsha regrequest err %v", err)
		return
	}
	request := &pb.Request{
		Data:      data,
		ServiceId: 0,
	}

	err = connector.Send(request)
	if err != nil {
		log.Errorf("reg svr err %v", err)
	}
	s.connector = connector
	//recv & send msg
	go s.handleRecv()
	go s.handleSend()
}

func (s *Stream) handleRecv() {
	log.Infof("service %v recv start", s.id)
	defer log.Infof("service %v recv end", s.id)
	for {
		reply, err := s.connector.Recv()
		if err != nil {
			close(s.die)
			return
		}
		s.chRecv <- reply.Data
	}
}

func (s *Stream) handleSend() {
	log.Infof("service %v send start", s.id)
	defer log.Infof("service %v send end", s.id)
	for {
		select {
		case req := <-s.chSend:
			err := s.connector.Send(req)
			if err != nil {
				return
			}
		case <-s.die:
			return
		}

	}
}
