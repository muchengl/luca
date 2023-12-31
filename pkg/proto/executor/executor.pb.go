// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: pkg/proto/executor/executor.proto

package executor

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_executor_executor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_executor_executor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_executor_executor_proto_rawDescGZIP(), []int{0}
}

func (x *HeartbeatRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type HeartbeatReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HeartbeatReply) Reset() {
	*x = HeartbeatReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_executor_executor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatReply) ProtoMessage() {}

func (x *HeartbeatReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_executor_executor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatReply.ProtoReflect.Descriptor instead.
func (*HeartbeatReply) Descriptor() ([]byte, []int) {
	return file_pkg_proto_executor_executor_proto_rawDescGZIP(), []int{1}
}

func (x *HeartbeatReply) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *HeartbeatReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type TaskInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskID          string `protobuf:"bytes,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
	FuncName        string `protobuf:"bytes,2,opt,name=funcName,proto3" json:"funcName,omitempty"`
	FuncID          string `protobuf:"bytes,8,opt,name=funcID,proto3" json:"funcID,omitempty"`
	FuncStorageType string `protobuf:"bytes,3,opt,name=funcStorageType,proto3" json:"funcStorageType,omitempty"` // s3
	TaskFuncPath    string `protobuf:"bytes,4,opt,name=taskFuncPath,proto3" json:"taskFuncPath,omitempty"`
	TaskRunningMode string `protobuf:"bytes,5,opt,name=taskRunningMode,proto3" json:"taskRunningMode,omitempty"` //sync or async
	FuncType        string `protobuf:"bytes,11,opt,name=funcType,proto3" json:"funcType,omitempty"`              // Go or python ......
	TaskCPUCore     int64  `protobuf:"varint,6,opt,name=taskCPUCore,proto3" json:"taskCPUCore,omitempty"`
	TaskMem         int64  `protobuf:"varint,7,opt,name=taskMem,proto3" json:"taskMem,omitempty"`
	TaskDiskSpace   int64  `protobuf:"varint,9,opt,name=taskDiskSpace,proto3" json:"taskDiskSpace,omitempty"`
	TaskMaxTime     int64  `protobuf:"varint,10,opt,name=taskMaxTime,proto3" json:"taskMaxTime,omitempty"`
}

func (x *TaskInstance) Reset() {
	*x = TaskInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_executor_executor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskInstance) ProtoMessage() {}

func (x *TaskInstance) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_executor_executor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskInstance.ProtoReflect.Descriptor instead.
func (*TaskInstance) Descriptor() ([]byte, []int) {
	return file_pkg_proto_executor_executor_proto_rawDescGZIP(), []int{2}
}

func (x *TaskInstance) GetTaskID() string {
	if x != nil {
		return x.TaskID
	}
	return ""
}

func (x *TaskInstance) GetFuncName() string {
	if x != nil {
		return x.FuncName
	}
	return ""
}

func (x *TaskInstance) GetFuncID() string {
	if x != nil {
		return x.FuncID
	}
	return ""
}

func (x *TaskInstance) GetFuncStorageType() string {
	if x != nil {
		return x.FuncStorageType
	}
	return ""
}

func (x *TaskInstance) GetTaskFuncPath() string {
	if x != nil {
		return x.TaskFuncPath
	}
	return ""
}

func (x *TaskInstance) GetTaskRunningMode() string {
	if x != nil {
		return x.TaskRunningMode
	}
	return ""
}

func (x *TaskInstance) GetFuncType() string {
	if x != nil {
		return x.FuncType
	}
	return ""
}

func (x *TaskInstance) GetTaskCPUCore() int64 {
	if x != nil {
		return x.TaskCPUCore
	}
	return 0
}

func (x *TaskInstance) GetTaskMem() int64 {
	if x != nil {
		return x.TaskMem
	}
	return 0
}

func (x *TaskInstance) GetTaskDiskSpace() int64 {
	if x != nil {
		return x.TaskDiskSpace
	}
	return 0
}

func (x *TaskInstance) GetTaskMaxTime() int64 {
	if x != nil {
		return x.TaskMaxTime
	}
	return 0
}

type InitTaskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *InitTaskReply) Reset() {
	*x = InitTaskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_executor_executor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitTaskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitTaskReply) ProtoMessage() {}

func (x *InitTaskReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_executor_executor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitTaskReply.ProtoReflect.Descriptor instead.
func (*InitTaskReply) Descriptor() ([]byte, []int) {
	return file_pkg_proto_executor_executor_proto_rawDescGZIP(), []int{3}
}

func (x *InitTaskReply) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *InitTaskReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type RunTaskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *RunTaskReply) Reset() {
	*x = RunTaskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_executor_executor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunTaskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunTaskReply) ProtoMessage() {}

func (x *RunTaskReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_executor_executor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunTaskReply.ProtoReflect.Descriptor instead.
func (*RunTaskReply) Descriptor() ([]byte, []int) {
	return file_pkg_proto_executor_executor_proto_rawDescGZIP(), []int{4}
}

func (x *RunTaskReply) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RunTaskReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type KillTaskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *KillTaskReply) Reset() {
	*x = KillTaskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_executor_executor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KillTaskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KillTaskReply) ProtoMessage() {}

func (x *KillTaskReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_executor_executor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KillTaskReply.ProtoReflect.Descriptor instead.
func (*KillTaskReply) Descriptor() ([]byte, []int) {
	return file_pkg_proto_executor_executor_proto_rawDescGZIP(), []int{5}
}

func (x *KillTaskReply) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *KillTaskReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_pkg_proto_executor_executor_proto protoreflect.FileDescriptor

var file_pkg_proto_executor_executor_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x6f, 0x72, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x22, 0x24, 0x0a,
	0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x3a, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22,
	0xf2, 0x02, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x75, 0x6e, 0x63, 0x49, 0x44, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x75, 0x6e, 0x63, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f,
	0x66, 0x75, 0x6e, 0x63, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x75, 0x6e, 0x63, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x46, 0x75,
	0x6e, 0x63, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x61,
	0x73, 0x6b, 0x46, 0x75, 0x6e, 0x63, 0x50, 0x61, 0x74, 0x68, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x61,
	0x73, 0x6b, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x43, 0x50, 0x55, 0x43, 0x6f, 0x72, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x43, 0x50, 0x55, 0x43, 0x6f,
	0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x6d, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x6d, 0x12, 0x24, 0x0a, 0x0d,
	0x74, 0x61, 0x73, 0x6b, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x70, 0x61, 0x63, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x4d, 0x61, 0x78, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x4d, 0x61, 0x78,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x0d, 0x49, 0x6e, 0x69, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22,
	0x38, 0x0a, 0x0c, 0x52, 0x75, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x39, 0x0a, 0x0d, 0x4b, 0x69, 0x6c,
	0x6c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x32, 0x8a, 0x02, 0x0a, 0x08, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x12, 0x43, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x1a,
	0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x08, 0x49, 0x6e, 0x69, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x16, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x17, 0x2e, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x16, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x16, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x6f, 0x72, 0x2e, 0x52, 0x75, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x08, 0x4b, 0x69, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x16,
	0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x17, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x2e, 0x4b, 0x69, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x1f, 0x5a, 0x1d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x2e, 0x66, 0x61, 0x61, 0x73, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_executor_executor_proto_rawDescOnce sync.Once
	file_pkg_proto_executor_executor_proto_rawDescData = file_pkg_proto_executor_executor_proto_rawDesc
)

func file_pkg_proto_executor_executor_proto_rawDescGZIP() []byte {
	file_pkg_proto_executor_executor_proto_rawDescOnce.Do(func() {
		file_pkg_proto_executor_executor_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_executor_executor_proto_rawDescData)
	})
	return file_pkg_proto_executor_executor_proto_rawDescData
}

var file_pkg_proto_executor_executor_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_proto_executor_executor_proto_goTypes = []interface{}{
	(*HeartbeatRequest)(nil), // 0: executor.HeartbeatRequest
	(*HeartbeatReply)(nil),   // 1: executor.HeartbeatReply
	(*TaskInstance)(nil),     // 2: executor.TaskInstance
	(*InitTaskReply)(nil),    // 3: executor.InitTaskReply
	(*RunTaskReply)(nil),     // 4: executor.RunTaskReply
	(*KillTaskReply)(nil),    // 5: executor.KillTaskReply
}
var file_pkg_proto_executor_executor_proto_depIdxs = []int32{
	0, // 0: executor.executor.Heartbeat:input_type -> executor.HeartbeatRequest
	2, // 1: executor.executor.InitTask:input_type -> executor.TaskInstance
	2, // 2: executor.executor.RunTask:input_type -> executor.TaskInstance
	2, // 3: executor.executor.KillTask:input_type -> executor.TaskInstance
	1, // 4: executor.executor.Heartbeat:output_type -> executor.HeartbeatReply
	3, // 5: executor.executor.InitTask:output_type -> executor.InitTaskReply
	4, // 6: executor.executor.RunTask:output_type -> executor.RunTaskReply
	5, // 7: executor.executor.KillTask:output_type -> executor.KillTaskReply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_executor_executor_proto_init() }
func file_pkg_proto_executor_executor_proto_init() {
	if File_pkg_proto_executor_executor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_executor_executor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_executor_executor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_executor_executor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskInstance); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_executor_executor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitTaskReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_executor_executor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunTaskReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_executor_executor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KillTaskReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_executor_executor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_executor_executor_proto_goTypes,
		DependencyIndexes: file_pkg_proto_executor_executor_proto_depIdxs,
		MessageInfos:      file_pkg_proto_executor_executor_proto_msgTypes,
	}.Build()
	File_pkg_proto_executor_executor_proto = out.File
	file_pkg_proto_executor_executor_proto_rawDesc = nil
	file_pkg_proto_executor_executor_proto_goTypes = nil
	file_pkg_proto_executor_executor_proto_depIdxs = nil
}
