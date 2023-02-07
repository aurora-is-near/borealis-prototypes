// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.4.0
// source: message.proto

package borealisproto

import (
	fmt "fmt"
	get_account_status "github.com/aurora-is-near/borealis-prototypes/go/payloads/get_account_status"
	near_block "github.com/aurora-is-near/borealis-prototypes/go/payloads/near_block"
	rpc "github.com/aurora-is-near/borealis-prototypes/go/payloads/rpc"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	bits "math/bits"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *Message) CloneVT() *Message {
	if m == nil {
		return (*Message)(nil)
	}
	r := &Message{
		Version: m.Version,
		Id:      m.Id,
	}
	if m.Payload != nil {
		r.Payload = m.Payload.(interface{ CloneVT() isMessage_Payload }).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *Message) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *Message_Generic) CloneVT() isMessage_Payload {
	if m == nil {
		return (*Message_Generic)(nil)
	}
	r := &Message_Generic{}
	if rhs := m.Generic; rhs != nil {
		tmpBytes := make([]byte, len(rhs))
		copy(tmpBytes, rhs)
		r.Generic = tmpBytes
	}
	return r
}

func (m *Message_RpcRequest) CloneVT() isMessage_Payload {
	if m == nil {
		return (*Message_RpcRequest)(nil)
	}
	r := &Message_RpcRequest{
		RpcRequest: m.RpcRequest.CloneVT(),
	}
	return r
}

func (m *Message_RpcResponse) CloneVT() isMessage_Payload {
	if m == nil {
		return (*Message_RpcResponse)(nil)
	}
	r := &Message_RpcResponse{
		RpcResponse: m.RpcResponse.CloneVT(),
	}
	return r
}

func (m *Message_GetAccountStatusRequest) CloneVT() isMessage_Payload {
	if m == nil {
		return (*Message_GetAccountStatusRequest)(nil)
	}
	r := &Message_GetAccountStatusRequest{
		GetAccountStatusRequest: m.GetAccountStatusRequest.CloneVT(),
	}
	return r
}

func (m *Message_GetAccountStatusResponse) CloneVT() isMessage_Payload {
	if m == nil {
		return (*Message_GetAccountStatusResponse)(nil)
	}
	r := &Message_GetAccountStatusResponse{
		GetAccountStatusResponse: m.GetAccountStatusResponse.CloneVT(),
	}
	return r
}

func (m *Message_NearBlockHeader) CloneVT() isMessage_Payload {
	if m == nil {
		return (*Message_NearBlockHeader)(nil)
	}
	r := &Message_NearBlockHeader{
		NearBlockHeader: m.NearBlockHeader.CloneVT(),
	}
	return r
}

func (m *Message_NearBlockShard) CloneVT() isMessage_Payload {
	if m == nil {
		return (*Message_NearBlockShard)(nil)
	}
	r := &Message_NearBlockShard{
		NearBlockShard: m.NearBlockShard.CloneVT(),
	}
	return r
}

func (this *Message) EqualVT(that *Message) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Payload == nil && that.Payload != nil {
		return false
	} else if this.Payload != nil {
		if that.Payload == nil {
			return false
		}
		if !this.Payload.(interface{ EqualVT(isMessage_Payload) bool }).EqualVT(that.Payload) {
			return false
		}
	}
	if this.Version != that.Version {
		return false
	}
	if this.Id != that.Id {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Message) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Message)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *Message_Generic) EqualVT(thatIface isMessage_Payload) bool {
	that, ok := thatIface.(*Message_Generic)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if string(this.Generic) != string(that.Generic) {
		return false
	}
	return true
}

