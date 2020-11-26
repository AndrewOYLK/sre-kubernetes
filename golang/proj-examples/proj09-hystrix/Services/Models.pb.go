// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Models.proto

package Services

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

// 商品模型
type ProdModel struct {
	// @inject_tag: json:"pid"
	ProdID int32 `protobuf:"varint,1,opt,name=ProdID,proto3" json:"pid"`
	// @inject_tag: json:"pname"
	ProdName             string   `protobuf:"bytes,2,opt,name=ProdName,proto3" json:"pname"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdModel) Reset()         { *m = ProdModel{} }
func (m *ProdModel) String() string { return proto.CompactTextString(m) }
func (*ProdModel) ProtoMessage()    {}
func (*ProdModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b05f67b8e9f86a, []int{0}
}

func (m *ProdModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdModel.Unmarshal(m, b)
}
func (m *ProdModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdModel.Marshal(b, m, deterministic)
}
func (m *ProdModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdModel.Merge(m, src)
}
func (m *ProdModel) XXX_Size() int {
	return xxx_messageInfo_ProdModel.Size(m)
}
func (m *ProdModel) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdModel.DiscardUnknown(m)
}

var xxx_messageInfo_ProdModel proto.InternalMessageInfo

func (m *ProdModel) GetProdID() int32 {
	if m != nil {
		return m.ProdID
	}
	return 0
}

func (m *ProdModel) GetProdName() string {
	if m != nil {
		return m.ProdName
	}
	return ""
}

func init() {
	proto.RegisterType((*ProdModel)(nil), "Services.ProdModel")
}

func init() {
	proto.RegisterFile("Models.proto", fileDescriptor_96b05f67b8e9f86a)
}

var fileDescriptor_96b05f67b8e9f86a = []byte{
	// 98 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xf1, 0xcd, 0x4f, 0x49,
	0xcd, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x08, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c,
	0x4e, 0x2d, 0x56, 0xb2, 0xe7, 0xe2, 0x0c, 0x28, 0xca, 0x4f, 0x01, 0xcb, 0x0a, 0x89, 0x71, 0xb1,
	0x81, 0x38, 0x9e, 0x2e, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x50, 0x9e, 0x90, 0x14, 0x17,
	0x07, 0x88, 0xe5, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe7, 0x27,
	0xb1, 0x81, 0x4d, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x72, 0xe5, 0xe8, 0x61, 0x00,
	0x00, 0x00,
}