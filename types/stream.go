package types

import (
	"context"
	"github.com/imnotanderson/X/log"
	"github.com/imnotanderson/X/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Stream struct {
	addr      string
	name      string
	chSend    chan *pb.Request
	chRecv    chan []byte
	die       chan struct{}
	connector pb.Connector_AcceptClient
	kv        map[string]string
}

func NewStream(addr string, name string, kv map[string]string) *Stream {
	s := &Stream{
		addr:   addr,
		name:   name,
		chSend: make(chan *pb.Request, 128),
		chRecv: make(chan []byte, 128),
		kv:     kv,
	}
	return s
}

func (s *Stream) Conn() error {
	conn, err := grpc.Dial(s.addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	c := pb.NewConnectorClient(conn)
	ctx := metadata.NewContext(context.Background(), metadata.New(s.kv))

	connector, err := c.Accept(ctx)
	if err != nil {
		return err
	}
	s.connector = connector
	s.die = make(chan struct{})
	//recv & send msg
	go s.handleRecv()
	go s.handleSend()
	return nil
}

func (s *Stream) handleRecv() {
	log.Infof("service %v recv start", s.name)
	defer log.Infof("service %v recv end", s.name)
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
	log.Infof("service %v send start", s.name)
	defer log.Infof("service %v send end", s.name)
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

func (s *Stream) Send(data []byte, svrId string) {
	s.chSend <- &pb.Request{
		Data:      data,
		ServiceId: svrId,
	}
}

func (s Stream) Recv() <-chan []byte {
	return s.chRecv
}

func (s Stream) Close() {
	s.connector.CloseSend()
}
