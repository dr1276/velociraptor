// Code generated by protoc-gen-go. DO NOT EDIT.
// source: artifacts.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto1 "www.velocidex.com/golang/velociraptor/actions/proto"
import _ "www.velocidex.com/golang/velociraptor/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetArtifactsRequest struct {
	// Deprecated
	IncludeEventArtifacts  bool     `protobuf:"varint,1,opt,name=include_event_artifacts,json=includeEventArtifacts,proto3" json:"include_event_artifacts,omitempty"`
	IncludeServerArtifacts bool     `protobuf:"varint,2,opt,name=include_server_artifacts,json=includeServerArtifacts,proto3" json:"include_server_artifacts,omitempty"`
	SearchTerm             string   `protobuf:"bytes,3,opt,name=search_term,json=searchTerm,proto3" json:"search_term,omitempty"`
	NumberOfResults        uint64   `protobuf:"varint,4,opt,name=number_of_results,json=numberOfResults,proto3" json:"number_of_results,omitempty"`
	Type                   string   `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Names                  []string `protobuf:"bytes,6,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *GetArtifactsRequest) Reset()         { *m = GetArtifactsRequest{} }
func (m *GetArtifactsRequest) String() string { return proto.CompactTextString(m) }
func (*GetArtifactsRequest) ProtoMessage()    {}
func (*GetArtifactsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_473c0b5cbda7b8e2, []int{0}
}
func (m *GetArtifactsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArtifactsRequest.Unmarshal(m, b)
}
func (m *GetArtifactsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArtifactsRequest.Marshal(b, m, deterministic)
}
func (dst *GetArtifactsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArtifactsRequest.Merge(dst, src)
}
func (m *GetArtifactsRequest) XXX_Size() int {
	return xxx_messageInfo_GetArtifactsRequest.Size(m)
}
func (m *GetArtifactsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArtifactsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetArtifactsRequest proto.InternalMessageInfo

func (m *GetArtifactsRequest) GetIncludeEventArtifacts() bool {
	if m != nil {
		return m.IncludeEventArtifacts
	}
	return false
}

func (m *GetArtifactsRequest) GetIncludeServerArtifacts() bool {
	if m != nil {
		return m.IncludeServerArtifacts
	}
	return false
}

func (m *GetArtifactsRequest) GetSearchTerm() string {
	if m != nil {
		return m.SearchTerm
	}
	return ""
}

func (m *GetArtifactsRequest) GetNumberOfResults() uint64 {
	if m != nil {
		return m.NumberOfResults
	}
	return 0
}

func (m *GetArtifactsRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetArtifactsRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type GetArtifactRequest struct {
	// Deprecated.
	VfsPath              string   `protobuf:"bytes,1,opt,name=vfs_path,json=vfsPath,proto3" json:"vfs_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetArtifactRequest) Reset()         { *m = GetArtifactRequest{} }
func (m *GetArtifactRequest) String() string { return proto.CompactTextString(m) }
func (*GetArtifactRequest) ProtoMessage()    {}
func (*GetArtifactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_473c0b5cbda7b8e2, []int{1}
}
func (m *GetArtifactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArtifactRequest.Unmarshal(m, b)
}
func (m *GetArtifactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArtifactRequest.Marshal(b, m, deterministic)
}
func (dst *GetArtifactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArtifactRequest.Merge(dst, src)
}
func (m *GetArtifactRequest) XXX_Size() int {
	return xxx_messageInfo_GetArtifactRequest.Size(m)
}
func (m *GetArtifactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArtifactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetArtifactRequest proto.InternalMessageInfo

func (m *GetArtifactRequest) GetVfsPath() string {
	if m != nil {
		return m.VfsPath
	}
	return ""
}

type GetArtifactResponse struct {
	Artifact             string   `protobuf:"bytes,1,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetArtifactResponse) Reset()         { *m = GetArtifactResponse{} }
func (m *GetArtifactResponse) String() string { return proto.CompactTextString(m) }
func (*GetArtifactResponse) ProtoMessage()    {}
func (*GetArtifactResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_473c0b5cbda7b8e2, []int{2}
}
func (m *GetArtifactResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArtifactResponse.Unmarshal(m, b)
}
func (m *GetArtifactResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArtifactResponse.Marshal(b, m, deterministic)
}
func (dst *GetArtifactResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArtifactResponse.Merge(dst, src)
}
func (m *GetArtifactResponse) XXX_Size() int {
	return xxx_messageInfo_GetArtifactResponse.Size(m)
}
func (m *GetArtifactResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArtifactResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetArtifactResponse proto.InternalMessageInfo

func (m *GetArtifactResponse) GetArtifact() string {
	if m != nil {
		return m.Artifact
	}
	return ""
}

type SetArtifactRequest struct {
	VfsPath              string   `protobuf:"bytes,1,opt,name=vfs_path,json=vfsPath,proto3" json:"vfs_path,omitempty"`
	Artifact             string   `protobuf:"bytes,2,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetArtifactRequest) Reset()         { *m = SetArtifactRequest{} }
func (m *SetArtifactRequest) String() string { return proto.CompactTextString(m) }
func (*SetArtifactRequest) ProtoMessage()    {}
func (*SetArtifactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_473c0b5cbda7b8e2, []int{3}
}
func (m *SetArtifactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetArtifactRequest.Unmarshal(m, b)
}
func (m *SetArtifactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetArtifactRequest.Marshal(b, m, deterministic)
}
func (dst *SetArtifactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetArtifactRequest.Merge(dst, src)
}
func (m *SetArtifactRequest) XXX_Size() int {
	return xxx_messageInfo_SetArtifactRequest.Size(m)
}
func (m *SetArtifactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetArtifactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetArtifactRequest proto.InternalMessageInfo

func (m *SetArtifactRequest) GetVfsPath() string {
	if m != nil {
		return m.VfsPath
	}
	return ""
}

func (m *SetArtifactRequest) GetArtifact() string {
	if m != nil {
		return m.Artifact
	}
	return ""
}

type APIResponse struct {
	Error                bool     `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *APIResponse) Reset()         { *m = APIResponse{} }
func (m *APIResponse) String() string { return proto.CompactTextString(m) }
func (*APIResponse) ProtoMessage()    {}
func (*APIResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_473c0b5cbda7b8e2, []int{4}
}
func (m *APIResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIResponse.Unmarshal(m, b)
}
func (m *APIResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIResponse.Marshal(b, m, deterministic)
}
func (dst *APIResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIResponse.Merge(dst, src)
}
func (m *APIResponse) XXX_Size() int {
	return xxx_messageInfo_APIResponse.Size(m)
}
func (m *APIResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_APIResponse.DiscardUnknown(m)
}

var xxx_messageInfo_APIResponse proto.InternalMessageInfo

func (m *APIResponse) GetError() bool {
	if m != nil {
		return m.Error
	}
	return false
}

func (m *APIResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type GetReportRequest struct {
	Artifact string `protobuf:"bytes,1,opt,name=artifact,proto3" json:"artifact,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Format   string `protobuf:"bytes,3,opt,name=format,proto3" json:"format,omitempty"`
	// Common parameters
	ClientId string `protobuf:"bytes,5,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// Parameters for MONITORING_DAILY
	DayName   string `protobuf:"bytes,6,opt,name=day_name,json=dayName,proto3" json:"day_name,omitempty"`
	StartTime uint64 `protobuf:"varint,8,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   uint64 `protobuf:"varint,9,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Parameters for CLIENT
	FlowId               string           `protobuf:"bytes,7,opt,name=flow_id,json=flowId,proto3" json:"flow_id,omitempty"`
	Parameters           []*proto1.VQLEnv `protobuf:"bytes,4,rep,name=parameters,proto3" json:"parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetReportRequest) Reset()         { *m = GetReportRequest{} }
func (m *GetReportRequest) String() string { return proto.CompactTextString(m) }
func (*GetReportRequest) ProtoMessage()    {}
func (*GetReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_473c0b5cbda7b8e2, []int{5}
}
func (m *GetReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReportRequest.Unmarshal(m, b)
}
func (m *GetReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReportRequest.Marshal(b, m, deterministic)
}
func (dst *GetReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReportRequest.Merge(dst, src)
}
func (m *GetReportRequest) XXX_Size() int {
	return xxx_messageInfo_GetReportRequest.Size(m)
}
func (m *GetReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetReportRequest proto.InternalMessageInfo

func (m *GetReportRequest) GetArtifact() string {
	if m != nil {
		return m.Artifact
	}
	return ""
}

func (m *GetReportRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetReportRequest) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *GetReportRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *GetReportRequest) GetDayName() string {
	if m != nil {
		return m.DayName
	}
	return ""
}

func (m *GetReportRequest) GetStartTime() uint64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *GetReportRequest) GetEndTime() uint64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *GetReportRequest) GetFlowId() string {
	if m != nil {
		return m.FlowId
	}
	return ""
}

func (m *GetReportRequest) GetParameters() []*proto1.VQLEnv {
	if m != nil {
		return m.Parameters
	}
	return nil
}

// This presents the report in a form that can be rendered in the
// GUI. The data is presented in two parts - first "data" contains a
// json encoded object, then "template" is an angular template to be
// evaluated with the data.
type GetReportResponse struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Template             string   `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	Messages             []string `protobuf:"bytes,3,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReportResponse) Reset()         { *m = GetReportResponse{} }
func (m *GetReportResponse) String() string { return proto.CompactTextString(m) }
func (*GetReportResponse) ProtoMessage()    {}
func (*GetReportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_473c0b5cbda7b8e2, []int{6}
}
func (m *GetReportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReportResponse.Unmarshal(m, b)
}
func (m *GetReportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReportResponse.Marshal(b, m, deterministic)
}
func (dst *GetReportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReportResponse.Merge(dst, src)
}
func (m *GetReportResponse) XXX_Size() int {
	return xxx_messageInfo_GetReportResponse.Size(m)
}
func (m *GetReportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetReportResponse proto.InternalMessageInfo

func (m *GetReportResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *GetReportResponse) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *GetReportResponse) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

// Deprecated.
type ArtifactCompressionDict struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtifactCompressionDict) Reset()         { *m = ArtifactCompressionDict{} }
func (m *ArtifactCompressionDict) String() string { return proto.CompactTextString(m) }
func (*ArtifactCompressionDict) ProtoMessage()    {}
func (*ArtifactCompressionDict) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_473c0b5cbda7b8e2, []int{7}
}
func (m *ArtifactCompressionDict) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactCompressionDict.Unmarshal(m, b)
}
func (m *ArtifactCompressionDict) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactCompressionDict.Marshal(b, m, deterministic)
}
func (dst *ArtifactCompressionDict) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactCompressionDict.Merge(dst, src)
}
func (m *ArtifactCompressionDict) XXX_Size() int {
	return xxx_messageInfo_ArtifactCompressionDict.Size(m)
}
func (m *ArtifactCompressionDict) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactCompressionDict.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactCompressionDict proto.InternalMessageInfo

type ListAvailableEventResultsRequest struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAvailableEventResultsRequest) Reset()         { *m = ListAvailableEventResultsRequest{} }
func (m *ListAvailableEventResultsRequest) String() string { return proto.CompactTextString(m) }
func (*ListAvailableEventResultsRequest) ProtoMessage()    {}
func (*ListAvailableEventResultsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_473c0b5cbda7b8e2, []int{8}
}
func (m *ListAvailableEventResultsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAvailableEventResultsRequest.Unmarshal(m, b)
}
func (m *ListAvailableEventResultsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAvailableEventResultsRequest.Marshal(b, m, deterministic)
}
func (dst *ListAvailableEventResultsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAvailableEventResultsRequest.Merge(dst, src)
}
func (m *ListAvailableEventResultsRequest) XXX_Size() int {
	return xxx_messageInfo_ListAvailableEventResultsRequest.Size(m)
}
func (m *ListAvailableEventResultsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAvailableEventResultsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAvailableEventResultsRequest proto.InternalMessageInfo

func (m *ListAvailableEventResultsRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type AvailableEvent struct {
	Artifact             string   `protobuf:"bytes,1,opt,name=artifact,proto3" json:"artifact,omitempty"`
	Timestamps           []int32  `protobuf:"varint,2,rep,packed,name=timestamps,proto3" json:"timestamps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AvailableEvent) Reset()         { *m = AvailableEvent{} }
func (m *AvailableEvent) String() string { return proto.CompactTextString(m) }
func (*AvailableEvent) ProtoMessage()    {}
func (*AvailableEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_473c0b5cbda7b8e2, []int{9}
}
func (m *AvailableEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvailableEvent.Unmarshal(m, b)
}
func (m *AvailableEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvailableEvent.Marshal(b, m, deterministic)
}
func (dst *AvailableEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvailableEvent.Merge(dst, src)
}
func (m *AvailableEvent) XXX_Size() int {
	return xxx_messageInfo_AvailableEvent.Size(m)
}
func (m *AvailableEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AvailableEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AvailableEvent proto.InternalMessageInfo

func (m *AvailableEvent) GetArtifact() string {
	if m != nil {
		return m.Artifact
	}
	return ""
}

func (m *AvailableEvent) GetTimestamps() []int32 {
	if m != nil {
		return m.Timestamps
	}
	return nil
}

type ListAvailableEventResultsResponse struct {
	Logs                 []*AvailableEvent `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListAvailableEventResultsResponse) Reset()         { *m = ListAvailableEventResultsResponse{} }
func (m *ListAvailableEventResultsResponse) String() string { return proto.CompactTextString(m) }
func (*ListAvailableEventResultsResponse) ProtoMessage()    {}
func (*ListAvailableEventResultsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_473c0b5cbda7b8e2, []int{10}
}
func (m *ListAvailableEventResultsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAvailableEventResultsResponse.Unmarshal(m, b)
}
func (m *ListAvailableEventResultsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAvailableEventResultsResponse.Marshal(b, m, deterministic)
}
func (dst *ListAvailableEventResultsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAvailableEventResultsResponse.Merge(dst, src)
}
func (m *ListAvailableEventResultsResponse) XXX_Size() int {
	return xxx_messageInfo_ListAvailableEventResultsResponse.Size(m)
}
func (m *ListAvailableEventResultsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAvailableEventResultsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAvailableEventResultsResponse proto.InternalMessageInfo

func (m *ListAvailableEventResultsResponse) GetLogs() []*AvailableEvent {
	if m != nil {
		return m.Logs
	}
	return nil
}

func init() {
	proto.RegisterType((*GetArtifactsRequest)(nil), "proto.GetArtifactsRequest")
	proto.RegisterType((*GetArtifactRequest)(nil), "proto.GetArtifactRequest")
	proto.RegisterType((*GetArtifactResponse)(nil), "proto.GetArtifactResponse")
	proto.RegisterType((*SetArtifactRequest)(nil), "proto.SetArtifactRequest")
	proto.RegisterType((*APIResponse)(nil), "proto.APIResponse")
	proto.RegisterType((*GetReportRequest)(nil), "proto.GetReportRequest")
	proto.RegisterType((*GetReportResponse)(nil), "proto.GetReportResponse")
	proto.RegisterType((*ArtifactCompressionDict)(nil), "proto.ArtifactCompressionDict")
	proto.RegisterType((*ListAvailableEventResultsRequest)(nil), "proto.ListAvailableEventResultsRequest")
	proto.RegisterType((*AvailableEvent)(nil), "proto.AvailableEvent")
	proto.RegisterType((*ListAvailableEventResultsResponse)(nil), "proto.ListAvailableEventResultsResponse")
}

func init() { proto.RegisterFile("artifacts.proto", fileDescriptor_artifacts_473c0b5cbda7b8e2) }

var fileDescriptor_artifacts_473c0b5cbda7b8e2 = []byte{
	// 946 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xdd, 0x6e, 0xdc, 0x44,
	0x14, 0x96, 0xb3, 0xd9, 0xbf, 0x09, 0xa5, 0xed, 0xa0, 0x12, 0x37, 0x94, 0x32, 0x18, 0x04, 0x5b,
	0x84, 0xbc, 0xfc, 0x48, 0x50, 0x45, 0x80, 0xd8, 0x25, 0x21, 0x5a, 0x91, 0x26, 0xc5, 0x5d, 0x21,
	0xa1, 0x5e, 0x58, 0x13, 0xfb, 0x78, 0x77, 0x84, 0x3d, 0xe3, 0xcc, 0xcc, 0x7a, 0x9b, 0x6b, 0x6e,
	0x78, 0x10, 0x6e, 0xb9, 0xe6, 0x05, 0x78, 0x05, 0x5e, 0x00, 0x5e, 0x83, 0x0b, 0x34, 0x33, 0xf6,
	0xc6, 0x0b, 0x12, 0xaa, 0xb8, 0xe8, 0x55, 0x66, 0xcf, 0x77, 0xbe, 0x6f, 0x8e, 0xcf, 0xf9, 0xce,
	0x04, 0xdd, 0xa4, 0x52, 0xb3, 0x8c, 0x26, 0x5a, 0x85, 0xa5, 0x14, 0x5a, 0xe0, 0xae, 0xfd, 0x73,
	0x70, 0xb8, 0x5e, 0xaf, 0xc3, 0x0a, 0x72, 0x91, 0xb0, 0x14, 0x9e, 0x85, 0x89, 0x28, 0xc6, 0x0b,
	0x91, 0x53, 0xbe, 0x18, 0xbb, 0xa0, 0xa4, 0xa5, 0x16, 0x72, 0x6c, 0x93, 0xc7, 0x0a, 0x0a, 0xca,
	0x35, 0x4b, 0x9c, 0xc4, 0xc1, 0xe7, 0xcf, 0xc7, 0xa5, 0x89, 0x66, 0x82, 0xab, 0x5a, 0xa3, 0xba,
	0xcc, 0x1d, 0x3d, 0xf8, 0x79, 0x07, 0xbd, 0x72, 0x02, 0x7a, 0xd2, 0x14, 0x16, 0xc1, 0xe5, 0x0a,
	0x94, 0xc6, 0x9f, 0xa0, 0x7d, 0xc6, 0x93, 0x7c, 0x95, 0x42, 0x0c, 0x15, 0x70, 0x1d, 0x6f, 0x4a,
	0xf7, 0x3d, 0xe2, 0x8d, 0x06, 0xd1, 0x9d, 0x1a, 0x3e, 0x36, 0xe8, 0x86, 0x8e, 0x1f, 0x22, 0xbf,
	0xe1, 0x29, 0x90, 0x15, 0xc8, 0x16, 0x71, 0xc7, 0x12, 0x5f, 0xad, 0xf1, 0x27, 0x16, 0xbe, 0x66,
	0xbe, 0x81, 0xf6, 0x14, 0x50, 0x99, 0x2c, 0x63, 0x0d, 0xb2, 0xf0, 0x3b, 0xc4, 0x1b, 0x0d, 0x23,
	0xe4, 0x42, 0x73, 0x90, 0x05, 0x7e, 0x0f, 0xdd, 0xe6, 0xab, 0xe2, 0x02, 0x64, 0x2c, 0xb2, 0x58,
	0x82, 0x5a, 0xe5, 0x5a, 0xf9, 0xbb, 0xc4, 0x1b, 0xed, 0x46, 0x37, 0x1d, 0x70, 0x9e, 0x45, 0x2e,
	0x8c, 0x31, 0xda, 0xd5, 0x57, 0x25, 0xf8, 0x5d, 0xab, 0x62, 0xcf, 0xf8, 0x10, 0x75, 0x39, 0x2d,
	0x40, 0xf9, 0x3d, 0xd2, 0x19, 0x0d, 0xa7, 0x6f, 0xff, 0xf1, 0xd7, 0x9f, 0xbf, 0x79, 0xf7, 0xf1,
	0xbd, 0x09, 0xc9, 0x99, 0xd2, 0x44, 0x64, 0x64, 0x53, 0x2a, 0xd1, 0x82, 0x64, 0xa0, 0x93, 0x65,
	0xe4, 0x28, 0xc1, 0x25, 0xc2, 0xad, 0x2e, 0x35, 0x4d, 0x7a, 0x8a, 0x06, 0x55, 0xa6, 0xe2, 0x92,
	0xea, 0xa5, 0xed, 0xca, 0x70, 0xfa, 0xa5, 0x15, 0x3d, 0xc4, 0x0f, 0xe7, 0x4b, 0x20, 0x55, 0xa6,
	0x88, 0xc1, 0x88, 0x84, 0x9c, 0x6a, 0x56, 0x81, 0x91, 0xd5, 0x4b, 0x68, 0xdd, 0x93, 0x42, 0xc6,
	0x38, 0x33, 0xa3, 0x21, 0x4a, 0x0b, 0x09, 0x61, 0xd4, 0xaf, 0x32, 0xf5, 0x98, 0xea, 0x65, 0xf0,
	0x74, 0x6b, 0x30, 0x11, 0xa8, 0x52, 0x70, 0x05, 0xf8, 0x08, 0x0d, 0x1a, 0x7a, 0x7d, 0xe7, 0xc8,
	0xde, 0x19, 0x60, 0x32, 0x6f, 0x49, 0x93, 0x94, 0x6a, 0xfa, 0x3e, 0x11, 0x92, 0x50, 0x73, 0x09,
	0x5d, 0xe5, 0x3a, 0x8c, 0x36, 0xcc, 0xe0, 0x57, 0x0f, 0xe1, 0x27, 0x2f, 0xf6, 0x83, 0xb6, 0x2a,
	0xdf, 0xf9, 0xdf, 0x95, 0x3f, 0x43, 0x7b, 0x93, 0xc7, 0xb3, 0x56, 0x3b, 0xba, 0x20, 0xa5, 0x90,
	0xce, 0x95, 0xd3, 0xd0, 0x2a, 0x8e, 0xf0, 0x3b, 0x13, 0x4e, 0x6c, 0x9c, 0x88, 0x24, 0x59, 0x49,
	0x48, 0x89, 0x02, 0xad, 0x19, 0x5f, 0x6c, 0x95, 0x1b, 0x46, 0x8e, 0x8c, 0xdf, 0x42, 0x37, 0xec,
	0x21, 0x2e, 0x40, 0x29, 0xba, 0x00, 0x57, 0x5f, 0xf4, 0x92, 0x0d, 0x3e, 0x72, 0xb1, 0xe0, 0xf7,
	0x0e, 0xba, 0x75, 0x02, 0x3a, 0x82, 0x52, 0xc8, 0x4d, 0xc7, 0x9e, 0x77, 0x1c, 0x99, 0x90, 0x64,
	0xbd, 0x64, 0xc9, 0x92, 0xac, 0x81, 0x48, 0x27, 0xb1, 0x61, 0xe2, 0x6f, 0x6a, 0xbb, 0xba, 0xb6,
	0x7c, 0x6a, 0x15, 0x3e, 0xc4, 0x63, 0xa3, 0xe0, 0x72, 0x89, 0x81, 0x0d, 0x95, 0x03, 0xa4, 0x64,
	0x04, 0xe1, 0x22, 0x24, 0x8f, 0xce, 0xcf, 0x66, 0xf3, 0xf3, 0x68, 0x76, 0x76, 0x12, 0x1f, 0x4d,
	0x66, 0xa7, 0xdf, 0x3f, 0xa8, 0x7d, 0xfe, 0x01, 0xea, 0x65, 0x42, 0x16, 0x54, 0xbb, 0x1d, 0x9a,
	0xfa, 0x56, 0x0e, 0xe3, 0x5b, 0x5f, 0xdb, 0x28, 0xb1, 0xe4, 0xa5, 0x2e, 0xf2, 0xa8, 0xce, 0xc3,
	0xaf, 0xa1, 0x61, 0x92, 0x33, 0xb3, 0xe5, 0x2c, 0xad, 0x57, 0x66, 0xe0, 0x02, 0xb3, 0x14, 0xdf,
	0x45, 0x83, 0x94, 0x5e, 0xc5, 0x66, 0x0f, 0xfc, 0x9e, 0xc5, 0xfa, 0x29, 0xbd, 0x3a, 0xa3, 0x05,
	0xe0, 0xd7, 0x11, 0x52, 0x9a, 0x4a, 0x1d, 0x6b, 0x56, 0x80, 0x3f, 0xb0, 0xab, 0x38, 0xb4, 0x91,
	0x39, 0x2b, 0xc0, 0x30, 0x81, 0xa7, 0x0e, 0x1c, 0x5a, 0xb0, 0x0f, 0x3c, 0xb5, 0xd0, 0x3e, 0xea,
	0x67, 0xb9, 0x58, 0x9b, 0xfb, 0xfa, 0x56, 0xb3, 0x67, 0x7e, 0xce, 0x52, 0x7c, 0x89, 0x50, 0x49,
	0x25, 0x2d, 0x40, 0x83, 0x34, 0xdb, 0xdd, 0x19, 0xed, 0x7d, 0x74, 0xc3, 0xbd, 0x55, 0xe1, 0x77,
	0xdf, 0x9e, 0x1e, 0xf3, 0x6a, 0x3a, 0xb5, 0xdf, 0xf3, 0x19, 0x3e, 0x74, 0x93, 0x20, 0xd7, 0xf9,
	0xa1, 0xe9, 0xb9, 0x02, 0x92, 0x42, 0x09, 0x3c, 0x25, 0x82, 0xdb, 0x69, 0xdb, 0xd6, 0x89, 0xcc,
	0x9e, 0x5d, 0x37, 0xc3, 0xa8, 0x75, 0x49, 0xf0, 0xa3, 0x87, 0x6e, 0xb7, 0xe6, 0x5a, 0x1b, 0x0b,
	0xa3, 0x5d, 0x63, 0x44, 0x37, 0xd4, 0xc8, 0x9e, 0xf1, 0x01, 0x1a, 0x68, 0x28, 0xca, 0x9c, 0xea,
	0xc6, 0x21, 0x9b, 0xdf, 0xf8, 0x0b, 0x34, 0xa8, 0xcd, 0xa3, 0xfc, 0x8e, 0x7d, 0x60, 0x02, 0x5b,
	0xe7, 0x3d, 0x7c, 0x70, 0xec, 0x8c, 0x28, 0xc9, 0x9a, 0x4a, 0x6e, 0x3c, 0xd8, 0x24, 0x86, 0xd1,
	0x86, 0x13, 0xdc, 0x45, 0xfb, 0xcd, 0x36, 0x7e, 0x25, 0x8a, 0x52, 0x82, 0x52, 0x4c, 0xf0, 0x23,
	0x96, 0xe8, 0xe0, 0x17, 0x0f, 0x91, 0x53, 0xa6, 0xf4, 0xa4, 0xa2, 0x2c, 0xa7, 0x17, 0xb9, 0x7b,
	0x73, 0xeb, 0xa7, 0xae, 0x31, 0xe2, 0x4f, 0x5e, 0x7b, 0x88, 0xce, 0x8a, 0x3f, 0xd8, 0x0a, 0x00,
	0x27, 0xc6, 0x48, 0x0e, 0x24, 0xb3, 0x23, 0x63, 0x23, 0xfb, 0xe2, 0xd9, 0x55, 0x68, 0x34, 0x49,
	0x21, 0x38, 0xd3, 0x42, 0x9a, 0xfa, 0x72, 0xb1, 0x50, 0x21, 0x99, 0x65, 0x04, 0x8a, 0x52, 0x5f,
	0x6d, 0xe5, 0xbb, 0x47, 0xfd, 0x5d, 0xf5, 0xaf, 0xf4, 0x6b, 0xc7, 0x04, 0xa7, 0xe8, 0xe5, 0xed,
	0x52, 0x4d, 0xe3, 0xb6, 0xb7, 0xa4, 0xe5, 0xfd, 0xfb, 0x08, 0x19, 0x87, 0x28, 0x4d, 0x8b, 0xd2,
	0xfc, 0x8f, 0xe8, 0x8c, 0xba, 0x51, 0x2b, 0x12, 0x9c, 0xa1, 0x37, 0xff, 0xe3, 0xe3, 0xeb, 0x69,
	0x3d, 0x40, 0xbb, 0xa6, 0x0a, 0xdf, 0xb3, 0x86, 0xb9, 0x53, 0x1b, 0xe6, 0x1f, 0x1c, 0x9b, 0x72,
	0xd1, 0xb3, 0xd8, 0xc7, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x62, 0x03, 0x3b, 0xf4, 0x8d, 0x07,
	0x00, 0x00,
}
