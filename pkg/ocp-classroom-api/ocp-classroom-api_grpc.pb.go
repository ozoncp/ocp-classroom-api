// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ocp_classroom_api

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

// OcpClassroomApiClient is the client API for OcpClassroomApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OcpClassroomApiClient interface {
	// Возвращает список учебных комнат
	ListClassroomsV1(ctx context.Context, in *ListClassroomsV1Request, opts ...grpc.CallOption) (*ListClassroomsV1Response, error)
	// Возвращает описание учебной комнаты по ее идентификатору
	DescribeClassroomV1(ctx context.Context, in *DescribeClassroomV1Request, opts ...grpc.CallOption) (*DescribeClassroomV1Response, error)
	// Создает учебную комнату
	CreateClassroomV1(ctx context.Context, in *CreateClassroomV1Request, opts ...grpc.CallOption) (*CreateClassroomV1Response, error)
	// Создает множество учебных комнат
	MultiCreateClassroomV1(ctx context.Context, in *MultiCreateClassroomV1Request, opts ...grpc.CallOption) (*MultiCreateClassroomV1Response, error)
	// Обновляет данные в учебной комнате
	UpdateClassroomV1(ctx context.Context, in *UpdateClassroomV1Request, opts ...grpc.CallOption) (*UpdateClassroomV1Response, error)
	// Удаляет учебную комнату по её идентификатору
	RemoveClassroomV1(ctx context.Context, in *RemoveClassroomV1Request, opts ...grpc.CallOption) (*RemoveClassroomV1Response, error)
}

type ocpClassroomApiClient struct {
	cc grpc.ClientConnInterface
}

func NewOcpClassroomApiClient(cc grpc.ClientConnInterface) OcpClassroomApiClient {
	return &ocpClassroomApiClient{cc}
}

