package gate

import "github.com/imnotanderson/X/types"

type session struct {
	uuid   string
	stream *types.Stream
	die    chan struct{}
}
