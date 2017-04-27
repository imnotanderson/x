package connector

//import (
//	"github.com/imnotanderson/X/connector/pb"
//	"google.golang.org/grpc"
//	"net"
//)
//
//type IConnector interface {
//	ReadFull([]byte) error
//	Write([]byte) error
//}
//
//type gconnector struct {
//	recvData chan []byte
//	sendData chan []byte
//}
//
//func (g *gconnector) Accept(c pb.Connector_AcceptServer) error {
//	return nil
//}
//
//func (g *gconnector) ReadFull(buf []byte) error {
//	return nil
//}
//
//func (g *gconnector) Write(data []byte) error {
//	return nil
//}
//
//func (c *gconnector) Listen(laddr string) error {
//	lsn, err := net.Listen("tcp", laddr)
//	if err != nil {
//		return err
//	}
//	s := grpc.NewServer()
//	pb.RegisterConnectorServer(s, &gconnector{})
//	err = s.Serve(lsn)
//	if err != nil {
//		return err
//	}
//	return nil
//}
