// Code generated by protoc-gen-go.
// source: brands.proto
// DO NOT EDIT!

/*
Package brands is a generated protocol buffer package.

It is generated from these files:
	brands.proto

It has these top-level messages:
	CreateRequest
	ListRequest
	BrandsResponse
*/
package brands

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import fcommon "github.com/fortifi/proto-go/fcommon"
import fortifi_ftypes "github.com/fortifi/proto-go/ftypes/fortifi"

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
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
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

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type BrandsResponse struct {
	Brands map[string]*fortifi_ftypes.Brand `protobuf:"bytes,1,rep,name=brands" json:"brands,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *BrandsResponse) Reset()                    { *m = BrandsResponse{} }
func (m *BrandsResponse) String() string            { return proto.CompactTextString(m) }
func (*BrandsResponse) ProtoMessage()               {}
func (*BrandsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BrandsResponse) GetBrands() map[string]*fortifi_ftypes.Brand {
	if m != nil {
		return m.Brands
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "brands.CreateRequest")
	proto.RegisterType((*ListRequest)(nil), "brands.ListRequest")
	proto.RegisterType((*BrandsResponse)(nil), "brands.BrandsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Brands service

type BrandsClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*fortifi_ftypes.Brand, error)
	Retrieve(ctx context.Context, in *fcommon.IDRequest, opts ...grpc.CallOption) (*fortifi_ftypes.Brand, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*BrandsResponse, error)
	Archive(ctx context.Context, in *fcommon.IDRequest, opts ...grpc.CallOption) (*fcommon.BoolResponse, error)
}

type brandsClient struct {
	cc *grpc.ClientConn
}

func NewBrandsClient(cc *grpc.ClientConn) BrandsClient {
	return &brandsClient{cc}
}

func (c *brandsClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*fortifi_ftypes.Brand, error) {
	out := new(fortifi_ftypes.Brand)
	err := grpc.Invoke(ctx, "/brands.Brands/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandsClient) Retrieve(ctx context.Context, in *fcommon.IDRequest, opts ...grpc.CallOption) (*fortifi_ftypes.Brand, error) {
	out := new(fortifi_ftypes.Brand)
	err := grpc.Invoke(ctx, "/brands.Brands/Retrieve", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandsClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*BrandsResponse, error) {
	out := new(BrandsResponse)
	err := grpc.Invoke(ctx, "/brands.Brands/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandsClient) Archive(ctx context.Context, in *fcommon.IDRequest, opts ...grpc.CallOption) (*fcommon.BoolResponse, error) {
	out := new(fcommon.BoolResponse)
	err := grpc.Invoke(ctx, "/brands.Brands/Archive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Brands service

type BrandsServer interface {
	Create(context.Context, *CreateRequest) (*fortifi_ftypes.Brand, error)
	Retrieve(context.Context, *fcommon.IDRequest) (*fortifi_ftypes.Brand, error)
	List(context.Context, *ListRequest) (*BrandsResponse, error)
	Archive(context.Context, *fcommon.IDRequest) (*fcommon.BoolResponse, error)
}

func RegisterBrandsServer(s *grpc.Server, srv BrandsServer) {
	s.RegisterService(&_Brands_serviceDesc, srv)
}

func _Brands_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brands.Brands/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brands_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fcommon.IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brands.Brands/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).Retrieve(ctx, req.(*fcommon.IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brands_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brands.Brands/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brands_Archive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fcommon.IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).Archive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brands.Brands/Archive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).Archive(ctx, req.(*fcommon.IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Brands_serviceDesc = grpc.ServiceDesc{
	ServiceName: "brands.Brands",
	HandlerType: (*BrandsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Brands_Create_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _Brands_Retrieve_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Brands_List_Handler,
		},
		{
			MethodName: "Archive",
			Handler:    _Brands_Archive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "brands.proto",
}

func init() { proto.RegisterFile("brands.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x51, 0xdf, 0x4b, 0x3a, 0x41,
	0x10, 0x67, 0x4f, 0xbf, 0xf7, 0xad, 0xb9, 0x94, 0x98, 0x32, 0xe4, 0x9e, 0xe4, 0xe8, 0x41, 0x88,
	0x4e, 0xd0, 0x8a, 0xf0, 0x2d, 0xab, 0x87, 0xa0, 0x87, 0xd8, 0xc7, 0xde, 0x4e, 0x9d, 0xd3, 0x25,
	0xef, 0xd6, 0x76, 0x57, 0xc1, 0x7f, 0xa6, 0x3f, 0xae, 0xbf, 0x24, 0xbc, 0xdd, 0x05, 0x05, 0xa3,
	0xb7, 0x99, 0x0f, 0x33, 0xf3, 0xf9, 0x31, 0x70, 0x32, 0x56, 0x59, 0x39, 0xd5, 0xe9, 0x52, 0x49,
	0x23, 0x31, 0xb4, 0x5d, 0xdc, 0xc8, 0x27, 0xb2, 0x28, 0x64, 0x69, 0xe1, 0xf8, 0x3c, 0x37, 0x9b,
	0x25, 0xe9, 0x5e, 0x2e, 0x95, 0x11, 0xb9, 0xb0, 0x68, 0x32, 0x80, 0xc6, 0xa3, 0xa2, 0xcc, 0x10,
	0xa7, 0xcf, 0x15, 0x69, 0x83, 0x4d, 0x08, 0xc4, 0xb4, 0xcd, 0x3a, 0xac, 0x7b, 0xcc, 0x03, 0x31,
	0x45, 0x84, 0x7a, 0x99, 0x15, 0xd4, 0x0e, 0x2a, 0xa4, 0xaa, 0x93, 0x06, 0x44, 0xaf, 0x42, 0x1b,
	0xb7, 0x92, 0x7c, 0x31, 0x68, 0x8e, 0x2a, 0x4e, 0x4e, 0x7a, 0x29, 0x4b, 0x4d, 0x38, 0x04, 0xa7,
	0xa2, 0xcd, 0x3a, 0xb5, 0x6e, 0xd4, 0x4f, 0x52, 0x27, 0x71, 0x7f, 0xce, 0xb5, 0xcf, 0xa5, 0x51,
	0x1b, 0xee, 0x75, 0xbf, 0x41, 0xb4, 0x03, 0xe3, 0x29, 0xd4, 0x3e, 0x68, 0xe3, 0x14, 0x6d, 0x4b,
	0xbc, 0x82, 0x7f, 0xeb, 0x6c, 0xb1, 0xb2, 0x9a, 0xa2, 0x7e, 0x2b, 0xf5, 0x96, 0xac, 0x43, 0x7b,
	0x94, 0xdb, 0x99, 0x61, 0x70, 0xcf, 0xfa, 0xdf, 0x0c, 0x42, 0x7b, 0x12, 0xef, 0x20, 0xb4, 0x7e,
	0xb1, 0xe5, 0x25, 0xed, 0xf9, 0x8f, 0x0f, 0x5f, 0xc3, 0x5b, 0x38, 0xe2, 0x64, 0x94, 0xa0, 0x35,
	0x21, 0xa6, 0x3e, 0xd9, 0x97, 0xa7, 0x3f, 0xd6, 0x06, 0x50, 0xdf, 0x26, 0x85, 0x67, 0x9e, 0x6c,
	0x27, 0xb7, 0xf8, 0xe2, 0x70, 0x28, 0x78, 0x03, 0xff, 0x1f, 0xd4, 0x64, 0x2e, 0x7e, 0xa7, 0x72,
	0xd8, 0x48, 0xca, 0x85, 0xdf, 0x1a, 0x5d, 0xbe, 0x27, 0x33, 0x61, 0xe6, 0xab, 0x71, 0x3a, 0x91,
	0x85, 0xff, 0x72, 0xaf, 0xfa, 0xf2, 0xf5, 0x4c, 0xf6, 0x2c, 0xd5, 0x38, 0xac, 0x80, 0xc1, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x1e, 0xee, 0x53, 0x33, 0x02, 0x00, 0x00,
}