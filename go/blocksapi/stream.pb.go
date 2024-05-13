// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v3.21.12
// source: blocksapi/stream.proto

package blocksapi

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

// Defines initial stream seek behavior for given target message ID
type GetBlockStreamRequest_StartPolicy int32

const (
	// Start on earliest available message
	GetBlockStreamRequest_ON_EARLIEST_AVAILABLE GetBlockStreamRequest_StartPolicy = 0
	// Start on latest available message
	GetBlockStreamRequest_ON_LATEST_AVAILABLE GetBlockStreamRequest_StartPolicy = 1
	// Start exactly on target, return error if no such target
	GetBlockStreamRequest_EXACTLY_ON_TARGET GetBlockStreamRequest_StartPolicy = 2
	// Start on message which comes exactly after target, return error if no such target
	GetBlockStreamRequest_EXACTLY_AFTER_TARGET GetBlockStreamRequest_StartPolicy = 3
	// Start on earliest available message that is greater or equal to target
	GetBlockStreamRequest_ON_CLOSEST_TO_TARGET GetBlockStreamRequest_StartPolicy = 4
	// Start on earliest available message that is strictly greater than target
	GetBlockStreamRequest_ON_EARLIEST_AFTER_TARGET GetBlockStreamRequest_StartPolicy = 5
)

// Enum value maps for GetBlockStreamRequest_StartPolicy.
var (
	GetBlockStreamRequest_StartPolicy_name = map[int32]string{
		0: "ON_EARLIEST_AVAILABLE",
		1: "ON_LATEST_AVAILABLE",
		2: "EXACTLY_ON_TARGET",
		3: "EXACTLY_AFTER_TARGET",
		4: "ON_CLOSEST_TO_TARGET",
		5: "ON_EARLIEST_AFTER_TARGET",
	}
	GetBlockStreamRequest_StartPolicy_value = map[string]int32{
		"ON_EARLIEST_AVAILABLE":    0,
		"ON_LATEST_AVAILABLE":      1,
		"EXACTLY_ON_TARGET":        2,
		"EXACTLY_AFTER_TARGET":     3,
		"ON_CLOSEST_TO_TARGET":     4,
		"ON_EARLIEST_AFTER_TARGET": 5,
	}
)

func (x GetBlockStreamRequest_StartPolicy) Enum() *GetBlockStreamRequest_StartPolicy {
	p := new(GetBlockStreamRequest_StartPolicy)
	*p = x
	return p
}

func (x GetBlockStreamRequest_StartPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetBlockStreamRequest_StartPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_blocksapi_stream_proto_enumTypes[0].Descriptor()
}

func (GetBlockStreamRequest_StartPolicy) Type() protoreflect.EnumType {
	return &file_blocksapi_stream_proto_enumTypes[0]
}

