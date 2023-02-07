// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: payloads/near_block/transaction/transaction.proto

package transaction

import (
	common "github.com/aurora-is-near/borealis-prototypes/go/common"
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

type TransactionWithOutcome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *SignedTransactionView               `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Outcome     *ExecutionOutcomeWithOptionalReceipt `protobuf:"bytes,2,opt,name=outcome,proto3" json:"outcome,omitempty"`
}

func (x *TransactionWithOutcome) Reset() {
	*x = TransactionWithOutcome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_block_transaction_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionWithOutcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionWithOutcome) ProtoMessage() {}

func (x *TransactionWithOutcome) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_block_transaction_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionWithOutcome.ProtoReflect.Descriptor instead.
func (*TransactionWithOutcome) Descriptor() ([]byte, []int) {
	return file_payloads_near_block_transaction_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionWithOutcome) GetTransaction() *SignedTransactionView {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *TransactionWithOutcome) GetOutcome() *ExecutionOutcomeWithOptionalReceipt {
	if x != nil {
		return x.Outcome
	}
	return nil
}

type SignedTransactionView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignerId   string            `protobuf:"bytes,1,opt,name=signer_id,json=signerId,proto3" json:"signer_id,omitempty"`
	PublicKey  *common.PublicKey `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Nonce      uint64            `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ReceiverId string            `protobuf:"bytes,4,opt,name=receiver_id,json=receiverId,proto3" json:"receiver_id,omitempty"`
	Actions    []*ActionView     `protobuf:"bytes,5,rep,name=actions,proto3" json:"actions,omitempty"`
	Signature  *common.Signature `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	H256Hash   []byte            `protobuf:"bytes,7,opt,name=h256_hash,json=h256Hash,proto3" json:"h256_hash,omitempty"`
}

func (x *SignedTransactionView) Reset() {
	*x = SignedTransactionView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_block_transaction_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedTransactionView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedTransactionView) ProtoMessage() {}

func (x *SignedTransactionView) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_block_transaction_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedTransactionView.ProtoReflect.Descriptor instead.
func (*SignedTransactionView) Descriptor() ([]byte, []int) {
	return file_payloads_near_block_transaction_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *SignedTransactionView) GetSignerId() string {
	if x != nil {
		return x.SignerId
	}
	return ""
}

func (x *SignedTransactionView) GetPublicKey() *common.PublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *SignedTransactionView) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *SignedTransactionView) GetReceiverId() string {
	if x != nil {
		return x.ReceiverId
	}
	return ""
}

func (x *SignedTransactionView) GetActions() []*ActionView {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *SignedTransactionView) GetSignature() *common.Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *SignedTransactionView) GetH256Hash() []byte {
	if x != nil {
		return x.H256Hash
	}
	return nil
}

var File_payloads_near_block_transaction_transaction_proto protoreflect.FileDescriptor

var file_payloads_near_block_transaction_transaction_proto_rawDesc = []byte{
	0x0a, 0x31, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2f, 0x6e, 0x65, 0x61, 0x72, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x73, 0x2f, 0x6e, 0x65, 0x61, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x2f, 0x6e, 0x65, 0x61, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x16, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x75, 0x74, 0x63,
	0x6f, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x69, 0x65, 0x77,
	0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a,
	0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d,
	0x65, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x22, 0x84, 0x02,
	0x0a, 0x15, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x56, 0x69, 0x65, 0x77, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x56, 0x69, 0x65, 0x77, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x32, 0x35, 0x36, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x68, 0x32, 0x35, 0x36,
	0x48, 0x61, 0x73, 0x68, 0x42, 0x5e, 0x5a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x72, 0x6f, 0x72, 0x61, 0x2d, 0x69, 0x73, 0x2d, 0x6e, 0x65, 0x61,
	0x72, 0x2f, 0x62, 0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x73, 0x2f, 0x6e, 0x65, 0x61, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payloads_near_block_transaction_transaction_proto_rawDescOnce sync.Once
	file_payloads_near_block_transaction_transaction_proto_rawDescData = file_payloads_near_block_transaction_transaction_proto_rawDesc
)

func file_payloads_near_block_transaction_transaction_proto_rawDescGZIP() []byte {
	file_payloads_near_block_transaction_transaction_proto_rawDescOnce.Do(func() {
		file_payloads_near_block_transaction_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_payloads_near_block_transaction_transaction_proto_rawDescData)
	})
	return file_payloads_near_block_transaction_transaction_proto_rawDescData
}

var file_payloads_near_block_transaction_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_payloads_near_block_transaction_transaction_proto_goTypes = []interface{}{
	(*TransactionWithOutcome)(nil),              // 0: TransactionWithOutcome
	(*SignedTransactionView)(nil),               // 1: SignedTransactionView
	(*ExecutionOutcomeWithOptionalReceipt)(nil), // 2: ExecutionOutcomeWithOptionalReceipt
	(*common.PublicKey)(nil),                    // 3: PublicKey
	(*ActionView)(nil),                          // 4: ActionView
	(*common.Signature)(nil),                    // 5: Signature
}
var file_payloads_near_block_transaction_transaction_proto_depIdxs = []int32{
	1, // 0: TransactionWithOutcome.transaction:type_name -> SignedTransactionView
	2, // 1: TransactionWithOutcome.outcome:type_name -> ExecutionOutcomeWithOptionalReceipt
	3, // 2: SignedTransactionView.public_key:type_name -> PublicKey
	4, // 3: SignedTransactionView.actions:type_name -> ActionView
	5, // 4: SignedTransactionView.signature:type_name -> Signature
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_payloads_near_block_transaction_transaction_proto_init() }
func file_payloads_near_block_transaction_transaction_proto_init() {
	if File_payloads_near_block_transaction_transaction_proto != nil {
		return
	}
	file_payloads_near_block_transaction_execution_proto_init()
	file_payloads_near_block_transaction_receipt_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_payloads_near_block_transaction_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionWithOutcome); i {
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
		file_payloads_near_block_transaction_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedTransactionView); i {
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
			RawDescriptor: file_payloads_near_block_transaction_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payloads_near_block_transaction_transaction_proto_goTypes,
		DependencyIndexes: file_payloads_near_block_transaction_transaction_proto_depIdxs,
		MessageInfos:      file_payloads_near_block_transaction_transaction_proto_msgTypes,
	}.Build()
	File_payloads_near_block_transaction_transaction_proto = out.File
	file_payloads_near_block_transaction_transaction_proto_rawDesc = nil
	file_payloads_near_block_transaction_transaction_proto_goTypes = nil
	file_payloads_near_block_transaction_transaction_proto_depIdxs = nil
}
