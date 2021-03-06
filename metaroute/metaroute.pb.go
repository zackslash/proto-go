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
	Host           string                                `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	Path           string                                `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Method         string                                `protobuf:"bytes,3,opt,name=method" json:"method,omitempty"`
	PostParams     map[string]*HTTPRequest_HTTPParameter `protobuf:"bytes,4,rep,name=post_params,json=postParams" json:"post_params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	QueryString    string                                `protobuf:"bytes,5,opt,name=query_string,json=queryString" json:"query_string,omitempty"`
	Headers        map[string]*HTTPRequest_HTTPParameter `protobuf:"bytes,6,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body           string                                `protobuf:"bytes,7,opt,name=body" json:"body,omitempty"`
	Language       string                                `protobuf:"bytes,8,opt,name=language" json:"language,omitempty"`
	ContentType    string                                `protobuf:"bytes,9,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	OrganisationId string                                `protobuf:"bytes,10,opt,name=organisation_id,json=organisationId" json:"organisation_id,omitempty"`
	Customer       *HTTPRequest_CustomerInformation      `protobuf:"bytes,11,opt,name=customer" json:"customer,omitempty"`
}

func (m *HTTPRequest) Reset()                    { *m = HTTPRequest{} }
func (m *HTTPRequest) String() string            { return proto.CompactTextString(m) }
func (*HTTPRequest) ProtoMessage()               {}
func (*HTTPRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HTTPRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HTTPRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HTTPRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *HTTPRequest) GetPostParams() map[string]*HTTPRequest_HTTPParameter {
	if m != nil {
		return m.PostParams
	}
	return nil
}

func (m *HTTPRequest) GetQueryString() string {
	if m != nil {
		return m.QueryString
	}
	return ""
}

func (m *HTTPRequest) GetHeaders() map[string]*HTTPRequest_HTTPParameter {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTTPRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *HTTPRequest) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *HTTPRequest) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *HTTPRequest) GetOrganisationId() string {
	if m != nil {
		return m.OrganisationId
	}
	return ""
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

func (m *HTTPRequest_HTTPParameter) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

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
	return fileDescriptor0, []int{0, 3}
}

func (m *HTTPRequest_CustomerInformation) GetBrandId() string {
	if m != nil {
		return m.BrandId
	}
	return ""
}

func (m *HTTPRequest_CustomerInformation) GetCustomerFid() string {
	if m != nil {
		return m.CustomerFid
	}
	return ""
}

func (m *HTTPRequest_CustomerInformation) GetCustomerName() string {
	if m != nil {
		return m.CustomerName
	}
	return ""
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

func (m *HTTPResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *HTTPResponse) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *HTTPResponse) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *HTTPResponse) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

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

func (m *HTTPResponse_HTTPHeaderParameter) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
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
const _ = grpc.SupportPackageIsVersion4

// Client API for MetaRoute service

type MetaRouteClient interface {
	HandleHTTPRequest(ctx context.Context, in *HTTPRequest, opts ...grpc.CallOption) (*HTTPResponse, error)
}

type metaRouteClient struct {
	cc *grpc.ClientConn
}

func NewMetaRouteClient(cc *grpc.ClientConn) MetaRouteClient {
	return &metaRouteClient{cc}
}

func (c *metaRouteClient) HandleHTTPRequest(ctx context.Context, in *HTTPRequest, opts ...grpc.CallOption) (*HTTPResponse, error) {
	out := new(HTTPResponse)
	err := grpc.Invoke(ctx, "/metaroute.MetaRoute/HandleHTTPRequest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MetaRoute service

type MetaRouteServer interface {
	HandleHTTPRequest(context.Context, *HTTPRequest) (*HTTPResponse, error)
}

func RegisterMetaRouteServer(s *grpc.Server, srv MetaRouteServer) {
	s.RegisterService(&_MetaRoute_serviceDesc, srv)
}

func _MetaRoute_HandleHTTPRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaRouteServer).HandleHTTPRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metaroute.MetaRoute/HandleHTTPRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaRouteServer).HandleHTTPRequest(ctx, req.(*HTTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetaRoute_serviceDesc = grpc.ServiceDesc{
	ServiceName: "metaroute.MetaRoute",
	HandlerType: (*MetaRouteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleHTTPRequest",
			Handler:    _MetaRoute_HandleHTTPRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metaroute.proto",
}

func init() { proto.RegisterFile("metaroute.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x54, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0x55, 0xff, 0xad, 0xed, 0xcd, 0xf6, 0xdb, 0x0f, 0x0f, 0x0d, 0x93, 0x17, 0xca, 0x06, 0x6c,
	0x02, 0xad, 0x93, 0x0a, 0x12, 0x68, 0x12, 0x48, 0x30, 0x18, 0xdd, 0x03, 0x68, 0x84, 0x3d, 0xf1,
	0x52, 0xdc, 0xfa, 0x36, 0x8d, 0x58, 0xec, 0xcc, 0x76, 0x90, 0xfa, 0x95, 0xf8, 0x48, 0x7c, 0x15,
	0x5e, 0x90, 0xed, 0x2c, 0x0b, 0x5b, 0xab, 0xf1, 0xc2, 0x9b, 0x7d, 0x72, 0x6f, 0xce, 0xf5, 0x39,
	0xc7, 0x86, 0xf5, 0x14, 0x0d, 0x53, 0x32, 0x37, 0xd8, 0xcf, 0x94, 0x34, 0x92, 0x74, 0x4b, 0x60,
	0xeb, 0x47, 0x1b, 0x82, 0xe1, 0xe9, 0xe9, 0x49, 0x84, 0xe7, 0x39, 0x6a, 0x43, 0x08, 0x34, 0x67,
	0x52, 0x1b, 0x5a, 0xeb, 0xd5, 0x76, 0xbb, 0x91, 0x5b, 0x5b, 0x2c, 0x63, 0x66, 0x46, 0xeb, 0x1e,
	0xb3, 0x6b, 0xb2, 0x09, 0x2b, 0x29, 0x9a, 0x99, 0xe4, 0xb4, 0xe1, 0xd0, 0x62, 0x47, 0xde, 0x43,
	0x90, 0x49, 0x6d, 0x46, 0x19, 0x53, 0x2c, 0xd5, 0xb4, 0xd9, 0x6b, 0xec, 0x06, 0x83, 0x47, 0xfd,
	0xcb, 0x09, 0x2a, 0x64, 0xfd, 0x13, 0xa9, 0xcd, 0x89, 0x2b, 0x7c, 0x27, 0x8c, 0x9a, 0x47, 0x90,
	0x95, 0x00, 0xb9, 0x0f, 0xab, 0xe7, 0x39, 0xaa, 0xf9, 0x48, 0x1b, 0x95, 0x88, 0x98, 0xb6, 0x1c,
	0x4d, 0xe0, 0xb0, 0xcf, 0x0e, 0x22, 0x2f, 0xa1, 0x3d, 0x43, 0xc6, 0x51, 0x69, 0xba, 0xe2, 0x78,
	0xb6, 0x97, 0xf0, 0x0c, 0x7d, 0x95, 0x27, 0xb9, 0xe8, 0xb1, 0xc7, 0x1a, 0x4b, 0x3e, 0xa7, 0x6d,
	0x7f, 0x2c, 0xbb, 0x26, 0x21, 0x74, 0xce, 0x98, 0x88, 0x73, 0x16, 0x23, 0xed, 0x38, 0xbc, 0xdc,
	0xdb, 0x89, 0x26, 0x52, 0x18, 0x14, 0x66, 0x64, 0xe6, 0x19, 0xd2, 0xae, 0x9f, 0xa8, 0xc0, 0x4e,
	0xe7, 0x19, 0x92, 0x1d, 0x58, 0x97, 0x2a, 0x66, 0x22, 0xd1, 0xcc, 0x24, 0x52, 0x8c, 0x12, 0x4e,
	0xc1, 0x55, 0xfd, 0x57, 0x85, 0x8f, 0x39, 0x39, 0x82, 0xce, 0x24, 0xd7, 0x46, 0xa6, 0xa8, 0x68,
	0xd0, 0xab, 0xed, 0x06, 0x83, 0xc7, 0x4b, 0x66, 0x3f, 0x2c, 0xca, 0x8e, 0xc5, 0x54, 0xaa, 0xd4,
	0xf5, 0x47, 0x65, 0x6f, 0xb8, 0x03, 0x6b, 0xb6, 0xd8, 0x69, 0x86, 0x06, 0x95, 0xf5, 0xe5, 0x3b,
	0x3b, 0xcb, 0x51, 0xd3, 0x5a, 0xaf, 0x61, 0x7d, 0xf1, 0xbb, 0x70, 0x02, 0xeb, 0x57, 0xd4, 0x26,
	0xff, 0x43, 0xe3, 0x1b, 0xce, 0x0b, 0xa7, 0xed, 0x92, 0x1c, 0x40, 0xcb, 0x95, 0x3b, 0xa7, 0x83,
	0xc1, 0x83, 0x65, 0x72, 0x56, 0x19, 0x23, 0xdf, 0x72, 0x50, 0x7f, 0x51, 0x0b, 0xbf, 0xc2, 0x6a,
	0x55, 0xea, 0x7f, 0xc0, 0xf0, 0xab, 0x06, 0x1b, 0x0b, 0x14, 0x21, 0x77, 0xa1, 0x33, 0x56, 0x4c,
	0x70, 0xab, 0xb8, 0xa7, 0x6b, 0xbb, 0xfd, 0x31, 0x77, 0xb6, 0x15, 0x1d, 0xa3, 0x69, 0xc2, 0x8b,
	0x14, 0x07, 0x17, 0xd8, 0x51, 0xc2, 0xc9, 0x36, 0xac, 0x95, 0x25, 0x82, 0xa5, 0x58, 0x64, 0xba,
	0xec, 0xfb, 0xc8, 0x52, 0x24, 0x43, 0x68, 0xda, 0x61, 0x8b, 0x48, 0x3f, 0xfb, 0x7b, 0xbb, 0xfa,
	0x1f, 0xd0, 0x30, 0x9f, 0x3d, 0xf7, 0x87, 0xf0, 0x39, 0x74, 0x4b, 0x68, 0x81, 0x46, 0xb7, 0xab,
	0x1a, 0x75, 0x2b, 0xa7, 0xdf, 0xfa, 0x59, 0x87, 0x55, 0x4f, 0xa6, 0x33, 0x29, 0x34, 0x92, 0x7b,
	0x10, 0x68, 0xc3, 0x4c, 0xae, 0x47, 0x13, 0xc9, 0xd1, 0xfd, 0xa4, 0x15, 0x81, 0x87, 0x0e, 0x25,
	0xc7, 0x32, 0xe3, 0xf5, 0x4a, 0xc6, 0xaf, 0xe6, 0xb8, 0x71, 0x3d, 0xc7, 0xd5, 0x6b, 0xd0, 0xbc,
	0x72, 0x0d, 0x5e, 0x5d, 0xde, 0xba, 0x96, 0x93, 0xe2, 0xba, 0x89, 0x7e, 0xba, 0xc5, 0xd7, 0x2e,
	0xdc, 0x83, 0x0d, 0x5b, 0xe5, 0x3f, 0xde, 0x1c, 0xdc, 0xf8, 0xc6, 0x4c, 0xbd, 0xfe, 0x33, 0x53,
	0x4f, 0x96, 0x8e, 0x73, 0x9d, 0xb5, 0x22, 0xee, 0xe0, 0x93, 0x77, 0x25, 0xb2, 0x8d, 0xe4, 0x2d,
	0xdc, 0x1a, 0x32, 0xc1, 0xcf, 0xb0, 0xfa, 0x36, 0x6e, 0x2e, 0xf6, 0x3c, 0xbc, 0xb3, 0x84, 0xf1,
	0xcd, 0xce, 0x97, 0x87, 0x71, 0x62, 0x66, 0xf9, 0xb8, 0x3f, 0x91, 0xe9, 0xfe, 0x54, 0x2a, 0x93,
	0x4c, 0x93, 0x7d, 0xf7, 0x06, 0xef, 0xc5, 0x72, 0xbf, 0xec, 0x1a, 0xaf, 0x38, 0xec, 0xe9, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x84, 0xdb, 0x17, 0xaa, 0x05, 0x00, 0x00,
}
