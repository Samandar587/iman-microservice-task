// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: internal/datafetcher/ports/grpc/proto/posts.proto

package pb

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

// SavePostsServiceClient is the client API for SavePostsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SavePostsServiceClient interface {
	CollectPosts(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type savePostsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSavePostsServiceClient(cc grpc.ClientConnInterface) SavePostsServiceClient {
	return &savePostsServiceClient{cc}
}

func (c *savePostsServiceClient) CollectPosts(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.SavePostsService/CollectPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SavePostsServiceServer is the server API for SavePostsService service.
// All implementations must embed UnimplementedSavePostsServiceServer
// for forward compatibility
type SavePostsServiceServer interface {
	CollectPosts(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedSavePostsServiceServer()
}

// UnimplementedSavePostsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSavePostsServiceServer struct {
}

func (UnimplementedSavePostsServiceServer) CollectPosts(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectPosts not implemented")
}
func (UnimplementedSavePostsServiceServer) mustEmbedUnimplementedSavePostsServiceServer() {}

// UnsafeSavePostsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SavePostsServiceServer will
// result in compilation errors.
type UnsafeSavePostsServiceServer interface {
	mustEmbedUnimplementedSavePostsServiceServer()
}

func RegisterSavePostsServiceServer(s grpc.ServiceRegistrar, srv SavePostsServiceServer) {
	s.RegisterService(&SavePostsService_ServiceDesc, srv)
}

func _SavePostsService_CollectPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SavePostsServiceServer).CollectPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.SavePostsService/CollectPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SavePostsServiceServer).CollectPosts(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// SavePostsService_ServiceDesc is the grpc.ServiceDesc for SavePostsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SavePostsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.SavePostsService",
	HandlerType: (*SavePostsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CollectPosts",
			Handler:    _SavePostsService_CollectPosts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/datafetcher/ports/grpc/proto/posts.proto",
}