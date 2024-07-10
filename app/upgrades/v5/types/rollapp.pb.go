// COPIED FROM main

package types

import (
	"fmt"
	"io"
	"math"
	math_bits "math/bits"

	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	"github.com/cosmos/gogoproto/proto"
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

// RollappGenesisState is a partial repr of the state the hub can expect the rollapp to be in upon genesis
type RollappGenesisState struct {
	// If true, then full usage of the canonical ibc transfer channel is enabled.
	// Note: in v3.1.0 and prior this field marked the completion of the 'genesis event'
	// Keeping and renaming the field enables a seamless upgrade https://www.notion.so/dymension/ADR-x-Genesis-Bridge-Phase-2-89769aa551b5440b9ed403a101775ce1?pvs=4#89698384d815435b87393dbe45bc5a74
	// to the new genesis transfer protocol
	// Note: if this field is false, ibc transfers may still be allowed in one or either direction.
	TransfersEnabled bool `protobuf:"varint,2,opt,name=transfers_enabled,json=transfersEnabled,proto3" json:"transfers_enabled,omitempty"`
}

func (m *RollappGenesisState) Reset()         { *m = RollappGenesisState{} }
func (m *RollappGenesisState) String() string { return proto.CompactTextString(m) }
func (*RollappGenesisState) ProtoMessage()    {}
func (*RollappGenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4ef2bec3aea5528, []int{0}
}
func (m *RollappGenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RollappGenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RollappGenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RollappGenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RollappGenesisState.Merge(m, src)
}
func (m *RollappGenesisState) XXX_Size() int {
	return m.Size()
}
func (m *RollappGenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_RollappGenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_RollappGenesisState proto.InternalMessageInfo

func (m *RollappGenesisState) GetTransfersEnabled() bool {
	if m != nil {
		return m.TransfersEnabled
	}
	return false
}

// Rollapp defines a rollapp object. First the RollApp is created and then
// sequencers can be created and attached. The RollApp is identified by rollappId
type Rollapp struct {
	// The unique identifier of the rollapp chain.
	// The rollappId follows the same standard as cosmos chain_id.
	RollappId string `protobuf:"bytes,1,opt,name=rollappId,proto3" json:"rollappId,omitempty"`
	// creator is the bech32-encoded address of the rollapp creator.
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	// version is the software and configuration version.
	// starts from 1 and increases by one on every MsgUpdateState
	Version uint64 `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	// maxSequencers is the maximum number of sequencers.
	MaxSequencers uint64 `protobuf:"varint,4,opt,name=maxSequencers,proto3" json:"maxSequencers,omitempty"`
	// permissionedAddresses is a bech32-encoded address list of the sequencers that are allowed to serve this rollappId.
	// In the case of an empty list, the rollapp is considered permissionless.
	PermissionedAddresses []string `protobuf:"bytes,5,rep,name=permissionedAddresses,proto3" json:"permissionedAddresses,omitempty"`
	// genesis_state is a partial repr of the state the hub can expect the rollapp to be in upon genesis
	GenesisState RollappGenesisState `protobuf:"bytes,7,opt,name=genesis_state,json=genesisState,proto3" json:"genesis_state"`
	// channel_id will be set to the canonical IBC channel of the rollapp.
	ChannelId string `protobuf:"bytes,8,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// frozen is a boolean that indicates if the rollapp is frozen.
	Frozen bool `protobuf:"varint,9,opt,name=frozen,proto3" json:"frozen,omitempty"`
	// registeredDenoms is a list of registered denom bases on this rollapp
	RegisteredDenoms []string `protobuf:"bytes,10,rep,name=registeredDenoms,proto3" json:"registeredDenoms,omitempty"`
}

func (m *Rollapp) Reset()         { *m = Rollapp{} }
func (m *Rollapp) String() string { return proto.CompactTextString(m) }
func (*Rollapp) ProtoMessage()    {}
func (*Rollapp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4ef2bec3aea5528, []int{1}
}
func (m *Rollapp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Rollapp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Rollapp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Rollapp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rollapp.Merge(m, src)
}
func (m *Rollapp) XXX_Size() int {
	return m.Size()
}
func (m *Rollapp) XXX_DiscardUnknown() {
	xxx_messageInfo_Rollapp.DiscardUnknown(m)
}

var xxx_messageInfo_Rollapp proto.InternalMessageInfo

func (m *Rollapp) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *Rollapp) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Rollapp) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Rollapp) GetMaxSequencers() uint64 {
	if m != nil {
		return m.MaxSequencers
	}
	return 0
}

func (m *Rollapp) GetPermissionedAddresses() []string {
	if m != nil {
		return m.PermissionedAddresses
	}
	return nil
}

func (m *Rollapp) GetGenesisState() RollappGenesisState {
	if m != nil {
		return m.GenesisState
	}
	return RollappGenesisState{}
}

