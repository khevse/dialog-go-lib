// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test.proto

package test

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xce, 0x48, 0x4d, 0xce, 0x4e, 0x2d, 0x92,
	0x92, 0x4e, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x0b, 0x27, 0x95, 0xa6, 0xe9, 0xa7, 0xe6,
	0x16, 0x94, 0x54, 0x42, 0x54, 0x19, 0x39, 0x73, 0xb1, 0x3b, 0x43, 0xd4, 0x09, 0x59, 0x70, 0xb1,
	0x04, 0x64, 0xe6, 0xa5, 0x0b, 0x89, 0xe9, 0x41, 0x34, 0xe8, 0xc1, 0x34, 0xe8, 0xb9, 0x82, 0x34,
	0x48, 0xe1, 0x10, 0x57, 0x62, 0x70, 0x92, 0x3b, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39,
	0x86, 0x28, 0x16, 0x90, 0x83, 0x92, 0xd8, 0xc0, 0x3a, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xbd, 0xc2, 0x6c, 0xec, 0x9f, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CheckerClient is the client API for Checker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CheckerClient interface {
	Ping(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error)
}

type checkerClient struct {
	cc *grpc.ClientConn
}

func NewCheckerClient(cc *grpc.ClientConn) CheckerClient {
	return &checkerClient{cc}
}

func (c *checkerClient) Ping(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/checker.Checker/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckerServer is the server API for Checker service.
type CheckerServer interface {
	Ping(context.Context, *types.Empty) (*types.Empty, error)
}

func RegisterCheckerServer(s *grpc.Server, srv CheckerServer) {
	s.RegisterService(&_Checker_serviceDesc, srv)
}

func _Checker_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/checker.Checker/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).Ping(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Checker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "checker.Checker",
	HandlerType: (*CheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Checker_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
