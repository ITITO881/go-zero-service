// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: manager_ito.proto

package pbmgrito

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

const (
	MgrController_AccountDept_FullMethodName             = "/manager_service.manager_ito.MgrController/AccountDept"
	MgrController_AccountStaffTypeOptions_FullMethodName = "/manager_service.manager_ito.MgrController/AccountStaffTypeOptions"
	MgrController_Create_FullMethodName                  = "/manager_service.manager_ito.MgrController/Create"
	MgrController_CreateWebRoute_FullMethodName          = "/manager_service.manager_ito.MgrController/CreateWebRoute"
	MgrController_DeleteWebRoute_FullMethodName          = "/manager_service.manager_ito.MgrController/DeleteWebRoute"
	MgrController_Destroy_FullMethodName                 = "/manager_service.manager_ito.MgrController/Destroy"
	MgrController_List_FullMethodName                    = "/manager_service.manager_ito.MgrController/List"
	MgrController_ListWebRoute_FullMethodName            = "/manager_service.manager_ito.MgrController/ListWebRoute"
	MgrController_Login_FullMethodName                   = "/manager_service.manager_ito.MgrController/Login"
	MgrController_ManagerWebRoute_FullMethodName         = "/manager_service.manager_ito.MgrController/ManagerWebRoute"
	MgrController_PartialUpdate_FullMethodName           = "/manager_service.manager_ito.MgrController/PartialUpdate"
	MgrController_Retrieve_FullMethodName                = "/manager_service.manager_ito.MgrController/Retrieve"
	MgrController_Update_FullMethodName                  = "/manager_service.manager_ito.MgrController/Update"
	MgrController_UpdateWebRoute_FullMethodName          = "/manager_service.manager_ito.MgrController/UpdateWebRoute"
)

