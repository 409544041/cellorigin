// Code generated by protoc-gen-go.
// source: channel.proto
// DO NOT EDIT!

/*
Package gamedef is a generated protocol buffer package.

It is generated from these files:
	channel.proto
	service.proto
	tool.proto

It has these top-level messages:
	ChannelDefine
	ChannelFile
	ServiceDefine
	ServiceFile
	LocalConfig
	TableCodeOption
*/
package gamedef

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// 通信线路定义
type ChannelDefine struct {
	Name    string `protobuf:"bytes,1,opt,name=Name,json=name" json:"Name,omitempty"`
	Port    int32  `protobuf:"varint,2,opt,name=Port,json=port" json:"Port,omitempty"`
	IPIndex int32  `protobuf:"varint,3,opt,name=IPIndex,json=iPIndex" json:"IPIndex,omitempty"`
}

func (m *ChannelDefine) Reset()                    { *m = ChannelDefine{} }
func (m *ChannelDefine) String() string            { return proto.CompactTextString(m) }
func (*ChannelDefine) ProtoMessage()               {}
func (*ChannelDefine) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 通信线路文件
type ChannelFile struct {
	Channel []*ChannelDefine `protobuf:"bytes,1,rep,name=Channel,json=channel" json:"Channel,omitempty"`
}

func (m *ChannelFile) Reset()                    { *m = ChannelFile{} }
func (m *ChannelFile) String() string            { return proto.CompactTextString(m) }
func (*ChannelFile) ProtoMessage()               {}
func (*ChannelFile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ChannelFile) GetChannel() []*ChannelDefine {
	if m != nil {
		return m.Channel
	}
	return nil
}

func init() {
	proto.RegisterType((*ChannelDefine)(nil), "gamedef.ChannelDefine")
	proto.RegisterType((*ChannelFile)(nil), "gamedef.ChannelFile")
}

var fileDescriptor0 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0x48, 0xcc,
	0xcb, 0x4b, 0xcd, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x4f, 0xcc, 0x4d, 0x4d,
	0x49, 0x4d, 0x53, 0x0a, 0xe4, 0xe2, 0x75, 0x86, 0xc8, 0xb8, 0xa4, 0xa6, 0x65, 0xe6, 0xa5, 0x0a,
	0x09, 0x71, 0xb1, 0xf8, 0x01, 0xe5, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x58, 0xf2, 0x80,
	0x6c, 0x90, 0x58, 0x40, 0x7e, 0x51, 0x89, 0x04, 0x13, 0x50, 0x8c, 0x35, 0x88, 0xa5, 0x00, 0xc8,
	0x16, 0x92, 0xe0, 0x62, 0xf7, 0x0c, 0xf0, 0xcc, 0x4b, 0x49, 0xad, 0x90, 0x60, 0x06, 0x0b, 0xb3,
	0x67, 0x42, 0xb8, 0x4a, 0xf6, 0x5c, 0xdc, 0x50, 0x23, 0xdd, 0x32, 0x73, 0x52, 0x85, 0x0c, 0xb8,
	0xd8, 0xa1, 0x5c, 0xa0, 0x99, 0xcc, 0x1a, 0xdc, 0x46, 0x62, 0x7a, 0x50, 0xcb, 0xf5, 0x50, 0x6c,
	0x0e, 0x62, 0x87, 0x3a, 0x31, 0x89, 0x0d, 0xec, 0x46, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x5b, 0x41, 0xe2, 0x68, 0xb4, 0x00, 0x00, 0x00,
}
