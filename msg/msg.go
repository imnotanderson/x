package msg

import (
	"github.com/chrislonng/starx/log"
	"github.com/imnotanderson/X/types"
	"sync"
)

type Msg struct {
	dataMsgLock      sync.RWMutex
	dataMsgMap       map[interface{}][]func([]byte)
	playerMsgMapLock sync.RWMutex
	playerMsgMap     map[interface{}][]func(*types.Player, int32, []byte)
}

var m *Msg = &Msg{
	dataMsgMap:   make(map[interface{}][]func([]byte)),
	playerMsgMap: make(map[interface{}][]func(*types.Player, int32, []byte)),
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
	m.dataMsgLock.RLock()
	defer m.dataMsgLock.RUnlock()
	if m.dataMsgMap[sign] != nil {
		log.Errorf("send dataMsg err no listener:%v", sign)
		return
	}
	for _, v := range m.dataMsgMap[sign] {
		v(data)
	}
}

func ListenPlayerMsg(pktType interface{}, f func(pPlayer *types.Player, msgType int32, data []byte)) {
	m.playerMsgMapLock.Lock()
	defer m.playerMsgMapLock.Unlock()
	if m.playerMsgMap[pktType] == nil {
		m.playerMsgMap[pktType] = []func(*types.Player, int32, []byte){f}
		return
	}
	m.playerMsgMap[pktType] = append(m.playerMsgMap[pktType], f)
}

func SendPlayerMsg(pPlayer *types.Player, msgType int32, data []byte) {
	m.playerMsgMapLock.RLock()
	defer m.playerMsgMapLock.RUnlock()
	if m.playerMsgMap[msgType] != nil {
		log.Errorf("send dataMsg err no listener:%v", msgType)
		return
	}
	for _, v := range m.playerMsgMap[msgType] {
		v(pPlayer, msgType, data)
	}
}
