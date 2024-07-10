// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.2
// source: proto/echo.proto

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

// PersonClient is the client API for Person service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PersonClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetPerson, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreatePerson, error)
	ChangeName(ctx context.Context, in *ChangeNameRequest, opts ...grpc.CallOption) (*ChangePersonName, error)
	ChangeAmount(ctx context.Context, in *ChangeAmountRequest, opts ...grpc.CallOption) (*ChangePersonAmount, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeletePerson, error)
}

type personClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonClient(cc grpc.ClientConnInterface) PersonClient {
	return &personClient{cc}
}

func (c *personClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetPerson, error) {
	out := new(GetPerson)
	err := c.cc.Invoke(ctx, "/proto.Person/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreatePerson, error) {
	out := new(CreatePerson)
	err := c.cc.Invoke(ctx, "/proto.Person/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personClient) ChangeName(ctx context.Context, in *ChangeNameRequest, opts ...grpc.CallOption) (*ChangePersonName, error) {
	out := new(ChangePersonName)
	err := c.cc.Invoke(ctx, "/proto.Person/ChangeName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personClient) ChangeAmount(ctx context.Context, in *ChangeAmountRequest, opts ...grpc.CallOption) (*ChangePersonAmount, error) {
	out := new(ChangePersonAmount)
	err := c.cc.Invoke(ctx, "/proto.Person/ChangeAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeletePerson, error) {
	out := new(DeletePerson)
	err := c.cc.Invoke(ctx, "/proto.Person/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersonServer is the server API for Person service.
// All implementations must embed UnimplementedPersonServer
// for forward compatibility
type PersonServer interface {
	Get(context.Context, *GetRequest) (*GetPerson, error)
	Create(context.Context, *CreateRequest) (*CreatePerson, error)
	ChangeName(context.Context, *ChangeNameRequest) (*ChangePersonName, error)
	ChangeAmount(context.Context, *ChangeAmountRequest) (*ChangePersonAmount, error)
	Delete(context.Context, *DeleteRequest) (*DeletePerson, error)
	mustEmbedUnimplementedPersonServer()
}

// UnimplementedPersonServer must be embedded to have forward compatible implementations.
type UnimplementedPersonServer struct {
}

func (UnimplementedPersonServer) Get(context.Context, *GetRequest) (*GetPerson, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPersonServer) Create(context.Context, *CreateRequest) (*CreatePerson, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPersonServer) ChangeName(context.Context, *ChangeNameRequest) (*ChangePersonName, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeName not implemented")
}
func (UnimplementedPersonServer) ChangeAmount(context.Context, *ChangeAmountRequest) (*ChangePersonAmount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeAmount not implemented")
}
func (UnimplementedPersonServer) Delete(context.Context, *DeleteRequest) (*DeletePerson, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPersonServer) mustEmbedUnimplementedPersonServer() {}

// UnsafePersonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PersonServer will
// result in compilation errors.
type UnsafePersonServer interface {
	mustEmbedUnimplementedPersonServer()
}

func RegisterPersonServer(s grpc.ServiceRegistrar, srv PersonServer) {
	s.RegisterService(&Person_ServiceDesc, srv)
}

func _Person_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Person/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Person_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Person/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Person_ChangeName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServer).ChangeName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Person/ChangeName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServer).ChangeName(ctx, req.(*ChangeNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Person_ChangeAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeAmountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServer).ChangeAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Person/ChangeAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServer).ChangeAmount(ctx, req.(*ChangeAmountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Person_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Person/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Person_ServiceDesc is the grpc.ServiceDesc for Person service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Person_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Person",
	HandlerType: (*PersonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Person_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Person_Create_Handler,
		},
		{
			MethodName: "ChangeName",
			Handler:    _Person_ChangeName_Handler,
		},
		{
			MethodName: "ChangeAmount",
			Handler:    _Person_ChangeAmount_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Person_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/echo.proto",
}