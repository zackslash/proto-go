// Code generated by protoc-gen-go.
// source: metaroute.proto
// DO NOT EDIT!

/*
Package metaroute is a generated protocol buffer package.

It is generated from these files:
	metaroute.proto

It has these top-level messages:
	HTTPRequest
	HTTPResponse
*/
package metaroute

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

type HTTPRequest struct {
	Subdomain   string                                `protobuf:"bytes,1,opt,name=subdomain" json:"subdomain,omitempty"`
	Domain      string                                `protobuf:"bytes,2,opt,name=domain" json:"domain,omitempty"`
	Path        string                                `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	Method      string                                `protobuf:"bytes,4,opt,name=method" json:"method,omitempty"`
	PostParams  map[string]*HTTPRequest_HTTPParameter `protobuf:"bytes,5,rep,name=post_params,json=postParams" json:"post_params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	QueryParams map[string]*HTTPRequest_HTTPParameter `protobuf:"bytes,6,rep,name=query_params,json=queryParams" json:"query_params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Headers     map[string]*HTTPRequest_HTTPParameter `protobuf:"bytes,7,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body        string                                `protobuf:"bytes,8,opt,name=body" json:"body,omitempty"`
	Language    string                                `protobuf:"bytes,9,opt,name=language" json:"language,omitempty"`
	ContentType string                                `protobuf:"bytes,10,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	Customer    *HTTPRequest_CustomerInformation      `protobuf:"bytes,11,opt,name=customer" json:"customer,omitempty"`
}

func (m *HTTPRequest) Reset()                    { *m = HTTPRequest{} }
func (m *HTTPRequest) String() string            { return proto.CompactTextString(m) }
func (*HTTPRequest) ProtoMessage()               {}
func (*HTTPRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HTTPRequest) GetPostParams() map[string]*HTTPRequest_HTTPParameter {
	if m != nil {
		return m.PostParams
	}
	return nil
}

func (m *HTTPRequest) GetQueryParams() map[string]*HTTPRequest_HTTPParameter {
	if m != nil {
		return m.QueryParams
	}
	return nil
}

func (m *HTTPRequest) GetHeaders() map[string]*HTTPRequest_HTTPParameter {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTTPRequest) GetCustomer() *HTTPRequest_CustomerInformation {
	if m != nil {
		return m.Customer
	}
	return nil
}

type HTTPRequest_HTTPParameter struct {
	Values []string `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *HTTPRequest_HTTPParameter) Reset()                    { *m = HTTPRequest_HTTPParameter{} }
func (m *HTTPRequest_HTTPParameter) String() string            { return proto.CompactTextString(m) }
func (*HTTPRequest_HTTPParameter) ProtoMessage()               {}
func (*HTTPRequest_HTTPParameter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type HTTPRequest_CustomerInformation struct {
	BrandId      string            `protobuf:"bytes,1,opt,name=brand_id,json=brandId" json:"brand_id,omitempty"`
	CustomerFid  string            `protobuf:"bytes,2,opt,name=customer_fid,json=customerFid" json:"customer_fid,omitempty"`
	CustomerName string            `protobuf:"bytes,3,opt,name=customer_name,json=customerName" json:"customer_name,omitempty"`
	Meta         map[string]string `protobuf:"bytes,4,rep,name=meta" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *HTTPRequest_CustomerInformation) Reset()         { *m = HTTPRequest_CustomerInformation{} }
func (m *HTTPRequest_CustomerInformation) String() string { return proto.CompactTextString(m) }
func (*HTTPRequest_CustomerInformation) ProtoMessage()    {}
func (*HTTPRequest_CustomerInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 4}
}

func (m *HTTPRequest_CustomerInformation) GetMeta() map[string]string {
	if m != nil {
		return m.Meta
	}
	return nil
}