func (this *Message_RpcRequest) EqualVT(thatIface isMessage_Payload) bool {
	that, ok := thatIface.(*Message_RpcRequest)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.RpcRequest, that.RpcRequest; p != q {
		if p == nil {
			p = &rpc.RpcMessage{}
		}
		if q == nil {
			q = &rpc.RpcMessage{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *Message_RpcResponse) EqualVT(thatIface isMessage_Payload) bool {
	that, ok := thatIface.(*Message_RpcResponse)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.RpcResponse, that.RpcResponse; p != q {
		if p == nil {
			p = &rpc.RpcMessage{}
		}
		if q == nil {
			q = &rpc.RpcMessage{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *Message_GetAccountStatusRequest) EqualVT(thatIface isMessage_Payload) bool {
	that, ok := thatIface.(*Message_GetAccountStatusRequest)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.GetAccountStatusRequest, that.GetAccountStatusRequest; p != q {
		if p == nil {
			p = &get_account_status.GetAccountStatusRequest{}
		}
		if q == nil {
			q = &get_account_status.GetAccountStatusRequest{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *Message_GetAccountStatusResponse) EqualVT(thatIface isMessage_Payload) bool {
	that, ok := thatIface.(*Message_GetAccountStatusResponse)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.GetAccountStatusResponse, that.GetAccountStatusResponse; p != q {
		if p == nil {
			p = &get_account_status.GetAccountStatusResponse{}
		}
		if q == nil {
			q = &get_account_status.GetAccountStatusResponse{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *Message_NearBlockHeader) EqualVT(thatIface isMessage_Payload) bool {
	that, ok := thatIface.(*Message_NearBlockHeader)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.NearBlockHeader, that.NearBlockHeader; p != q {
		if p == nil {
			p = &near_block.BlockHeaderView{}
		}
		if q == nil {
			q = &near_block.BlockHeaderView{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *Message_NearBlockShard) EqualVT(thatIface isMessage_Payload) bool {
	that, ok := thatIface.(*Message_NearBlockShard)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.NearBlockShard, that.NearBlockShard; p != q {
		if p == nil {
			p = &near_block.BlockShard{}
		}
		if q == nil {
			q = &near_block.BlockShard{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (m *Message) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Message) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if vtmsg, ok := m.Payload.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarint(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if m.Version != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Message_Generic) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Message_Generic) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Generic)
	copy(dAtA[i:], m.Generic)
	i = encodeVarint(dAtA, i, uint64(len(m.Generic)))
	i--
	dAtA[i] = 0x1a
	return len(dAtA) - i, nil
}
func (m *Message_RpcRequest) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Message_RpcRequest) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.RpcRequest != nil {
		size, err := m.RpcRequest.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *Message_RpcResponse) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Message_RpcResponse) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.RpcResponse != nil {
		size, err := m.RpcResponse.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *Message_GetAccountStatusRequest) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Message_GetAccountStatusRequest) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.GetAccountStatusRequest != nil {
		size, err := m.GetAccountStatusRequest.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *Message_GetAccountStatusResponse) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Message_GetAccountStatusResponse) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.GetAccountStatusResponse != nil {
		size, err := m.GetAccountStatusResponse.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *Message_NearBlockHeader) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Message_NearBlockHeader) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NearBlockHeader != nil {
		size, err := m.NearBlockHeader.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x42
	}
	return len(dAtA) - i, nil
}
func (m *Message_NearBlockShard) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Message_NearBlockShard) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NearBlockShard != nil {
		size, err := m.NearBlockShard.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x4a
	}
	return len(dAtA) - i, nil
}
func encodeVarint(dAtA []byte, offset int, v uint64) int {
	offset -= sov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Message) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Message) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if msg, ok := m.Payload.(*Message_NearBlockShard); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.Payload.(*Message_NearBlockHeader); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.Payload.(*Message_GetAccountStatusResponse); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.Payload.(*Message_GetAccountStatusRequest); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.Payload.(*Message_RpcResponse); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.Payload.(*Message_RpcRequest); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.Payload.(*Message_Generic); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarint(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if m.Version != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Message_Generic) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Message_Generic) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Generic)
	copy(dAtA[i:], m.Generic)
	i = encodeVarint(dAtA, i, uint64(len(m.Generic)))
	i--
	dAtA[i] = 0x1a
	return len(dAtA) - i, nil
}
func (m *Message_RpcRequest) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Message_RpcRequest) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.RpcRequest != nil {
		size, err := m.RpcRequest.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *Message_RpcResponse) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Message_RpcResponse) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.RpcResponse != nil {
		size, err := m.RpcResponse.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *Message_GetAccountStatusRequest) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Message_GetAccountStatusRequest) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.GetAccountStatusRequest != nil {
		size, err := m.GetAccountStatusRequest.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *Message_GetAccountStatusResponse) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Message_GetAccountStatusResponse) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.GetAccountStatusResponse != nil {
		size, err := m.GetAccountStatusResponse.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *Message_NearBlockHeader) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Message_NearBlockHeader) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NearBlockHeader != nil {
		size, err := m.NearBlockHeader.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x42
	}
	return len(dAtA) - i, nil
}
func (m *Message_NearBlockShard) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Message_NearBlockShard) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NearBlockShard != nil {
		size, err := m.NearBlockShard.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x4a
	}
	return len(dAtA) - i, nil
}
func (m *Message) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sov(uint64(m.Version))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if vtmsg, ok := m.Payload.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *Message_Generic) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Generic)
	n += 1 + l + sov(uint64(l))
	return n
}
func (m *Message_RpcRequest) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RpcRequest != nil {
		l = m.RpcRequest.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *Message_RpcResponse) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RpcResponse != nil {
		l = m.RpcResponse.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *Message_GetAccountStatusRequest) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GetAccountStatusRequest != nil {
		l = m.GetAccountStatusRequest.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *Message_GetAccountStatusResponse) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GetAccountStatusResponse != nil {
		l = m.GetAccountStatusResponse.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *Message_NearBlockHeader) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NearBlockHeader != nil {
		l = m.NearBlockHeader.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *Message_NearBlockShard) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NearBlockShard != nil {
		l = m.NearBlockShard.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}

