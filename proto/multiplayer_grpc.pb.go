// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package __

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

// MultiplayerServiceClient is the client API for MultiplayerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MultiplayerServiceClient interface {
	SetPosition(ctx context.Context, in *SetPositionRequest, opts ...grpc.CallOption) (*SetPositionResponse, error)
	GetPositions(ctx context.Context, in *GetPositionsRequest, opts ...grpc.CallOption) (*GetPlayerPositions, error)
	BiDirectSetPositions(ctx context.Context, opts ...grpc.CallOption) (MultiplayerService_BiDirectSetPositionsClient, error)
}

type multiplayerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMultiplayerServiceClient(cc grpc.ClientConnInterface) MultiplayerServiceClient {
	return &multiplayerServiceClient{cc}
}

func (c *multiplayerServiceClient) SetPosition(ctx context.Context, in *SetPositionRequest, opts ...grpc.CallOption) (*SetPositionResponse, error) {
	out := new(SetPositionResponse)
	err := c.cc.Invoke(ctx, "/pulsey.protobuf.MultiplayerService/SetPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multiplayerServiceClient) GetPositions(ctx context.Context, in *GetPositionsRequest, opts ...grpc.CallOption) (*GetPlayerPositions, error) {
	out := new(GetPlayerPositions)
	err := c.cc.Invoke(ctx, "/pulsey.protobuf.MultiplayerService/GetPositions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multiplayerServiceClient) BiDirectSetPositions(ctx context.Context, opts ...grpc.CallOption) (MultiplayerService_BiDirectSetPositionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &MultiplayerService_ServiceDesc.Streams[0], "/pulsey.protobuf.MultiplayerService/BiDirectSetPositions", opts...)
	if err != nil {
		return nil, err
	}
	x := &multiplayerServiceBiDirectSetPositionsClient{stream}
	return x, nil
}

type MultiplayerService_BiDirectSetPositionsClient interface {
	Send(*SetPositionRequest) error
	Recv() (*GetPlayerPositions, error)
	grpc.ClientStream
}

type multiplayerServiceBiDirectSetPositionsClient struct {
	grpc.ClientStream
}

func (x *multiplayerServiceBiDirectSetPositionsClient) Send(m *SetPositionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *multiplayerServiceBiDirectSetPositionsClient) Recv() (*GetPlayerPositions, error) {
	m := new(GetPlayerPositions)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MultiplayerServiceServer is the server API for MultiplayerService service.
// All implementations must embed UnimplementedMultiplayerServiceServer
// for forward compatibility
type MultiplayerServiceServer interface {
	SetPosition(context.Context, *SetPositionRequest) (*SetPositionResponse, error)
	GetPositions(context.Context, *GetPositionsRequest) (*GetPlayerPositions, error)
	BiDirectSetPositions(MultiplayerService_BiDirectSetPositionsServer) error
	mustEmbedUnimplementedMultiplayerServiceServer()
}

// UnimplementedMultiplayerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMultiplayerServiceServer struct {
}

func (UnimplementedMultiplayerServiceServer) SetPosition(context.Context, *SetPositionRequest) (*SetPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPosition not implemented")
}
func (UnimplementedMultiplayerServiceServer) GetPositions(context.Context, *GetPositionsRequest) (*GetPlayerPositions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPositions not implemented")
}
func (UnimplementedMultiplayerServiceServer) BiDirectSetPositions(MultiplayerService_BiDirectSetPositionsServer) error {
	return status.Errorf(codes.Unimplemented, "method BiDirectSetPositions not implemented")
}
func (UnimplementedMultiplayerServiceServer) mustEmbedUnimplementedMultiplayerServiceServer() {}

// UnsafeMultiplayerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MultiplayerServiceServer will
// result in compilation errors.
type UnsafeMultiplayerServiceServer interface {
	mustEmbedUnimplementedMultiplayerServiceServer()
}

func RegisterMultiplayerServiceServer(s grpc.ServiceRegistrar, srv MultiplayerServiceServer) {
	s.RegisterService(&MultiplayerService_ServiceDesc, srv)
}

func _MultiplayerService_SetPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultiplayerServiceServer).SetPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulsey.protobuf.MultiplayerService/SetPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultiplayerServiceServer).SetPosition(ctx, req.(*SetPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultiplayerService_GetPositions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultiplayerServiceServer).GetPositions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulsey.protobuf.MultiplayerService/GetPositions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultiplayerServiceServer).GetPositions(ctx, req.(*GetPositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultiplayerService_BiDirectSetPositions_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MultiplayerServiceServer).BiDirectSetPositions(&multiplayerServiceBiDirectSetPositionsServer{stream})
}

type MultiplayerService_BiDirectSetPositionsServer interface {
	Send(*GetPlayerPositions) error
	Recv() (*SetPositionRequest, error)
	grpc.ServerStream
}

type multiplayerServiceBiDirectSetPositionsServer struct {
	grpc.ServerStream
}

func (x *multiplayerServiceBiDirectSetPositionsServer) Send(m *GetPlayerPositions) error {
	return x.ServerStream.SendMsg(m)
}

func (x *multiplayerServiceBiDirectSetPositionsServer) Recv() (*SetPositionRequest, error) {
	m := new(SetPositionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MultiplayerService_ServiceDesc is the grpc.ServiceDesc for MultiplayerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MultiplayerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pulsey.protobuf.MultiplayerService",
	HandlerType: (*MultiplayerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetPosition",
			Handler:    _MultiplayerService_SetPosition_Handler,
		},
		{
			MethodName: "GetPositions",
			Handler:    _MultiplayerService_GetPositions_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BiDirectSetPositions",
			Handler:       _MultiplayerService_BiDirectSetPositions_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/multiplayer.proto",
}
