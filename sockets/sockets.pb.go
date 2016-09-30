// Code generated by protoc-gen-go.
// source: sockets.proto
// DO NOT EDIT!

/*
Package sockets is a generated protocol buffer package.

It is generated from these files:
	sockets.proto

It has these top-level messages:
	SocketMessage
	SubscribeMessage
	PublishResponse
*/
package sockets

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

type SocketMessage struct {
	AppId   string `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	Channel string `protobuf:"bytes,2,opt,name=channel" json:"channel,omitempty"`
	Action  string `protobuf:"bytes,3,opt,name=action" json:"action,omitempty"`
	Payload string `protobuf:"bytes,4,opt,name=payload" json:"payload,omitempty"`
}

func (m *SocketMessage) Reset()                    { *m = SocketMessage{} }
func (m *SocketMessage) String() string            { return proto.CompactTextString(m) }
func (*SocketMessage) ProtoMessage()               {}
func (*SocketMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SubscribeMessage struct {
	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	AppId     string `protobuf:"bytes,2,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	Channel   string `protobuf:"bytes,3,opt,name=channel" json:"channel,omitempty"`
}

func (m *SubscribeMessage) Reset()                    { *m = SubscribeMessage{} }
func (m *SubscribeMessage) String() string            { return proto.CompactTextString(m) }
func (*SubscribeMessage) ProtoMessage()               {}
func (*SubscribeMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type PublishResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *PublishResponse) Reset()                    { *m = PublishResponse{} }
func (m *PublishResponse) String() string            { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()               {}
func (*PublishResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*SocketMessage)(nil), "sockets.SocketMessage")
	proto.RegisterType((*SubscribeMessage)(nil), "sockets.SubscribeMessage")
	proto.RegisterType((*PublishResponse)(nil), "sockets.PublishResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Sockets service

type SocketsClient interface {
	Publish(ctx context.Context, in *SocketMessage, opts ...grpc.CallOption) (*PublishResponse, error)
	Subscribe(ctx context.Context, in *SubscribeMessage, opts ...grpc.CallOption) (*PublishResponse, error)
	UnSubscribe(ctx context.Context, in *SubscribeMessage, opts ...grpc.CallOption) (*PublishResponse, error)
}

type socketsClient struct {
	cc *grpc.ClientConn
}

func NewSocketsClient(cc *grpc.ClientConn) SocketsClient {
	return &socketsClient{cc}
}

func (c *socketsClient) Publish(ctx context.Context, in *SocketMessage, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := grpc.Invoke(ctx, "/sockets.Sockets/Publish", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socketsClient) Subscribe(ctx context.Context, in *SubscribeMessage, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := grpc.Invoke(ctx, "/sockets.Sockets/Subscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socketsClient) UnSubscribe(ctx context.Context, in *SubscribeMessage, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := grpc.Invoke(ctx, "/sockets.Sockets/UnSubscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Sockets service

type SocketsServer interface {
	Publish(context.Context, *SocketMessage) (*PublishResponse, error)
	Subscribe(context.Context, *SubscribeMessage) (*PublishResponse, error)
	UnSubscribe(context.Context, *SubscribeMessage) (*PublishResponse, error)
}

func RegisterSocketsServer(s *grpc.Server, srv SocketsServer) {
	s.RegisterService(&_Sockets_serviceDesc, srv)
}

func _Sockets_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SocketMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocketsServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sockets.Sockets/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocketsServer).Publish(ctx, req.(*SocketMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sockets_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocketsServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sockets.Sockets/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocketsServer).Subscribe(ctx, req.(*SubscribeMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sockets_UnSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocketsServer).UnSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sockets.Sockets/UnSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocketsServer).UnSubscribe(ctx, req.(*SubscribeMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sockets_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sockets.Sockets",
	HandlerType: (*SocketsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Sockets_Publish_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _Sockets_Subscribe_Handler,
		},
		{
			MethodName: "UnSubscribe",
			Handler:    _Sockets_UnSubscribe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("sockets.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x91, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0x96, 0x36, 0xe4, 0x50, 0x05, 0xb2, 0xa0, 0x32, 0x48, 0x48, 0x28, 0x13, 0x53,
	0x07, 0x18, 0x59, 0x80, 0x8d, 0x01, 0x09, 0xa5, 0x62, 0x46, 0xb6, 0x73, 0xa2, 0x11, 0x91, 0x6d,
	0xf9, 0x92, 0x81, 0xff, 0xc8, 0x8f, 0xc2, 0x75, 0xed, 0x46, 0xad, 0x04, 0x4b, 0xb7, 0xfb, 0x2e,
	0x79, 0xef, 0x5d, 0x5e, 0x60, 0x46, 0x46, 0x7d, 0x61, 0x47, 0x0b, 0xeb, 0x4c, 0x67, 0x58, 0x1e,
	0xb1, 0x74, 0x30, 0x5b, 0x86, 0xf1, 0x15, 0x89, 0xc4, 0x27, 0xb2, 0x0b, 0x98, 0x0a, 0x6b, 0x3f,
	0x9a, 0x9a, 0x67, 0x37, 0xd9, 0x6d, 0x51, 0x4d, 0x3c, 0xbd, 0xd4, 0x8c, 0x43, 0xae, 0x56, 0x42,
	0x6b, 0x6c, 0xf9, 0x28, 0xec, 0x13, 0xb2, 0xb9, 0x17, 0xa8, 0xae, 0x31, 0x9a, 0x8f, 0xc3, 0x83,
	0x48, 0x6b, 0x85, 0x15, 0xdf, 0xad, 0x11, 0x35, 0x3f, 0xda, 0x28, 0x22, 0x96, 0x12, 0xce, 0x96,
	0xbd, 0x24, 0xe5, 0x1a, 0x89, 0x29, 0xf6, 0x1a, 0x80, 0xfc, 0xe8, 0x85, 0x43, 0x74, 0x11, 0x37,
	0x3e, 0x7e, 0xb8, 0x6a, 0xf4, 0xc7, 0x55, 0xe3, 0x9d, 0xab, 0xca, 0x27, 0x38, 0x7d, 0xeb, 0x65,
	0xdb, 0xd0, 0xaa, 0x42, 0xb2, 0x46, 0x13, 0xae, 0x5f, 0xa6, 0x5e, 0x29, 0xef, 0x19, 0xfc, 0x8f,
	0xab, 0x84, 0xec, 0x1c, 0x26, 0xe8, 0x9c, 0x71, 0xc9, 0x3c, 0xc0, 0xdd, 0x4f, 0x06, 0xf9, 0xa6,
	0x1b, 0x62, 0x0f, 0x90, 0x47, 0x3b, 0x36, 0x5f, 0xa4, 0x2a, 0x77, 0x8a, 0xbb, 0xe2, 0xdb, 0xfd,
	0x7e, 0xf0, 0x23, 0x14, 0xdb, 0xef, 0x65, 0x97, 0x83, 0x7c, 0xaf, 0x83, 0x7f, 0x1c, 0x9e, 0xe1,
	0xe4, 0x5d, 0x1f, 0xe6, 0x21, 0xa7, 0xe1, 0xcf, 0xdf, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xc6,
	0x6b, 0xa1, 0x47, 0x0a, 0x02, 0x00, 0x00,
}
