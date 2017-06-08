//client session manager
package gate

import (
	"github.com/imnotanderson/X/conf"
	"github.com/imnotanderson/X/log"
	"github.com/imnotanderson/X/packet"
	"github.com/imnotanderson/X/types"
	"net"
	"sync"
	"time"
)

type Gate struct {
	addr           string
	sessionMap     map[string]*session
	sessionMapLock sync.RWMutex
	auth           func(conn net.Conn, data []byte) (uuid string, err error)
}

var Module *Gate = &Gate{
	addr:       conf.Gate_addr,
	auth:       conf.Auth,
	sessionMap: map[string]*session{},
}

func (g *Gate) Init() {

}

func (g *Gate) Run(closeSign <-chan struct{}) {
	go g.waitForClient()
	<-closeSign
}

func (g *Gate) waitForClient() {
	lsn, err := net.Listen("tcp", g.addr)
	if checkErr(err) {
		return
	}
	for {
		conn, err := lsn.Accept()
		if checkErr(err) {
			return
		}
		go g.handleConn(conn)
	}
}

func (g *Gate) handleConn(conn net.Conn) {
	defer conn.Close()
	chRecv := make(chan []byte)
	packet := packet.NewPacket()
	conn_die := make(chan struct{})
	go func() {
		defer close(conn_die)
		for {
			data, err := packet.ReadData(conn)
			if checkErr(err) {
				return
			}
			data = g.decode(data)
			chRecv <- data
		}
	}()

	authData := <-chRecv
	if authData == nil {
		return
	}
	uuid, err := g.auth(conn, authData)
	if err != nil {
		return
	}
	session, err := g.createSession(uuid)
	if err != nil {
		return
	}
	defer func() {
		close(session.die)
	}()

	go func() {
		for {
			select {
			case data := <-session.stream.Recv():
				conn.Write(data)
			case <-session.die:
				return
			}
		}
	}()
	for {
		select {
		case data := <-chRecv:
			session.stream.Send(data, "")
		case <-conn_die:
			return
		}
	}
}

func (g *Gate) createSession(uuid string) (*session, error) {
	kv := map[string]string{
		"uuid": uuid,
	}
	gameSvrAddr := g.selectGameSvr(uuid)
	stream := types.NewStream(gameSvrAddr, "clientsession", kv)
	session := &session{
		uuid:   uuid,
		stream: stream,
		die:    make(chan struct{}),
	}

	g.sessionMapLock.Lock()
	defer g.sessionMapLock.Unlock()
	oldSession := g.sessionMap[uuid]
	if oldSession != nil {
		if oldSession.stream != nil {
			oldSession.stream.Close()
		}
	}
	go func() {
		for {
			select {
			case <-session.stream.Conn():
				<-time.After(time.Second)
			case <-session.die:
				return
			}
		}
	}()
	g.sessionMap[uuid] = session
	log.Debugf("add session %v", uuid)
	go func() {
		<-session.die
		g.sessionMapLock.Lock()
		defer g.sessionMapLock.Unlock()
		if g.sessionMap[session.uuid] == session {
			delete(g.sessionMap, session.uuid)
			log.Debugf("remove session %v", uuid)
		}
		if session.stream != nil {
			session.stream.Close()
		}
	}()
	return session, nil
}

func checkErr(err error) bool {
	if err != nil {
		log.Errorf("err %v", err)
		return true
	}
	return false
}

func (g *Gate) decode(data []byte) []byte {
	return data
}

func (g *Gate) selectGameSvr(uuid string) string {
	return conf.Game_addr
}