type HTTPResponse struct {
	StatusCode  int32                                        `protobuf:"varint,1,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	Body        string                                       `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
	ContentType string                                       `protobuf:"bytes,3,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	Language    string                                       `protobuf:"bytes,4,opt,name=language" json:"language,omitempty"`
	Headers     map[string]*HTTPResponse_HTTPHeaderParameter `protobuf:"bytes,5,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *HTTPResponse) Reset()                    { *m = HTTPResponse{} }
func (m *HTTPResponse) String() string            { return proto.CompactTextString(m) }
func (*HTTPResponse) ProtoMessage()               {}
func (*HTTPResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HTTPResponse) GetHeaders() map[string]*HTTPResponse_HTTPHeaderParameter {
	if m != nil {
		return m.Headers
	}
	return nil
}

type HTTPResponse_HTTPHeaderParameter struct {
	Values []string `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *HTTPResponse_HTTPHeaderParameter) Reset()         { *m = HTTPResponse_HTTPHeaderParameter{} }
func (m *HTTPResponse_HTTPHeaderParameter) String() string { return proto.CompactTextString(m) }
func (*HTTPResponse_HTTPHeaderParameter) ProtoMessage()    {}
func (*HTTPResponse_HTTPHeaderParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

func init() {
	proto.RegisterType((*HTTPRequest)(nil), "metaroute.HTTPRequest")
	proto.RegisterType((*HTTPRequest_HTTPParameter)(nil), "metaroute.HTTPRequest.HTTPParameter")
	proto.RegisterType((*HTTPRequest_CustomerInformation)(nil), "metaroute.HTTPRequest.CustomerInformation")
	proto.RegisterType((*HTTPResponse)(nil), "metaroute.HTTPResponse")
	proto.RegisterType((*HTTPResponse_HTTPHeaderParameter)(nil), "metaroute.HTTPResponse.HTTPHeaderParameter")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for FortifiPlatform service

type FortifiPlatformClient interface {
	HandleHTTPRequest(ctx context.Context, in *HTTPRequest, opts ...grpc.CallOption) (*HTTPResponse, error)
}

type fortifiPlatformClient struct {
	cc *grpc.ClientConn
}

func NewFortifiPlatformClient(cc *grpc.ClientConn) FortifiPlatformClient {
	return &fortifiPlatformClient{cc}
}

func (c *fortifiPlatformClient) HandleHTTPRequest(ctx context.Context, in *HTTPRequest, opts ...grpc.CallOption) (*HTTPResponse, error) {
	out := new(HTTPResponse)
	err := grpc.Invoke(ctx, "/metaroute.FortifiPlatform/HandleHTTPRequest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FortifiPlatform service

type FortifiPlatformServer interface {
	HandleHTTPRequest(context.Context, *HTTPRequest) (*HTTPResponse, error)
}

func RegisterFortifiPlatformServer(s *grpc.Server, srv FortifiPlatformServer) {
	s.RegisterService(&_FortifiPlatform_serviceDesc, srv)
}

func _FortifiPlatform_HandleHTTPRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FortifiPlatformServer).HandleHTTPRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metaroute.FortifiPlatform/HandleHTTPRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FortifiPlatformServer).HandleHTTPRequest(ctx, req.(*HTTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FortifiPlatform_serviceDesc = grpc.ServiceDesc{
	ServiceName: "metaroute.FortifiPlatform",
	HandlerType: (*FortifiPlatformServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleHTTPRequest",
			Handler:    _FortifiPlatform_HandleHTTPRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("metaroute.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x55, 0x6d, 0x6b, 0xd3, 0x50,
	0x14, 0xa6, 0x2f, 0x6b, 0x9b, 0x93, 0x8e, 0xce, 0x3b, 0x99, 0x31, 0x08, 0xce, 0x4d, 0xed, 0x50,
	0xd6, 0x42, 0x15, 0x94, 0x81, 0x82, 0x4e, 0x6b, 0x27, 0x28, 0x35, 0x14, 0x04, 0xbf, 0xd4, 0xdb,
	0xe6, 0xb6, 0x0d, 0x36, 0xb9, 0xd9, 0xcd, 0x8d, 0x90, 0xbf, 0xe5, 0xcf, 0xf1, 0xaf, 0xf8, 0xc5,
	0xfb, 0x92, 0xa5, 0xb1, 0x6d, 0x98, 0x5f, 0xf6, 0xed, 0xde, 0xe7, 0x3c, 0xe7, 0xe5, 0x9e, 0x73,
	0x9e, 0x04, 0x5a, 0x3e, 0xe1, 0x98, 0xd1, 0x98, 0x93, 0x4e, 0xc8, 0x28, 0xa7, 0xc8, 0xc8, 0x80,
	0xa3, 0x5f, 0x0d, 0x30, 0x07, 0xa3, 0xd1, 0xd0, 0x21, 0x97, 0x31, 0x89, 0x38, 0xba, 0x07, 0x46,
	0x14, 0x4f, 0x5c, 0xea, 0x63, 0x2f, 0xb0, 0x4a, 0x87, 0xa5, 0x13, 0xc3, 0x59, 0x01, 0xe8, 0x00,
	0x6a, 0xa9, 0xa9, 0xac, 0x4c, 0xe9, 0x0d, 0x21, 0xa8, 0x86, 0x98, 0x2f, 0xac, 0x8a, 0x42, 0xd5,
	0x59, 0x72, 0x45, 0x9a, 0x05, 0x75, 0xad, 0xaa, 0xe6, 0xea, 0x1b, 0xfa, 0x00, 0x66, 0x48, 0x23,
	0x3e, 0x0e, 0x31, 0xc3, 0x7e, 0x64, 0xed, 0x1c, 0x56, 0x4e, 0xcc, 0xde, 0xe3, 0xce, 0xaa, 0xc6,
	0x5c, 0x39, 0x9d, 0xa1, 0x60, 0x0e, 0x15, 0xf1, 0x7d, 0xc0, 0x59, 0xe2, 0x40, 0x98, 0x01, 0xe8,
	0x23, 0x34, 0x05, 0x89, 0x25, 0x57, 0x91, 0x6a, 0x2a, 0x52, 0xbb, 0x20, 0xd2, 0x17, 0x49, 0xcd,
	0x87, 0x32, 0x2f, 0x57, 0x08, 0x7a, 0x05, 0xf5, 0x05, 0xc1, 0x2e, 0x61, 0x91, 0x55, 0x57, 0x61,
	0x8e, 0x0b, 0xc2, 0x0c, 0x34, 0x4b, 0x87, 0xb8, 0xf2, 0x91, 0xef, 0x9f, 0x50, 0x37, 0xb1, 0x1a,
	0xfa, 0xfd, 0xf2, 0x8c, 0x6c, 0x68, 0x2c, 0x71, 0x30, 0x8f, 0xf1, 0x9c, 0x58, 0x86, 0xc2, 0xb3,
	0x3b, 0x7a, 0x00, 0xcd, 0x29, 0x0d, 0x38, 0x09, 0xf8, 0x98, 0x27, 0x21, 0xb1, 0x40, 0xd9, 0xcd,
	0x14, 0x1b, 0x09, 0x08, 0xf5, 0xa1, 0x31, 0x8d, 0x23, 0x4e, 0x7d, 0xc2, 0x2c, 0x53, 0x98, 0xcd,
	0xde, 0x93, 0x82, 0x92, 0xce, 0x53, 0xda, 0x45, 0x30, 0xa3, 0xcc, 0xc7, 0xdc, 0xa3, 0x81, 0x93,
	0xf9, 0xda, 0x6d, 0xd8, 0x95, 0x64, 0xf5, 0x4e, 0xc2, 0x09, 0x93, 0x73, 0xf9, 0x89, 0x97, 0xc2,
	0x53, 0x8c, 0xb7, 0x22, 0xe7, 0xa2, 0x6f, 0xf6, 0x14, 0x5a, 0x6b, 0xdd, 0x46, 0x7b, 0x50, 0xf9,
	0x41, 0x92, 0x74, 0x0d, 0xe4, 0x11, 0x9d, 0xc1, 0x8e, 0xa2, 0xab, 0xf9, 0x9b, 0xbd, 0x87, 0x45,
	0x5d, 0xca, 0x67, 0x74, 0xb4, 0xcb, 0x59, 0xf9, 0x65, 0xc9, 0x76, 0x61, 0x6f, 0x7d, 0x10, 0x37,
	0x90, 0xe5, 0x3b, 0x34, 0xf3, 0x73, 0xba, 0x81, 0x0c, 0x7f, 0x4a, 0xb0, 0xbf, 0xa5, 0xef, 0xe8,
	0x2e, 0x34, 0x26, 0x0c, 0x07, 0xee, 0xd8, 0x73, 0xd3, 0x74, 0x75, 0x75, 0xbf, 0x70, 0xd5, 0xcc,
	0x53, 0x8f, 0xf1, 0x4c, 0x98, 0xcb, 0xe9, 0xcc, 0x53, 0xac, 0xef, 0xb9, 0xe8, 0x18, 0x76, 0x33,
	0x4a, 0x20, 0x92, 0xa6, 0x7a, 0xca, 0xfc, 0x3e, 0x0b, 0x0c, 0x0d, 0xa0, 0x2a, 0x8b, 0x15, 0xaa,
	0x92, 0x7b, 0xfa, 0xfc, 0xff, 0x97, 0xa2, 0xf3, 0x49, 0x30, 0xf5, 0xe2, 0xaa, 0x08, 0xf6, 0x0b,
	0x30, 0x32, 0x68, 0x4b, 0x8f, 0x6e, 0xe7, 0x7b, 0x64, 0xe4, 0x5e, 0x7f, 0xf4, 0xbb, 0x2c, 0x1a,
	0xac, 0x92, 0x45, 0x21, 0x0d, 0x22, 0x82, 0xee, 0x83, 0x19, 0x71, 0xcc, 0xe3, 0x68, 0x3c, 0xa5,
	0x2e, 0x51, 0x41, 0x76, 0x1c, 0xd0, 0xd0, 0xb9, 0x40, 0x32, 0x81, 0x94, 0x73, 0x02, 0x59, 0x17,
	0x41, 0x65, 0x53, 0x04, 0x79, 0x0d, 0x55, 0xd7, 0x34, 0xf4, 0x7a, 0x25, 0x59, 0xfd, 0x0d, 0xd9,
	0x1c, 0xa2, 0xae, 0x6e, 0xbb, 0x66, 0xed, 0x53, 0xd8, 0x97, 0x2c, 0x6d, 0xbc, 0x5e, 0x1e, 0xf3,
	0x6b, 0x77, 0xea, 0xcd, 0xbf, 0x3b, 0xf5, 0xb4, 0xb0, 0x9c, 0xcd, 0xac, 0xb9, 0xe6, 0xf6, 0xbe,
	0x42, 0xab, 0x4f, 0x19, 0xf7, 0x66, 0xde, 0x70, 0x89, 0xb9, 0x9c, 0x1e, 0x7a, 0x07, 0xb7, 0x06,
	0x62, 0x89, 0x96, 0x24, 0xff, 0xa5, 0x3e, 0xd8, 0x3e, 0x79, 0xfb, 0x4e, 0x41, 0xde, 0xb7, 0xed,
	0x6f, 0x8f, 0xe6, 0x1e, 0x5f, 0xc4, 0x93, 0xce, 0x94, 0xfa, 0xdd, 0x99, 0xce, 0xd1, 0x55, 0x7f,
	0x84, 0xd3, 0x39, 0xed, 0x66, 0x5e, 0x93, 0x9a, 0xc2, 0x9e, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0x28, 0x71, 0x7a, 0x8e, 0x38, 0x06, 0x00, 0x00,
}