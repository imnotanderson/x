package gate

import "github.com/imnotanderson/X/types"

type session struct {
	uuid   string
	stream *types.Stream
	die    chan struct{}
}

//func (s *session) start() {
//	defer s.stream.Close()
//	go func() {
//		for {
//			select {
//			case <-s.die:
//				return
//			case <-s.stream.Conn():
//			}
//		}
//	}()
//}
//
//func (s *session) close() {
//	close(s.die)
//}
