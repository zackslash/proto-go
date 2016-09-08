// Code generated by protoc-gen-go.
// source: domains.proto
// DO NOT EDIT!

/*
Package domains is a generated protocol buffer package.

It is generated from these files:
	domains.proto

It has these top-level messages:
	LookupRequest
	GetDomainsRequest
	CreateRequest
	DomainResponse
	DomainsResponse
	VerifyRequest
	DomainServiceResponse
	SetServiceRequest
	ServiceRequest
	DomainServicesResponse
*/
package domains

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

type VerificationMode int32

const (
	VerificationMode_DNS    VerificationMode = 0
	VerificationMode_FILE   VerificationMode = 1
	VerificationMode_MANUAL VerificationMode = 2
)

var VerificationMode_name = map[int32]string{
	0: "DNS",
	1: "FILE",
	2: "MANUAL",
}
var VerificationMode_value = map[string]int32{
	"DNS":    0,
	"FILE":   1,
	"MANUAL": 2,
}

func (x VerificationMode) String() string {
	return proto.EnumName(VerificationMode_name, int32(x))
}
func (VerificationMode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DomainServiceType int32

const (
	DomainServiceType_SUPPORT_PORTAL            DomainServiceType = 0
	DomainServiceType_COMMUNICATION_PREFERENCES DomainServiceType = 1
	DomainServiceType_CUSTOMER_PORTAL           DomainServiceType = 2
	DomainServiceType_SHORT_URL                 DomainServiceType = 3
	DomainServiceType_REFERRAL_TRACKING         DomainServiceType = 4
	DomainServiceType_PUBLISHER_PORTAL          DomainServiceType = 5
	DomainServiceType_LIVE_CHAT                 DomainServiceType = 6
)

var DomainServiceType_name = map[int32]string{
	0: "SUPPORT_PORTAL",
	1: "COMMUNICATION_PREFERENCES",
	2: "CUSTOMER_PORTAL",
	3: "SHORT_URL",
	4: "REFERRAL_TRACKING",
	5: "PUBLISHER_PORTAL",
	6: "LIVE_CHAT",
}
var DomainServiceType_value = map[string]int32{
	"SUPPORT_PORTAL":            0,
	"COMMUNICATION_PREFERENCES": 1,
	"CUSTOMER_PORTAL":           2,
	"SHORT_URL":                 3,
	"REFERRAL_TRACKING":         4,
	"PUBLISHER_PORTAL":          5,
	"LIVE_CHAT":                 6,
}

func (x DomainServiceType) String() string {
	return proto.EnumName(DomainServiceType_name, int32(x))
}
func (DomainServiceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type DomainServiceSchema int32

const (
	DomainServiceSchema_HTTP  DomainServiceSchema = 0
	DomainServiceSchema_HTTPS DomainServiceSchema = 1
)

var DomainServiceSchema_name = map[int32]string{
	0: "HTTP",
	1: "HTTPS",
}
var DomainServiceSchema_value = map[string]int32{
	"HTTP":  0,
	"HTTPS": 1,
}

func (x DomainServiceSchema) String() string {
	return proto.EnumName(DomainServiceSchema_name, int32(x))
}
func (DomainServiceSchema) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type LookupRequest struct {
	Domain string `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
}

func (m *LookupRequest) Reset()                    { *m = LookupRequest{} }
func (m *LookupRequest) String() string            { return proto.CompactTextString(m) }
func (*LookupRequest) ProtoMessage()               {}
func (*LookupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetDomainsRequest struct {
	BrandId string `protobuf:"bytes,1,opt,name=brand_id,json=brandId" json:"brand_id,omitempty"`
}

func (m *GetDomainsRequest) Reset()                    { *m = GetDomainsRequest{} }
func (m *GetDomainsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDomainsRequest) ProtoMessage()               {}
func (*GetDomainsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CreateRequest struct {
	Domain  string `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	BrandId string `protobuf:"bytes,2,opt,name=brand_id,json=brandId" json:"brand_id,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type DomainResponse struct {
	Domain           string                     `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	BrandId          string                     `protobuf:"bytes,2,opt,name=brand_id,json=brandId" json:"brand_id,omitempty"`
	Verification     string                     `protobuf:"bytes,3,opt,name=verification" json:"verification,omitempty"`
	VerificationMode string                     `protobuf:"bytes,4,opt,name=verification_mode,json=verificationMode" json:"verification_mode,omitempty"`
	IsVerified       string                     `protobuf:"bytes,5,opt,name=is_verified,json=isVerified" json:"is_verified,omitempty"`
	DateVerified     *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=date_verified,json=dateVerified" json:"date_verified,omitempty"`
	Language         string                     `protobuf:"bytes,7,opt,name=language" json:"language,omitempty"`
	Country          string                     `protobuf:"bytes,8,opt,name=country" json:"country,omitempty"`
}

func (m *DomainResponse) Reset()                    { *m = DomainResponse{} }
func (m *DomainResponse) String() string            { return proto.CompactTextString(m) }
func (*DomainResponse) ProtoMessage()               {}
func (*DomainResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DomainResponse) GetDateVerified() *google_protobuf.Timestamp {
	if m != nil {
		return m.DateVerified
	}
	return nil
}

type DomainsResponse struct {
	Domains map[string]*DomainResponse `protobuf:"bytes,1,rep,name=domains" json:"domains,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *DomainsResponse) Reset()                    { *m = DomainsResponse{} }
func (m *DomainsResponse) String() string            { return proto.CompactTextString(m) }
func (*DomainsResponse) ProtoMessage()               {}
func (*DomainsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DomainsResponse) GetDomains() map[string]*DomainResponse {
	if m != nil {
		return m.Domains
	}
	return nil
}

type VerifyRequest struct {
	Domain string `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
}

func (m *VerifyRequest) Reset()                    { *m = VerifyRequest{} }
func (m *VerifyRequest) String() string            { return proto.CompactTextString(m) }
func (*VerifyRequest) ProtoMessage()               {}
func (*VerifyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type DomainServiceResponse struct {
	Domain         string                `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	Subdomain      string                `protobuf:"bytes,2,opt,name=subdomain" json:"subdomain,omitempty"`
	Type           DomainServiceType     `protobuf:"varint,3,opt,name=type,enum=domains.DomainServiceType" json:"type,omitempty"`
	IsEnabled      bool                  `protobuf:"varint,4,opt,name=is_enabled,json=isEnabled" json:"is_enabled,omitempty"`
	BaseUrl        string                `protobuf:"bytes,5,opt,name=base_url,json=baseUrl" json:"base_url,omitempty"`
	DefaultSchema  DomainServiceSchema   `protobuf:"varint,6,opt,name=default_schema,json=defaultSchema,enum=domains.DomainServiceSchema" json:"default_schema,omitempty"`
	AvailbleSchema []DomainServiceSchema `protobuf:"varint,7,rep,name=availble_schema,json=availbleSchema,enum=domains.DomainServiceSchema" json:"availble_schema,omitempty"`
}

func (m *DomainServiceResponse) Reset()                    { *m = DomainServiceResponse{} }
func (m *DomainServiceResponse) String() string            { return proto.CompactTextString(m) }
func (*DomainServiceResponse) ProtoMessage()               {}
func (*DomainServiceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type SetServiceRequest struct {
	Domain    string            `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	Subdomain string            `protobuf:"bytes,2,opt,name=subdomain" json:"subdomain,omitempty"`
	Type      DomainServiceType `protobuf:"varint,3,opt,name=type,enum=domains.DomainServiceType" json:"type,omitempty"`
}

func (m *SetServiceRequest) Reset()                    { *m = SetServiceRequest{} }
func (m *SetServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*SetServiceRequest) ProtoMessage()               {}
func (*SetServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type ServiceRequest struct {
	Domain string            `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	Type   DomainServiceType `protobuf:"varint,2,opt,name=type,enum=domains.DomainServiceType" json:"type,omitempty"`
}

func (m *ServiceRequest) Reset()                    { *m = ServiceRequest{} }
func (m *ServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*ServiceRequest) ProtoMessage()               {}
func (*ServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type DomainServicesResponse struct {
	Services map[string]*DomainServiceResponse `protobuf:"bytes,1,rep,name=services" json:"services,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *DomainServicesResponse) Reset()                    { *m = DomainServicesResponse{} }
func (m *DomainServicesResponse) String() string            { return proto.CompactTextString(m) }
func (*DomainServicesResponse) ProtoMessage()               {}
func (*DomainServicesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DomainServicesResponse) GetServices() map[string]*DomainServiceResponse {
	if m != nil {
		return m.Services
	}
	return nil
}

func init() {
	proto.RegisterType((*LookupRequest)(nil), "domains.LookupRequest")
	proto.RegisterType((*GetDomainsRequest)(nil), "domains.GetDomainsRequest")
	proto.RegisterType((*CreateRequest)(nil), "domains.CreateRequest")
	proto.RegisterType((*DomainResponse)(nil), "domains.DomainResponse")
	proto.RegisterType((*DomainsResponse)(nil), "domains.DomainsResponse")
	proto.RegisterType((*VerifyRequest)(nil), "domains.VerifyRequest")
	proto.RegisterType((*DomainServiceResponse)(nil), "domains.DomainServiceResponse")
	proto.RegisterType((*SetServiceRequest)(nil), "domains.SetServiceRequest")
	proto.RegisterType((*ServiceRequest)(nil), "domains.ServiceRequest")
	proto.RegisterType((*DomainServicesResponse)(nil), "domains.DomainServicesResponse")
	proto.RegisterEnum("domains.VerificationMode", VerificationMode_name, VerificationMode_value)
	proto.RegisterEnum("domains.DomainServiceType", DomainServiceType_name, DomainServiceType_value)
	proto.RegisterEnum("domains.DomainServiceSchema", DomainServiceSchema_name, DomainServiceSchema_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Domains service

type DomainsClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*DomainResponse, error)
	GetDomains(ctx context.Context, in *GetDomainsRequest, opts ...grpc.CallOption) (*DomainsResponse, error)
	Retrieve(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*DomainResponse, error)
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*DomainResponse, error)
	SetService(ctx context.Context, in *SetServiceRequest, opts ...grpc.CallOption) (*DomainServiceResponse, error)
	RetrieveService(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*DomainServiceResponse, error)
	GetServices(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*DomainServicesResponse, error)
	RemoveService(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*DomainServiceResponse, error)
	DisableService(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*DomainServiceResponse, error)
	EnableService(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*DomainServiceResponse, error)
}

type domainsClient struct {
	cc *grpc.ClientConn
}

func NewDomainsClient(cc *grpc.ClientConn) DomainsClient {
	return &domainsClient{cc}
}

func (c *domainsClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*DomainResponse, error) {
	out := new(DomainResponse)
	err := grpc.Invoke(ctx, "/domains.Domains/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainsClient) GetDomains(ctx context.Context, in *GetDomainsRequest, opts ...grpc.CallOption) (*DomainsResponse, error) {
	out := new(DomainsResponse)
	err := grpc.Invoke(ctx, "/domains.Domains/GetDomains", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainsClient) Retrieve(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*DomainResponse, error) {
	out := new(DomainResponse)
	err := grpc.Invoke(ctx, "/domains.Domains/Retrieve", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainsClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*DomainResponse, error) {
	out := new(DomainResponse)
	err := grpc.Invoke(ctx, "/domains.Domains/Verify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainsClient) SetService(ctx context.Context, in *SetServiceRequest, opts ...grpc.CallOption) (*DomainServiceResponse, error) {
	out := new(DomainServiceResponse)
	err := grpc.Invoke(ctx, "/domains.Domains/SetService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainsClient) RetrieveService(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*DomainServiceResponse, error) {
	out := new(DomainServiceResponse)
	err := grpc.Invoke(ctx, "/domains.Domains/RetrieveService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainsClient) GetServices(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*DomainServicesResponse, error) {
	out := new(DomainServicesResponse)
	err := grpc.Invoke(ctx, "/domains.Domains/GetServices", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainsClient) RemoveService(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*DomainServiceResponse, error) {
	out := new(DomainServiceResponse)
	err := grpc.Invoke(ctx, "/domains.Domains/RemoveService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainsClient) DisableService(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*DomainServiceResponse, error) {
	out := new(DomainServiceResponse)
	err := grpc.Invoke(ctx, "/domains.Domains/DisableService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainsClient) EnableService(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*DomainServiceResponse, error) {
	out := new(DomainServiceResponse)
	err := grpc.Invoke(ctx, "/domains.Domains/EnableService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Domains service

type DomainsServer interface {
	Create(context.Context, *CreateRequest) (*DomainResponse, error)
	GetDomains(context.Context, *GetDomainsRequest) (*DomainsResponse, error)
	Retrieve(context.Context, *LookupRequest) (*DomainResponse, error)
	Verify(context.Context, *VerifyRequest) (*DomainResponse, error)
	SetService(context.Context, *SetServiceRequest) (*DomainServiceResponse, error)
	RetrieveService(context.Context, *ServiceRequest) (*DomainServiceResponse, error)
	GetServices(context.Context, *ServiceRequest) (*DomainServicesResponse, error)
	RemoveService(context.Context, *ServiceRequest) (*DomainServiceResponse, error)
	DisableService(context.Context, *ServiceRequest) (*DomainServiceResponse, error)
	EnableService(context.Context, *ServiceRequest) (*DomainServiceResponse, error)
}

func RegisterDomainsServer(s *grpc.Server, srv DomainsServer) {
	s.RegisterService(&_Domains_serviceDesc, srv)
}

func _Domains_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domains.Domains/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainsServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Domains_GetDomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDomainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainsServer).GetDomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domains.Domains/GetDomains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainsServer).GetDomains(ctx, req.(*GetDomainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Domains_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainsServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domains.Domains/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainsServer).Retrieve(ctx, req.(*LookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Domains_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainsServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domains.Domains/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainsServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Domains_SetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainsServer).SetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domains.Domains/SetService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainsServer).SetService(ctx, req.(*SetServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Domains_RetrieveService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainsServer).RetrieveService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domains.Domains/RetrieveService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainsServer).RetrieveService(ctx, req.(*ServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Domains_GetServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainsServer).GetServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domains.Domains/GetServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainsServer).GetServices(ctx, req.(*ServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Domains_RemoveService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainsServer).RemoveService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domains.Domains/RemoveService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainsServer).RemoveService(ctx, req.(*ServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Domains_DisableService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainsServer).DisableService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domains.Domains/DisableService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainsServer).DisableService(ctx, req.(*ServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Domains_EnableService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainsServer).EnableService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domains.Domains/EnableService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainsServer).EnableService(ctx, req.(*ServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Domains_serviceDesc = grpc.ServiceDesc{
	ServiceName: "domains.Domains",
	HandlerType: (*DomainsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Domains_Create_Handler,
		},
		{
			MethodName: "GetDomains",
			Handler:    _Domains_GetDomains_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _Domains_Retrieve_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Domains_Verify_Handler,
		},
		{
			MethodName: "SetService",
			Handler:    _Domains_SetService_Handler,
		},
		{
			MethodName: "RetrieveService",
			Handler:    _Domains_RetrieveService_Handler,
		},
		{
			MethodName: "GetServices",
			Handler:    _Domains_GetServices_Handler,
		},
		{
			MethodName: "RemoveService",
			Handler:    _Domains_RemoveService_Handler,
		},
		{
			MethodName: "DisableService",
			Handler:    _Domains_DisableService_Handler,
		},
		{
			MethodName: "EnableService",
			Handler:    _Domains_EnableService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("domains.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 873 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x56, 0x5b, 0x8f, 0xda, 0x46,
	0x14, 0x0e, 0xb0, 0x60, 0x38, 0xbb, 0xb0, 0x66, 0xd2, 0x6c, 0x1c, 0x94, 0x76, 0x2b, 0xa4, 0xaa,
	0xd5, 0x56, 0x21, 0x2a, 0xed, 0x43, 0x2f, 0x0f, 0x11, 0xeb, 0x38, 0xe0, 0x96, 0x9b, 0xc6, 0x66,
	0x55, 0xa9, 0x0f, 0x96, 0x81, 0x59, 0x6a, 0xc5, 0x60, 0xea, 0x0b, 0x12, 0x7f, 0xa8, 0x7d, 0xe9,
	0x6f, 0xa8, 0x2a, 0xf5, 0x8f, 0x75, 0x3c, 0xe3, 0x0b, 0xde, 0x85, 0x92, 0x26, 0xfb, 0x82, 0x7c,
	0x6e, 0xdf, 0x39, 0x73, 0xce, 0x37, 0x67, 0x80, 0xea, 0xdc, 0x59, 0x9a, 0xd6, 0xca, 0x6b, 0xad,
	0x5d, 0xc7, 0x77, 0x90, 0x10, 0x89, 0x8d, 0xcb, 0x85, 0xe3, 0x2c, 0x6c, 0xf2, 0x92, 0xa9, 0xa7,
	0xc1, 0xed, 0x4b, 0xdf, 0x5a, 0x12, 0xcf, 0x37, 0x97, 0x6b, 0xee, 0xd9, 0xfc, 0x1c, 0xaa, 0x7d,
	0xc7, 0x79, 0x1b, 0xac, 0x31, 0xf9, 0x2d, 0xa0, 0x16, 0x74, 0x01, 0x25, 0x1e, 0x2c, 0xe5, 0x3e,
	0xcd, 0x7d, 0x51, 0xc1, 0x91, 0xd4, 0x6c, 0x41, 0xbd, 0x4b, 0xfc, 0xd7, 0x1c, 0x37, 0x76, 0x7e,
	0x06, 0xe5, 0xa9, 0x6b, 0xae, 0xe6, 0x86, 0x35, 0x8f, 0xdc, 0x05, 0x26, 0xab, 0xf3, 0xe6, 0x35,
	0x54, 0x65, 0x97, 0x98, 0x3e, 0x39, 0x02, 0x9c, 0xc1, 0xc8, 0x67, 0x31, 0xfe, 0xcc, 0x43, 0x8d,
	0x67, 0xc4, 0xc4, 0x5b, 0x3b, 0x2b, 0x8f, 0xbc, 0x07, 0x0a, 0x6a, 0xc2, 0xd9, 0x86, 0xb8, 0xd6,
	0xad, 0x35, 0x33, 0x7d, 0xcb, 0x59, 0x49, 0x05, 0x66, 0xce, 0xe8, 0xd0, 0x97, 0x50, 0xdf, 0x95,
	0x8d, 0xa5, 0x33, 0x27, 0xd2, 0x09, 0x73, 0x14, 0x77, 0x0d, 0x03, 0xaa, 0x47, 0x97, 0x70, 0x6a,
	0x79, 0x06, 0x57, 0x93, 0xb9, 0x54, 0x64, 0x6e, 0x60, 0x79, 0x37, 0x91, 0x06, 0xbd, 0xa2, 0xf3,
	0xa0, 0x27, 0x4f, 0x5d, 0x4a, 0xd4, 0xe5, 0xb4, 0xdd, 0x68, 0xf1, 0x69, 0xb4, 0xe2, 0x69, 0xb4,
	0xf4, 0x78, 0x1a, 0xf8, 0x2c, 0x0c, 0x48, 0x00, 0x1a, 0x50, 0xb6, 0xcd, 0xd5, 0x22, 0x30, 0x17,
	0x44, 0x12, 0x18, 0x7c, 0x22, 0x23, 0x09, 0x84, 0x99, 0x13, 0xac, 0x7c, 0x77, 0x2b, 0x95, 0xf9,
	0x41, 0x23, 0xb1, 0xf9, 0x47, 0x0e, 0xce, 0x93, 0x01, 0x45, 0xfd, 0x7a, 0x05, 0x31, 0x17, 0x68,
	0xc3, 0x0a, 0xb4, 0x88, 0xcf, 0x5a, 0x31, 0x55, 0xee, 0xb8, 0xc6, 0xb2, 0x12, 0x62, 0xe1, 0x84,
	0x41, 0x1a, 0x9c, 0xed, 0x1a, 0x90, 0x08, 0x85, 0xb7, 0x64, 0x1b, 0x75, 0x3f, 0xfc, 0x44, 0x2f,
	0xa0, 0xb8, 0x31, 0xed, 0x80, 0xb0, 0xbe, 0x9f, 0xb6, 0x9f, 0xde, 0x49, 0x10, 0xe3, 0x63, 0xee,
	0xf5, 0x7d, 0xfe, 0xdb, 0x5c, 0xc8, 0x3a, 0x76, 0xd6, 0xed, 0x31, 0xd6, 0xfd, 0x93, 0x87, 0x27,
	0x1c, 0x46, 0x23, 0xee, 0xc6, 0x9a, 0x91, 0xa3, 0x44, 0x78, 0x0e, 0x15, 0x2f, 0x98, 0x46, 0x26,
	0xce, 0x84, 0x54, 0x81, 0x5a, 0x70, 0xe2, 0x6f, 0xd7, 0x84, 0x71, 0xa0, 0x46, 0x07, 0x92, 0x2d,
	0x35, 0xca, 0xa1, 0x53, 0x0f, 0xcc, 0xfc, 0xd0, 0xc7, 0x40, 0xe7, 0x6a, 0x90, 0x95, 0x39, 0xb5,
	0xe9, 0x18, 0x43, 0x42, 0x94, 0x71, 0xc5, 0xa2, 0xad, 0x60, 0x0a, 0xc6, 0x3a, 0xd3, 0x23, 0x46,
	0xe0, 0xda, 0x11, 0x0d, 0x84, 0x50, 0x9e, 0xb8, 0x36, 0x92, 0xa1, 0x36, 0x27, 0xb7, 0x66, 0x60,
	0xfb, 0x86, 0x37, 0xfb, 0x95, 0x2c, 0x4d, 0x46, 0x82, 0x5a, 0xfb, 0xf9, 0xfe, 0x9c, 0x1a, 0xf3,
	0xc1, 0xd5, 0x28, 0x86, 0x8b, 0x48, 0x81, 0x73, 0x73, 0x63, 0x5a, 0x36, 0x4d, 0x16, 0xa3, 0x08,
	0x74, 0x8a, 0xc7, 0x50, 0x6a, 0x71, 0x10, 0x97, 0x9b, 0x5b, 0xa8, 0x6b, 0xc4, 0x4f, 0x3a, 0xf8,
	0xdf, 0xf7, 0xf1, 0x41, 0x1b, 0xd8, 0xfc, 0x19, 0x6a, 0xef, 0x98, 0x37, 0x46, 0xce, 0xbf, 0x23,
	0xf2, 0xdf, 0x39, 0xb8, 0xc8, 0xd8, 0x52, 0xd2, 0xab, 0x50, 0xf6, 0x22, 0x5d, 0xc4, 0xfa, 0x17,
	0xfb, 0xe1, 0x52, 0xf2, 0xc7, 0x0a, 0xce, 0xfe, 0x24, 0xbc, 0xf1, 0x0b, 0x54, 0x33, 0xa6, 0x3d,
	0xfc, 0xff, 0x26, 0xcb, 0xff, 0x4f, 0xf6, 0xa7, 0xda, 0x73, 0x0d, 0xae, 0xbe, 0x02, 0xf1, 0xe6,
	0xee, 0x72, 0x11, 0xa0, 0xf0, 0x7a, 0xa8, 0x89, 0x8f, 0x50, 0x19, 0x4e, 0xde, 0xa8, 0x7d, 0x45,
	0xcc, 0x21, 0x80, 0xd2, 0xa0, 0x33, 0x9c, 0x74, 0xfa, 0x62, 0xfe, 0xea, 0xf7, 0x1c, 0xd4, 0xef,
	0x75, 0x04, 0x21, 0xda, 0xe5, 0xc9, 0x78, 0x3c, 0xc2, 0xba, 0x11, 0xfe, 0x50, 0xcf, 0x47, 0x94,
	0xba, 0xcf, 0xe4, 0xd1, 0x60, 0x30, 0x19, 0xaa, 0x72, 0x47, 0x57, 0x47, 0x43, 0x63, 0x8c, 0x95,
	0x37, 0x0a, 0x56, 0x86, 0xb2, 0xa2, 0x51, 0xd0, 0xc7, 0x70, 0x2e, 0x4f, 0x34, 0x7d, 0x34, 0x50,
	0x70, 0x1c, 0x93, 0x47, 0x55, 0xa8, 0x68, 0xbd, 0x10, 0x65, 0x82, 0xfb, 0x62, 0x01, 0x3d, 0x81,
	0x3a, 0x8b, 0xc1, 0x9d, 0xbe, 0xa1, 0xe3, 0x8e, 0xfc, 0x93, 0x3a, 0xec, 0x8a, 0x27, 0xe8, 0x23,
	0x10, 0xc7, 0x93, 0xeb, 0xbe, 0xaa, 0xf5, 0xd2, 0xd8, 0x62, 0x18, 0xdb, 0x57, 0x6f, 0x14, 0x43,
	0xee, 0x75, 0x74, 0xb1, 0x74, 0x75, 0x05, 0x8f, 0xf7, 0x50, 0x33, 0x3c, 0x55, 0x4f, 0xd7, 0xc7,
	0xb4, 0xbe, 0x0a, 0x14, 0xc3, 0x2f, 0x5a, 0x4b, 0xfb, 0xaf, 0x22, 0x08, 0xd1, 0x92, 0x41, 0xdf,
	0x41, 0x89, 0xbf, 0x1b, 0xe8, 0x22, 0x69, 0x64, 0xe6, 0x21, 0x69, 0x1c, 0x5a, 0x30, 0xe8, 0x1a,
	0x20, 0x7d, 0xa2, 0x50, 0xca, 0xa0, 0x7b, 0xef, 0x56, 0x43, 0x3a, 0xb4, 0x04, 0xd1, 0x0f, 0x50,
	0xc6, 0xc4, 0x77, 0x2d, 0xb2, 0xd9, 0x2d, 0x20, 0xf3, 0x44, 0x1e, 0x2e, 0x80, 0xd6, 0xce, 0xd7,
	0xda, 0x4e, 0x68, 0x66, 0xcf, 0x1d, 0x0e, 0xed, 0x01, 0xa4, 0x57, 0x74, 0xa7, 0xf6, 0x7b, 0xf7,
	0xb6, 0x71, 0x84, 0x5f, 0xe8, 0x47, 0x38, 0x8f, 0x4f, 0x10, 0xc3, 0x3d, 0xdd, 0x81, 0xfb, 0x5f,
	0x58, 0x5d, 0x38, 0xed, 0x26, 0x05, 0x78, 0x87, 0x71, 0x2e, 0x8f, 0x5c, 0x2f, 0x7a, 0xbc, 0x2a,
	0x26, 0x4b, 0xe7, 0x01, 0x4a, 0x52, 0xe9, 0x5f, 0x02, 0xcb, 0x0b, 0xd7, 0xef, 0x07, 0x43, 0xd1,
	0xa2, 0xf8, 0x22, 0xff, 0x50, 0xa4, 0x69, 0x89, 0xbd, 0xe8, 0x5f, 0xff, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0x0a, 0x09, 0xf9, 0xff, 0x87, 0x09, 0x00, 0x00,
}
