// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: topic.proto

package topic

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	TopicService_CreateTopic_FullMethodName  = "/topic.TopicService/CreateTopic"
	TopicService_UpdateTopic_FullMethodName  = "/topic.TopicService/UpdateTopic"
	TopicService_DeleteTopic_FullMethodName  = "/topic.TopicService/DeleteTopic"
	TopicService_GetAllTopics_FullMethodName = "/topic.TopicService/GetAllTopics"
)

// TopicServiceClient is the client API for TopicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TopicServiceClient interface {
	CreateTopic(ctx context.Context, in *CreateTopicReq, opts ...grpc.CallOption) (*CreateTopicResp, error)
	UpdateTopic(ctx context.Context, in *UpdateTopicReq, opts ...grpc.CallOption) (*UpdateTopicResp, error)
	DeleteTopic(ctx context.Context, in *DeleteTopicReq, opts ...grpc.CallOption) (*DeleteTopicResp, error)
	GetAllTopics(ctx context.Context, in *GetAllTopicsReq, opts ...grpc.CallOption) (*GetAllTopicsResp, error)
}

type topicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTopicServiceClient(cc grpc.ClientConnInterface) TopicServiceClient {
	return &topicServiceClient{cc}
}

func (c *topicServiceClient) CreateTopic(ctx context.Context, in *CreateTopicReq, opts ...grpc.CallOption) (*CreateTopicResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTopicResp)
	err := c.cc.Invoke(ctx, TopicService_CreateTopic_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) UpdateTopic(ctx context.Context, in *UpdateTopicReq, opts ...grpc.CallOption) (*UpdateTopicResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTopicResp)
	err := c.cc.Invoke(ctx, TopicService_UpdateTopic_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) DeleteTopic(ctx context.Context, in *DeleteTopicReq, opts ...grpc.CallOption) (*DeleteTopicResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTopicResp)
	err := c.cc.Invoke(ctx, TopicService_DeleteTopic_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) GetAllTopics(ctx context.Context, in *GetAllTopicsReq, opts ...grpc.CallOption) (*GetAllTopicsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllTopicsResp)
	err := c.cc.Invoke(ctx, TopicService_GetAllTopics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopicServiceServer is the server API for TopicService service.
// All implementations must embed UnimplementedTopicServiceServer
// for forward compatibility
type TopicServiceServer interface {
	CreateTopic(context.Context, *CreateTopicReq) (*CreateTopicResp, error)
	UpdateTopic(context.Context, *UpdateTopicReq) (*UpdateTopicResp, error)
	DeleteTopic(context.Context, *DeleteTopicReq) (*DeleteTopicResp, error)
	GetAllTopics(context.Context, *GetAllTopicsReq) (*GetAllTopicsResp, error)
	mustEmbedUnimplementedTopicServiceServer()
}

// UnimplementedTopicServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTopicServiceServer struct {
}

func (UnimplementedTopicServiceServer) CreateTopic(context.Context, *CreateTopicReq) (*CreateTopicResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopic not implemented")
}
func (UnimplementedTopicServiceServer) UpdateTopic(context.Context, *UpdateTopicReq) (*UpdateTopicResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTopic not implemented")
}
func (UnimplementedTopicServiceServer) DeleteTopic(context.Context, *DeleteTopicReq) (*DeleteTopicResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTopic not implemented")
}
func (UnimplementedTopicServiceServer) GetAllTopics(context.Context, *GetAllTopicsReq) (*GetAllTopicsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTopics not implemented")
}
func (UnimplementedTopicServiceServer) mustEmbedUnimplementedTopicServiceServer() {}

// UnsafeTopicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TopicServiceServer will
// result in compilation errors.
type UnsafeTopicServiceServer interface {
	mustEmbedUnimplementedTopicServiceServer()
}

func RegisterTopicServiceServer(s grpc.ServiceRegistrar, srv TopicServiceServer) {
	s.RegisterService(&TopicService_ServiceDesc, srv)
}

func _TopicService_CreateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTopicReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).CreateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_CreateTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).CreateTopic(ctx, req.(*CreateTopicReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_UpdateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTopicReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).UpdateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_UpdateTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).UpdateTopic(ctx, req.(*UpdateTopicReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_DeleteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTopicReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).DeleteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_DeleteTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).DeleteTopic(ctx, req.(*DeleteTopicReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_GetAllTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllTopicsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).GetAllTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_GetAllTopics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).GetAllTopics(ctx, req.(*GetAllTopicsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TopicService_ServiceDesc is the grpc.ServiceDesc for TopicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TopicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "topic.TopicService",
	HandlerType: (*TopicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTopic",
			Handler:    _TopicService_CreateTopic_Handler,
		},
		{
			MethodName: "UpdateTopic",
			Handler:    _TopicService_UpdateTopic_Handler,
		},
		{
			MethodName: "DeleteTopic",
			Handler:    _TopicService_DeleteTopic_Handler,
		},
		{
			MethodName: "GetAllTopics",
			Handler:    _TopicService_GetAllTopics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "topic.proto",
}
