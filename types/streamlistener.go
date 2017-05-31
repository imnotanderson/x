package types

import (
	"errors"
	"github.com/chrislonng/starx/service"
	"github.com/golang/protobuf/proto"
	"github.com/imnotanderson/X/log"
	"github.com/imnotanderson/X/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
)

type StreamListener struct {
	addr   string
	chRecv chan []byte
	chSend chan []byte
}

func (s *StreamListener) Listen() {
	lsn, err := net.Listen("tcp", s.addr)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterConnectorServer(s, s)
	err = s.Serve(lsn)
	if err != nil {
		panic(err)
	}
}

func (s *StreamListener) Accept(conn pb.Connector_AcceptServer) error {
	//regRequest := &pb.ServiceRegRequest{}
	//err = proto.Unmarshal(request.Data, regRequest)
	//if checkErr(err) {
	//	return err
	//}
	service_die := make(chan struct{})
	defer close(service_die)
	//service := s.regService(regRequest.ServiceId)
	//defer s.removeService(service.id)
	ch_recv := s.start_recv(conn, service_die)
	//
	for {
		select {
		case req, ok := <-ch_recv:
			if ok == false {
				return errors.New("ch_recv err")
			}
			//targetService := (*Service)(nil)
			//s.mapLock.Lock()
			//targetService = s.serviceMap[req.ServiceId]
			//s.mapLock.Unlock()
			//if targetService != nil {
			s.chSend <- req.Data
			//} else {
			//	log.Errorf("no service found id:%v", req.ServiceId)
			//}
		case sendData := <-s.chSend:
			err := conn.Send(&pb.Reply{
				Data: sendData,
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *StreamListener) start_recv(conn pb.Connector_AcceptServer, die <-chan struct{}) <-chan *pb.Request {
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

func checkErr(err error) bool {
	if err != nil {
		log.Errorf("agent err:%v", err)
		return true
	}
	return false
}
