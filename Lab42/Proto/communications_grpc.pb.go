// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: communications.proto

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
	Mercenary_StartFloor_FullMethodName = "/Proto.Mercenary/StartFloor"
	Mercenary_ReportMove_FullMethodName = "/Proto.Mercenary/ReportMove"
)

// MercenaryClient is the client API for Mercenary service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MercenaryClient interface {
	StartFloor(ctx context.Context, in *FloorInfo, opts ...grpc.CallOption) (*Empty, error)
	ReportMove(ctx context.Context, in *Move, opts ...grpc.CallOption) (*Empty, error)
}

type mercenaryClient struct {
	cc grpc.ClientConnInterface
}

func NewMercenaryClient(cc grpc.ClientConnInterface) MercenaryClient {
	return &mercenaryClient{cc}
}

func (c *mercenaryClient) StartFloor(ctx context.Context, in *FloorInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Mercenary_StartFloor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mercenaryClient) ReportMove(ctx context.Context, in *Move, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Mercenary_ReportMove_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MercenaryServer is the server API for Mercenary service.
// All implementations must embed UnimplementedMercenaryServer
// for forward compatibility
type MercenaryServer interface {
	StartFloor(context.Context, *FloorInfo) (*Empty, error)
	ReportMove(context.Context, *Move) (*Empty, error)
	mustEmbedUnimplementedMercenaryServer()
}

// UnimplementedMercenaryServer must be embedded to have forward compatible implementations.
type UnimplementedMercenaryServer struct {
}

func (UnimplementedMercenaryServer) StartFloor(context.Context, *FloorInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFloor not implemented")
}
func (UnimplementedMercenaryServer) ReportMove(context.Context, *Move) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportMove not implemented")
}
func (UnimplementedMercenaryServer) mustEmbedUnimplementedMercenaryServer() {}

// UnsafeMercenaryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MercenaryServer will
// result in compilation errors.
type UnsafeMercenaryServer interface {
	mustEmbedUnimplementedMercenaryServer()
}

func RegisterMercenaryServer(s grpc.ServiceRegistrar, srv MercenaryServer) {
	s.RegisterService(&Mercenary_ServiceDesc, srv)
}

func _Mercenary_StartFloor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FloorInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MercenaryServer).StartFloor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mercenary_StartFloor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MercenaryServer).StartFloor(ctx, req.(*FloorInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mercenary_ReportMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Move)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MercenaryServer).ReportMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mercenary_ReportMove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MercenaryServer).ReportMove(ctx, req.(*Move))
	}
	return interceptor(ctx, in, info, handler)
}

// Mercenary_ServiceDesc is the grpc.ServiceDesc for Mercenary service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mercenary_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Proto.Mercenary",
	HandlerType: (*MercenaryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartFloor",
			Handler:    _Mercenary_StartFloor_Handler,
		},
		{
			MethodName: "ReportMove",
			Handler:    _Mercenary_ReportMove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "communications.proto",
}

const (
	Director_AddMercenary_FullMethodName             = "/Proto.Director/AddMercenary"
	Director_ReportDeath_FullMethodName              = "/Proto.Director/ReportDeath"
	Director_ReportMove_FullMethodName               = "/Proto.Director/ReportMove"
	Director_CheckStatus_FullMethodName              = "/Proto.Director/CheckStatus"
	Director_RequestAccumulatedAmount_FullMethodName = "/Proto.Director/RequestAccumulatedAmount"
)

// DirectorClient is the client API for Director service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DirectorClient interface {
	AddMercenary(ctx context.Context, in *MercenaryInfo, opts ...grpc.CallOption) (*Empty, error)
	ReportDeath(ctx context.Context, in *DeathInfo, opts ...grpc.CallOption) (*Empty, error)
	ReportMove(ctx context.Context, in *MoveInfo, opts ...grpc.CallOption) (*Empty, error)
	CheckStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StatusResponse, error)
	RequestAccumulatedAmount(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AccumulatedAmountResponse, error)
}

type directorClient struct {
	cc grpc.ClientConnInterface
}

func NewDirectorClient(cc grpc.ClientConnInterface) DirectorClient {
	return &directorClient{cc}
}

