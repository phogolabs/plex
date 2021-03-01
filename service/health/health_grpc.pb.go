// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package health

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

// HeartbeatAPIClient is the client API for HeartbeatAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeartbeatAPIClient interface {
	// CheckLive checks the live state
	CheckLive(ctx context.Context, in *CheckLiveRequest, opts ...grpc.CallOption) (*CheckLiveResponse, error)
	// CheckReady checks the ready state
	CheckReady(ctx context.Context, in *CheckReadyRequest, opts ...grpc.CallOption) (*CheckReadyResponse, error)
}

type heartbeatAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewHeartbeatAPIClient(cc grpc.ClientConnInterface) HeartbeatAPIClient {
	return &heartbeatAPIClient{cc}
}

func (c *heartbeatAPIClient) CheckLive(ctx context.Context, in *CheckLiveRequest, opts ...grpc.CallOption) (*CheckLiveResponse, error) {
	out := new(CheckLiveResponse)
	err := c.cc.Invoke(ctx, "/phogolabs.plex.health.HeartbeatAPI/CheckLive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heartbeatAPIClient) CheckReady(ctx context.Context, in *CheckReadyRequest, opts ...grpc.CallOption) (*CheckReadyResponse, error) {
	out := new(CheckReadyResponse)
	err := c.cc.Invoke(ctx, "/phogolabs.plex.health.HeartbeatAPI/CheckReady", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeartbeatAPIServer is the server API for HeartbeatAPI service.
// All implementations should embed UnimplementedHeartbeatAPIServer
// for forward compatibility
type HeartbeatAPIServer interface {
	// CheckLive checks the live state
	CheckLive(context.Context, *CheckLiveRequest) (*CheckLiveResponse, error)
	// CheckReady checks the ready state
	CheckReady(context.Context, *CheckReadyRequest) (*CheckReadyResponse, error)
}

// UnimplementedHeartbeatAPIServer should be embedded to have forward compatible implementations.
type UnimplementedHeartbeatAPIServer struct {
}

func (UnimplementedHeartbeatAPIServer) CheckLive(context.Context, *CheckLiveRequest) (*CheckLiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckLive not implemented")
}
func (UnimplementedHeartbeatAPIServer) CheckReady(context.Context, *CheckReadyRequest) (*CheckReadyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckReady not implemented")
}

// UnsafeHeartbeatAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeartbeatAPIServer will
// result in compilation errors.
type UnsafeHeartbeatAPIServer interface {
	mustEmbedUnimplementedHeartbeatAPIServer()
}

func RegisterHeartbeatAPIServer(s grpc.ServiceRegistrar, srv HeartbeatAPIServer) {
	s.RegisterService(&HeartbeatAPI_ServiceDesc, srv)
}

func _HeartbeatAPI_CheckLive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckLiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartbeatAPIServer).CheckLive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/phogolabs.plex.health.HeartbeatAPI/CheckLive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartbeatAPIServer).CheckLive(ctx, req.(*CheckLiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeartbeatAPI_CheckReady_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckReadyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartbeatAPIServer).CheckReady(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/phogolabs.plex.health.HeartbeatAPI/CheckReady",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartbeatAPIServer).CheckReady(ctx, req.(*CheckReadyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HeartbeatAPI_ServiceDesc is the grpc.ServiceDesc for HeartbeatAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HeartbeatAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "phogolabs.plex.health.HeartbeatAPI",
	HandlerType: (*HeartbeatAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckLive",
			Handler:    _HeartbeatAPI_CheckLive_Handler,
		},
		{
			MethodName: "CheckReady",
			Handler:    _HeartbeatAPI_CheckReady_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "health.proto",
}