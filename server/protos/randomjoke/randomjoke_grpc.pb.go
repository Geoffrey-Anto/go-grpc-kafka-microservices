// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: protos/randomjoke.proto

package randomjoke

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
	RandomJokeService_GetRandomJoke_FullMethodName = "/RandomJokeService/GetRandomJoke"
)

// RandomJokeServiceClient is the client API for RandomJokeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RandomJokeServiceClient interface {
	GetRandomJoke(ctx context.Context, in *RandomJokeRequest, opts ...grpc.CallOption) (*RandomJokeResponse, error)
}

type randomJokeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRandomJokeServiceClient(cc grpc.ClientConnInterface) RandomJokeServiceClient {
	return &randomJokeServiceClient{cc}
}

func (c *randomJokeServiceClient) GetRandomJoke(ctx context.Context, in *RandomJokeRequest, opts ...grpc.CallOption) (*RandomJokeResponse, error) {
	out := new(RandomJokeResponse)
	err := c.cc.Invoke(ctx, RandomJokeService_GetRandomJoke_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RandomJokeServiceServer is the server API for RandomJokeService service.
// All implementations must embed UnimplementedRandomJokeServiceServer
// for forward compatibility
type RandomJokeServiceServer interface {
	GetRandomJoke(context.Context, *RandomJokeRequest) (*RandomJokeResponse, error)
	mustEmbedUnimplementedRandomJokeServiceServer()
}

// UnimplementedRandomJokeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRandomJokeServiceServer struct {
}

func (UnimplementedRandomJokeServiceServer) GetRandomJoke(context.Context, *RandomJokeRequest) (*RandomJokeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandomJoke not implemented")
}
func (UnimplementedRandomJokeServiceServer) mustEmbedUnimplementedRandomJokeServiceServer() {}

// UnsafeRandomJokeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RandomJokeServiceServer will
// result in compilation errors.
type UnsafeRandomJokeServiceServer interface {
	mustEmbedUnimplementedRandomJokeServiceServer()
}

func RegisterRandomJokeServiceServer(s grpc.ServiceRegistrar, srv RandomJokeServiceServer) {
	s.RegisterService(&RandomJokeService_ServiceDesc, srv)
}

func _RandomJokeService_GetRandomJoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandomJokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomJokeServiceServer).GetRandomJoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RandomJokeService_GetRandomJoke_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomJokeServiceServer).GetRandomJoke(ctx, req.(*RandomJokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RandomJokeService_ServiceDesc is the grpc.ServiceDesc for RandomJokeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RandomJokeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RandomJokeService",
	HandlerType: (*RandomJokeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRandomJoke",
			Handler:    _RandomJokeService_GetRandomJoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/randomjoke.proto",
}