func (c *directorClient) AddMercenary(ctx context.Context, in *MercenaryInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Director_AddMercenary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directorClient) ReportDeath(ctx context.Context, in *DeathInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Director_ReportDeath_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directorClient) ReportMove(ctx context.Context, in *MoveInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Director_ReportMove_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directorClient) CheckStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, Director_CheckStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directorClient) RequestAccumulatedAmount(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AccumulatedAmountResponse, error) {
	out := new(AccumulatedAmountResponse)
	err := c.cc.Invoke(ctx, Director_RequestAccumulatedAmount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DirectorServer is the server API for Director service.
// All implementations must embed UnimplementedDirectorServer
// for forward compatibility
type DirectorServer interface {
	AddMercenary(context.Context, *MercenaryInfo) (*Empty, error)
	ReportDeath(context.Context, *DeathInfo) (*Empty, error)
	ReportMove(context.Context, *MoveInfo) (*Empty, error)
	CheckStatus(context.Context, *Empty) (*StatusResponse, error)
	RequestAccumulatedAmount(context.Context, *Empty) (*AccumulatedAmountResponse, error)
	mustEmbedUnimplementedDirectorServer()
}

// UnimplementedDirectorServer must be embedded to have forward compatible implementations.
type UnimplementedDirectorServer struct {
}

func (UnimplementedDirectorServer) AddMercenary(context.Context, *MercenaryInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMercenary not implemented")
}
func (UnimplementedDirectorServer) ReportDeath(context.Context, *DeathInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportDeath not implemented")
}
func (UnimplementedDirectorServer) ReportMove(context.Context, *MoveInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportMove not implemented")
}
func (UnimplementedDirectorServer) CheckStatus(context.Context, *Empty) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckStatus not implemented")
}
func (UnimplementedDirectorServer) RequestAccumulatedAmount(context.Context, *Empty) (*AccumulatedAmountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestAccumulatedAmount not implemented")
}
func (UnimplementedDirectorServer) mustEmbedUnimplementedDirectorServer() {}

// UnsafeDirectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DirectorServer will
// result in compilation errors.
type UnsafeDirectorServer interface {
	mustEmbedUnimplementedDirectorServer()
}

func RegisterDirectorServer(s grpc.ServiceRegistrar, srv DirectorServer) {
	s.RegisterService(&Director_ServiceDesc, srv)
}

func _Director_AddMercenary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MercenaryInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectorServer).AddMercenary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Director_AddMercenary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectorServer).AddMercenary(ctx, req.(*MercenaryInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Director_ReportDeath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeathInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectorServer).ReportDeath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Director_ReportDeath_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectorServer).ReportDeath(ctx, req.(*DeathInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Director_ReportMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectorServer).ReportMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Director_ReportMove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectorServer).ReportMove(ctx, req.(*MoveInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Director_CheckStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectorServer).CheckStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Director_CheckStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectorServer).CheckStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Director_RequestAccumulatedAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectorServer).RequestAccumulatedAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Director_RequestAccumulatedAmount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectorServer).RequestAccumulatedAmount(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Director_ServiceDesc is the grpc.ServiceDesc for Director service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Director_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Proto.Director",
	HandlerType: (*DirectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMercenary",
			Handler:    _Director_AddMercenary_Handler,
		},
		{
			MethodName: "ReportDeath",
			Handler:    _Director_ReportDeath_Handler,
		},
		{
			MethodName: "ReportMove",
			Handler:    _Director_ReportMove_Handler,
		},
		{
			MethodName: "CheckStatus",
			Handler:    _Director_CheckStatus_Handler,
		},
		{
			MethodName: "RequestAccumulatedAmount",
			Handler:    _Director_RequestAccumulatedAmount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "communications.proto",
}

const (
	DoshBank_GetAccumulatedAmount_FullMethodName = "/Proto.DoshBank/GetAccumulatedAmount"
)

// DoshBankClient is the client API for DoshBank service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DoshBankClient interface {
	GetAccumulatedAmount(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AccumulatedAmountResponse, error)
}

type doshBankClient struct {
	cc grpc.ClientConnInterface
}

func NewDoshBankClient(cc grpc.ClientConnInterface) DoshBankClient {
	return &doshBankClient{cc}
}

func (c *doshBankClient) GetAccumulatedAmount(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AccumulatedAmountResponse, error) {
	out := new(AccumulatedAmountResponse)
	err := c.cc.Invoke(ctx, DoshBank_GetAccumulatedAmount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DoshBankServer is the server API for DoshBank service.
// All implementations must embed UnimplementedDoshBankServer
// for forward compatibility
type DoshBankServer interface {
	GetAccumulatedAmount(context.Context, *Empty) (*AccumulatedAmountResponse, error)
	mustEmbedUnimplementedDoshBankServer()
}

// UnimplementedDoshBankServer must be embedded to have forward compatible implementations.
type UnimplementedDoshBankServer struct {
}

func (UnimplementedDoshBankServer) GetAccumulatedAmount(context.Context, *Empty) (*AccumulatedAmountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccumulatedAmount not implemented")
}
func (UnimplementedDoshBankServer) mustEmbedUnimplementedDoshBankServer() {}

// UnsafeDoshBankServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DoshBankServer will
// result in compilation errors.
type UnsafeDoshBankServer interface {
	mustEmbedUnimplementedDoshBankServer()
}

func RegisterDoshBankServer(s grpc.ServiceRegistrar, srv DoshBankServer) {
	s.RegisterService(&DoshBank_ServiceDesc, srv)
}

func _DoshBank_GetAccumulatedAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoshBankServer).GetAccumulatedAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoshBank_GetAccumulatedAmount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoshBankServer).GetAccumulatedAmount(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// DoshBank_ServiceDesc is the grpc.ServiceDesc for DoshBank service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DoshBank_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Proto.DoshBank",
	HandlerType: (*DoshBankServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccumulatedAmount",
			Handler:    _DoshBank_GetAccumulatedAmount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "communications.proto",
}
