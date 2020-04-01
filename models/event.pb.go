// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package proto

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

type Event struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	EventName            string   `protobuf:"bytes,2,opt,name=EventName,proto3" json:"EventName,omitempty"`
	Publisher            string   `protobuf:"bytes,3,opt,name=Publisher,proto3" json:"Publisher,omitempty"`
	Timestamp            int64    `protobuf:"varint,4,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Payload              []byte   `protobuf:"bytes,5,opt,name=Payload,proto3" json:"Payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Event) GetEventName() string {
	if m != nil {
		return m.EventName
	}
	return ""
}

func (m *Event) GetPublisher() string {
	if m != nil {
		return m.Publisher
	}
	return ""
}

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Event) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "proto.Event")
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_2d17a9d3f0ddf27e) }

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2d, 0x4b, 0xcd,
	0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xdd, 0x8c, 0x5c, 0xac,
	0xae, 0x20, 0x61, 0x21, 0x3e, 0x2e, 0x26, 0x4f, 0x17, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0x26, 0x4f, 0x17, 0x21, 0x19, 0x2e, 0x4e, 0xb0, 0x84, 0x5f, 0x62, 0x6e, 0xaa, 0x04, 0x13, 0x58,
	0x18, 0x21, 0x00, 0x92, 0x0d, 0x28, 0x4d, 0xca, 0xc9, 0x2c, 0xce, 0x48, 0x2d, 0x92, 0x60, 0x86,
	0xc8, 0xc2, 0x05, 0x40, 0xb2, 0x21, 0x99, 0xb9, 0xa9, 0xc5, 0x25, 0x89, 0xb9, 0x05, 0x12, 0x2c,
	0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x08, 0x01, 0x21, 0x09, 0x2e, 0xf6, 0x80, 0xc4, 0xca, 0x9c, 0xfc,
	0xc4, 0x14, 0x09, 0x56, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x18, 0x37, 0x89, 0x0d, 0xec, 0x28, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0xf3, 0x50, 0x21, 0xaa, 0x00, 0x00, 0x00,
}