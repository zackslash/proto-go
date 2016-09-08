// Code generated by protoc-gen-go.
// source: organisations.proto
// DO NOT EDIT!

/*
Package organisations is a generated protocol buffer package.

It is generated from these files:
	organisations.proto

It has these top-level messages:
	CreateRequest
	RetrieveRequest
	OrganisationResponse
	AddMemberRequest
	RemoveMemberRequest
	InviteRequest
	InviteStatusResponse
	ActionInviteRequest
	MemberStatusResponse
	GetMembersRequest
	GetOrganisationsRequest
	MembersResponse
	OrganisationsResponse
*/
package organisations

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

type Mode int32

const (
	Mode_SANDBOX    Mode = 0
	Mode_PRODUCTION Mode = 1
)

var Mode_name = map[int32]string{
	0: "SANDBOX",
	1: "PRODUCTION",
}
var Mode_value = map[string]int32{
	"SANDBOX":    0,
	"PRODUCTION": 1,
}

func (x Mode) String() string {
	return proto.EnumName(Mode_name, int32(x))
}
func (Mode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Region int32

const (
	Region_GLOBAL Region = 0
	Region_US     Region = 1
	Region_EUROPE Region = 2
	Region_ASIA   Region = 3
)

var Region_name = map[int32]string{
	0: "GLOBAL",
	1: "US",
	2: "EUROPE",
	3: "ASIA",
}
var Region_value = map[string]int32{
	"GLOBAL": 0,
	"US":     1,
	"EUROPE": 2,
	"ASIA":   3,
}

func (x Region) String() string {
	return proto.EnumName(Region_name, int32(x))
}
func (Region) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type MemberType int32

const (
	MemberType_OWNER   MemberType = 0
	MemberType_MEMBER  MemberType = 1
	MemberType_VIEWER  MemberType = 2
	MemberType_SUPPORT MemberType = 3
)

var MemberType_name = map[int32]string{
	0: "OWNER",
	1: "MEMBER",
	2: "VIEWER",
	3: "SUPPORT",
}
var MemberType_value = map[string]int32{
	"OWNER":   0,
	"MEMBER":  1,
	"VIEWER":  2,
	"SUPPORT": 3,
}

func (x MemberType) String() string {
	return proto.EnumName(MemberType_name, int32(x))
}
func (MemberType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type InviteStatus int32

const (
	InviteStatus_PENDING   InviteStatus = 0
	InviteStatus_ACCEPTED  InviteStatus = 1
	InviteStatus_REJECTED  InviteStatus = 2
	InviteStatus_EXPIRED   InviteStatus = 3
	InviteStatus_NOT_FOUND InviteStatus = 4
)

var InviteStatus_name = map[int32]string{
	0: "PENDING",
	1: "ACCEPTED",
	2: "REJECTED",
	3: "EXPIRED",
	4: "NOT_FOUND",
}
var InviteStatus_value = map[string]int32{
	"PENDING":   0,
	"ACCEPTED":  1,
	"REJECTED":  2,
	"EXPIRED":   3,
	"NOT_FOUND": 4,
}

func (x InviteStatus) String() string {
	return proto.EnumName(InviteStatus_name, int32(x))
}
func (InviteStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type CreateRequest struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	ResellerId string `protobuf:"bytes,3,opt,name=reseller_id,json=resellerId" json:"reseller_id,omitempty"`
	Mode       Mode   `protobuf:"varint,4,opt,name=mode,enum=organisations.Mode" json:"mode,omitempty"`
	Region     Region `protobuf:"varint,5,opt,name=region,enum=organisations.Region" json:"region,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RetrieveRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *RetrieveRequest) Reset()                    { *m = RetrieveRequest{} }
func (m *RetrieveRequest) String() string            { return proto.CompactTextString(m) }
func (*RetrieveRequest) ProtoMessage()               {}
func (*RetrieveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type OrganisationResponse struct {
	Id         string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name       string                     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	ResellerId string                     `protobuf:"bytes,3,opt,name=reseller_id,json=resellerId" json:"reseller_id,omitempty"`
	Mode       Mode                       `protobuf:"varint,4,opt,name=mode,enum=organisations.Mode" json:"mode,omitempty"`
	Region     Region                     `protobuf:"varint,5,opt,name=region,enum=organisations.Region" json:"region,omitempty"`
	Created    *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=created" json:"created,omitempty"`
}

func (m *OrganisationResponse) Reset()                    { *m = OrganisationResponse{} }
func (m *OrganisationResponse) String() string            { return proto.CompactTextString(m) }
func (*OrganisationResponse) ProtoMessage()               {}
func (*OrganisationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *OrganisationResponse) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

type AddMemberRequest struct {
	MemberId       string                     `protobuf:"bytes,1,opt,name=member_id,json=memberId" json:"member_id,omitempty"`
	OrganisationId string                     `protobuf:"bytes,2,opt,name=organisation_id,json=organisationId" json:"organisation_id,omitempty"`
	Type           MemberType                 `protobuf:"varint,3,opt,name=type,enum=organisations.MemberType" json:"type,omitempty"`
	Expiry         *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=expiry" json:"expiry,omitempty"`
}

func (m *AddMemberRequest) Reset()                    { *m = AddMemberRequest{} }
func (m *AddMemberRequest) String() string            { return proto.CompactTextString(m) }
func (*AddMemberRequest) ProtoMessage()               {}
func (*AddMemberRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AddMemberRequest) GetExpiry() *google_protobuf.Timestamp {
	if m != nil {
		return m.Expiry
	}
	return nil
}

type RemoveMemberRequest struct {
	MemberId       string `protobuf:"bytes,1,opt,name=member_id,json=memberId" json:"member_id,omitempty"`
	OrganisationId string `protobuf:"bytes,2,opt,name=organisation_id,json=organisationId" json:"organisation_id,omitempty"`
}

func (m *RemoveMemberRequest) Reset()                    { *m = RemoveMemberRequest{} }
func (m *RemoveMemberRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveMemberRequest) ProtoMessage()               {}
func (*RemoveMemberRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type InviteRequest struct {
	Email          string                     `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	OrganisationId string                     `protobuf:"bytes,2,opt,name=organisation_id,json=organisationId" json:"organisation_id,omitempty"`
	Type           MemberType                 `protobuf:"varint,3,opt,name=type,enum=organisations.MemberType" json:"type,omitempty"`
	Expiry         *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=expiry" json:"expiry,omitempty"`
}

func (m *InviteRequest) Reset()                    { *m = InviteRequest{} }
func (m *InviteRequest) String() string            { return proto.CompactTextString(m) }
func (*InviteRequest) ProtoMessage()               {}
func (*InviteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *InviteRequest) GetExpiry() *google_protobuf.Timestamp {
	if m != nil {
		return m.Expiry
	}
	return nil
}

type InviteStatusResponse struct {
	Email        string                     `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Status       InviteStatus               `protobuf:"varint,2,opt,name=status,enum=organisations.InviteStatus" json:"status,omitempty"`
	Verification string                     `protobuf:"bytes,3,opt,name=verification" json:"verification,omitempty"`
	Created      *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=created" json:"created,omitempty"`
	Actioned     *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=actioned" json:"actioned,omitempty"`
}

func (m *InviteStatusResponse) Reset()                    { *m = InviteStatusResponse{} }
func (m *InviteStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*InviteStatusResponse) ProtoMessage()               {}
func (*InviteStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *InviteStatusResponse) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *InviteStatusResponse) GetActioned() *google_protobuf.Timestamp {
	if m != nil {
		return m.Actioned
	}
	return nil
}

type ActionInviteRequest struct {
	Email          string       `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	OrganisationId string       `protobuf:"bytes,2,opt,name=organisation_id,json=organisationId" json:"organisation_id,omitempty"`
	MemberId       string       `protobuf:"bytes,3,opt,name=member_id,json=memberId" json:"member_id,omitempty"`
	Status         InviteStatus `protobuf:"varint,4,opt,name=status,enum=organisations.InviteStatus" json:"status,omitempty"`
	Verification   string       `protobuf:"bytes,5,opt,name=verification" json:"verification,omitempty"`
}

func (m *ActionInviteRequest) Reset()                    { *m = ActionInviteRequest{} }
func (m *ActionInviteRequest) String() string            { return proto.CompactTextString(m) }
func (*ActionInviteRequest) ProtoMessage()               {}
func (*ActionInviteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type MemberStatusResponse struct {
	MemberId       string                     `protobuf:"bytes,1,opt,name=member_id,json=memberId" json:"member_id,omitempty"`
	OrganisationId string                     `protobuf:"bytes,2,opt,name=organisation_id,json=organisationId" json:"organisation_id,omitempty"`
	Type           MemberType                 `protobuf:"varint,3,opt,name=type,enum=organisations.MemberType" json:"type,omitempty"`
	Invited        *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=invited" json:"invited,omitempty"`
	Expiry         *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=expiry" json:"expiry,omitempty"`
	Joined         *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=joined" json:"joined,omitempty"`
	Detached       *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=detached" json:"detached,omitempty"`
}

func (m *MemberStatusResponse) Reset()                    { *m = MemberStatusResponse{} }
func (m *MemberStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*MemberStatusResponse) ProtoMessage()               {}
func (*MemberStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *MemberStatusResponse) GetInvited() *google_protobuf.Timestamp {
	if m != nil {
		return m.Invited
	}
	return nil
}

func (m *MemberStatusResponse) GetExpiry() *google_protobuf.Timestamp {
	if m != nil {
		return m.Expiry
	}
	return nil
}

func (m *MemberStatusResponse) GetJoined() *google_protobuf.Timestamp {
	if m != nil {
		return m.Joined
	}
	return nil
}

func (m *MemberStatusResponse) GetDetached() *google_protobuf.Timestamp {
	if m != nil {
		return m.Detached
	}
	return nil
}

type GetMembersRequest struct {
	OrganisationId string `protobuf:"bytes,1,opt,name=organisation_id,json=organisationId" json:"organisation_id,omitempty"`
}

func (m *GetMembersRequest) Reset()                    { *m = GetMembersRequest{} }
func (m *GetMembersRequest) String() string            { return proto.CompactTextString(m) }
func (*GetMembersRequest) ProtoMessage()               {}
func (*GetMembersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type GetOrganisationsRequest struct {
	MemberId string `protobuf:"bytes,1,opt,name=member_id,json=memberId" json:"member_id,omitempty"`
}

func (m *GetOrganisationsRequest) Reset()                    { *m = GetOrganisationsRequest{} }
func (m *GetOrganisationsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetOrganisationsRequest) ProtoMessage()               {}
func (*GetOrganisationsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type MembersResponse struct {
	Members map[string]*MemberStatusResponse `protobuf:"bytes,1,rep,name=members" json:"members,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *MembersResponse) Reset()                    { *m = MembersResponse{} }
func (m *MembersResponse) String() string            { return proto.CompactTextString(m) }
func (*MembersResponse) ProtoMessage()               {}
func (*MembersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *MembersResponse) GetMembers() map[string]*MemberStatusResponse {
	if m != nil {
		return m.Members
	}
	return nil
}

type OrganisationsResponse struct {
	Organisations map[string]*OrganisationResponse `protobuf:"bytes,1,rep,name=organisations" json:"organisations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *OrganisationsResponse) Reset()                    { *m = OrganisationsResponse{} }
func (m *OrganisationsResponse) String() string            { return proto.CompactTextString(m) }
func (*OrganisationsResponse) ProtoMessage()               {}
func (*OrganisationsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *OrganisationsResponse) GetOrganisations() map[string]*OrganisationResponse {
	if m != nil {
		return m.Organisations
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "organisations.CreateRequest")
	proto.RegisterType((*RetrieveRequest)(nil), "organisations.RetrieveRequest")
	proto.RegisterType((*OrganisationResponse)(nil), "organisations.OrganisationResponse")
	proto.RegisterType((*AddMemberRequest)(nil), "organisations.AddMemberRequest")
	proto.RegisterType((*RemoveMemberRequest)(nil), "organisations.RemoveMemberRequest")
	proto.RegisterType((*InviteRequest)(nil), "organisations.InviteRequest")
	proto.RegisterType((*InviteStatusResponse)(nil), "organisations.InviteStatusResponse")
	proto.RegisterType((*ActionInviteRequest)(nil), "organisations.ActionInviteRequest")
	proto.RegisterType((*MemberStatusResponse)(nil), "organisations.MemberStatusResponse")
	proto.RegisterType((*GetMembersRequest)(nil), "organisations.GetMembersRequest")
	proto.RegisterType((*GetOrganisationsRequest)(nil), "organisations.GetOrganisationsRequest")
	proto.RegisterType((*MembersResponse)(nil), "organisations.MembersResponse")
	proto.RegisterType((*OrganisationsResponse)(nil), "organisations.OrganisationsResponse")
	proto.RegisterEnum("organisations.Mode", Mode_name, Mode_value)
	proto.RegisterEnum("organisations.Region", Region_name, Region_value)
	proto.RegisterEnum("organisations.MemberType", MemberType_name, MemberType_value)
	proto.RegisterEnum("organisations.InviteStatus", InviteStatus_name, InviteStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Organisations service

type OrganisationsClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*OrganisationResponse, error)
	Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*OrganisationResponse, error)
	GetOrganisations(ctx context.Context, in *GetOrganisationsRequest, opts ...grpc.CallOption) (*OrganisationsResponse, error)
	AddMember(ctx context.Context, in *AddMemberRequest, opts ...grpc.CallOption) (*MemberStatusResponse, error)
	RemoveMember(ctx context.Context, in *RemoveMemberRequest, opts ...grpc.CallOption) (*MemberStatusResponse, error)
	GetMembers(ctx context.Context, in *GetMembersRequest, opts ...grpc.CallOption) (*MembersResponse, error)
	Invite(ctx context.Context, in *InviteRequest, opts ...grpc.CallOption) (*InviteStatusResponse, error)
	ActionInvite(ctx context.Context, in *ActionInviteRequest, opts ...grpc.CallOption) (*MemberStatusResponse, error)
}

type organisationsClient struct {
	cc *grpc.ClientConn
}

func NewOrganisationsClient(cc *grpc.ClientConn) OrganisationsClient {
	return &organisationsClient{cc}
}

func (c *organisationsClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*OrganisationResponse, error) {
	out := new(OrganisationResponse)
	err := grpc.Invoke(ctx, "/organisations.Organisations/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organisationsClient) Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*OrganisationResponse, error) {
	out := new(OrganisationResponse)
	err := grpc.Invoke(ctx, "/organisations.Organisations/Retrieve", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organisationsClient) GetOrganisations(ctx context.Context, in *GetOrganisationsRequest, opts ...grpc.CallOption) (*OrganisationsResponse, error) {
	out := new(OrganisationsResponse)
	err := grpc.Invoke(ctx, "/organisations.Organisations/GetOrganisations", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organisationsClient) AddMember(ctx context.Context, in *AddMemberRequest, opts ...grpc.CallOption) (*MemberStatusResponse, error) {
	out := new(MemberStatusResponse)
	err := grpc.Invoke(ctx, "/organisations.Organisations/AddMember", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organisationsClient) RemoveMember(ctx context.Context, in *RemoveMemberRequest, opts ...grpc.CallOption) (*MemberStatusResponse, error) {
	out := new(MemberStatusResponse)
	err := grpc.Invoke(ctx, "/organisations.Organisations/RemoveMember", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organisationsClient) GetMembers(ctx context.Context, in *GetMembersRequest, opts ...grpc.CallOption) (*MembersResponse, error) {
	out := new(MembersResponse)
	err := grpc.Invoke(ctx, "/organisations.Organisations/GetMembers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organisationsClient) Invite(ctx context.Context, in *InviteRequest, opts ...grpc.CallOption) (*InviteStatusResponse, error) {
	out := new(InviteStatusResponse)
	err := grpc.Invoke(ctx, "/organisations.Organisations/Invite", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organisationsClient) ActionInvite(ctx context.Context, in *ActionInviteRequest, opts ...grpc.CallOption) (*MemberStatusResponse, error) {
	out := new(MemberStatusResponse)
	err := grpc.Invoke(ctx, "/organisations.Organisations/ActionInvite", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Organisations service

type OrganisationsServer interface {
	Create(context.Context, *CreateRequest) (*OrganisationResponse, error)
	Retrieve(context.Context, *RetrieveRequest) (*OrganisationResponse, error)
	GetOrganisations(context.Context, *GetOrganisationsRequest) (*OrganisationsResponse, error)
	AddMember(context.Context, *AddMemberRequest) (*MemberStatusResponse, error)
	RemoveMember(context.Context, *RemoveMemberRequest) (*MemberStatusResponse, error)
	GetMembers(context.Context, *GetMembersRequest) (*MembersResponse, error)
	Invite(context.Context, *InviteRequest) (*InviteStatusResponse, error)
	ActionInvite(context.Context, *ActionInviteRequest) (*MemberStatusResponse, error)
}

func RegisterOrganisationsServer(s *grpc.Server, srv OrganisationsServer) {
	s.RegisterService(&_Organisations_serviceDesc, srv)
}

func _Organisations_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganisationsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organisations.Organisations/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganisationsServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organisations_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganisationsServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organisations.Organisations/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganisationsServer).Retrieve(ctx, req.(*RetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organisations_GetOrganisations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrganisationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganisationsServer).GetOrganisations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organisations.Organisations/GetOrganisations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganisationsServer).GetOrganisations(ctx, req.(*GetOrganisationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organisations_AddMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganisationsServer).AddMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organisations.Organisations/AddMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganisationsServer).AddMember(ctx, req.(*AddMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organisations_RemoveMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganisationsServer).RemoveMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organisations.Organisations/RemoveMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganisationsServer).RemoveMember(ctx, req.(*RemoveMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organisations_GetMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganisationsServer).GetMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organisations.Organisations/GetMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganisationsServer).GetMembers(ctx, req.(*GetMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organisations_Invite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganisationsServer).Invite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organisations.Organisations/Invite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganisationsServer).Invite(ctx, req.(*InviteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organisations_ActionInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionInviteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganisationsServer).ActionInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organisations.Organisations/ActionInvite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganisationsServer).ActionInvite(ctx, req.(*ActionInviteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Organisations_serviceDesc = grpc.ServiceDesc{
	ServiceName: "organisations.Organisations",
	HandlerType: (*OrganisationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Organisations_Create_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _Organisations_Retrieve_Handler,
		},
		{
			MethodName: "GetOrganisations",
			Handler:    _Organisations_GetOrganisations_Handler,
		},
		{
			MethodName: "AddMember",
			Handler:    _Organisations_AddMember_Handler,
		},
		{
			MethodName: "RemoveMember",
			Handler:    _Organisations_RemoveMember_Handler,
		},
		{
			MethodName: "GetMembers",
			Handler:    _Organisations_GetMembers_Handler,
		},
		{
			MethodName: "Invite",
			Handler:    _Organisations_Invite_Handler,
		},
		{
			MethodName: "ActionInvite",
			Handler:    _Organisations_ActionInvite_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("organisations.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 949 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x56, 0x5f, 0x6f, 0xdb, 0x54,
	0x14, 0xc7, 0x89, 0xed, 0x36, 0xa7, 0x69, 0x6a, 0x6e, 0x3b, 0x11, 0x32, 0xb4, 0x0e, 0x0f, 0x01,
	0x2a, 0x5a, 0x26, 0x65, 0x68, 0xfc, 0xd1, 0x5e, 0xd2, 0xc4, 0x54, 0x81, 0xd5, 0x0e, 0x37, 0x09,
	0x9d, 0x84, 0x50, 0x71, 0x9b, 0xbb, 0x62, 0x88, 0xe3, 0x60, 0xbb, 0x11, 0xf9, 0x4e, 0xbc, 0xec,
	0x11, 0xf1, 0x0d, 0xf8, 0x0e, 0x3c, 0xf0, 0xc8, 0x33, 0x5f, 0x80, 0xe3, 0x6b, 0x3b, 0xf1, 0xbf,
	0x6d, 0x06, 0x81, 0xb6, 0x37, 0xdf, 0x73, 0x7e, 0xe7, 0xf8, 0xfc, 0xf9, 0xdd, 0x7b, 0x0e, 0xec,
	0x3b, 0xee, 0x95, 0x39, 0xb7, 0x3c, 0xd3, 0xb7, 0x9c, 0xb9, 0xd7, 0x5e, 0xb8, 0x8e, 0xef, 0x90,
	0xdd, 0x94, 0xb0, 0x75, 0x78, 0xe5, 0x38, 0x57, 0x33, 0x76, 0x8f, 0x2b, 0x2f, 0xae, 0x9f, 0xdc,
	0xf3, 0x2d, 0x9b, 0x79, 0xbe, 0x69, 0x2f, 0x42, 0xbc, 0xfa, 0xb3, 0x00, 0xbb, 0x3d, 0x97, 0x99,
	0x3e, 0xa3, 0xec, 0xc7, 0x6b, 0x54, 0x91, 0x06, 0x54, 0xac, 0x69, 0x53, 0xb8, 0x2d, 0xbc, 0x5f,
	0xa3, 0xf8, 0x45, 0x08, 0x88, 0x73, 0xd3, 0x66, 0xcd, 0x0a, 0x97, 0xf0, 0x6f, 0x72, 0x08, 0x3b,
	0x2e, 0xf3, 0xd8, 0x6c, 0xc6, 0xdc, 0x73, 0x04, 0x57, 0xb9, 0x0a, 0x62, 0xd1, 0x60, 0x4a, 0xde,
	0x03, 0xd1, 0x76, 0xa6, 0xac, 0x29, 0xa2, 0xa6, 0xd1, 0xd9, 0x6f, 0xa7, 0x43, 0x3d, 0x45, 0x15,
	0xe5, 0x00, 0x72, 0x17, 0x64, 0x97, 0x5d, 0xa1, 0xb4, 0x29, 0x71, 0xe8, 0x8d, 0x0c, 0x94, 0x72,
	0x25, 0x8d, 0x40, 0xea, 0xdb, 0xb0, 0x47, 0x99, 0xef, 0x5a, 0x6c, 0xf9, 0xac, 0x78, 0xd5, 0x3f,
	0x05, 0x38, 0x30, 0x12, 0x3e, 0x28, 0xf3, 0x16, 0xe8, 0x89, 0xbd, 0xd2, 0x89, 0x91, 0x0f, 0x61,
	0xeb, 0x92, 0xb7, 0x61, 0xda, 0x94, 0x11, 0xbf, 0xd3, 0x69, 0xb5, 0xc3, 0xd6, 0xb5, 0xe3, 0xd6,
	0xb5, 0xc7, 0x71, 0xeb, 0x68, 0x0c, 0x55, 0x7f, 0x15, 0x40, 0xe9, 0x4e, 0xa7, 0xa7, 0xcc, 0xbe,
	0x60, 0x6e, 0x5c, 0x90, 0x9b, 0x50, 0xb3, 0xb9, 0xe0, 0x7c, 0x9d, 0xee, 0x76, 0x28, 0xe0, 0xf1,
	0xef, 0x25, 0xe3, 0x08, 0x20, 0x61, 0xfe, 0x8d, 0xa4, 0x18, 0x81, 0x77, 0x41, 0xf4, 0x57, 0x0b,
	0xc6, 0x4b, 0xd0, 0xe8, 0xbc, 0x99, 0x4d, 0x94, 0xfb, 0x1b, 0x23, 0x80, 0x72, 0x18, 0xe9, 0x80,
	0xcc, 0x7e, 0x5a, 0x58, 0xee, 0x8a, 0x57, 0xe6, 0xf9, 0xe1, 0x47, 0x48, 0xf5, 0x6b, 0xd8, 0xa7,
	0xcc, 0x76, 0x96, 0xec, 0x7f, 0x88, 0x5f, 0x7d, 0x8a, 0xc4, 0x1e, 0xcc, 0x97, 0xd6, 0x86, 0xd8,
	0x07, 0x20, 0x31, 0xdb, 0xb4, 0x66, 0x91, 0xcf, 0xf0, 0xf0, 0x4a, 0x15, 0xe4, 0x2f, 0xa4, 0x6e,
	0x18, 0xf3, 0xc8, 0x37, 0xfd, 0x6b, 0x6f, 0x4d, 0xdd, 0xe2, 0xd0, 0xef, 0x83, 0xec, 0x71, 0x1c,
	0x8f, 0xb8, 0xd1, 0xb9, 0x99, 0x89, 0x29, 0xe5, 0x2a, 0x82, 0x12, 0x15, 0xea, 0x4b, 0xe6, 0x5a,
	0x4f, 0xac, 0x4b, 0x8e, 0x8a, 0x28, 0x9e, 0x92, 0x25, 0xc9, 0x28, 0x96, 0x26, 0x23, 0x79, 0x00,
	0xdb, 0xe6, 0x65, 0x60, 0x8f, 0x66, 0xd2, 0x0b, 0xcd, 0xd6, 0x58, 0xf5, 0x37, 0x01, 0xf6, 0xbb,
	0xfc, 0xf0, 0x9f, 0xf6, 0x2b, 0x45, 0xa3, 0x6a, 0x86, 0x46, 0x9b, 0xd2, 0x89, 0xff, 0xbe, 0x74,
	0x52, 0xbe, 0x74, 0xea, 0x1f, 0x15, 0x38, 0x08, 0xb9, 0x90, 0x69, 0xe1, 0x4b, 0xb9, 0x95, 0xd8,
	0x48, 0x8b, 0x67, 0x52, 0xaa, 0x91, 0x11, 0x34, 0x41, 0x5d, 0xa9, 0x2c, 0x75, 0x03, 0x9b, 0xef,
	0x1d, 0x6b, 0x5e, 0xea, 0xf9, 0x8a, 0x90, 0x01, 0x61, 0xa6, 0xcc, 0x37, 0x2f, 0xbf, 0x43, 0xab,
	0xad, 0x17, 0x13, 0x26, 0xc6, 0xaa, 0x0f, 0xe1, 0xf5, 0x13, 0xe6, 0x87, 0xc9, 0x7a, 0x31, 0x5b,
	0x0a, 0x4a, 0x28, 0x14, 0x3e, 0x0c, 0x0f, 0xe0, 0x0d, 0xb4, 0x4e, 0x4e, 0x08, 0xaf, 0xcc, 0xcb,
	0xa3, 0xfe, 0x22, 0xc0, 0xde, 0xfa, 0x9f, 0x51, 0x53, 0x35, 0xd8, 0x0a, 0xf5, 0x1e, 0xc2, 0xab,
	0x98, 0xc0, 0x07, 0x85, 0x1d, 0x59, 0x1b, 0xc4, 0x67, 0x6d, 0xee, 0xbb, 0x2b, 0x1a, 0xdb, 0xb6,
	0xce, 0xa1, 0x9e, 0x54, 0x10, 0x05, 0xaa, 0x3f, 0xb0, 0x55, 0x14, 0x41, 0xf0, 0x49, 0x3e, 0x01,
	0x69, 0x69, 0xce, 0xae, 0xc3, 0x61, 0xb5, 0xd3, 0xb9, 0x53, 0xf8, 0x9b, 0x34, 0xe3, 0x68, 0x68,
	0xf1, 0x69, 0xe5, 0x63, 0x41, 0xfd, 0x5d, 0x80, 0x1b, 0x99, 0x8c, 0xa3, 0x0c, 0xbe, 0x81, 0xf4,
	0xc6, 0x10, 0xe5, 0xf1, 0x51, 0xe6, 0x07, 0x85, 0xc6, 0x69, 0x69, 0x98, 0x53, 0x66, 0xff, 0x60,
	0x40, 0xf2, 0xa0, 0x7f, 0x9e, 0x5f, 0xd1, 0x3c, 0x4f, 0xe4, 0x77, 0x74, 0x07, 0xc4, 0x60, 0xf4,
	0x92, 0x1d, 0xd8, 0x1a, 0x75, 0xf5, 0xfe, 0xb1, 0xf1, 0x58, 0x79, 0x0d, 0xe7, 0x3d, 0x0c, 0xa9,
	0xd1, 0x9f, 0xf4, 0xc6, 0x03, 0x43, 0x57, 0x84, 0x23, 0xa4, 0x68, 0x38, 0x74, 0x09, 0x80, 0x7c,
	0xf2, 0xc8, 0x38, 0xee, 0x3e, 0x42, 0x94, 0x0c, 0x95, 0xc9, 0x48, 0x11, 0x02, 0x99, 0x36, 0xa1,
	0xc6, 0x50, 0x53, 0x2a, 0x64, 0x1b, 0xc4, 0xee, 0x68, 0xd0, 0x55, 0xaa, 0x47, 0x0f, 0x01, 0x36,
	0x97, 0x8a, 0xd4, 0x40, 0x32, 0xce, 0x74, 0x8d, 0xa2, 0x19, 0xc2, 0x4f, 0xb5, 0xd3, 0x63, 0xfc,
	0xe6, 0xa6, 0x5f, 0x0d, 0xb4, 0x33, 0xfc, 0xae, 0xf0, 0x08, 0x26, 0xc3, 0xa1, 0x41, 0xc7, 0x68,
	0x3d, 0x82, 0x7a, 0xf2, 0x21, 0x09, 0x94, 0x43, 0x4d, 0xef, 0x0f, 0xf4, 0x13, 0xf4, 0x50, 0x87,
	0xed, 0x6e, 0xaf, 0xa7, 0x0d, 0xc7, 0x5a, 0x1f, 0x7d, 0xe0, 0x89, 0x6a, 0x9f, 0x6b, 0xbd, 0xe0,
	0xc4, 0xbd, 0x68, 0x8f, 0x87, 0x03, 0x8a, 0x87, 0x2a, 0xd9, 0x85, 0x9a, 0x6e, 0x8c, 0xcf, 0x3f,
	0x33, 0x26, 0x7a, 0x5f, 0x11, 0x3b, 0x4f, 0x25, 0xd8, 0x4d, 0xd5, 0x94, 0x7c, 0x01, 0x72, 0xb8,
	0xc2, 0x91, 0xb7, 0x32, 0x75, 0x4b, 0x6d, 0x76, 0xad, 0x32, 0x55, 0x25, 0x06, 0x06, 0x12, 0x6d,
	0x58, 0xe4, 0x56, 0x6e, 0x67, 0x49, 0xad, 0x5e, 0xe5, 0x1c, 0x7e, 0x0b, 0x4a, 0xf6, 0xbe, 0x91,
	0x77, 0x33, 0x86, 0xcf, 0xb8, 0x90, 0xad, 0x77, 0xca, 0xd0, 0x90, 0x7c, 0x09, 0xb5, 0xf5, 0x12,
	0x44, 0x0e, 0x33, 0x26, 0xd9, 0xf5, 0xa8, 0x55, 0xe6, 0xee, 0x90, 0x33, 0xa8, 0x27, 0x57, 0x13,
	0xa2, 0xe6, 0x2a, 0x91, 0xdb, 0x5b, 0xca, 0x39, 0xd6, 0x01, 0x36, 0x6f, 0x17, 0xb9, 0x9d, 0xaf,
	0x43, 0xfa, 0x59, 0x6b, 0xdd, 0x7a, 0xfe, 0x83, 0x12, 0xf4, 0x3e, 0xa4, 0x58, 0xae, 0xf7, 0xa9,
	0x61, 0x9a, 0x0b, 0xae, 0x70, 0xcd, 0xc0, 0xac, 0x93, 0x83, 0x38, 0x97, 0x75, 0xc1, 0x94, 0x2e,
	0x95, 0xf5, 0x85, 0xcc, 0xdf, 0xf3, 0xfb, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x13, 0x06, 0x59,
	0x24, 0xb3, 0x0c, 0x00, 0x00,
}
