package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/imnotanderson/X/agent"
	"github.com/imnotanderson/X/demo/pb"
	"github.com/imnotanderson/X/gate"
	"github.com/imnotanderson/X/launcher"
	"net"
)

func init() {
	gate.Module.Auth = func(conn net.Conn, data []byte) (uuid string, err error) {
		pkt := &pb.LoginRequest{}
		err = proto.Unmarshal(data, pkt)
		if err != nil {
			fmt.Printf("err:%v", err)
			return "", err
		}
		uuid = pkt.Uuid
		return
	}
}

func main() {
	l := launcher.Launcher{}
	l.Start(
		agent.Module,
		gate.Module,
	)
}
