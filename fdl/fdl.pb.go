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
	PropertyType_META    PropertyType = 0
	PropertyType_DATA    PropertyType = 1
	PropertyType_COUNTER PropertyType = 2
	PropertyType_SET     PropertyType = 3
)

var PropertyType_name = map[int32]string{
	0: "META",
	1: "DATA",
	2: "COUNTER",
	3: "SET",
}
var PropertyType_value = map[string]int32{
	"META":    0,
	"DATA":    1,
	"COUNTER": 2,
	"SET":     3,
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

func (m *Property) GetProperty() string {
	if m != nil {
		return m.Property
	}
	return ""
}

func (m *Property) GetType() PropertyType {
	if m != nil {
		return m.Type
	}
	return PropertyType_META
}

func (m *Property) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Property) GetMode() MutationMode {
	if m != nil {
		return m.Mode
	}
	return MutationMode_NONE
}

type MutationRequest struct {
	Fid        string      `protobuf:"bytes,1,opt,name=fid" json:"fid,omitempty"`
	Properties []*Property `protobuf:"bytes,2,rep,name=properties" json:"properties,omitempty"`
	MemberId   string      `protobuf:"bytes,3,opt,name=member_id,json=memberId" json:"member_id,omitempty"`
}

func (m *MutationRequest) Reset()                    { *m = MutationRequest{} }
func (m *MutationRequest) String() string            { return proto.CompactTextString(m) }
func (*MutationRequest) ProtoMessage()               {}
func (*MutationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MutationRequest) GetFid() string {
	if m != nil {
		return m.Fid
	}
	return ""
}

func (m *MutationRequest) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *MutationRequest) GetMemberId() string {
	if m != nil {
		return m.MemberId
	}
	return ""
}

