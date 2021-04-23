// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package consumer

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

// ConsumerClient is the client API for Consumer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConsumerClient interface {
	GetConsumerStatus(ctx context.Context, in *ConsumerRequest, opts ...grpc.CallOption) (*ConsumerResponse, error)
	Start(ctx context.Context, in *ConsumerRequest, opts ...grpc.CallOption) (*ConsumerResponse, error)
	Stop(ctx context.Context, in *ConsumerRequest, opts ...grpc.CallOption) (*ConsumerResponse, error)
}

type consumerClient struct {
	cc grpc.ClientConnInterface
}

func NewConsumerClient(cc grpc.ClientConnInterface) ConsumerClient {
	return &consumerClient{cc}
}

func (c *consumerClient) GetConsumerStatus(ctx context.Context, in *ConsumerRequest, opts ...grpc.CallOption) (*ConsumerResponse, error) {
	out := new(ConsumerResponse)
	err := c.cc.Invoke(ctx, "/v1.consumer.Consumer/GetConsumerStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerClient) Start(ctx context.Context, in *ConsumerRequest, opts ...grpc.CallOption) (*ConsumerResponse, error) {
	out := new(ConsumerResponse)
	err := c.cc.Invoke(ctx, "/v1.consumer.Consumer/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerClient) Stop(ctx context.Context, in *ConsumerRequest, opts ...grpc.CallOption) (*ConsumerResponse, error) {
	out := new(ConsumerResponse)
	err := c.cc.Invoke(ctx, "/v1.consumer.Consumer/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsumerServer is the server API for Consumer service.
// All implementations must embed UnimplementedConsumerServer
// for forward compatibility
type ConsumerServer interface {
	GetConsumerStatus(context.Context, *ConsumerRequest) (*ConsumerResponse, error)
	Start(context.Context, *ConsumerRequest) (*ConsumerResponse, error)
	Stop(context.Context, *ConsumerRequest) (*ConsumerResponse, error)
	mustEmbedUnimplementedConsumerServer()
}

// UnimplementedConsumerServer must be embedded to have forward compatible implementations.
type UnimplementedConsumerServer struct {
}

func (UnimplementedConsumerServer) GetConsumerStatus(context.Context, *ConsumerRequest) (*ConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsumerStatus not implemented")
}
func (UnimplementedConsumerServer) Start(context.Context, *ConsumerRequest) (*ConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedConsumerServer) Stop(context.Context, *ConsumerRequest) (*ConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedConsumerServer) mustEmbedUnimplementedConsumerServer() {}

// UnsafeConsumerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConsumerServer will
// result in compilation errors.
type UnsafeConsumerServer interface {
	mustEmbedUnimplementedConsumerServer()
}

func RegisterConsumerServer(s grpc.ServiceRegistrar, srv ConsumerServer) {
	s.RegisterService(&Consumer_ServiceDesc, srv)
}

func _Consumer_GetConsumerStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServer).GetConsumerStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.consumer.Consumer/GetConsumerStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServer).GetConsumerStatus(ctx, req.(*ConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Consumer_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.consumer.Consumer/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServer).Start(ctx, req.(*ConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Consumer_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.consumer.Consumer/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServer).Stop(ctx, req.(*ConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Consumer_ServiceDesc is the grpc.ServiceDesc for Consumer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Consumer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.consumer.Consumer",
	HandlerType: (*ConsumerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConsumerStatus",
			Handler:    _Consumer_GetConsumerStatus_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Consumer_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Consumer_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/consumer/consumer.proto",
}
