package main

import (
	"fmt"
	"github.com/imnotanderson/X/conf"
	"github.com/imnotanderson/X/demo/pb"
	"github.com/imnotanderson/X/packet"
	"net"
)

var cSend chan []byte = make(chan []byte)
var cRecv chan []byte = make(chan []byte)

func main() {
	conn, err := net.Dial("tcp", conf.Gate_addr)
	checkErr(err)
	go func() {
		for {
			_, err := conn.Write(<-cSend)
			checkErr(err)
		}
	}()
	go func() {
		p := packet.NewPacket()
		for {
			data, err := p.ReadData(conn)
			checkErr(err)
			cRecv <- data
		}
	}()
	cSend <- packet.NewPacket().Write(&pb.LoginRequest{"uuuuuid"})
	for {
		fmt.Println(<-cRecv)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
