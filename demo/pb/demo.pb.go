// Code generated by protoc-gen-go.
// source: demo.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	demo.proto

It has these top-level messages:
	LoginRequest
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LoginRequest struct {
	Uuid string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*LoginRequest)(nil), "pb.LoginRequest")
}

func init() { proto.RegisterFile("demo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 78 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0xcd, 0xcd,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x52, 0xe2, 0xe2, 0xf1, 0xc9,
	0x4f, 0xcf, 0xcc, 0x0b, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0x29, 0x2d,
	0xcd, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x93, 0xd8, 0xc0, 0xca, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x4e, 0x55, 0xd9, 0x3c, 0x00, 0x00, 0x00,
}
