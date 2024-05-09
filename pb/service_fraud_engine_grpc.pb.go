// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: service_fraud_engine.proto

package pb

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

const (
	FraudEngine_CreateUser_FullMethodName = "/pb.FraudEngine/CreateUser"
)

// FraudEngineClient is the client API for FraudEngine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FraudEngineClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type fraudEngineClient struct {
	cc grpc.ClientConnInterface
}

func NewFraudEngineClient(cc grpc.ClientConnInterface) FraudEngineClient {
	return &fraudEngineClient{cc}
}

func (c *fraudEngineClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, FraudEngine_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FraudEngineServer is the server API for FraudEngine service.
// All implementations must embed UnimplementedFraudEngineServer
// for forward compatibility
type FraudEngineServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	mustEmbedUnimplementedFraudEngineServer()
}

// UnimplementedFraudEngineServer must be embedded to have forward compatible implementations.
type UnimplementedFraudEngineServer struct {
}

func (UnimplementedFraudEngineServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedFraudEngineServer) mustEmbedUnimplementedFraudEngineServer() {}

// UnsafeFraudEngineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FraudEngineServer will
// result in compilation errors.
type UnsafeFraudEngineServer interface {
	mustEmbedUnimplementedFraudEngineServer()
}

func RegisterFraudEngineServer(s grpc.ServiceRegistrar, srv FraudEngineServer) {
	s.RegisterService(&FraudEngine_ServiceDesc, srv)
}

func _FraudEngine_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FraudEngineServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FraudEngine_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FraudEngineServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FraudEngine_ServiceDesc is the grpc.ServiceDesc for FraudEngine service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FraudEngine_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.FraudEngine",
	HandlerType: (*FraudEngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _FraudEngine_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_fraud_engine.proto",
}
