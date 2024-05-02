// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v3.21.12
// source: blocksapi/message.proto

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

type BlockMessage_Kind int32

const (
	BlockMessage_MSG_WHOLE  BlockMessage_Kind = 0
	BlockMessage_MSG_HEADER BlockMessage_Kind = 1
	BlockMessage_MSG_CHUNK  BlockMessage_Kind = 2
)

// Enum value maps for BlockMessage_Kind.
var (
	BlockMessage_Kind_name = map[int32]string{
		0: "MSG_WHOLE",
		1: "MSG_HEADER",
		2: "MSG_CHUNK",
	}
	BlockMessage_Kind_value = map[string]int32{
		"MSG_WHOLE":  0,
		"MSG_HEADER": 1,
		"MSG_CHUNK":  2,
	}
)

func (x BlockMessage_Kind) Enum() *BlockMessage_Kind {
	p := new(BlockMessage_Kind)
	*p = x
	return p
}

func (x BlockMessage_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockMessage_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_blocksapi_message_proto_enumTypes[0].Descriptor()
}

func (BlockMessage_Kind) Type() protoreflect.EnumType {
	return &file_blocksapi_message_proto_enumTypes[0]
}

func (x BlockMessage_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlockMessage_Kind.Descriptor instead.
func (BlockMessage_Kind) EnumDescriptor() ([]byte, []int) {
	return file_blocksapi_message_proto_rawDescGZIP(), []int{0, 0}
}

type BlockMessage_Format int32

const (
	// JSON -> LZ4 -> CBOR bytes -> borealis envelope (whole blocks)
	BlockMessage_PAYLOAD_NEAR_BLOCK_V2 BlockMessage_Format = 0
	// CBOR STRUCT -> borealis envelope (whole blocks)
	BlockMessage_PAYLOAD_AURORA_BLOCK_V2 BlockMessage_Format = 1
	// protobuf (block headers + block shards)
	BlockMessage_PAYLOAD_NEAR_BLOCK_V3 BlockMessage_Format = 2
)

// Enum value maps for BlockMessage_Format.
var (
	BlockMessage_Format_name = map[int32]string{
		0: "PAYLOAD_NEAR_BLOCK_V2",
		1: "PAYLOAD_AURORA_BLOCK_V2",
		2: "PAYLOAD_NEAR_BLOCK_V3",
	}
	BlockMessage_Format_value = map[string]int32{
		"PAYLOAD_NEAR_BLOCK_V2":   0,
		"PAYLOAD_AURORA_BLOCK_V2": 1,
		"PAYLOAD_NEAR_BLOCK_V3":   2,
	}
)

func (x BlockMessage_Format) Enum() *BlockMessage_Format {
	p := new(BlockMessage_Format)
	*p = x
	return p
}

func (x BlockMessage_Format) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockMessage_Format) Descriptor() protoreflect.EnumDescriptor {
	return file_blocksapi_message_proto_enumTypes[1].Descriptor()
}

func (BlockMessage_Format) Type() protoreflect.EnumType {
	return &file_blocksapi_message_proto_enumTypes[1]
}

func (x BlockMessage_Format) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlockMessage_Format.Descriptor instead.
func (BlockMessage_Format) EnumDescriptor() ([]byte, []int) {
	return file_blocksapi_message_proto_rawDescGZIP(), []int{0, 1}
}

type BlockMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             *BlockMessage_ID    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Format         BlockMessage_Format `protobuf:"varint,2,opt,name=format,proto3,enum=BlockMessage_Format" json:"format,omitempty"`
	ZstdCompressed bool                `protobuf:"varint,3,opt,name=zstd_compressed,json=zstdCompressed,proto3" json:"zstd_compressed,omitempty"`
	// Types that are assignable to Payload:
	//
	//	*BlockMessage_RawPayload
	Payload isBlockMessage_Payload `protobuf_oneof:"payload"`
}

func (x *BlockMessage) Reset() {
	*x = BlockMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocksapi_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockMessage) ProtoMessage() {}

