// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymensionxyz/dymension/lightclient/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type HeaderSigner struct {
	// acc addr
	SequencerAddress string `protobuf:"bytes,1,opt,name=sequencer_address,json=sequencerAddress,proto3" json:"sequencer_address,omitempty"`
	ClientId         string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Height           uint64 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *HeaderSigner) Reset()         { *m = HeaderSigner{} }
func (m *HeaderSigner) String() string { return proto.CompactTextString(m) }
func (*HeaderSigner) ProtoMessage()    {}
func (*HeaderSigner) Descriptor() ([]byte, []int) {
	return fileDescriptor_5520440548912168, []int{0}
}
func (m *HeaderSigner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeaderSigner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HeaderSigner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HeaderSigner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderSigner.Merge(m, src)
}
func (m *HeaderSigner) XXX_Size() int {
	return m.Size()
}
func (m *HeaderSigner) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderSigner.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderSigner proto.InternalMessageInfo

func (m *HeaderSigner) GetSequencerAddress() string {
	if m != nil {
		return m.SequencerAddress
	}
	return ""
}

func (m *HeaderSigner) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *HeaderSigner) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type ClientHeightToSigner struct {
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Height   uint64 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	// acc addr
	SequencerAddress string `protobuf:"bytes,3,opt,name=sequencer_address,json=sequencerAddress,proto3" json:"sequencer_address,omitempty"`
}

func (m *ClientHeightToSigner) Reset()         { *m = ClientHeightToSigner{} }
func (m *ClientHeightToSigner) String() string { return proto.CompactTextString(m) }
func (*ClientHeightToSigner) ProtoMessage()    {}
func (*ClientHeightToSigner) Descriptor() ([]byte, []int) {
	return fileDescriptor_5520440548912168, []int{1}
}
func (m *ClientHeightToSigner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientHeightToSigner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientHeightToSigner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientHeightToSigner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientHeightToSigner.Merge(m, src)
}
func (m *ClientHeightToSigner) XXX_Size() int {
	return m.Size()
}
func (m *ClientHeightToSigner) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientHeightToSigner.DiscardUnknown(m)
}

var xxx_messageInfo_ClientHeightToSigner proto.InternalMessageInfo

func (m *ClientHeightToSigner) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *ClientHeightToSigner) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ClientHeightToSigner) GetSequencerAddress() string {
	if m != nil {
		return m.SequencerAddress
	}
	return ""
}

