package main

import (
	"github.com/imnotanderson/X/types"
)

func main() {
	s := types.NewStream("127.0.0.1:9999", 1)
	for {
		s.Conn()
		println("reconn")
	}
}
