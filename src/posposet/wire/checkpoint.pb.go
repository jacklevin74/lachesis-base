// Code generated by protoc-gen-go. DO NOT EDIT.
// source: checkpoint.proto

package wire

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Checkpoint struct {
	SuperFrameN          uint64   `protobuf:"varint,1,opt,name=SuperFrameN,proto3" json:"SuperFrameN,omitempty"`
	LastBlockN           uint64   `protobuf:"varint,2,opt,name=LastBlockN,proto3" json:"LastBlockN,omitempty"`
	TotalCap             uint64   `protobuf:"varint,3,opt,name=TotalCap,proto3" json:"TotalCap,omitempty"`
	LastConsensusTime    uint64   `protobuf:"varint,4,opt,name=LastConsensusTime,proto3" json:"LastConsensusTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Checkpoint) Reset()         { *m = Checkpoint{} }
func (m *Checkpoint) String() string { return proto.CompactTextString(m) }
func (*Checkpoint) ProtoMessage()    {}
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bab050ffa824783, []int{0}
}

func (m *Checkpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Checkpoint.Unmarshal(m, b)
}
func (m *Checkpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Checkpoint.Marshal(b, m, deterministic)
}
func (m *Checkpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Checkpoint.Merge(m, src)
}
func (m *Checkpoint) XXX_Size() int {
	return xxx_messageInfo_Checkpoint.Size(m)
}
func (m *Checkpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Checkpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Checkpoint proto.InternalMessageInfo

func (m *Checkpoint) GetSuperFrameN() uint64 {
	if m != nil {
		return m.SuperFrameN
	}
	return 0
}

func (m *Checkpoint) GetLastBlockN() uint64 {
	if m != nil {
		return m.LastBlockN
	}
	return 0
}

func (m *Checkpoint) GetTotalCap() uint64 {
	if m != nil {
		return m.TotalCap
	}
	return 0
}

func (m *Checkpoint) GetLastConsensusTime() uint64 {
	if m != nil {
		return m.LastConsensusTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Checkpoint)(nil), "wire.Checkpoint")
}

func init() { proto.RegisterFile("checkpoint.proto", fileDescriptor_9bab050ffa824783) }

var fileDescriptor_9bab050ffa824783 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xce, 0x48, 0x4d,
	0xce, 0x2e, 0xc8, 0xcf, 0xcc, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0xcf,
	0x2c, 0x4a, 0x55, 0x9a, 0xc1, 0xc8, 0xc5, 0xe5, 0x0c, 0x97, 0x12, 0x52, 0xe0, 0xe2, 0x0e, 0x2e,
	0x2d, 0x48, 0x2d, 0x72, 0x2b, 0x4a, 0xcc, 0x4d, 0xf5, 0x93, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x09,
	0x42, 0x16, 0x12, 0x92, 0xe3, 0xe2, 0xf2, 0x49, 0x2c, 0x2e, 0x71, 0xca, 0xc9, 0x4f, 0xce, 0xf6,
	0x93, 0x60, 0x02, 0x2b, 0x40, 0x12, 0x11, 0x92, 0xe2, 0xe2, 0x08, 0xc9, 0x2f, 0x49, 0xcc, 0x71,
	0x4e, 0x2c, 0x90, 0x60, 0x06, 0xcb, 0xc2, 0xf9, 0x42, 0x3a, 0x5c, 0x82, 0x20, 0x95, 0xce, 0xf9,
	0x79, 0xc5, 0xa9, 0x79, 0xc5, 0xa5, 0xc5, 0x21, 0x99, 0xb9, 0xa9, 0x12, 0x2c, 0x60, 0x45, 0x98,
	0x12, 0x49, 0x6c, 0x60, 0x77, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x06, 0x9c, 0x35,
	0xbb, 0x00, 0x00, 0x00,
}
