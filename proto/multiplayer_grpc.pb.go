// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// MultiplayerServiceClient is the client API for MultiplayerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MultiplayerServiceClient interface {
	RoomConnect(ctx context.Context, in *RoomConnectRequest, opts ...grpc.CallOption) (*RoomConnectResponse, error)
	RoomStream(ctx context.Context, opts ...grpc.CallOption) (MultiplayerService_RoomStreamClient, error)
}

type multiplayerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMultiplayerServiceClient(cc grpc.ClientConnInterface) MultiplayerServiceClient {
	return &multiplayerServiceClient{cc}
}

func (c *multiplayerServiceClient) RoomConnect(ctx context.Context, in *RoomConnectRequest, opts ...grpc.CallOption) (*RoomConnectResponse, error) {
	out := new(RoomConnectResponse)
	err := c.cc.Invoke(ctx, "/MultiplayerService/RoomConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multiplayerServiceClient) RoomStream(ctx context.Context, opts ...grpc.CallOption) (MultiplayerService_RoomStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &MultiplayerService_ServiceDesc.Streams[0], "/MultiplayerService/RoomStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &multiplayerServiceRoomStreamClient{stream}
	return x, nil
}

type MultiplayerService_RoomStreamClient interface {
	Send(*RoomStreamRequest) error
	Recv() (*RoomStreamResponse, error)
	grpc.ClientStream
}

type multiplayerServiceRoomStreamClient struct {
	grpc.ClientStream
}

func (x *multiplayerServiceRoomStreamClient) Send(m *RoomStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *multiplayerServiceRoomStreamClient) Recv() (*RoomStreamResponse, error) {
	m := new(RoomStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MultiplayerServiceServer is the server API for MultiplayerService service.
// All implementations must embed UnimplementedMultiplayerServiceServer
// for forward compatibility
type MultiplayerServiceServer interface {
	RoomConnect(context.Context, *RoomConnectRequest) (*RoomConnectResponse, error)
	RoomStream(MultiplayerService_RoomStreamServer) error
	mustEmbedUnimplementedMultiplayerServiceServer()
}

// UnimplementedMultiplayerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMultiplayerServiceServer struct {
}

func (UnimplementedMultiplayerServiceServer) RoomConnect(context.Context, *RoomConnectRequest) (*RoomConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoomConnect not implemented")
}
func (UnimplementedMultiplayerServiceServer) RoomStream(MultiplayerService_RoomStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method RoomStream not implemented")
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

func _MultiplayerService_RoomConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultiplayerServiceServer).RoomConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MultiplayerService/RoomConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultiplayerServiceServer).RoomConnect(ctx, req.(*RoomConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultiplayerService_RoomStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MultiplayerServiceServer).RoomStream(&multiplayerServiceRoomStreamServer{stream})
}

type MultiplayerService_RoomStreamServer interface {
	Send(*RoomStreamResponse) error
	Recv() (*RoomStreamRequest, error)
	grpc.ServerStream
}

type multiplayerServiceRoomStreamServer struct {
	grpc.ServerStream
}

func (x *multiplayerServiceRoomStreamServer) Send(m *RoomStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *multiplayerServiceRoomStreamServer) Recv() (*RoomStreamRequest, error) {
	m := new(RoomStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MultiplayerService_ServiceDesc is the grpc.ServiceDesc for MultiplayerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MultiplayerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MultiplayerService",
	HandlerType: (*MultiplayerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RoomConnect",
			Handler:    _MultiplayerService_RoomConnect_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RoomStream",
			Handler:       _MultiplayerService_RoomStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/multiplayer.proto",
}
