// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: shopagolic/currency/v1beta1/currency_service.proto

package v1beta1

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

// CurrencyServiceClient is the client API for CurrencyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CurrencyServiceClient interface {
	GetSupportedCurrencies(ctx context.Context, in *GetSupportedCurrenciesRequest, opts ...grpc.CallOption) (*GetSupportedCurrenciesResponse, error)
	Convert(ctx context.Context, in *ConvertRequest, opts ...grpc.CallOption) (*ConvertResponse, error)
	PutCurrency(ctx context.Context, in *PutCurrencyRequest, opts ...grpc.CallOption) (*PutCurrencyResponse, error)
	DeleteCurrency(ctx context.Context, in *DeleteCurrencyRequest, opts ...grpc.CallOption) (*DeleteCurrencyResponse, error)
}

type currencyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyServiceClient(cc grpc.ClientConnInterface) CurrencyServiceClient {
	return &currencyServiceClient{cc}
}

func (c *currencyServiceClient) GetSupportedCurrencies(ctx context.Context, in *GetSupportedCurrenciesRequest, opts ...grpc.CallOption) (*GetSupportedCurrenciesResponse, error) {
	out := new(GetSupportedCurrenciesResponse)
	err := c.cc.Invoke(ctx, "/shopagolic.currency.v1beta1.CurrencyService/GetSupportedCurrencies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyServiceClient) Convert(ctx context.Context, in *ConvertRequest, opts ...grpc.CallOption) (*ConvertResponse, error) {
	out := new(ConvertResponse)
	err := c.cc.Invoke(ctx, "/shopagolic.currency.v1beta1.CurrencyService/Convert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyServiceClient) PutCurrency(ctx context.Context, in *PutCurrencyRequest, opts ...grpc.CallOption) (*PutCurrencyResponse, error) {
	out := new(PutCurrencyResponse)
	err := c.cc.Invoke(ctx, "/shopagolic.currency.v1beta1.CurrencyService/PutCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyServiceClient) DeleteCurrency(ctx context.Context, in *DeleteCurrencyRequest, opts ...grpc.CallOption) (*DeleteCurrencyResponse, error) {
	out := new(DeleteCurrencyResponse)
	err := c.cc.Invoke(ctx, "/shopagolic.currency.v1beta1.CurrencyService/DeleteCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrencyServiceServer is the server API for CurrencyService service.
// All implementations must embed UnimplementedCurrencyServiceServer
// for forward compatibility
type CurrencyServiceServer interface {
	GetSupportedCurrencies(context.Context, *GetSupportedCurrenciesRequest) (*GetSupportedCurrenciesResponse, error)
	Convert(context.Context, *ConvertRequest) (*ConvertResponse, error)
	PutCurrency(context.Context, *PutCurrencyRequest) (*PutCurrencyResponse, error)
	DeleteCurrency(context.Context, *DeleteCurrencyRequest) (*DeleteCurrencyResponse, error)
	mustEmbedUnimplementedCurrencyServiceServer()
}

// UnimplementedCurrencyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCurrencyServiceServer struct {
}

func (UnimplementedCurrencyServiceServer) GetSupportedCurrencies(context.Context, *GetSupportedCurrenciesRequest) (*GetSupportedCurrenciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupportedCurrencies not implemented")
}
func (UnimplementedCurrencyServiceServer) Convert(context.Context, *ConvertRequest) (*ConvertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Convert not implemented")
}
func (UnimplementedCurrencyServiceServer) PutCurrency(context.Context, *PutCurrencyRequest) (*PutCurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutCurrency not implemented")
}
func (UnimplementedCurrencyServiceServer) DeleteCurrency(context.Context, *DeleteCurrencyRequest) (*DeleteCurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCurrency not implemented")
}
func (UnimplementedCurrencyServiceServer) mustEmbedUnimplementedCurrencyServiceServer() {}

// UnsafeCurrencyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CurrencyServiceServer will
// result in compilation errors.
type UnsafeCurrencyServiceServer interface {
	mustEmbedUnimplementedCurrencyServiceServer()
}

func RegisterCurrencyServiceServer(s grpc.ServiceRegistrar, srv CurrencyServiceServer) {
	s.RegisterService(&CurrencyService_ServiceDesc, srv)
}

func _CurrencyService_GetSupportedCurrencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSupportedCurrenciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).GetSupportedCurrencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shopagolic.currency.v1beta1.CurrencyService/GetSupportedCurrencies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).GetSupportedCurrencies(ctx, req.(*GetSupportedCurrenciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyService_Convert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConvertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).Convert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shopagolic.currency.v1beta1.CurrencyService/Convert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).Convert(ctx, req.(*ConvertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyService_PutCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).PutCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shopagolic.currency.v1beta1.CurrencyService/PutCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).PutCurrency(ctx, req.(*PutCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyService_DeleteCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).DeleteCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shopagolic.currency.v1beta1.CurrencyService/DeleteCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).DeleteCurrency(ctx, req.(*DeleteCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CurrencyService_ServiceDesc is the grpc.ServiceDesc for CurrencyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CurrencyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shopagolic.currency.v1beta1.CurrencyService",
	HandlerType: (*CurrencyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSupportedCurrencies",
			Handler:    _CurrencyService_GetSupportedCurrencies_Handler,
		},
		{
			MethodName: "Convert",
			Handler:    _CurrencyService_Convert_Handler,
		},
		{
			MethodName: "PutCurrency",
			Handler:    _CurrencyService_PutCurrency_Handler,
		},
		{
			MethodName: "DeleteCurrency",
			Handler:    _CurrencyService_DeleteCurrency_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shopagolic/currency/v1beta1/currency_service.proto",
}
