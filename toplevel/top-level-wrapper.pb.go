// Code generated by protoc-gen-go. DO NOT EDIT.
// source: top-level-wrapper.proto

package toplevel

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

type Nested1 struct {
	First                string   `protobuf:"bytes,1,opt,name=first,proto3" json:"first,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nested1) Reset()         { *m = Nested1{} }
func (m *Nested1) String() string { return proto.CompactTextString(m) }
func (*Nested1) ProtoMessage()    {}
func (*Nested1) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe245448f29f3d4c, []int{0}
}

func (m *Nested1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nested1.Unmarshal(m, b)
}
func (m *Nested1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nested1.Marshal(b, m, deterministic)
}
func (m *Nested1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nested1.Merge(m, src)
}
func (m *Nested1) XXX_Size() int {
	return xxx_messageInfo_Nested1.Size(m)
}
func (m *Nested1) XXX_DiscardUnknown() {
	xxx_messageInfo_Nested1.DiscardUnknown(m)
}

var xxx_messageInfo_Nested1 proto.InternalMessageInfo

func (m *Nested1) GetFirst() string {
	if m != nil {
		return m.First
	}
	return ""
}

type Nested2 struct {
	Second               string   `protobuf:"bytes,1,opt,name=second,proto3" json:"second,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nested2) Reset()         { *m = Nested2{} }
func (m *Nested2) String() string { return proto.CompactTextString(m) }
func (*Nested2) ProtoMessage()    {}
func (*Nested2) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe245448f29f3d4c, []int{1}
}

func (m *Nested2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nested2.Unmarshal(m, b)
}
func (m *Nested2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nested2.Marshal(b, m, deterministic)
}
func (m *Nested2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nested2.Merge(m, src)
}
func (m *Nested2) XXX_Size() int {
	return xxx_messageInfo_Nested2.Size(m)
}
func (m *Nested2) XXX_DiscardUnknown() {
	xxx_messageInfo_Nested2.DiscardUnknown(m)
}

var xxx_messageInfo_Nested2 proto.InternalMessageInfo

func (m *Nested2) GetSecond() string {
	if m != nil {
		return m.Second
	}
	return ""
}

type TopWrapper struct {
	Foo                  string   `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	Bar                  int64    `protobuf:"varint,2,opt,name=bar,proto3" json:"bar,omitempty"`
	Baz                  bool     `protobuf:"varint,3,opt,name=baz,proto3" json:"baz,omitempty"`
	Nested1              *Nested1 `protobuf:"bytes,4,opt,name=nested1,proto3" json:"nested1,omitempty"`
	Nested2              *Nested2 `protobuf:"bytes,5,opt,name=nested2,proto3" json:"nested2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopWrapper) Reset()         { *m = TopWrapper{} }
func (m *TopWrapper) String() string { return proto.CompactTextString(m) }
func (*TopWrapper) ProtoMessage()    {}
func (*TopWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe245448f29f3d4c, []int{2}
}

func (m *TopWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopWrapper.Unmarshal(m, b)
}
func (m *TopWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopWrapper.Marshal(b, m, deterministic)
}
func (m *TopWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopWrapper.Merge(m, src)
}
func (m *TopWrapper) XXX_Size() int {
	return xxx_messageInfo_TopWrapper.Size(m)
}
func (m *TopWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_TopWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_TopWrapper proto.InternalMessageInfo

func (m *TopWrapper) GetFoo() string {
	if m != nil {
		return m.Foo
	}
	return ""
}

func (m *TopWrapper) GetBar() int64 {
	if m != nil {
		return m.Bar
	}
	return 0
}

func (m *TopWrapper) GetBaz() bool {
	if m != nil {
		return m.Baz
	}
	return false
}

func (m *TopWrapper) GetNested1() *Nested1 {
	if m != nil {
		return m.Nested1
	}
	return nil
}

func (m *TopWrapper) GetNested2() *Nested2 {
	if m != nil {
		return m.Nested2
	}
	return nil
}

func init() {
	proto.RegisterType((*Nested1)(nil), "toplevel.Nested1")
	proto.RegisterType((*Nested2)(nil), "toplevel.Nested2")
	proto.RegisterType((*TopWrapper)(nil), "toplevel.TopWrapper")
}

func init() { proto.RegisterFile("top-level-wrapper.proto", fileDescriptor_fe245448f29f3d4c) }

var fileDescriptor_fe245448f29f3d4c = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0xc9, 0x2f, 0xd0,
	0xcd, 0x49, 0x2d, 0x4b, 0xcd, 0xd1, 0x2d, 0x2f, 0x4a, 0x2c, 0x28, 0x48, 0x2d, 0xd2, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0xc9, 0x2f, 0x00, 0x8b, 0x2b, 0xc9, 0x73, 0xb1, 0xfb, 0xa5,
	0x16, 0x97, 0xa4, 0xa6, 0x18, 0x0a, 0x89, 0x70, 0xb1, 0xa6, 0x65, 0x16, 0x15, 0x97, 0x48, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x4a, 0x8a, 0x30, 0x05, 0x46, 0x42, 0x62, 0x5c, 0x6c,
	0xc5, 0xa9, 0xc9, 0xf9, 0x79, 0x29, 0x50, 0x15, 0x50, 0x9e, 0xd2, 0x1c, 0x46, 0x2e, 0xae, 0x90,
	0xfc, 0x82, 0x70, 0x88, 0x15, 0x42, 0x02, 0x5c, 0xcc, 0x69, 0xf9, 0xf9, 0x50, 0x35, 0x20, 0x26,
	0x48, 0x24, 0x29, 0xb1, 0x48, 0x82, 0x49, 0x81, 0x51, 0x83, 0x39, 0x08, 0xc4, 0x84, 0x88, 0x54,
	0x49, 0x30, 0x2b, 0x30, 0x6a, 0x70, 0x80, 0x44, 0xaa, 0x84, 0xb4, 0xb9, 0xd8, 0xf3, 0x20, 0x0e,
	0x91, 0x60, 0x51, 0x60, 0xd4, 0xe0, 0x36, 0x12, 0xd4, 0x83, 0x39, 0x52, 0x0f, 0xea, 0xc2, 0x20,
	0x98, 0x0a, 0x84, 0x62, 0x23, 0x09, 0x56, 0xec, 0x8a, 0x8d, 0x60, 0x8a, 0x8d, 0x9c, 0xb8, 0xa2,
	0xe0, 0xde, 0x4d, 0x62, 0x03, 0xfb, 0xdf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x76, 0x13, 0x68,
	0x62, 0x1a, 0x01, 0x00, 0x00,
}
