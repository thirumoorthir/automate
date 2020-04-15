// Code generated by protoc-gen-go. DO NOT EDIT.
// source: health.proto

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

type VersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdbebe66dda7cb29, []int{0}
}

func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (m *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(m, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

type VersionResponse struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdbebe66dda7cb29, []int{1}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*VersionRequest)(nil), "notifications.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "notifications.VersionResponse")
}

func init() {
	proto.RegisterFile("health.proto", fileDescriptor_fdbebe66dda7cb29)
}

var fileDescriptor_fdbebe66dda7cb29 = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x48, 0x4d, 0xcc,
	0x29, 0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcd, 0xcb, 0x2f, 0xc9, 0x4c, 0xcb,
	0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x56, 0x12, 0xe0, 0xe2, 0x0b, 0x4b, 0x2d, 0x2a, 0xce,
	0xcc, 0xcf, 0x0b, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0xd2, 0xe6, 0xe2, 0x87, 0x8b, 0x14,
	0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x49, 0x70, 0xb1, 0x97, 0x41, 0x84, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x60, 0x5c, 0x27, 0xd6, 0x28, 0xe6, 0xc4, 0x82, 0xcc, 0x24, 0x36, 0xb0, 0xd9,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0x73, 0x15, 0xb3, 0x6b, 0x00, 0x00, 0x00,
}
