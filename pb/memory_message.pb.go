// Code generated by protoc-gen-go. DO NOT EDIT.
// source: memory_message.proto

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

type Memory_Unit int32

const (
	Memory_UNKNOWN  Memory_Unit = 0
	Memory_BIT      Memory_Unit = 1
	Memory_BYTE     Memory_Unit = 2
	Memory_KILOBYTE Memory_Unit = 3
	Memory_MEGABYTE Memory_Unit = 4
	Memory_GIGABYTE Memory_Unit = 5
	Memory_TERABYTE Memory_Unit = 6
)

var Memory_Unit_name = map[int32]string{
	0: "UNKNOWN",
	1: "BIT",
	2: "BYTE",
	3: "KILOBYTE",
	4: "MEGABYTE",
	5: "GIGABYTE",
	6: "TERABYTE",
}

var Memory_Unit_value = map[string]int32{
	"UNKNOWN":  0,
	"BIT":      1,
	"BYTE":     2,
	"KILOBYTE": 3,
	"MEGABYTE": 4,
	"GIGABYTE": 5,
	"TERABYTE": 6,
}

func (x Memory_Unit) String() string {
	return proto.EnumName(Memory_Unit_name, int32(x))
}

func (Memory_Unit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c0c7f919ccc765da, []int{0, 0}
}

type Memory struct {
	Value                uint64      `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Unit                 Memory_Unit `protobuf:"varint,2,opt,name=unit,proto3,enum=pcbook.proto.Memory_Unit" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Memory) Reset()         { *m = Memory{} }
func (m *Memory) String() string { return proto.CompactTextString(m) }
func (*Memory) ProtoMessage()    {}
func (*Memory) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0c7f919ccc765da, []int{0}
}

func (m *Memory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Memory.Unmarshal(m, b)
}
func (m *Memory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Memory.Marshal(b, m, deterministic)
}
func (m *Memory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Memory.Merge(m, src)
}
func (m *Memory) XXX_Size() int {
	return xxx_messageInfo_Memory.Size(m)
}
func (m *Memory) XXX_DiscardUnknown() {
	xxx_messageInfo_Memory.DiscardUnknown(m)
}

var xxx_messageInfo_Memory proto.InternalMessageInfo

func (m *Memory) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Memory) GetUnit() Memory_Unit {
	if m != nil {
		return m.Unit
	}
	return Memory_UNKNOWN
}

func init() {
	proto.RegisterEnum("pcbook.proto.Memory_Unit", Memory_Unit_name, Memory_Unit_value)
	proto.RegisterType((*Memory)(nil), "pcbook.proto.Memory")
}

func init() { proto.RegisterFile("memory_message.proto", fileDescriptor_c0c7f919ccc765da) }

var fileDescriptor_c0c7f919ccc765da = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0x4d, 0xcd, 0xcd,
	0x2f, 0xaa, 0x8c, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x29, 0x48, 0x4e, 0xca, 0xcf, 0xcf, 0x86, 0xf0, 0x94, 0xd6, 0x32, 0x72, 0xb1, 0xf9,
	0x82, 0x95, 0x09, 0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0xb0, 0x04, 0x41, 0x38, 0x42, 0xba, 0x5c, 0x2c, 0xa5, 0x79, 0x99, 0x25, 0x12, 0x4c, 0x0a, 0x8c,
	0x1a, 0x7c, 0x46, 0x92, 0x7a, 0xc8, 0xba, 0xf5, 0x20, 0x3a, 0xf5, 0x42, 0xf3, 0x32, 0x4b, 0x82,
	0xc0, 0xca, 0x94, 0xe2, 0xb8, 0x58, 0x40, 0x3c, 0x21, 0x6e, 0x2e, 0xf6, 0x50, 0x3f, 0x6f, 0x3f,
	0xff, 0x70, 0x3f, 0x01, 0x06, 0x21, 0x76, 0x2e, 0x66, 0x27, 0xcf, 0x10, 0x01, 0x46, 0x21, 0x0e,
	0x2e, 0x16, 0xa7, 0xc8, 0x10, 0x57, 0x01, 0x26, 0x21, 0x1e, 0x2e, 0x0e, 0x6f, 0x4f, 0x1f, 0x7f,
	0x30, 0x8f, 0x19, 0xc4, 0xf3, 0x75, 0x75, 0x77, 0x04, 0xf3, 0x58, 0x40, 0x3c, 0x77, 0x4f, 0x28,
	0x8f, 0x15, 0xc4, 0x0b, 0x71, 0x0d, 0x82, 0xf0, 0xd8, 0x9c, 0x58, 0xa2, 0x98, 0x0a, 0x92, 0x92,
	0xd8, 0xc0, 0xd6, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x98, 0xd3, 0xa8, 0xf4, 0xe2, 0x00,
	0x00, 0x00,
}
