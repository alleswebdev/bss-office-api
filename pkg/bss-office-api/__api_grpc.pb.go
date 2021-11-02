// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package omp_template_api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OmpTemplateApiServiceClient is the client API for OmpTemplateApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OmpTemplateApiServiceClient interface {
	// DescribeTemplateV1 - Describe a template
	DescribeTemplateV1(ctx context.Context, in *DescribeTemplateV1Request, opts ...grpc.CallOption) (*DescribeTemplateV1Response, error)
}

type ompTemplateApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOmpTemplateApiServiceClient(cc grpc.ClientConnInterface) OmpTemplateApiServiceClient {
	return &ompTemplateApiServiceClient{cc}
}

func (c *ompTemplateApiServiceClient) DescribeTemplateV1(ctx context.Context, in *DescribeTemplateV1Request, opts ...grpc.CallOption) (*DescribeTemplateV1Response, error) {
	out := new(DescribeTemplateV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.omp_template_api.v1.OmpTemplateApiService/DescribeTemplateV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OmpTemplateApiServiceServer is the server API for OmpTemplateApiService service.
// All implementations must embed UnimplementedOmpTemplateApiServiceServer
// for forward compatibility
type OmpTemplateApiServiceServer interface {
	// DescribeTemplateV1 - Describe a template
	DescribeTemplateV1(context.Context, *DescribeTemplateV1Request) (*DescribeTemplateV1Response, error)
	mustEmbedUnimplementedOmpTemplateApiServiceServer()
}

// UnimplementedOmpTemplateApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOmpTemplateApiServiceServer struct {
}

func (UnimplementedOmpTemplateApiServiceServer) DescribeTemplateV1(context.Context, *DescribeTemplateV1Request) (*DescribeTemplateV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeTemplateV1 not implemented")
}
func (UnimplementedOmpTemplateApiServiceServer) mustEmbedUnimplementedOmpTemplateApiServiceServer() {}

// UnsafeOmpTemplateApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OmpTemplateApiServiceServer will
// result in compilation errors.
type UnsafeOmpTemplateApiServiceServer interface {
	mustEmbedUnimplementedOmpTemplateApiServiceServer()
}

func RegisterOmpTemplateApiServiceServer(s grpc.ServiceRegistrar, srv OmpTemplateApiServiceServer) {
	s.RegisterService(&OmpTemplateApiService_ServiceDesc, srv)
}

func _OmpTemplateApiService_DescribeTemplateV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeTemplateV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmpTemplateApiServiceServer).DescribeTemplateV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.omp_template_api.v1.OmpTemplateApiService/DescribeTemplateV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmpTemplateApiServiceServer).DescribeTemplateV1(ctx, req.(*DescribeTemplateV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OmpTemplateApiService_ServiceDesc is the grpc.ServiceDesc for OmpTemplateApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OmpTemplateApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ozonmp.omp_template_api.v1.OmpTemplateApiService",
	HandlerType: (*OmpTemplateApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeTemplateV1",
			Handler:    _OmpTemplateApiService_DescribeTemplateV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ozonmp/__api/v1/__api.proto",
}