func (m *Rollapp) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *Rollapp) GetFrozen() bool {
	if m != nil {
		return m.Frozen
	}
	return false
}

func (m *Rollapp) GetRegisteredDenoms() []string {
	if m != nil {
		return m.RegisteredDenoms
	}
	return nil
}

// Rollapp summary is a compact representation of Rollapp
type RollappSummary struct {
	// The unique identifier of the rollapp chain.
	// The rollappId follows the same standard as cosmos chain_id.
	RollappId string `protobuf:"bytes,1,opt,name=rollappId,proto3" json:"rollappId,omitempty"`
	// Defines the index of the last rollapp UpdateState.
	LatestStateIndex *StateInfoIndex `protobuf:"bytes,2,opt,name=latestStateIndex,proto3" json:"latestStateIndex,omitempty"`
	// Defines the index of the last rollapp UpdateState that was finalized.
	LatestFinalizedStateIndex *StateInfoIndex `protobuf:"bytes,3,opt,name=latestFinalizedStateIndex,proto3" json:"latestFinalizedStateIndex,omitempty"`
}

func (m *RollappSummary) Reset()         { *m = RollappSummary{} }
func (m *RollappSummary) String() string { return proto.CompactTextString(m) }
func (*RollappSummary) ProtoMessage()    {}
func (*RollappSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4ef2bec3aea5528, []int{2}
}
func (m *RollappSummary) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RollappSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RollappSummary.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RollappSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RollappSummary.Merge(m, src)
}
func (m *RollappSummary) XXX_Size() int {
	return m.Size()
}
func (m *RollappSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_RollappSummary.DiscardUnknown(m)
}

var xxx_messageInfo_RollappSummary proto.InternalMessageInfo

func (m *RollappSummary) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *RollappSummary) GetLatestStateIndex() *StateInfoIndex {
	if m != nil {
		return m.LatestStateIndex
	}
	return nil
}

func (m *RollappSummary) GetLatestFinalizedStateIndex() *StateInfoIndex {
	if m != nil {
		return m.LatestFinalizedStateIndex
	}
	return nil
}

func init() {
	proto.RegisterType((*RollappGenesisState)(nil), "dymensionxyz.dymension.rollapp.RollappGenesisState")
	proto.RegisterType((*Rollapp)(nil), "dymensionxyz.dymension.rollapp.Rollapp")
	proto.RegisterType((*RollappSummary)(nil), "dymensionxyz.dymension.rollapp.RollappSummary")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/rollapp/rollapp.proto", fileDescriptor_d4ef2bec3aea5528)
}

