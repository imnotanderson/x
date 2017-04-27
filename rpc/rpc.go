package rpc

import "github.com/imnotanderson/X/log"

type Rpc struct {
	fMap map[string]func()
}

func NewRpc() *Rpc {
	return &Rpc{
		fMap: make(map[string]func()),
	}
}

func (r *Rpc) RegFunc(name string, f func()) {
	if r.fMap[name] != nil {
		log.Error("same name ", name)
		return
	}
	r.fMap[name] = f
}

func (r *Rpc) Call(name string) {
	if r.fMap[name] == nil {
		log.Error("no func ", name)
		return
	}
	r.fMap[name]()
}
