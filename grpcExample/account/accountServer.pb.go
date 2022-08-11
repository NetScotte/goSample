// Code generated by protoc-gen-go. DO NOT EDIT.
// source: accountServer.proto

package account

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

type Id struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_104f608dc9682142, []int{0}
}

func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (m *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(m, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Sex                  string   `protobuf:"bytes,3,opt,name=sex,proto3" json:"sex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_104f608dc9682142, []int{1}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserInfo) GetSex() string {
	if m != nil {
		return m.Sex
	}
	return ""
}

func init() {
	proto.RegisterType((*Id)(nil), "account.Id")
	proto.RegisterType((*UserInfo)(nil), "account.UserInfo")
}

func init() { proto.RegisterFile("accountServer.proto", fileDescriptor_104f608dc9682142) }

var fileDescriptor_104f608dc9682142 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x4c, 0x4e, 0xce,
	0x2f, 0xcd, 0x2b, 0x09, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x87, 0x0a, 0x2a, 0x89, 0x70, 0x31, 0x79, 0xa6, 0x08, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48,
	0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x31, 0x65, 0xa6, 0x28, 0x39, 0x71, 0x71, 0x84, 0x16, 0xa7,
	0x16, 0x79, 0xe6, 0xa5, 0xe5, 0x0b, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x82, 0x65, 0x39,
	0x83, 0xc0, 0x6c, 0x21, 0x01, 0x2e, 0xe6, 0xc4, 0xf4, 0x54, 0x09, 0x26, 0xb0, 0x06, 0x10, 0x13,
	0x24, 0x52, 0x9c, 0x5a, 0x21, 0xc1, 0x0c, 0x56, 0x04, 0x62, 0x1a, 0x39, 0x70, 0xf1, 0x3a, 0x22,
	0xdb, 0x2c, 0xa4, 0xcf, 0xc5, 0xed, 0x9e, 0x5a, 0x02, 0x32, 0xd7, 0xa9, 0xd2, 0x33, 0x45, 0x88,
	0x5b, 0x0f, 0xea, 0x06, 0x3d, 0xcf, 0x14, 0x29, 0x41, 0x38, 0x07, 0x66, 0xaf, 0x12, 0x43, 0x12,
	0x1b, 0xd8, 0xad, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xdf, 0xeb, 0x9c, 0xc2, 0x00,
	0x00, 0x00,
}
