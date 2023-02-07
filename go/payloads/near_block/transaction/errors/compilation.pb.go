// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: payloads/near_block/transaction/errors/compilation.proto

package errors

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

// Error that can occur while preparing or executing Wasm smart-contract.
type PrepareError int32

const (
	// Error happened while serializing the module.
	PrepareError_SERIALIZATION PrepareError = 0
	// Error happened while deserializing the module.
	PrepareError_DESERIALIZATION PrepareError = 1
	// Internal memory declaration has been found in the module.
	PrepareError_INTERNAL_MEMORY_DECLARED PrepareError = 2
	// Gas instrumentation failed.
	//
	// This most likely indicates the module isn't valid.
	PrepareError_GAS_INSTRUMENTATION PrepareError = 3
	// Stack instrumentation failed.
	//
	// This  most likely indicates the module isn't valid.
	PrepareError_STACK_HEIGHT_INSTRUMENTATION PrepareError = 4
	// Error happened during instantiation.
	//
	// This might indicate that `start` function trapped, or module isn't
	// instantiable and/or unlinkable.
	PrepareError_INSTANTIATE PrepareError = 5
	// Error creating memory.
	PrepareError_MEMORY PrepareError = 6
	// Contract contains too many functions.
	PrepareError_TOO_MANY_FUNCTIONS PrepareError = 7
	// Contract contains too many locals.
	PrepareError_TOO_MANY_LOCALS PrepareError = 8
)

// Enum value maps for PrepareError.
var (
	PrepareError_name = map[int32]string{
		0: "SERIALIZATION",
		1: "DESERIALIZATION",
		2: "INTERNAL_MEMORY_DECLARED",
		3: "GAS_INSTRUMENTATION",
		4: "STACK_HEIGHT_INSTRUMENTATION",
		5: "INSTANTIATE",
		6: "MEMORY",
		7: "TOO_MANY_FUNCTIONS",
		8: "TOO_MANY_LOCALS",
	}
	PrepareError_value = map[string]int32{
		"SERIALIZATION":                0,
		"DESERIALIZATION":              1,
		"INTERNAL_MEMORY_DECLARED":     2,
		"GAS_INSTRUMENTATION":          3,
		"STACK_HEIGHT_INSTRUMENTATION": 4,
		"INSTANTIATE":                  5,
		"MEMORY":                       6,
		"TOO_MANY_FUNCTIONS":           7,
		"TOO_MANY_LOCALS":              8,
	}
)

func (x PrepareError) Enum() *PrepareError {
	p := new(PrepareError)
	*p = x
	return p
}

func (x PrepareError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PrepareError) Descriptor() protoreflect.EnumDescriptor {
	return file_payloads_near_block_transaction_errors_compilation_proto_enumTypes[0].Descriptor()
}

func (PrepareError) Type() protoreflect.EnumType {
	return &file_payloads_near_block_transaction_errors_compilation_proto_enumTypes[0]
}

func (x PrepareError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PrepareError.Descriptor instead.
func (PrepareError) EnumDescriptor() ([]byte, []int) {
	return file_payloads_near_block_transaction_errors_compilation_proto_rawDescGZIP(), []int{0}
}

type CompilationError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Variant:
	//
	//	*CompilationError_CodeDoesNotExist_
	//	*CompilationError_PrepareError_
	//	*CompilationError_WasmerCompileError_
	//	*CompilationError_UnsupportedCompiler_
	Variant isCompilationError_Variant `protobuf_oneof:"variant"`
}

func (x *CompilationError) Reset() {
	*x = CompilationError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_block_transaction_errors_compilation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompilationError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompilationError) ProtoMessage() {}

func (x *CompilationError) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_block_transaction_errors_compilation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompilationError.ProtoReflect.Descriptor instead.
func (*CompilationError) Descriptor() ([]byte, []int) {
	return file_payloads_near_block_transaction_errors_compilation_proto_rawDescGZIP(), []int{0}
}

func (m *CompilationError) GetVariant() isCompilationError_Variant {
	if m != nil {
		return m.Variant
	}
	return nil
}