func (c *ocpClassroomApiClient) ListClassroomsV1(ctx context.Context, in *ListClassroomsV1Request, opts ...grpc.CallOption) (*ListClassroomsV1Response, error) {
	out := new(ListClassroomsV1Response)
	err := c.cc.Invoke(ctx, "/ocp.classroom.api.OcpClassroomApi/ListClassroomsV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpClassroomApiClient) DescribeClassroomV1(ctx context.Context, in *DescribeClassroomV1Request, opts ...grpc.CallOption) (*DescribeClassroomV1Response, error) {
	out := new(DescribeClassroomV1Response)
	err := c.cc.Invoke(ctx, "/ocp.classroom.api.OcpClassroomApi/DescribeClassroomV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpClassroomApiClient) CreateClassroomV1(ctx context.Context, in *CreateClassroomV1Request, opts ...grpc.CallOption) (*CreateClassroomV1Response, error) {
	out := new(CreateClassroomV1Response)
	err := c.cc.Invoke(ctx, "/ocp.classroom.api.OcpClassroomApi/CreateClassroomV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpClassroomApiClient) MultiCreateClassroomV1(ctx context.Context, in *MultiCreateClassroomV1Request, opts ...grpc.CallOption) (*MultiCreateClassroomV1Response, error) {
	out := new(MultiCreateClassroomV1Response)
	err := c.cc.Invoke(ctx, "/ocp.classroom.api.OcpClassroomApi/MultiCreateClassroomV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpClassroomApiClient) UpdateClassroomV1(ctx context.Context, in *UpdateClassroomV1Request, opts ...grpc.CallOption) (*UpdateClassroomV1Response, error) {
	out := new(UpdateClassroomV1Response)
	err := c.cc.Invoke(ctx, "/ocp.classroom.api.OcpClassroomApi/UpdateClassroomV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpClassroomApiClient) RemoveClassroomV1(ctx context.Context, in *RemoveClassroomV1Request, opts ...grpc.CallOption) (*RemoveClassroomV1Response, error) {
	out := new(RemoveClassroomV1Response)
	err := c.cc.Invoke(ctx, "/ocp.classroom.api.OcpClassroomApi/RemoveClassroomV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OcpClassroomApiServer is the server API for OcpClassroomApi service.
// All implementations must embed UnimplementedOcpClassroomApiServer
// for forward compatibility
type OcpClassroomApiServer interface {
	// Возвращает список учебных комнат
	ListClassroomsV1(context.Context, *ListClassroomsV1Request) (*ListClassroomsV1Response, error)
	// Возвращает описание учебной комнаты по ее идентификатору
	DescribeClassroomV1(context.Context, *DescribeClassroomV1Request) (*DescribeClassroomV1Response, error)
	// Создает учебную комнату
	CreateClassroomV1(context.Context, *CreateClassroomV1Request) (*CreateClassroomV1Response, error)
	// Создает множество учебных комнат
	MultiCreateClassroomV1(context.Context, *MultiCreateClassroomV1Request) (*MultiCreateClassroomV1Response, error)
	// Обновляет данные в учебной комнате
	UpdateClassroomV1(context.Context, *UpdateClassroomV1Request) (*UpdateClassroomV1Response, error)
	// Удаляет учебную комнату по её идентификатору
	RemoveClassroomV1(context.Context, *RemoveClassroomV1Request) (*RemoveClassroomV1Response, error)
	mustEmbedUnimplementedOcpClassroomApiServer()
}

// UnimplementedOcpClassroomApiServer must be embedded to have forward compatible implementations.
type UnimplementedOcpClassroomApiServer struct {
}

func (UnimplementedOcpClassroomApiServer) ListClassroomsV1(context.Context, *ListClassroomsV1Request) (*ListClassroomsV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClassroomsV1 not implemented")
}
func (UnimplementedOcpClassroomApiServer) DescribeClassroomV1(context.Context, *DescribeClassroomV1Request) (*DescribeClassroomV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeClassroomV1 not implemented")
}
func (UnimplementedOcpClassroomApiServer) CreateClassroomV1(context.Context, *CreateClassroomV1Request) (*CreateClassroomV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClassroomV1 not implemented")
}
func (UnimplementedOcpClassroomApiServer) MultiCreateClassroomV1(context.Context, *MultiCreateClassroomV1Request) (*MultiCreateClassroomV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiCreateClassroomV1 not implemented")
}
func (UnimplementedOcpClassroomApiServer) UpdateClassroomV1(context.Context, *UpdateClassroomV1Request) (*UpdateClassroomV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClassroomV1 not implemented")
}
func (UnimplementedOcpClassroomApiServer) RemoveClassroomV1(context.Context, *RemoveClassroomV1Request) (*RemoveClassroomV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveClassroomV1 not implemented")
}
func (UnimplementedOcpClassroomApiServer) mustEmbedUnimplementedOcpClassroomApiServer() {}

// UnsafeOcpClassroomApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OcpClassroomApiServer will
// result in compilation errors.
type UnsafeOcpClassroomApiServer interface {
	mustEmbedUnimplementedOcpClassroomApiServer()
}

func RegisterOcpClassroomApiServer(s grpc.ServiceRegistrar, srv OcpClassroomApiServer) {
	s.RegisterService(&OcpClassroomApi_ServiceDesc, srv)
}

func _OcpClassroomApi_ListClassroomsV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClassroomsV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpClassroomApiServer).ListClassroomsV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.classroom.api.OcpClassroomApi/ListClassroomsV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpClassroomApiServer).ListClassroomsV1(ctx, req.(*ListClassroomsV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpClassroomApi_DescribeClassroomV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeClassroomV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpClassroomApiServer).DescribeClassroomV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.classroom.api.OcpClassroomApi/DescribeClassroomV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpClassroomApiServer).DescribeClassroomV1(ctx, req.(*DescribeClassroomV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpClassroomApi_CreateClassroomV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClassroomV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpClassroomApiServer).CreateClassroomV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.classroom.api.OcpClassroomApi/CreateClassroomV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpClassroomApiServer).CreateClassroomV1(ctx, req.(*CreateClassroomV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpClassroomApi_MultiCreateClassroomV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiCreateClassroomV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpClassroomApiServer).MultiCreateClassroomV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.classroom.api.OcpClassroomApi/MultiCreateClassroomV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpClassroomApiServer).MultiCreateClassroomV1(ctx, req.(*MultiCreateClassroomV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpClassroomApi_UpdateClassroomV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClassroomV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpClassroomApiServer).UpdateClassroomV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.classroom.api.OcpClassroomApi/UpdateClassroomV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpClassroomApiServer).UpdateClassroomV1(ctx, req.(*UpdateClassroomV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpClassroomApi_RemoveClassroomV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveClassroomV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpClassroomApiServer).RemoveClassroomV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.classroom.api.OcpClassroomApi/RemoveClassroomV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpClassroomApiServer).RemoveClassroomV1(ctx, req.(*RemoveClassroomV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OcpClassroomApi_ServiceDesc is the grpc.ServiceDesc for OcpClassroomApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OcpClassroomApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ocp.classroom.api.OcpClassroomApi",
	HandlerType: (*OcpClassroomApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListClassroomsV1",
			Handler:    _OcpClassroomApi_ListClassroomsV1_Handler,
		},
		{
			MethodName: "DescribeClassroomV1",
			Handler:    _OcpClassroomApi_DescribeClassroomV1_Handler,
		},
		{
			MethodName: "CreateClassroomV1",
			Handler:    _OcpClassroomApi_CreateClassroomV1_Handler,
		},
		{
			MethodName: "MultiCreateClassroomV1",
			Handler:    _OcpClassroomApi_MultiCreateClassroomV1_Handler,
		},
		{
			MethodName: "UpdateClassroomV1",
			Handler:    _OcpClassroomApi_UpdateClassroomV1_Handler,
		},
		{
			MethodName: "RemoveClassroomV1",
			Handler:    _OcpClassroomApi_RemoveClassroomV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ocp-classroom-api/ocp-classroom-api.proto",
}
