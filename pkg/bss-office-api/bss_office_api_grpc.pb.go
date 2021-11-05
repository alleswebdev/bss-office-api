// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package bss_office_api

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

// BssOfficeApiServiceClient is the client API for BssOfficeApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BssOfficeApiServiceClient interface {
	// DescribeOfficeV1 - Describe a office
	DescribeOfficeV1(ctx context.Context, in *DescribeOfficeV1Request, opts ...grpc.CallOption) (*DescribeOfficeV1Response, error)
	// CreateOfficeV1 - Create new office
	CreateOfficeV1(ctx context.Context, in *CreateOfficeV1Request, opts ...grpc.CallOption) (*CreateOfficeV1Response, error)
	// RemoveOfficeV1 - delete the office by id
	RemoveOfficeV1(ctx context.Context, in *RemoveOfficeV1Request, opts ...grpc.CallOption) (*RemoveOfficeV1Response, error)
	// ListOfficeV1 - list of offices
	ListOfficesV1(ctx context.Context, in *ListOfficesV1Request, opts ...grpc.CallOption) (*ListOfficesV1Response, error)
}

type bssOfficeApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBssOfficeApiServiceClient(cc grpc.ClientConnInterface) BssOfficeApiServiceClient {
	return &bssOfficeApiServiceClient{cc}
}

func (c *bssOfficeApiServiceClient) DescribeOfficeV1(ctx context.Context, in *DescribeOfficeV1Request, opts ...grpc.CallOption) (*DescribeOfficeV1Response, error) {
	out := new(DescribeOfficeV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.bss_office_api.v1.BssOfficeApiService/DescribeOfficeV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bssOfficeApiServiceClient) CreateOfficeV1(ctx context.Context, in *CreateOfficeV1Request, opts ...grpc.CallOption) (*CreateOfficeV1Response, error) {
	out := new(CreateOfficeV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.bss_office_api.v1.BssOfficeApiService/CreateOfficeV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bssOfficeApiServiceClient) RemoveOfficeV1(ctx context.Context, in *RemoveOfficeV1Request, opts ...grpc.CallOption) (*RemoveOfficeV1Response, error) {
	out := new(RemoveOfficeV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.bss_office_api.v1.BssOfficeApiService/RemoveOfficeV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bssOfficeApiServiceClient) ListOfficesV1(ctx context.Context, in *ListOfficesV1Request, opts ...grpc.CallOption) (*ListOfficesV1Response, error) {
	out := new(ListOfficesV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.bss_office_api.v1.BssOfficeApiService/ListOfficesV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BssOfficeApiServiceServer is the server API for BssOfficeApiService service.
// All implementations must embed UnimplementedBssOfficeApiServiceServer
// for forward compatibility
type BssOfficeApiServiceServer interface {
	// DescribeOfficeV1 - Describe a office
	DescribeOfficeV1(context.Context, *DescribeOfficeV1Request) (*DescribeOfficeV1Response, error)
	// CreateOfficeV1 - Create new office
	CreateOfficeV1(context.Context, *CreateOfficeV1Request) (*CreateOfficeV1Response, error)
	// RemoveOfficeV1 - delete the office by id
	RemoveOfficeV1(context.Context, *RemoveOfficeV1Request) (*RemoveOfficeV1Response, error)
	// ListOfficeV1 - list of offices
	ListOfficesV1(context.Context, *ListOfficesV1Request) (*ListOfficesV1Response, error)
	mustEmbedUnimplementedBssOfficeApiServiceServer()
}

// UnimplementedBssOfficeApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBssOfficeApiServiceServer struct {
}

func (UnimplementedBssOfficeApiServiceServer) DescribeOfficeV1(context.Context, *DescribeOfficeV1Request) (*DescribeOfficeV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeOfficeV1 not implemented")
}
func (UnimplementedBssOfficeApiServiceServer) CreateOfficeV1(context.Context, *CreateOfficeV1Request) (*CreateOfficeV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOfficeV1 not implemented")
}
func (UnimplementedBssOfficeApiServiceServer) RemoveOfficeV1(context.Context, *RemoveOfficeV1Request) (*RemoveOfficeV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveOfficeV1 not implemented")
}
func (UnimplementedBssOfficeApiServiceServer) ListOfficesV1(context.Context, *ListOfficesV1Request) (*ListOfficesV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOfficesV1 not implemented")
}
func (UnimplementedBssOfficeApiServiceServer) mustEmbedUnimplementedBssOfficeApiServiceServer() {}

// UnsafeBssOfficeApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BssOfficeApiServiceServer will
// result in compilation errors.
type UnsafeBssOfficeApiServiceServer interface {
	mustEmbedUnimplementedBssOfficeApiServiceServer()
}

func RegisterBssOfficeApiServiceServer(s grpc.ServiceRegistrar, srv BssOfficeApiServiceServer) {
	s.RegisterService(&BssOfficeApiService_ServiceDesc, srv)
}

func _BssOfficeApiService_DescribeOfficeV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeOfficeV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BssOfficeApiServiceServer).DescribeOfficeV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.bss_office_api.v1.BssOfficeApiService/DescribeOfficeV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BssOfficeApiServiceServer).DescribeOfficeV1(ctx, req.(*DescribeOfficeV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BssOfficeApiService_CreateOfficeV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOfficeV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BssOfficeApiServiceServer).CreateOfficeV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.bss_office_api.v1.BssOfficeApiService/CreateOfficeV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BssOfficeApiServiceServer).CreateOfficeV1(ctx, req.(*CreateOfficeV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BssOfficeApiService_RemoveOfficeV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveOfficeV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BssOfficeApiServiceServer).RemoveOfficeV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.bss_office_api.v1.BssOfficeApiService/RemoveOfficeV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BssOfficeApiServiceServer).RemoveOfficeV1(ctx, req.(*RemoveOfficeV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BssOfficeApiService_ListOfficesV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOfficesV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BssOfficeApiServiceServer).ListOfficesV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.bss_office_api.v1.BssOfficeApiService/ListOfficesV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BssOfficeApiServiceServer).ListOfficesV1(ctx, req.(*ListOfficesV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// BssOfficeApiService_ServiceDesc is the grpc.ServiceDesc for BssOfficeApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BssOfficeApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ozonmp.bss_office_api.v1.BssOfficeApiService",
	HandlerType: (*BssOfficeApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeOfficeV1",
			Handler:    _BssOfficeApiService_DescribeOfficeV1_Handler,
		},
		{
			MethodName: "CreateOfficeV1",
			Handler:    _BssOfficeApiService_CreateOfficeV1_Handler,
		},
		{
			MethodName: "RemoveOfficeV1",
			Handler:    _BssOfficeApiService_RemoveOfficeV1_Handler,
		},
		{
			MethodName: "ListOfficesV1",
			Handler:    _BssOfficeApiService_ListOfficesV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ozonmp/bss_office_api/v1/bss_office_api.proto",
}
