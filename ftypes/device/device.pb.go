// Code generated by protoc-gen-go.
// source: ftypes/device.proto
// DO NOT EDIT!

/*
Package device is a generated protocol buffer package.

It is generated from these files:
	ftypes/device.proto

It has these top-level messages:
*/
package device

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

type OperatingSystem int32

const (
	OperatingSystem_WINDOWS    OperatingSystem = 0
	OperatingSystem_OSX        OperatingSystem = 1
	OperatingSystem_LINUX      OperatingSystem = 2
	OperatingSystem_ANDROID    OperatingSystem = 3
	OperatingSystem_BLACKBERRY OperatingSystem = 4
	OperatingSystem_IOS        OperatingSystem = 5
	OperatingSystem_FIREFOX_OS OperatingSystem = 6
	OperatingSystem_CHROME_OS  OperatingSystem = 7
	OperatingSystem_WEB_OS     OperatingSystem = 8
	OperatingSystem_PALM_OS    OperatingSystem = 9
	OperatingSystem_SYMBIAN    OperatingSystem = 10
	OperatingSystem_TIZEN      OperatingSystem = 11
	OperatingSystem_BSD        OperatingSystem = 12
)

var OperatingSystem_name = map[int32]string{
	0:  "WINDOWS",
	1:  "OSX",
	2:  "LINUX",
	3:  "ANDROID",
	4:  "BLACKBERRY",
	5:  "IOS",
	6:  "FIREFOX_OS",
	7:  "CHROME_OS",
	8:  "WEB_OS",
	9:  "PALM_OS",
	10: "SYMBIAN",
	11: "TIZEN",
	12: "BSD",
}
var OperatingSystem_value = map[string]int32{
	"WINDOWS":    0,
	"OSX":        1,
	"LINUX":      2,
	"ANDROID":    3,
	"BLACKBERRY": 4,
	"IOS":        5,
	"FIREFOX_OS": 6,
	"CHROME_OS":  7,
	"WEB_OS":     8,
	"PALM_OS":    9,
	"SYMBIAN":    10,
	"TIZEN":      11,
	"BSD":        12,
}

func (x OperatingSystem) String() string {
	return proto.EnumName(OperatingSystem_name, int32(x))
}
func (OperatingSystem) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DeviceType int32

const (
	DeviceType_DESKTOP      DeviceType = 0
	DeviceType_TABLET       DeviceType = 1
	DeviceType_PHONE        DeviceType = 2
	DeviceType_TV           DeviceType = 3
	DeviceType_CAMERA       DeviceType = 4
	DeviceType_CAR          DeviceType = 5
	DeviceType_CONSOLE      DeviceType = 6
	DeviceType_MEDIA_PLAYER DeviceType = 7
	DeviceType_GAME_CONSOLE DeviceType = 8
	DeviceType_BOT          DeviceType = 9
)

var DeviceType_name = map[int32]string{
	0: "DESKTOP",
	1: "TABLET",
	2: "PHONE",
	3: "TV",
	4: "CAMERA",
	5: "CAR",
	6: "CONSOLE",
	7: "MEDIA_PLAYER",
	8: "GAME_CONSOLE",
	9: "BOT",
}
var DeviceType_value = map[string]int32{
	"DESKTOP":      0,
	"TABLET":       1,
	"PHONE":        2,
	"TV":           3,
	"CAMERA":       4,
	"CAR":          5,
	"CONSOLE":      6,
	"MEDIA_PLAYER": 7,
	"GAME_CONSOLE": 8,
	"BOT":          9,
}

func (x DeviceType) String() string {
	return proto.EnumName(DeviceType_name, int32(x))
}
func (DeviceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type InstallAction int32

const (
	InstallAction_INSTALL   InstallAction = 0
	InstallAction_UNINSTALL InstallAction = 1
	InstallAction_UPGRADE   InstallAction = 2
	InstallAction_REINSTALL InstallAction = 3
)

var InstallAction_name = map[int32]string{
	0: "INSTALL",
	1: "UNINSTALL",
	2: "UPGRADE",
	3: "REINSTALL",
}
var InstallAction_value = map[string]int32{
	"INSTALL":   0,
	"UNINSTALL": 1,
	"UPGRADE":   2,
	"REINSTALL": 3,
}

func (x InstallAction) String() string {
	return proto.EnumName(InstallAction_name, int32(x))
}
func (InstallAction) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterEnum("fortifi.device.OperatingSystem", OperatingSystem_name, OperatingSystem_value)
	proto.RegisterEnum("fortifi.device.DeviceType", DeviceType_name, DeviceType_value)
	proto.RegisterEnum("fortifi.device.InstallAction", InstallAction_name, InstallAction_value)
}

