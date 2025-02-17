// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: contacts.proto

package contacts

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ContactService_AddContact_FullMethodName                 = "/ContactService/AddContact"
	ContactService_RemoveContact_FullMethodName              = "/ContactService/RemoveContact"
	ContactService_AcceptContact_FullMethodName              = "/ContactService/AcceptContact"
	ContactService_GetContacts_FullMethodName                = "/ContactService/GetContacts"
	ContactService_GetPendingSentContacts_FullMethodName     = "/ContactService/GetPendingSentContacts"
	ContactService_GetPendingReceivedContacts_FullMethodName = "/ContactService/GetPendingReceivedContacts"
	ContactService_RejectContact_FullMethodName              = "/ContactService/RejectContact"
	ContactService_GetContactsNotInGroup_FullMethodName      = "/ContactService/GetContactsNotInGroup"
)

// ContactServiceClient is the client API for ContactService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContactServiceClient interface {
	AddContact(ctx context.Context, in *AddContactRequest, opts ...grpc.CallOption) (*AddContactResponse, error)
	RemoveContact(ctx context.Context, in *RemoveContactRequest, opts ...grpc.CallOption) (*RemoveContactResponse, error)
	AcceptContact(ctx context.Context, in *AcceptContactRequest, opts ...grpc.CallOption) (*AcceptContactResponse, error)
	GetContacts(ctx context.Context, in *GetContactsRequest, opts ...grpc.CallOption) (*GetContactsResponse, error)
	GetPendingSentContacts(ctx context.Context, in *GetPendingSentContactsRequest, opts ...grpc.CallOption) (*GetPendingSentContactsResponse, error)
	GetPendingReceivedContacts(ctx context.Context, in *GetPendingReceivedContactsRequest, opts ...grpc.CallOption) (*GetPendingReceivedContactsResponse, error)
	RejectContact(ctx context.Context, in *RejectContactRequest, opts ...grpc.CallOption) (*RejectContactResponse, error)
	GetContactsNotInGroup(ctx context.Context, in *GetContactsNotInGroupRequest, opts ...grpc.CallOption) (*GetContactsNotInGroupResponse, error)
}

type contactServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContactServiceClient(cc grpc.ClientConnInterface) ContactServiceClient {
	return &contactServiceClient{cc}
}

