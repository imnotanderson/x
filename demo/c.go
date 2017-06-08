package main

import (
	"github.com/imnotanderson/X/types"
	"time"
)

func main() {
	C("1", "2", []byte("from c1"))
	C("2", "1", []byte("from c2"))
	select {}
}

func C(id string, toId string, data []byte) {
	kv := map[string]string{
		"id": id,
	}
	s := types.NewStream("127.0.0.1:9999", id, kv)
	go func() {
		for {
			<-s.Conn()
			println("reconn")
		}
	}()
	go func() {
		data := <-s.Recv()
		println(id, "======>recv", string(data))
	}()
	<-time.After(time.Second)
	s.Send(data, toId)

}
