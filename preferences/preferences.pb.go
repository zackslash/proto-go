// Code generated by protoc-gen-go.
// source: preferences.proto
// DO NOT EDIT!

/*
Package preferences is a generated protocol buffer package.

It is generated from these files:
	preferences.proto

It has these top-level messages:
	Preference
	PreferenceResponse
*/
package preferences

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

type Preference struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Preference) Reset()                    { *m = Preference{} }
func (m *Preference) String() string            { return proto.CompactTextString(m) }
func (*Preference) ProtoMessage()               {}
func (*Preference) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Preference) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Preference) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type PreferenceResponse struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value   string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Status  int32  `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
	Message string `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
	Version string `protobuf:"bytes,5,opt,name=version" json:"version,omitempty"`
}

func (m *PreferenceResponse) Reset()                    { *m = PreferenceResponse{} }
func (m *PreferenceResponse) String() string            { return proto.CompactTextString(m) }
func (*PreferenceResponse) ProtoMessage()               {}
func (*PreferenceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PreferenceResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PreferenceResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PreferenceResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *PreferenceResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PreferenceResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*Preference)(nil), "preferences.Preference")
	proto.RegisterType((*PreferenceResponse)(nil), "preferences.PreferenceResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Preferences service

type PreferencesClient interface {
	Mutate(ctx context.Context, in *Preference, opts ...grpc.CallOption) (*PreferenceResponse, error)
	Retrieve(ctx context.Context, in *Preference, opts ...grpc.CallOption) (*PreferenceResponse, error)
	Delete(ctx context.Context, in *Preference, opts ...grpc.CallOption) (*PreferenceResponse, error)
}

type preferencesClient struct {
	cc *grpc.ClientConn
}

func NewPreferencesClient(cc *grpc.ClientConn) PreferencesClient {
	return &preferencesClient{cc}
}

func (c *preferencesClient) Mutate(ctx context.Context, in *Preference, opts ...grpc.CallOption) (*PreferenceResponse, error) {
	out := new(PreferenceResponse)
	err := grpc.Invoke(ctx, "/preferences.Preferences/Mutate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *preferencesClient) Retrieve(ctx context.Context, in *Preference, opts ...grpc.CallOption) (*PreferenceResponse, error) {
	out := new(PreferenceResponse)
	err := grpc.Invoke(ctx, "/preferences.Preferences/Retrieve", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *preferencesClient) Delete(ctx context.Context, in *Preference, opts ...grpc.CallOption) (*PreferenceResponse, error) {
	out := new(PreferenceResponse)
	err := grpc.Invoke(ctx, "/preferences.Preferences/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Preferences service

type PreferencesServer interface {
	Mutate(context.Context, *Preference) (*PreferenceResponse, error)
	Retrieve(context.Context, *Preference) (*PreferenceResponse, error)
	Delete(context.Context, *Preference) (*PreferenceResponse, error)
}

func RegisterPreferencesServer(s *grpc.Server, srv PreferencesServer) {
	s.RegisterService(&_Preferences_serviceDesc, srv)
}

func _Preferences_Mutate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Preference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PreferencesServer).Mutate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/preferences.Preferences/Mutate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PreferencesServer).Mutate(ctx, req.(*Preference))
	}
	return interceptor(ctx, in, info, handler)
}

func _Preferences_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Preference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PreferencesServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/preferences.Preferences/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PreferencesServer).Retrieve(ctx, req.(*Preference))
	}
	return interceptor(ctx, in, info, handler)
}

func _Preferences_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Preference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PreferencesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/preferences.Preferences/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PreferencesServer).Delete(ctx, req.(*Preference))
	}
	return interceptor(ctx, in, info, handler)
}

var _Preferences_serviceDesc = grpc.ServiceDesc{
	ServiceName: "preferences.Preferences",
	HandlerType: (*PreferencesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Mutate",
			Handler:    _Preferences_Mutate_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _Preferences_Retrieve_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Preferences_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "preferences.proto",
}

func init() { proto.RegisterFile("preferences.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x59, 0x6d, 0xa2, 0x4e, 0x4f, 0x0e, 0xa2, 0x8b, 0x17, 0x4b, 0x2f, 0xd6, 0x83, 0x09,
	0x28, 0xf8, 0x00, 0xa1, 0x57, 0x41, 0x72, 0xf4, 0xb6, 0x2d, 0x93, 0xb8, 0xd0, 0x64, 0xc3, 0xce,
	0x6c, 0x1e, 0xc2, 0x57, 0xf4, 0x65, 0xc4, 0x6d, 0x6d, 0x72, 0xf1, 0x20, 0xb9, 0xed, 0xf7, 0xff,
	0x33, 0xc3, 0xff, 0xb3, 0x70, 0xd9, 0x79, 0xaa, 0xc8, 0x53, 0xbb, 0x25, 0xce, 0x3a, 0xef, 0xc4,
	0xe1, 0x7c, 0x24, 0x2d, 0x5f, 0x00, 0xde, 0x8e, 0x88, 0x08, 0xb3, 0xd6, 0x34, 0xa4, 0xd5, 0x42,
	0xad, 0x2e, 0xca, 0xf8, 0xc6, 0x2b, 0x48, 0x7a, 0xb3, 0x0b, 0xa4, 0x4f, 0xa2, 0xb8, 0x87, 0xe5,
	0xa7, 0x02, 0x1c, 0x16, 0x4b, 0xe2, 0xce, 0xb5, 0xfc, 0x8f, 0x03, 0x78, 0x0d, 0x29, 0x8b, 0x91,
	0xc0, 0xfa, 0x74, 0xa1, 0x56, 0x49, 0x79, 0x20, 0xd4, 0x70, 0xd6, 0x10, 0xb3, 0xa9, 0x49, 0xcf,
	0xe2, 0xfc, 0x2f, 0xfe, 0x38, 0x3d, 0x79, 0xb6, 0xae, 0xd5, 0xc9, 0xde, 0x39, 0xe0, 0xd3, 0x97,
	0x82, 0xf9, 0x10, 0x86, 0xb1, 0x80, 0xf4, 0x35, 0x88, 0x11, 0xc2, 0x9b, 0x6c, 0xdc, 0x7f, 0x98,
	0xb9, 0xbd, 0xfb, 0xc3, 0x38, 0x36, 0x59, 0xc3, 0x79, 0x49, 0xe2, 0x2d, 0xf5, 0x53, 0xae, 0x14,
	0x90, 0xae, 0x69, 0x47, 0x53, 0x92, 0x14, 0x0f, 0xef, 0xf7, 0xb5, 0x95, 0x8f, 0xb0, 0xc9, 0xb6,
	0xae, 0xc9, 0x2b, 0xe7, 0xc5, 0x56, 0x36, 0x8f, 0x7f, 0xf9, 0x58, 0xbb, 0x7c, 0xb4, 0xbd, 0x49,
	0xa3, 0xfa, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x75, 0xb5, 0xe7, 0x96, 0xf6, 0x01, 0x00, 0x00,
}