func sov(x uint64) (n int) {
	return (bits.Len64(x|1) + 6) / 7
}
func soz(x uint64) (n int) {
	return sov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Generic", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Payload = &Message_Generic{Generic: v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RpcRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if oneof, ok := m.Payload.(*Message_RpcRequest); ok {
				if err := oneof.RpcRequest.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := &rpc.RpcMessage{}
				if err := v.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.Payload = &Message_RpcRequest{RpcRequest: v}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RpcResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if oneof, ok := m.Payload.(*Message_RpcResponse); ok {
				if err := oneof.RpcResponse.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := &rpc.RpcMessage{}
				if err := v.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.Payload = &Message_RpcResponse{RpcResponse: v}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GetAccountStatusRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if oneof, ok := m.Payload.(*Message_GetAccountStatusRequest); ok {
				if err := oneof.GetAccountStatusRequest.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := &get_account_status.GetAccountStatusRequest{}
				if err := v.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.Payload = &Message_GetAccountStatusRequest{GetAccountStatusRequest: v}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GetAccountStatusResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if oneof, ok := m.Payload.(*Message_GetAccountStatusResponse); ok {
				if err := oneof.GetAccountStatusResponse.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := &get_account_status.GetAccountStatusResponse{}
				if err := v.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.Payload = &Message_GetAccountStatusResponse{GetAccountStatusResponse: v}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NearBlockHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if oneof, ok := m.Payload.(*Message_NearBlockHeader); ok {
				if err := oneof.NearBlockHeader.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := &near_block.BlockHeaderView{}
				if err := v.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.Payload = &Message_NearBlockHeader{NearBlockHeader: v}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NearBlockShard", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if oneof, ok := m.Payload.(*Message_NearBlockShard); ok {
				if err := oneof.NearBlockShard.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := &near_block.BlockShard{}
				if err := v.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.Payload = &Message_NearBlockShard{NearBlockShard: v}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func skip(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflow
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflow
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflow
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLength
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroup
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLength
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLength        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflow          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroup = fmt.Errorf("proto: unexpected end of group")
)