// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: parser/parser.proto

package common_parser

import (
	context "context"
	schema "github.com/MyWeHub/plugin-sdk/gen/schema"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ParserService_Parse_FullMethodName = "/parser.ParserService/Parse"
)

// ParserServiceClient is the client API for ParserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ParserServiceClient interface {
	Parse(ctx context.Context, in *ParserDataRequest, opts ...grpc.CallOption) (*schema.Versions, error)
}

type parserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewParserServiceClient(cc grpc.ClientConnInterface) ParserServiceClient {
	return &parserServiceClient{cc}
}

func (c *parserServiceClient) Parse(ctx context.Context, in *ParserDataRequest, opts ...grpc.CallOption) (*schema.Versions, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(schema.Versions)
	err := c.cc.Invoke(ctx, ParserService_Parse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParserServiceServer is the server API for ParserService service.
// All implementations must embed UnimplementedParserServiceServer
// for forward compatibility.
type ParserServiceServer interface {
	Parse(context.Context, *ParserDataRequest) (*schema.Versions, error)
	mustEmbedUnimplementedParserServiceServer()
}

// UnimplementedParserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedParserServiceServer struct{}

func (UnimplementedParserServiceServer) Parse(context.Context, *ParserDataRequest) (*schema.Versions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Parse not implemented")
}
func (UnimplementedParserServiceServer) mustEmbedUnimplementedParserServiceServer() {}
func (UnimplementedParserServiceServer) testEmbeddedByValue()                       {}

// UnsafeParserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParserServiceServer will
// result in compilation errors.
type UnsafeParserServiceServer interface {
	mustEmbedUnimplementedParserServiceServer()
}

func RegisterParserServiceServer(s grpc.ServiceRegistrar, srv ParserServiceServer) {
	// If the following call pancis, it indicates UnimplementedParserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ParserService_ServiceDesc, srv)
}

func _ParserService_Parse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParserDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParserServiceServer).Parse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParserService_Parse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParserServiceServer).Parse(ctx, req.(*ParserDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ParserService_ServiceDesc is the grpc.ServiceDesc for ParserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ParserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "parser.ParserService",
	HandlerType: (*ParserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Parse",
			Handler:    _ParserService_Parse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "parser/parser.proto",
}
