// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: beef.proto

package beefpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BeefService_Get_FullMethodName        = "/beef.BeefService/Get"
	BeefService_GetSummary_FullMethodName = "/beef.BeefService/GetSummary"
)

// BeefServiceClient is the client API for BeefService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BeefServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetSummary(ctx context.Context, in *GetSummaryRequest, opts ...grpc.CallOption) (*GetSummaryResponse, error)
}

type beefServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBeefServiceClient(cc grpc.ClientConnInterface) BeefServiceClient {
	return &beefServiceClient{cc}
}

func (c *beefServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, BeefService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beefServiceClient) GetSummary(ctx context.Context, in *GetSummaryRequest, opts ...grpc.CallOption) (*GetSummaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSummaryResponse)
	err := c.cc.Invoke(ctx, BeefService_GetSummary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BeefServiceServer is the server API for BeefService service.
// All implementations must embed UnimplementedBeefServiceServer
// for forward compatibility.
type BeefServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetSummary(context.Context, *GetSummaryRequest) (*GetSummaryResponse, error)
	mustEmbedUnimplementedBeefServiceServer()
}

// UnimplementedBeefServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBeefServiceServer struct{}

func (UnimplementedBeefServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBeefServiceServer) GetSummary(context.Context, *GetSummaryRequest) (*GetSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSummary not implemented")
}
func (UnimplementedBeefServiceServer) mustEmbedUnimplementedBeefServiceServer() {}
func (UnimplementedBeefServiceServer) testEmbeddedByValue()                     {}

// UnsafeBeefServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BeefServiceServer will
// result in compilation errors.
type UnsafeBeefServiceServer interface {
	mustEmbedUnimplementedBeefServiceServer()
}

func RegisterBeefServiceServer(s grpc.ServiceRegistrar, srv BeefServiceServer) {
	// If the following call pancis, it indicates UnimplementedBeefServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BeefService_ServiceDesc, srv)
}

func _BeefService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeefServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BeefService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeefServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeefService_GetSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeefServiceServer).GetSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BeefService_GetSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeefServiceServer).GetSummary(ctx, req.(*GetSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BeefService_ServiceDesc is the grpc.ServiceDesc for BeefService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BeefService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "beef.BeefService",
	HandlerType: (*BeefServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _BeefService_Get_Handler,
		},
		{
			MethodName: "GetSummary",
			Handler:    _BeefService_GetSummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "beef.proto",
}
