// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/product.proto

package api

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

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductClient interface {
	// create product endpoint
	ProductCreate(ctx context.Context, in *ProductCreateRequest, opts ...grpc.CallOption) (*ProductCreateResponse, error)
	// list products endpoint
	ProductList(ctx context.Context, in *ProductListRequest, opts ...grpc.CallOption) (*ProductListResponse, error)
	// update product endpoint
	ProductUpdate(ctx context.Context, in *ProductUpdateRequest, opts ...grpc.CallOption) (*ProductUpdateResponse, error)
	// delete product endpoint
	ProductDelete(ctx context.Context, in *ProductDeleteRequest, opts ...grpc.CallOption) (*ProductDeleteResponse, error)
}

type productClient struct {
	cc grpc.ClientConnInterface
}

func NewProductClient(cc grpc.ClientConnInterface) ProductClient {
	return &productClient{cc}
}

func (c *productClient) ProductCreate(ctx context.Context, in *ProductCreateRequest, opts ...grpc.CallOption) (*ProductCreateResponse, error) {
	out := new(ProductCreateResponse)
	err := c.cc.Invoke(ctx, "/homework.api.Product/ProductCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) ProductList(ctx context.Context, in *ProductListRequest, opts ...grpc.CallOption) (*ProductListResponse, error) {
	out := new(ProductListResponse)
	err := c.cc.Invoke(ctx, "/homework.api.Product/ProductList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) ProductUpdate(ctx context.Context, in *ProductUpdateRequest, opts ...grpc.CallOption) (*ProductUpdateResponse, error) {
	out := new(ProductUpdateResponse)
	err := c.cc.Invoke(ctx, "/homework.api.Product/ProductUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) ProductDelete(ctx context.Context, in *ProductDeleteRequest, opts ...grpc.CallOption) (*ProductDeleteResponse, error) {
	out := new(ProductDeleteResponse)
	err := c.cc.Invoke(ctx, "/homework.api.Product/ProductDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
// All implementations must embed UnimplementedProductServer
// for forward compatibility
type ProductServer interface {
	// create product endpoint
	ProductCreate(context.Context, *ProductCreateRequest) (*ProductCreateResponse, error)
	// list products endpoint
	ProductList(context.Context, *ProductListRequest) (*ProductListResponse, error)
	// update product endpoint
	ProductUpdate(context.Context, *ProductUpdateRequest) (*ProductUpdateResponse, error)
	// delete product endpoint
	ProductDelete(context.Context, *ProductDeleteRequest) (*ProductDeleteResponse, error)
	mustEmbedUnimplementedProductServer()
}

// UnimplementedProductServer must be embedded to have forward compatible implementations.
type UnimplementedProductServer struct {
}

func (UnimplementedProductServer) ProductCreate(context.Context, *ProductCreateRequest) (*ProductCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductCreate not implemented")
}
func (UnimplementedProductServer) ProductList(context.Context, *ProductListRequest) (*ProductListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductList not implemented")
}
func (UnimplementedProductServer) ProductUpdate(context.Context, *ProductUpdateRequest) (*ProductUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductUpdate not implemented")
}
func (UnimplementedProductServer) ProductDelete(context.Context, *ProductDeleteRequest) (*ProductDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductDelete not implemented")
}
func (UnimplementedProductServer) mustEmbedUnimplementedProductServer() {}

// UnsafeProductServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServer will
// result in compilation errors.
type UnsafeProductServer interface {
	mustEmbedUnimplementedProductServer()
}

func RegisterProductServer(s grpc.ServiceRegistrar, srv ProductServer) {
	s.RegisterService(&Product_ServiceDesc, srv)
}

func _Product_ProductCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).ProductCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/homework.api.Product/ProductCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).ProductCreate(ctx, req.(*ProductCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_ProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).ProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/homework.api.Product/ProductList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).ProductList(ctx, req.(*ProductListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_ProductUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).ProductUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/homework.api.Product/ProductUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).ProductUpdate(ctx, req.(*ProductUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_ProductDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).ProductDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/homework.api.Product/ProductDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).ProductDelete(ctx, req.(*ProductDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Product_ServiceDesc is the grpc.ServiceDesc for Product service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Product_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "homework.api.Product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProductCreate",
			Handler:    _Product_ProductCreate_Handler,
		},
		{
			MethodName: "ProductList",
			Handler:    _Product_ProductList_Handler,
		},
		{
			MethodName: "ProductUpdate",
			Handler:    _Product_ProductUpdate_Handler,
		},
		{
			MethodName: "ProductDelete",
			Handler:    _Product_ProductDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/product.proto",
}
