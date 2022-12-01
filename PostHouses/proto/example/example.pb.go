// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_PostHouses

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

type Request struct {
	Sessionid            string   `protobuf:"bytes,1,opt,name=sessionid,proto3" json:"sessionid,omitempty"`
	Max                  []byte   `protobuf:"bytes,2,opt,name=max,proto3" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetSessionid() string {
	if m != nil {
		return m.Sessionid
	}
	return ""
}

func (m *Request) GetMax() []byte {
	if m != nil {
		return m.Max
	}
	return nil
}

type Response struct {
	Errno                string   `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	HouseId              int64    `protobuf:"varint,3,opt,name=house_id,json=houseId,proto3" json:"house_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Response) GetHouseId() int64 {
	if m != nil {
		return m.HouseId
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.PostHouses.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.PostHouses.Response")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4b, 0xc4, 0x40,
	0x0c, 0xc5, 0x1d, 0x8b, 0xdb, 0x6d, 0xf0, 0x20, 0x41, 0xb4, 0xfe, 0x39, 0x8c, 0x3d, 0xf5, 0x34,
	0x82, 0x9e, 0xfc, 0x00, 0x82, 0xde, 0x64, 0xf6, 0xae, 0x54, 0x1b, 0xea, 0x80, 0xd3, 0xd4, 0x49,
	0x2b, 0xfd, 0xf8, 0xd2, 0x69, 0xa5, 0x27, 0xf7, 0x94, 0xbc, 0x47, 0x1e, 0xf9, 0x25, 0x70, 0xd5,
	0x05, 0xee, 0xf9, 0x96, 0xc6, 0xca, 0x77, 0x5f, 0xf4, 0x57, 0x4d, 0x74, 0xf1, 0xbc, 0x61, 0xe3,
	0xdd, 0x47, 0x60, 0x23, 0xe1, 0xc7, 0xbc, 0xb0, 0xf4, 0x4f, 0x3c, 0x08, 0x49, 0xf1, 0x00, 0xa9,
	0xa5, 0xef, 0x81, 0xa4, 0xc7, 0x6b, 0xc8, 0x84, 0x44, 0x1c, 0xb7, 0xae, 0xce, 0x95, 0x56, 0x65,
	0x66, 0x57, 0x03, 0x4f, 0x20, 0xf1, 0xd5, 0x98, 0x1f, 0x6a, 0x55, 0x1e, 0xdb, 0xa9, 0x2d, 0x76,
	0xb0, 0xb5, 0x24, 0x1d, 0xb7, 0x42, 0x78, 0x0a, 0x47, 0x14, 0x42, 0xcb, 0x4b, 0x6e, 0x16, 0x78,
	0x06, 0x1b, 0x0a, 0xc1, 0x4b, 0x13, 0x63, 0x99, 0x5d, 0x14, 0x5e, 0xc0, 0xf6, 0x73, 0x5a, 0xff,
	0xe6, 0xea, 0x3c, 0xd1, 0xaa, 0x4c, 0x6c, 0x1a, 0xf5, 0x73, 0x7d, 0xf7, 0x0a, 0xe9, 0xe3, 0x4c,
	0x8e, 0x3b, 0x80, 0x15, 0x14, 0xb5, 0xf9, 0xe7, 0x04, 0xb3, 0xf0, 0x5f, 0xde, 0xec, 0x99, 0x98,
	0x31, 0x8b, 0x83, 0xf7, 0x4d, 0xfc, 0xc7, 0xfd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0x97,
	0x52, 0xec, 0x2e, 0x01, 0x00, 0x00,
}
