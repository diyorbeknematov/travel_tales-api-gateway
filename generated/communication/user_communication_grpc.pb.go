// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: user_communication.proto

package communication

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

// CommunicationServiceClient is the client API for CommunicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommunicationServiceClient interface {
	SendMessageUser(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	ListMessage(ctx context.Context, in *ListMessageRequest, opts ...grpc.CallOption) (*ListMessageResponse, error)
	AddTravelTips(ctx context.Context, in *AddTravelTipsRequest, opts ...grpc.CallOption) (*AddTravelTipsResponse, error)
	GetTravelTips(ctx context.Context, in *GetTravelTipsRequest, opts ...grpc.CallOption) (*GetTravelTipsResponse, error)
	GetUserStatics(ctx context.Context, in *GetUserStaticsRequest, opts ...grpc.CallOption) (*GetUserStaticsResponse, error)
}

type communicationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommunicationServiceClient(cc grpc.ClientConnInterface) CommunicationServiceClient {
	return &communicationServiceClient{cc}
}

func (c *communicationServiceClient) SendMessageUser(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, "/user_communication.CommunicationService/SendMessageUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationServiceClient) ListMessage(ctx context.Context, in *ListMessageRequest, opts ...grpc.CallOption) (*ListMessageResponse, error) {
	out := new(ListMessageResponse)
	err := c.cc.Invoke(ctx, "/user_communication.CommunicationService/ListMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationServiceClient) AddTravelTips(ctx context.Context, in *AddTravelTipsRequest, opts ...grpc.CallOption) (*AddTravelTipsResponse, error) {
	out := new(AddTravelTipsResponse)
	err := c.cc.Invoke(ctx, "/user_communication.CommunicationService/AddTravelTips", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationServiceClient) GetTravelTips(ctx context.Context, in *GetTravelTipsRequest, opts ...grpc.CallOption) (*GetTravelTipsResponse, error) {
	out := new(GetTravelTipsResponse)
	err := c.cc.Invoke(ctx, "/user_communication.CommunicationService/GetTravelTips", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationServiceClient) GetUserStatics(ctx context.Context, in *GetUserStaticsRequest, opts ...grpc.CallOption) (*GetUserStaticsResponse, error) {
	out := new(GetUserStaticsResponse)
	err := c.cc.Invoke(ctx, "/user_communication.CommunicationService/GetUserStatics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunicationServiceServer is the server API for CommunicationService service.
// All implementations must embed UnimplementedCommunicationServiceServer
// for forward compatibility
type CommunicationServiceServer interface {
	SendMessageUser(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	ListMessage(context.Context, *ListMessageRequest) (*ListMessageResponse, error)
	AddTravelTips(context.Context, *AddTravelTipsRequest) (*AddTravelTipsResponse, error)
	GetTravelTips(context.Context, *GetTravelTipsRequest) (*GetTravelTipsResponse, error)
	GetUserStatics(context.Context, *GetUserStaticsRequest) (*GetUserStaticsResponse, error)
	mustEmbedUnimplementedCommunicationServiceServer()
}

// UnimplementedCommunicationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommunicationServiceServer struct {
}

func (UnimplementedCommunicationServiceServer) SendMessageUser(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageUser not implemented")
}
func (UnimplementedCommunicationServiceServer) ListMessage(context.Context, *ListMessageRequest) (*ListMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMessage not implemented")
}
func (UnimplementedCommunicationServiceServer) AddTravelTips(context.Context, *AddTravelTipsRequest) (*AddTravelTipsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTravelTips not implemented")
}
func (UnimplementedCommunicationServiceServer) GetTravelTips(context.Context, *GetTravelTipsRequest) (*GetTravelTipsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTravelTips not implemented")
}
func (UnimplementedCommunicationServiceServer) GetUserStatics(context.Context, *GetUserStaticsRequest) (*GetUserStaticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserStatics not implemented")
}
func (UnimplementedCommunicationServiceServer) mustEmbedUnimplementedCommunicationServiceServer() {}

// UnsafeCommunicationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommunicationServiceServer will
// result in compilation errors.
type UnsafeCommunicationServiceServer interface {
	mustEmbedUnimplementedCommunicationServiceServer()
}

func RegisterCommunicationServiceServer(s grpc.ServiceRegistrar, srv CommunicationServiceServer) {
	s.RegisterService(&CommunicationService_ServiceDesc, srv)
}

func _CommunicationService_SendMessageUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServiceServer).SendMessageUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_communication.CommunicationService/SendMessageUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServiceServer).SendMessageUser(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationService_ListMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServiceServer).ListMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_communication.CommunicationService/ListMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServiceServer).ListMessage(ctx, req.(*ListMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationService_AddTravelTips_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTravelTipsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServiceServer).AddTravelTips(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_communication.CommunicationService/AddTravelTips",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServiceServer).AddTravelTips(ctx, req.(*AddTravelTipsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationService_GetTravelTips_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTravelTipsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServiceServer).GetTravelTips(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_communication.CommunicationService/GetTravelTips",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServiceServer).GetTravelTips(ctx, req.(*GetTravelTipsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationService_GetUserStatics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserStaticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServiceServer).GetUserStatics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_communication.CommunicationService/GetUserStatics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServiceServer).GetUserStatics(ctx, req.(*GetUserStaticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommunicationService_ServiceDesc is the grpc.ServiceDesc for CommunicationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommunicationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_communication.CommunicationService",
	HandlerType: (*CommunicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessageUser",
			Handler:    _CommunicationService_SendMessageUser_Handler,
		},
		{
			MethodName: "ListMessage",
			Handler:    _CommunicationService_ListMessage_Handler,
		},
		{
			MethodName: "AddTravelTips",
			Handler:    _CommunicationService_AddTravelTips_Handler,
		},
		{
			MethodName: "GetTravelTips",
			Handler:    _CommunicationService_GetTravelTips_Handler,
		},
		{
			MethodName: "GetUserStatics",
			Handler:    _CommunicationService_GetUserStatics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_communication.proto",
}
