// Code generated by protoc-gen-go.
// source: ftypes/domain.proto
// DO NOT EDIT!

/*
Package domain is a generated protocol buffer package.

It is generated from these files:
	ftypes/domain.proto

It has these top-level messages:
	DomainTarget
	DomainService
	Domain
*/
package domain

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

type DomainServiceType int32

const (
	DomainServiceType_CUSTOMER      DomainServiceType = 0
	DomainServiceType_COMMUNICATION DomainServiceType = 1
	DomainServiceType_SUPPORT       DomainServiceType = 2
	DomainServiceType_CHAT          DomainServiceType = 3
)

var DomainServiceType_name = map[int32]string{
	0: "CUSTOMER",
	1: "COMMUNICATION",
	2: "SUPPORT",
	3: "CHAT",
}
var DomainServiceType_value = map[string]int32{
	"CUSTOMER":      0,
	"COMMUNICATION": 1,
	"SUPPORT":       2,
	"CHAT":          3,
}

func (x DomainServiceType) String() string {
	return proto.EnumName(DomainServiceType_name, int32(x))
}
func (DomainServiceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DomainTargetType int32

const (
	DomainTargetType_NONE        DomainTargetType = 0
	DomainTargetType_COUNTRY     DomainTargetType = 1
	DomainTargetType_LANGUAGE    DomainTargetType = 2
	DomainTargetType_PRODUCT     DomainTargetType = 3
	DomainTargetType_DEVELOPMENT DomainTargetType = 4
)

var DomainTargetType_name = map[int32]string{
	0: "NONE",
	1: "COUNTRY",
	2: "LANGUAGE",
	3: "PRODUCT",
	4: "DEVELOPMENT",
}
var DomainTargetType_value = map[string]int32{
	"NONE":        0,
	"COUNTRY":     1,
	"LANGUAGE":    2,
	"PRODUCT":     3,
	"DEVELOPMENT": 4,
}

func (x DomainTargetType) String() string {
	return proto.EnumName(DomainTargetType_name, int32(x))
}
func (DomainTargetType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type DomainTarget struct {
	Type  DomainTargetType `protobuf:"varint,1,opt,name=type,enum=fortifi.domain.DomainTargetType" json:"type,omitempty"`
	Value string           `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *DomainTarget) Reset()                    { *m = DomainTarget{} }
func (m *DomainTarget) String() string            { return proto.CompactTextString(m) }
func (*DomainTarget) ProtoMessage()               {}
func (*DomainTarget) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DomainTarget) GetType() DomainTargetType {
	if m != nil {
		return m.Type
	}
	return DomainTargetType_NONE
}

func (m *DomainTarget) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type DomainService struct {
	Domain    string            `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	Service   DomainServiceType `protobuf:"varint,2,opt,name=service,enum=fortifi.domain.DomainServiceType" json:"service,omitempty"`
	Subdomain string            `protobuf:"bytes,3,opt,name=subdomain" json:"subdomain,omitempty"`
}

func (m *DomainService) Reset()                    { *m = DomainService{} }
func (m *DomainService) String() string            { return proto.CompactTextString(m) }
func (*DomainService) ProtoMessage()               {}
func (*DomainService) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DomainService) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *DomainService) GetService() DomainServiceType {
	if m != nil {
		return m.Service
	}
	return DomainServiceType_CUSTOMER
}

func (m *DomainService) GetSubdomain() string {
	if m != nil {
		return m.Subdomain
	}
	return ""
}

