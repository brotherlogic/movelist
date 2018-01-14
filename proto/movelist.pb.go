// Code generated by protoc-gen-go. DO NOT EDIT.
// source: movelist.proto

/*
Package movelist is a generated protocol buffer package.

It is generated from these files:
	movelist.proto

It has these top-level messages:
	RecordMove
*/
package movelist

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

type RecordMove struct {
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *RecordMove) Reset()                    { *m = RecordMove{} }
func (m *RecordMove) String() string            { return proto.CompactTextString(m) }
func (*RecordMove) ProtoMessage()               {}
func (*RecordMove) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RecordMove) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*RecordMove)(nil), "movelist.RecordMove")
}

func init() { proto.RegisterFile("movelist.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 85 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xcd, 0x2f, 0x4b,
	0xcd, 0xc9, 0x2c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0xb4,
	0xb8, 0xb8, 0x82, 0x52, 0x93, 0xf3, 0x8b, 0x52, 0x7c, 0xf3, 0xcb, 0x52, 0x85, 0x64, 0xb8, 0x38,
	0x4b, 0x32, 0x73, 0x53, 0x8b, 0x4b, 0x12, 0x73, 0x0b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83,
	0x10, 0x02, 0x49, 0x6c, 0x60, 0xcd, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x76, 0x89,
	0x01, 0x4e, 0x00, 0x00, 0x00,
}