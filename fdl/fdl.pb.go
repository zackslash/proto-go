// Code generated by protoc-gen-go.
// source: fdl.proto
// DO NOT EDIT!

/*
Package fdl is a generated protocol buffer package.

It is generated from these files:
	fdl.proto

It has these top-level messages:
	Property
	MutationRequest
	ReadRequest
	MutationResponse
	PropertiesResponse
*/
package fdl

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

type PropertyType int32

const (
	PropertyType_META PropertyType = 0
	PropertyType_DATA PropertyType = 1
	PropertyType_LIST PropertyType = 2
)

var PropertyType_name = map[int32]string{
	0: "META",
	1: "DATA",
	2: "LIST",
}
var PropertyType_value = map[string]int32{
	"META": 0,
	"DATA": 1,
	"LIST": 2,
}

func (x PropertyType) String() string {
	return proto.EnumName(PropertyType_name, int32(x))
}
func (PropertyType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type MutationMode int32

const (
	MutationMode_NONE   MutationMode = 0
	MutationMode_WRITE  MutationMode = 1
	MutationMode_APPEND MutationMode = 2
	MutationMode_REMOVE MutationMode = 3
	MutationMode_DELETE MutationMode = 4
)

var MutationMode_name = map[int32]string{
	0: "NONE",
	1: "WRITE",
	2: "APPEND",
	3: "REMOVE",
	4: "DELETE",
}
var MutationMode_value = map[string]int32{
	"NONE":   0,
	"WRITE":  1,
	"APPEND": 2,
	"REMOVE": 3,
	"DELETE": 4,
}

func (x MutationMode) String() string {
	return proto.EnumName(MutationMode_name, int32(x))
}
func (MutationMode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type MutationStatus int32

const (
	MutationStatus_FAILURE MutationStatus = 0
	MutationStatus_SUCCESS MutationStatus = 1
)

var MutationStatus_name = map[int32]string{
	0: "FAILURE",
	1: "SUCCESS",
}
var MutationStatus_value = map[string]int32{
	"FAILURE": 0,
	"SUCCESS": 1,
}

func (x MutationStatus) String() string {
	return proto.EnumName(MutationStatus_name, int32(x))
}
func (MutationStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Property struct {
	Property string       `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	Type     PropertyType `protobuf:"varint,2,opt,name=type,enum=fdl.PropertyType" json:"type,omitempty"`
	Value    string       `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	Mode     MutationMode `protobuf:"varint,4,opt,name=mode,enum=fdl.MutationMode" json:"mode,omitempty"`
}

func (m *Property) Reset()                    { *m = Property{} }
func (m *Property) String() string            { return proto.CompactTextString(m) }
func (*Property) ProtoMessage()               {}
func (*Property) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type MutationRequest struct {
	Fid        string      `protobuf:"bytes,1,opt,name=fid" json:"fid,omitempty"`
	Properties []*Property `protobuf:"bytes,2,rep,name=properties" json:"properties,omitempty"`
	MemberId   string      `protobuf:"bytes,3,opt,name=member_id,json=memberId" json:"member_id,omitempty"`
}

func (m *MutationRequest) Reset()                    { *m = MutationRequest{} }
func (m *MutationRequest) String() string            { return proto.CompactTextString(m) }
func (*MutationRequest) ProtoMessage()               {}
func (*MutationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MutationRequest) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

type ReadRequest struct {
	Fid        string      `protobuf:"bytes,1,opt,name=fid" json:"fid,omitempty"`
	Properties []*Property `protobuf:"bytes,2,rep,name=properties" json:"properties,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadRequest) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

type MutationResponse struct {
	Fid           string         `protobuf:"bytes,1,opt,name=fid" json:"fid,omitempty"`
	Status        MutationStatus `protobuf:"varint,2,opt,name=status,enum=fdl.MutationStatus" json:"status,omitempty"`
	StatusMessage string         `protobuf:"bytes,3,opt,name=status_message,json=statusMessage" json:"status_message,omitempty"`
}

func (m *MutationResponse) Reset()                    { *m = MutationResponse{} }
func (m *MutationResponse) String() string            { return proto.CompactTextString(m) }
func (*MutationResponse) ProtoMessage()               {}
func (*MutationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type PropertiesResponse struct {
	Fid        string      `protobuf:"bytes,1,opt,name=fid" json:"fid,omitempty"`
	Properties []*Property `protobuf:"bytes,2,rep,name=properties" json:"properties,omitempty"`
}

func (m *PropertiesResponse) Reset()                    { *m = PropertiesResponse{} }
func (m *PropertiesResponse) String() string            { return proto.CompactTextString(m) }
func (*PropertiesResponse) ProtoMessage()               {}
func (*PropertiesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PropertiesResponse) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

func init() {
	proto.RegisterType((*Property)(nil), "fdl.Property")
	proto.RegisterType((*MutationRequest)(nil), "fdl.MutationRequest")
	proto.RegisterType((*ReadRequest)(nil), "fdl.ReadRequest")
	proto.RegisterType((*MutationResponse)(nil), "fdl.MutationResponse")
	proto.RegisterType((*PropertiesResponse)(nil), "fdl.PropertiesResponse")
	proto.RegisterEnum("fdl.PropertyType", PropertyType_name, PropertyType_value)
	proto.RegisterEnum("fdl.MutationMode", MutationMode_name, MutationMode_value)
	proto.RegisterEnum("fdl.MutationStatus", MutationStatus_name, MutationStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Fdl service

type FdlClient interface {
	Mutate(ctx context.Context, in *MutationRequest, opts ...grpc.CallOption) (*MutationResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*PropertiesResponse, error)
}

type fdlClient struct {
	cc *grpc.ClientConn
}

func NewFdlClient(cc *grpc.ClientConn) FdlClient {
	return &fdlClient{cc}
}

func (c *fdlClient) Mutate(ctx context.Context, in *MutationRequest, opts ...grpc.CallOption) (*MutationResponse, error) {
	out := new(MutationResponse)
	err := grpc.Invoke(ctx, "/fdl.Fdl/Mutate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fdlClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*PropertiesResponse, error) {
	out := new(PropertiesResponse)
	err := grpc.Invoke(ctx, "/fdl.Fdl/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Fdl service

type FdlServer interface {
	Mutate(context.Context, *MutationRequest) (*MutationResponse, error)
	Read(context.Context, *ReadRequest) (*PropertiesResponse, error)
}

func RegisterFdlServer(s *grpc.Server, srv FdlServer) {
	s.RegisterService(&_Fdl_serviceDesc, srv)
}

func _Fdl_Mutate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FdlServer).Mutate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fdl.Fdl/Mutate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FdlServer).Mutate(ctx, req.(*MutationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fdl_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FdlServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fdl.Fdl/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FdlServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Fdl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fdl.Fdl",
	HandlerType: (*FdlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Mutate",
			Handler:    _Fdl_Mutate_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Fdl_Read_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("fdl.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x53, 0xd1, 0x8e, 0x93, 0x40,
	0x14, 0x95, 0xc2, 0x22, 0xdc, 0xee, 0xd6, 0x71, 0xd4, 0x48, 0xea, 0x8b, 0x21, 0x31, 0xd9, 0xa0,
	0x6e, 0x62, 0x8d, 0x1f, 0x40, 0xb6, 0xb3, 0x86, 0xa4, 0xb0, 0x64, 0xa0, 0xfa, 0xb8, 0x61, 0xc3,
	0xd4, 0x34, 0x29, 0x1d, 0x64, 0xc0, 0xd8, 0x5f, 0xf0, 0xab, 0x9d, 0x01, 0x5a, 0xc1, 0x18, 0x13,
	0x13, 0xdf, 0xce, 0xdc, 0x73, 0xee, 0xe5, 0xcc, 0x99, 0x0b, 0xd8, 0x9b, 0x7c, 0x77, 0x55, 0x56,
	0xbc, 0xe6, 0x58, 0x97, 0xd0, 0xfd, 0xa1, 0x81, 0x15, 0x57, 0xbc, 0x64, 0x55, 0x7d, 0xc0, 0x73,
	0xb0, 0xca, 0x1e, 0x3b, 0xda, 0x4b, 0xed, 0xd2, 0xa6, 0xa7, 0x33, 0x7e, 0x05, 0x46, 0x7d, 0x28,
	0x99, 0x33, 0x91, 0xf5, 0xd9, 0xe2, 0xf1, 0x95, 0x9a, 0x73, 0x6c, 0x4c, 0x25, 0x41, 0x5b, 0x1a,
	0x3f, 0x85, 0xb3, 0x6f, 0xd9, 0xae, 0x61, 0x8e, 0xde, 0xf6, 0x77, 0x07, 0xd5, 0x5c, 0xf0, 0x9c,
	0x39, 0xc6, 0xa0, 0x39, 0x6c, 0xea, 0xac, 0xde, 0xf2, 0x7d, 0x28, 0x09, 0xda, 0xd2, 0x2e, 0x87,
	0x47, 0xc7, 0x2a, 0x65, 0x5f, 0x1b, 0x26, 0x6a, 0x8c, 0x40, 0xdf, 0x6c, 0xf3, 0xde, 0x8d, 0x82,
	0xf8, 0x2d, 0x40, 0x6f, 0x6a, 0xcb, 0x84, 0xb4, 0xa3, 0x5f, 0x4e, 0x17, 0x17, 0x23, 0x3b, 0x74,
	0x20, 0xc0, 0x2f, 0xc0, 0x2e, 0x58, 0x71, 0xcf, 0xaa, 0x3b, 0x39, 0xa6, 0x33, 0x65, 0x75, 0x85,
	0x20, 0x77, 0x23, 0x98, 0x52, 0x96, 0xe5, 0xff, 0xeb, 0x63, 0xee, 0x77, 0x40, 0xbf, 0x2e, 0x20,
	0x4a, 0xbe, 0x17, 0xec, 0x0f, 0x43, 0x5f, 0x83, 0x29, 0xa4, 0xa8, 0x11, 0x7d, 0x98, 0x4f, 0x46,
	0x79, 0x24, 0x2d, 0x45, 0x7b, 0x89, 0x8c, 0x6e, 0xd6, 0xa1, 0xbb, 0x82, 0x09, 0x91, 0x7d, 0x39,
	0x26, 0x7b, 0xd1, 0x55, 0xc3, 0xae, 0xe8, 0xae, 0x01, 0xc7, 0x27, 0x1f, 0x7f, 0xf9, 0xf6, 0xbf,
	0x5d, 0xc8, 0x7b, 0x03, 0xe7, 0xc3, 0x47, 0xc6, 0x16, 0x18, 0x21, 0x49, 0x7d, 0xf4, 0x40, 0xa1,
	0xa5, 0x2f, 0x91, 0xa6, 0xd0, 0x2a, 0x48, 0x52, 0x34, 0xf1, 0x3e, 0xc2, 0xf9, 0xf0, 0x55, 0x15,
	0x13, 0xdd, 0x46, 0x44, 0xaa, 0x6d, 0x38, 0xfb, 0x4c, 0x83, 0x94, 0x48, 0x39, 0x80, 0xe9, 0xc7,
	0x31, 0x89, 0x96, 0x68, 0xa2, 0x30, 0x25, 0xe1, 0xed, 0x27, 0x82, 0x74, 0x85, 0x97, 0x64, 0x45,
	0xa4, 0xc6, 0xf0, 0x3c, 0x98, 0x8d, 0xe3, 0xc0, 0x53, 0x78, 0x78, 0xe3, 0x07, 0xab, 0x35, 0x55,
	0xd3, 0xe4, 0x21, 0x59, 0x5f, 0x5f, 0x93, 0x24, 0x41, 0xda, 0x82, 0x83, 0x7e, 0x93, 0xef, 0xf0,
	0x07, 0x30, 0xdb, 0x16, 0xb9, 0x82, 0xa3, 0x38, 0xfb, 0xb7, 0x9d, 0x3f, 0xfb, 0xad, 0xda, 0x27,
	0xf4, 0x0e, 0x0c, 0xb5, 0x01, 0x18, 0xb5, 0xf4, 0x60, 0x19, 0xe6, 0xcf, 0x87, 0xa9, 0x0c, 0x42,
	0xbd, 0x37, 0xdb, 0xdf, 0xe7, 0xfd, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x73, 0x25, 0x56, 0x29,
	0x4b, 0x03, 0x00, 0x00,
}
