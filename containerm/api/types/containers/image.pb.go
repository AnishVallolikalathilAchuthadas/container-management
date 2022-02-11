// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/types/containers/image.proto

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

// Image contains the reference of the image used to build the
// specification and snapshots for running this container.
//
// If this field is updated, the spec and rootfs needed to updated, as well.
type Image struct {
	//Image name or id
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4e55ada209af028, []int{0}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Image)(nil), "github.com.eclipse_kanto.container_management.containerm.api.types.containers.Image")
}

func init() {
	proto.RegisterFile("api/types/containers/image.proto", fileDescriptor_d4e55ada209af028)
}

var fileDescriptor_d4e55ada209af028 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0x2c, 0xc8, 0xd4,
	0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xcc, 0xcc, 0x4b, 0x2d, 0x2a,
	0xd6, 0xcf, 0xcc, 0x4d, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xf2, 0x4d, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4b, 0x4d, 0xce, 0xc9, 0x2c, 0x28, 0x4e,
	0x8d, 0xcf, 0x4e, 0xcc, 0x2b, 0xc9, 0xd7, 0x83, 0x6b, 0x88, 0xcf, 0x4d, 0xcc, 0x4b, 0x4c, 0x4f,
	0xcd, 0x4d, 0xcd, 0x2b, 0x41, 0x08, 0xe6, 0xea, 0x25, 0x16, 0x64, 0xea, 0x81, 0x8d, 0x46, 0x08,
	0x16, 0x2b, 0x49, 0x73, 0xb1, 0x7a, 0x82, 0x4c, 0x17, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d,
	0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x9d, 0xa2, 0xa2, 0x22, 0x10, 0xb6, 0xe9,
	0x43, 0x6d, 0xd3, 0x05, 0xdb, 0x86, 0x70, 0x9e, 0x2e, 0xc2, 0x36, 0x84, 0x60, 0xae, 0x3e, 0x36,
	0x8f, 0x58, 0x23, 0x98, 0x49, 0x6c, 0x60, 0xef, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x15,
	0xc4, 0x78, 0x46, 0xf2, 0x00, 0x00, 0x00,
}
