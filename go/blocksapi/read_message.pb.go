// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v3.21.12
// source: blocksapi/read_message.proto

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

type ReadBlockMessageResponse_Error_Kind int32

const (
	// Default error kind
	ReadBlockMessageResponse_Error_UNKNOWN ReadBlockMessageResponse_Error_Kind = 0
	// Given message ID is lower than earliest available
	ReadBlockMessageResponse_Error_LOW_ID ReadBlockMessageResponse_Error_Kind = 1
	// Given message ID is greater than latest available
	ReadBlockMessageResponse_Error_HIGH_ID ReadBlockMessageResponse_Error_Kind = 2
	// Message ID is within given range but not present
	ReadBlockMessageResponse_Error_NOT_FOUND ReadBlockMessageResponse_Error_Kind = 3
)

// Enum value maps for ReadBlockMessageResponse_Error_Kind.
var (
	ReadBlockMessageResponse_Error_Kind_name = map[int32]string{
		0: "UNKNOWN",
		1: "LOW_ID",
		2: "HIGH_ID",
		3: "NOT_FOUND",
	}
	ReadBlockMessageResponse_Error_Kind_value = map[string]int32{
		"UNKNOWN":   0,
		"LOW_ID":    1,
		"HIGH_ID":   2,
		"NOT_FOUND": 3,
	}
)

func (x ReadBlockMessageResponse_Error_Kind) Enum() *ReadBlockMessageResponse_Error_Kind {
	p := new(ReadBlockMessageResponse_Error_Kind)
	*p = x
	return p
}

func (x ReadBlockMessageResponse_Error_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReadBlockMessageResponse_Error_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_blocksapi_read_message_proto_enumTypes[0].Descriptor()
}

func (ReadBlockMessageResponse_Error_Kind) Type() protoreflect.EnumType {
	return &file_blocksapi_read_message_proto_enumTypes[0]
}

func (x ReadBlockMessageResponse_Error_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReadBlockMessageResponse_Error_Kind.Descriptor instead.
func (ReadBlockMessageResponse_Error_Kind) EnumDescriptor() ([]byte, []int) {
	return file_blocksapi_read_message_proto_rawDescGZIP(), []int{1, 1, 0}
}

type ReadBlockMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamName       string                        `protobuf:"bytes,1,opt,name=stream_name,json=streamName,proto3" json:"stream_name,omitempty"`
	MessageId        *BlockMessage_ID              `protobuf:"bytes,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	DeliverySettings *BlockMessageDeliverySettings `protobuf:"bytes,3,opt,name=delivery_settings,json=deliverySettings,proto3" json:"delivery_settings,omitempty"`
}

func (x *ReadBlockMessageRequest) Reset() {
	*x = ReadBlockMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocksapi_read_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadBlockMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadBlockMessageRequest) ProtoMessage() {}

func (x *ReadBlockMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blocksapi_read_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadBlockMessageRequest.ProtoReflect.Descriptor instead.
func (*ReadBlockMessageRequest) Descriptor() ([]byte, []int) {
	return file_blocksapi_read_message_proto_rawDescGZIP(), []int{0}
}

func (x *ReadBlockMessageRequest) GetStreamName() string {
	if x != nil {
		return x.StreamName
	}
	return ""
}

func (x *ReadBlockMessageRequest) GetMessageId() *BlockMessage_ID {
	if x != nil {
		return x.MessageId
	}
	return nil
}

func (x *ReadBlockMessageRequest) GetDeliverySettings() *BlockMessageDeliverySettings {
	if x != nil {
		return x.DeliverySettings
	}
	return nil
}

type ReadBlockMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//
	//	*ReadBlockMessageResponse_Result_
	//	*ReadBlockMessageResponse_Error_
	Response isReadBlockMessageResponse_Response `protobuf_oneof:"response"`
}

func (x *ReadBlockMessageResponse) Reset() {
	*x = ReadBlockMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocksapi_read_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadBlockMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadBlockMessageResponse) ProtoMessage() {}

func (x *ReadBlockMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blocksapi_read_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadBlockMessageResponse.ProtoReflect.Descriptor instead.
func (*ReadBlockMessageResponse) Descriptor() ([]byte, []int) {
	return file_blocksapi_read_message_proto_rawDescGZIP(), []int{1}
}

func (m *ReadBlockMessageResponse) GetResponse() isReadBlockMessageResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *ReadBlockMessageResponse) GetResult() *ReadBlockMessageResponse_Result {
	if x, ok := x.GetResponse().(*ReadBlockMessageResponse_Result_); ok {
		return x.Result
	}
	return nil
}

func (x *ReadBlockMessageResponse) GetError() *ReadBlockMessageResponse_Error {
	if x, ok := x.GetResponse().(*ReadBlockMessageResponse_Error_); ok {
		return x.Error
	}
	return nil
}

type isReadBlockMessageResponse_Response interface {
	isReadBlockMessageResponse_Response()
}

type ReadBlockMessageResponse_Result_ struct {
	Result *ReadBlockMessageResponse_Result `protobuf:"bytes,1,opt,name=result,proto3,oneof"`
}

type ReadBlockMessageResponse_Error_ struct {
	Error *ReadBlockMessageResponse_Error `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*ReadBlockMessageResponse_Result_) isReadBlockMessageResponse_Response() {}

func (*ReadBlockMessageResponse_Error_) isReadBlockMessageResponse_Response() {}

