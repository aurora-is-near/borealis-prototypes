// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v3.21.12
// source: blocksapi/services.proto

package blocksapi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_blocksapi_services_proto protoreflect.FileDescriptor

var file_blocksapi_services_proto_rawDesc = []byte{
	0x0a, 0x18, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe7, 0x01,
	0x0a, 0x0e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x12, 0x47, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x12, 0x18, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x10, 0x52, 0x65, 0x61,
	0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x72, 0x6f, 0x72, 0x61, 0x2d, 0x69, 0x73, 0x2d,
	0x6e, 0x65, 0x61, 0x72, 0x2f, 0x62, 0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x61, 0x70, 0x69, 0x3b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_blocksapi_services_proto_goTypes = []interface{}{
	(*ListBlockStreamsRequest)(nil),  // 0: ListBlockStreamsRequest
	(*ReadBlockMessageRequest)(nil),  // 1: ReadBlockMessageRequest
	(*GetBlockStreamRequest)(nil),    // 2: GetBlockStreamRequest
	(*ListBlockStreamsResponse)(nil), // 3: ListBlockStreamsResponse
	(*ReadBlockMessageResponse)(nil), // 4: ReadBlockMessageResponse
	(*GetBlockStreamResponse)(nil),   // 5: GetBlockStreamResponse
}
var file_blocksapi_services_proto_depIdxs = []int32{
	0, // 0: BlocksProvider.ListBlockStreams:input_type -> ListBlockStreamsRequest
	1, // 1: BlocksProvider.ReadBlockMessage:input_type -> ReadBlockMessageRequest
	2, // 2: BlocksProvider.GetBlockStream:input_type -> GetBlockStreamRequest
	3, // 3: BlocksProvider.ListBlockStreams:output_type -> ListBlockStreamsResponse
	4, // 4: BlocksProvider.ReadBlockMessage:output_type -> ReadBlockMessageResponse
	5, // 5: BlocksProvider.GetBlockStream:output_type -> GetBlockStreamResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_blocksapi_services_proto_init() }
func file_blocksapi_services_proto_init() {
	if File_blocksapi_services_proto != nil {
		return
	}
	file_blocksapi_read_message_proto_init()
	file_blocksapi_stream_list_proto_init()
	file_blocksapi_stream_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blocksapi_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blocksapi_services_proto_goTypes,
		DependencyIndexes: file_blocksapi_services_proto_depIdxs,
	}.Build()
	File_blocksapi_services_proto = out.File
	file_blocksapi_services_proto_rawDesc = nil
	file_blocksapi_services_proto_goTypes = nil
	file_blocksapi_services_proto_depIdxs = nil
}
