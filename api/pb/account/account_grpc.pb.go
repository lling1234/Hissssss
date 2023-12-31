// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: account.proto

package account

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

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpReply, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInReply, error)
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectReply, error)
	DisConnect(ctx context.Context, in *DisConnectRequest, opts ...grpc.CallOption) (*DisConnectReply, error)
}

type accountClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountClient(cc grpc.ClientConnInterface) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpReply, error) {
	out := new(SignUpReply)
	err := c.cc.Invoke(ctx, "/account.Account/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInReply, error) {
	out := new(SignInReply)
	err := c.cc.Invoke(ctx, "/account.Account/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectReply, error) {
	out := new(ConnectReply)
	err := c.cc.Invoke(ctx, "/account.Account/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) DisConnect(ctx context.Context, in *DisConnectRequest, opts ...grpc.CallOption) (*DisConnectReply, error) {
	out := new(DisConnectReply)
	err := c.cc.Invoke(ctx, "/account.Account/DisConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
// All implementations must embed UnimplementedAccountServer
// for forward compatibility
type AccountServer interface {
	SignUp(context.Context, *SignUpRequest) (*SignUpReply, error)
	SignIn(context.Context, *SignInRequest) (*SignInReply, error)
	Connect(context.Context, *ConnectRequest) (*ConnectReply, error)
	DisConnect(context.Context, *DisConnectRequest) (*DisConnectReply, error)
	mustEmbedUnimplementedAccountServer()
}

// UnimplementedAccountServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (UnimplementedAccountServer) SignUp(context.Context, *SignUpRequest) (*SignUpReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedAccountServer) SignIn(context.Context, *SignInRequest) (*SignInReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedAccountServer) Connect(context.Context, *ConnectRequest) (*ConnectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedAccountServer) DisConnect(context.Context, *DisConnectRequest) (*DisConnectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisConnect not implemented")
}
func (UnimplementedAccountServer) mustEmbedUnimplementedAccountServer() {}

// UnsafeAccountServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServer will
// result in compilation errors.
type UnsafeAccountServer interface {
	mustEmbedUnimplementedAccountServer()
}

func RegisterAccountServer(s grpc.ServiceRegistrar, srv AccountServer) {
	s.RegisterService(&Account_ServiceDesc, srv)
}

func _Account_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_DisConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).DisConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/DisConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).DisConnect(ctx, req.(*DisConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Account_ServiceDesc is the grpc.ServiceDesc for Account service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Account_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _Account_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _Account_SignIn_Handler,
		},
		{
			MethodName: "Connect",
			Handler:    _Account_Connect_Handler,
		},
		{
			MethodName: "DisConnect",
			Handler:    _Account_DisConnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}
