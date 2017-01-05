// Code generated by protoc-gen-go.
// source: departments.proto
// DO NOT EDIT!

/*
Package departments is a generated protocol buffer package.

It is generated from these files:
	departments.proto

It has these top-level messages:
	CreateRequest
	ListRequest
	DepartmentResponse
	DepartmentsResponse
*/
package departments

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import fcommon "github.com/fortifi/proto-go/fcommon"

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

type CreateRequest struct {
	Id             string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	OrganisationId string `protobuf:"bytes,2,opt,name=organisation_id,json=organisationId" json:"organisation_id,omitempty"`
	Name           string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateRequest) GetOrganisationId() string {
	if m != nil {
		return m.OrganisationId
	}
	return ""
}

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListRequest struct {
	OrganisationId string `protobuf:"bytes,1,opt,name=organisation_id,json=organisationId" json:"organisation_id,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListRequest) GetOrganisationId() string {
	if m != nil {
		return m.OrganisationId
	}
	return ""
}

type DepartmentResponse struct {
	Id             string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	OrganisationId string `protobuf:"bytes,2,opt,name=organisation_id,json=organisationId" json:"organisation_id,omitempty"`
	Name           string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *DepartmentResponse) Reset()                    { *m = DepartmentResponse{} }
func (m *DepartmentResponse) String() string            { return proto.CompactTextString(m) }
func (*DepartmentResponse) ProtoMessage()               {}
func (*DepartmentResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DepartmentResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DepartmentResponse) GetOrganisationId() string {
	if m != nil {
		return m.OrganisationId
	}
	return ""
}

func (m *DepartmentResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DepartmentsResponse struct {
	Departments map[string]*DepartmentResponse `protobuf:"bytes,1,rep,name=departments" json:"departments,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *DepartmentsResponse) Reset()                    { *m = DepartmentsResponse{} }
func (m *DepartmentsResponse) String() string            { return proto.CompactTextString(m) }
func (*DepartmentsResponse) ProtoMessage()               {}
func (*DepartmentsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DepartmentsResponse) GetDepartments() map[string]*DepartmentResponse {
	if m != nil {
		return m.Departments
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "departments.CreateRequest")
	proto.RegisterType((*ListRequest)(nil), "departments.ListRequest")
	proto.RegisterType((*DepartmentResponse)(nil), "departments.DepartmentResponse")
	proto.RegisterType((*DepartmentsResponse)(nil), "departments.DepartmentsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Departments service

type DepartmentsClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*DepartmentResponse, error)
	Retrieve(ctx context.Context, in *fcommon.IDRequest, opts ...grpc.CallOption) (*DepartmentResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*DepartmentsResponse, error)
	Archive(ctx context.Context, in *fcommon.IDRequest, opts ...grpc.CallOption) (*fcommon.BoolResponse, error)
}

type departmentsClient struct {
	cc *grpc.ClientConn
}

func NewDepartmentsClient(cc *grpc.ClientConn) DepartmentsClient {
	return &departmentsClient{cc}
}

func (c *departmentsClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*DepartmentResponse, error) {
	out := new(DepartmentResponse)
	err := grpc.Invoke(ctx, "/departments.Departments/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentsClient) Retrieve(ctx context.Context, in *fcommon.IDRequest, opts ...grpc.CallOption) (*DepartmentResponse, error) {
	out := new(DepartmentResponse)
	err := grpc.Invoke(ctx, "/departments.Departments/Retrieve", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentsClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*DepartmentsResponse, error) {
	out := new(DepartmentsResponse)
	err := grpc.Invoke(ctx, "/departments.Departments/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentsClient) Archive(ctx context.Context, in *fcommon.IDRequest, opts ...grpc.CallOption) (*fcommon.BoolResponse, error) {
	out := new(fcommon.BoolResponse)
	err := grpc.Invoke(ctx, "/departments.Departments/Archive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Departments service

type DepartmentsServer interface {
	Create(context.Context, *CreateRequest) (*DepartmentResponse, error)
	Retrieve(context.Context, *fcommon.IDRequest) (*DepartmentResponse, error)
	List(context.Context, *ListRequest) (*DepartmentsResponse, error)
	Archive(context.Context, *fcommon.IDRequest) (*fcommon.BoolResponse, error)
}

func RegisterDepartmentsServer(s *grpc.Server, srv DepartmentsServer) {
	s.RegisterService(&_Departments_serviceDesc, srv)
}

func _Departments_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/departments.Departments/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentsServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Departments_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fcommon.IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentsServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/departments.Departments/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentsServer).Retrieve(ctx, req.(*fcommon.IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Departments_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/departments.Departments/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentsServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Departments_Archive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fcommon.IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentsServer).Archive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/departments.Departments/Archive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentsServer).Archive(ctx, req.(*fcommon.IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Departments_serviceDesc = grpc.ServiceDesc{
	ServiceName: "departments.Departments",
	HandlerType: (*DepartmentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Departments_Create_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _Departments_Retrieve_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Departments_List_Handler,
		},
		{
			MethodName: "Archive",
			Handler:    _Departments_Archive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "departments.proto",
}

func init() { proto.RegisterFile("departments.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x52, 0x4f, 0x4f, 0xfa, 0x40,
	0x10, 0x4d, 0x0b, 0x3f, 0x7e, 0x3a, 0x0d, 0x88, 0x63, 0x4c, 0x9a, 0x5e, 0x24, 0x5c, 0xc0, 0x83,
	0x25, 0xe2, 0x9f, 0x18, 0x2f, 0x46, 0x84, 0x03, 0x89, 0xa7, 0x7a, 0x33, 0x26, 0xa4, 0xd0, 0x01,
	0x36, 0xd2, 0x5d, 0xdc, 0x5d, 0x48, 0xf8, 0x14, 0x7e, 0x30, 0xbf, 0x94, 0xa1, 0x58, 0xdc, 0x1a,
	0x08, 0x1e, 0xbc, 0x6d, 0xde, 0xcc, 0xbe, 0x37, 0xf3, 0xe6, 0xc1, 0x61, 0x44, 0xd3, 0x50, 0xea,
	0x98, 0xb8, 0x56, 0xfe, 0x54, 0x0a, 0x2d, 0xd0, 0x31, 0x20, 0xaf, 0x38, 0x1c, 0x88, 0x38, 0x16,
	0x7c, 0x55, 0xab, 0xbe, 0x40, 0xf1, 0x41, 0x52, 0xa8, 0x29, 0xa0, 0xb7, 0x19, 0x29, 0x8d, 0x25,
	0xb0, 0x59, 0xe4, 0x5a, 0x15, 0xab, 0xbe, 0x1f, 0xd8, 0x2c, 0xc2, 0x1a, 0x1c, 0x08, 0x39, 0x0a,
	0x39, 0x53, 0xa1, 0x66, 0x82, 0xf7, 0x58, 0xe4, 0xda, 0x49, 0xb1, 0x64, 0xc2, 0xdd, 0x08, 0x11,
	0xf2, 0x3c, 0x8c, 0xc9, 0xcd, 0x25, 0xd5, 0xe4, 0x5d, 0xbd, 0x06, 0xe7, 0x91, 0x29, 0x9d, 0x72,
	0x6f, 0xe0, 0xb2, 0x36, 0x71, 0x55, 0x43, 0xc0, 0xf6, 0x7a, 0xe6, 0x80, 0xd4, 0x54, 0x70, 0x45,
	0x7f, 0x3b, 0xda, 0x87, 0x05, 0x47, 0xdf, 0x1a, 0x6a, 0x2d, 0xf2, 0x04, 0xa6, 0x5d, 0xae, 0x55,
	0xc9, 0xd5, 0x9d, 0xe6, 0xb9, 0x6f, 0xba, 0xba, 0xe1, 0x9b, 0x89, 0x75, 0xb8, 0x96, 0x8b, 0x20,
	0x63, 0x7a, 0x0f, 0xca, 0x3f, 0x1b, 0xb0, 0x0c, 0xb9, 0x57, 0x5a, 0x7c, 0xad, 0xb3, 0x7c, 0xe2,
	0x15, 0xfc, 0x9b, 0x87, 0x93, 0x19, 0x25, 0x5b, 0x38, 0xcd, 0x93, 0x2d, 0xa2, 0xa9, 0x66, 0xb0,
	0xea, 0xbe, 0xb5, 0x6f, 0xac, 0xe6, 0xbb, 0x0d, 0x8e, 0xa1, 0x80, 0x1d, 0x28, 0xac, 0xce, 0x8a,
	0x5e, 0x86, 0x25, 0x73, 0x6b, 0x6f, 0x97, 0x02, 0xde, 0xc1, 0x5e, 0x40, 0x5a, 0x32, 0x9a, 0x13,
	0xa2, 0x9f, 0x26, 0xa7, 0xdb, 0xfe, 0x35, 0x41, 0x0b, 0xf2, 0xcb, 0x00, 0xa0, 0x9b, 0x69, 0x34,
	0x32, 0xe1, 0x55, 0x76, 0x59, 0x8b, 0x97, 0xf0, 0xff, 0x5e, 0x0e, 0xc6, 0x6c, 0xcb, 0x0c, 0xc7,
	0x6b, 0xac, 0x25, 0xc4, 0x24, 0xfd, 0xd5, 0x3a, 0x7d, 0xae, 0x8d, 0x98, 0x1e, 0xcf, 0xfa, 0xfe,
	0x40, 0xc4, 0x8d, 0xa1, 0x90, 0x9a, 0x0d, 0x59, 0x23, 0x09, 0xfd, 0xd9, 0x48, 0x34, 0x0c, 0xd1,
	0x7e, 0x21, 0x41, 0x2f, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xdc, 0x43, 0x51, 0x3b, 0x03,
	0x00, 0x00,
}
