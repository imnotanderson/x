package types

import (
	"errors"
	"github.com/imnotanderson/X/log"
	"github.com/imnotanderson/X/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
	"proto"
)

type StreamListener struct {
	addr string
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
	request, err := conn.Recv()
	md, ok := metadata.FromContext(conn.Context())
	if ok == false {
		return nil
	}
	uuid := md["uuid"][0]
	if checkErr(err) {
		return err
	}

	//regRequest := &pb.ServiceRegRequest{}
	//err = proto.Unmarshal(request.Data, regRequest)
	//if checkErr(err) {
	//	return err
	//}
	//service_die := make(chan struct{})
	//defer close(service_die)
	//service := s.regService(regRequest.ServiceId)
	//defer s.removeService(service.id)
	//ch_recv := service.start_recv(conn, service_die)
	//
	//for {
	//	select {
	//	case req, ok := <-ch_recv:
	//		if ok == false {
	//			return errors.New("ch_recv err")
	//		}
	//		targetService := (*Service)(nil)
	//		s.mapLock.Lock()
	//		targetService = s.serviceMap[req.ServiceId]
	//		s.mapLock.Unlock()
	//		if targetService != nil {
	//			targetService.chSend <- req.Data
	//		} else {
	//			log.Errorf("no service found id:%v", req.ServiceId)
	//		}
	//	case sendData := <-service.chSend:
	//		err := conn.Send(&pb.Reply{
	//			Data: sendData,
	//		})
	//		if err != nil {
	//			return err
	//		}
	//	}
	//}
	return nil
}

func checkErr(err error) bool {
	if err != nil {
		log.Errorf("agent err:%v", err)
		return true
	}
	return false
}