type ReadBlockMessageResponse_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *BlockMessage `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ReadBlockMessageResponse_Result) Reset() {
	*x = ReadBlockMessageResponse_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocksapi_read_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadBlockMessageResponse_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadBlockMessageResponse_Result) ProtoMessage() {}

func (x *ReadBlockMessageResponse_Result) ProtoReflect() protoreflect.Message {
	mi := &file_blocksapi_read_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadBlockMessageResponse_Result.ProtoReflect.Descriptor instead.
func (*ReadBlockMessageResponse_Result) Descriptor() ([]byte, []int) {
	return file_blocksapi_read_message_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ReadBlockMessageResponse_Result) GetMessage() *BlockMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

type ReadBlockMessageResponse_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind        ReadBlockMessageResponse_Error_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=ReadBlockMessageResponse_Error_Kind" json:"kind,omitempty"`
	Description string                              `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ReadBlockMessageResponse_Error) Reset() {
	*x = ReadBlockMessageResponse_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocksapi_read_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadBlockMessageResponse_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadBlockMessageResponse_Error) ProtoMessage() {}

func (x *ReadBlockMessageResponse_Error) ProtoReflect() protoreflect.Message {
	mi := &file_blocksapi_read_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadBlockMessageResponse_Error.ProtoReflect.Descriptor instead.
func (*ReadBlockMessageResponse_Error) Descriptor() ([]byte, []int) {
	return file_blocksapi_read_message_proto_rawDescGZIP(), []int{1, 1}
}

func (x *ReadBlockMessageResponse_Error) GetKind() ReadBlockMessageResponse_Error_Kind {
	if x != nil {
		return x.Kind
	}
	return ReadBlockMessageResponse_Error_UNKNOWN
}

func (x *ReadBlockMessageResponse_Error) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_blocksapi_read_message_proto protoreflect.FileDescriptor

var file_blocksapi_read_message_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x61, 0x64,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x61, 0x64,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x44, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52,
	0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x22, 0xf1, 0x02, 0x0a, 0x18, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x37, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x1a, 0x31, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x27, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0xa0, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x38, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3b, 0x0a, 0x04,
	0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x4f, 0x57, 0x5f, 0x49, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x48, 0x49, 0x47, 0x48, 0x5f, 0x49, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f,
	0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x72, 0x6f, 0x72, 0x61, 0x2d, 0x69, 0x73, 0x2d, 0x6e, 0x65,
	0x61, 0x72, 0x2f, 0x62, 0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x61, 0x70, 0x69, 0x3b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blocksapi_read_message_proto_rawDescOnce sync.Once
	file_blocksapi_read_message_proto_rawDescData = file_blocksapi_read_message_proto_rawDesc
)

func file_blocksapi_read_message_proto_rawDescGZIP() []byte {
	file_blocksapi_read_message_proto_rawDescOnce.Do(func() {
		file_blocksapi_read_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_blocksapi_read_message_proto_rawDescData)
	})
	return file_blocksapi_read_message_proto_rawDescData
}

var file_blocksapi_read_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_blocksapi_read_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_blocksapi_read_message_proto_goTypes = []interface{}{
	(ReadBlockMessageResponse_Error_Kind)(0), // 0: ReadBlockMessageResponse.Error.Kind
	(*ReadBlockMessageRequest)(nil),          // 1: ReadBlockMessageRequest
	(*ReadBlockMessageResponse)(nil),         // 2: ReadBlockMessageResponse
	(*ReadBlockMessageResponse_Result)(nil),  // 3: ReadBlockMessageResponse.Result
	(*ReadBlockMessageResponse_Error)(nil),   // 4: ReadBlockMessageResponse.Error
	(*BlockMessage_ID)(nil),                  // 5: BlockMessage.ID
	(*BlockMessageDeliverySettings)(nil),     // 6: BlockMessageDeliverySettings
	(*BlockMessage)(nil),                     // 7: BlockMessage
}
var file_blocksapi_read_message_proto_depIdxs = []int32{
	5, // 0: ReadBlockMessageRequest.message_id:type_name -> BlockMessage.ID
	6, // 1: ReadBlockMessageRequest.delivery_settings:type_name -> BlockMessageDeliverySettings
	3, // 2: ReadBlockMessageResponse.result:type_name -> ReadBlockMessageResponse.Result
	4, // 3: ReadBlockMessageResponse.error:type_name -> ReadBlockMessageResponse.Error
	7, // 4: ReadBlockMessageResponse.Result.message:type_name -> BlockMessage
	0, // 5: ReadBlockMessageResponse.Error.kind:type_name -> ReadBlockMessageResponse.Error.Kind
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_blocksapi_read_message_proto_init() }
func file_blocksapi_read_message_proto_init() {
	if File_blocksapi_read_message_proto != nil {
		return
	}
	file_blocksapi_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_blocksapi_read_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadBlockMessageRequest); i {
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
		file_blocksapi_read_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadBlockMessageResponse); i {
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
		file_blocksapi_read_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadBlockMessageResponse_Result); i {
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
		file_blocksapi_read_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadBlockMessageResponse_Error); i {
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
	file_blocksapi_read_message_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ReadBlockMessageResponse_Result_)(nil),
		(*ReadBlockMessageResponse_Error_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blocksapi_read_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blocksapi_read_message_proto_goTypes,
		DependencyIndexes: file_blocksapi_read_message_proto_depIdxs,
		EnumInfos:         file_blocksapi_read_message_proto_enumTypes,
		MessageInfos:      file_blocksapi_read_message_proto_msgTypes,
	}.Build()
	File_blocksapi_read_message_proto = out.File
	file_blocksapi_read_message_proto_rawDesc = nil
	file_blocksapi_read_message_proto_goTypes = nil
	file_blocksapi_read_message_proto_depIdxs = nil
}
