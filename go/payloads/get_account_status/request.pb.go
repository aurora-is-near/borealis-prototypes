// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: payloads/get_account_status/request.proto

package getaccountstatus

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

type GetAccountStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *GetAccountStatusRequest) Reset() {
	*x = GetAccountStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_get_account_status_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountStatusRequest) ProtoMessage() {}

func (x *GetAccountStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_get_account_status_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountStatusRequest.ProtoReflect.Descriptor instead.
func (*GetAccountStatusRequest) Descriptor() ([]byte, []int) {
	return file_payloads_get_account_status_request_proto_rawDescGZIP(), []int{0}
}

func (x *GetAccountStatusRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

var File_payloads_get_account_status_request_proto protoreflect.FileDescriptor

var file_payloads_get_account_status_request_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x5f, 0x5a, 0x5d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x75, 0x72, 0x6f, 0x72, 0x61, 0x2d, 0x69, 0x73, 0x2d, 0x6e, 0x65, 0x61, 0x72, 0x2f, 0x62, 0x6f,
	0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x65,
	0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x3b, 0x67, 0x65, 0x74, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payloads_get_account_status_request_proto_rawDescOnce sync.Once
	file_payloads_get_account_status_request_proto_rawDescData = file_payloads_get_account_status_request_proto_rawDesc
)

func file_payloads_get_account_status_request_proto_rawDescGZIP() []byte {
	file_payloads_get_account_status_request_proto_rawDescOnce.Do(func() {
		file_payloads_get_account_status_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_payloads_get_account_status_request_proto_rawDescData)
	})
	return file_payloads_get_account_status_request_proto_rawDescData
}

var file_payloads_get_account_status_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_payloads_get_account_status_request_proto_goTypes = []interface{}{
	(*GetAccountStatusRequest)(nil), // 0: GetAccountStatusRequest
}
var file_payloads_get_account_status_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_payloads_get_account_status_request_proto_init() }
func file_payloads_get_account_status_request_proto_init() {
	if File_payloads_get_account_status_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payloads_get_account_status_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountStatusRequest); i {
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
			RawDescriptor: file_payloads_get_account_status_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payloads_get_account_status_request_proto_goTypes,
		DependencyIndexes: file_payloads_get_account_status_request_proto_depIdxs,
		MessageInfos:      file_payloads_get_account_status_request_proto_msgTypes,
	}.Build()
	File_payloads_get_account_status_request_proto = out.File
	file_payloads_get_account_status_request_proto_rawDesc = nil
	file_payloads_get_account_status_request_proto_goTypes = nil
	file_payloads_get_account_status_request_proto_depIdxs = nil
}