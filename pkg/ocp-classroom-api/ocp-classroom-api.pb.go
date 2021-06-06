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

type DescribeClassroomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassroomId uint64 `protobuf:"varint,1,opt,name=classroom_id,json=classroomId,proto3" json:"classroom_id,omitempty"`
	Verbose     bool   `protobuf:"varint,2,opt,name=verbose,proto3" json:"verbose,omitempty"`
}

func (x *DescribeClassroomRequest) Reset() {
	*x = DescribeClassroomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeClassroomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeClassroomRequest) ProtoMessage() {}

func (x *DescribeClassroomRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DescribeClassroomRequest.ProtoReflect.Descriptor instead.
func (*DescribeClassroomRequest) Descriptor() ([]byte, []int) {
	return file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDescGZIP(), []int{0}
}

func (x *DescribeClassroomRequest) GetClassroomId() uint64 {
	if x != nil {
		return x.ClassroomId
	}
	return 0
}

func (x *DescribeClassroomRequest) GetVerbose() bool {
	if x != nil {
		return x.Verbose
	}
	return false
}

type DescribeClassroomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Classroom *Classroom `protobuf:"bytes,1,opt,name=classroom,proto3" json:"classroom,omitempty"`
}

func (x *DescribeClassroomResponse) Reset() {
	*x = DescribeClassroomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeClassroomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeClassroomResponse) ProtoMessage() {}

func (x *DescribeClassroomResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DescribeClassroomResponse.ProtoReflect.Descriptor instead.
func (*DescribeClassroomResponse) Descriptor() ([]byte, []int) {
	return file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeClassroomResponse) GetClassroom() *Classroom {
	if x != nil {
		return x.Classroom
	}
	return nil
}

type Classroom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassroomId uint64 `protobuf:"varint,1,opt,name=classroom_id,json=classroomId,proto3" json:"classroom_id,omitempty"`
	TenantId    uint64 `protobuf:"varint,2,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	CalendarId  uint64 `protobuf:"varint,3,opt,name=calendarId,proto3" json:"calendarId,omitempty"`
}

func (x *Classroom) Reset() {
	*x = Classroom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Classroom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Classroom) ProtoMessage() {}

func (x *Classroom) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Classroom.ProtoReflect.Descriptor instead.
func (*Classroom) Descriptor() ([]byte, []int) {
	return file_api_ocp_classroom_api_ocp_classroom_api_proto_rawDescGZIP(), []int{2}
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
	0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x18, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x22, 0x57, 0x0a, 0x19,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6f,
	0x63, 0x70, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x6a, 0x0a, 0x09, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f,
	0x6f, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72,
	0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x49,
	0x64, 0x32, 0xa6, 0x01, 0x0a, 0x0f, 0x4f, 0x63, 0x70, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f,
	0x6f, 0x6d, 0x41, 0x70, 0x69, 0x12, 0x92, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x2b, 0x2e, 0x6f, 0x63,
	0x70, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x1a,
	0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x63, 0x70, 0x2f,
	0x6f, 0x63, 0x70, 0x2d, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72,
	0x6f, 0x6f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x6f, 0x63, 0x70, 0x5f, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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

var file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_ocp_classroom_api_ocp_classroom_api_proto_goTypes = []interface{}{
	(*DescribeClassroomRequest)(nil),  // 0: ocp.classroom.api.DescribeClassroomRequest
	(*DescribeClassroomResponse)(nil), // 1: ocp.classroom.api.DescribeClassroomResponse
	(*Classroom)(nil),                 // 2: ocp.classroom.api.Classroom
}
var file_api_ocp_classroom_api_ocp_classroom_api_proto_depIdxs = []int32{
	2, // 0: ocp.classroom.api.DescribeClassroomResponse.classroom:type_name -> ocp.classroom.api.Classroom
	0, // 1: ocp.classroom.api.OcpClassroomApi.DescribeClassroom:input_type -> ocp.classroom.api.DescribeClassroomRequest
	1, // 2: ocp.classroom.api.OcpClassroomApi.DescribeClassroom:output_type -> ocp.classroom.api.DescribeClassroomResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_ocp_classroom_api_ocp_classroom_api_proto_init() }
func file_api_ocp_classroom_api_ocp_classroom_api_proto_init() {
	if File_api_ocp_classroom_api_ocp_classroom_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ocp_classroom_api_ocp_classroom_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeClassroomRequest); i {
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
			switch v := v.(*DescribeClassroomResponse); i {
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
			NumMessages:   3,
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
