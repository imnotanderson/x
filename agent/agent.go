//relay between services
package agent

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/imnotanderson/X/conf"
	"github.com/imnotanderson/X/log"
	"github.com/imnotanderson/X/pb"
	"google.golang.org/grpc"
	"net"
	"sync"
)

type Agent struct {
	addr       string
	mapLock    *sync.RWMutex
	serviceMap map[uint32]*Service
}

var Module *Agent = &Agent{
	addr:       conf.Agent_addr,
	serviceMap: make(map[uint32]*Service),
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
	request, err := conn.Recv()
	if checkErr(err) {
		return err
	}
	regRequest := &pb.ServiceRegRequest{}
	err = proto.Unmarshal(request.Data, regRequest)
	if checkErr(err) {
		return err
	}
	service_die := make(chan struct{})
	defer close(service_die)
	service := a.regService(regRequest.ServiceId)
	defer a.removeService(service.id)
	ch_recv := service.start_recv(conn, service_die)

	for {
		select {
		case req, ok := <-ch_recv:
			if ok == false {
				return errors.New("ch_recv err")
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

func (a *Agent) regService(serviceId uint32) *Service {
	a.mapLock.Lock()
	defer a.mapLock.Unlock()
	if nil != a.serviceMap[serviceId] {
		log.Errorf("reg service err: exist same id %v", serviceId)
		return nil
	}
	service := NewService(serviceId)
	a.serviceMap[serviceId] = service
	log.Infof("service %v reg", serviceId)
	return service
}

func (a *Agent) removeService(serviceId uint32) {
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
