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
		log.Errorf("same name %v", name)
		return
	}
	r.fMap[name] = f
}

func (r *Rpc) Call(name string) {
	if r.fMap[name] == nil {
		log.Errorf("no func %v", name)
		return
	}
	r.fMap[name]()
}
