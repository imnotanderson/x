//relay between services
package agent

import (
	"errors"
	"github.com/imnotanderson/X/conf"
	"github.com/imnotanderson/X/log"
	"github.com/imnotanderson/X/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
	"sync"
)

var (
	REG_SERVOCE_ERR_SAME_ID = errors.New("reg service err: exist same id")
	CH_RECV_ERR             = errors.New("ch_recv err")
)

type Agent struct {
	addr       string
	mapLock    *sync.RWMutex
	serviceMap map[string]*Service
}

var Module *Agent = &Agent{
	addr:       conf.Agent_addr,
	serviceMap: make(map[string]*Service),
	mapLock:    new(sync.RWMutex),
}

func (a *Agent) Init() {

}

func (a *Agent) Run(closeSign <-chan struct{}) {
	go a.handleService()
	<-closeSign
	a.handleDestroy()
}

func (a *Agent) handleService() {
	lsn, err := net.Listen("tcp", a.addr)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterConnectorServer(s, a)
	err = s.Serve(lsn)
	if err != nil {
		panic(err)
	}
}

func (a *Agent) handleDestroy() {

}

func (a *Agent) Accept(conn pb.Connector_AcceptServer) error {

	md, ok := metadata.FromContext(conn.Context())
	if ok == false {
		log.Errorf("ok == false")
		return nil
	}
	if len(md["id"]) == 0 {
		log.Errorf("no id %+v", md)
		return nil
	}
	serviceId := md["id"][0]

	service_die := make(chan struct{})
	defer close(service_die)
	service, err := a.regService(serviceId)
	if err != nil {
		return err
	}
	defer a.removeService(service.id)
	ch_recv := service.start_recv(conn, service_die)

	for {
		select {
		case req, ok := <-ch_recv:
			if ok == false {
				return CH_RECV_ERR
			}
			targetService := (*Service)(nil)
			a.mapLock.Lock()
			targetService = a.serviceMap[req.ServiceId]
			a.mapLock.Unlock()
			if targetService != nil {
				targetService.chSend <- req.Data
			} else {
				log.Errorf("no service found id:%v", req.ServiceId)
			}
		case sendData := <-service.chSend:
			err := conn.Send(&pb.Reply{
				Data: sendData,
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *Agent) regService(serviceId string) (*Service, error) {
	a.mapLock.Lock()
	defer a.mapLock.Unlock()
	if nil != a.serviceMap[serviceId] {
		log.Errorf("reg service err: exist same id %v", serviceId)
		return nil, REG_SERVOCE_ERR_SAME_ID
	}
	service := NewService(serviceId)
	a.serviceMap[serviceId] = service
	log.Infof("service %v reg", serviceId)
	return service, nil
}

func (a *Agent) removeService(serviceId string) {
	a.mapLock.Lock()
	defer a.mapLock.Unlock()
	delete(a.serviceMap, serviceId)
	log.Infof("service %v remove", serviceId)
}

func checkErr(err error) bool {
	if err != nil {
		log.Errorf("agent err:%v", err)
		return true
	}
	return false
}
