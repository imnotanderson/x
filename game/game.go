package game

import (
	"errors"
	"github.com/imnotanderson/X/conf"
	"github.com/imnotanderson/X/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
)

var (
	MD_PARSE_ERR         = errors.New("md parse err")
	MD_PARSE_ERR_NO_UUID = errors.New("md parse err no uuid")
)

type Game struct {
	addr string
}

var Module *Game = &Game{
	addr: conf.Game_port,
}

func (g *Game) Init() {

}
func (g *Game) Run(closeSign <-chan struct{}) {
	go g.handleClient()
	<-closeSign
}

func (g *Game) handleClient() {
	lsn, err := net.Listen("tcp", g.addr)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterConnectorServer(s, g)
	err = s.Serve(lsn)
	if err != nil {
		panic(err)
	}
}

func (g *Game) Accept(connector pb.Connector_AcceptServer) error {
	md, ok := metadata.FromContext(connector.Context())
	if ok == false {
		return MD_PARSE_ERR
	}
	if len(md["uuid"]) <= 0 {
		return MD_PARSE_ERR_NO_UUID
	}

	return nil
}
