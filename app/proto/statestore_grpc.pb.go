// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: app/proto/statestore.proto

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

// StateStoreClient is the client API for StateStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StateStoreClient interface {
	Store(ctx context.Context, in *State, opts ...grpc.CallOption) (*Confirm, error)
	Retrieve(ctx context.Context, in *State, opts ...grpc.CallOption) (*State, error)
}

type stateStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewStateStoreClient(cc grpc.ClientConnInterface) StateStoreClient {
	return &stateStoreClient{cc}
}

func (c *stateStoreClient) Store(ctx context.Context, in *State, opts ...grpc.CallOption) (*Confirm, error) {
	out := new(Confirm)
	err := c.cc.Invoke(ctx, "/dapr.StateStore/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateStoreClient) Retrieve(ctx context.Context, in *State, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/dapr.StateStore/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StateStoreServer is the server API for StateStore service.
// All implementations must embed UnimplementedStateStoreServer
// for forward compatibility
type StateStoreServer interface {
	Store(context.Context, *State) (*Confirm, error)
	Retrieve(context.Context, *State) (*State, error)
	mustEmbedUnimplementedStateStoreServer()
}

// UnimplementedStateStoreServer must be embedded to have forward compatible implementations.
type UnimplementedStateStoreServer struct {
}

func (UnimplementedStateStoreServer) Store(context.Context, *State) (*Confirm, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (UnimplementedStateStoreServer) Retrieve(context.Context, *State) (*State, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedStateStoreServer) mustEmbedUnimplementedStateStoreServer() {}

// UnsafeStateStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StateStoreServer will
// result in compilation errors.
type UnsafeStateStoreServer interface {
	mustEmbedUnimplementedStateStoreServer()
}

func RegisterStateStoreServer(s grpc.ServiceRegistrar, srv StateStoreServer) {
	s.RegisterService(&StateStore_ServiceDesc, srv)
}

func _StateStore_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(State)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateStoreServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.StateStore/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateStoreServer).Store(ctx, req.(*State))
	}
	return interceptor(ctx, in, info, handler)
}

func _StateStore_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(State)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateStoreServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.StateStore/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateStoreServer).Retrieve(ctx, req.(*State))
	}
	return interceptor(ctx, in, info, handler)
}

// StateStore_ServiceDesc is the grpc.ServiceDesc for StateStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StateStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dapr.StateStore",
	HandlerType: (*StateStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _StateStore_Store_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _StateStore_Retrieve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/proto/statestore.proto",
}