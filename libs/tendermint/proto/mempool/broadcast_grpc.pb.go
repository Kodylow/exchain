// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: libs/tendermint/proto/mempool/broadcast.proto

package mempool

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MempoolTxReceiverClient is the client API for MempoolTxReceiver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MempoolTxReceiverClient interface {
	Receive(ctx context.Context, in *TxsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ReceiveSentry(ctx context.Context, in *SentryTxs, opts ...grpc.CallOption) (*emptypb.Empty, error)
	TxIndices(ctx context.Context, in *IndicesRequest, opts ...grpc.CallOption) (*IndicesResponse, error)
}

type mempoolTxReceiverClient struct {
	cc grpc.ClientConnInterface
}

func NewMempoolTxReceiverClient(cc grpc.ClientConnInterface) MempoolTxReceiverClient {
	return &mempoolTxReceiverClient{cc}
}

func (c *mempoolTxReceiverClient) Receive(ctx context.Context, in *TxsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mempool.MempoolTxReceiver/Receive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mempoolTxReceiverClient) ReceiveSentry(ctx context.Context, in *SentryTxs, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mempool.MempoolTxReceiver/ReceiveSentry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mempoolTxReceiverClient) TxIndices(ctx context.Context, in *IndicesRequest, opts ...grpc.CallOption) (*IndicesResponse, error) {
	out := new(IndicesResponse)
	err := c.cc.Invoke(ctx, "/mempool.MempoolTxReceiver/TxIndices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MempoolTxReceiverServer is the server API for MempoolTxReceiver service.
// All implementations must embed UnimplementedMempoolTxReceiverServer
// for forward compatibility
type MempoolTxReceiverServer interface {
	Receive(context.Context, *TxsRequest) (*emptypb.Empty, error)
	ReceiveSentry(context.Context, *SentryTxs) (*emptypb.Empty, error)
	TxIndices(context.Context, *IndicesRequest) (*IndicesResponse, error)
	mustEmbedUnimplementedMempoolTxReceiverServer()
}

// UnimplementedMempoolTxReceiverServer must be embedded to have forward compatible implementations.
type UnimplementedMempoolTxReceiverServer struct {
}

func (UnimplementedMempoolTxReceiverServer) Receive(context.Context, *TxsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (UnimplementedMempoolTxReceiverServer) ReceiveSentry(context.Context, *SentryTxs) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveSentry not implemented")
}
func (UnimplementedMempoolTxReceiverServer) TxIndices(context.Context, *IndicesRequest) (*IndicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TxIndices not implemented")
}
func (UnimplementedMempoolTxReceiverServer) mustEmbedUnimplementedMempoolTxReceiverServer() {}

// UnsafeMempoolTxReceiverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MempoolTxReceiverServer will
// result in compilation errors.
type UnsafeMempoolTxReceiverServer interface {
	mustEmbedUnimplementedMempoolTxReceiverServer()
}

func RegisterMempoolTxReceiverServer(s grpc.ServiceRegistrar, srv MempoolTxReceiverServer) {
	s.RegisterService(&MempoolTxReceiver_ServiceDesc, srv)
}

func _MempoolTxReceiver_Receive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MempoolTxReceiverServer).Receive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mempool.MempoolTxReceiver/Receive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MempoolTxReceiverServer).Receive(ctx, req.(*TxsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MempoolTxReceiver_ReceiveSentry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SentryTxs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MempoolTxReceiverServer).ReceiveSentry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mempool.MempoolTxReceiver/ReceiveSentry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MempoolTxReceiverServer).ReceiveSentry(ctx, req.(*SentryTxs))
	}
	return interceptor(ctx, in, info, handler)
}

func _MempoolTxReceiver_TxIndices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MempoolTxReceiverServer).TxIndices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mempool.MempoolTxReceiver/TxIndices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MempoolTxReceiverServer).TxIndices(ctx, req.(*IndicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MempoolTxReceiver_ServiceDesc is the grpc.ServiceDesc for MempoolTxReceiver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MempoolTxReceiver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mempool.MempoolTxReceiver",
	HandlerType: (*MempoolTxReceiverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Receive",
			Handler:    _MempoolTxReceiver_Receive_Handler,
		},
		{
			MethodName: "ReceiveSentry",
			Handler:    _MempoolTxReceiver_ReceiveSentry_Handler,
		},
		{
			MethodName: "TxIndices",
			Handler:    _MempoolTxReceiver_TxIndices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "libs/tendermint/proto/mempool/broadcast.proto",
}