var fileDescriptor_d4ef2bec3aea5528 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x8e, 0x9b, 0x90, 0x1f, 0x97, 0xa2, 0x60, 0x7e, 0x64, 0x2a, 0x58, 0xa2, 0x88, 0x43, 0x04,
	0x68, 0x57, 0x6d, 0xfa, 0x02, 0x54, 0xfc, 0xa5, 0x07, 0x0e, 0x9b, 0x5b, 0x0f, 0x44, 0xce, 0x7a,
	0xb2, 0xb5, 0xb4, 0x6b, 0x07, 0xdb, 0x89, 0x92, 0x9c, 0x78, 0x04, 0x1e, 0xab, 0xc7, 0x1e, 0x39,
	0x21, 0x94, 0xbc, 0x04, 0x47, 0xb4, 0xde, 0x4d, 0x08, 0x6a, 0x4b, 0x24, 0x4e, 0xbb, 0xf3, 0x8d,
	0xbf, 0x6f, 0x66, 0xbe, 0xd1, 0xe0, 0xd7, 0x7c, 0x9e, 0x82, 0x34, 0x42, 0xc9, 0xd9, 0x7c, 0x11,
	0x6c, 0x82, 0x40, 0xab, 0x24, 0x61, 0xe3, 0xf1, 0xfa, 0xeb, 0x8f, 0xb5, 0xb2, 0x8a, 0x78, 0xdb,
	0xaf, 0xfd, 0x4d, 0xe0, 0x17, 0xaf, 0x0e, 0x1f, 0xc6, 0x2a, 0x56, 0xee, 0x69, 0x90, 0xfd, 0xe5,
	0xac, 0xc3, 0x60, 0x47, 0x0d, 0x63, 0x99, 0x85, 0x81, 0x90, 0xa3, 0x35, 0xc1, 0x8b, 0x94, 0x49,
	0x95, 0x09, 0x86, 0xcc, 0x40, 0x30, 0x3d, 0x1a, 0x82, 0x65, 0x47, 0x41, 0xa4, 0x84, 0xcc, 0xf3,
	0xed, 0x8f, 0xf8, 0x41, 0x98, 0x73, 0x3f, 0x80, 0x04, 0x23, 0x4c, 0x3f, 0x53, 0x20, 0xaf, 0xf0,
	0x7d, 0xab, 0x99, 0x34, 0x23, 0xd0, 0x66, 0x00, 0x92, 0x0d, 0x13, 0xe0, 0x74, 0xaf, 0x85, 0x3a,
	0xf5, 0xb0, 0xb9, 0x49, 0xbc, 0xcb, 0xf1, 0xb3, 0x4a, 0x1d, 0x35, 0xf7, 0xda, 0x5f, 0xcb, 0xb8,
	0x56, 0x48, 0x91, 0xa7, 0xb8, 0x51, 0x74, 0xd4, 0xe3, 0x14, 0xb5, 0x50, 0xa7, 0x11, 0xfe, 0x01,
	0x08, 0xc5, 0xb5, 0x48, 0x03, 0xb3, 0x4a, 0x3b, 0xc9, 0x46, 0xb8, 0x0e, 0xb3, 0xcc, 0x14, 0x74,
	0x36, 0x11, 0x2d, 0xb7, 0x50, 0xa7, 0x12, 0xae, 0x43, 0xf2, 0x02, 0x1f, 0xa4, 0x6c, 0xd6, 0x87,
	0x2f, 0x13, 0x90, 0x11, 0x68, 0x43, 0x2b, 0x2e, 0xff, 0x37, 0x48, 0x4e, 0xf0, 0xa3, 0x31, 0xe8,
	0x54, 0x98, 0x8c, 0x03, 0xfc, 0x0d, 0xe7, 0x1a, 0x8c, 0x01, 0x43, 0xef, 0xb4, 0xca, 0x9d, 0x46,
	0x78, 0x73, 0x92, 0x7c, 0xc6, 0x07, 0x71, 0x3e, 0xfc, 0xc0, 0xf9, 0x47, 0x6b, 0x2d, 0xd4, 0xd9,
	0x3f, 0xee, 0xfa, 0xff, 0x5e, 0x91, 0x7f, 0x83, 0x71, 0xa7, 0x95, 0xcb, 0x1f, 0xcf, 0x4b, 0xe1,
	0xdd, 0x78, 0xdb, 0xcc, 0x67, 0x18, 0x47, 0x17, 0x4c, 0x4a, 0x48, 0x06, 0x82, 0xd3, 0x7a, 0x6e,
	0x47, 0x81, 0xf4, 0x38, 0x79, 0x8c, 0xab, 0x23, 0xad, 0x16, 0x20, 0x69, 0xc3, 0x19, 0x5c, 0x44,
	0xe4, 0x25, 0x6e, 0x6a, 0x88, 0x85, 0xb1, 0xa0, 0x81, 0xbf, 0x05, 0xa9, 0x52, 0x43, 0xb1, 0x9b,
	0xe3, 0x1a, 0x7e, 0x56, 0xa9, 0x57, 0x9b, 0xb5, 0xf6, 0x2f, 0x84, 0xef, 0x15, 0x4d, 0xf5, 0x27,
	0x69, 0xca, 0xf4, 0x7c, 0xc7, 0x26, 0xce, 0x71, 0x33, 0x61, 0x16, 0x8c, 0x75, 0x8d, 0xf6, 0x24,
	0x87, 0x99, 0x5b, 0xc9, 0xfe, 0xb1, 0xbf, 0x6b, 0xf8, 0x82, 0x31, 0x52, 0x8e, 0x15, 0x5e, 0xd3,
	0x21, 0x09, 0x7e, 0x92, 0x63, 0xef, 0x85, 0x64, 0x89, 0x58, 0x00, 0xdf, 0x2a, 0x52, 0xfe, 0xaf,
	0x22, 0xb7, 0x0b, 0x9e, 0x7e, 0xba, 0x5c, 0x7a, 0xe8, 0x6a, 0xe9, 0xa1, 0x9f, 0x4b, 0x0f, 0x7d,
	0x5b, 0x79, 0xa5, 0xab, 0x95, 0x57, 0xfa, 0xbe, 0xf2, 0x4a, 0xe7, 0x27, 0xb1, 0xb0, 0x17, 0x93,
	0xa1, 0x1f, 0xa9, 0xf4, 0xb6, 0xeb, 0x99, 0x76, 0x83, 0xd9, 0xe6, 0x84, 0xec, 0x7c, 0x0c, 0x66,
	0x58, 0x75, 0xe7, 0xd1, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xf3, 0x6f, 0x10, 0xd5, 0x03,
	0x00, 0x00,
}

