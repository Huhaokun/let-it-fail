// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

package contract

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

type Endpoint struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NetworkId            string            `protobuf:"bytes,2,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	Host                 string            `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Ports                []int32           `protobuf:"varint,4,rep,packed,name=ports,proto3" json:"ports,omitempty"`
	Labels               map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0}
}

func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Endpoint) GetNetworkId() string {
	if m != nil {
		return m.NetworkId
	}
	return ""
}

func (m *Endpoint) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Endpoint) GetPorts() []int32 {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *Endpoint) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "letitfail.Endpoint")
	proto.RegisterMapType((map[string]string)(nil), "letitfail.Endpoint.LabelsEntry")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_4c16552f9fdb66d8) }

var fileDescriptor_4c16552f9fdb66d8 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xcd, 0x4a, 0xc0, 0x30,
	0x10, 0x84, 0x49, 0xd2, 0x16, 0xbb, 0x05, 0x91, 0xc5, 0x43, 0x10, 0xc4, 0xe2, 0xa9, 0xa7, 0x1c,
	0xf4, 0xe0, 0xcf, 0xbd, 0x07, 0xc1, 0x53, 0x5f, 0x40, 0x5a, 0x12, 0x31, 0x34, 0x26, 0x25, 0x5d,
	0x95, 0x3e, 0xa9, 0xaf, 0x23, 0x4d, 0xa3, 0x78, 0x9b, 0x99, 0x1d, 0x98, 0x6f, 0xa1, 0x79, 0x0f,
	0xda, 0x38, 0xb5, 0xc4, 0x40, 0x01, 0x6b, 0x67, 0xc8, 0xd2, 0xeb, 0x68, 0xdd, 0xf5, 0x37, 0x83,
	0x93, 0xde, 0xeb, 0x25, 0x58, 0x4f, 0x78, 0x0a, 0xdc, 0x6a, 0xc9, 0x5a, 0xd6, 0xd5, 0x03, 0xb7,
	0x1a, 0x2f, 0x01, 0xbc, 0xa1, 0xaf, 0x10, 0xe7, 0x17, 0xab, 0x25, 0x4f, 0x79, 0x9d, 0x93, 0x27,
	0x8d, 0x08, 0xc5, 0x5b, 0x58, 0x49, 0x8a, 0x74, 0x48, 0x1a, 0xcf, 0xa1, 0x5c, 0x42, 0xa4, 0x55,
	0x16, 0xad, 0xe8, 0xca, 0xe1, 0x30, 0x78, 0x07, 0x95, 0x1b, 0x27, 0xe3, 0x56, 0x59, 0xb6, 0xa2,
	0x6b, 0x6e, 0xae, 0xd4, 0x1f, 0x81, 0xfa, 0x5d, 0x57, 0xcf, 0xa9, 0xd1, 0x7b, 0x8a, 0xdb, 0x90,
	0xeb, 0x17, 0x0f, 0xd0, 0xfc, 0x8b, 0xf1, 0x0c, 0xc4, 0x6c, 0xb6, 0x4c, 0xb8, 0xcb, 0x7d, 0xef,
	0x73, 0x74, 0x1f, 0x26, 0xd3, 0x1d, 0xe6, 0x91, 0xdf, 0xb3, 0xa9, 0x4a, 0xbf, 0xde, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x1d, 0x22, 0x63, 0xe2, 0xfa, 0x00, 0x00, 0x00,
}