func (c *contactServiceClient) AddContact(ctx context.Context, in *AddContactRequest, opts ...grpc.CallOption) (*AddContactResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddContactResponse)
	err := c.cc.Invoke(ctx, ContactService_AddContact_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) RemoveContact(ctx context.Context, in *RemoveContactRequest, opts ...grpc.CallOption) (*RemoveContactResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveContactResponse)
	err := c.cc.Invoke(ctx, ContactService_RemoveContact_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) AcceptContact(ctx context.Context, in *AcceptContactRequest, opts ...grpc.CallOption) (*AcceptContactResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AcceptContactResponse)
	err := c.cc.Invoke(ctx, ContactService_AcceptContact_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) GetContacts(ctx context.Context, in *GetContactsRequest, opts ...grpc.CallOption) (*GetContactsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetContactsResponse)
	err := c.cc.Invoke(ctx, ContactService_GetContacts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) GetPendingSentContacts(ctx context.Context, in *GetPendingSentContactsRequest, opts ...grpc.CallOption) (*GetPendingSentContactsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPendingSentContactsResponse)
	err := c.cc.Invoke(ctx, ContactService_GetPendingSentContacts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) GetPendingReceivedContacts(ctx context.Context, in *GetPendingReceivedContactsRequest, opts ...grpc.CallOption) (*GetPendingReceivedContactsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPendingReceivedContactsResponse)
	err := c.cc.Invoke(ctx, ContactService_GetPendingReceivedContacts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) RejectContact(ctx context.Context, in *RejectContactRequest, opts ...grpc.CallOption) (*RejectContactResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RejectContactResponse)
	err := c.cc.Invoke(ctx, ContactService_RejectContact_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) GetContactsNotInGroup(ctx context.Context, in *GetContactsNotInGroupRequest, opts ...grpc.CallOption) (*GetContactsNotInGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetContactsNotInGroupResponse)
	err := c.cc.Invoke(ctx, ContactService_GetContactsNotInGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactServiceServer is the server API for ContactService service.
// All implementations must embed UnimplementedContactServiceServer
// for forward compatibility.
type ContactServiceServer interface {
	AddContact(context.Context, *AddContactRequest) (*AddContactResponse, error)
	RemoveContact(context.Context, *RemoveContactRequest) (*RemoveContactResponse, error)
	AcceptContact(context.Context, *AcceptContactRequest) (*AcceptContactResponse, error)
	GetContacts(context.Context, *GetContactsRequest) (*GetContactsResponse, error)
	GetPendingSentContacts(context.Context, *GetPendingSentContactsRequest) (*GetPendingSentContactsResponse, error)
	GetPendingReceivedContacts(context.Context, *GetPendingReceivedContactsRequest) (*GetPendingReceivedContactsResponse, error)
	RejectContact(context.Context, *RejectContactRequest) (*RejectContactResponse, error)
	GetContactsNotInGroup(context.Context, *GetContactsNotInGroupRequest) (*GetContactsNotInGroupResponse, error)
	mustEmbedUnimplementedContactServiceServer()
}

// UnimplementedContactServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedContactServiceServer struct{}

func (UnimplementedContactServiceServer) AddContact(context.Context, *AddContactRequest) (*AddContactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddContact not implemented")
}
func (UnimplementedContactServiceServer) RemoveContact(context.Context, *RemoveContactRequest) (*RemoveContactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveContact not implemented")
}
func (UnimplementedContactServiceServer) AcceptContact(context.Context, *AcceptContactRequest) (*AcceptContactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptContact not implemented")
}
func (UnimplementedContactServiceServer) GetContacts(context.Context, *GetContactsRequest) (*GetContactsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContacts not implemented")
}
func (UnimplementedContactServiceServer) GetPendingSentContacts(context.Context, *GetPendingSentContactsRequest) (*GetPendingSentContactsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPendingSentContacts not implemented")
}
func (UnimplementedContactServiceServer) GetPendingReceivedContacts(context.Context, *GetPendingReceivedContactsRequest) (*GetPendingReceivedContactsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPendingReceivedContacts not implemented")
}
func (UnimplementedContactServiceServer) RejectContact(context.Context, *RejectContactRequest) (*RejectContactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectContact not implemented")
}
func (UnimplementedContactServiceServer) GetContactsNotInGroup(context.Context, *GetContactsNotInGroupRequest) (*GetContactsNotInGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContactsNotInGroup not implemented")
}
func (UnimplementedContactServiceServer) mustEmbedUnimplementedContactServiceServer() {}
func (UnimplementedContactServiceServer) testEmbeddedByValue()                        {}

// UnsafeContactServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContactServiceServer will
// result in compilation errors.
type UnsafeContactServiceServer interface {
	mustEmbedUnimplementedContactServiceServer()
}

func RegisterContactServiceServer(s grpc.ServiceRegistrar, srv ContactServiceServer) {
	// If the following call pancis, it indicates UnimplementedContactServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ContactService_ServiceDesc, srv)
}

func _ContactService_AddContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).AddContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContactService_AddContact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).AddContact(ctx, req.(*AddContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_RemoveContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).RemoveContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContactService_RemoveContact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).RemoveContact(ctx, req.(*RemoveContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_AcceptContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).AcceptContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContactService_AcceptContact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).AcceptContact(ctx, req.(*AcceptContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_GetContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContactsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).GetContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContactService_GetContacts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).GetContacts(ctx, req.(*GetContactsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_GetPendingSentContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPendingSentContactsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).GetPendingSentContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContactService_GetPendingSentContacts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).GetPendingSentContacts(ctx, req.(*GetPendingSentContactsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_GetPendingReceivedContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPendingReceivedContactsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).GetPendingReceivedContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContactService_GetPendingReceivedContacts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).GetPendingReceivedContacts(ctx, req.(*GetPendingReceivedContactsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_RejectContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).RejectContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContactService_RejectContact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).RejectContact(ctx, req.(*RejectContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_GetContactsNotInGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContactsNotInGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).GetContactsNotInGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContactService_GetContactsNotInGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).GetContactsNotInGroup(ctx, req.(*GetContactsNotInGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContactService_ServiceDesc is the grpc.ServiceDesc for ContactService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContactService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ContactService",
	HandlerType: (*ContactServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddContact",
			Handler:    _ContactService_AddContact_Handler,
		},
		{
			MethodName: "RemoveContact",
			Handler:    _ContactService_RemoveContact_Handler,
		},
		{
			MethodName: "AcceptContact",
			Handler:    _ContactService_AcceptContact_Handler,
		},
		{
			MethodName: "GetContacts",
			Handler:    _ContactService_GetContacts_Handler,
		},
		{
			MethodName: "GetPendingSentContacts",
			Handler:    _ContactService_GetPendingSentContacts_Handler,
		},
		{
			MethodName: "GetPendingReceivedContacts",
			Handler:    _ContactService_GetPendingReceivedContacts_Handler,
		},
		{
			MethodName: "RejectContact",
			Handler:    _ContactService_RejectContact_Handler,
		},
		{
			MethodName: "GetContactsNotInGroup",
			Handler:    _ContactService_GetContactsNotInGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contacts.proto",
}
