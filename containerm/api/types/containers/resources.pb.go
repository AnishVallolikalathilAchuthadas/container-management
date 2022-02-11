// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/types/containers/resources.proto

package containers

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

// Represents resource limitation of a container
type Resources struct {
	// Hard memory limit
	Memory string `protobuf:"bytes,1,opt,name=memory,proto3" json:"memory,omitempty"`
	// Soft memory limit
	MemoryReservation string `protobuf:"bytes,2,opt,name=memory_reservation,json=memoryReservation,proto3" json:"memory_reservation,omitempty"`
	// Swap memory limit(memory + swap)
	MemorySwap           string   `protobuf:"bytes,3,opt,name=memory_swap,json=memorySwap,proto3" json:"memory_swap,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resources) Reset()         { *m = Resources{} }
func (m *Resources) String() string { return proto.CompactTextString(m) }
func (*Resources) ProtoMessage()    {}
func (*Resources) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7c067718a7b8fce, []int{0}
}

func (m *Resources) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resources.Unmarshal(m, b)
}
func (m *Resources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resources.Marshal(b, m, deterministic)
}
func (m *Resources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resources.Merge(m, src)
}
func (m *Resources) XXX_Size() int {
	return xxx_messageInfo_Resources.Size(m)
}
func (m *Resources) XXX_DiscardUnknown() {
	xxx_messageInfo_Resources.DiscardUnknown(m)
}

var xxx_messageInfo_Resources proto.InternalMessageInfo

func (m *Resources) GetMemory() string {
	if m != nil {
		return m.Memory
	}
	return ""
}

func (m *Resources) GetMemoryReservation() string {
	if m != nil {
		return m.MemoryReservation
	}
	return ""
}

func (m *Resources) GetMemorySwap() string {
	if m != nil {
		return m.MemorySwap
	}
	return ""
}

func init() {
	proto.RegisterType((*Resources)(nil), "github.com.eclipse_kanto.container_management.containerm.api.types.containers.Resources")
}

func init() {
	proto.RegisterFile("api/types/containers/resources.proto", fileDescriptor_a7c067718a7b8fce)
}

var fileDescriptor_a7c067718a7b8fce = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4b, 0xc6, 0x30,
	0x10, 0x86, 0xf9, 0x14, 0x3e, 0x68, 0x9c, 0xcc, 0x20, 0xdd, 0x14, 0x71, 0x70, 0x69, 0x32, 0x38,
	0xba, 0xb9, 0xbb, 0xd4, 0x45, 0xba, 0x94, 0x6b, 0x38, 0x6a, 0xd0, 0xe4, 0x42, 0x2e, 0xb5, 0xf4,
	0xdf, 0x0b, 0x69, 0xf5, 0x16, 0xb7, 0xcb, 0x73, 0x6f, 0x12, 0x9e, 0x57, 0x3d, 0x40, 0xf2, 0xb6,
	0x6c, 0x09, 0xd9, 0x3a, 0x8a, 0x05, 0x7c, 0xc4, 0xcc, 0x36, 0x23, 0xd3, 0x92, 0x1d, 0xb2, 0x49,
	0x99, 0x0a, 0xe9, 0xd7, 0xd9, 0x97, 0x8f, 0x65, 0x32, 0x8e, 0x82, 0x41, 0xf7, 0xe5, 0x13, 0xe3,
	0xf8, 0x09, 0xb1, 0x90, 0xf9, 0xbb, 0x34, 0x06, 0x88, 0x30, 0x63, 0xc0, 0x58, 0x04, 0x06, 0x03,
	0xc9, 0x9b, 0xfa, 0xbc, 0x40, 0xbe, 0x67, 0xd5, 0xf4, 0xbf, 0x3f, 0xe8, 0x1b, 0x75, 0x0e, 0x18,
	0x28, 0x6f, 0xed, 0xe9, 0xee, 0xf4, 0xd8, 0xf4, 0xc7, 0x49, 0x77, 0x4a, 0xef, 0xd3, 0x98, 0x91,
	0x31, 0x7f, 0x43, 0xf1, 0x14, 0xdb, 0x8b, 0x9a, 0xb9, 0xde, 0x37, 0xbd, 0x2c, 0xf4, 0xad, 0xba,
	0x3a, 0xe2, 0xbc, 0x42, 0x6a, 0x2f, 0x6b, 0x4e, 0xed, 0xe8, 0x6d, 0x85, 0xf4, 0x32, 0x0c, 0xef,
	0x62, 0x61, 0x0f, 0x8b, 0xae, 0x5a, 0x88, 0x7a, 0x27, 0x16, 0x02, 0x83, 0xfd, 0xaf, 0xa4, 0x67,
	0x19, 0xa7, 0x73, 0xad, 0xe9, 0xe9, 0x27, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x58, 0xfe, 0x1c, 0x4e,
	0x01, 0x00, 0x00,
}
