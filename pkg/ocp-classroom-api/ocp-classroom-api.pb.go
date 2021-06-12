// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.2
// source: api/ocp-classroom-api/ocp-classroom-api.proto

package ocp_classroom_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListClassroomsV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListClassroomsV1Request) Reset() {
	*x = ListClassroomsV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClassroomsV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClassroomsV1Request) ProtoMessage() {}

func (x *ListClassroomsV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClassroomsV1Request.ProtoReflect.Descriptor instead.
func (*ListClassroomsV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDescGZIP(), []int{0}
}

func (x *ListClassroomsV1Request) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListClassroomsV1Request) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListClassroomsV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Classrooms []*Classroom `protobuf:"bytes,1,rep,name=classrooms,proto3" json:"classrooms,omitempty"`
}

func (x *ListClassroomsV1Response) Reset() {
	*x = ListClassroomsV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClassroomsV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClassroomsV1Response) ProtoMessage() {}

func (x *ListClassroomsV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClassroomsV1Response.ProtoReflect.Descriptor instead.
func (*ListClassroomsV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDescGZIP(), []int{1}
}

func (x *ListClassroomsV1Response) GetClassrooms() []*Classroom {
	if x != nil {
		return x.Classrooms
	}
	return nil
}

type DescribeClassroomV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassroomId uint64 `protobuf:"varint,1,opt,name=classroom_id,json=classroomId,proto3" json:"classroom_id,omitempty"`
}

func (x *DescribeClassroomV1Request) Reset() {
	*x = DescribeClassroomV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeClassroomV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeClassroomV1Request) ProtoMessage() {}

func (x *DescribeClassroomV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeClassroomV1Request.ProtoReflect.Descriptor instead.
func (*DescribeClassroomV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeClassroomV1Request) GetClassroomId() uint64 {
	if x != nil {
		return x.ClassroomId
	}
	return 0
}

type DescribeClassroomV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Classroom *Classroom `protobuf:"bytes,1,opt,name=classroom,proto3" json:"classroom,omitempty"`
}

func (x *DescribeClassroomV1Response) Reset() {
	*x = DescribeClassroomV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeClassroomV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeClassroomV1Response) ProtoMessage() {}

func (x *DescribeClassroomV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeClassroomV1Response.ProtoReflect.Descriptor instead.
func (*DescribeClassroomV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeClassroomV1Response) GetClassroom() *Classroom {
	if x != nil {
		return x.Classroom
	}
	return nil
}

type CreateClassroomV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId   uint64 `protobuf:"varint,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	CalendarId uint64 `protobuf:"varint,2,opt,name=calendar_id,json=calendarId,proto3" json:"calendar_id,omitempty"`
}

func (x *CreateClassroomV1Request) Reset() {
	*x = CreateClassroomV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClassroomV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClassroomV1Request) ProtoMessage() {}

func (x *CreateClassroomV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClassroomV1Request.ProtoReflect.Descriptor instead.
func (*CreateClassroomV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDescGZIP(), []int{4}
}

func (x *CreateClassroomV1Request) GetTenantId() uint64 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *CreateClassroomV1Request) GetCalendarId() uint64 {
	if x != nil {
		return x.CalendarId
	}
	return 0
}

type CreateClassroomV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassroomId uint64 `protobuf:"varint,1,opt,name=classroom_id,json=classroomId,proto3" json:"classroom_id,omitempty"`
}

func (x *CreateClassroomV1Response) Reset() {
	*x = CreateClassroomV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClassroomV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClassroomV1Response) ProtoMessage() {}

func (x *CreateClassroomV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClassroomV1Response.ProtoReflect.Descriptor instead.
func (*CreateClassroomV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDescGZIP(), []int{5}
}

func (x *CreateClassroomV1Response) GetClassroomId() uint64 {
	if x != nil {
		return x.ClassroomId
	}
	return 0
}

type RemoveClassroomV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassroomId uint64 `protobuf:"varint,1,opt,name=classroom_id,json=classroomId,proto3" json:"classroom_id,omitempty"`
}

func (x *RemoveClassroomV1Request) Reset() {
	*x = RemoveClassroomV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveClassroomV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveClassroomV1Request) ProtoMessage() {}

func (x *RemoveClassroomV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveClassroomV1Request.ProtoReflect.Descriptor instead.
func (*RemoveClassroomV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveClassroomV1Request) GetClassroomId() uint64 {
	if x != nil {
		return x.ClassroomId
	}
	return 0
}

type RemoveClassroomV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Found bool `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
}

