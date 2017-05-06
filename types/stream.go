package types

import (
	"../log"
	"../pb"
	"context"
	"google.golang.org/grpc"
)

type Stream struct {
	addr string
	id   uint32
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

	//recv & send msg
}