// MgrControllerClient is the client API for MgrController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MgrControllerClient interface {
	AccountDept(ctx context.Context, in *MgrWebLoginRequest, opts ...grpc.CallOption) (*AccountsUsersinfoResponse, error)
	AccountStaffTypeOptions(ctx context.Context, in *AccountStaffTypeOptionsRequest, opts ...grpc.CallOption) (*AccountStaffTypeOptionsListResponse, error)
	Create(ctx context.Context, in *ManagerWebModelRequest, opts ...grpc.CallOption) (*ManagerWebModelResponse, error)
	CreateWebRoute(ctx context.Context, in *WebRouteModelRequest, opts ...grpc.CallOption) (*WebRouteModelResponse, error)
	DeleteWebRoute(ctx context.Context, in *MgrDeleteWebRouteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Destroy(ctx context.Context, in *ManagerWebModelDestroyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	List(ctx context.Context, in *ManagerWebModelListRequest, opts ...grpc.CallOption) (*ManagerWebModelListResponse, error)
	ListWebRoute(ctx context.Context, in *WebRouteModelRequest, opts ...grpc.CallOption) (*WebRouteModelListResponse, error)
	Login(ctx context.Context, in *MgrWebLoginRequest, opts ...grpc.CallOption) (*ManagerWebModelResponse, error)
	ManagerWebRoute(ctx context.Context, in *MgrManagerWebRouteRequest, opts ...grpc.CallOption) (*ManagerWebModelResponse, error)
	PartialUpdate(ctx context.Context, in *ManagerWebModelPartialUpdateRequest, opts ...grpc.CallOption) (*ManagerWebModelResponse, error)
	Retrieve(ctx context.Context, in *ManagerWebModelRetrieveRequest, opts ...grpc.CallOption) (*ManagerWebModelResponse, error)
	Update(ctx context.Context, in *ManagerWebModelRequest, opts ...grpc.CallOption) (*ManagerWebModelResponse, error)
	UpdateWebRoute(ctx context.Context, in *PartialUpdateRequestRequest, opts ...grpc.CallOption) (*WebRouteModelResponse, error)
}

type mgrControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewMgrControllerClient(cc grpc.ClientConnInterface) MgrControllerClient {
	return &mgrControllerClient{cc}
}

func (c *mgrControllerClient) AccountDept(ctx context.Context, in *MgrWebLoginRequest, opts ...grpc.CallOption) (*AccountsUsersinfoResponse, error) {
	out := new(AccountsUsersinfoResponse)
	err := c.cc.Invoke(ctx, MgrController_AccountDept_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgrControllerClient) AccountStaffTypeOptions(ctx context.Context, in *AccountStaffTypeOptionsRequest, opts ...grpc.CallOption) (*AccountStaffTypeOptionsListResponse, error) {
	out := new(AccountStaffTypeOptionsListResponse)
	err := c.cc.Invoke(ctx, MgrController_AccountStaffTypeOptions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgrControllerClient) Create(ctx context.Context, in *ManagerWebModelRequest, opts ...grpc.CallOption) (*ManagerWebModelResponse, error) {
	out := new(ManagerWebModelResponse)
	err := c.cc.Invoke(ctx, MgrController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgrControllerClient) CreateWebRoute(ctx context.Context, in *WebRouteModelRequest, opts ...grpc.CallOption) (*WebRouteModelResponse, error) {
	out := new(WebRouteModelResponse)
	err := c.cc.Invoke(ctx, MgrController_CreateWebRoute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgrControllerClient) DeleteWebRoute(ctx context.Context, in *MgrDeleteWebRouteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MgrController_DeleteWebRoute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgrControllerClient) Destroy(ctx context.Context, in *ManagerWebModelDestroyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MgrController_Destroy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgrControllerClient) List(ctx context.Context, in *ManagerWebModelListRequest, opts ...grpc.CallOption) (*ManagerWebModelListResponse, error) {
	out := new(ManagerWebModelListResponse)
	err := c.cc.Invoke(ctx, MgrController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgrControllerClient) ListWebRoute(ctx context.Context, in *WebRouteModelRequest, opts ...grpc.CallOption) (*WebRouteModelListResponse, error) {
	out := new(WebRouteModelListResponse)
	err := c.cc.Invoke(ctx, MgrController_ListWebRoute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgrControllerClient) Login(ctx context.Context, in *MgrWebLoginRequest, opts ...grpc.CallOption) (*ManagerWebModelResponse, error) {
	out := new(ManagerWebModelResponse)
	err := c.cc.Invoke(ctx, MgrController_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgrControllerClient) ManagerWebRoute(ctx context.Context, in *MgrManagerWebRouteRequest, opts ...grpc.CallOption) (*ManagerWebModelResponse, error) {
	out := new(ManagerWebModelResponse)
	err := c.cc.Invoke(ctx, MgrController_ManagerWebRoute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgrControllerClient) PartialUpdate(ctx context.Context, in *ManagerWebModelPartialUpdateRequest, opts ...grpc.CallOption) (*ManagerWebModelResponse, error) {
	out := new(ManagerWebModelResponse)
	err := c.cc.Invoke(ctx, MgrController_PartialUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgrControllerClient) Retrieve(ctx context.Context, in *ManagerWebModelRetrieveRequest, opts ...grpc.CallOption) (*ManagerWebModelResponse, error) {
	out := new(ManagerWebModelResponse)
	err := c.cc.Invoke(ctx, MgrController_Retrieve_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgrControllerClient) Update(ctx context.Context, in *ManagerWebModelRequest, opts ...grpc.CallOption) (*ManagerWebModelResponse, error) {
	out := new(ManagerWebModelResponse)
	err := c.cc.Invoke(ctx, MgrController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgrControllerClient) UpdateWebRoute(ctx context.Context, in *PartialUpdateRequestRequest, opts ...grpc.CallOption) (*WebRouteModelResponse, error) {
	out := new(WebRouteModelResponse)
	err := c.cc.Invoke(ctx, MgrController_UpdateWebRoute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MgrControllerServer is the server API for MgrController service.
// All implementations must embed UnimplementedMgrControllerServer
// for forward compatibility
type MgrControllerServer interface {
	AccountDept(context.Context, *MgrWebLoginRequest) (*AccountsUsersinfoResponse, error)
	AccountStaffTypeOptions(context.Context, *AccountStaffTypeOptionsRequest) (*AccountStaffTypeOptionsListResponse, error)
	Create(context.Context, *ManagerWebModelRequest) (*ManagerWebModelResponse, error)
	CreateWebRoute(context.Context, *WebRouteModelRequest) (*WebRouteModelResponse, error)
	DeleteWebRoute(context.Context, *MgrDeleteWebRouteRequest) (*emptypb.Empty, error)
	Destroy(context.Context, *ManagerWebModelDestroyRequest) (*emptypb.Empty, error)
	List(context.Context, *ManagerWebModelListRequest) (*ManagerWebModelListResponse, error)
	ListWebRoute(context.Context, *WebRouteModelRequest) (*WebRouteModelListResponse, error)
	Login(context.Context, *MgrWebLoginRequest) (*ManagerWebModelResponse, error)
	ManagerWebRoute(context.Context, *MgrManagerWebRouteRequest) (*ManagerWebModelResponse, error)
	PartialUpdate(context.Context, *ManagerWebModelPartialUpdateRequest) (*ManagerWebModelResponse, error)
	Retrieve(context.Context, *ManagerWebModelRetrieveRequest) (*ManagerWebModelResponse, error)
	Update(context.Context, *ManagerWebModelRequest) (*ManagerWebModelResponse, error)
	UpdateWebRoute(context.Context, *PartialUpdateRequestRequest) (*WebRouteModelResponse, error)
	mustEmbedUnimplementedMgrControllerServer()
}

// UnimplementedMgrControllerServer must be embedded to have forward compatible implementations.
type UnimplementedMgrControllerServer struct {
}

func (UnimplementedMgrControllerServer) AccountDept(context.Context, *MgrWebLoginRequest) (*AccountsUsersinfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountDept not implemented")
}
func (UnimplementedMgrControllerServer) AccountStaffTypeOptions(context.Context, *AccountStaffTypeOptionsRequest) (*AccountStaffTypeOptionsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountStaffTypeOptions not implemented")
}
func (UnimplementedMgrControllerServer) Create(context.Context, *ManagerWebModelRequest) (*ManagerWebModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMgrControllerServer) CreateWebRoute(context.Context, *WebRouteModelRequest) (*WebRouteModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWebRoute not implemented")
}
func (UnimplementedMgrControllerServer) DeleteWebRoute(context.Context, *MgrDeleteWebRouteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWebRoute not implemented")
}
func (UnimplementedMgrControllerServer) Destroy(context.Context, *ManagerWebModelDestroyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}
func (UnimplementedMgrControllerServer) List(context.Context, *ManagerWebModelListRequest) (*ManagerWebModelListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedMgrControllerServer) ListWebRoute(context.Context, *WebRouteModelRequest) (*WebRouteModelListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWebRoute not implemented")
}
func (UnimplementedMgrControllerServer) Login(context.Context, *MgrWebLoginRequest) (*ManagerWebModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedMgrControllerServer) ManagerWebRoute(context.Context, *MgrManagerWebRouteRequest) (*ManagerWebModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManagerWebRoute not implemented")
}
func (UnimplementedMgrControllerServer) PartialUpdate(context.Context, *ManagerWebModelPartialUpdateRequest) (*ManagerWebModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PartialUpdate not implemented")
}
func (UnimplementedMgrControllerServer) Retrieve(context.Context, *ManagerWebModelRetrieveRequest) (*ManagerWebModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedMgrControllerServer) Update(context.Context, *ManagerWebModelRequest) (*ManagerWebModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMgrControllerServer) UpdateWebRoute(context.Context, *PartialUpdateRequestRequest) (*WebRouteModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWebRoute not implemented")
}
func (UnimplementedMgrControllerServer) mustEmbedUnimplementedMgrControllerServer() {}

// UnsafeMgrControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MgrControllerServer will
// result in compilation errors.
type UnsafeMgrControllerServer interface {
	mustEmbedUnimplementedMgrControllerServer()
}

func RegisterMgrControllerServer(s grpc.ServiceRegistrar, srv MgrControllerServer) {
	s.RegisterService(&MgrController_ServiceDesc, srv)
}

func _MgrController_AccountDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MgrWebLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgrControllerServer).AccountDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgrController_AccountDept_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgrControllerServer).AccountDept(ctx, req.(*MgrWebLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgrController_AccountStaffTypeOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountStaffTypeOptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgrControllerServer).AccountStaffTypeOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgrController_AccountStaffTypeOptions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgrControllerServer).AccountStaffTypeOptions(ctx, req.(*AccountStaffTypeOptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgrController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManagerWebModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgrControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgrController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgrControllerServer).Create(ctx, req.(*ManagerWebModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgrController_CreateWebRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebRouteModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgrControllerServer).CreateWebRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgrController_CreateWebRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgrControllerServer).CreateWebRoute(ctx, req.(*WebRouteModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgrController_DeleteWebRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MgrDeleteWebRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgrControllerServer).DeleteWebRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgrController_DeleteWebRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgrControllerServer).DeleteWebRoute(ctx, req.(*MgrDeleteWebRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgrController_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManagerWebModelDestroyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgrControllerServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgrController_Destroy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgrControllerServer).Destroy(ctx, req.(*ManagerWebModelDestroyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgrController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManagerWebModelListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgrControllerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgrController_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgrControllerServer).List(ctx, req.(*ManagerWebModelListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgrController_ListWebRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebRouteModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgrControllerServer).ListWebRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgrController_ListWebRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgrControllerServer).ListWebRoute(ctx, req.(*WebRouteModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgrController_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MgrWebLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgrControllerServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgrController_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgrControllerServer).Login(ctx, req.(*MgrWebLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgrController_ManagerWebRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MgrManagerWebRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgrControllerServer).ManagerWebRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgrController_ManagerWebRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgrControllerServer).ManagerWebRoute(ctx, req.(*MgrManagerWebRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgrController_PartialUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManagerWebModelPartialUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgrControllerServer).PartialUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgrController_PartialUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgrControllerServer).PartialUpdate(ctx, req.(*ManagerWebModelPartialUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgrController_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManagerWebModelRetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgrControllerServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgrController_Retrieve_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgrControllerServer).Retrieve(ctx, req.(*ManagerWebModelRetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgrController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManagerWebModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgrControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgrController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgrControllerServer).Update(ctx, req.(*ManagerWebModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgrController_UpdateWebRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PartialUpdateRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgrControllerServer).UpdateWebRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgrController_UpdateWebRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgrControllerServer).UpdateWebRoute(ctx, req.(*PartialUpdateRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MgrController_ServiceDesc is the grpc.ServiceDesc for MgrController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MgrController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manager_service.manager_ito.MgrController",
	HandlerType: (*MgrControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AccountDept",
			Handler:    _MgrController_AccountDept_Handler,
		},
		{
			MethodName: "AccountStaffTypeOptions",
			Handler:    _MgrController_AccountStaffTypeOptions_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _MgrController_Create_Handler,
		},
		{
			MethodName: "CreateWebRoute",
			Handler:    _MgrController_CreateWebRoute_Handler,
		},
		{
			MethodName: "DeleteWebRoute",
			Handler:    _MgrController_DeleteWebRoute_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _MgrController_Destroy_Handler,
		},
		{
			MethodName: "List",
			Handler:    _MgrController_List_Handler,
		},
		{
			MethodName: "ListWebRoute",
			Handler:    _MgrController_ListWebRoute_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _MgrController_Login_Handler,
		},
		{
			MethodName: "ManagerWebRoute",
			Handler:    _MgrController_ManagerWebRoute_Handler,
		},
		{
			MethodName: "PartialUpdate",
			Handler:    _MgrController_PartialUpdate_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _MgrController_Retrieve_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _MgrController_Update_Handler,
		},
		{
			MethodName: "UpdateWebRoute",
			Handler:    _MgrController_UpdateWebRoute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager_ito.proto",
}
