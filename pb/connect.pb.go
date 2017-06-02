// Code generated by protoc-gen-go.
// source: connect.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	connect.proto

It has these top-level messages:
	Request
	Reply
	ServiceRegRequest
	Frame
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	// send to service id
	ServiceId uint32 `protobuf:"varint,1,opt,name=serviceId" json:"serviceId,omitempty"`
	Data      []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Reply struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ServiceRegRequest struct {
	ServiceId uint32 `protobuf:"varint,1,opt,name=serviceId" json:"serviceId,omitempty"`
}

func (m *ServiceRegRequest) Reset()                    { *m = ServiceRegRequest{} }
func (m *ServiceRegRequest) String() string            { return proto.CompactTextString(m) }
func (*ServiceRegRequest) ProtoMessage()               {}
func (*ServiceRegRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Frame struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Frame) Reset()                    { *m = Frame{} }
func (m *Frame) String() string            { return proto.CompactTextString(m) }
func (*Frame) ProtoMessage()               {}
func (*Frame) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Request)(nil), "pb.Request")
	proto.RegisterType((*Reply)(nil), "pb.Reply")
	proto.RegisterType((*ServiceRegRequest)(nil), "pb.ServiceRegRequest")
	proto.RegisterType((*Frame)(nil), "pb.Frame")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Connector service

type ConnectorClient interface {
	Accept(ctx context.Context, opts ...grpc.CallOption) (Connector_AcceptClient, error)
}

type connectorClient struct {
	cc *grpc.ClientConn
}

func NewConnectorClient(cc *grpc.ClientConn) ConnectorClient {
	return &connectorClient{cc}
}

func (c *connectorClient) Accept(ctx context.Context, opts ...grpc.CallOption) (Connector_AcceptClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Connector_serviceDesc.Streams[0], c.cc, "/pb.Connector/Accept", opts...)
	if err != nil {
		return nil, err
	}
	x := &connectorAcceptClient{stream}
	return x, nil
}

type Connector_AcceptClient interface {
	Send(*Request) error
	Recv() (*Reply, error)
	grpc.ClientStream
}

type connectorAcceptClient struct {
	grpc.ClientStream
}

func (x *connectorAcceptClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *connectorAcceptClient) Recv() (*Reply, error) {
	m := new(Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Connector service

type ConnectorServer interface {
	Accept(Connector_AcceptServer) error
}

func RegisterConnectorServer(s *grpc.Server, srv ConnectorServer) {
	s.RegisterService(&_Connector_serviceDesc, srv)
}

func _Connector_Accept_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConnectorServer).Accept(&connectorAcceptServer{stream})
}

