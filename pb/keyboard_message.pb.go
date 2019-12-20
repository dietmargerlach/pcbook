// Code generated by protoc-gen-go. DO NOT EDIT.
// source: keyboard_message.proto

package pb

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

type Keyboard_Layout int32

const (
	Keyboard_UNKNOWN Keyboard_Layout = 0
	Keyboard_QWERTY  Keyboard_Layout = 1
	Keyboard_QWERTZ  Keyboard_Layout = 2
	Keyboard_AWERTY  Keyboard_Layout = 3
)

var Keyboard_Layout_name = map[int32]string{
	0: "UNKNOWN",
	1: "QWERTY",
	2: "QWERTZ",
	3: "AWERTY",
}

var Keyboard_Layout_value = map[string]int32{
	"UNKNOWN": 0,
	"QWERTY":  1,
	"QWERTZ":  2,
	"AWERTY":  3,
}

func (x Keyboard_Layout) String() string {
	return proto.EnumName(Keyboard_Layout_name, int32(x))
}

func (Keyboard_Layout) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8fd4226b8635fe5d, []int{0, 0}
}

type Keyboard struct {
	Layout               Keyboard_Layout `protobuf:"varint,1,opt,name=layout,proto3,enum=pcbook.proto.Keyboard_Layout" json:"layout,omitempty"`
	Backlit              bool            `protobuf:"varint,2,opt,name=backlit,proto3" json:"backlit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Keyboard) Reset()         { *m = Keyboard{} }
func (m *Keyboard) String() string { return proto.CompactTextString(m) }
func (*Keyboard) ProtoMessage()    {}
func (*Keyboard) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fd4226b8635fe5d, []int{0}
}

func (m *Keyboard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Keyboard.Unmarshal(m, b)
}
func (m *Keyboard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Keyboard.Marshal(b, m, deterministic)
}
func (m *Keyboard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keyboard.Merge(m, src)
}
func (m *Keyboard) XXX_Size() int {
	return xxx_messageInfo_Keyboard.Size(m)
}
func (m *Keyboard) XXX_DiscardUnknown() {
	xxx_messageInfo_Keyboard.DiscardUnknown(m)
}

var xxx_messageInfo_Keyboard proto.InternalMessageInfo

func (m *Keyboard) GetLayout() Keyboard_Layout {
	if m != nil {
		return m.Layout
	}
	return Keyboard_UNKNOWN
}

func (m *Keyboard) GetBacklit() bool {
	if m != nil {
		return m.Backlit
	}
	return false
}

func init() {
	proto.RegisterEnum("pcbook.proto.Keyboard_Layout", Keyboard_Layout_name, Keyboard_Layout_value)
	proto.RegisterType((*Keyboard)(nil), "pcbook.proto.Keyboard")
}

func init() { proto.RegisterFile("keyboard_message.proto", fileDescriptor_8fd4226b8635fe5d) }

var fileDescriptor_8fd4226b8635fe5d = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0x4e, 0xad, 0x4c,
	0xca, 0x4f, 0x2c, 0x4a, 0x89, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x29, 0x48, 0x4e, 0xca, 0xcf, 0xcf, 0x86, 0xf0, 0x94, 0xa6, 0x31, 0x72,
	0x71, 0x78, 0x43, 0x15, 0x0a, 0x99, 0x72, 0xb1, 0xe5, 0x24, 0x56, 0xe6, 0x97, 0x96, 0x48, 0x30,
	0x2a, 0x30, 0x6a, 0xf0, 0x19, 0xc9, 0xea, 0x21, 0xab, 0xd5, 0x83, 0xa9, 0xd3, 0xf3, 0x01, 0x2b,
	0x0a, 0x82, 0x2a, 0x16, 0x92, 0xe0, 0x62, 0x4f, 0x4a, 0x4c, 0xce, 0xce, 0xc9, 0x2c, 0x91, 0x60,
	0x52, 0x60, 0xd4, 0xe0, 0x08, 0x82, 0x71, 0x95, 0x2c, 0xb9, 0xd8, 0x20, 0x6a, 0x85, 0xb8, 0xb9,
	0xd8, 0x43, 0xfd, 0xbc, 0xfd, 0xfc, 0xc3, 0xfd, 0x04, 0x18, 0x84, 0xb8, 0xb8, 0xd8, 0x02, 0xc3,
	0x5d, 0x83, 0x42, 0x22, 0x05, 0x18, 0xe1, 0xec, 0x28, 0x01, 0x26, 0x10, 0xdb, 0x11, 0x22, 0xce,
	0xec, 0xc4, 0x12, 0xc5, 0x54, 0x90, 0x94, 0xc4, 0x06, 0xb6, 0xd9, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x66, 0x93, 0x21, 0x33, 0xcd, 0x00, 0x00, 0x00,
}