func (x *RemoveClassroomV1Response) Reset() {
	*x = RemoveClassroomV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveClassroomV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveClassroomV1Response) ProtoMessage() {}

func (x *RemoveClassroomV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveClassroomV1Response.ProtoReflect.Descriptor instead.
func (*RemoveClassroomV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveClassroomV1Response) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

type Classroom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassroomId uint64 `protobuf:"varint,1,opt,name=classroom_id,json=classroomId,proto3" json:"classroom_id,omitempty"`
	TenantId    uint64 `protobuf:"varint,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	CalendarId  uint64 `protobuf:"varint,3,opt,name=calendar_id,json=calendarId,proto3" json:"calendar_id,omitempty"`
}

func (x *Classroom) Reset() {
	*x = Classroom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Classroom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Classroom) ProtoMessage() {}

func (x *Classroom) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Classroom.ProtoReflect.Descriptor instead.
func (*Classroom) Descriptor() ([]byte, []int) {
	return file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDescGZIP(), []int{8}
}

func (x *Classroom) GetClassroomId() uint64 {
	if x != nil {
		return x.ClassroomId
	}
	return 0
}

func (x *Classroom) GetTenantId() uint64 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *Classroom) GetCalendarId() uint64 {
	if x != nil {
		return x.CalendarId
	}
	return 0
}

var File_api_ocp_classroom_api_ocp_classroom_api_proto protoreflect.FileDescriptor

var file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72,
	0x6f, 0x6f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x6f, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x61,
	0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x58, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72,
	0x6f, 0x6f, 0x6d, 0x52, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x22,
	0x48, 0x0a, 0x1a, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x72, 0x6f, 0x6f, 0x6d, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x0c, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x0b, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x59, 0x0a, 0x1b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x56, 0x31,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6f, 0x63,
	0x70, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x6a, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x32, 0x02, 0x20, 0x00, 0x52, 0x0a, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x49, 0x64,
	0x22, 0x3e, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72,
	0x6f, 0x6f, 0x6d, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x22, 0x46, 0x0a, 0x18, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72,
	0x6f, 0x6f, 0x6d, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x0c,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x0b, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x19, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x56, 0x31, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x87, 0x01, 0x0a, 0x09,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x2a, 0x0a, 0x0c, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72,
	0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20,
	0x00, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0b, 0x63,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x0a, 0x63, 0x61, 0x6c, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x49, 0x64, 0x32, 0xd6, 0x04, 0x0a, 0x0f, 0x4f, 0x63, 0x70, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x41, 0x70, 0x69, 0x12, 0x83, 0x01, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x56, 0x31, 0x12, 0x2a,
	0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d,
	0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6f, 0x63, 0x70,
	0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x56, 0x31, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12,
	0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x12,
	0x9b, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x56, 0x31, 0x12, 0x2d, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x56, 0x31, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x56, 0x31, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2f, 0x7b,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x86, 0x01,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f,
	0x6d, 0x56, 0x31, 0x12, 0x2b, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72,
	0x6f, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2c, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x72, 0x6f, 0x6f, 0x6d, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x95, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x56, 0x31, 0x12, 0x2b, 0x2e, 0x6f,
	0x63, 0x70, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d,
	0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6f, 0x63, 0x70, 0x2e,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x56, 0x31, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x2a,
	0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2f,
	0x7b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x4d,
	0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6f,
	0x6e, 0x63, 0x70, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f,
	0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x6f, 0x63, 0x70, 0x5f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDescOnce sync.Once
	file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDescData = file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDesc
)

func file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDescGZIP() []byte {
	file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDescOnce.Do(func() {
		file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDescData)
	})
	return file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDescData
}

var file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_ocp_classroom_api_ocp_classroom_api_proto_goTypes = []interface{}{
	(*ListClassroomsV1Request)(nil),     // 0: ocp.classroom.api.ListClassroomsV1Request
	(*ListClassroomsV1Response)(nil),    // 1: ocp.classroom.api.ListClassroomsV1Response
	(*DescribeClassroomV1Request)(nil),  // 2: ocp.classroom.api.DescribeClassroomV1Request
	(*DescribeClassroomV1Response)(nil), // 3: ocp.classroom.api.DescribeClassroomV1Response
	(*CreateClassroomV1Request)(nil),    // 4: ocp.classroom.api.CreateClassroomV1Request
	(*CreateClassroomV1Response)(nil),   // 5: ocp.classroom.api.CreateClassroomV1Response
	(*RemoveClassroomV1Request)(nil),    // 6: ocp.classroom.api.RemoveClassroomV1Request
	(*RemoveClassroomV1Response)(nil),   // 7: ocp.classroom.api.RemoveClassroomV1Response
	(*Classroom)(nil),                   // 8: ocp.classroom.api.Classroom
}
var file_api_ocp_classroom_api_ocp_classroom_api_proto_depIdxs = []int32{
	8, // 0: ocp.classroom.api.ListClassroomsV1Response.classrooms:type_name -> ocp.classroom.api.Classroom
	8, // 1: ocp.classroom.api.DescribeClassroomV1Response.classroom:type_name -> ocp.classroom.api.Classroom
	0, // 2: ocp.classroom.api.OcpClassroomApi.ListClassroomsV1:input_type -> ocp.classroom.api.ListClassroomsV1Request
	2, // 3: ocp.classroom.api.OcpClassroomApi.DescribeClassroomV1:input_type -> ocp.classroom.api.DescribeClassroomV1Request
	4, // 4: ocp.classroom.api.OcpClassroomApi.CreateClassroomV1:input_type -> ocp.classroom.api.CreateClassroomV1Request
	6, // 5: ocp.classroom.api.OcpClassroomApi.RemoveClassroomV1:input_type -> ocp.classroom.api.RemoveClassroomV1Request
	1, // 6: ocp.classroom.api.OcpClassroomApi.ListClassroomsV1:output_type -> ocp.classroom.api.ListClassroomsV1Response
	3, // 7: ocp.classroom.api.OcpClassroomApi.DescribeClassroomV1:output_type -> ocp.classroom.api.DescribeClassroomV1Response
	5, // 8: ocp.classroom.api.OcpClassroomApi.CreateClassroomV1:output_type -> ocp.classroom.api.CreateClassroomV1Response
	7, // 9: ocp.classroom.api.OcpClassroomApi.RemoveClassroomV1:output_type -> ocp.classroom.api.RemoveClassroomV1Response
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_ocp_classroom_api_ocp_classroom_api_proto_init() }
func file_api_ocp_classroom_api_ocp_classroom_api_proto_init() {
	if File_api_ocp_classroom_api_ocp_classroom_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClassroomsV1Request); i {
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
		file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClassroomsV1Response); i {
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
		file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeClassroomV1Request); i {
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
		file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeClassroomV1Response); i {
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
		file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClassroomV1Request); i {
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
		file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClassroomV1Response); i {
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
		file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveClassroomV1Request); i {
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
		file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveClassroomV1Response); i {
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
		file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Classroom); i {
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
			RawDescriptor: file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ocp_classroom_api_ocp_classroom_api_proto_goTypes,
		DependencyIndexes: file_api_ocp_classroom_api_ocp_classroom_api_proto_depIdxs,
		MessageInfos:      file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes,
	}.Build()
	File_api_ocp_classroom_api_ocp_classroom_api_proto = out.File
	file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDesc = nil
	file_api_ocp_classroom_api_ocp_classroom_api_proto_goTypes = nil
	file_api_ocp_classroom_api_ocp_classroom_api_proto_depIdxs = nil
}