type ReadRequest struct {
	Fid        string      `protobuf:"bytes,1,opt,name=fid" json:"fid,omitempty"`
	Properties []*Property `protobuf:"bytes,2,rep,name=properties" json:"properties,omitempty"`
	MemberId   string      `protobuf:"bytes,3,opt,name=member_id,json=memberId" json:"member_id,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadRequest) GetFid() string {
	if m != nil {
		return m.Fid
	}
	return ""
}

func (m *ReadRequest) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *ReadRequest) GetMemberId() string {
	if m != nil {
		return m.MemberId
	}
	return ""
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

func (m *MutationResponse) GetFid() string {
	if m != nil {
		return m.Fid
	}
	return ""
}

func (m *MutationResponse) GetStatus() MutationStatus {
	if m != nil {
		return m.Status
	}
	return MutationStatus_FAILURE
}

func (m *MutationResponse) GetStatusMessage() string {
	if m != nil {
		return m.StatusMessage
	}
	return ""
}

type PropertiesResponse struct {
	Fid        string      `protobuf:"bytes,1,opt,name=fid" json:"fid,omitempty"`
	Properties []*Property `protobuf:"bytes,2,rep,name=properties" json:"properties,omitempty"`
}

func (m *PropertiesResponse) Reset()                    { *m = PropertiesResponse{} }
func (m *PropertiesResponse) String() string            { return proto.CompactTextString(m) }
func (*PropertiesResponse) ProtoMessage()               {}
func (*PropertiesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PropertiesResponse) GetFid() string {
	if m != nil {
		return m.Fid
	}
	return ""
}

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
const _ = grpc.SupportPackageIsVersion4

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
	Metadata: "fdl.proto",
}

func init() { proto.RegisterFile("fdl.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x53, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0xdd, 0x34, 0xd9, 0x7e, 0xdc, 0xee, 0xd6, 0x71, 0x5c, 0x31, 0xd4, 0x07, 0x6b, 0x61, 0xa1,
	0x54, 0xb6, 0xc5, 0x8a, 0xe0, 0x6b, 0x6c, 0x67, 0xa5, 0xb0, 0x69, 0xcb, 0x24, 0x55, 0xf0, 0x65,
	0x49, 0x9d, 0x49, 0x0d, 0x26, 0x3b, 0x31, 0x99, 0x88, 0xfd, 0x0b, 0xfe, 0x6a, 0x99, 0x24, 0x5d,
	0x13, 0x11, 0xc1, 0x17, 0xdf, 0xce, 0x9c, 0x7b, 0xcf, 0xdc, 0x33, 0x27, 0x37, 0xd0, 0xf1, 0x59,
	0x38, 0x89, 0x13, 0x21, 0x05, 0xd6, 0x7d, 0x16, 0x0e, 0x7f, 0x68, 0xd0, 0xde, 0x24, 0x22, 0xe6,
	0x89, 0x3c, 0xe0, 0x3e, 0xb4, 0xe3, 0x12, 0x9b, 0xda, 0x40, 0x1b, 0x75, 0xe8, 0xfd, 0x19, 0x5f,
	0x82, 0x21, 0x0f, 0x31, 0x37, 0x1b, 0x03, 0x6d, 0xd4, 0x9b, 0x3d, 0x9c, 0xa8, 0x7b, 0x8e, 0x42,
	0xf7, 0x10, 0x73, 0x9a, 0x97, 0xf1, 0x05, 0x9c, 0x7e, 0xf3, 0xc2, 0x8c, 0x9b, 0x7a, 0xae, 0x2f,
	0x0e, 0x4a, 0x1c, 0x09, 0xc6, 0x4d, 0xa3, 0x22, 0xb6, 0x33, 0xe9, 0xc9, 0x40, 0xdc, 0xd9, 0x82,
	0x71, 0x9a, 0x97, 0x87, 0x02, 0x1e, 0x1c, 0x59, 0xca, 0xbf, 0x66, 0x3c, 0x95, 0x18, 0x81, 0xee,
	0x07, 0xac, 0x74, 0xa3, 0x20, 0xbe, 0x02, 0x28, 0x4d, 0x05, 0x3c, 0x35, 0x1b, 0x03, 0x7d, 0xd4,
	0x9d, 0x9d, 0xd7, 0xec, 0xd0, 0x4a, 0x03, 0x7e, 0x0a, 0x9d, 0x88, 0x47, 0x3b, 0x9e, 0xdc, 0x06,
	0xac, 0x34, 0xd5, 0x2e, 0x88, 0x25, 0x1b, 0x7e, 0x81, 0x2e, 0xe5, 0x1e, 0xfb, 0x3f, 0xc3, 0xbe,
	0x03, 0xfa, 0xf5, 0xba, 0x34, 0x16, 0x77, 0x29, 0xff, 0xc3, 0xc4, 0x17, 0xd0, 0x4c, 0xa5, 0x27,
	0xb3, 0xb4, 0x4c, 0xfa, 0x51, 0x2d, 0x2c, 0x27, 0x2f, 0xd1, 0xb2, 0x05, 0x5f, 0x42, 0xaf, 0x40,
	0xb7, 0x11, 0x4f, 0x53, 0x6f, 0x7f, 0x8c, 0xfd, 0xbc, 0x60, 0xed, 0x82, 0x1c, 0x6e, 0x01, 0x6f,
	0xee, 0x4d, 0xfe, 0x65, 0xf6, 0xbf, 0xbd, 0x76, 0xfc, 0x06, 0xce, 0xaa, 0x1b, 0x80, 0xdb, 0x60,
	0xd8, 0xc4, 0xb5, 0xd0, 0x89, 0x42, 0x0b, 0xcb, 0xb5, 0x90, 0x86, 0xbb, 0xd0, 0x9a, 0xaf, 0xb7,
	0x2b, 0x97, 0x50, 0xd4, 0xc0, 0x2d, 0xd0, 0x1d, 0xe2, 0x22, 0x7d, 0xfc, 0x0e, 0xce, 0xaa, 0x9f,
	0x5f, 0xf5, 0xaf, 0xd6, 0x2b, 0x82, 0x4e, 0x70, 0x07, 0x4e, 0x3f, 0xd0, 0xa5, 0x4b, 0x90, 0x86,
	0x01, 0x9a, 0xd6, 0x66, 0x43, 0x56, 0x0b, 0xd4, 0x50, 0x98, 0x12, 0x7b, 0xfd, 0x9e, 0x20, 0x5d,
	0xe1, 0x05, 0xb9, 0x21, 0x2e, 0x41, 0xc6, 0x78, 0x0c, 0xbd, 0x7a, 0x34, 0x6a, 0xe0, 0xb5, 0xb5,
	0xbc, 0xd9, 0x52, 0x75, 0x5b, 0x17, 0x5a, 0xce, 0x76, 0x3e, 0x27, 0x8e, 0x83, 0xb4, 0x99, 0x00,
	0xfd, 0x9a, 0x85, 0xf8, 0x35, 0x34, 0x73, 0x09, 0xc7, 0x17, 0xb5, 0x68, 0xcb, 0x25, 0xe8, 0x3f,
	0xfe, 0x8d, 0x2d, 0xd3, 0x7a, 0x09, 0x86, 0x5a, 0x15, 0x8c, 0xf2, 0x72, 0x65, 0x6b, 0xfa, 0x4f,
	0xaa, 0x09, 0x55, 0x02, 0x7e, 0xfb, 0xfc, 0xe3, 0xb3, 0x7d, 0x20, 0x3f, 0x67, 0xbb, 0xc9, 0x27,
	0x11, 0x4d, 0x7d, 0x91, 0xc8, 0xc0, 0x0f, 0xa6, 0xf9, 0xcf, 0x77, 0xb5, 0x17, 0x53, 0x9f, 0x85,
	0xbb, 0x66, 0x7e, 0x7a, 0xf5, 0x33, 0x00, 0x00, 0xff, 0xff, 0x71, 0x28, 0x75, 0xf9, 0x97, 0x03,
	0x00, 0x00,
}