type Domain struct {
	Domain   string           `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	BrandId  string           `protobuf:"bytes,2,opt,name=brand_id,json=brandId" json:"brand_id,omitempty"`
	Targets  []*DomainTarget  `protobuf:"bytes,3,rep,name=targets" json:"targets,omitempty"`
	Services []*DomainService `protobuf:"bytes,4,rep,name=services" json:"services,omitempty"`
}

func (m *Domain) Reset()                    { *m = Domain{} }
func (m *Domain) String() string            { return proto.CompactTextString(m) }
func (*Domain) ProtoMessage()               {}
func (*Domain) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Domain) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Domain) GetBrandId() string {
	if m != nil {
		return m.BrandId
	}
	return ""
}

func (m *Domain) GetTargets() []*DomainTarget {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *Domain) GetServices() []*DomainService {
	if m != nil {
		return m.Services
	}
	return nil
}

func init() {
	proto.RegisterType((*DomainTarget)(nil), "fortifi.domain.DomainTarget")
	proto.RegisterType((*DomainService)(nil), "fortifi.domain.DomainService")
	proto.RegisterType((*Domain)(nil), "fortifi.domain.Domain")
	proto.RegisterEnum("fortifi.domain.DomainServiceType", DomainServiceType_name, DomainServiceType_value)
	proto.RegisterEnum("fortifi.domain.DomainTargetType", DomainTargetType_name, DomainTargetType_value)
}

func init() { proto.RegisterFile("ftypes/domain.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x52, 0xef, 0xef, 0x92, 0x40,
	0x18, 0xff, 0x22, 0x24, 0xf8, 0xf8, 0xa3, 0xf3, 0x6a, 0x8d, 0x36, 0xdb, 0xc8, 0x57, 0x66, 0x0b,
	0x36, 0x6b, 0x6d, 0xad, 0x57, 0x84, 0xcc, 0x5c, 0xc2, 0xb1, 0xf3, 0x68, 0xcb, 0x37, 0x0d, 0x04,
	0x8d, 0x2d, 0xc5, 0x01, 0xba, 0xf9, 0xb6, 0x3f, 0xa8, 0xbf, 0xb1, 0x71, 0x60, 0x65, 0x3f, 0x7c,
	0x77, 0xcf, 0xf3, 0x7c, 0x7e, 0xdc, 0xe7, 0xee, 0x81, 0x07, 0x9b, 0xe2, 0x7c, 0x88, 0x73, 0x23,
	0x4a, 0x77, 0x41, 0xb2, 0xd7, 0x0f, 0x59, 0x5a, 0xa4, 0xb8, 0xb7, 0x49, 0xb3, 0x22, 0xd9, 0x24,
	0x7a, 0xd5, 0x1d, 0xae, 0xa0, 0x33, 0xe5, 0x27, 0x16, 0x64, 0xdb, 0xb8, 0xc0, 0xaf, 0x40, 0x2a,
	0x59, 0xaa, 0xa0, 0x09, 0xa3, 0xde, 0x44, 0xd3, 0xaf, 0xe1, 0xfa, 0xef, 0x58, 0x76, 0x3e, 0xc4,
	0x94, 0xa3, 0xf1, 0x43, 0xb8, 0x77, 0x0a, 0xbe, 0x1e, 0x63, 0xb5, 0xa1, 0x09, 0xa3, 0x16, 0xad,
	0x8a, 0xe1, 0x37, 0x01, 0xba, 0x15, 0x61, 0x19, 0x67, 0xa7, 0x64, 0x1d, 0xe3, 0x47, 0xd0, 0xac,
	0x84, 0xb8, 0x7e, 0x8b, 0xd6, 0x15, 0x7e, 0x0b, 0x72, 0x5e, 0x41, 0xb8, 0x42, 0x6f, 0xf2, 0xf4,
	0xdf, 0xc6, 0xb5, 0x0e, 0x77, 0xbe, 0x30, 0xf0, 0x00, 0x5a, 0xf9, 0x31, 0xac, 0x75, 0x45, 0xae,
	0xfb, 0xab, 0x31, 0xfc, 0x2e, 0x40, 0xb3, 0x22, 0xff, 0xd7, 0xfd, 0x31, 0x28, 0x61, 0x16, 0xec,
	0xa3, 0xcf, 0x49, 0x54, 0x07, 0x90, 0x79, 0x3d, 0x8f, 0xf0, 0x6b, 0x90, 0x0b, 0x1e, 0x36, 0x57,
	0x45, 0x4d, 0x1c, 0xb5, 0x27, 0x83, 0x5b, 0x2f, 0x42, 0x2f, 0x60, 0xfc, 0x06, 0x94, 0xfa, 0x7a,
	0xb9, 0x2a, 0x71, 0xe2, 0x93, 0x9b, 0x89, 0xe8, 0x4f, 0xf8, 0xf8, 0x03, 0xf4, 0xff, 0x0a, 0x8b,
	0x3b, 0xa0, 0x58, 0xfe, 0x92, 0x11, 0xc7, 0xa6, 0xe8, 0x0e, 0xf7, 0xa1, 0x6b, 0x11, 0xc7, 0xf1,
	0xdd, 0xb9, 0x65, 0xb2, 0x39, 0x71, 0x91, 0x80, 0xdb, 0x20, 0x2f, 0x7d, 0xcf, 0x23, 0x94, 0xa1,
	0x06, 0x56, 0x40, 0xb2, 0xde, 0x9b, 0x0c, 0x89, 0x63, 0x1f, 0xd0, 0x9f, 0x5f, 0x56, 0x4e, 0x5d,
	0xe2, 0xda, 0xe8, 0xae, 0x24, 0x59, 0xc4, 0x77, 0x19, 0xfd, 0x84, 0x84, 0xd2, 0x62, 0x61, 0xba,
	0x33, 0xdf, 0x9c, 0xd9, 0xa8, 0x51, 0x8e, 0x3c, 0x4a, 0xa6, 0xbe, 0xc5, 0x90, 0x88, 0xef, 0x43,
	0x7b, 0x6a, 0x7f, 0xb4, 0x17, 0xc4, 0x73, 0x6c, 0x97, 0x21, 0xe9, 0xdd, 0xf3, 0xd5, 0xb3, 0x6d,
	0x52, 0x7c, 0x39, 0x86, 0xfa, 0x3a, 0xdd, 0x19, 0x75, 0x30, 0x83, 0x6f, 0xd8, 0x8b, 0x6d, 0x6a,
	0x5c, 0x2d, 0x5e, 0xd8, 0xe4, 0xfd, 0x97, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x00, 0x7d, 0xdc,
	0xb9, 0x90, 0x02, 0x00, 0x00,
}