func init() { proto.RegisterFile("ftypes/device.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x91, 0x5f, 0x8e, 0xda, 0x30,
	0x10, 0xc6, 0x9b, 0xa5, 0x1b, 0x36, 0xb3, 0x7f, 0x6a, 0xa5, 0xa7, 0x28, 0x55, 0xc9, 0x43, 0x4f,
	0xe0, 0xc4, 0x5e, 0xd6, 0x5a, 0xc7, 0x8e, 0x6c, 0x53, 0x60, 0x5f, 0x10, 0xd0, 0x40, 0x23, 0x01,
	0x41, 0xe0, 0x56, 0xe2, 0x04, 0x3d, 0x4f, 0x6f, 0xd8, 0x71, 0x28, 0x0f, 0x7d, 0xf3, 0x37, 0xdf,
	0xcc, 0xe7, 0xdf, 0x68, 0xe0, 0xe3, 0xda, 0x9f, 0x0f, 0xf5, 0x29, 0xfb, 0x5e, 0xff, 0x6a, 0x56,
	0xf5, 0xf0, 0x70, 0x6c, 0x7d, 0x9b, 0x3e, 0xad, 0xdb, 0xa3, 0x6f, 0xd6, 0xcd, 0xf0, 0x52, 0x1d,
	0xfc, 0x89, 0xe0, 0x83, 0x3e, 0xd4, 0xc7, 0x85, 0x6f, 0xf6, 0x1b, 0x7b, 0x3e, 0xf9, 0x7a, 0x97,
	0xde, 0x43, 0x7f, 0x22, 0x14, 0xd3, 0x13, 0x4b, 0xde, 0xa5, 0x7d, 0xe8, 0x69, 0x3b, 0x25, 0x51,
	0x9a, 0xc0, 0xad, 0x14, 0x6a, 0x3c, 0x25, 0x37, 0xa1, 0x81, 0x2a, 0x66, 0xb4, 0x60, 0xa4, 0x97,
	0x3e, 0x01, 0xe4, 0x92, 0x16, 0xaf, 0x39, 0x37, 0x66, 0x46, 0xde, 0x87, 0x01, 0xa1, 0x2d, 0xb9,
	0x0d, 0xc6, 0xb3, 0x30, 0xfc, 0x59, 0x4f, 0xe7, 0xa8, 0xe3, 0xf4, 0x11, 0x92, 0xe2, 0xc5, 0xe8,
	0x92, 0x07, 0xd9, 0x4f, 0x01, 0xe2, 0x09, 0xcf, 0xc3, 0xfb, 0x2e, 0x04, 0x56, 0x54, 0x96, 0x41,
	0x24, 0x41, 0xd8, 0x59, 0x99, 0x0b, 0xaa, 0x08, 0x84, 0x5f, 0x9d, 0x78, 0xe3, 0x8a, 0xdc, 0x87,
	0xe0, 0xdc, 0x32, 0xf2, 0x30, 0xf8, 0x1d, 0x01, 0xb0, 0x0e, 0xdf, 0xe1, 0x82, 0xa1, 0x9f, 0x71,
	0xfb, 0xea, 0x74, 0x85, 0xb8, 0x98, 0xea, 0x68, 0x2e, 0xb9, 0xbb, 0x10, 0x57, 0x2f, 0x5a, 0x71,
	0x24, 0x8e, 0xe1, 0xc6, 0x7d, 0x43, 0x58, 0xb4, 0x0b, 0x5a, 0x72, 0x43, 0x2f, 0xa0, 0x05, 0x35,
	0x08, 0x8a, 0x01, 0x85, 0x56, 0x56, 0x4b, 0x8e, 0x94, 0x04, 0x1e, 0x4a, 0xce, 0x04, 0x9d, 0x57,
	0x92, 0xce, 0xb8, 0x41, 0x50, 0xac, 0x8c, 0x70, 0x66, 0x7e, 0xed, 0xb9, 0xeb, 0x48, 0xb4, 0x23,
	0xc9, 0x60, 0x04, 0x8f, 0x62, 0x7f, 0xf2, 0x8b, 0xed, 0x96, 0xae, 0x7c, 0xd3, 0xee, 0x43, 0x94,
	0x50, 0xd6, 0x51, 0x29, 0x91, 0x05, 0x17, 0x1e, 0xab, 0xab, 0x8c, 0x82, 0x37, 0xae, 0x46, 0x86,
	0xb2, 0x00, 0x84, 0x9e, 0xe1, 0x57, 0xaf, 0x97, 0x7f, 0x7e, 0xfb, 0xb4, 0x69, 0xfc, 0x8f, 0x9f,
	0xcb, 0xe1, 0xaa, 0xdd, 0x65, 0xff, 0x6e, 0x94, 0x75, 0x27, 0xfb, 0xb2, 0x69, 0xb3, 0xff, 0x2e,
	0xb9, 0x8c, 0xbb, 0xfa, 0xd7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x16, 0x43, 0x0f, 0xb5, 0xe1,
	0x01, 0x00, 0x00,
}
