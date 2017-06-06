package packet

import (
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"io"
)

type Packet struct {
	header []byte
}

func NewPacket() *Packet {
	return &Packet{
		header: make([]byte, 2),
	}
}

func (p *Packet) Read(r io.Reader) ([]byte, error) {
	_, err := io.ReadFull(r, p.header)
	if err != nil {
		return nil, err
	}
	size := binary.BigEndian.Uint16(p.header)
	data := make([]byte, size)
	_, err = io.ReadFull(r, data)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (p *Packet) Write(msg proto.Message) []byte {
	data, err := proto.Marshal(msg)
	if err != nil {
		panic(err)
	}
	data = append([]byte{0, 0}, data...)
	binary.BigEndian.PutUint16(data, uint16(len(data)-2))
	return data
}
