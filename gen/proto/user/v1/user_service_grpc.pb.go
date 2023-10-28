// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: user/v1/user_service.proto

package userV1

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
	UserService_GetInfo_FullMethodName              = "/user.v1.UserService/GetInfo"
	UserService_SendEmailVerifyCode_FullMethodName  = "/user.v1.UserService/SendEmailVerifyCode"
	UserService_CheckEmailVerifyCode_FullMethodName = "/user.v1.UserService/CheckEmailVerifyCode"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error)
	SendEmailVerifyCode(ctx context.Context, in *SendEmailVerifyCodeRequest, opts ...grpc.CallOption) (*SendEmailVerifyCodeResponse, error)
	CheckEmailVerifyCode(ctx context.Context, in *CheckEmailVerifyCodeRequest, opts ...grpc.CallOption) (*CheckEmailVerifyCodeResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error) {
	out := new(GetInfoResponse)
	err := c.cc.Invoke(ctx, UserService_GetInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SendEmailVerifyCode(ctx context.Context, in *SendEmailVerifyCodeRequest, opts ...grpc.CallOption) (*SendEmailVerifyCodeResponse, error) {
	out := new(SendEmailVerifyCodeResponse)
	err := c.cc.Invoke(ctx, UserService_SendEmailVerifyCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CheckEmailVerifyCode(ctx context.Context, in *CheckEmailVerifyCodeRequest, opts ...grpc.CallOption) (*CheckEmailVerifyCodeResponse, error) {
	out := new(CheckEmailVerifyCodeResponse)
	err := c.cc.Invoke(ctx, UserService_CheckEmailVerifyCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error)
	SendEmailVerifyCode(context.Context, *SendEmailVerifyCodeRequest) (*SendEmailVerifyCodeResponse, error)
	CheckEmailVerifyCode(context.Context, *CheckEmailVerifyCodeRequest) (*CheckEmailVerifyCodeResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (UnimplementedUserServiceServer) SendEmailVerifyCode(context.Context, *SendEmailVerifyCodeRequest) (*SendEmailVerifyCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailVerifyCode not implemented")
}
func (UnimplementedUserServiceServer) CheckEmailVerifyCode(context.Context, *CheckEmailVerifyCodeRequest) (*CheckEmailVerifyCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckEmailVerifyCode not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SendEmailVerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailVerifyCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SendEmailVerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_SendEmailVerifyCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SendEmailVerifyCode(ctx, req.(*SendEmailVerifyCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CheckEmailVerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckEmailVerifyCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CheckEmailVerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CheckEmailVerifyCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CheckEmailVerifyCode(ctx, req.(*CheckEmailVerifyCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _UserService_GetInfo_Handler,
		},
		{
			MethodName: "SendEmailVerifyCode",
			Handler:    _UserService_SendEmailVerifyCode_Handler,
		},
		{
			MethodName: "CheckEmailVerifyCode",
			Handler:    _UserService_CheckEmailVerifyCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/v1/user_service.proto",
}