func (m *RollappGenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RollappGenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RollappGenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TransfersEnabled {
		i--
		if m.TransfersEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	return len(dAtA) - i, nil
}

func (m *Rollapp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Rollapp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Rollapp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RegisteredDenoms) > 0 {
		for iNdEx := len(m.RegisteredDenoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.RegisteredDenoms[iNdEx])
			copy(dAtA[i:], m.RegisteredDenoms[iNdEx])
			i = encodeVarintRollapp(dAtA, i, uint64(len(m.RegisteredDenoms[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if m.Frozen {
		i--
		if m.Frozen {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintRollapp(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x42
	}
	{
		size, err := m.GenesisState.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintRollapp(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.PermissionedAddresses) > 0 {
		for iNdEx := len(m.PermissionedAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PermissionedAddresses[iNdEx])
			copy(dAtA[i:], m.PermissionedAddresses[iNdEx])
			i = encodeVarintRollapp(dAtA, i, uint64(len(m.PermissionedAddresses[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.MaxSequencers != 0 {
		i = encodeVarintRollapp(dAtA, i, uint64(m.MaxSequencers))
		i--
		dAtA[i] = 0x20
	}
	if m.Version != 0 {
		i = encodeVarintRollapp(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintRollapp(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintRollapp(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RollappSummary) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RollappSummary) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RollappSummary) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LatestFinalizedStateIndex != nil {
		{
			size, err := m.LatestFinalizedStateIndex.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRollapp(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.LatestStateIndex != nil {
		{
			size, err := m.LatestStateIndex.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRollapp(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintRollapp(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRollapp(dAtA []byte, offset int, v uint64) int {
	offset -= sovRollapp(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RollappGenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TransfersEnabled {
		n += 2
	}
	return n
}

func (m *Rollapp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovRollapp(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovRollapp(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovRollapp(uint64(m.Version))
	}
	if m.MaxSequencers != 0 {
		n += 1 + sovRollapp(uint64(m.MaxSequencers))
	}
	if len(m.PermissionedAddresses) > 0 {
		for _, s := range m.PermissionedAddresses {
			l = len(s)
			n += 1 + l + sovRollapp(uint64(l))
		}
	}
	l = m.GenesisState.Size()
	n += 1 + l + sovRollapp(uint64(l))
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovRollapp(uint64(l))
	}
	if m.Frozen {
		n += 2
	}
	if len(m.RegisteredDenoms) > 0 {
		for _, s := range m.RegisteredDenoms {
			l = len(s)
			n += 1 + l + sovRollapp(uint64(l))
		}
	}
	return n
}

func (m *RollappSummary) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovRollapp(uint64(l))
	}
	if m.LatestStateIndex != nil {
		l = m.LatestStateIndex.Size()
		n += 1 + l + sovRollapp(uint64(l))
	}
	if m.LatestFinalizedStateIndex != nil {
		l = m.LatestFinalizedStateIndex.Size()
		n += 1 + l + sovRollapp(uint64(l))
	}
	return n
}

func sovRollapp(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRollapp(x uint64) (n int) {
	return sovRollapp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RollappGenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRollapp
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
			return fmt.Errorf("proto: RollappGenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RollappGenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransfersEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.TransfersEnabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRollapp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRollapp
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
func (m *Rollapp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRollapp
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
			return fmt.Errorf("proto: Rollapp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Rollapp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
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
				return ErrInvalidLengthRollapp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
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
				return ErrInvalidLengthRollapp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSequencers", wireType)
			}
			m.MaxSequencers = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSequencers |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PermissionedAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
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
				return ErrInvalidLengthRollapp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PermissionedAddresses = append(m.PermissionedAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
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
				return ErrInvalidLengthRollapp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GenesisState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
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
				return ErrInvalidLengthRollapp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Frozen", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Frozen = bool(v != 0)
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegisteredDenoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
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
				return ErrInvalidLengthRollapp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RegisteredDenoms = append(m.RegisteredDenoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRollapp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRollapp
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
func (m *RollappSummary) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRollapp
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
			return fmt.Errorf("proto: RollappSummary: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RollappSummary: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
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
				return ErrInvalidLengthRollapp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestStateIndex", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
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
				return ErrInvalidLengthRollapp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LatestStateIndex == nil {
				m.LatestStateIndex = &StateInfoIndex{}
			}
			if err := m.LatestStateIndex.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestFinalizedStateIndex", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
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
				return ErrInvalidLengthRollapp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LatestFinalizedStateIndex == nil {
				m.LatestFinalizedStateIndex = &StateInfoIndex{}
			}
			if err := m.LatestFinalizedStateIndex.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRollapp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRollapp
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
func skipRollapp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRollapp
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
					return 0, ErrIntOverflowRollapp
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
					return 0, ErrIntOverflowRollapp
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
				return 0, ErrInvalidLengthRollapp
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRollapp
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRollapp
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRollapp        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRollapp          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRollapp = fmt.Errorf("proto: unexpected end of group")
)