func (x *CompilationError) GetCodeDoesNotExist() *CompilationError_CodeDoesNotExist {
	if x, ok := x.GetVariant().(*CompilationError_CodeDoesNotExist_); ok {
		return x.CodeDoesNotExist
	}
	return nil
}

func (x *CompilationError) GetPrepareError() *CompilationError_PrepareError {
	if x, ok := x.GetVariant().(*CompilationError_PrepareError_); ok {
		return x.PrepareError
	}
	return nil
}

func (x *CompilationError) GetWasmerCompileError() *CompilationError_WasmerCompileError {
	if x, ok := x.GetVariant().(*CompilationError_WasmerCompileError_); ok {
		return x.WasmerCompileError
	}
	return nil
}

func (x *CompilationError) GetUnsupportedCompiler() *CompilationError_UnsupportedCompiler {
	if x, ok := x.GetVariant().(*CompilationError_UnsupportedCompiler_); ok {
		return x.UnsupportedCompiler
	}
	return nil
}

type isCompilationError_Variant interface {
	isCompilationError_Variant()
}

type CompilationError_CodeDoesNotExist_ struct {
	CodeDoesNotExist *CompilationError_CodeDoesNotExist `protobuf:"bytes,1,opt,name=code_does_not_exist,json=codeDoesNotExist,proto3,oneof"`
}

type CompilationError_PrepareError_ struct {
	PrepareError *CompilationError_PrepareError `protobuf:"bytes,2,opt,name=prepare_error,json=prepareError,proto3,oneof"`
}

type CompilationError_WasmerCompileError_ struct {
	WasmerCompileError *CompilationError_WasmerCompileError `protobuf:"bytes,3,opt,name=wasmer_compile_error,json=wasmerCompileError,proto3,oneof"`
}

type CompilationError_UnsupportedCompiler_ struct {
	UnsupportedCompiler *CompilationError_UnsupportedCompiler `protobuf:"bytes,4,opt,name=unsupported_compiler,json=unsupportedCompiler,proto3,oneof"`
}

func (*CompilationError_CodeDoesNotExist_) isCompilationError_Variant() {}

func (*CompilationError_PrepareError_) isCompilationError_Variant() {}

func (*CompilationError_WasmerCompileError_) isCompilationError_Variant() {}

func (*CompilationError_UnsupportedCompiler_) isCompilationError_Variant() {}

type CompilationError_CodeDoesNotExist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *CompilationError_CodeDoesNotExist) Reset() {
	*x = CompilationError_CodeDoesNotExist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_block_transaction_errors_compilation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompilationError_CodeDoesNotExist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompilationError_CodeDoesNotExist) ProtoMessage() {}

func (x *CompilationError_CodeDoesNotExist) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_block_transaction_errors_compilation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompilationError_CodeDoesNotExist.ProtoReflect.Descriptor instead.
func (*CompilationError_CodeDoesNotExist) Descriptor() ([]byte, []int) {
	return file_payloads_near_block_transaction_errors_compilation_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CompilationError_CodeDoesNotExist) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type CompilationError_PrepareError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error PrepareError `protobuf:"varint,1,opt,name=error,proto3,enum=PrepareError" json:"error,omitempty"`
}

func (x *CompilationError_PrepareError) Reset() {
	*x = CompilationError_PrepareError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_block_transaction_errors_compilation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompilationError_PrepareError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompilationError_PrepareError) ProtoMessage() {}

func (x *CompilationError_PrepareError) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_block_transaction_errors_compilation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompilationError_PrepareError.ProtoReflect.Descriptor instead.
func (*CompilationError_PrepareError) Descriptor() ([]byte, []int) {
	return file_payloads_near_block_transaction_errors_compilation_proto_rawDescGZIP(), []int{0, 1}
}

func (x *CompilationError_PrepareError) GetError() PrepareError {
	if x != nil {
		return x.Error
	}
	return PrepareError_SERIALIZATION
}

