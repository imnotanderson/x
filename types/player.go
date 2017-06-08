package types

import (
	"github.com/imnotanderson/X/pb"
)

type Player struct {
	Uuid   string
	Die    chan struct{}
	ChSend chan []*pb.Msg
	ChRecv chan []*pb.Msg
}
