// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userpass.proto

/*
Package catenasup is a generated protocol buffer package.

It is generated from these files:
	userpass.proto

It has these top-level messages:
	UserPass
	User
	Status
*/
package catenasup

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

type Status_StatusType int32

const (
	Status_SUCCESS Status_StatusType = 0
	Status_FAIL    Status_StatusType = 1
)

var Status_StatusType_name = map[int32]string{
	0: "SUCCESS",
	1: "FAIL",
}
var Status_StatusType_value = map[string]int32{
	"SUCCESS": 0,
	"FAIL":    1,
}

func (x Status_StatusType) String() string {
	return proto.EnumName(Status_StatusType_name, int32(x))
}
func (Status_StatusType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type UserPass struct {
	User string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Pass string `protobuf:"bytes,2,opt,name=pass" json:"pass,omitempty"`
}

func (m *UserPass) Reset()                    { *m = UserPass{} }
func (m *UserPass) String() string            { return proto.CompactTextString(m) }
func (*UserPass) ProtoMessage()               {}
func (*UserPass) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserPass) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *UserPass) GetPass() string {
	if m != nil {
		return m.Pass
	}
	return ""
}

type User struct {
	User string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *User) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type Status struct {
	Result Status_StatusType `protobuf:"varint,1,opt,name=result,enum=catenasup.Status_StatusType" json:"result,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Status) GetResult() Status_StatusType {
	if m != nil {
		return m.Result
	}
	return Status_SUCCESS
}

func init() {
	proto.RegisterType((*UserPass)(nil), "catenasup.UserPass")
	proto.RegisterType((*User)(nil), "catenasup.User")
	proto.RegisterType((*Status)(nil), "catenasup.Status")
	proto.RegisterEnum("catenasup.Status_StatusType", Status_StatusType_name, Status_StatusType_value)
}

func init() { proto.RegisterFile("userpass.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x2d, 0x4e, 0x2d,
	0x2a, 0x48, 0x2c, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0x4e, 0x2c, 0x49,
	0xcd, 0x4b, 0x2c, 0x2e, 0x2d, 0x50, 0x32, 0xe2, 0xe2, 0x08, 0x2d, 0x4e, 0x2d, 0x0a, 0x48, 0x2c,
	0x2e, 0x16, 0x12, 0xe2, 0x62, 0x01, 0x29, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3,
	0x41, 0x62, 0x20, 0x8d, 0x12, 0x4c, 0x10, 0x31, 0x10, 0x5b, 0x49, 0x8a, 0x8b, 0x25, 0x14, 0x2a,
	0x87, 0xae, 0x5e, 0x29, 0x99, 0x8b, 0x2d, 0xb8, 0x24, 0xb1, 0xa4, 0xb4, 0x58, 0xc8, 0x84, 0x8b,
	0xad, 0x28, 0xb5, 0xb8, 0x34, 0xa7, 0x04, 0x2c, 0xcf, 0x67, 0x24, 0xa3, 0x07, 0xb7, 0x55, 0x0f,
	0xa2, 0x04, 0x4a, 0x85, 0x54, 0x16, 0xa4, 0x06, 0x41, 0xd5, 0x2a, 0x29, 0x73, 0x71, 0x21, 0x44,
	0x85, 0xb8, 0xb9, 0xd8, 0x83, 0x43, 0x9d, 0x9d, 0x5d, 0x83, 0x83, 0x05, 0x18, 0x84, 0x38, 0xb8,
	0x58, 0xdc, 0x1c, 0x3d, 0x7d, 0x04, 0x18, 0x8d, 0xde, 0x30, 0x72, 0xf1, 0x39, 0x83, 0x0d, 0x83,
	0xbb, 0xdd, 0x98, 0x8b, 0xdd, 0x31, 0x25, 0x05, 0xec, 0x2c, 0x61, 0x24, 0x8b, 0x60, 0xf2, 0x52,
	0x82, 0x18, 0xb6, 0x2b, 0x31, 0x08, 0x99, 0x71, 0x71, 0x39, 0x67, 0x24, 0xe6, 0xa5, 0xa7, 0x82,
	0x8d, 0x20, 0x5e, 0x9f, 0x25, 0x17, 0xaf, 0x73, 0x46, 0x6a, 0x72, 0x36, 0xdc, 0x76, 0xe2, 0xb5,
	0x1a, 0x71, 0x71, 0xb9, 0xa4, 0xe6, 0xa4, 0x96, 0xa4, 0x82, 0x9d, 0xca, 0x8f, 0xa6, 0x0f, 0xab,
	0x9e, 0x24, 0x36, 0x70, 0xac, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x24, 0x54, 0x34,
	0xc7, 0x01, 0x00, 0x00,
}
