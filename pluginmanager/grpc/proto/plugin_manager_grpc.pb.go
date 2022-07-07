// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: plugin_manager.proto

package proto

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

// PluginManagerClient is the client API for PluginManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginManagerClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error)
}

type pluginManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginManagerClient(cc grpc.ClientConnInterface) PluginManagerClient {
	return &pluginManagerClient{cc}
}

func (c *pluginManagerClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/proto.PluginManager/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginManagerClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error) {
	out := new(ShutdownResponse)
	err := c.cc.Invoke(ctx, "/proto.PluginManager/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginManagerServer is the server API for PluginManager service.
// All implementations must embed UnimplementedPluginManagerServer
// for forward compatibility
type PluginManagerServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error)
	mustEmbedUnimplementedPluginManagerServer()
}

// UnimplementedPluginManagerServer must be embedded to have forward compatible implementations.
type UnimplementedPluginManagerServer struct {
}

func (UnimplementedPluginManagerServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPluginManagerServer) Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}
func (UnimplementedPluginManagerServer) mustEmbedUnimplementedPluginManagerServer() {}

// UnsafePluginManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PluginManagerServer will
// result in compilation errors.
type UnsafePluginManagerServer interface {
	mustEmbedUnimplementedPluginManagerServer()
}

func RegisterPluginManagerServer(s grpc.ServiceRegistrar, srv PluginManagerServer) {
	s.RegisterService(&PluginManager_ServiceDesc, srv)
}

func _PluginManager_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginManagerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PluginManager/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginManagerServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginManager_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginManagerServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PluginManager/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginManagerServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PluginManager_ServiceDesc is the grpc.ServiceDesc for PluginManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PluginManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PluginManager",
	HandlerType: (*PluginManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PluginManager_Get_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _PluginManager_Shutdown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugin_manager.proto",
}
