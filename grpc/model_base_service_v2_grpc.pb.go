// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.2
// source: services/model_base_service_v2.proto

package grpc

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

// ModelBaseServiceV2Client is the client API for ModelBaseServiceV2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelBaseServiceV2Client interface {
	GetById(ctx context.Context, in *ModelServiceV2GetByIdRequest, opts ...grpc.CallOption) (*Response, error)
	GetOne(ctx context.Context, in *ModelServiceV2GetOneRequest, opts ...grpc.CallOption) (*Response, error)
	GetMany(ctx context.Context, in *ModelServiceV2GetManyRequest, opts ...grpc.CallOption) (*Response, error)
	DeleteById(ctx context.Context, in *ModelServiceV2DeleteByIdRequest, opts ...grpc.CallOption) (*Response, error)
	DeleteOne(ctx context.Context, in *ModelServiceV2DeleteOneRequest, opts ...grpc.CallOption) (*Response, error)
	DeleteMany(ctx context.Context, in *ModelServiceV2DeleteManyRequest, opts ...grpc.CallOption) (*Response, error)
	UpdateById(ctx context.Context, in *ModelServiceV2UpdateByIdRequest, opts ...grpc.CallOption) (*Response, error)
	UpdateOne(ctx context.Context, in *ModelServiceV2UpdateOneRequest, opts ...grpc.CallOption) (*Response, error)
	UpdateMany(ctx context.Context, in *ModelServiceV2UpdateManyRequest, opts ...grpc.CallOption) (*Response, error)
	ReplaceById(ctx context.Context, in *ModelServiceV2ReplaceByIdRequest, opts ...grpc.CallOption) (*Response, error)
	ReplaceOne(ctx context.Context, in *ModelServiceV2ReplaceOneRequest, opts ...grpc.CallOption) (*Response, error)
	InsertOne(ctx context.Context, in *ModelServiceV2InsertOneRequest, opts ...grpc.CallOption) (*Response, error)
	InsertMany(ctx context.Context, in *ModelServiceV2InsertManyRequest, opts ...grpc.CallOption) (*Response, error)
	Count(ctx context.Context, in *ModelServiceV2CountRequest, opts ...grpc.CallOption) (*Response, error)
}

type modelBaseServiceV2Client struct {
	cc grpc.ClientConnInterface
}

func NewModelBaseServiceV2Client(cc grpc.ClientConnInterface) ModelBaseServiceV2Client {
	return &modelBaseServiceV2Client{cc}
}