type CompilationError_WasmerCompileError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CompilationError_WasmerCompileError) Reset() {
	*x = CompilationError_WasmerCompileError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_block_transaction_errors_compilation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompilationError_WasmerCompileError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompilationError_WasmerCompileError) ProtoMessage() {}

func (x *CompilationError_WasmerCompileError) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_block_transaction_errors_compilation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompilationError_WasmerCompileError.ProtoReflect.Descriptor instead.
func (*CompilationError_WasmerCompileError) Descriptor() ([]byte, []int) {
	return file_payloads_near_block_transaction_errors_compilation_proto_rawDescGZIP(), []int{0, 2}
}

func (x *CompilationError_WasmerCompileError) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type CompilationError_UnsupportedCompiler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CompilationError_UnsupportedCompiler) Reset() {
	*x = CompilationError_UnsupportedCompiler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payloads_near_block_transaction_errors_compilation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompilationError_UnsupportedCompiler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompilationError_UnsupportedCompiler) ProtoMessage() {}

func (x *CompilationError_UnsupportedCompiler) ProtoReflect() protoreflect.Message {
	mi := &file_payloads_near_block_transaction_errors_compilation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompilationError_UnsupportedCompiler.ProtoReflect.Descriptor instead.
func (*CompilationError_UnsupportedCompiler) Descriptor() ([]byte, []int) {
	return file_payloads_near_block_transaction_errors_compilation_proto_rawDescGZIP(), []int{0, 3}
}

func (x *CompilationError_UnsupportedCompiler) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_payloads_near_block_transaction_errors_compilation_proto protoreflect.FileDescriptor

var file_payloads_near_block_transaction_errors_compilation_proto_rawDesc = []byte{
	0x0a, 0x38, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2f, 0x6e, 0x65, 0x61, 0x72, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x04, 0x0a, 0x10, 0x43,
	0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x53, 0x0a, 0x13, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x64, 0x6f, 0x65, 0x73, 0x5f, 0x6e, 0x6f, 0x74,
	0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x44, 0x6f, 0x65, 0x73, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x48, 0x00, 0x52, 0x10, 0x63, 0x6f, 0x64, 0x65, 0x44, 0x6f, 0x65, 0x73, 0x4e, 0x6f, 0x74, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x5f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x50,
	0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x0c, 0x70,
	0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x58, 0x0a, 0x14, 0x77,
	0x61, 0x73, 0x6d, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x57, 0x61, 0x73,
	0x6d, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48,
	0x00, 0x52, 0x12, 0x77, 0x61, 0x73, 0x6d, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x5a, 0x0a, 0x14, 0x75, 0x6e, 0x73, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x55, 0x6e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x48, 0x00, 0x52, 0x13, 0x75, 0x6e,
	0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65,
	0x72, 0x1a, 0x31, 0x0a, 0x10, 0x43, 0x6f, 0x64, 0x65, 0x44, 0x6f, 0x65, 0x73, 0x4e, 0x6f, 0x74,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x1a, 0x33, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x26, 0x0a, 0x12, 0x57, 0x61, 0x73,
	0x6d, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x1a, 0x27, 0x0a, 0x13, 0x55, 0x6e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64,
	0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x09, 0x0a, 0x07, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x2a, 0xd9, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72,
	0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x45, 0x52, 0x49, 0x41, 0x4c,
	0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x45, 0x53,
	0x45, 0x52, 0x49, 0x41, 0x4c, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x1c,
	0x0a, 0x18, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x45, 0x4d, 0x4f, 0x52,
	0x59, 0x5f, 0x44, 0x45, 0x43, 0x4c, 0x41, 0x52, 0x45, 0x44, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13,
	0x47, 0x41, 0x53, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x48,
	0x45, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d, 0x45, 0x4e, 0x54,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x53, 0x54, 0x41,
	0x4e, 0x54, 0x49, 0x41, 0x54, 0x45, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x45, 0x4d, 0x4f,
	0x52, 0x59, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59,
	0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f,
	0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x53, 0x10,
	0x08, 0x42, 0x60, 0x5a, 0x5e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x75, 0x72, 0x6f, 0x72, 0x61, 0x2d, 0x69, 0x73, 0x2d, 0x6e, 0x65, 0x61, 0x72, 0x2f, 0x62,
	0x6f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2f, 0x6e,
	0x65, 0x61, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payloads_near_block_transaction_errors_compilation_proto_rawDescOnce sync.Once
	file_payloads_near_block_transaction_errors_compilation_proto_rawDescData = file_payloads_near_block_transaction_errors_compilation_proto_rawDesc
)

func file_payloads_near_block_transaction_errors_compilation_proto_rawDescGZIP() []byte {
	file_payloads_near_block_transaction_errors_compilation_proto_rawDescOnce.Do(func() {
		file_payloads_near_block_transaction_errors_compilation_proto_rawDescData = protoimpl.X.CompressGZIP(file_payloads_near_block_transaction_errors_compilation_proto_rawDescData)
	})
	return file_payloads_near_block_transaction_errors_compilation_proto_rawDescData
}

var file_payloads_near_block_transaction_errors_compilation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_payloads_near_block_transaction_errors_compilation_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_payloads_near_block_transaction_errors_compilation_proto_goTypes = []interface{}{
	(PrepareError)(0),                            // 0: PrepareError
	(*CompilationError)(nil),                     // 1: CompilationError
	(*CompilationError_CodeDoesNotExist)(nil),    // 2: CompilationError.CodeDoesNotExist
	(*CompilationError_PrepareError)(nil),        // 3: CompilationError.PrepareError
	(*CompilationError_WasmerCompileError)(nil),  // 4: CompilationError.WasmerCompileError
	(*CompilationError_UnsupportedCompiler)(nil), // 5: CompilationError.UnsupportedCompiler
}
var file_payloads_near_block_transaction_errors_compilation_proto_depIdxs = []int32{
	2, // 0: CompilationError.code_does_not_exist:type_name -> CompilationError.CodeDoesNotExist
	3, // 1: CompilationError.prepare_error:type_name -> CompilationError.PrepareError
	4, // 2: CompilationError.wasmer_compile_error:type_name -> CompilationError.WasmerCompileError
	5, // 3: CompilationError.unsupported_compiler:type_name -> CompilationError.UnsupportedCompiler
	0, // 4: CompilationError.PrepareError.error:type_name -> PrepareError
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_payloads_near_block_transaction_errors_compilation_proto_init() }
func file_payloads_near_block_transaction_errors_compilation_proto_init() {
	if File_payloads_near_block_transaction_errors_compilation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payloads_near_block_transaction_errors_compilation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompilationError); i {
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
		file_payloads_near_block_transaction_errors_compilation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompilationError_CodeDoesNotExist); i {
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
		file_payloads_near_block_transaction_errors_compilation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompilationError_PrepareError); i {
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
		file_payloads_near_block_transaction_errors_compilation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompilationError_WasmerCompileError); i {
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
		file_payloads_near_block_transaction_errors_compilation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompilationError_UnsupportedCompiler); i {
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
	file_payloads_near_block_transaction_errors_compilation_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CompilationError_CodeDoesNotExist_)(nil),
		(*CompilationError_PrepareError_)(nil),
		(*CompilationError_WasmerCompileError_)(nil),
		(*CompilationError_UnsupportedCompiler_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_payloads_near_block_transaction_errors_compilation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payloads_near_block_transaction_errors_compilation_proto_goTypes,
		DependencyIndexes: file_payloads_near_block_transaction_errors_compilation_proto_depIdxs,
		EnumInfos:         file_payloads_near_block_transaction_errors_compilation_proto_enumTypes,
		MessageInfos:      file_payloads_near_block_transaction_errors_compilation_proto_msgTypes,
	}.Build()
	File_payloads_near_block_transaction_errors_compilation_proto = out.File
	file_payloads_near_block_transaction_errors_compilation_proto_rawDesc = nil
	file_payloads_near_block_transaction_errors_compilation_proto_goTypes = nil
	file_payloads_near_block_transaction_errors_compilation_proto_depIdxs = nil
}
