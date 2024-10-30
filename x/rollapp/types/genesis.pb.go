// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymensionxyz/dymension/rollapp/genesis.proto

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

// GenesisState defines the rollapp module's genesis state.
type GenesisState struct {
	Params                             Params                           `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	RollappList                        []Rollapp                        `protobuf:"bytes,2,rep,name=rollappList,proto3" json:"rollappList"`
	StateInfoList                      []StateInfo                      `protobuf:"bytes,3,rep,name=stateInfoList,proto3" json:"stateInfoList"`
	LatestStateInfoIndexList           []StateInfoIndex                 `protobuf:"bytes,4,rep,name=latestStateInfoIndexList,proto3" json:"latestStateInfoIndexList"`
	LatestFinalizedStateIndexList      []StateInfoIndex                 `protobuf:"bytes,5,rep,name=latestFinalizedStateIndexList,proto3" json:"latestFinalizedStateIndexList"`
	BlockHeightToFinalizationQueueList []BlockHeightToFinalizationQueue `protobuf:"bytes,6,rep,name=blockHeightToFinalizationQueueList,proto3" json:"blockHeightToFinalizationQueueList"`
	// LivenessEvents are scheduled upcoming liveness events
	LivenessEvents   []LivenessEvent           `protobuf:"bytes,7,rep,name=livenessEvents,proto3" json:"livenessEvents"`
	AppList          []App                     `protobuf:"bytes,8,rep,name=appList,proto3" json:"appList"`
	RegisteredDenoms []RollappRegisteredDenoms `protobuf:"bytes,9,rep,name=registeredDenoms,proto3" json:"registeredDenoms"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b76890aebc09aa04, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetRollappList() []Rollapp {
	if m != nil {
		return m.RollappList
	}
	return nil
}

func (m *GenesisState) GetStateInfoList() []StateInfo {
	if m != nil {
		return m.StateInfoList
	}
	return nil
}

func (m *GenesisState) GetLatestStateInfoIndexList() []StateInfoIndex {
	if m != nil {
		return m.LatestStateInfoIndexList
	}
	return nil
}

func (m *GenesisState) GetLatestFinalizedStateIndexList() []StateInfoIndex {
	if m != nil {
		return m.LatestFinalizedStateIndexList
	}
	return nil
}

func (m *GenesisState) GetBlockHeightToFinalizationQueueList() []BlockHeightToFinalizationQueue {
	if m != nil {
		return m.BlockHeightToFinalizationQueueList
	}
	return nil
}

func (m *GenesisState) GetLivenessEvents() []LivenessEvent {
	if m != nil {
		return m.LivenessEvents
	}
	return nil
}

func (m *GenesisState) GetAppList() []App {
	if m != nil {
		return m.AppList
	}
	return nil
}

func (m *GenesisState) GetRegisteredDenoms() []RollappRegisteredDenoms {
	if m != nil {
		return m.RegisteredDenoms
	}
	return nil
}

type RollappRegisteredDenoms struct {
	RollappId string   `protobuf:"bytes,1,opt,name=rollapp_id,json=rollappId,proto3" json:"rollapp_id,omitempty"`
	Denoms    []string `protobuf:"bytes,2,rep,name=denoms,proto3" json:"denoms,omitempty"`
}

func (m *RollappRegisteredDenoms) Reset()         { *m = RollappRegisteredDenoms{} }
func (m *RollappRegisteredDenoms) String() string { return proto.CompactTextString(m) }
func (*RollappRegisteredDenoms) ProtoMessage()    {}
func (*RollappRegisteredDenoms) Descriptor() ([]byte, []int) {
	return fileDescriptor_b76890aebc09aa04, []int{1}
}
func (m *RollappRegisteredDenoms) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RollappRegisteredDenoms) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RollappRegisteredDenoms.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RollappRegisteredDenoms) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RollappRegisteredDenoms.Merge(m, src)
}
func (m *RollappRegisteredDenoms) XXX_Size() int {
	return m.Size()
}
func (m *RollappRegisteredDenoms) XXX_DiscardUnknown() {
	xxx_messageInfo_RollappRegisteredDenoms.DiscardUnknown(m)
}

