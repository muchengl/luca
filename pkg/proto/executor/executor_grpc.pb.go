// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: pkg/proto/executor/executor.proto

package executor

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

// ExecutorClient is the client API for Executor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExecutorClient interface {
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatReply, error)
	// init the running env of a task.
	InitTask(ctx context.Context, in *TaskInstance, opts ...grpc.CallOption) (*InitTaskReply, error)
	RunTask(ctx context.Context, in *TaskInstance, opts ...grpc.CallOption) (*RunTaskReply, error)
	KillTask(ctx context.Context, in *TaskInstance, opts ...grpc.CallOption) (*KillTaskReply, error)
}

type executorClient struct {
	cc grpc.ClientConnInterface
}

func NewExecutorClient(cc grpc.ClientConnInterface) ExecutorClient {
	return &executorClient{cc}
}

func (c *executorClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatReply, error) {
	out := new(HeartbeatReply)
	err := c.cc.Invoke(ctx, "/executor.executor/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) InitTask(ctx context.Context, in *TaskInstance, opts ...grpc.CallOption) (*InitTaskReply, error) {
	out := new(InitTaskReply)
	err := c.cc.Invoke(ctx, "/executor.executor/InitTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) RunTask(ctx context.Context, in *TaskInstance, opts ...grpc.CallOption) (*RunTaskReply, error) {
	out := new(RunTaskReply)
	err := c.cc.Invoke(ctx, "/executor.executor/RunTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) KillTask(ctx context.Context, in *TaskInstance, opts ...grpc.CallOption) (*KillTaskReply, error) {
	out := new(KillTaskReply)
	err := c.cc.Invoke(ctx, "/executor.executor/KillTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutorServer is the server API for Executor service.
// All implementations must embed UnimplementedExecutorServer
// for forward compatibility
type ExecutorServer interface {
	Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatReply, error)
	// init the running env of a task.
	InitTask(context.Context, *TaskInstance) (*InitTaskReply, error)
	RunTask(context.Context, *TaskInstance) (*RunTaskReply, error)
	KillTask(context.Context, *TaskInstance) (*KillTaskReply, error)
	mustEmbedUnimplementedExecutorServer()
}

// UnimplementedExecutorServer must be embedded to have forward compatible implementations.
type UnimplementedExecutorServer struct {
}

func (UnimplementedExecutorServer) Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedExecutorServer) InitTask(context.Context, *TaskInstance) (*InitTaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitTask not implemented")
}
func (UnimplementedExecutorServer) RunTask(context.Context, *TaskInstance) (*RunTaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunTask not implemented")
}
func (UnimplementedExecutorServer) KillTask(context.Context, *TaskInstance) (*KillTaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KillTask not implemented")
}
func (UnimplementedExecutorServer) mustEmbedUnimplementedExecutorServer() {}

// UnsafeExecutorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExecutorServer will
// result in compilation errors.
type UnsafeExecutorServer interface {
	mustEmbedUnimplementedExecutorServer()
}

func RegisterExecutorServer(s grpc.ServiceRegistrar, srv ExecutorServer) {
	s.RegisterService(&Executor_ServiceDesc, srv)
}

func _Executor_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/executor.executor/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_InitTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskInstance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).InitTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/executor.executor/InitTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).InitTask(ctx, req.(*TaskInstance))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_RunTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskInstance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).RunTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/executor.executor/RunTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).RunTask(ctx, req.(*TaskInstance))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_KillTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskInstance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).KillTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/executor.executor/KillTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).KillTask(ctx, req.(*TaskInstance))
	}
	return interceptor(ctx, in, info, handler)
}

// Executor_ServiceDesc is the grpc.ServiceDesc for Executor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Executor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "executor.executor",
	HandlerType: (*ExecutorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heartbeat",
			Handler:    _Executor_Heartbeat_Handler,
		},
		{
			MethodName: "InitTask",
			Handler:    _Executor_InitTask_Handler,
		},
		{
			MethodName: "RunTask",
			Handler:    _Executor_RunTask_Handler,
		},
		{
			MethodName: "KillTask",
			Handler:    _Executor_KillTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/executor/executor.proto",
}
