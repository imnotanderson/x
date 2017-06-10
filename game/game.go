package game

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/imnotanderson/X/conf"
	"github.com/imnotanderson/X/log"
	. "github.com/imnotanderson/X/msg"
	"github.com/imnotanderson/X/pb"
	. "github.com/imnotanderson/X/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
	"sync"
)

var (
	MD_PARSE_ERR         = errors.New("md parse err")
	MD_PARSE_ERR_NO_UUID = errors.New("md parse err no uuid")
)

type Game struct {
	addr          string
	playerMap     map[string]*Player
	playerMapLock sync.RWMutex
	stream        *Stream
}

var Module *Game = &Game{
	addr:      conf.Game_port,
	playerMap: map[string]*Player{},
}

func (g *Game) Init() {

}

func (g *Game) Run(closeSign <-chan struct{}) {
	go g.handleClient()

	//handle stream
	kv := map[string]string{"id": "game"}
	g.stream = NewStream(conf.Agent_addr, "game", kv)
	go g.handleRecvStream()
	select {
	case <-g.stream.Conn():
	case <-closeSign:
		return
	}
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
	uuid := md["uuid"][0]

	player := g.addPlayer(uuid)
	log.Debugf("player %v join", uuid)
	defer close(player.Die)
	for {
		req, err := connector.Recv()
		if err != nil {
			return err
		}
		//todo:msg pool
		msg := &pb.Msg{}
		err = proto.Unmarshal(req.Data, msg)
		if err != nil {
			return err
		}
		SendPlayerMsg(player, msg.MsgType, msg.MsgData)
	}
	return nil
}

func (g *Game) addPlayer(uuid string) *Player {
	g.playerMapLock.Lock()
	defer g.playerMapLock.Unlock()
	oldPlayer := g.playerMap[uuid]
	if oldPlayer != nil {
		panic(fmt.Sprintf("same uuid %v", uuid))
	}
	player := &Player{
		Uuid:   uuid,
		Die:    make(chan struct{}),
		ChSend: make(chan []*pb.Msg, 128),
		ChRecv: make(chan []*pb.Msg, 128),
	}
	g.playerMap[uuid] = player
	return player
}

func (g *Game) handleRecvStream() {
	go func() {
		for {
			data, ok := <-g.stream.Recv()
			if ok == false {
				return
			}
			msg := &pb.Msg{}
			err := proto.Unmarshal(data, msg)
			if err != nil {
				return
			}
			SendDataMsg(msg.MsgType, msg.MsgData)
		}
	}()
}

func (g *Game) SendMsgToService(svrId string, msgType int32, msgData []byte) {
	msg := &pb.Msg{MsgType: msgType, MsgData: msgData}
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Errorf("marshal msg err %v", err)
		return
	}
	g.stream.Send(data, svrId)
}
