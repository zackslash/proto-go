// Code generated by protoc-gen-go.
// source: discovery.proto
// DO NOT EDIT!

/*
Package discovery is a generated protocol buffer package.

It is generated from these files:
	discovery.proto

It has these top-level messages:
	RegisterRequest
	DeRegisterRequest
	HeartBeatRequest
	LocationRequest
	StatusRequest
	DiscoveryResponse
	ServiceLocation
*/
package discovery

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

type ServiceStatus int32

const (
	ServiceStatus_OFFLINE              ServiceStatus = 0
	ServiceStatus_ONLINE               ServiceStatus = 1
	ServiceStatus_DEGRADED_PERFORMANCE ServiceStatus = 2
	ServiceStatus_PARTIAL_OUTAGE       ServiceStatus = 3
	ServiceStatus_UNDER_MAINTENANCE    ServiceStatus = 5
)

var ServiceStatus_name = map[int32]string{
	0: "OFFLINE",
	1: "ONLINE",
	2: "DEGRADED_PERFORMANCE",
	3: "PARTIAL_OUTAGE",
	5: "UNDER_MAINTENANCE",
}
var ServiceStatus_value = map[string]int32{
	"OFFLINE":              0,
	"ONLINE":               1,
	"DEGRADED_PERFORMANCE": 2,
	"PARTIAL_OUTAGE":       3,
	"UNDER_MAINTENANCE":    5,
}

