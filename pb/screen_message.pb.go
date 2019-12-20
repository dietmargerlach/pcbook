// Code generated by protoc-gen-go. DO NOT EDIT.
// source: screen_message.proto

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

type Screen_Panel int32

const (
	Screen_UNKNOWN Screen_Panel = 0
	Screen_IPS     Screen_Panel = 1
	Screen_OLED    Screen_Panel = 2
)

var Screen_Panel_name = map[int32]string{
	0: "UNKNOWN",
	1: "IPS",
	2: "OLED",
}

var Screen_Panel_value = map[string]int32{
	"UNKNOWN": 0,
	"IPS":     1,
	"OLED":    2,
}

func (x Screen_Panel) String() string {
	return proto.EnumName(Screen_Panel_name, int32(x))
}

func (Screen_Panel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0, 0}
}

type Screen struct {
	SizeInch             float32            `protobuf:"fixed32,1,opt,name=size_inch,json=sizeInch,proto3" json:"size_inch,omitempty"`
	Resolution           *Screen_Resolution `protobuf:"bytes,2,opt,name=resolution,proto3" json:"resolution,omitempty"`
	Panel                Screen_Panel       `protobuf:"varint,3,opt,name=panel,proto3,enum=pcbook.proto.Screen_Panel" json:"panel,omitempty"`
	Multitouch           bool               `protobuf:"varint,4,opt,name=multitouch,proto3" json:"multitouch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Screen) Reset()         { *m = Screen{} }
func (m *Screen) String() string { return proto.CompactTextString(m) }
func (*Screen) ProtoMessage()    {}
func (*Screen) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0}
}

func (m *Screen) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Screen.Unmarshal(m, b)
}
func (m *Screen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Screen.Marshal(b, m, deterministic)
}
func (m *Screen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Screen.Merge(m, src)
}
func (m *Screen) XXX_Size() int {
	return xxx_messageInfo_Screen.Size(m)
}
func (m *Screen) XXX_DiscardUnknown() {
	xxx_messageInfo_Screen.DiscardUnknown(m)
}

var xxx_messageInfo_Screen proto.InternalMessageInfo

func (m *Screen) GetSizeInch() float32 {
	if m != nil {
		return m.SizeInch
	}
	return 0
}

func (m *Screen) GetResolution() *Screen_Resolution {
	if m != nil {
		return m.Resolution
	}
	return nil
}

func (m *Screen) GetPanel() Screen_Panel {
	if m != nil {
		return m.Panel
	}
	return Screen_UNKNOWN
}

func (m *Screen) GetMultitouch() bool {
	if m != nil {
		return m.Multitouch
	}
	return false
}

type Screen_Resolution struct {
	Width                uint32   `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height               uint32   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Screen_Resolution) Reset()         { *m = Screen_Resolution{} }
func (m *Screen_Resolution) String() string { return proto.CompactTextString(m) }
func (*Screen_Resolution) ProtoMessage()    {}
func (*Screen_Resolution) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0, 0}
}

func (m *Screen_Resolution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Screen_Resolution.Unmarshal(m, b)
}
func (m *Screen_Resolution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Screen_Resolution.Marshal(b, m, deterministic)
}
func (m *Screen_Resolution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Screen_Resolution.Merge(m, src)
}
func (m *Screen_Resolution) XXX_Size() int {
	return xxx_messageInfo_Screen_Resolution.Size(m)
}
func (m *Screen_Resolution) XXX_DiscardUnknown() {
	xxx_messageInfo_Screen_Resolution.DiscardUnknown(m)
}

var xxx_messageInfo_Screen_Resolution proto.InternalMessageInfo

func (m *Screen_Resolution) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Screen_Resolution) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterEnum("pcbook.proto.Screen_Panel", Screen_Panel_name, Screen_Panel_value)
	proto.RegisterType((*Screen)(nil), "pcbook.proto.Screen")
	proto.RegisterType((*Screen_Resolution)(nil), "pcbook.proto.Screen.Resolution")
}

func init() { proto.RegisterFile("screen_message.proto", fileDescriptor_8a2814cd02b8aba7) }

var fileDescriptor_8a2814cd02b8aba7 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8e, 0x4f, 0x4b, 0xfb, 0x30,
	0x1c, 0xc6, 0x7f, 0xe9, 0xda, 0xae, 0xbf, 0xef, 0x9c, 0x94, 0x2f, 0x43, 0xca, 0x04, 0x2d, 0xbb,
	0xd8, 0x53, 0x91, 0x79, 0xf3, 0x22, 0x88, 0x1e, 0x86, 0xd2, 0x8d, 0x0c, 0x11, 0xbc, 0x8c, 0x35,
	0x86, 0x25, 0xd8, 0x25, 0xa5, 0x49, 0x11, 0x7c, 0x2f, 0xbe, 0x57, 0x59, 0x23, 0xda, 0x83, 0xc7,
	0xe7, 0xe1, 0xf3, 0xfc, 0x81, 0x89, 0x61, 0x0d, 0xe7, 0x6a, 0xb3, 0xe7, 0xc6, 0x6c, 0x77, 0x3c,
	0xaf, 0x1b, 0x6d, 0x35, 0x1e, 0xd5, 0xac, 0xd4, 0xfa, 0xcd, 0xa9, 0xd9, 0xa7, 0x07, 0xe1, 0xba,
	0xc3, 0xf0, 0x14, 0xfe, 0x1b, 0xf9, 0xc1, 0x37, 0x52, 0x31, 0x91, 0x90, 0x94, 0x64, 0x1e, 0x8d,
	0x0e, 0xc6, 0x42, 0x31, 0x81, 0x37, 0x00, 0x0d, 0x37, 0xba, 0x6a, 0xad, 0xd4, 0x2a, 0xf1, 0x52,
	0x92, 0x8d, 0xe6, 0xe7, 0x79, 0xbf, 0x2a, 0x77, 0x35, 0x39, 0xfd, 0xc1, 0x68, 0x2f, 0x82, 0x97,
	0x10, 0xd4, 0x5b, 0xc5, 0xab, 0x64, 0x90, 0x92, 0xec, 0x78, 0x3e, 0xfd, 0x33, 0xbb, 0x3a, 0x10,
	0xd4, 0x81, 0x78, 0x06, 0xb0, 0x6f, 0x2b, 0x2b, 0xad, 0x6e, 0x99, 0x48, 0xfc, 0x94, 0x64, 0x11,
	0xed, 0x39, 0xd3, 0x6b, 0x80, 0xdf, 0x2d, 0x9c, 0x40, 0xf0, 0x2e, 0x5f, 0xad, 0x7b, 0x3e, 0xa6,
	0x4e, 0xe0, 0x09, 0x84, 0x82, 0xcb, 0x9d, 0xb0, 0xdd, 0xe5, 0x31, 0xfd, 0x56, 0xb3, 0x0b, 0x08,
	0xba, 0x2d, 0x1c, 0xc1, 0xf0, 0xa9, 0x78, 0x28, 0x96, 0xcf, 0x45, 0xfc, 0x0f, 0x87, 0x30, 0x58,
	0xac, 0xd6, 0x31, 0xc1, 0x08, 0xfc, 0xe5, 0xe3, 0xfd, 0x5d, 0xec, 0xdd, 0xfa, 0x2f, 0x5e, 0x5d,
	0x96, 0x61, 0xf7, 0xf2, 0xea, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x93, 0xb0, 0xd1, 0xf4, 0x52, 0x01,
	0x00, 0x00,
}
