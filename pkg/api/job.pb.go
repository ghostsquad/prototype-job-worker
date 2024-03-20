// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: pkg/api/job.proto

package prototype_job_worker

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

type JobStartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command *string `protobuf:"bytes,1,req,name=command" json:"command,omitempty"`
}

func (x *JobStartRequest) Reset() {
	*x = JobStartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobStartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStartRequest) ProtoMessage() {}

func (x *JobStartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStartRequest.ProtoReflect.Descriptor instead.
func (*JobStartRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_job_proto_rawDescGZIP(), []int{0}
}

func (x *JobStartRequest) GetCommand() string {
	if x != nil && x.Command != nil {
		return *x.Command
	}
	return ""
}

type JobGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
}

func (x *JobGetRequest) Reset() {
	*x = JobGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobGetRequest) ProtoMessage() {}

func (x *JobGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobGetRequest.ProtoReflect.Descriptor instead.
func (*JobGetRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_job_proto_rawDescGZIP(), []int{1}
}

func (x *JobGetRequest) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

type JobGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Command  *string `protobuf:"bytes,2,req,name=command" json:"command,omitempty"`
	Start    *string `protobuf:"bytes,3,req,name=start" json:"start,omitempty"`
	End      *string `protobuf:"bytes,4,opt,name=end" json:"end,omitempty"`
	ExitCode *int32  `protobuf:"varint,5,opt,name=exit_code,json=exitCode" json:"exit_code,omitempty"`
}

func (x *JobGetResponse) Reset() {
	*x = JobGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_job_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobGetResponse) ProtoMessage() {}

func (x *JobGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_job_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobGetResponse.ProtoReflect.Descriptor instead.
func (*JobGetResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_job_proto_rawDescGZIP(), []int{2}
}

func (x *JobGetResponse) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *JobGetResponse) GetCommand() string {
	if x != nil && x.Command != nil {
		return *x.Command
	}
	return ""
}

func (x *JobGetResponse) GetStart() string {
	if x != nil && x.Start != nil {
		return *x.Start
	}
	return ""
}

func (x *JobGetResponse) GetEnd() string {
	if x != nil && x.End != nil {
		return *x.End
	}
	return ""
}

func (x *JobGetResponse) GetExitCode() int32 {
	if x != nil && x.ExitCode != nil {
		return *x.ExitCode
	}
	return 0
}

type JobListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JobListRequest) Reset() {
	*x = JobListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_job_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobListRequest) ProtoMessage() {}

