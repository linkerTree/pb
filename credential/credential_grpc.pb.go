// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package credential

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

// UserCredentialValidatorClient is the client API for UserCredentialValidator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserCredentialValidatorClient interface {
	// 获取一个public key
	GetPublicKey(ctx context.Context, in *GetPublicKeyReq, opts ...grpc.CallOption) (*GetPublicKeyRsp, error)
	// 尝试登陆的时候来验证密码是否正确
	ValidatePassWord(ctx context.Context, in *ValidatePassWordReq, opts ...grpc.CallOption) (*ValidatePassWordRsp, error)
	// 根据session ID检查登陆状态
	CheckIsLoggingIn(ctx context.Context, in *CheckIsLoggingInReq, opts ...grpc.CallOption) (*CheckIsLoggingInRsp, error)
}

type userCredentialValidatorClient struct {
	cc grpc.ClientConnInterface
}

func NewUserCredentialValidatorClient(cc grpc.ClientConnInterface) UserCredentialValidatorClient {
	return &userCredentialValidatorClient{cc}
}

func (c *userCredentialValidatorClient) GetPublicKey(ctx context.Context, in *GetPublicKeyReq, opts ...grpc.CallOption) (*GetPublicKeyRsp, error) {
	out := new(GetPublicKeyRsp)
	err := c.cc.Invoke(ctx, "/credential.UserCredentialValidator/GetPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCredentialValidatorClient) ValidatePassWord(ctx context.Context, in *ValidatePassWordReq, opts ...grpc.CallOption) (*ValidatePassWordRsp, error) {
	out := new(ValidatePassWordRsp)
	err := c.cc.Invoke(ctx, "/credential.UserCredentialValidator/ValidatePassWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCredentialValidatorClient) CheckIsLoggingIn(ctx context.Context, in *CheckIsLoggingInReq, opts ...grpc.CallOption) (*CheckIsLoggingInRsp, error) {
	out := new(CheckIsLoggingInRsp)
	err := c.cc.Invoke(ctx, "/credential.UserCredentialValidator/CheckIsLoggingIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserCredentialValidatorServer is the server API for UserCredentialValidator service.
// All implementations must embed UnimplementedUserCredentialValidatorServer
// for forward compatibility
type UserCredentialValidatorServer interface {
	// 获取一个public key
	GetPublicKey(context.Context, *GetPublicKeyReq) (*GetPublicKeyRsp, error)
	// 尝试登陆的时候来验证密码是否正确
	ValidatePassWord(context.Context, *ValidatePassWordReq) (*ValidatePassWordRsp, error)
	// 根据session ID检查登陆状态
	CheckIsLoggingIn(context.Context, *CheckIsLoggingInReq) (*CheckIsLoggingInRsp, error)
	mustEmbedUnimplementedUserCredentialValidatorServer()
}

// UnimplementedUserCredentialValidatorServer must be embedded to have forward compatible implementations.
type UnimplementedUserCredentialValidatorServer struct {
}

func (UnimplementedUserCredentialValidatorServer) GetPublicKey(context.Context, *GetPublicKeyReq) (*GetPublicKeyRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicKey not implemented")
}
func (UnimplementedUserCredentialValidatorServer) ValidatePassWord(context.Context, *ValidatePassWordReq) (*ValidatePassWordRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatePassWord not implemented")
}
func (UnimplementedUserCredentialValidatorServer) CheckIsLoggingIn(context.Context, *CheckIsLoggingInReq) (*CheckIsLoggingInRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIsLoggingIn not implemented")
}
func (UnimplementedUserCredentialValidatorServer) mustEmbedUnimplementedUserCredentialValidatorServer() {
}

// UnsafeUserCredentialValidatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserCredentialValidatorServer will
// result in compilation errors.
type UnsafeUserCredentialValidatorServer interface {
	mustEmbedUnimplementedUserCredentialValidatorServer()
}

func RegisterUserCredentialValidatorServer(s grpc.ServiceRegistrar, srv UserCredentialValidatorServer) {
	s.RegisterService(&UserCredentialValidator_ServiceDesc, srv)
}

func _UserCredentialValidator_GetPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicKeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCredentialValidatorServer).GetPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.UserCredentialValidator/GetPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCredentialValidatorServer).GetPublicKey(ctx, req.(*GetPublicKeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCredentialValidator_ValidatePassWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidatePassWordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCredentialValidatorServer).ValidatePassWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.UserCredentialValidator/ValidatePassWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCredentialValidatorServer).ValidatePassWord(ctx, req.(*ValidatePassWordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCredentialValidator_CheckIsLoggingIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckIsLoggingInReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCredentialValidatorServer).CheckIsLoggingIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.UserCredentialValidator/CheckIsLoggingIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCredentialValidatorServer).CheckIsLoggingIn(ctx, req.(*CheckIsLoggingInReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserCredentialValidator_ServiceDesc is the grpc.ServiceDesc for UserCredentialValidator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserCredentialValidator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "credential.UserCredentialValidator",
	HandlerType: (*UserCredentialValidatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPublicKey",
			Handler:    _UserCredentialValidator_GetPublicKey_Handler,
		},
		{
			MethodName: "ValidatePassWord",
			Handler:    _UserCredentialValidator_ValidatePassWord_Handler,
		},
		{
			MethodName: "CheckIsLoggingIn",
			Handler:    _UserCredentialValidator_CheckIsLoggingIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "credential.proto",
}