func (x ServiceStatus) String() string {
	return proto.EnumName(ServiceStatus_name, int32(x))
}
func (ServiceStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AppVersion int32

const (
	AppVersion_STABLE AppVersion = 0
	AppVersion_BETA   AppVersion = 1
	AppVersion_ALPHA  AppVersion = 2
)

var AppVersion_name = map[int32]string{
	0: "STABLE",
	1: "BETA",
	2: "ALPHA",
}
var AppVersion_value = map[string]int32{
	"STABLE": 0,
	"BETA":   1,
	"ALPHA":  2,
}

func (x AppVersion) String() string {
	return proto.EnumName(AppVersion_name, int32(x))
}
func (AppVersion) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type StatusTarget int32

const (
	StatusTarget_INSTANCE StatusTarget = 0
	StatusTarget_SERVICE  StatusTarget = 1
	StatusTarget_BOTH     StatusTarget = 2
)

var StatusTarget_name = map[int32]string{
	0: "INSTANCE",
	1: "SERVICE",
	2: "BOTH",
}
var StatusTarget_value = map[string]int32{
	"INSTANCE": 0,
	"SERVICE":  1,
	"BOTH":     2,
}

func (x StatusTarget) String() string {
	return proto.EnumName(StatusTarget_name, int32(x))
}
func (StatusTarget) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type RegisterRequest struct {
	AppId        string     `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	InstanceUuid string     `protobuf:"bytes,2,opt,name=instance_uuid,json=instanceUuid" json:"instance_uuid,omitempty"`
	ServiceHost  string     `protobuf:"bytes,3,opt,name=service_host,json=serviceHost" json:"service_host,omitempty"`
	ServicePort  int32      `protobuf:"varint,4,opt,name=service_port,json=servicePort" json:"service_port,omitempty"`
	Version      AppVersion `protobuf:"varint,5,opt,name=version,enum=discovery.AppVersion" json:"version,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RegisterRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *RegisterRequest) GetInstanceUuid() string {
	if m != nil {
		return m.InstanceUuid
	}
	return ""
}

func (m *RegisterRequest) GetServiceHost() string {
	if m != nil {
		return m.ServiceHost
	}
	return ""
}

func (m *RegisterRequest) GetServicePort() int32 {
	if m != nil {
		return m.ServicePort
	}
	return 0
}

func (m *RegisterRequest) GetVersion() AppVersion {
	if m != nil {
		return m.Version
	}
	return AppVersion_STABLE
}

type DeRegisterRequest struct {
	AppId        string     `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	InstanceUuid string     `protobuf:"bytes,2,opt,name=instance_uuid,json=instanceUuid" json:"instance_uuid,omitempty"`
	Version      AppVersion `protobuf:"varint,3,opt,name=version,enum=discovery.AppVersion" json:"version,omitempty"`
}

func (m *DeRegisterRequest) Reset()                    { *m = DeRegisterRequest{} }
func (m *DeRegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*DeRegisterRequest) ProtoMessage()               {}
func (*DeRegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeRegisterRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *DeRegisterRequest) GetInstanceUuid() string {
	if m != nil {
		return m.InstanceUuid
	}
	return ""
}

func (m *DeRegisterRequest) GetVersion() AppVersion {
	if m != nil {
		return m.Version
	}
	return AppVersion_STABLE
}

type HeartBeatRequest struct {
	AppId        string     `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	InstanceUuid string     `protobuf:"bytes,2,opt,name=instance_uuid,json=instanceUuid" json:"instance_uuid,omitempty"`
	Version      AppVersion `protobuf:"varint,3,opt,name=version,enum=discovery.AppVersion" json:"version,omitempty"`
}

func (m *HeartBeatRequest) Reset()                    { *m = HeartBeatRequest{} }
func (m *HeartBeatRequest) String() string            { return proto.CompactTextString(m) }
func (*HeartBeatRequest) ProtoMessage()               {}
func (*HeartBeatRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HeartBeatRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *HeartBeatRequest) GetInstanceUuid() string {
	if m != nil {
		return m.InstanceUuid
	}
	return ""
}

func (m *HeartBeatRequest) GetVersion() AppVersion {
	if m != nil {
		return m.Version
	}
	return AppVersion_STABLE
}

type LocationRequest struct {
	AppId   string     `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	Version AppVersion `protobuf:"varint,2,opt,name=version,enum=discovery.AppVersion" json:"version,omitempty"`
}

func (m *LocationRequest) Reset()                    { *m = LocationRequest{} }
func (m *LocationRequest) String() string            { return proto.CompactTextString(m) }
func (*LocationRequest) ProtoMessage()               {}
func (*LocationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LocationRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *LocationRequest) GetVersion() AppVersion {
	if m != nil {
		return m.Version
	}
	return AppVersion_STABLE
}

type StatusRequest struct {
	AppId        string        `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	InstanceUuid string        `protobuf:"bytes,2,opt,name=instance_uuid,json=instanceUuid" json:"instance_uuid,omitempty"`
	Message      string        `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Status       ServiceStatus `protobuf:"varint,4,opt,name=status,enum=discovery.ServiceStatus" json:"status,omitempty"`
	Target       StatusTarget  `protobuf:"varint,5,opt,name=target,enum=discovery.StatusTarget" json:"target,omitempty"`
	Version      AppVersion    `protobuf:"varint,6,opt,name=version,enum=discovery.AppVersion" json:"version,omitempty"`
}

func (m *StatusRequest) Reset()                    { *m = StatusRequest{} }
func (m *StatusRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()               {}
func (*StatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StatusRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *StatusRequest) GetInstanceUuid() string {
	if m != nil {
		return m.InstanceUuid
	}
	return ""
}

func (m *StatusRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *StatusRequest) GetStatus() ServiceStatus {
	if m != nil {
		return m.Status
	}
	return ServiceStatus_OFFLINE
}

func (m *StatusRequest) GetTarget() StatusTarget {
	if m != nil {
		return m.Target
	}
	return StatusTarget_INSTANCE
}

func (m *StatusRequest) GetVersion() AppVersion {
	if m != nil {
		return m.Version
	}
	return AppVersion_STABLE
}

type DiscoveryResponse struct {
	Recorded bool `protobuf:"varint,1,opt,name=recorded" json:"recorded,omitempty"`
}

func (m *DiscoveryResponse) Reset()                    { *m = DiscoveryResponse{} }
func (m *DiscoveryResponse) String() string            { return proto.CompactTextString(m) }
func (*DiscoveryResponse) ProtoMessage()               {}
func (*DiscoveryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DiscoveryResponse) GetRecorded() bool {
	if m != nil {
		return m.Recorded
	}
	return false
}

type ServiceLocation struct {
	AppId       string        `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	ServiceHost string        `protobuf:"bytes,2,opt,name=service_host,json=serviceHost" json:"service_host,omitempty"`
	ServicePort int32         `protobuf:"varint,3,opt,name=service_port,json=servicePort" json:"service_port,omitempty"`
	Status      ServiceStatus `protobuf:"varint,4,opt,name=status,enum=discovery.ServiceStatus" json:"status,omitempty"`
	Version     AppVersion    `protobuf:"varint,5,opt,name=version,enum=discovery.AppVersion" json:"version,omitempty"`
}

func (m *ServiceLocation) Reset()                    { *m = ServiceLocation{} }
func (m *ServiceLocation) String() string            { return proto.CompactTextString(m) }
func (*ServiceLocation) ProtoMessage()               {}
func (*ServiceLocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ServiceLocation) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *ServiceLocation) GetServiceHost() string {
	if m != nil {
		return m.ServiceHost
	}
	return ""
}

