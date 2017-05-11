package main

import (
	"github.com/imnotanderson/X/types"
	"time"
)

func main() {
	types.NewStream("127.0.0.1:9999", 1).Conn()
	<-time.After(time.Hour)
}
