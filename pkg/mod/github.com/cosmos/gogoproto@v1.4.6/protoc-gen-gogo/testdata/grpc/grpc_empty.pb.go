// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: grpc/grpc_empty.proto

package testing

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
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

func init() { proto.RegisterFile("grpc/grpc_empty.proto", fileDescriptor_c580a37f1c90e9b1) }

var fileDescriptor_c580a37f1c90e9b1 = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0x2f, 0x2a, 0x48,
	0xd6, 0x07, 0x11, 0xf1, 0xa9, 0xb9, 0x05, 0x25, 0x95, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0x2c, 0x20, 0x11, 0x23, 0x3e, 0x2e, 0x1e, 0x57, 0x90, 0x60, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72,
	0xaa, 0x93, 0x73, 0x94, 0x63, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e,
	0x72, 0x7e, 0x71, 0x6e, 0x7e, 0xb1, 0x7e, 0x7a, 0x7e, 0x7a, 0x3e, 0x58, 0x93, 0x3e, 0x98, 0x4c,
	0xd6, 0x4d, 0x4f, 0xcd, 0xd3, 0x05, 0x09, 0xea, 0x97, 0xa4, 0x16, 0x97, 0xa4, 0x24, 0x96, 0x24,
	0x82, 0x8d, 0xb7, 0x06, 0xf1, 0x32, 0xf3, 0xd2, 0x93, 0xd8, 0xc0, 0xca, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xdb, 0xd8, 0xed, 0x92, 0x7a, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EmptyServiceClient is the client API for EmptyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmptyServiceClient interface {
}

type emptyServiceClient struct {
	cc grpc1.ClientConn
}

func NewEmptyServiceClient(cc grpc1.ClientConn) EmptyServiceClient {
	return &emptyServiceClient{cc}
}

// EmptyServiceServer is the server API for EmptyService service.
type EmptyServiceServer interface {
}

// UnimplementedEmptyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEmptyServiceServer struct {
}

func RegisterEmptyServiceServer(s grpc1.Server, srv EmptyServiceServer) {
	s.RegisterService(&_EmptyService_serviceDesc, srv)
}

var _EmptyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.EmptyService",
	HandlerType: (*EmptyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "grpc/grpc_empty.proto",
}