type Connector_AcceptServer interface {
	Send(*Reply) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type connectorAcceptServer struct {
	grpc.ServerStream
}

func (x *connectorAcceptServer) Send(m *Reply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *connectorAcceptServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Connector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Connector",
	HandlerType: (*ConnectorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Accept",
			Handler:       _Connector_Accept_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

// Client API for ClientConnector service

type ClientConnectorClient interface {
	Accept(ctx context.Context, opts ...grpc.CallOption) (ClientConnector_AcceptClient, error)
}

type clientConnectorClient struct {
	cc *grpc.ClientConn
}

func NewClientConnectorClient(cc *grpc.ClientConn) ClientConnectorClient {
	return &clientConnectorClient{cc}
}

func (c *clientConnectorClient) Accept(ctx context.Context, opts ...grpc.CallOption) (ClientConnector_AcceptClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ClientConnector_serviceDesc.Streams[0], c.cc, "/pb.ClientConnector/Accept", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientConnectorAcceptClient{stream}
	return x, nil
}

type ClientConnector_AcceptClient interface {
	Send(*Frame) error
	Recv() (*Frame, error)
	grpc.ClientStream
}

type clientConnectorAcceptClient struct {
	grpc.ClientStream
}

func (x *clientConnectorAcceptClient) Send(m *Frame) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientConnectorAcceptClient) Recv() (*Frame, error) {
	m := new(Frame)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ClientConnector service

type ClientConnectorServer interface {
	Accept(ClientConnector_AcceptServer) error
}

func RegisterClientConnectorServer(s *grpc.Server, srv ClientConnectorServer) {
	s.RegisterService(&_ClientConnector_serviceDesc, srv)
}

func _ClientConnector_Accept_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientConnectorServer).Accept(&clientConnectorAcceptServer{stream})
}

type ClientConnector_AcceptServer interface {
	Send(*Frame) error
	Recv() (*Frame, error)
	grpc.ServerStream
}

type clientConnectorAcceptServer struct {
	grpc.ServerStream
}

func (x *clientConnectorAcceptServer) Send(m *Frame) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientConnectorAcceptServer) Recv() (*Frame, error) {
	m := new(Frame)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ClientConnector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ClientConnector",
	HandlerType: (*ClientConnectorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Accept",
			Handler:       _ClientConnector_Accept_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("connect.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x90, 0xbf, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x71, 0x05, 0x45, 0x39, 0xa8, 0x2a, 0x3c, 0x55, 0x2d, 0x43, 0x15, 0x21, 0x94, 0xc9,
	0x2a, 0xed, 0xc0, 0xc0, 0x44, 0x2b, 0x21, 0xd8, 0x2a, 0x33, 0x30, 0x3b, 0xce, 0xa9, 0x44, 0x72,
	0xe3, 0xc3, 0x31, 0x3f, 0xfa, 0xdf, 0x23, 0xdb, 0x28, 0x61, 0x60, 0x60, 0xfb, 0xec, 0xf3, 0xf7,
	0xce, 0x7a, 0x30, 0xd2, 0xb6, 0x69, 0x50, 0x7b, 0x41, 0xce, 0x7a, 0xcb, 0x07, 0x54, 0xe6, 0x77,
	0x70, 0x2a, 0xf1, 0xed, 0x1d, 0x5b, 0xcf, 0x2f, 0x21, 0x6b, 0xd1, 0x7d, 0xd4, 0x1a, 0x9f, 0xaa,
	0x09, 0x9b, 0xb3, 0x62, 0x24, 0xfb, 0x0b, 0xce, 0xe1, 0xb8, 0x52, 0x5e, 0x4d, 0x06, 0x73, 0x56,
	0x9c, 0xcb, 0xc8, 0xf9, 0x0c, 0x4e, 0x24, 0x92, 0x39, 0x74, 0x43, 0xf6, 0x6b, 0x78, 0x03, 0x17,
	0xcf, 0xc9, 0x96, 0xb8, 0xfb, 0xd7, 0x8e, 0x90, 0xf7, 0xe0, 0xd4, 0x1e, 0xff, 0xca, 0x5b, 0xae,
	0x20, 0xdb, 0xa4, 0xef, 0x5b, 0xc7, 0xaf, 0x61, 0x78, 0xaf, 0x35, 0x92, 0xe7, 0x67, 0x82, 0x4a,
	0xf1, 0x13, 0x3f, 0xcd, 0xd2, 0x81, 0xcc, 0x21, 0x3f, 0x2a, 0xd8, 0x82, 0x2d, 0x6f, 0x61, 0xbc,
	0x31, 0x35, 0x36, 0xbe, 0x57, 0xaf, 0x3a, 0x35, 0xbe, 0x8e, 0x0b, 0xa7, 0x3d, 0x26, 0x71, 0xbd,
	0x80, 0x59, 0x6d, 0xc5, 0xce, 0x91, 0x16, 0xf8, 0xa5, 0xf6, 0x64, 0xb0, 0x15, 0xaf, 0x68, 0x8c,
	0xfd, 0xb4, 0xce, 0x54, 0xeb, 0xf1, 0x63, 0xe0, 0x97, 0xc0, 0xdb, 0xd0, 0xe5, 0x96, 0x95, 0xc3,
	0x58, 0xea, 0xea, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xa4, 0x6b, 0xee, 0x65, 0x01, 0x00, 0x00,
}
