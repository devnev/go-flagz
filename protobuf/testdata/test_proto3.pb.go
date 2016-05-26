// Code generated by protoc-gen-go.
// source: test_proto3.proto
// DO NOT EDIT!

/*
Package mwitkow_testproto is a generated protocol buffer package.

It is generated from these files:
	test_proto3.proto

It has these top-level messages:
	SomeMsg
*/
package mwitkow_testproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type SomeEnum int32

const (
	SomeEnum_OPT_1 SomeEnum = 0
	SomeEnum_OPT_2 SomeEnum = 1
)

var SomeEnum_name = map[int32]string{
	0: "OPT_1",
	1: "OPT_2",
}
var SomeEnum_value = map[string]int32{
	"OPT_1": 0,
	"OPT_2": 1,
}

func (x SomeEnum) String() string {
	return proto.EnumName(SomeEnum_name, int32(x))
}
func (SomeEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SomeMsg struct {
	SomeString string           `protobuf:"bytes,1,opt,name=some_string,json=someString" json:"some_string,omitempty"`
	SomeEnum   SomeEnum         `protobuf:"varint,2,opt,name=some_enum,json=someEnum,enum=mwitkow.testproto.SomeEnum" json:"some_enum,omitempty"`
	SomeMap    map[string]int32 `protobuf:"bytes,3,rep,name=some_map,json=someMap" json:"some_map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *SomeMsg) Reset()                    { *m = SomeMsg{} }
func (m *SomeMsg) String() string            { return proto.CompactTextString(m) }
func (*SomeMsg) ProtoMessage()               {}
func (*SomeMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SomeMsg) GetSomeMap() map[string]int32 {
	if m != nil {
		return m.SomeMap
	}
	return nil
}

func init() {
	proto.RegisterType((*SomeMsg)(nil), "mwitkow.testproto.SomeMsg")
	proto.RegisterEnum("mwitkow.testproto.SomeEnum", SomeEnum_name, SomeEnum_value)
}

var fileDescriptor0 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x49, 0x2d, 0x2e,
	0x89, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x37, 0xd6, 0x03, 0x53, 0x42, 0x82, 0xb9, 0xe5, 0x99, 0x25,
	0xd9, 0xf9, 0xe5, 0x7a, 0x20, 0x29, 0xb0, 0x90, 0xd2, 0x13, 0x46, 0x2e, 0xf6, 0xe0, 0xfc, 0xdc,
	0x54, 0xdf, 0xe2, 0x74, 0x21, 0x79, 0x2e, 0xee, 0x62, 0x20, 0x33, 0xbe, 0xb8, 0xa4, 0x28, 0x33,
	0x2f, 0x5d, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88, 0x0b, 0x24, 0x14, 0x0c, 0x16, 0x11, 0xb2,
	0xe0, 0xe2, 0x04, 0x2b, 0x48, 0xcd, 0x2b, 0xcd, 0x95, 0x60, 0x02, 0x4a, 0xf3, 0x19, 0x49, 0xeb,
	0x61, 0x98, 0xa9, 0x07, 0x32, 0xcf, 0x15, 0xa8, 0x24, 0x88, 0xa3, 0x18, 0xca, 0x12, 0x72, 0xe2,
	0x02, 0xb3, 0xe3, 0x73, 0x13, 0x0b, 0x24, 0x98, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0xd4, 0x71, 0x68,
	0x04, 0x3a, 0x04, 0x42, 0x27, 0x16, 0xb8, 0xe6, 0x95, 0x14, 0x55, 0x06, 0xb1, 0x17, 0x43, 0x78,
	0x52, 0x56, 0x5c, 0x3c, 0xc8, 0x12, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x50, 0x67, 0x82,
	0x98, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x60, 0xb7, 0xb1, 0x06, 0x41, 0x38,
	0x56, 0x4c, 0x16, 0x8c, 0x5a, 0x0a, 0x5c, 0x1c, 0x30, 0x57, 0x09, 0x71, 0x72, 0xb1, 0xfa, 0x07,
	0x84, 0xc4, 0x1b, 0x0a, 0x30, 0xc0, 0x98, 0x46, 0x02, 0x8c, 0x49, 0x6c, 0x90, 0x90, 0x02, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x2f, 0xf1, 0x31, 0x8c, 0x37, 0x01, 0x00, 0x00,
}
