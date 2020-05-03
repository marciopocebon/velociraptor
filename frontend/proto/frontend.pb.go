// Code generated by protoc-gen-go. DO NOT EDIT.
// source: frontend.proto

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

// Describe the frontend state.
type FrontendState struct {
	// The unique name of the frontend.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Updated periodically to indicate the frontend is alive (in seconds).
	Heartbeat int64 `protobuf:"varint,2,opt,name=heartbeat,proto3" json:"heartbeat,omitempty"`
	// The URL over which the frontend is accessible to clients.
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrontendState) Reset()         { *m = FrontendState{} }
func (m *FrontendState) String() string { return proto.CompactTextString(m) }
func (*FrontendState) ProtoMessage()    {}
func (*FrontendState) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{0}
}

func (m *FrontendState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrontendState.Unmarshal(m, b)
}
func (m *FrontendState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrontendState.Marshal(b, m, deterministic)
}
func (m *FrontendState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrontendState.Merge(m, src)
}
func (m *FrontendState) XXX_Size() int {
	return xxx_messageInfo_FrontendState.Size(m)
}
func (m *FrontendState) XXX_DiscardUnknown() {
	xxx_messageInfo_FrontendState.DiscardUnknown(m)
}

var xxx_messageInfo_FrontendState proto.InternalMessageInfo

func (m *FrontendState) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FrontendState) GetHeartbeat() int64 {
	if m != nil {
		return m.Heartbeat
	}
	return 0
}

func (m *FrontendState) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*FrontendState)(nil), "proto.FrontendState")
}

func init() {
	proto.RegisterFile("frontend.proto", fileDescriptor_eca3873955a29cfe)
}

var fileDescriptor_eca3873955a29cfe = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x2b, 0xca, 0xcf,
	0x2b, 0x49, 0xcd, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xc1,
	0x5c, 0xbc, 0x6e, 0x50, 0x89, 0xe0, 0x92, 0xc4, 0x92, 0x54, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4,
	0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x86, 0x8b, 0x33, 0x23,
	0x35, 0xb1, 0xa8, 0x24, 0x29, 0x35, 0xb1, 0x44, 0x82, 0x49, 0x81, 0x51, 0x83, 0x39, 0x08, 0x21,
	0x20, 0x24, 0xc0, 0xc5, 0x5c, 0x5a, 0x94, 0x23, 0xc1, 0x0c, 0xd6, 0x00, 0x62, 0x3a, 0x99, 0x45,
	0x99, 0x94, 0x97, 0x97, 0xeb, 0x95, 0xa5, 0xe6, 0xe4, 0x27, 0x67, 0xa6, 0xa4, 0x56, 0xe8, 0x25,
	0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xe7, 0x24, 0xe6, 0xa5, 0xeb, 0x43, 0x04, 0x8b, 0x12, 0x0b, 0x4a,
	0xf2, 0x8b, 0xf4, 0x61, 0x6e, 0xd2, 0x07, 0x3b, 0x26, 0x89, 0x0d, 0x4c, 0x19, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x68, 0x68, 0x66, 0xa5, 0xac, 0x00, 0x00, 0x00,
}