type GenesisState struct {
	CanonicalClients      []CanonicalClient      `protobuf:"bytes,1,rep,name=canonical_clients,json=canonicalClients,proto3" json:"canonical_clients"`
	HeaderSigners         []HeaderSigner         `protobuf:"bytes,2,rep,name=header_signers,json=headerSigners,proto3" json:"header_signers"`
	ClientHeightToSigners []ClientHeightToSigner `protobuf:"bytes,3,rep,name=client_height_to_signers,json=clientHeightToSigners,proto3" json:"client_height_to_signers"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5520440548912168, []int{2}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetCanonicalClients() []CanonicalClient {
	if m != nil {
		return m.CanonicalClients
	}
	return nil
}

func (m *GenesisState) GetHeaderSigners() []HeaderSigner {
	if m != nil {
		return m.HeaderSigners
	}
	return nil
}

func (m *GenesisState) GetClientHeightToSigners() []ClientHeightToSigner {
	if m != nil {
		return m.ClientHeightToSigners
	}
	return nil
}

type CanonicalClient struct {
	RollappId   string `protobuf:"bytes,1,opt,name=rollapp_id,json=rollappId,proto3" json:"rollapp_id,omitempty"`
	IbcClientId string `protobuf:"bytes,2,opt,name=ibc_client_id,json=ibcClientId,proto3" json:"ibc_client_id,omitempty"`
}

func (m *CanonicalClient) Reset()         { *m = CanonicalClient{} }
func (m *CanonicalClient) String() string { return proto.CompactTextString(m) }
func (*CanonicalClient) ProtoMessage()    {}
func (*CanonicalClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_5520440548912168, []int{3}
}
func (m *CanonicalClient) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CanonicalClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CanonicalClient.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CanonicalClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CanonicalClient.Merge(m, src)
}
func (m *CanonicalClient) XXX_Size() int {
	return m.Size()
}
func (m *CanonicalClient) XXX_DiscardUnknown() {
	xxx_messageInfo_CanonicalClient.DiscardUnknown(m)
}

var xxx_messageInfo_CanonicalClient proto.InternalMessageInfo

func (m *CanonicalClient) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *CanonicalClient) GetIbcClientId() string {
	if m != nil {
		return m.IbcClientId
	}
	return ""
}

func init() {
	proto.RegisterType((*HeaderSigner)(nil), "dymensionxyz.dymension.lightclient.HeaderSigner")
	proto.RegisterType((*ClientHeightToSigner)(nil), "dymensionxyz.dymension.lightclient.ClientHeightToSigner")
	proto.RegisterType((*GenesisState)(nil), "dymensionxyz.dymension.lightclient.GenesisState")
	proto.RegisterType((*CanonicalClient)(nil), "dymensionxyz.dymension.lightclient.CanonicalClient")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/lightclient/genesis.proto", fileDescriptor_5520440548912168)
}

var fileDescriptor_5520440548912168 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0xae, 0xd2, 0x40,
	0x14, 0xc6, 0xfb, 0xe7, 0xe6, 0x46, 0xe6, 0xde, 0xab, 0x30, 0x41, 0xd3, 0x68, 0xac, 0xa4, 0x2b,
	0x12, 0x93, 0x96, 0xc8, 0x86, 0xad, 0xb0, 0x10, 0xb6, 0x85, 0x95, 0x89, 0x69, 0xda, 0xe9, 0xd8,
	0x4e, 0x52, 0x66, 0x6a, 0x67, 0x50, 0xf0, 0x29, 0x7c, 0x2c, 0x56, 0x86, 0xa5, 0x2b, 0x63, 0xe0,
	0x45, 0x0c, 0x33, 0xb5, 0x16, 0x52, 0x72, 0xd9, 0xf5, 0x7c, 0x3d, 0xe7, 0xfc, 0xe6, 0x7c, 0x67,
	0x06, 0x0c, 0xe2, 0xcd, 0x12, 0x53, 0x4e, 0x18, 0x5d, 0x6f, 0xbe, 0x7b, 0x55, 0xe0, 0x65, 0x24,
	0x49, 0x05, 0xca, 0x08, 0xa6, 0xc2, 0x4b, 0x30, 0xc5, 0x9c, 0x70, 0x37, 0x2f, 0x98, 0x60, 0xd0,
	0xa9, 0x57, 0xb8, 0x55, 0xe0, 0xd6, 0x2a, 0x5e, 0x76, 0x13, 0x96, 0x30, 0x99, 0xee, 0x1d, 0xbf,
	0x54, 0xa5, 0x93, 0x83, 0xfb, 0x29, 0x0e, 0x63, 0x5c, 0xcc, 0x49, 0x42, 0x71, 0x01, 0xdf, 0x82,
	0x0e, 0xc7, 0x5f, 0x56, 0x98, 0x22, 0x5c, 0x04, 0x61, 0x1c, 0x17, 0x98, 0x73, 0x4b, 0xef, 0xe9,
	0xfd, 0x96, 0xdf, 0xae, 0x7e, 0xbc, 0x57, 0x3a, 0x7c, 0x05, 0x5a, 0xaa, 0x79, 0x40, 0x62, 0xcb,
	0x90, 0x49, 0x4f, 0x94, 0x30, 0x8b, 0xe1, 0x0b, 0x70, 0x9b, 0xe2, 0x23, 0xdf, 0x32, 0x7b, 0x7a,
	0xff, 0xc6, 0x2f, 0x23, 0x67, 0x0d, 0xba, 0x13, 0x99, 0x33, 0x95, 0xf1, 0x82, 0x95, 0xe4, 0x93,
	0x66, 0xfa, 0xc5, 0x66, 0x46, 0xbd, 0x59, 0xf3, 0x71, 0xcd, 0xe6, 0xe3, 0x3a, 0x3f, 0x0d, 0x70,
	0xff, 0x41, 0xf9, 0x36, 0x17, 0xa1, 0xc0, 0xf0, 0x33, 0xe8, 0xa0, 0x90, 0x32, 0x4a, 0x50, 0x98,
	0x05, 0x8a, 0x75, 0x1c, 0xd6, 0xec, 0xdf, 0xbd, 0x1b, 0xba, 0x8f, 0x5b, 0xea, 0x4e, 0xfe, 0x15,
	0xab, 0x81, 0xc6, 0x37, 0xdb, 0xdf, 0x6f, 0x34, 0xbf, 0x8d, 0x4e, 0x65, 0x0e, 0x3f, 0x81, 0xa7,
	0xa9, 0x34, 0x39, 0xe0, 0x72, 0x56, 0x6e, 0x19, 0x12, 0x32, 0xb8, 0x06, 0x52, 0x5f, 0x4f, 0x49,
	0x78, 0x48, 0x6b, 0x1a, 0x87, 0xdf, 0x80, 0x55, 0x3a, 0xa7, 0x5c, 0x09, 0x04, 0xab, 0x40, 0xa6,
	0x04, 0x8d, 0xae, 0x9a, 0xa6, 0x61, 0x2b, 0x25, 0xf0, 0x39, 0x6a, 0xf8, 0xc7, 0x9d, 0x05, 0x78,
	0x76, 0x66, 0x01, 0x7c, 0x0d, 0x40, 0xc1, 0xb2, 0x2c, 0xcc, 0xf3, 0xff, 0x6b, 0x6c, 0x95, 0xca,
	0x2c, 0x86, 0x0e, 0x78, 0x20, 0x11, 0x0a, 0xce, 0x6f, 0xcd, 0x1d, 0x89, 0xd0, 0xa4, 0xdc, 0xf5,
	0xd8, 0xdf, 0xee, 0x6d, 0x7d, 0xb7, 0xb7, 0xf5, 0x3f, 0x7b, 0x5b, 0xff, 0x71, 0xb0, 0xb5, 0xdd,
	0xc1, 0xd6, 0x7e, 0x1d, 0x6c, 0xed, 0xe3, 0x28, 0x21, 0x22, 0x5d, 0x45, 0x2e, 0x62, 0x4b, 0xef,
	0xc2, 0x1b, 0xf9, 0x3a, 0xf4, 0xd6, 0x27, 0x0f, 0x45, 0x6c, 0x72, 0xcc, 0xa3, 0x5b, 0x79, 0xdb,
	0x87, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x63, 0xe6, 0x41, 0x5b, 0x03, 0x00, 0x00,
}

func (m *HeaderSigner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeaderSigner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HeaderSigner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SequencerAddress) > 0 {
		i -= len(m.SequencerAddress)
		copy(dAtA[i:], m.SequencerAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.SequencerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClientHeightToSigner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientHeightToSigner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientHeightToSigner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SequencerAddress) > 0 {
		i -= len(m.SequencerAddress)
		copy(dAtA[i:], m.SequencerAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.SequencerAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Height != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClientHeightToSigners) > 0 {
		for iNdEx := len(m.ClientHeightToSigners) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClientHeightToSigners[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.HeaderSigners) > 0 {
		for iNdEx := len(m.HeaderSigners) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HeaderSigners[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.CanonicalClients) > 0 {
		for iNdEx := len(m.CanonicalClients) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CanonicalClients[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CanonicalClient) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CanonicalClient) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CanonicalClient) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IbcClientId) > 0 {
		i -= len(m.IbcClientId)
		copy(dAtA[i:], m.IbcClientId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.IbcClientId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HeaderSigner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SequencerAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovGenesis(uint64(m.Height))
	}
	return n
}

func (m *ClientHeightToSigner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovGenesis(uint64(m.Height))
	}
	l = len(m.SequencerAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CanonicalClients) > 0 {
		for _, e := range m.CanonicalClients {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.HeaderSigners) > 0 {
		for _, e := range m.HeaderSigners {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ClientHeightToSigners) > 0 {
		for _, e := range m.ClientHeightToSigners {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *CanonicalClient) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.IbcClientId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HeaderSigner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: HeaderSigner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeaderSigner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequencerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SequencerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClientHeightToSigner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ClientHeightToSigner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientHeightToSigner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequencerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SequencerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanonicalClients", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CanonicalClients = append(m.CanonicalClients, CanonicalClient{})
			if err := m.CanonicalClients[len(m.CanonicalClients)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderSigners", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeaderSigners = append(m.HeaderSigners, HeaderSigner{})
			if err := m.HeaderSigners[len(m.HeaderSigners)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientHeightToSigners", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientHeightToSigners = append(m.ClientHeightToSigners, ClientHeightToSigner{})
			if err := m.ClientHeightToSigners[len(m.ClientHeightToSigners)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CanonicalClient) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: CanonicalClient: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CanonicalClient: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