func (x *BlockMessage) ProtoReflect() protoreflect.Message {
	mi := &file_blocksapi_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockMessage.ProtoReflect.Descriptor instead.
func (*BlockMessage) Descriptor() ([]byte, []int) {
	return file_blocksapi_message_proto_rawDescGZIP(), []int{0}
}

func (x *BlockMessage) GetId() *BlockMessage_ID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *BlockMessage) GetFormat() BlockMessage_Format {
	if x != nil {
		return x.Format
	}
	return BlockMessage_PAYLOAD_NEAR_BLOCK_V2
}

func (x *BlockMessage) GetZstdCompressed() bool {
	if x != nil {
		return x.ZstdCompressed
	}
	return false
}

func (m *BlockMessage) GetPayload() isBlockMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *BlockMessage) GetRawPayload() []byte {
	if x, ok := x.GetPayload().(*BlockMessage_RawPayload); ok {
		return x.RawPayload
	}
	return nil
}

type isBlockMessage_Payload interface {
	isBlockMessage_Payload()
}

type BlockMessage_RawPayload struct {
	RawPayload []byte `protobuf:"bytes,4,opt,name=raw_payload,json=rawPayload,proto3,oneof"`
}

func (*BlockMessage_RawPayload) isBlockMessage_Payload() {}

type BlockMessageDeliverySettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExcludePayload  bool                 `protobuf:"varint,1,opt,name=exclude_payload,json=excludePayload,proto3" json:"exclude_payload,omitempty"`
	ZstdCompression bool                 `protobuf:"varint,2,opt,name=zstd_compression,json=zstdCompression,proto3" json:"zstd_compression,omitempty"`
	RequireFormat   *BlockMessage_Format `protobuf:"varint,3,opt,name=require_format,json=requireFormat,proto3,enum=BlockMessage_Format,oneof" json:"require_format,omitempty"`
}

func (x *BlockMessageDeliverySettings) Reset() {
	*x = BlockMessageDeliverySettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocksapi_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockMessageDeliverySettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockMessageDeliverySettings) ProtoMessage() {}

