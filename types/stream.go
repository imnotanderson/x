package types

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/imnotanderson/X/log"
	"github.com/imnotanderson/X/pb"
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

func NewStream(addr string, serviceId uint32) *Stream {
	return &Stream{
		addr:   addr,
		id:     serviceId,
		chSend: make(chan *pb.Request, 128),
		chRecv: make(chan []byte, 128),
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
	ctx := context.Background()

	connector, err := c.Accept(ctx)
	if err != nil {
		log.Errorf("accecp err %v", err)
		return
	}
	//reg id
	regRequest := &pb.ServiceRegRequest{
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
	s.die = make(chan struct{})
	//recv & send msg
	go s.handleRecv()
	go s.handleSend()
	<-s.die
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

func (s *Stream) Send(data []byte, svrId uint32) {
	s.chSend <- &pb.Request{
		Data:      data,
		ServiceId: svrId,
	}
}

func (s Stream) Recv() <-chan []byte {
	return s.chRecv
}