func (m *ServiceLocation) GetServicePort() int32 {
	if m != nil {
		return m.ServicePort
	}
	return 0
}

func (m *ServiceLocation) GetStatus() ServiceStatus {
	if m != nil {
		return m.Status
	}
	return ServiceStatus_OFFLINE
}

func (m *ServiceLocation) GetVersion() AppVersion {
	if m != nil {
		return m.Version
	}
	return AppVersion_STABLE
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "discovery.RegisterRequest")
	proto.RegisterType((*DeRegisterRequest)(nil), "discovery.DeRegisterRequest")
	proto.RegisterType((*HeartBeatRequest)(nil), "discovery.HeartBeatRequest")
	proto.RegisterType((*LocationRequest)(nil), "discovery.LocationRequest")
	proto.RegisterType((*StatusRequest)(nil), "discovery.StatusRequest")
	proto.RegisterType((*DiscoveryResponse)(nil), "discovery.DiscoveryResponse")
	proto.RegisterType((*ServiceLocation)(nil), "discovery.ServiceLocation")
	proto.RegisterEnum("discovery.ServiceStatus", ServiceStatus_name, ServiceStatus_value)
	proto.RegisterEnum("discovery.AppVersion", AppVersion_name, AppVersion_value)
	proto.RegisterEnum("discovery.StatusTarget", StatusTarget_name, StatusTarget_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Discovery service

type DiscoveryClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
	DeRegister(ctx context.Context, in *DeRegisterRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
	HeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
	GetLocation(ctx context.Context, in *LocationRequest, opts ...grpc.CallOption) (*ServiceLocation, error)
}

type discoveryClient struct {
	cc *grpc.ClientConn
}

func NewDiscoveryClient(cc *grpc.ClientConn) DiscoveryClient {
	return &discoveryClient{cc}
}

func (c *discoveryClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := grpc.Invoke(ctx, "/discovery.Discovery/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) DeRegister(ctx context.Context, in *DeRegisterRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := grpc.Invoke(ctx, "/discovery.Discovery/DeRegister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) HeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := grpc.Invoke(ctx, "/discovery.Discovery/HeartBeat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := grpc.Invoke(ctx, "/discovery.Discovery/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) GetLocation(ctx context.Context, in *LocationRequest, opts ...grpc.CallOption) (*ServiceLocation, error) {
	out := new(ServiceLocation)
	err := grpc.Invoke(ctx, "/discovery.Discovery/GetLocation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Discovery service

type DiscoveryServer interface {
	Register(context.Context, *RegisterRequest) (*DiscoveryResponse, error)
	DeRegister(context.Context, *DeRegisterRequest) (*DiscoveryResponse, error)
	HeartBeat(context.Context, *HeartBeatRequest) (*DiscoveryResponse, error)
	Status(context.Context, *StatusRequest) (*DiscoveryResponse, error)
	GetLocation(context.Context, *LocationRequest) (*ServiceLocation, error)
}

func RegisterDiscoveryServer(s *grpc.Server, srv DiscoveryServer) {
	s.RegisterService(&_Discovery_serviceDesc, srv)
}

func _Discovery_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.Discovery/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discovery_DeRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).DeRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.Discovery/DeRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).DeRegister(ctx, req.(*DeRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discovery_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.Discovery/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).HeartBeat(ctx, req.(*HeartBeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discovery_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.Discovery/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discovery_GetLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).GetLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.Discovery/GetLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).GetLocation(ctx, req.(*LocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Discovery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discovery.Discovery",
	HandlerType: (*DiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Discovery_Register_Handler,
		},
		{
			MethodName: "DeRegister",
			Handler:    _Discovery_DeRegister_Handler,
		},
		{
			MethodName: "HeartBeat",
			Handler:    _Discovery_HeartBeat_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Discovery_Status_Handler,
		},
		{
			MethodName: "GetLocation",
			Handler:    _Discovery_GetLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discovery.proto",
}

func init() { proto.RegisterFile("discovery.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 621 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x55, 0xcd, 0x4e, 0xdb, 0x4c,
	0x14, 0xc5, 0x0e, 0x31, 0xc9, 0x25, 0x10, 0x33, 0xfa, 0xd0, 0x67, 0xa5, 0x5d, 0x50, 0xaa, 0xaa,
	0x08, 0x09, 0x5c, 0xc1, 0x0b, 0xd4, 0xc1, 0x0e, 0xb1, 0x14, 0x9c, 0x68, 0x62, 0x90, 0xda, 0x8d,
	0x65, 0xec, 0x21, 0x78, 0x41, 0xc6, 0x9d, 0x19, 0x23, 0x75, 0xd5, 0x3e, 0x5c, 0x5f, 0xa0, 0xef,
	0xd1, 0x17, 0xe8, 0xae, 0xb2, 0x63, 0x3b, 0x26, 0x51, 0x4b, 0xa8, 0x90, 0xba, 0xf3, 0xdc, 0x7b,
	0xe6, 0xcc, 0xb9, 0xbf, 0x86, 0x76, 0x18, 0xf1, 0x80, 0xde, 0x13, 0xf6, 0xf9, 0x38, 0x66, 0x54,
	0x50, 0xd4, 0x2c, 0x0d, 0xfb, 0xdf, 0x24, 0x68, 0x63, 0x32, 0x89, 0xb8, 0x20, 0x0c, 0x93, 0x4f,
	0x09, 0xe1, 0x02, 0xed, 0x82, 0xe2, 0xc7, 0xb1, 0x17, 0x85, 0x9a, 0xb4, 0x27, 0x1d, 0x34, 0x71,
	0xdd, 0x8f, 0x63, 0x3b, 0x44, 0xaf, 0x61, 0x2b, 0x9a, 0x72, 0xe1, 0x4f, 0x03, 0xe2, 0x25, 0x49,
	0x14, 0x6a, 0x72, 0xe6, 0x6d, 0x15, 0xc6, 0xcb, 0x24, 0x0a, 0xd1, 0x2b, 0x68, 0x71, 0xc2, 0xee,
	0xa3, 0x80, 0x78, 0xb7, 0x94, 0x0b, 0xad, 0x96, 0x61, 0x36, 0x73, 0x5b, 0x9f, 0x72, 0x51, 0x85,
	0xc4, 0x94, 0x09, 0x6d, 0x7d, 0x4f, 0x3a, 0xa8, 0x97, 0x90, 0x11, 0x65, 0x02, 0xe9, 0xb0, 0x71,
	0x4f, 0x18, 0x8f, 0xe8, 0x54, 0xab, 0xef, 0x49, 0x07, 0xdb, 0x27, 0xbb, 0xc7, 0xf3, 0x18, 0x8c,
	0x38, 0xbe, 0x9a, 0x39, 0x71, 0x81, 0xda, 0xff, 0x2a, 0xc1, 0x8e, 0x49, 0x9e, 0x33, 0x90, 0x8a,
	0x84, 0xda, 0x4a, 0x12, 0xbe, 0x80, 0xda, 0x27, 0x3e, 0x13, 0x5d, 0xe2, 0x8b, 0x7f, 0x22, 0xe0,
	0x03, 0xb4, 0x07, 0x34, 0xf0, 0x45, 0x6a, 0xfc, 0xf3, 0xfb, 0x15, 0x6a, 0x79, 0x25, 0xea, 0x9f,
	0x12, 0x6c, 0x8d, 0x85, 0x2f, 0x12, 0xfe, 0x1c, 0x91, 0x69, 0xb0, 0x71, 0x47, 0x38, 0xf7, 0x27,
	0x24, 0x6f, 0x8f, 0xe2, 0x88, 0xde, 0x81, 0xc2, 0xb3, 0x67, 0xb2, 0xa6, 0xd8, 0x3e, 0xd1, 0x2a,
	0xba, 0xc6, 0xb3, 0xfe, 0xc8, 0x65, 0xe4, 0x38, 0xa4, 0x83, 0x22, 0x7c, 0x36, 0x21, 0x22, 0x6f,
	0x94, 0xff, 0xab, 0x37, 0x32, 0x88, 0x9b, 0xb9, 0x71, 0x0e, 0xab, 0xc6, 0xae, 0xac, 0x14, 0xbb,
	0x0e, 0x3b, 0x66, 0x01, 0xc0, 0x84, 0xc7, 0x74, 0xca, 0x09, 0xea, 0x40, 0x83, 0x91, 0x80, 0xb2,
	0x90, 0xcc, 0x12, 0xd0, 0xc0, 0xe5, 0x79, 0xff, 0xbb, 0x04, 0xed, 0x5c, 0x6c, 0x51, 0x8f, 0xdf,
	0xa5, 0x6b, 0x71, 0x5a, 0xe4, 0xc7, 0xa7, 0xa5, 0xb6, 0x3c, 0x2d, 0x7f, 0x93, 0xb5, 0xa7, 0xcd,
	0xd7, 0xe1, 0x1d, 0x6c, 0x3d, 0x60, 0x42, 0x9b, 0xb0, 0x31, 0xec, 0xf5, 0x06, 0xb6, 0x63, 0xa9,
	0x6b, 0x08, 0x40, 0x19, 0x3a, 0xd9, 0xb7, 0x84, 0x34, 0xf8, 0xcf, 0xb4, 0xce, 0xb1, 0x61, 0x5a,
	0xa6, 0x37, 0xb2, 0x70, 0x6f, 0x88, 0x2f, 0x0c, 0xe7, 0xcc, 0x52, 0x65, 0x84, 0x60, 0x7b, 0x64,
	0x60, 0xd7, 0x36, 0x06, 0xde, 0xf0, 0xd2, 0x35, 0xce, 0x2d, 0xb5, 0x86, 0x76, 0x61, 0xe7, 0xd2,
	0x31, 0x2d, 0xec, 0x5d, 0x18, 0xb6, 0xe3, 0x5a, 0x4e, 0x06, 0xad, 0x1f, 0x1e, 0x01, 0xcc, 0x55,
	0xa4, 0xf4, 0x63, 0xd7, 0xe8, 0x0e, 0xd2, 0xa7, 0x1a, 0xb0, 0xde, 0xb5, 0x5c, 0x43, 0x95, 0x50,
	0x13, 0xea, 0xc6, 0x60, 0xd4, 0x37, 0x54, 0xf9, 0xf0, 0x14, 0x5a, 0xd5, 0x5a, 0xa3, 0x16, 0x34,
	0x6c, 0x67, 0xec, 0x66, 0x64, 0x6b, 0xa9, 0xd4, 0xb1, 0x85, 0xaf, 0xec, 0xb3, 0x54, 0x5e, 0x7a,
	0x7f, 0xe8, 0xf6, 0x55, 0xf9, 0xe4, 0x87, 0x0c, 0xcd, 0xb2, 0xb0, 0xc8, 0x84, 0x46, 0xb1, 0x3d,
	0x50, 0xa7, 0x92, 0x8c, 0x85, 0x95, 0xd2, 0x79, 0x59, 0xf1, 0x2d, 0xb7, 0x45, 0x1f, 0x60, 0xbe,
	0x85, 0xd0, 0x03, 0x2c, 0x79, 0x1a, 0x53, 0x0f, 0x9a, 0xe5, 0x36, 0x41, 0x2f, 0x2a, 0xd0, 0xc5,
	0x1d, 0xf3, 0x08, 0xcf, 0x7b, 0x50, 0xf2, 0x8a, 0x69, 0x4b, 0x93, 0xb1, 0x1a, 0x83, 0x05, 0x9b,
	0xe7, 0x44, 0x94, 0x9d, 0x5c, 0x4d, 0xce, 0xc2, 0xba, 0xe9, 0x74, 0x96, 0x1b, 0xaf, 0x80, 0x74,
	0xdf, 0x7e, 0x7c, 0x33, 0x89, 0xc4, 0x6d, 0x72, 0x7d, 0x1c, 0xd0, 0x3b, 0xfd, 0x86, 0x32, 0x11,
	0xdd, 0x44, 0x7a, 0xf6, 0x3f, 0x3a, 0x9a, 0x50, 0xbd, 0xbc, 0x78, 0xad, 0x64, 0xb6, 0xd3, 0x5f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x4b, 0x70, 0x70, 0xb6, 0x06, 0x00, 0x00,
}