func (x *BlockMessageDeliverySettings) ProtoReflect() protoreflect.Message {
	mi := &file_blocksapi_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockMessageDeliverySettings.ProtoReflect.Descriptor instead.
func (*BlockMessageDeliverySettings) Descriptor() ([]byte, []int) {
	return file_blocksapi_message_proto_rawDescGZIP(), []int{1}
}

func (x *BlockMessageDeliverySettings) GetExcludePayload() bool {
	if x != nil {
		return x.ExcludePayload
	}
	return false
}

func (x *BlockMessageDeliverySettings) GetZstdCompression() bool {
	if x != nil {
		return x.ZstdCompression
	}
	return false
}

func (x *BlockMessageDeliverySettings) GetRequireFormat() BlockMessage_Format {
	if x != nil && x.RequireFormat != nil {
		return *x.RequireFormat
	}
	return BlockMessage_PAYLOAD_NEAR_BLOCK_V2
}

type BlockMessage_ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind    BlockMessage_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=BlockMessage_Kind" json:"kind,omitempty"`
	Height  uint64            `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	ShardId uint64            `protobuf:"varint,3,opt,name=shard_id,json=shardId,proto3" json:"shard_id,omitempty"`
}

func (x *BlockMessage_ID) Reset() {
	*x = BlockMessage_ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocksapi_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockMessage_ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockMessage_ID) ProtoMessage() {}

func (x *BlockMessage_ID) ProtoReflect() protoreflect.Message {
	mi := &file_blocksapi_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockMessage_ID.ProtoReflect.Descriptor instead.
func (*BlockMessage_ID) Descriptor() ([]byte, []int) {
	return file_blocksapi_message_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BlockMessage_ID) GetKind() BlockMessage_Kind {
	if x != nil {
		return x.Kind
	}
	return BlockMessage_MSG_WHOLE
}

func (x *BlockMessage_ID) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *BlockMessage_ID) GetShardId() uint64 {
	if x != nil {
		return x.ShardId
	}
	return 0
}

var File_blocksapi_message_proto protoreflect.FileDescriptor

var file_blocksapi_message_proto_rawDesc = []byte{
	0x0a, 0x17, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x03, 0x0a, 0x0c, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x06,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x7a, 0x73,
	0x74, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x7a, 0x73, 0x74, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0b, 0x72, 0x61, 0x77, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0a, 0x72, 0x61, 0x77, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x5f, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x73, 0x68, 0x61, 0x72, 0x64, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12,
	0x0d, 0x0a, 0x09, 0x4d, 0x53, 0x47, 0x5f, 0x57, 0x48, 0x4f, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0e,
	0x0a, 0x0a, 0x4d, 0x53, 0x47, 0x5f, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0d,
	0x0a, 0x09, 0x4d, 0x53, 0x47, 0x5f, 0x43, 0x48, 0x55, 0x4e, 0x4b, 0x10, 0x02, 0x22, 0x5b, 0x0a,
	0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x41, 0x59, 0x4c, 0x4f,
	0x41, 0x44, 0x5f, 0x4e, 0x45, 0x41, 0x52, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x56, 0x32,
	0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x41, 0x59, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x41, 0x55,
	0x52, 0x4f, 0x52, 0x41, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x56, 0x32, 0x10, 0x01, 0x12,
	0x19, 0x0a, 0x15, 0x50, 0x41, 0x59, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x4e, 0x45, 0x41, 0x52, 0x5f,
	0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x56, 0x33, 0x10, 0x02, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xc7, 0x01, 0x0a, 0x1c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0e, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x29, 0x0a, 0x10, 0x7a, 0x73, 0x74, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x7a, 0x73, 0x74, 0x64, 0x43,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0e, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x14, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x42,
	0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75,
	0x72, 0x6f, 0x72, 0x61, 0x2d, 0x69, 0x73, 0x2d, 0x6e, 0x65, 0x61, 0x72, 0x2f, 0x62, 0x6f, 0x72,
	0x65, 0x61, 0x6c, 0x69, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x3b, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blocksapi_message_proto_rawDescOnce sync.Once
	file_blocksapi_message_proto_rawDescData = file_blocksapi_message_proto_rawDesc
)

func file_blocksapi_message_proto_rawDescGZIP() []byte {
	file_blocksapi_message_proto_rawDescOnce.Do(func() {
		file_blocksapi_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_blocksapi_message_proto_rawDescData)
	})
	return file_blocksapi_message_proto_rawDescData
}

var file_blocksapi_message_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_blocksapi_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_blocksapi_message_proto_goTypes = []interface{}{
	(BlockMessage_Kind)(0),               // 0: BlockMessage.Kind
	(BlockMessage_Format)(0),             // 1: BlockMessage.Format
	(*BlockMessage)(nil),                 // 2: BlockMessage
	(*BlockMessageDeliverySettings)(nil), // 3: BlockMessageDeliverySettings
	(*BlockMessage_ID)(nil),              // 4: BlockMessage.ID
}
var file_blocksapi_message_proto_depIdxs = []int32{
	4, // 0: BlockMessage.id:type_name -> BlockMessage.ID
	1, // 1: BlockMessage.format:type_name -> BlockMessage.Format
	1, // 2: BlockMessageDeliverySettings.require_format:type_name -> BlockMessage.Format
	0, // 3: BlockMessage.ID.kind:type_name -> BlockMessage.Kind
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_blocksapi_message_proto_init() }
func file_blocksapi_message_proto_init() {
	if File_blocksapi_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blocksapi_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockMessage); i {
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
		file_blocksapi_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockMessageDeliverySettings); i {
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
		file_blocksapi_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockMessage_ID); i {
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
	file_blocksapi_message_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*BlockMessage_RawPayload)(nil),
	}
	file_blocksapi_message_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blocksapi_message_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blocksapi_message_proto_goTypes,
		DependencyIndexes: file_blocksapi_message_proto_depIdxs,
		EnumInfos:         file_blocksapi_message_proto_enumTypes,
		MessageInfos:      file_blocksapi_message_proto_msgTypes,
	}.Build()
	File_blocksapi_message_proto = out.File
	file_blocksapi_message_proto_rawDesc = nil
	file_blocksapi_message_proto_goTypes = nil
	file_blocksapi_message_proto_depIdxs = nil
}
