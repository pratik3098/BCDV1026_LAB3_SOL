// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto3/proto3.proto

package proto3

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Request_Flavour int32

const (
	Request_SWEET         Request_Flavour = 0
	Request_SOUR          Request_Flavour = 1
	Request_UMAMI         Request_Flavour = 2
	Request_GOPHERLICIOUS Request_Flavour = 3
)

var Request_Flavour_name = map[int32]string{
	0: "SWEET",
	1: "SOUR",
	2: "UMAMI",
	3: "GOPHERLICIOUS",
}

var Request_Flavour_value = map[string]int32{
	"SWEET":         0,
	"SOUR":          1,
	"UMAMI":         2,
	"GOPHERLICIOUS": 3,
}

func (x Request_Flavour) String() string {
	return proto.EnumName(Request_Flavour_name, int32(x))
}

func (Request_Flavour) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ab04eb4084a521db, []int{0, 0}
}

type Request struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Key                  []int64         `protobuf:"varint,2,rep,packed,name=key,proto3" json:"key,omitempty"`
	Taste                Request_Flavour `protobuf:"varint,3,opt,name=taste,proto3,enum=proto3.Request_Flavour" json:"taste,omitempty"`
	Book                 *Book           `protobuf:"bytes,4,opt,name=book,proto3" json:"book,omitempty"`
	Unpacked             []int64         `protobuf:"varint,5,rep,name=unpacked,proto3" json:"unpacked,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab04eb4084a521db, []int{0}
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

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetKey() []int64 {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Request) GetTaste() Request_Flavour {
	if m != nil {
		return m.Taste
	}
	return Request_SWEET
}

func (m *Request) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

func (m *Request) GetUnpacked() []int64 {
	if m != nil {
		return m.Unpacked
	}
	return nil
}

type Book struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	RawData              []byte   `protobuf:"bytes,2,opt,name=raw_data,json=rawData,proto3" json:"raw_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab04eb4084a521db, []int{1}
}
func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetRawData() []byte {
	if m != nil {
		return m.RawData
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto3.Request_Flavour", Request_Flavour_name, Request_Flavour_value)
	proto.RegisterType((*Request)(nil), "proto3.Request")
	proto.RegisterType((*Book)(nil), "proto3.Book")
}

func init() { proto.RegisterFile("proto3/proto3.proto", fileDescriptor_ab04eb4084a521db) }

var fileDescriptor_ab04eb4084a521db = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0x4f, 0x4f, 0xc2, 0x40,
	0x10, 0xc5, 0xd9, 0xfe, 0x11, 0x18, 0xd1, 0xd4, 0xd1, 0xc4, 0x7a, 0x31, 0x0d, 0xa7, 0x5e, 0x68,
	0x13, 0x3c, 0x78, 0x50, 0x0f, 0xa2, 0xa8, 0x24, 0x12, 0xcc, 0x22, 0x31, 0xf1, 0x62, 0x96, 0xb2,
	0xa9, 0xa4, 0xd0, 0xc1, 0x76, 0x2b, 0xf1, 0xcb, 0xfa, 0x59, 0xcc, 0xb6, 0xc5, 0xd3, 0xcc, 0xbc,
	0x99, 0xfc, 0x26, 0xef, 0xc1, 0xf1, 0x26, 0x23, 0x45, 0x17, 0x61, 0x55, 0x82, 0xb2, 0xe0, 0x5e,
	0x35, 0x75, 0x7f, 0x19, 0x34, 0xb9, 0xfc, 0x2a, 0x64, 0xae, 0x10, 0xc1, 0x4a, 0xc5, 0x5a, 0xba,
	0xcc, 0x63, 0x7e, 0x9b, 0x97, 0x3d, 0x3a, 0x60, 0x26, 0xf2, 0xc7, 0x35, 0x3c, 0xd3, 0x37, 0xb9,
	0x6e, 0xb1, 0x07, 0xb6, 0x12, 0xb9, 0x92, 0xae, 0xe9, 0x31, 0xff, 0xb0, 0x7f, 0x1a, 0xd4, 0xdc,
	0x9a, 0x12, 0x3c, 0xac, 0xc4, 0x37, 0x15, 0x19, 0xaf, 0xae, 0xd0, 0x03, 0x6b, 0x4e, 0x94, 0xb8,
	0x96, 0xc7, 0xfc, 0xfd, 0x7e, 0x67, 0x77, 0x3d, 0x20, 0x4a, 0x78, 0xb9, 0xc1, 0x73, 0x68, 0x15,
	0xe9, 0x46, 0x44, 0x89, 0x5c, 0xb8, 0xb6, 0xfe, 0x33, 0x30, 0x9c, 0x06, 0xff, 0xd7, 0xba, 0xd7,
	0xd0, 0xac, 0x99, 0xd8, 0x06, 0x7b, 0xfa, 0x36, 0x1c, 0xbe, 0x3a, 0x0d, 0x6c, 0x81, 0x35, 0x9d,
	0xcc, 0xb8, 0xc3, 0xb4, 0x38, 0x1b, 0xdf, 0x8e, 0x47, 0x8e, 0x81, 0x47, 0x70, 0xf0, 0x38, 0x79,
	0x79, 0x1a, 0xf2, 0xe7, 0xd1, 0xdd, 0x68, 0x32, 0x9b, 0x3a, 0x66, 0xf7, 0x12, 0x2c, 0xfd, 0x0b,
	0x4f, 0xc0, 0x56, 0x4b, 0xb5, 0xda, 0xb9, 0xab, 0x06, 0x3c, 0x83, 0x56, 0x26, 0xb6, 0x1f, 0x0b,
	0xa1, 0x84, 0x6b, 0x78, 0xcc, 0xef, 0xf0, 0x66, 0x26, 0xb6, 0xf7, 0x42, 0x89, 0xc1, 0xcd, 0xfb,
	0x55, 0xbc, 0x54, 0x9f, 0xc5, 0x3c, 0x88, 0x68, 0x1d, 0x46, 0x94, 0xaf, 0x29, 0x0f, 0x63, 0x8a,
	0xa9, 0x74, 0x50, 0xa5, 0x19, 0xf5, 0x62, 0x99, 0xf6, 0xb4, 0x18, 0x2a, 0x99, 0x2b, 0x8d, 0xa9,
	0x63, 0x9e, 0xd7, 0x01, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x31, 0x30, 0xc4, 0x7e, 0x01,
	0x00, 0x00,
}