var xxx_messageInfo_RollappRegisteredDenoms proto.InternalMessageInfo

func (m *RollappRegisteredDenoms) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *RollappRegisteredDenoms) GetDenoms() []string {
	if m != nil {
		return m.Denoms
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "dymensionxyz.dymension.rollapp.GenesisState")
	proto.RegisterType((*RollappRegisteredDenoms)(nil), "dymensionxyz.dymension.rollapp.RollappRegisteredDenoms")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/rollapp/genesis.proto", fileDescriptor_b76890aebc09aa04)
}

var fileDescriptor_b76890aebc09aa04 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x4d, 0x6b, 0x13, 0x41,
	0x1c, 0xc6, 0xb3, 0xb6, 0x26, 0x66, 0xa2, 0x22, 0x83, 0xe0, 0x12, 0xec, 0x5a, 0x22, 0x68, 0x44,
	0xbb, 0x0b, 0xad, 0xe0, 0x4d, 0x30, 0xd6, 0x97, 0x40, 0xf1, 0x25, 0x55, 0x0f, 0x7a, 0x28, 0x9b,
	0xee, 0xbf, 0xdb, 0xc1, 0xcd, 0xcc, 0xb2, 0x33, 0x09, 0x49, 0x3e, 0x84, 0x78, 0xf0, 0x43, 0xf5,
	0xd8, 0xa3, 0x27, 0x91, 0xe4, 0x8b, 0xc8, 0xce, 0xfe, 0x67, 0x89, 0x95, 0x64, 0x02, 0x9e, 0xf6,
	0xed, 0x79, 0x7e, 0xcf, 0xb3, 0xc3, 0x7f, 0x86, 0x3c, 0x8a, 0x26, 0x03, 0xe0, 0x92, 0x09, 0x3e,
	0x9e, 0x4c, 0x83, 0xf2, 0x21, 0xc8, 0x44, 0x92, 0x84, 0x69, 0x1a, 0xc4, 0xc0, 0x41, 0x32, 0xe9,
	0xa7, 0x99, 0x50, 0x82, 0x7a, 0x8b, 0x6a, 0xbf, 0x7c, 0xf0, 0x51, 0xdd, 0xbc, 0x19, 0x8b, 0x58,
	0x68, 0x69, 0x90, 0xdf, 0x15, 0xae, 0xe6, 0x43, 0x4b, 0x46, 0x1a, 0x66, 0xe1, 0x00, 0x23, 0x9a,
	0xb6, 0x42, 0x78, 0x45, 0x75, 0x60, 0x51, 0x4b, 0x15, 0x2a, 0x38, 0x62, 0xfc, 0xc4, 0x74, 0xd9,
	0xb1, 0x18, 0x12, 0x36, 0xca, 0xff, 0xd8, 0xb4, 0x69, 0x5b, 0xe4, 0x65, 0x93, 0xd6, 0xb7, 0x1a,
	0xb9, 0xfa, 0xaa, 0x58, 0xac, 0xc3, 0x3c, 0x94, 0xee, 0x93, 0x6a, 0xf1, 0x63, 0xae, 0xb3, 0xed,
	0xb4, 0x1b, 0xbb, 0xf7, 0xfc, 0xd5, 0x8b, 0xe7, 0xbf, 0xd3, 0xea, 0xce, 0xe6, 0xd9, 0xaf, 0x3b,
	0x95, 0x1e, 0x7a, 0xe9, 0x5b, 0xd2, 0xc0, 0xef, 0x07, 0x4c, 0x2a, 0xf7, 0xd2, 0xf6, 0x46, 0xbb,
	0xb1, 0x7b, 0xdf, 0x86, 0xea, 0x15, 0x57, 0x64, 0x2d, 0x12, 0xe8, 0x47, 0x72, 0x4d, 0x2f, 0x4a,
	0x97, 0x9f, 0x08, 0x8d, 0xdc, 0xd0, 0xc8, 0x07, 0x36, 0xe4, 0xa1, 0x31, 0x21, 0xf4, 0x6f, 0x0a,
	0x4d, 0x89, 0x9b, 0x84, 0x0a, 0xa4, 0x2a, 0x75, 0x5d, 0x1e, 0xc1, 0x58, 0x27, 0x6c, 0xea, 0x04,
	0x7f, 0xed, 0x04, 0xed, 0xc4, 0x98, 0xa5, 0x54, 0x3a, 0x25, 0x5b, 0xc5, 0xb7, 0x97, 0x8c, 0x87,
	0x09, 0x9b, 0x42, 0x84, 0x22, 0x13, 0x7b, 0xf9, 0x3f, 0x62, 0x57, 0xa3, 0xe9, 0x0f, 0x87, 0xb4,
	0xfa, 0x89, 0x38, 0xfe, 0xfa, 0x1a, 0x58, 0x7c, 0xaa, 0x3e, 0x08, 0x14, 0x86, 0x8a, 0x09, 0xfe,
	0x7e, 0x08, 0x43, 0xd0, 0x0d, 0xaa, 0xba, 0xc1, 0x53, 0x5b, 0x83, 0xce, 0x4a, 0x12, 0x36, 0x5a,
	0x23, 0x8f, 0x7e, 0x21, 0xd7, 0xcd, 0xfc, 0xbe, 0x18, 0x01, 0x57, 0xd2, 0xad, 0xe9, 0x06, 0x3b,
	0xb6, 0x06, 0x07, 0x8b, 0x2e, 0x0c, 0xbc, 0x80, 0xa2, 0xcf, 0x49, 0xcd, 0x4c, 0xe1, 0x15, 0x4d,
	0xbd, 0x6b, 0xa3, 0x3e, 0x2b, 0x27, 0xd0, 0x38, 0x29, 0x23, 0x37, 0x32, 0x88, 0x99, 0x54, 0x90,
	0x41, 0xb4, 0x0f, 0x5c, 0x0c, 0xa4, 0x5b, 0xd7, 0xb4, 0x27, 0x6b, 0xce, 0x74, 0xef, 0x82, 0x1d,
	0x13, 0xfe, 0xc1, 0xb6, 0x3e, 0x91, 0x5b, 0x4b, 0x2c, 0x74, 0x8b, 0x10, 0xa4, 0x1e, 0xb1, 0x48,
	0x6f, 0xcf, 0x7a, 0xaf, 0x8e, 0x6f, 0xba, 0x11, 0xbd, 0x4d, 0xaa, 0x51, 0x51, 0x2d, 0xdf, 0x6e,
	0x75, 0xb3, 0x23, 0x8b, 0x77, 0x9d, 0x37, 0x67, 0x33, 0xcf, 0x39, 0x9f, 0x79, 0xce, 0xef, 0x99,
	0xe7, 0x7c, 0x9f, 0x7b, 0x95, 0xf3, 0xb9, 0x57, 0xf9, 0x39, 0xf7, 0x2a, 0x9f, 0x1f, 0xc7, 0x4c,
	0x9d, 0x0e, 0xfb, 0xfe, 0xb1, 0x18, 0x2c, 0x3b, 0x97, 0x46, 0x7b, 0xc1, 0xb8, 0x3c, 0x3c, 0xd4,
	0x24, 0x05, 0xd9, 0xaf, 0xea, 0xf3, 0x63, 0xef, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xc7,
	0x8b, 0xbc, 0x8a, 0x05, 0x00, 0x00,
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
	if len(m.RegisteredDenoms) > 0 {
		for iNdEx := len(m.RegisteredDenoms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RegisteredDenoms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.AppList) > 0 {
		for iNdEx := len(m.AppList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AppList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.LivenessEvents) > 0 {
		for iNdEx := len(m.LivenessEvents) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LivenessEvents[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.BlockHeightToFinalizationQueueList) > 0 {
		for iNdEx := len(m.BlockHeightToFinalizationQueueList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BlockHeightToFinalizationQueueList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.LatestFinalizedStateIndexList) > 0 {
		for iNdEx := len(m.LatestFinalizedStateIndexList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LatestFinalizedStateIndexList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.LatestStateInfoIndexList) > 0 {
		for iNdEx := len(m.LatestStateInfoIndexList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LatestStateInfoIndexList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.StateInfoList) > 0 {
		for iNdEx := len(m.StateInfoList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StateInfoList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.RollappList) > 0 {
		for iNdEx := len(m.RollappList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RollappList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *RollappRegisteredDenoms) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RollappRegisteredDenoms) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RollappRegisteredDenoms) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denoms) > 0 {
		for iNdEx := len(m.Denoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Denoms[iNdEx])
			copy(dAtA[i:], m.Denoms[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.Denoms[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
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
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.RollappList) > 0 {
		for _, e := range m.RollappList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.StateInfoList) > 0 {
		for _, e := range m.StateInfoList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LatestStateInfoIndexList) > 0 {
		for _, e := range m.LatestStateInfoIndexList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LatestFinalizedStateIndexList) > 0 {
		for _, e := range m.LatestFinalizedStateIndexList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BlockHeightToFinalizationQueueList) > 0 {
		for _, e := range m.BlockHeightToFinalizationQueueList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LivenessEvents) > 0 {
		for _, e := range m.LivenessEvents {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AppList) > 0 {
		for _, e := range m.AppList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RegisteredDenoms) > 0 {
		for _, e := range m.RegisteredDenoms {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *RollappRegisteredDenoms) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Denoms) > 0 {
		for _, s := range m.Denoms {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappList", wireType)
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
			m.RollappList = append(m.RollappList, Rollapp{})
			if err := m.RollappList[len(m.RollappList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateInfoList", wireType)
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
			m.StateInfoList = append(m.StateInfoList, StateInfo{})
			if err := m.StateInfoList[len(m.StateInfoList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestStateInfoIndexList", wireType)
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
			m.LatestStateInfoIndexList = append(m.LatestStateInfoIndexList, StateInfoIndex{})
			if err := m.LatestStateInfoIndexList[len(m.LatestStateInfoIndexList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestFinalizedStateIndexList", wireType)
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
			m.LatestFinalizedStateIndexList = append(m.LatestFinalizedStateIndexList, StateInfoIndex{})
			if err := m.LatestFinalizedStateIndexList[len(m.LatestFinalizedStateIndexList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeightToFinalizationQueueList", wireType)
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
			m.BlockHeightToFinalizationQueueList = append(m.BlockHeightToFinalizationQueueList, BlockHeightToFinalizationQueue{})
			if err := m.BlockHeightToFinalizationQueueList[len(m.BlockHeightToFinalizationQueueList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LivenessEvents", wireType)
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
			m.LivenessEvents = append(m.LivenessEvents, LivenessEvent{})
			if err := m.LivenessEvents[len(m.LivenessEvents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppList", wireType)
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
			m.AppList = append(m.AppList, App{})
			if err := m.AppList[len(m.AppList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegisteredDenoms", wireType)
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
			m.RegisteredDenoms = append(m.RegisteredDenoms, RollappRegisteredDenoms{})
			if err := m.RegisteredDenoms[len(m.RegisteredDenoms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *RollappRegisteredDenoms) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: RollappRegisteredDenoms: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RollappRegisteredDenoms: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Denoms", wireType)
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
			m.Denoms = append(m.Denoms, string(dAtA[iNdEx:postIndex]))
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