func (c *modelBaseServiceV2Client) GetById(ctx context.Context, in *ModelServiceV2GetByIdRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseServiceV2/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceV2Client) GetOne(ctx context.Context, in *ModelServiceV2GetOneRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseServiceV2/GetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceV2Client) GetMany(ctx context.Context, in *ModelServiceV2GetManyRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseServiceV2/GetMany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceV2Client) DeleteById(ctx context.Context, in *ModelServiceV2DeleteByIdRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseServiceV2/DeleteById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceV2Client) DeleteOne(ctx context.Context, in *ModelServiceV2DeleteOneRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseServiceV2/DeleteOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceV2Client) DeleteMany(ctx context.Context, in *ModelServiceV2DeleteManyRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseServiceV2/DeleteMany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceV2Client) UpdateById(ctx context.Context, in *ModelServiceV2UpdateByIdRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseServiceV2/UpdateById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceV2Client) UpdateOne(ctx context.Context, in *ModelServiceV2UpdateOneRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseServiceV2/UpdateOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceV2Client) UpdateMany(ctx context.Context, in *ModelServiceV2UpdateManyRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseServiceV2/UpdateMany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceV2Client) ReplaceById(ctx context.Context, in *ModelServiceV2ReplaceByIdRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseServiceV2/ReplaceById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceV2Client) ReplaceOne(ctx context.Context, in *ModelServiceV2ReplaceOneRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseServiceV2/ReplaceOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceV2Client) InsertOne(ctx context.Context, in *ModelServiceV2InsertOneRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseServiceV2/InsertOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceV2Client) InsertMany(ctx context.Context, in *ModelServiceV2InsertManyRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseServiceV2/InsertMany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelBaseServiceV2Client) Count(ctx context.Context, in *ModelServiceV2CountRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.ModelBaseServiceV2/Count", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelBaseServiceV2Server is the server API for ModelBaseServiceV2 service.
// All implementations must embed UnimplementedModelBaseServiceV2Server
// for forward compatibility
type ModelBaseServiceV2Server interface {
	GetById(context.Context, *ModelServiceV2GetByIdRequest) (*Response, error)
	GetOne(context.Context, *ModelServiceV2GetOneRequest) (*Response, error)
	GetMany(context.Context, *ModelServiceV2GetManyRequest) (*Response, error)
	DeleteById(context.Context, *ModelServiceV2DeleteByIdRequest) (*Response, error)
	DeleteOne(context.Context, *ModelServiceV2DeleteOneRequest) (*Response, error)
	DeleteMany(context.Context, *ModelServiceV2DeleteManyRequest) (*Response, error)
	UpdateById(context.Context, *ModelServiceV2UpdateByIdRequest) (*Response, error)
	UpdateOne(context.Context, *ModelServiceV2UpdateOneRequest) (*Response, error)
	UpdateMany(context.Context, *ModelServiceV2UpdateManyRequest) (*Response, error)
	ReplaceById(context.Context, *ModelServiceV2ReplaceByIdRequest) (*Response, error)
	ReplaceOne(context.Context, *ModelServiceV2ReplaceOneRequest) (*Response, error)
	InsertOne(context.Context, *ModelServiceV2InsertOneRequest) (*Response, error)
	InsertMany(context.Context, *ModelServiceV2InsertManyRequest) (*Response, error)
	Count(context.Context, *ModelServiceV2CountRequest) (*Response, error)
	mustEmbedUnimplementedModelBaseServiceV2Server()
}

// UnimplementedModelBaseServiceV2Server must be embedded to have forward compatible implementations.
type UnimplementedModelBaseServiceV2Server struct {
}

func (UnimplementedModelBaseServiceV2Server) GetById(context.Context, *ModelServiceV2GetByIdRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedModelBaseServiceV2Server) GetOne(context.Context, *ModelServiceV2GetOneRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedModelBaseServiceV2Server) GetMany(context.Context, *ModelServiceV2GetManyRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMany not implemented")
}
func (UnimplementedModelBaseServiceV2Server) DeleteById(context.Context, *ModelServiceV2DeleteByIdRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteById not implemented")
}
func (UnimplementedModelBaseServiceV2Server) DeleteOne(context.Context, *ModelServiceV2DeleteOneRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOne not implemented")
}
func (UnimplementedModelBaseServiceV2Server) DeleteMany(context.Context, *ModelServiceV2DeleteManyRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMany not implemented")
}
func (UnimplementedModelBaseServiceV2Server) UpdateById(context.Context, *ModelServiceV2UpdateByIdRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateById not implemented")
}
func (UnimplementedModelBaseServiceV2Server) UpdateOne(context.Context, *ModelServiceV2UpdateOneRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOne not implemented")
}
func (UnimplementedModelBaseServiceV2Server) UpdateMany(context.Context, *ModelServiceV2UpdateManyRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMany not implemented")
}
func (UnimplementedModelBaseServiceV2Server) ReplaceById(context.Context, *ModelServiceV2ReplaceByIdRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceById not implemented")
}
func (UnimplementedModelBaseServiceV2Server) ReplaceOne(context.Context, *ModelServiceV2ReplaceOneRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceOne not implemented")
}
func (UnimplementedModelBaseServiceV2Server) InsertOne(context.Context, *ModelServiceV2InsertOneRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertOne not implemented")
}
func (UnimplementedModelBaseServiceV2Server) InsertMany(context.Context, *ModelServiceV2InsertManyRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertMany not implemented")
}
func (UnimplementedModelBaseServiceV2Server) Count(context.Context, *ModelServiceV2CountRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (UnimplementedModelBaseServiceV2Server) mustEmbedUnimplementedModelBaseServiceV2Server() {}

// UnsafeModelBaseServiceV2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelBaseServiceV2Server will
// result in compilation errors.
type UnsafeModelBaseServiceV2Server interface {
	mustEmbedUnimplementedModelBaseServiceV2Server()
}

func RegisterModelBaseServiceV2Server(s grpc.ServiceRegistrar, srv ModelBaseServiceV2Server) {
	s.RegisterService(&ModelBaseServiceV2_ServiceDesc, srv)
}

func _ModelBaseServiceV2_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelServiceV2GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceV2Server).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseServiceV2/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceV2Server).GetById(ctx, req.(*ModelServiceV2GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseServiceV2_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelServiceV2GetOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceV2Server).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseServiceV2/GetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceV2Server).GetOne(ctx, req.(*ModelServiceV2GetOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseServiceV2_GetMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelServiceV2GetManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceV2Server).GetMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseServiceV2/GetMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceV2Server).GetMany(ctx, req.(*ModelServiceV2GetManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseServiceV2_DeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelServiceV2DeleteByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceV2Server).DeleteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseServiceV2/DeleteById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceV2Server).DeleteById(ctx, req.(*ModelServiceV2DeleteByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseServiceV2_DeleteOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelServiceV2DeleteOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceV2Server).DeleteOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseServiceV2/DeleteOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceV2Server).DeleteOne(ctx, req.(*ModelServiceV2DeleteOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseServiceV2_DeleteMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelServiceV2DeleteManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceV2Server).DeleteMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseServiceV2/DeleteMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceV2Server).DeleteMany(ctx, req.(*ModelServiceV2DeleteManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseServiceV2_UpdateById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelServiceV2UpdateByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceV2Server).UpdateById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseServiceV2/UpdateById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceV2Server).UpdateById(ctx, req.(*ModelServiceV2UpdateByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseServiceV2_UpdateOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelServiceV2UpdateOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceV2Server).UpdateOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseServiceV2/UpdateOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceV2Server).UpdateOne(ctx, req.(*ModelServiceV2UpdateOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseServiceV2_UpdateMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelServiceV2UpdateManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceV2Server).UpdateMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseServiceV2/UpdateMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceV2Server).UpdateMany(ctx, req.(*ModelServiceV2UpdateManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseServiceV2_ReplaceById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelServiceV2ReplaceByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceV2Server).ReplaceById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseServiceV2/ReplaceById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceV2Server).ReplaceById(ctx, req.(*ModelServiceV2ReplaceByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseServiceV2_ReplaceOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelServiceV2ReplaceOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceV2Server).ReplaceOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseServiceV2/ReplaceOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceV2Server).ReplaceOne(ctx, req.(*ModelServiceV2ReplaceOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseServiceV2_InsertOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelServiceV2InsertOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceV2Server).InsertOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseServiceV2/InsertOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceV2Server).InsertOne(ctx, req.(*ModelServiceV2InsertOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseServiceV2_InsertMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelServiceV2InsertManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceV2Server).InsertMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseServiceV2/InsertMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceV2Server).InsertMany(ctx, req.(*ModelServiceV2InsertManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelBaseServiceV2_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelServiceV2CountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBaseServiceV2Server).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ModelBaseServiceV2/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBaseServiceV2Server).Count(ctx, req.(*ModelServiceV2CountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ModelBaseServiceV2_ServiceDesc is the grpc.ServiceDesc for ModelBaseServiceV2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelBaseServiceV2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.ModelBaseServiceV2",
	HandlerType: (*ModelBaseServiceV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _ModelBaseServiceV2_GetById_Handler,
		},
		{
			MethodName: "GetOne",
			Handler:    _ModelBaseServiceV2_GetOne_Handler,
		},
		{
			MethodName: "GetMany",
			Handler:    _ModelBaseServiceV2_GetMany_Handler,
		},
		{
			MethodName: "DeleteById",
			Handler:    _ModelBaseServiceV2_DeleteById_Handler,
		},
		{
			MethodName: "DeleteOne",
			Handler:    _ModelBaseServiceV2_DeleteOne_Handler,
		},
		{
			MethodName: "DeleteMany",
			Handler:    _ModelBaseServiceV2_DeleteMany_Handler,
		},
		{
			MethodName: "UpdateById",
			Handler:    _ModelBaseServiceV2_UpdateById_Handler,
		},
		{
			MethodName: "UpdateOne",
			Handler:    _ModelBaseServiceV2_UpdateOne_Handler,
		},
		{
			MethodName: "UpdateMany",
			Handler:    _ModelBaseServiceV2_UpdateMany_Handler,
		},
		{
			MethodName: "ReplaceById",
			Handler:    _ModelBaseServiceV2_ReplaceById_Handler,
		},
		{
			MethodName: "ReplaceOne",
			Handler:    _ModelBaseServiceV2_ReplaceOne_Handler,
		},
		{
			MethodName: "InsertOne",
			Handler:    _ModelBaseServiceV2_InsertOne_Handler,
		},
		{
			MethodName: "InsertMany",
			Handler:    _ModelBaseServiceV2_InsertMany_Handler,
		},
		{
			MethodName: "Count",
			Handler:    _ModelBaseServiceV2_Count_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/model_base_service_v2.proto",
}
