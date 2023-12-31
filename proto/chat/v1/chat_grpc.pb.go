// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.1
// source: proto/chat.proto

package chatv1

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

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	Join(ctx context.Context, in *User, opts ...grpc.CallOption) (*JoinResponse, error)
	SendMsg(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (*Empty, error)
	ReceiveMsg(ctx context.Context, in *ReceiveMsgRequest, opts ...grpc.CallOption) (ChatService_ReceiveMsgClient, error)
	GetAllUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserList, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Join(ctx context.Context, in *User, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := c.cc.Invoke(ctx, "/main.ChatService/join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SendMsg(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/main.ChatService/sendMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) ReceiveMsg(ctx context.Context, in *ReceiveMsgRequest, opts ...grpc.CallOption) (ChatService_ReceiveMsgClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], "/main.ChatService/receiveMsg", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceReceiveMsgClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_ReceiveMsgClient interface {
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type chatServiceReceiveMsgClient struct {
	grpc.ClientStream
}

func (x *chatServiceReceiveMsgClient) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) GetAllUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := c.cc.Invoke(ctx, "/main.ChatService/getAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	Join(context.Context, *User) (*JoinResponse, error)
	SendMsg(context.Context, *ChatMessage) (*Empty, error)
	ReceiveMsg(*ReceiveMsgRequest, ChatService_ReceiveMsgServer) error
	GetAllUsers(context.Context, *Empty) (*UserList, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) Join(context.Context, *User) (*JoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedChatServiceServer) SendMsg(context.Context, *ChatMessage) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsg not implemented")
}
func (UnimplementedChatServiceServer) ReceiveMsg(*ReceiveMsgRequest, ChatService_ReceiveMsgServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveMsg not implemented")
}
func (UnimplementedChatServiceServer) GetAllUsers(context.Context, *Empty) (*UserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ChatService/join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Join(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ChatService/sendMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SendMsg(ctx, req.(*ChatMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_ReceiveMsg_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiveMsgRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).ReceiveMsg(m, &chatServiceReceiveMsgServer{stream})
}

type ChatService_ReceiveMsgServer interface {
	Send(*ChatMessage) error
	grpc.ServerStream
}

type chatServiceReceiveMsgServer struct {
	grpc.ServerStream
}

func (x *chatServiceReceiveMsgServer) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ChatService/getAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetAllUsers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "join",
			Handler:    _ChatService_Join_Handler,
		},
		{
			MethodName: "sendMsg",
			Handler:    _ChatService_SendMsg_Handler,
		},
		{
			MethodName: "getAllUsers",
			Handler:    _ChatService_GetAllUsers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "receiveMsg",
			Handler:       _ChatService_ReceiveMsg_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/chat.proto",
}