func (x GetBlockStreamRequest_StartPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetBlockStreamRequest_StartPolicy.Descriptor instead.
func (GetBlockStreamRequest_StartPolicy) EnumDescriptor() ([]byte, []int) {
	return file_blocksapi_stream_proto_rawDescGZIP(), []int{2, 0}
}

// Defines how service should behave if start target is not yet available
type GetBlockStreamRequest_CatchupPolicy int32

const (
	// Return error if catch up needed
	GetBlockStreamRequest_PANIC GetBlockStreamRequest_CatchupPolicy = 0
	// Don't send anything until catch up
	GetBlockStreamRequest_WAIT GetBlockStreamRequest_CatchupPolicy = 1
	// Stream normally from whatever is available before start target
	GetBlockStreamRequest_STREAM GetBlockStreamRequest_CatchupPolicy = 2
)

// Enum value maps for GetBlockStreamRequest_CatchupPolicy.
var (
	GetBlockStreamRequest_CatchupPolicy_name = map[int32]string{
		0: "PANIC",
		1: "WAIT",
		2: "STREAM",
	}
	GetBlockStreamRequest_CatchupPolicy_value = map[string]int32{
		"PANIC":  0,
		"WAIT":   1,
		"STREAM": 2,
	}
)

func (x GetBlockStreamRequest_CatchupPolicy) Enum() *GetBlockStreamRequest_CatchupPolicy {
	p := new(GetBlockStreamRequest_CatchupPolicy)
	*p = x
	return p
}

func (x GetBlockStreamRequest_CatchupPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetBlockStreamRequest_CatchupPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_blocksapi_stream_proto_enumTypes[1].Descriptor()
}

func (GetBlockStreamRequest_CatchupPolicy) Type() protoreflect.EnumType {
	return &file_blocksapi_stream_proto_enumTypes[1]
}

func (x GetBlockStreamRequest_CatchupPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetBlockStreamRequest_CatchupPolicy.Descriptor instead.
func (GetBlockStreamRequest_CatchupPolicy) EnumDescriptor() ([]byte, []int) {
	return file_blocksapi_stream_proto_rawDescGZIP(), []int{2, 1}
}

// Defines when stream has to stop
type GetBlockStreamRequest_StopPolicy int32

const (
	// Follow new blocks
	GetBlockStreamRequest_NEVER GetBlockStreamRequest_StopPolicy = 0
	// Don't send messages greater than target
	GetBlockStreamRequest_AFTER_TARGET GetBlockStreamRequest_StopPolicy = 1
	// Don't send messages greater or equal to target
	GetBlockStreamRequest_BEFORE_TARGET GetBlockStreamRequest_StopPolicy = 2
)

// Enum value maps for GetBlockStreamRequest_StopPolicy.
var (
	GetBlockStreamRequest_StopPolicy_name = map[int32]string{
		0: "NEVER",
		1: "AFTER_TARGET",
		2: "BEFORE_TARGET",
	}
	GetBlockStreamRequest_StopPolicy_value = map[string]int32{
		"NEVER":         0,
		"AFTER_TARGET":  1,
		"BEFORE_TARGET": 2,
	}
)

func (x GetBlockStreamRequest_StopPolicy) Enum() *GetBlockStreamRequest_StopPolicy {
	p := new(GetBlockStreamRequest_StopPolicy)
	*p = x
	return p
}

func (x GetBlockStreamRequest_StopPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetBlockStreamRequest_StopPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_blocksapi_stream_proto_enumTypes[2].Descriptor()
}

func (GetBlockStreamRequest_StopPolicy) Type() protoreflect.EnumType {
	return &file_blocksapi_stream_proto_enumTypes[2]
}

func (x GetBlockStreamRequest_StopPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetBlockStreamRequest_StopPolicy.Descriptor instead.
func (GetBlockStreamRequest_StopPolicy) EnumDescriptor() ([]byte, []int) {
	return file_blocksapi_stream_proto_rawDescGZIP(), []int{2, 2}
}

type GetBlockStreamResponse_Error_Kind int32

const (
	// Default error class
	GetBlockStreamResponse_Error_UNKNOWN GetBlockStreamResponse_Error_Kind = 0
	// Catch up required, but catchup policy is PANIC
	GetBlockStreamResponse_Error_CATCHUP_REQUIRED GetBlockStreamResponse_Error_Kind = 1
	// Request is constructed in a wrong way
	GetBlockStreamResponse_Error_BAD_REQUEST GetBlockStreamResponse_Error_Kind = 2
)

// Enum value maps for GetBlockStreamResponse_Error_Kind.
var (
	GetBlockStreamResponse_Error_Kind_name = map[int32]string{
		0: "UNKNOWN",
		1: "CATCHUP_REQUIRED",
		2: "BAD_REQUEST",
	}
	GetBlockStreamResponse_Error_Kind_value = map[string]int32{
		"UNKNOWN":          0,
		"CATCHUP_REQUIRED": 1,
		"BAD_REQUEST":      2,
	}
)

func (x GetBlockStreamResponse_Error_Kind) Enum() *GetBlockStreamResponse_Error_Kind {
	p := new(GetBlockStreamResponse_Error_Kind)
	*p = x
	return p
}

func (x GetBlockStreamResponse_Error_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetBlockStreamResponse_Error_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_blocksapi_stream_proto_enumTypes[3].Descriptor()
}

func (GetBlockStreamResponse_Error_Kind) Type() protoreflect.EnumType {
	return &file_blocksapi_stream_proto_enumTypes[3]
}

func (x GetBlockStreamResponse_Error_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetBlockStreamResponse_Error_Kind.Descriptor instead.
func (GetBlockStreamResponse_Error_Kind) EnumDescriptor() ([]byte, []int) {
	return file_blocksapi_stream_proto_rawDescGZIP(), []int{3, 2, 0}
}

type BlockStreamFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExcludeShards bool     `protobuf:"varint,1,opt,name=exclude_shards,json=excludeShards,proto3" json:"exclude_shards,omitempty"`
	FilterShards  []uint64 `protobuf:"varint,2,rep,packed,name=filter_shards,json=filterShards,proto3" json:"filter_shards,omitempty"`
}

func (x *BlockStreamFilter) Reset() {
	*x = BlockStreamFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocksapi_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockStreamFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockStreamFilter) ProtoMessage() {}

func (x *BlockStreamFilter) ProtoReflect() protoreflect.Message {
	mi := &file_blocksapi_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockStreamFilter.ProtoReflect.Descriptor instead.
func (*BlockStreamFilter) Descriptor() ([]byte, []int) {
	return file_blocksapi_stream_proto_rawDescGZIP(), []int{0}
}

func (x *BlockStreamFilter) GetExcludeShards() bool {
	if x != nil {
		return x.ExcludeShards
	}
	return false
}

func (x *BlockStreamFilter) GetFilterShards() []uint64 {
	if x != nil {
		return x.FilterShards
	}
	return nil
}

type BlockStreamDeliverySettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter  *BlockStreamFilter            `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	Content *BlockMessageDeliverySettings `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *BlockStreamDeliverySettings) Reset() {
	*x = BlockStreamDeliverySettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocksapi_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockStreamDeliverySettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockStreamDeliverySettings) ProtoMessage() {}

func (x *BlockStreamDeliverySettings) ProtoReflect() protoreflect.Message {
	mi := &file_blocksapi_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockStreamDeliverySettings.ProtoReflect.Descriptor instead.
func (*BlockStreamDeliverySettings) Descriptor() ([]byte, []int) {
	return file_blocksapi_stream_proto_rawDescGZIP(), []int{1}
}

func (x *BlockStreamDeliverySettings) GetFilter() *BlockStreamFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *BlockStreamDeliverySettings) GetContent() *BlockMessageDeliverySettings {
	if x != nil {
		return x.Content
	}
	return nil
}

type GetBlockStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamName       string                              `protobuf:"bytes,1,opt,name=stream_name,json=streamName,proto3" json:"stream_name,omitempty"`
	StartPolicy      GetBlockStreamRequest_StartPolicy   `protobuf:"varint,2,opt,name=start_policy,json=startPolicy,proto3,enum=GetBlockStreamRequest_StartPolicy" json:"start_policy,omitempty"`
	StartTarget      *BlockMessage_ID                    `protobuf:"bytes,3,opt,name=start_target,json=startTarget,proto3,oneof" json:"start_target,omitempty"`
	StopPolicy       GetBlockStreamRequest_StopPolicy    `protobuf:"varint,4,opt,name=stop_policy,json=stopPolicy,proto3,enum=GetBlockStreamRequest_StopPolicy" json:"stop_policy,omitempty"`
	StopTarget       *BlockMessage_ID                    `protobuf:"bytes,5,opt,name=stop_target,json=stopTarget,proto3,oneof" json:"stop_target,omitempty"`
	DeliverySettings *BlockStreamDeliverySettings        `protobuf:"bytes,6,opt,name=delivery_settings,json=deliverySettings,proto3" json:"delivery_settings,omitempty"`
	CatchupPolicy    GetBlockStreamRequest_CatchupPolicy `protobuf:"varint,7,opt,name=catchup_policy,json=catchupPolicy,proto3,enum=GetBlockStreamRequest_CatchupPolicy" json:"catchup_policy,omitempty"`
	// If not provided - default delivery settings are used during catchup
	CatchupDeliverySettings *BlockStreamDeliverySettings `protobuf:"bytes,8,opt,name=catchup_delivery_settings,json=catchupDeliverySettings,proto3,oneof" json:"catchup_delivery_settings,omitempty"`
}

func (x *GetBlockStreamRequest) Reset() {
	*x = GetBlockStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocksapi_stream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockStreamRequest) ProtoMessage() {}

func (x *GetBlockStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blocksapi_stream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockStreamRequest.ProtoReflect.Descriptor instead.
func (*GetBlockStreamRequest) Descriptor() ([]byte, []int) {
	return file_blocksapi_stream_proto_rawDescGZIP(), []int{2}
}

func (x *GetBlockStreamRequest) GetStreamName() string {
	if x != nil {
		return x.StreamName
	}
	return ""
}

func (x *GetBlockStreamRequest) GetStartPolicy() GetBlockStreamRequest_StartPolicy {
	if x != nil {
		return x.StartPolicy
	}
	return GetBlockStreamRequest_ON_EARLIEST_AVAILABLE
}

func (x *GetBlockStreamRequest) GetStartTarget() *BlockMessage_ID {
	if x != nil {
		return x.StartTarget
	}
	return nil
}

func (x *GetBlockStreamRequest) GetStopPolicy() GetBlockStreamRequest_StopPolicy {
	if x != nil {
		return x.StopPolicy
	}
	return GetBlockStreamRequest_NEVER
}

func (x *GetBlockStreamRequest) GetStopTarget() *BlockMessage_ID {
	if x != nil {
		return x.StopTarget
	}
	return nil
}

func (x *GetBlockStreamRequest) GetDeliverySettings() *BlockStreamDeliverySettings {
	if x != nil {
		return x.DeliverySettings
	}
	return nil
}

func (x *GetBlockStreamRequest) GetCatchupPolicy() GetBlockStreamRequest_CatchupPolicy {
	if x != nil {
		return x.CatchupPolicy
	}
	return GetBlockStreamRequest_PANIC
}

func (x *GetBlockStreamRequest) GetCatchupDeliverySettings() *BlockStreamDeliverySettings {
	if x != nil {
		return x.CatchupDeliverySettings
	}
	return nil
}

type GetBlockStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//
	//	*GetBlockStreamResponse_Message
	//	*GetBlockStreamResponse_Done_
	//	*GetBlockStreamResponse_Error_
	Response isGetBlockStreamResponse_Response `protobuf_oneof:"response"`
}

func (x *GetBlockStreamResponse) Reset() {
	*x = GetBlockStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocksapi_stream_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockStreamResponse) ProtoMessage() {}

func (x *GetBlockStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blocksapi_stream_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockStreamResponse.ProtoReflect.Descriptor instead.
func (*GetBlockStreamResponse) Descriptor() ([]byte, []int) {
	return file_blocksapi_stream_proto_rawDescGZIP(), []int{3}
}

func (m *GetBlockStreamResponse) GetResponse() isGetBlockStreamResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *GetBlockStreamResponse) GetMessage() *GetBlockStreamResponse_Result {
	if x, ok := x.GetResponse().(*GetBlockStreamResponse_Message); ok {
		return x.Message
	}
	return nil
}

func (x *GetBlockStreamResponse) GetDone() *GetBlockStreamResponse_Done {
	if x, ok := x.GetResponse().(*GetBlockStreamResponse_Done_); ok {
		return x.Done
	}
	return nil
}

func (x *GetBlockStreamResponse) GetError() *GetBlockStreamResponse_Error {
	if x, ok := x.GetResponse().(*GetBlockStreamResponse_Error_); ok {
		return x.Error
	}
	return nil
}

type isGetBlockStreamResponse_Response interface {
	isGetBlockStreamResponse_Response()
}

type GetBlockStreamResponse_Message struct {
	Message *GetBlockStreamResponse_Result `protobuf:"bytes,1,opt,name=message,proto3,oneof"`
}

type GetBlockStreamResponse_Done_ struct {
	Done *GetBlockStreamResponse_Done `protobuf:"bytes,2,opt,name=done,proto3,oneof"`
}

type GetBlockStreamResponse_Error_ struct {
	Error *GetBlockStreamResponse_Error `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*GetBlockStreamResponse_Message) isGetBlockStreamResponse_Response() {}

func (*GetBlockStreamResponse_Done_) isGetBlockStreamResponse_Response() {}

func (*GetBlockStreamResponse_Error_) isGetBlockStreamResponse_Response() {}

type GetBlockStreamResponse_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message           *BlockMessage `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	CatchupInProgress bool          `protobuf:"varint,2,opt,name=catchup_in_progress,json=catchupInProgress,proto3" json:"catchup_in_progress,omitempty"` // TODO: maybe add auxiliary info that helps understanding distance to latest block
}

func (x *GetBlockStreamResponse_Result) Reset() {
	*x = GetBlockStreamResponse_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocksapi_stream_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockStreamResponse_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockStreamResponse_Result) ProtoMessage() {}

func (x *GetBlockStreamResponse_Result) ProtoReflect() protoreflect.Message {
	mi := &file_blocksapi_stream_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockStreamResponse_Result.ProtoReflect.Descriptor instead.
func (*GetBlockStreamResponse_Result) Descriptor() ([]byte, []int) {
	return file_blocksapi_stream_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetBlockStreamResponse_Result) GetMessage() *BlockMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *GetBlockStreamResponse_Result) GetCatchupInProgress() bool {
	if x != nil {
		return x.CatchupInProgress
	}
	return false
}

type GetBlockStreamResponse_Done struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *GetBlockStreamResponse_Done) Reset() {
	*x = GetBlockStreamResponse_Done{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocksapi_stream_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockStreamResponse_Done) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockStreamResponse_Done) ProtoMessage() {}

func (x *GetBlockStreamResponse_Done) ProtoReflect() protoreflect.Message {
	mi := &file_blocksapi_stream_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockStreamResponse_Done.ProtoReflect.Descriptor instead.
func (*GetBlockStreamResponse_Done) Descriptor() ([]byte, []int) {
	return file_blocksapi_stream_proto_rawDescGZIP(), []int{3, 1}
}

func (x *GetBlockStreamResponse_Done) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetBlockStreamResponse_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind        GetBlockStreamResponse_Error_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=GetBlockStreamResponse_Error_Kind" json:"kind,omitempty"`
	Description string                            `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *GetBlockStreamResponse_Error) Reset() {
	*x = GetBlockStreamResponse_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocksapi_stream_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockStreamResponse_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockStreamResponse_Error) ProtoMessage() {}

func (x *GetBlockStreamResponse_Error) ProtoReflect() protoreflect.Message {
	mi := &file_blocksapi_stream_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockStreamResponse_Error.ProtoReflect.Descriptor instead.
func (*GetBlockStreamResponse_Error) Descriptor() ([]byte, []int) {
	return file_blocksapi_stream_proto_rawDescGZIP(), []int{3, 2}
}

func (x *GetBlockStreamResponse_Error) GetKind() GetBlockStreamResponse_Error_Kind {
	if x != nil {
		return x.Kind
	}
	return GetBlockStreamResponse_Error_UNKNOWN
}

func (x *GetBlockStreamResponse_Error) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_blocksapi_stream_proto protoreflect.FileDescriptor

var file_blocksapi_stream_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5f, 0x0a, 0x11, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x53, 0x68, 0x61, 0x72, 0x64, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x04, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x68, 0x61, 0x72,
	0x64, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x1b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x37,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x88, 0x07, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0b, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x38, 0x0a, 0x0c, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49,
	0x44, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x42, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x53, 0x74, 0x6f, 0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0a, 0x73, 0x74, 0x6f,
	0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x36, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x70, 0x5f,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x44, 0x48, 0x01,
	0x52, 0x0a, 0x73, 0x74, 0x6f, 0x70, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x49, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x4b, 0x0a, 0x0e, 0x63, 0x61,
	0x74, 0x63, 0x68, 0x75, 0x70, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x24, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x61, 0x74, 0x63, 0x68,
	0x75, 0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0d, 0x63, 0x61, 0x74, 0x63, 0x68, 0x75,
	0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x5d, 0x0a, 0x19, 0x63, 0x61, 0x74, 0x63, 0x68,
	0x75, 0x70, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x48, 0x02, 0x52, 0x17, 0x63, 0x61, 0x74, 0x63,
	0x68, 0x75, 0x70, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x88, 0x01, 0x01, 0x22, 0xaa, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x4e, 0x5f, 0x45, 0x41, 0x52,
	0x4c, 0x49, 0x45, 0x53, 0x54, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x00, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x4e, 0x5f, 0x4c, 0x41, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x41,
	0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x58,
	0x41, 0x43, 0x54, 0x4c, 0x59, 0x5f, 0x4f, 0x4e, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x10,
	0x02, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x58, 0x41, 0x43, 0x54, 0x4c, 0x59, 0x5f, 0x41, 0x46, 0x54,
	0x45, 0x52, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x4f,
	0x4e, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x53, 0x54, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x41, 0x52,
	0x47, 0x45, 0x54, 0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18, 0x4f, 0x4e, 0x5f, 0x45, 0x41, 0x52, 0x4c,
	0x49, 0x45, 0x53, 0x54, 0x5f, 0x41, 0x46, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45,
	0x54, 0x10, 0x05, 0x22, 0x30, 0x0a, 0x0d, 0x43, 0x61, 0x74, 0x63, 0x68, 0x75, 0x70, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41, 0x4e, 0x49, 0x43, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x57, 0x41, 0x49, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52,
	0x45, 0x41, 0x4d, 0x10, 0x02, 0x22, 0x3c, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x70, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x45, 0x56, 0x45, 0x52, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0c, 0x41, 0x46, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x10, 0x01,
	0x12, 0x11, 0x0a, 0x0d, 0x42, 0x45, 0x46, 0x4f, 0x52, 0x45, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45,
	0x54, 0x10, 0x02, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x42, 0x1c, 0x0a, 0x1a, 0x5f, 0x63, 0x61, 0x74, 0x63, 0x68, 0x75, 0x70,
	0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x22, 0xf8, 0x03, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x6f, 0x6e,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x44, 0x6f, 0x6e, 0x65, 0x48, 0x00, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x12, 0x35, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x1a, 0x61, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x27,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x61, 0x74, 0x63, 0x68,
	0x75, 0x70, 0x5f, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x63, 0x61, 0x74, 0x63, 0x68, 0x75, 0x70, 0x49, 0x6e, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x28, 0x0a, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x9d, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x36, 0x0a, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x41,
	0x54, 0x43, 0x48, 0x55, 0x50, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x41, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10,
	0x02, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x46, 0x5a,
	0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x72, 0x6f,
	0x72, 0x61, 0x2d, 0x69, 0x73, 0x2d, 0x6e, 0x65, 0x61, 0x72, 0x2f, 0x62, 0x6f, 0x72, 0x65, 0x61,
	0x6c, 0x69, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67,
	0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x3b, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blocksapi_stream_proto_rawDescOnce sync.Once
	file_blocksapi_stream_proto_rawDescData = file_blocksapi_stream_proto_rawDesc
)

func file_blocksapi_stream_proto_rawDescGZIP() []byte {
	file_blocksapi_stream_proto_rawDescOnce.Do(func() {
		file_blocksapi_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_blocksapi_stream_proto_rawDescData)
	})
	return file_blocksapi_stream_proto_rawDescData
}

var file_blocksapi_stream_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_blocksapi_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_blocksapi_stream_proto_goTypes = []interface{}{
	(GetBlockStreamRequest_StartPolicy)(0),   // 0: GetBlockStreamRequest.StartPolicy
	(GetBlockStreamRequest_CatchupPolicy)(0), // 1: GetBlockStreamRequest.CatchupPolicy
	(GetBlockStreamRequest_StopPolicy)(0),    // 2: GetBlockStreamRequest.StopPolicy
	(GetBlockStreamResponse_Error_Kind)(0),   // 3: GetBlockStreamResponse.Error.Kind
	(*BlockStreamFilter)(nil),                // 4: BlockStreamFilter
	(*BlockStreamDeliverySettings)(nil),      // 5: BlockStreamDeliverySettings
	(*GetBlockStreamRequest)(nil),            // 6: GetBlockStreamRequest
	(*GetBlockStreamResponse)(nil),           // 7: GetBlockStreamResponse
	(*GetBlockStreamResponse_Result)(nil),    // 8: GetBlockStreamResponse.Result
	(*GetBlockStreamResponse_Done)(nil),      // 9: GetBlockStreamResponse.Done
	(*GetBlockStreamResponse_Error)(nil),     // 10: GetBlockStreamResponse.Error
	(*BlockMessageDeliverySettings)(nil),     // 11: BlockMessageDeliverySettings
	(*BlockMessage_ID)(nil),                  // 12: BlockMessage.ID
	(*BlockMessage)(nil),                     // 13: BlockMessage
}
var file_blocksapi_stream_proto_depIdxs = []int32{
	4,  // 0: BlockStreamDeliverySettings.filter:type_name -> BlockStreamFilter
	11, // 1: BlockStreamDeliverySettings.content:type_name -> BlockMessageDeliverySettings
	0,  // 2: GetBlockStreamRequest.start_policy:type_name -> GetBlockStreamRequest.StartPolicy
	12, // 3: GetBlockStreamRequest.start_target:type_name -> BlockMessage.ID
	2,  // 4: GetBlockStreamRequest.stop_policy:type_name -> GetBlockStreamRequest.StopPolicy
	12, // 5: GetBlockStreamRequest.stop_target:type_name -> BlockMessage.ID
	5,  // 6: GetBlockStreamRequest.delivery_settings:type_name -> BlockStreamDeliverySettings
	1,  // 7: GetBlockStreamRequest.catchup_policy:type_name -> GetBlockStreamRequest.CatchupPolicy
	5,  // 8: GetBlockStreamRequest.catchup_delivery_settings:type_name -> BlockStreamDeliverySettings
	8,  // 9: GetBlockStreamResponse.message:type_name -> GetBlockStreamResponse.Result
	9,  // 10: GetBlockStreamResponse.done:type_name -> GetBlockStreamResponse.Done
	10, // 11: GetBlockStreamResponse.error:type_name -> GetBlockStreamResponse.Error
	13, // 12: GetBlockStreamResponse.Result.message:type_name -> BlockMessage
	3,  // 13: GetBlockStreamResponse.Error.kind:type_name -> GetBlockStreamResponse.Error.Kind
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_blocksapi_stream_proto_init() }
func file_blocksapi_stream_proto_init() {
	if File_blocksapi_stream_proto != nil {
		return
	}
	file_blocksapi_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_blocksapi_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockStreamFilter); i {
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
		file_blocksapi_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockStreamDeliverySettings); i {
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
		file_blocksapi_stream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockStreamRequest); i {
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
		file_blocksapi_stream_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockStreamResponse); i {
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
		file_blocksapi_stream_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockStreamResponse_Result); i {
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
		file_blocksapi_stream_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockStreamResponse_Done); i {
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
		file_blocksapi_stream_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockStreamResponse_Error); i {
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
	file_blocksapi_stream_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_blocksapi_stream_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*GetBlockStreamResponse_Message)(nil),
		(*GetBlockStreamResponse_Done_)(nil),
		(*GetBlockStreamResponse_Error_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blocksapi_stream_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blocksapi_stream_proto_goTypes,
		DependencyIndexes: file_blocksapi_stream_proto_depIdxs,
		EnumInfos:         file_blocksapi_stream_proto_enumTypes,
		MessageInfos:      file_blocksapi_stream_proto_msgTypes,
	}.Build()
	File_blocksapi_stream_proto = out.File
	file_blocksapi_stream_proto_rawDesc = nil
	file_blocksapi_stream_proto_goTypes = nil
	file_blocksapi_stream_proto_depIdxs = nil
}
