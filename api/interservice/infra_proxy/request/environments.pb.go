// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/infra_proxy/request/environments.proto

package request

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

type Environments struct {
	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId             string   `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Environments) Reset()         { *m = Environments{} }
func (m *Environments) String() string { return proto.CompactTextString(m) }
func (*Environments) ProtoMessage()    {}
func (*Environments) Descriptor() ([]byte, []int) {
	return fileDescriptor_e190c707558fe428, []int{0}
}

func (m *Environments) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Environments.Unmarshal(m, b)
}
func (m *Environments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Environments.Marshal(b, m, deterministic)
}
func (m *Environments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Environments.Merge(m, src)
}
func (m *Environments) XXX_Size() int {
	return xxx_messageInfo_Environments.Size(m)
}
func (m *Environments) XXX_DiscardUnknown() {
	xxx_messageInfo_Environments.DiscardUnknown(m)
}

var xxx_messageInfo_Environments proto.InternalMessageInfo

func (m *Environments) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *Environments) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

type Environment struct {
	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Environment name.
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Environment) Reset()         { *m = Environment{} }
func (m *Environment) String() string { return proto.CompactTextString(m) }
func (*Environment) ProtoMessage()    {}
func (*Environment) Descriptor() ([]byte, []int) {
	return fileDescriptor_e190c707558fe428, []int{1}
}

func (m *Environment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Environment.Unmarshal(m, b)
}
func (m *Environment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Environment.Marshal(b, m, deterministic)
}
func (m *Environment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Environment.Merge(m, src)
}
func (m *Environment) XXX_Size() int {
	return xxx_messageInfo_Environment.Size(m)
}
func (m *Environment) XXX_DiscardUnknown() {
	xxx_messageInfo_Environment.DiscardUnknown(m)
}

var xxx_messageInfo_Environment proto.InternalMessageInfo

func (m *Environment) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *Environment) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *Environment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateEnvironment struct {
	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Environment name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Environment description.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty" toml:"description,omitempty" mapstructure:"description,omitempty"`
	// Class name.
	JsonClass string `protobuf:"bytes,5,opt,name=json_class,json=jsonClass,proto3" json:"json_class,omitempty" toml:"json_class,omitempty" mapstructure:"json_class,omitempty"`
	// Environment versined cookbooks constraints.
	CookbookVersions map[string]string `protobuf:"bytes,6,rep,name=cookbook_versions,json=cookbookVersions,proto3" json:"cookbook_versions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" toml:"cookbook_versions,omitempty" mapstructure:"cookbook_versions,omitempty"`
	// Stringified json of the default attributes.
	DefaultAttributes string `protobuf:"bytes,7,opt,name=default_attributes,json=defaultAttributes,proto3" json:"default_attributes,omitempty" toml:"default_attributes,omitempty" mapstructure:"default_attributes,omitempty"`
	// Stringified json of the override attributes.
	OverrideAttributes   string   `protobuf:"bytes,8,opt,name=override_attributes,json=overrideAttributes,proto3" json:"override_attributes,omitempty" toml:"override_attributes,omitempty" mapstructure:"override_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *CreateEnvironment) Reset()         { *m = CreateEnvironment{} }
func (m *CreateEnvironment) String() string { return proto.CompactTextString(m) }
func (*CreateEnvironment) ProtoMessage()    {}
func (*CreateEnvironment) Descriptor() ([]byte, []int) {
	return fileDescriptor_e190c707558fe428, []int{2}
}

func (m *CreateEnvironment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEnvironment.Unmarshal(m, b)
}
func (m *CreateEnvironment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEnvironment.Marshal(b, m, deterministic)
}
func (m *CreateEnvironment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEnvironment.Merge(m, src)
}
func (m *CreateEnvironment) XXX_Size() int {
	return xxx_messageInfo_CreateEnvironment.Size(m)
}
func (m *CreateEnvironment) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEnvironment.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEnvironment proto.InternalMessageInfo

func (m *CreateEnvironment) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *CreateEnvironment) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *CreateEnvironment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateEnvironment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateEnvironment) GetJsonClass() string {
	if m != nil {
		return m.JsonClass
	}
	return ""
}

func (m *CreateEnvironment) GetCookbookVersions() map[string]string {
	if m != nil {
		return m.CookbookVersions
	}
	return nil
}

func (m *CreateEnvironment) GetDefaultAttributes() string {
	if m != nil {
		return m.DefaultAttributes
	}
	return ""
}

func (m *CreateEnvironment) GetOverrideAttributes() string {
	if m != nil {
		return m.OverrideAttributes
	}
	return ""
}

func init() {
	proto.RegisterType((*Environments)(nil), "chef.automate.domain.infra_proxy.request.Environments")
	proto.RegisterType((*Environment)(nil), "chef.automate.domain.infra_proxy.request.Environment")
	proto.RegisterType((*CreateEnvironment)(nil), "chef.automate.domain.infra_proxy.request.CreateEnvironment")
	proto.RegisterMapType((map[string]string)(nil), "chef.automate.domain.infra_proxy.request.CreateEnvironment.CookbookVersionsEntry")
}

func init() {
	proto.RegisterFile("api/interservice/infra_proxy/request/environments.proto", fileDescriptor_e190c707558fe428)
}

var fileDescriptor_e190c707558fe428 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xc1, 0x8a, 0x13, 0x41,
	0x10, 0x25, 0x3b, 0x9b, 0xb8, 0xa9, 0x78, 0xd8, 0xb4, 0x2e, 0x0c, 0x8a, 0x10, 0x72, 0xca, 0xc5,
	0x6e, 0xd0, 0x83, 0x22, 0x88, 0xb8, 0xc3, 0x1e, 0x72, 0x34, 0xa0, 0x07, 0x2f, 0x43, 0xcf, 0x4c,
	0x25, 0x69, 0x93, 0xe9, 0x1a, 0xab, 0x7b, 0x06, 0x73, 0xf1, 0xdf, 0xfc, 0x33, 0x99, 0xce, 0x44,
	0x07, 0xf5, 0x10, 0x84, 0xbd, 0x55, 0xbf, 0x57, 0xef, 0xf1, 0xfa, 0x51, 0xf0, 0x4a, 0x57, 0x46,
	0x19, 0xeb, 0x91, 0x1d, 0x72, 0x63, 0x72, 0x54, 0xc6, 0xae, 0x59, 0xa7, 0x15, 0xd3, 0xb7, 0x83,
	0x62, 0xfc, 0x5a, 0xa3, 0xf3, 0x0a, 0x6d, 0x63, 0x98, 0x6c, 0x89, 0xd6, 0x3b, 0x59, 0x31, 0x79,
	0x12, 0x8b, 0x7c, 0x8b, 0x6b, 0xa9, 0x6b, 0x4f, 0xa5, 0xf6, 0x28, 0x0b, 0x2a, 0xb5, 0xb1, 0xb2,
	0x27, 0x96, 0x9d, 0x78, 0x7e, 0x0b, 0x0f, 0xef, 0x7a, 0x7a, 0x71, 0x03, 0x23, 0xe2, 0x4d, 0x6a,
	0x8a, 0x78, 0x30, 0x1b, 0x2c, 0xc6, 0xab, 0x21, 0xf1, 0x66, 0x59, 0x88, 0xa7, 0x30, 0x6e, 0x23,
	0x20, 0xb7, 0xcc, 0x45, 0x60, 0xae, 0x8e, 0xc0, 0xb2, 0x98, 0x7f, 0x84, 0x49, 0xcf, 0xe3, 0x7f,
	0x2c, 0x84, 0x80, 0x4b, 0xab, 0x4b, 0x8c, 0xa3, 0x80, 0x87, 0x79, 0xfe, 0x23, 0x82, 0x69, 0xc2,
	0xa8, 0x3d, 0xde, 0x83, 0xbb, 0x98, 0xc1, 0xa4, 0x40, 0x97, 0xb3, 0xa9, 0xbc, 0x21, 0x1b, 0x5f,
	0x06, 0xaa, 0x0f, 0x89, 0x67, 0x00, 0x5f, 0x1c, 0xd9, 0x34, 0xdf, 0x6b, 0xe7, 0xe2, 0x61, 0x58,
	0x18, 0xb7, 0x48, 0xd2, 0x02, 0xe2, 0x3b, 0x4c, 0x73, 0xa2, 0x5d, 0x46, 0xb4, 0x4b, 0x1b, 0x64,
	0x67, 0xc8, 0xba, 0x78, 0x34, 0x8b, 0x16, 0x93, 0x17, 0x1f, 0xe4, 0xb9, 0xfd, 0xcb, 0xbf, 0x3e,
	0x28, 0x93, 0xce, 0xf4, 0x53, 0xe7, 0x79, 0x67, 0x3d, 0x1f, 0x56, 0xd7, 0xf9, 0x1f, 0xb0, 0x78,
	0x0e, 0xa2, 0xc0, 0xb5, 0xae, 0xf7, 0x3e, 0xd5, 0xde, 0xb3, 0xc9, 0x6a, 0x8f, 0x2e, 0x7e, 0x10,
	0x62, 0x4e, 0x3b, 0xe6, 0xfd, 0x2f, 0x42, 0x28, 0x78, 0x44, 0x0d, 0x32, 0x9b, 0x02, 0xfb, 0xfb,
	0x57, 0x61, 0x5f, 0x9c, 0xa8, 0xdf, 0x82, 0x27, 0x09, 0xdc, 0xfc, 0x33, 0x8a, 0xb8, 0x86, 0x68,
	0x87, 0x87, 0xae, 0xfe, 0x76, 0x14, 0x8f, 0x61, 0xd8, 0xe8, 0x7d, 0x8d, 0x5d, 0xf1, 0xc7, 0xc7,
	0x9b, 0x8b, 0xd7, 0x83, 0xdb, 0x77, 0x9f, 0xdf, 0x6e, 0x8c, 0xdf, 0xd6, 0x99, 0xcc, 0xa9, 0x54,
	0x6d, 0x2b, 0xea, 0xd4, 0x8a, 0x3a, 0xe7, 0xb8, 0xb3, 0x51, 0x38, 0xe8, 0x97, 0x3f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xa8, 0x01, 0x75, 0x18, 0x0b, 0x03, 0x00, 0x00,
}
