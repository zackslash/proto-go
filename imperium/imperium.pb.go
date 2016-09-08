// Code generated by protoc-gen-go.
// source: imperium.proto
// DO NOT EDIT!

/*
Package imperium is a generated protocol buffer package.

It is generated from these files:
	imperium.proto

It has these top-level messages:
	CertificateRequest
	CertificateResponse
*/
package imperium

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

type CertificateRequest struct {
	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
}

func (m *CertificateRequest) Reset()                    { *m = CertificateRequest{} }
func (m *CertificateRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateRequest) ProtoMessage()               {}
func (*CertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CertificateResponse struct {
	Certificate     string                     `protobuf:"bytes,1,opt,name=certificate" json:"certificate,omitempty"`
	PrivateKey      string                     `protobuf:"bytes,2,opt,name=private_key,json=privateKey" json:"private_key,omitempty"`
	RootCertificate string                     `protobuf:"bytes,3,opt,name=root_certificate,json=rootCertificate" json:"root_certificate,omitempty"`
	Hostname        string                     `protobuf:"bytes,4,opt,name=hostname" json:"hostname,omitempty"`
	ExpiryTime      *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=expiry_time,json=expiryTime" json:"expiry_time,omitempty"`
}

func (m *CertificateResponse) Reset()                    { *m = CertificateResponse{} }
func (m *CertificateResponse) String() string            { return proto.CompactTextString(m) }
func (*CertificateResponse) ProtoMessage()               {}
func (*CertificateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CertificateResponse) GetExpiryTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpiryTime
	}
	return nil
}

func init() {
	proto.RegisterType((*CertificateRequest)(nil), "imperium.CertificateRequest")
	proto.RegisterType((*CertificateResponse)(nil), "imperium.CertificateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Imperium service

type ImperiumClient interface {
	Request(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
}

type imperiumClient struct {
	cc *grpc.ClientConn
}

func NewImperiumClient(cc *grpc.ClientConn) ImperiumClient {
	return &imperiumClient{cc}
}

func (c *imperiumClient) Request(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := grpc.Invoke(ctx, "/imperium.Imperium/Request", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Imperium service

type ImperiumServer interface {
	Request(context.Context, *CertificateRequest) (*CertificateResponse, error)
}

func RegisterImperiumServer(s *grpc.Server, srv ImperiumServer) {
	s.RegisterService(&_Imperium_serviceDesc, srv)
}

func _Imperium_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImperiumServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imperium.Imperium/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImperiumServer).Request(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Imperium_serviceDesc = grpc.ServiceDesc{
	ServiceName: "imperium.Imperium",
	HandlerType: (*ImperiumServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _Imperium_Request_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("imperium.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xa9, 0xba, 0x59, 0x5f, 0x41, 0xe5, 0x89, 0x50, 0x8a, 0xb2, 0xb1, 0x93, 0x22, 0x64,
	0x30, 0x8f, 0x1e, 0x05, 0x61, 0x78, 0x2b, 0xde, 0x4b, 0xb7, 0xbd, 0xcd, 0xa0, 0x59, 0x9e, 0x69,
	0x2a, 0xf6, 0xbb, 0xfa, 0x61, 0x6c, 0xd3, 0x66, 0x56, 0x64, 0xb7, 0xbc, 0xff, 0xff, 0x97, 0x3c,
	0xf2, 0x83, 0x53, 0xa9, 0x98, 0x8c, 0x2c, 0x95, 0x60, 0xa3, 0xad, 0xc6, 0xd0, 0xcf, 0xc9, 0x68,
	0xa3, 0xf5, 0xe6, 0x9d, 0xa6, 0x2e, 0x5f, 0x94, 0xeb, 0xa9, 0x95, 0x8a, 0x0a, 0x9b, 0x2b, 0x6e,
	0xd1, 0xc9, 0x1d, 0xe0, 0x23, 0x19, 0x2b, 0xd7, 0x72, 0x99, 0x5b, 0x4a, 0xe9, 0xa3, 0xac, 0x6b,
	0xbc, 0x84, 0x61, 0xce, 0x9c, 0xc9, 0x55, 0x1c, 0x8c, 0x83, 0x9b, 0x93, 0x74, 0x50, 0x4f, 0xf3,
	0xd5, 0xe4, 0x3b, 0x80, 0x8b, 0x3f, 0x74, 0xc1, 0x7a, 0x5b, 0x10, 0x8e, 0x21, 0x5a, 0xfe, 0xc6,
	0xdd, 0x9d, 0x7e, 0x84, 0x23, 0x88, 0xd8, 0xc8, 0xcf, 0xfa, 0x98, 0xbd, 0x51, 0x15, 0x1f, 0x38,
	0x02, 0xba, 0xe8, 0x99, 0x2a, 0xbc, 0x85, 0x73, 0xa3, 0xb5, 0xcd, 0xfa, 0xef, 0x1c, 0x3a, 0xea,
	0xac, 0xc9, 0x7b, 0x5b, 0x31, 0x81, 0xf0, 0x55, 0x17, 0x76, 0x9b, 0x2b, 0x8a, 0x8f, 0x1c, 0xb2,
	0x9b, 0xf1, 0x01, 0x22, 0xfa, 0x62, 0x69, 0xaa, 0xac, 0xf9, 0x68, 0x3c, 0xa8, 0xeb, 0x68, 0x96,
	0x88, 0xd6, 0x82, 0xf0, 0x16, 0xc4, 0x8b, 0xb7, 0x90, 0x42, 0x8b, 0x37, 0xc1, 0x2c, 0x85, 0x70,
	0xde, 0x89, 0xc3, 0x27, 0x38, 0xf6, 0x32, 0xae, 0xc4, 0x4e, 0xef, 0x7f, 0x55, 0xc9, 0xf5, 0x9e,
	0xb6, 0x55, 0xb3, 0x18, 0xba, 0x9d, 0xf7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x37, 0x01,
	0xef, 0xa3, 0x01, 0x00, 0x00,
}
