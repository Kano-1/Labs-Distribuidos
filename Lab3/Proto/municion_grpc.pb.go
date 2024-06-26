// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0--rc1
// source: Proto/municion.proto

package Proto

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

const (
	ServicioMunicion_SolicitarM_FullMethodName = "/Proto.servicioMunicion/SolicitarM"
)

// ServicioMunicionClient is the client API for ServicioMunicion service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicioMunicionClient interface {
	SolicitarM(ctx context.Context, in *SolicitarMunicion, opts ...grpc.CallOption) (*RespuestaMunicion, error)
}

type servicioMunicionClient struct {
	cc grpc.ClientConnInterface
}

func NewServicioMunicionClient(cc grpc.ClientConnInterface) ServicioMunicionClient {
	return &servicioMunicionClient{cc}
}

func (c *servicioMunicionClient) SolicitarM(ctx context.Context, in *SolicitarMunicion, opts ...grpc.CallOption) (*RespuestaMunicion, error) {
	out := new(RespuestaMunicion)
	err := c.cc.Invoke(ctx, ServicioMunicion_SolicitarM_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicioMunicionServer is the server API for ServicioMunicion service.
// All implementations must embed UnimplementedServicioMunicionServer
// for forward compatibility
type ServicioMunicionServer interface {
	SolicitarM(context.Context, *SolicitarMunicion) (*RespuestaMunicion, error)
	mustEmbedUnimplementedServicioMunicionServer()
}

// UnimplementedServicioMunicionServer must be embedded to have forward compatible implementations.
type UnimplementedServicioMunicionServer struct {
}

func (UnimplementedServicioMunicionServer) SolicitarM(context.Context, *SolicitarMunicion) (*RespuestaMunicion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolicitarM not implemented")
}
func (UnimplementedServicioMunicionServer) mustEmbedUnimplementedServicioMunicionServer() {}

// UnsafeServicioMunicionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicioMunicionServer will
// result in compilation errors.
type UnsafeServicioMunicionServer interface {
	mustEmbedUnimplementedServicioMunicionServer()
}

func RegisterServicioMunicionServer(s grpc.ServiceRegistrar, srv ServicioMunicionServer) {
	s.RegisterService(&ServicioMunicion_ServiceDesc, srv)
}

func _ServicioMunicion_SolicitarM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitarMunicion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioMunicionServer).SolicitarM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicioMunicion_SolicitarM_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioMunicionServer).SolicitarM(ctx, req.(*SolicitarMunicion))
	}
	return interceptor(ctx, in, info, handler)
}

// ServicioMunicion_ServiceDesc is the grpc.ServiceDesc for ServicioMunicion service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServicioMunicion_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Proto.servicioMunicion",
	HandlerType: (*ServicioMunicionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SolicitarM",
			Handler:    _ServicioMunicion_SolicitarM_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Proto/municion.proto",
}
