// Code generated by protoc-gen-go.
// source: fcommon.proto
// DO NOT EDIT!

/*
Package fcommon is a generated protocol buffer package.

It is generated from these files:
	fcommon.proto

It has these top-level messages:
	BoolResponse
	IDRequest
*/
package fcommon

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BoolResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *BoolResponse) Reset()                    { *m = BoolResponse{} }
func (m *BoolResponse) String() string            { return proto.CompactTextString(m) }
func (*BoolResponse) ProtoMessage()               {}
func (*BoolResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BoolResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *BoolResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type IDRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *IDRequest) Reset()                    { *m = IDRequest{} }
func (m *IDRequest) String() string            { return proto.CompactTextString(m) }
func (*IDRequest) ProtoMessage()               {}
func (*IDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IDRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*BoolResponse)(nil), "fcommon.BoolResponse")
	proto.RegisterType((*IDRequest)(nil), "fcommon.IDRequest")
}

func init() { proto.RegisterFile("fcommon.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4b, 0xce, 0xcf,
	0xcd, 0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x9c, 0xb8,
	0x78, 0x9c, 0xf2, 0xf3, 0x73, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x24, 0xb8,
	0xd8, 0x8b, 0x4b, 0x93, 0x93, 0x53, 0x8b, 0x8b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x82, 0x60,
	0x5c, 0x90, 0x4c, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x8c, 0xab, 0x24, 0xcd, 0xc5, 0xe9, 0xe9, 0x12, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22,
	0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x02, 0xd6, 0xcb, 0x19, 0xc4, 0x94, 0x99, 0xe2, 0xa4, 0x1a, 0xa5,
	0x9c, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x96, 0x5f, 0x54, 0x92,
	0x99, 0x96, 0xa9, 0x0f, 0x76, 0x85, 0x6e, 0x7a, 0xbe, 0x3e, 0xd4, 0x1d, 0x49, 0x6c, 0x60, 0x11,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf7, 0xcd, 0x9c, 0x39, 0xa8, 0x00, 0x00, 0x00,
}
