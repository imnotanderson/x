package main

import (
	"github.com/imnotanderson/X/types"
)

func main() {
	C(1, 2, []byte("from c1"))
	C(2, 1, []byte("from c2"))
	select {}
}

func C(id uint32, toId uint32, data []byte) {
	s := types.NewStream("127.0.0.1:9999", id)
	go func() {
		for {
			s.Conn()
			println("reconn")
		}
	}()
	go func() {
		data := <-s.Recv()
		println(id, "======>recv", string(data))
	}()
	s.Send(data, toId)
}
