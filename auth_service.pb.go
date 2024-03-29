// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth_service.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateAccessTokenRequest struct {
	Identity             string   `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Room                 string   `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccessTokenRequest) Reset()         { *m = CreateAccessTokenRequest{} }
func (m *CreateAccessTokenRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAccessTokenRequest) ProtoMessage()    {}
func (*CreateAccessTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f39bb026ca10b68, []int{0}
}

func (m *CreateAccessTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccessTokenRequest.Unmarshal(m, b)
}
func (m *CreateAccessTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccessTokenRequest.Marshal(b, m, deterministic)
}
func (m *CreateAccessTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccessTokenRequest.Merge(m, src)
}
func (m *CreateAccessTokenRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAccessTokenRequest.Size(m)
}
func (m *CreateAccessTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccessTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccessTokenRequest proto.InternalMessageInfo

func (m *CreateAccessTokenRequest) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *CreateAccessTokenRequest) GetRoom() string {
	if m != nil {
		return m.Room
	}
	return ""
}

type CreateAccessTokenResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccessTokenResponse) Reset()         { *m = CreateAccessTokenResponse{} }
func (m *CreateAccessTokenResponse) String() string { return proto.CompactTextString(m) }
func (*CreateAccessTokenResponse) ProtoMessage()    {}
func (*CreateAccessTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f39bb026ca10b68, []int{1}
}

func (m *CreateAccessTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccessTokenResponse.Unmarshal(m, b)
}
func (m *CreateAccessTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccessTokenResponse.Marshal(b, m, deterministic)
}
func (m *CreateAccessTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccessTokenResponse.Merge(m, src)
}
func (m *CreateAccessTokenResponse) XXX_Size() int {
	return xxx_messageInfo_CreateAccessTokenResponse.Size(m)
}
func (m *CreateAccessTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccessTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccessTokenResponse proto.InternalMessageInfo

func (m *CreateAccessTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateAccessTokenRequest)(nil), "org.jetrtc.api.CreateAccessTokenRequest")
	proto.RegisterType((*CreateAccessTokenResponse)(nil), "org.jetrtc.api.CreateAccessTokenResponse")
}

func init() { proto.RegisterFile("auth_service.proto", fileDescriptor_0f39bb026ca10b68) }

var fileDescriptor_0f39bb026ca10b68 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x2c, 0x2d, 0xc9,
	0x88, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0xcb, 0x2f, 0x4a, 0xd7, 0xcb, 0x4a, 0x2d, 0x29, 0x2a, 0x49, 0xd6, 0x4b, 0x2c, 0xc8, 0x54, 0xf2,
	0xe2, 0x92, 0x70, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x75, 0x4c, 0x4e, 0x4e, 0x2d, 0x2e, 0x0e, 0xc9,
	0xcf, 0x4e, 0xcd, 0x0b, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe2, 0xe2, 0xc8, 0x4c,
	0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x85,
	0x84, 0xb8, 0x58, 0x8a, 0xf2, 0xf3, 0x73, 0x25, 0x98, 0xc0, 0xe2, 0x60, 0xb6, 0x92, 0x21, 0x97,
	0x24, 0x16, 0xb3, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x44, 0xb8, 0x58, 0x4b, 0x40, 0x02,
	0x50, 0x93, 0x20, 0x1c, 0x27, 0xd6, 0x28, 0xe6, 0xc4, 0x82, 0xcc, 0x24, 0x36, 0xb0, 0xe3, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x72, 0xfc, 0xbd, 0xc3, 0xb2, 0x00, 0x00, 0x00,
}
