// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// OtsUIClient is the client API for OtsUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OtsUIClient interface {
	// ListAgentSpecs returns a list of Mobile Object's Agent(s) that can be started through the UI.
	ListAgentSpecs(ctx context.Context, in *ListAgentSpecsRequest, opts ...grpc.CallOption) (OtsUI_ListAgentSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type otsUIClient struct {
	cc grpc.ClientConnInterface
}

func NewOtsUIClient(cc grpc.ClientConnInterface) OtsUIClient {
	return &otsUIClient{cc}
}

func (c *otsUIClient) ListAgentSpecs(ctx context.Context, in *ListAgentSpecsRequest, opts ...grpc.CallOption) (OtsUI_ListAgentSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &OtsUI_ServiceDesc.Streams[0], "/v1.OtsUI/ListAgentSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &otsUIListAgentSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OtsUI_ListAgentSpecsClient interface {
	Recv() (*ListAgentSpecsResponse, error)
	grpc.ClientStream
}

type otsUIListAgentSpecsClient struct {
	grpc.ClientStream
}

func (x *otsUIListAgentSpecsClient) Recv() (*ListAgentSpecsResponse, error) {
	m := new(ListAgentSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *otsUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.OtsUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OtsUIServer is the server API for OtsUI service.
// All implementations must embed UnimplementedOtsUIServer
// for forward compatibility
type OtsUIServer interface {
	// ListAgentSpecs returns a list of Mobile Object's Agent(s) that can be started through the UI.
	ListAgentSpecs(*ListAgentSpecsRequest, OtsUI_ListAgentSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedOtsUIServer()
}

// UnimplementedOtsUIServer must be embedded to have forward compatible implementations.
type UnimplementedOtsUIServer struct {
}

func (UnimplementedOtsUIServer) ListAgentSpecs(*ListAgentSpecsRequest, OtsUI_ListAgentSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAgentSpecs not implemented")
}
func (UnimplementedOtsUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedOtsUIServer) mustEmbedUnimplementedOtsUIServer() {}

// UnsafeOtsUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OtsUIServer will
// result in compilation errors.
type UnsafeOtsUIServer interface {
	mustEmbedUnimplementedOtsUIServer()
}

func RegisterOtsUIServer(s grpc.ServiceRegistrar, srv OtsUIServer) {
	s.RegisterService(&OtsUI_ServiceDesc, srv)
}

func _OtsUI_ListAgentSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListAgentSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OtsUIServer).ListAgentSpecs(m, &otsUIListAgentSpecsServer{stream})
}

type OtsUI_ListAgentSpecsServer interface {
	Send(*ListAgentSpecsResponse) error
	grpc.ServerStream
}

type otsUIListAgentSpecsServer struct {
	grpc.ServerStream
}

func (x *otsUIListAgentSpecsServer) Send(m *ListAgentSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _OtsUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OtsUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.OtsUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OtsUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OtsUI_ServiceDesc is the grpc.ServiceDesc for OtsUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OtsUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.OtsUI",
	HandlerType: (*OtsUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _OtsUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListAgentSpecs",
			Handler:       _OtsUI_ListAgentSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ots-ui.proto",
}