func (x *JobListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_job_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobListRequest.ProtoReflect.Descriptor instead.
func (*JobListRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_job_proto_rawDescGZIP(), []int{3}
}

type JobListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (x *JobListResponse) Reset() {
	*x = JobListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_job_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobListResponse) ProtoMessage() {}

func (x *JobListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_job_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobListResponse.ProtoReflect.Descriptor instead.
func (*JobListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_job_proto_rawDescGZIP(), []int{4}
}

func (x *JobListResponse) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type JobStopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
}

func (x *JobStopRequest) Reset() {
	*x = JobStopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_job_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobStopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStopRequest) ProtoMessage() {}

func (x *JobStopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_job_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStopRequest.ProtoReflect.Descriptor instead.
func (*JobStopRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_job_proto_rawDescGZIP(), []int{5}
}

func (x *JobStopRequest) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

type JobLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
}

func (x *JobLogsRequest) Reset() {
	*x = JobLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_job_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobLogsRequest) ProtoMessage() {}

func (x *JobLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_job_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobLogsRequest.ProtoReflect.Descriptor instead.
func (*JobLogsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_job_proto_rawDescGZIP(), []int{6}
}

func (x *JobLogsRequest) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

type JobLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Line *string `protobuf:"bytes,1,req,name=line" json:"line,omitempty"`
}

func (x *JobLogsResponse) Reset() {
	*x = JobLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_job_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobLogsResponse) ProtoMessage() {}

func (x *JobLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_job_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobLogsResponse.ProtoReflect.Descriptor instead.
func (*JobLogsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_job_proto_rawDescGZIP(), []int{7}
}

func (x *JobLogsResponse) GetLine() string {
	if x != nil && x.Line != nil {
		return *x.Line
	}
	return ""
}

var File_pkg_api_job_proto protoreflect.FileDescriptor

var file_pkg_api_job_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0f, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x22, 0x1f, 0x0a, 0x0d, 0x4a, 0x6f, 0x62, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x7f, 0x0a, 0x0e, 0x4a, 0x6f, 0x62, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x4a, 0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x23, 0x0a, 0x0f, 0x4a, 0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x4a, 0x6f, 0x62,
	0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x4a,
	0x6f, 0x62, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a,
	0x0f, 0x4a, 0x6f, 0x62, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04,
	0x6c, 0x69, 0x6e, 0x65, 0x32, 0xe8, 0x01, 0x0a, 0x06, 0x4a, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x12,
	0x2c, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x4a, 0x6f, 0x62,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2a, 0x0a,
	0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x0f, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x6f, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x4a, 0x6f, 0x62, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x0e, 0x2e, 0x4a, 0x6f, 0x62, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0f, 0x2e, 0x4a, 0x6f, 0x62, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x4a, 0x6f,
	0x62, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x4a,
	0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x2d, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x0f, 0x2e, 0x4a, 0x6f, 0x62, 0x4c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x4a, 0x6f, 0x62, 0x4c,
	0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x68,
	0x6f, 0x73, 0x74, 0x73, 0x71, 0x75, 0x61, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79,
	0x70, 0x65, 0x2d, 0x6a, 0x6f, 0x62, 0x2d, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
}

var (
	file_pkg_api_job_proto_rawDescOnce sync.Once
	file_pkg_api_job_proto_rawDescData = file_pkg_api_job_proto_rawDesc
)

func file_pkg_api_job_proto_rawDescGZIP() []byte {
	file_pkg_api_job_proto_rawDescOnce.Do(func() {
		file_pkg_api_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_api_job_proto_rawDescData)
	})
	return file_pkg_api_job_proto_rawDescData
}

var file_pkg_api_job_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_api_job_proto_goTypes = []interface{}{
	(*JobStartRequest)(nil), // 0: JobStartRequest
	(*JobGetRequest)(nil),   // 1: JobGetRequest
	(*JobGetResponse)(nil),  // 2: JobGetResponse
	(*JobListRequest)(nil),  // 3: JobListRequest
	(*JobListResponse)(nil), // 4: JobListResponse
	(*JobStopRequest)(nil),  // 5: JobStopRequest
	(*JobLogsRequest)(nil),  // 6: JobLogsRequest
	(*JobLogsResponse)(nil), // 7: JobLogsResponse
}
var file_pkg_api_job_proto_depIdxs = []int32{
	0, // 0: Jobber.Start:input_type -> JobStartRequest
	5, // 1: Jobber.Stop:input_type -> JobStopRequest
	1, // 2: Jobber.Get:input_type -> JobGetRequest
	3, // 3: Jobber.List:input_type -> JobListRequest
	6, // 4: Jobber.Logs:input_type -> JobLogsRequest
	2, // 5: Jobber.Start:output_type -> JobGetResponse
	2, // 6: Jobber.Stop:output_type -> JobGetResponse
	2, // 7: Jobber.Get:output_type -> JobGetResponse
	4, // 8: Jobber.List:output_type -> JobListResponse
	7, // 9: Jobber.Logs:output_type -> JobLogsResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_api_job_proto_init() }
func file_pkg_api_job_proto_init() {
	if File_pkg_api_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_api_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobStartRequest); i {
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
		file_pkg_api_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobGetRequest); i {
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
		file_pkg_api_job_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobGetResponse); i {
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
		file_pkg_api_job_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobListRequest); i {
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
		file_pkg_api_job_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobListResponse); i {
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
		file_pkg_api_job_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobStopRequest); i {
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
		file_pkg_api_job_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobLogsRequest); i {
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
		file_pkg_api_job_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobLogsResponse); i {
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
			RawDescriptor: file_pkg_api_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_job_proto_goTypes,
		DependencyIndexes: file_pkg_api_job_proto_depIdxs,
		MessageInfos:      file_pkg_api_job_proto_msgTypes,
	}.Build()
	File_pkg_api_job_proto = out.File
	file_pkg_api_job_proto_rawDesc = nil
	file_pkg_api_job_proto_goTypes = nil
	file_pkg_api_job_proto_depIdxs = nil
}
