package msg

import (
	"github.com/chrislonng/starx/log"
	"sync"
)

type Msg struct {
	dataMsgLock sync.Mutex
	dataMsgMap  map[interface{}][]func([]byte)
}

var m *Msg = &Msg{
	dataMsgMap: make(map[interface{}][]func([]byte)),
}

func ListenDataMsg(sign interface{}, f func([]byte)) {
	m.dataMsgLock.Lock()
	defer m.dataMsgLock.Unlock()
	if m.dataMsgMap[sign] == nil {
		m.dataMsgMap[sign] = []func([]byte){f}
		return
	}
	m.dataMsgMap[sign] = append(m.dataMsgMap[sign], f)
}

func SendDataMsg(sign interface{}, data []byte) {
	m.dataMsgLock.Lock()
	defer m.dataMsgLock.Unlock()
	if m.dataMsgMap[sign] != nil {
		log.Errorf("send dataMsg err no listener:%v", sign)
		return
	}
	for _, v := range m.dataMsgMap[sign] {
		v(data)
	}
}
