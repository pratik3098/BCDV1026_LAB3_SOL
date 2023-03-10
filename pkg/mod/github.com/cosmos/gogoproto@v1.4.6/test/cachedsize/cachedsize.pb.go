// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cachedsize/cachedsize.proto

package cachedsize

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	math "math"
	math_bits "math/bits"
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

type Foo struct {
	Field1               *Bar     `protobuf:"bytes,1,opt,name=field1" json:"field1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Foo) Reset()         { *m = Foo{} }
func (m *Foo) String() string { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()    {}
func (*Foo) Descriptor() ([]byte, []int) {
	return fileDescriptor_00ec3e3d9b714182, []int{0}
}
func (m *Foo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foo.Unmarshal(m, b)
}
func (m *Foo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Foo.Marshal(b, m, deterministic)
}
func (m *Foo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foo.Merge(m, src)
}
func (m *Foo) XXX_Size() int {
	return xxx_messageInfo_Foo.Size(m)
}
func (m *Foo) XXX_DiscardUnknown() {
	xxx_messageInfo_Foo.DiscardUnknown(m)
}

var xxx_messageInfo_Foo proto.InternalMessageInfo

func (m *Foo) GetField1() *Bar {
	if m != nil {
		return m.Field1
	}
	return nil
}

type Bar struct {
	Field2               bool     `protobuf:"varint,1,opt,name=field2" json:"field2"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bar) Reset()         { *m = Bar{} }
func (m *Bar) String() string { return proto.CompactTextString(m) }
func (*Bar) ProtoMessage()    {}
func (*Bar) Descriptor() ([]byte, []int) {
	return fileDescriptor_00ec3e3d9b714182, []int{1}
}
func (m *Bar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bar.Unmarshal(m, b)
}
func (m *Bar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bar.Marshal(b, m, deterministic)
}
func (m *Bar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bar.Merge(m, src)
}
func (m *Bar) XXX_Size() int {
	return xxx_messageInfo_Bar.Size(m)
}
func (m *Bar) XXX_DiscardUnknown() {
	xxx_messageInfo_Bar.DiscardUnknown(m)
}

var xxx_messageInfo_Bar proto.InternalMessageInfo

func (m *Bar) GetField2() bool {
	if m != nil {
		return m.Field2
	}
	return false
}

func init() {
	proto.RegisterType((*Foo)(nil), "cachedsize.Foo")
	proto.RegisterType((*Bar)(nil), "cachedsize.Bar")
}

func init() { proto.RegisterFile("cachedsize/cachedsize.proto", fileDescriptor_00ec3e3d9b714182) }

var fileDescriptor_00ec3e3d9b714182 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0x4e, 0x4c, 0xce,
	0x48, 0x4d, 0x29, 0xce, 0xac, 0x4a, 0xd5, 0x47, 0x30, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0xb8, 0x10, 0x22, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x61, 0x7d, 0x10, 0x0b, 0xa2, 0x42,
	0x49, 0x8f, 0x8b, 0xd9, 0x2d, 0x3f, 0x5f, 0x48, 0x9d, 0x8b, 0x2d, 0x2d, 0x33, 0x35, 0x27, 0xc5,
	0x50, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x88, 0x5f, 0x0f, 0xc9, 0x2c, 0xa7, 0xc4, 0xa2, 0x20,
	0xa8, 0xb4, 0x92, 0x32, 0x17, 0xb3, 0x53, 0x62, 0x91, 0x90, 0x0c, 0x54, 0xbd, 0x11, 0x58, 0x3d,
	0x87, 0x13, 0xcb, 0x89, 0x7b, 0xf2, 0x0c, 0x50, 0x45, 0x46, 0x4e, 0x12, 0x0f, 0x1e, 0xca, 0x31,
	0xae, 0x78, 0x24, 0xc7, 0x78, 0xe2, 0x91, 0x1c, 0xc3, 0x85, 0x47, 0x72, 0x0c, 0x0f, 0x1e, 0xc9,
	0x31, 0x7e, 0x78, 0x24, 0xc7, 0x08, 0x08, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x7e, 0x2f, 0x7c, 0xae,
	0x00, 0x00, 0x00,
}

func (this *Foo) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Foo)
	if !ok {
		that2, ok := that.(Foo)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Foo")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Foo but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Foo but is not nil && this == nil")
	}
	if !this.Field1.Equal(that1.Field1) {
		return fmt.Errorf("Field1 this(%v) Not Equal that(%v)", this.Field1, that1.Field1)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *Foo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Foo)
	if !ok {
		that2, ok := that.(Foo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Field1.Equal(that1.Field1) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Bar) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Bar)
	if !ok {
		that2, ok := that.(Bar)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Bar")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Bar but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Bar but is not nil && this == nil")
	}
	if this.Field2 != that1.Field2 {
		return fmt.Errorf("Field2 this(%v) Not Equal that(%v)", this.Field2, that1.Field2)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *Bar) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Bar)
	if !ok {
		that2, ok := that.(Bar)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Field2 != that1.Field2 {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *Foo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Field1 != nil {
		l = m.Field1.Size()
		n += 1 + l + sovCachedsize(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Bar) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 2
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCachedsize(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCachedsize(x uint64) (n int) {
	return sovCachedsize(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
