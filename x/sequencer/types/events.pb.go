// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymensionxyz/dymension/sequencer/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
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

// EventIncreasedBond is an event emitted when a sequencer's bond is increased.
type EventIncreasedBond struct {
	// sequencer is the bech32-encoded address of the sequencer
	Sequencer string `protobuf:"bytes,1,opt,name=sequencer,proto3" json:"sequencer,omitempty"`
	// added_amount is the amount of coins added to the sequencer's bond
	AddedAmount types.Coin `protobuf:"bytes,2,opt,name=added_amount,json=addedAmount,proto3" json:"added_amount"`
	// bond is the new active bond amount of the sequencer
	Bond github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=bond,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"bond"`
}

func (m *EventIncreasedBond) Reset()         { *m = EventIncreasedBond{} }
func (m *EventIncreasedBond) String() string { return proto.CompactTextString(m) }
func (*EventIncreasedBond) ProtoMessage()    {}
func (*EventIncreasedBond) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f8a63d7e7167eb3, []int{0}
}
func (m *EventIncreasedBond) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventIncreasedBond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventIncreasedBond.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventIncreasedBond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventIncreasedBond.Merge(m, src)
}
func (m *EventIncreasedBond) XXX_Size() int {
	return m.Size()
}
func (m *EventIncreasedBond) XXX_DiscardUnknown() {
	xxx_messageInfo_EventIncreasedBond.DiscardUnknown(m)
}

var xxx_messageInfo_EventIncreasedBond proto.InternalMessageInfo

func (m *EventIncreasedBond) GetSequencer() string {
	if m != nil {
		return m.Sequencer
	}
	return ""
}

func (m *EventIncreasedBond) GetAddedAmount() types.Coin {
	if m != nil {
		return m.AddedAmount
	}
	return types.Coin{}
}

func (m *EventIncreasedBond) GetBond() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Bond
	}
	return nil
}

// On a sequencer kicking the incumbent proposer
type EventKickedProposer struct {
	// Kicker is the bech32-encoded address of the sequencer who triggered the kick
	Kicker string `protobuf:"bytes,1,opt,name=kicker,proto3" json:"kicker,omitempty"`
	// Proposer is the bech32-encoded address of the proposer who was kicked
	Proposer string `protobuf:"bytes,2,opt,name=proposer,proto3" json:"proposer,omitempty"`
}

func (m *EventKickedProposer) Reset()         { *m = EventKickedProposer{} }
func (m *EventKickedProposer) String() string { return proto.CompactTextString(m) }
func (*EventKickedProposer) ProtoMessage()    {}
func (*EventKickedProposer) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f8a63d7e7167eb3, []int{1}
}
func (m *EventKickedProposer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventKickedProposer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventKickedProposer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventKickedProposer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventKickedProposer.Merge(m, src)
}
func (m *EventKickedProposer) XXX_Size() int {
	return m.Size()
}
func (m *EventKickedProposer) XXX_DiscardUnknown() {
	xxx_messageInfo_EventKickedProposer.DiscardUnknown(m)
}

var xxx_messageInfo_EventKickedProposer proto.InternalMessageInfo

func (m *EventKickedProposer) GetKicker() string {
	if m != nil {
		return m.Kicker
	}
	return ""
}

func (m *EventKickedProposer) GetProposer() string {
	if m != nil {
		return m.Proposer
	}
	return ""
}

// Whenever the proposer changes to a new proposer
type EventProposerChange struct {
	// Before is the bech32-encoded address of the old proposer
	Before string `protobuf:"bytes,1,opt,name=before,proto3" json:"before,omitempty"`
	// After is the bech32-encoded address of the new proposer
	After string `protobuf:"bytes,2,opt,name=after,proto3" json:"after,omitempty"`
}

func (m *EventProposerChange) Reset()         { *m = EventProposerChange{} }
func (m *EventProposerChange) String() string { return proto.CompactTextString(m) }
func (*EventProposerChange) ProtoMessage()    {}
func (*EventProposerChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f8a63d7e7167eb3, []int{2}
}
func (m *EventProposerChange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventProposerChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventProposerChange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventProposerChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventProposerChange.Merge(m, src)
}
func (m *EventProposerChange) XXX_Size() int {
	return m.Size()
}
func (m *EventProposerChange) XXX_DiscardUnknown() {
	xxx_messageInfo_EventProposerChange.DiscardUnknown(m)
}

var xxx_messageInfo_EventProposerChange proto.InternalMessageInfo

func (m *EventProposerChange) GetBefore() string {
	if m != nil {
		return m.Before
	}
	return ""
}

func (m *EventProposerChange) GetAfter() string {
	if m != nil {
		return m.After
	}
	return ""
}

// When a sequencer opt-in status changes
type OptInStatusChange struct {
	Before bool `protobuf:"varint,1,opt,name=before,proto3" json:"before,omitempty"`
	After  bool `protobuf:"varint,2,opt,name=after,proto3" json:"after,omitempty"`
}

func (m *OptInStatusChange) Reset()         { *m = OptInStatusChange{} }
func (m *OptInStatusChange) String() string { return proto.CompactTextString(m) }
func (*OptInStatusChange) ProtoMessage()    {}
func (*OptInStatusChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f8a63d7e7167eb3, []int{3}
}
func (m *OptInStatusChange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OptInStatusChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OptInStatusChange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OptInStatusChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OptInStatusChange.Merge(m, src)
}
func (m *OptInStatusChange) XXX_Size() int {
	return m.Size()
}
func (m *OptInStatusChange) XXX_DiscardUnknown() {
	xxx_messageInfo_OptInStatusChange.DiscardUnknown(m)
}

var xxx_messageInfo_OptInStatusChange proto.InternalMessageInfo

func (m *OptInStatusChange) GetBefore() bool {
	if m != nil {
		return m.Before
	}
	return false
}

func (m *OptInStatusChange) GetAfter() bool {
	if m != nil {
		return m.After
	}
	return false
}

func init() {
	proto.RegisterType((*EventIncreasedBond)(nil), "dymensionxyz.dymension.sequencer.EventIncreasedBond")
	proto.RegisterType((*EventKickedProposer)(nil), "dymensionxyz.dymension.sequencer.EventKickedProposer")
	proto.RegisterType((*EventProposerChange)(nil), "dymensionxyz.dymension.sequencer.EventProposerChange")
	proto.RegisterType((*OptInStatusChange)(nil), "dymensionxyz.dymension.sequencer.OptInStatusChange")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/sequencer/events.proto", fileDescriptor_1f8a63d7e7167eb3)
}

var fileDescriptor_1f8a63d7e7167eb3 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xb1, 0x72, 0xd3, 0x40,
	0x10, 0xb5, 0x48, 0xc8, 0x38, 0x17, 0x1a, 0x0e, 0x0f, 0x38, 0x29, 0x14, 0x8f, 0x2b, 0x37, 0xbe,
	0x8b, 0x09, 0x93, 0xde, 0xca, 0x50, 0x64, 0x28, 0xc8, 0x38, 0x1d, 0x8d, 0xe7, 0xa4, 0xdb, 0x28,
	0x1a, 0x8f, 0x6e, 0xc5, 0xdd, 0x59, 0xc4, 0xcc, 0xf0, 0x0f, 0x7c, 0x07, 0x35, 0x1f, 0x91, 0x32,
	0x43, 0x45, 0x05, 0x8c, 0xfd, 0x05, 0xfc, 0x01, 0xa3, 0xd3, 0x21, 0x44, 0x41, 0x42, 0x75, 0xda,
	0xdd, 0xf7, 0xf6, 0xbd, 0xd5, 0xdd, 0x92, 0xb1, 0x5c, 0xe5, 0xa0, 0x4c, 0x86, 0xea, 0x7a, 0xf5,
	0x9e, 0x37, 0x01, 0x37, 0xf0, 0x76, 0x09, 0x2a, 0x01, 0xcd, 0xa1, 0x04, 0x65, 0x0d, 0x2b, 0x34,
	0x5a, 0xa4, 0x83, 0x36, 0x9c, 0x35, 0x01, 0x6b, 0xe0, 0x07, 0xfb, 0x09, 0x9a, 0x1c, 0xcd, 0xdc,
	0xe1, 0x79, 0x1d, 0xd4, 0xe4, 0x83, 0x5e, 0x8a, 0x29, 0xd6, 0xf9, 0xea, 0xcb, 0x67, 0xc3, 0x1a,
	0xc3, 0x63, 0x61, 0x80, 0x97, 0x93, 0x18, 0xac, 0x98, 0xf0, 0x04, 0x33, 0xe5, 0xeb, 0xcf, 0x7c,
	0x3d, 0x37, 0x29, 0x2f, 0x27, 0xd5, 0x51, 0x17, 0x86, 0x3f, 0x03, 0x42, 0x5f, 0x56, 0xe6, 0xce,
	0x54, 0xa2, 0x41, 0x18, 0x90, 0x11, 0x2a, 0x49, 0x4f, 0xc8, 0x6e, 0xe3, 0xa6, 0x1f, 0x0c, 0x82,
	0xd1, 0x6e, 0xd4, 0xff, 0xf2, 0x79, 0xdc, 0xf3, 0x56, 0xa6, 0x52, 0x6a, 0x30, 0xe6, 0xc2, 0xea,
	0x4c, 0xa5, 0xb3, 0x3f, 0x50, 0x1a, 0x91, 0x47, 0x42, 0x4a, 0x90, 0x73, 0x91, 0xe3, 0x52, 0xd9,
	0xfe, 0x83, 0x41, 0x30, 0xda, 0x7b, 0xbe, 0xcf, 0x3c, 0xaf, 0xb2, 0xc7, 0xbc, 0x3d, 0x76, 0x8a,
	0x99, 0x8a, 0xb6, 0x6f, 0xbe, 0x1d, 0x76, 0x66, 0x7b, 0x8e, 0x34, 0x75, 0x1c, 0x3a, 0x27, 0xdb,
	0x31, 0x2a, 0xd9, 0xdf, 0x1a, 0x6c, 0xdd, 0xcd, 0x3d, 0xaa, 0xb8, 0x9f, 0xbe, 0x1f, 0x8e, 0xd2,
	0xcc, 0x5e, 0x2d, 0x63, 0x96, 0x60, 0xee, 0xff, 0x95, 0x3f, 0xc6, 0x46, 0x2e, 0xb8, 0x5d, 0x15,
	0x60, 0x1c, 0xc1, 0xcc, 0x5c, 0xe3, 0xe1, 0x07, 0xf2, 0xc4, 0x8d, 0xfc, 0x2a, 0x4b, 0x16, 0x20,
	0xcf, 0x35, 0x16, 0x68, 0x40, 0xd3, 0x23, 0xb2, 0xb3, 0xa8, 0x32, 0xf7, 0x0f, 0xec, 0x71, 0xf4,
	0x05, 0xe9, 0x16, 0x9e, 0xed, 0x26, 0xbd, 0x8b, 0xd3, 0x20, 0x87, 0xef, 0xbc, 0xfc, 0x6f, 0xe1,
	0xd3, 0x2b, 0xa1, 0x52, 0xa8, 0xe4, 0x63, 0xb8, 0x44, 0x0d, 0xf7, 0xcb, 0xd7, 0x38, 0xca, 0xc8,
	0x43, 0x71, 0x69, 0xff, 0x43, 0xbb, 0x86, 0x0d, 0xa7, 0xe4, 0xf1, 0xeb, 0xc2, 0x9e, 0xa9, 0x0b,
	0x2b, 0xec, 0xd2, 0x78, 0xd9, 0xa7, 0x7f, 0xc9, 0x76, 0x9b, 0xe6, 0xbd, 0x76, 0xf3, 0xae, 0x6f,
	0x11, 0x9d, 0xdf, 0xac, 0xc3, 0xe0, 0x76, 0x1d, 0x06, 0x3f, 0xd6, 0x61, 0xf0, 0x71, 0x13, 0x76,
	0x6e, 0x37, 0x61, 0xe7, 0xeb, 0x26, 0xec, 0xbc, 0x39, 0x69, 0x5d, 0xc2, 0x3f, 0xd6, 0xa1, 0x3c,
	0xe6, 0xd7, 0xad, 0x9d, 0x70, 0x17, 0x13, 0xef, 0xb8, 0x77, 0x78, 0xfc, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0x2c, 0xa4, 0xdb, 0xc3, 0x44, 0x03, 0x00, 0x00,
}

func (m *EventIncreasedBond) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventIncreasedBond) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventIncreasedBond) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bond) > 0 {
		for iNdEx := len(m.Bond) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Bond[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.AddedAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Sequencer) > 0 {
		i -= len(m.Sequencer)
		copy(dAtA[i:], m.Sequencer)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Sequencer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventKickedProposer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventKickedProposer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventKickedProposer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Kicker) > 0 {
		i -= len(m.Kicker)
		copy(dAtA[i:], m.Kicker)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Kicker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventProposerChange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventProposerChange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventProposerChange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.After) > 0 {
		i -= len(m.After)
		copy(dAtA[i:], m.After)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.After)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Before) > 0 {
		i -= len(m.Before)
		copy(dAtA[i:], m.Before)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Before)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OptInStatusChange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OptInStatusChange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OptInStatusChange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.After {
		i--
		if m.After {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Before {
		i--
		if m.Before {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventIncreasedBond) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sequencer)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.AddedAmount.Size()
	n += 1 + l + sovEvents(uint64(l))
	if len(m.Bond) > 0 {
		for _, e := range m.Bond {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func (m *EventKickedProposer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Kicker)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventProposerChange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Before)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.After)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *OptInStatusChange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Before {
		n += 2
	}
	if m.After {
		n += 2
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventIncreasedBond) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventIncreasedBond: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventIncreasedBond: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequencer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sequencer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddedAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AddedAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bond", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bond = append(m.Bond, types.Coin{})
			if err := m.Bond[len(m.Bond)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventKickedProposer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventKickedProposer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventKickedProposer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kicker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kicker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventProposerChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventProposerChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventProposerChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Before", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Before = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field After", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.After = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *OptInStatusChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: OptInStatusChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OptInStatusChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Before", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
			m.Before = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field After", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
			m.After = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
