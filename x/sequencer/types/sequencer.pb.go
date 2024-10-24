// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymensionxyz/dymension/sequencer/sequencer.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Sequencer defines a sequencer identified by its' address (sequencerAddress).
// The sequencer could be attached to only one rollapp (rollappId).
type Sequencer struct {
	// Address is the bech32-encoded address of the sequencer account which is the account that the message was sent from.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// DymintPubKey is the public key of the sequencers' dymint client, as a Protobuf Any.
	DymintPubKey *types.Any `protobuf:"bytes,2,opt,name=dymintPubKey,proto3" json:"dymintPubKey,omitempty"`
	// RollappId defines the rollapp to which the sequencer belongs.
	RollappId string `protobuf:"bytes,3,opt,name=rollappId,proto3" json:"rollappId,omitempty"`
	// SequencerMetadata defines the extra information for the sequencer.
	Metadata SequencerMetadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata"`
	Proposer bool              `protobuf:"varint,6,opt,name=proposer,proto3" json:"proposer,omitempty"` // Deprecated: Do not use.
	// OperatingStatus is the sequencer status (bonded/unbonded).
	Status OperatingStatus `protobuf:"varint,7,opt,name=status,proto3,enum=dymensionxyz.dymension.sequencer.OperatingStatus" json:"status,omitempty"`
	// OptedIn : when true and bonded, the sequencer can be chosen as proposer or successor
	OptedIn bool `protobuf:"varint,12,opt,name=opted_in,json=optedIn,proto3" json:"opted_in,omitempty"`
	// Tokens: A coins which should always be one dym coin. It's the amount of tokens the sequencer has given to the module.
	Tokens github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,8,rep,name=tokens,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"tokens"`
	// NoticePeriodTime defines the time when the sequencer will finish it's notice period. Zero means not started.
	NoticePeriodTime time.Time `protobuf:"bytes,11,opt,name=notice_period_time,json=noticePeriodTime,proto3,stdtime" json:"notice_period_time"`
}

func (m *Sequencer) Reset()         { *m = Sequencer{} }
func (m *Sequencer) String() string { return proto.CompactTextString(m) }
func (*Sequencer) ProtoMessage()    {}
func (*Sequencer) Descriptor() ([]byte, []int) {
	return fileDescriptor_997b8663a5fc0f58, []int{0}
}
func (m *Sequencer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Sequencer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Sequencer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Sequencer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sequencer.Merge(m, src)
}
func (m *Sequencer) XXX_Size() int {
	return m.Size()
}
func (m *Sequencer) XXX_DiscardUnknown() {
	xxx_messageInfo_Sequencer.DiscardUnknown(m)
}

var xxx_messageInfo_Sequencer proto.InternalMessageInfo

func (m *Sequencer) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Sequencer) GetDymintPubKey() *types.Any {
	if m != nil {
		return m.DymintPubKey
	}
	return nil
}

func (m *Sequencer) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *Sequencer) GetMetadata() SequencerMetadata {
	if m != nil {
		return m.Metadata
	}
	return SequencerMetadata{}
}

// Deprecated: Do not use.
func (m *Sequencer) GetProposer() bool {
	if m != nil {
		return m.Proposer
	}
	return false
}

func (m *Sequencer) GetStatus() OperatingStatus {
	if m != nil {
		return m.Status
	}
	return Unbonded
}

func (m *Sequencer) GetOptedIn() bool {
	if m != nil {
		return m.OptedIn
	}
	return false
}

func (m *Sequencer) GetTokens() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Tokens
	}
	return nil
}

func (m *Sequencer) GetNoticePeriodTime() time.Time {
	if m != nil {
		return m.NoticePeriodTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Sequencer)(nil), "dymensionxyz.dymension.sequencer.Sequencer")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/sequencer/sequencer.proto", fileDescriptor_997b8663a5fc0f58)
}

var fileDescriptor_997b8663a5fc0f58 = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x8e, 0xd3, 0x3c,
	0x14, 0x6d, 0x66, 0xf2, 0xb5, 0xa9, 0x3b, 0xfa, 0x54, 0x45, 0x95, 0x70, 0x2b, 0x94, 0x46, 0xac,
	0xb2, 0x19, 0x7b, 0xda, 0x4a, 0xb0, 0x26, 0xac, 0x5a, 0x84, 0xa8, 0x32, 0xb0, 0x61, 0x53, 0xe5,
	0xc7, 0x84, 0x68, 0x1a, 0x3b, 0xc4, 0x6e, 0x35, 0xe1, 0x29, 0xe6, 0x39, 0x58, 0xb3, 0x67, 0x3b,
	0x62, 0x35, 0x4b, 0x56, 0x0c, 0x6a, 0x5f, 0x04, 0xc5, 0x71, 0x32, 0x05, 0x84, 0xba, 0xb2, 0xef,
	0xcf, 0xb9, 0xd7, 0xe7, 0xf8, 0x80, 0x8b, 0xa8, 0x48, 0x09, 0xe5, 0x09, 0xa3, 0xd7, 0xc5, 0x27,
	0xdc, 0x04, 0x98, 0x93, 0x8f, 0x1b, 0x42, 0x43, 0x92, 0x3f, 0xdc, 0x50, 0x96, 0x33, 0xc1, 0x4c,
	0xfb, 0x10, 0x81, 0x9a, 0x00, 0x35, 0x7d, 0xa3, 0x61, 0xc8, 0x78, 0xca, 0xf8, 0x4a, 0xf6, 0xe3,
	0x2a, 0xa8, 0xc0, 0xa3, 0x61, 0xcc, 0x58, 0xbc, 0x26, 0x58, 0x46, 0xc1, 0xe6, 0x3d, 0xf6, 0x69,
	0xa1, 0x4a, 0x83, 0x98, 0xc5, 0xac, 0x82, 0x94, 0x37, 0x95, 0x1d, 0xff, 0x09, 0x10, 0x49, 0x4a,
	0xb8, 0xf0, 0xd3, 0x4c, 0x35, 0x58, 0xd5, 0x7c, 0x1c, 0xf8, 0x9c, 0xe0, 0xed, 0x24, 0x20, 0xc2,
	0x9f, 0xe0, 0x90, 0x25, 0x54, 0xd5, 0x1f, 0xa9, 0x7a, 0xca, 0x63, 0xbc, 0x9d, 0x94, 0x87, 0x2a,
	0xe0, 0xa3, 0xcc, 0x53, 0x22, 0xfc, 0xc8, 0x17, 0xbe, 0x02, 0x3c, 0x3b, 0x0a, 0x60, 0x19, 0xc9,
	0x7d, 0x91, 0xd0, 0x78, 0xc5, 0x85, 0x2f, 0x36, 0x8a, 0xf4, 0x93, 0xaf, 0x3a, 0xe8, 0x5e, 0xd6,
	0x4d, 0x26, 0x04, 0x1d, 0x3f, 0x8a, 0x72, 0xc2, 0x39, 0xd4, 0x6c, 0xcd, 0xe9, 0x7a, 0x75, 0x68,
	0x7a, 0xe0, 0x2c, 0x2a, 0xd2, 0x84, 0x8a, 0xe5, 0x26, 0x78, 0x49, 0x0a, 0x78, 0x62, 0x6b, 0x4e,
	0x6f, 0x3a, 0x40, 0x95, 0x04, 0xa8, 0x96, 0x00, 0x3d, 0xa7, 0x85, 0x0b, 0xbf, 0x7d, 0x39, 0x1f,
	0x28, 0x69, 0xc3, 0xbc, 0xc8, 0x04, 0x43, 0x15, 0xca, 0xfb, 0x6d, 0x86, 0xf9, 0x18, 0x74, 0x73,
	0xb6, 0x5e, 0xfb, 0x59, 0x36, 0x8f, 0xe0, 0xa9, 0xdc, 0xf7, 0x90, 0x30, 0xdf, 0x02, 0xa3, 0x26,
	0x09, 0x75, 0xb9, 0x6d, 0x86, 0x8e, 0x7d, 0x2f, 0x6a, 0xa8, 0xbc, 0x52, 0x50, 0x57, 0xbf, 0xfd,
	0x31, 0x6e, 0x79, 0xcd, 0x28, 0xd3, 0x02, 0x46, 0x96, 0xb3, 0x8c, 0x71, 0x92, 0xc3, 0xb6, 0xad,
	0x39, 0x86, 0x7b, 0x02, 0x35, 0xaf, 0xc9, 0x99, 0x73, 0xd0, 0xae, 0x04, 0x82, 0x1d, 0x5b, 0x73,
	0xfe, 0x9f, 0x4e, 0x8e, 0x2f, 0x7d, 0x5d, 0x4b, 0x7b, 0x29, 0x81, 0x9e, 0x1a, 0x60, 0x0e, 0x81,
	0xc1, 0x32, 0x41, 0xa2, 0x55, 0x42, 0xe1, 0x59, 0xb9, 0xca, 0xeb, 0xc8, 0x78, 0x4e, 0xcd, 0x10,
	0xb4, 0x05, 0xbb, 0x22, 0x94, 0x43, 0xc3, 0x3e, 0x75, 0x7a, 0xd3, 0x21, 0x52, 0x7a, 0x95, 0x56,
	0x41, 0xca, 0x2a, 0xe8, 0x05, 0x4b, 0xa8, 0x7b, 0x51, 0x12, 0xf8, 0x7c, 0x3f, 0x76, 0xe2, 0x44,
	0x7c, 0xd8, 0x04, 0x28, 0x64, 0xa9, 0xf2, 0xad, 0x3a, 0xce, 0x79, 0x74, 0x85, 0x45, 0x91, 0x11,
	0x2e, 0x01, 0xdc, 0x53, 0xa3, 0x4d, 0x0f, 0x98, 0x94, 0x89, 0x24, 0x24, 0xab, 0x8c, 0xe4, 0x09,
	0x8b, 0x56, 0xa5, 0x3f, 0x61, 0x4f, 0x6a, 0x39, 0xfa, 0xeb, 0xe7, 0xde, 0xd4, 0xe6, 0x75, 0x8d,
	0x72, 0xe3, 0xcd, 0xfd, 0x58, 0xf3, 0xfa, 0x15, 0x7e, 0x29, 0xe1, 0x65, 0xc3, 0x42, 0x37, 0xfe,
	0xeb, 0xb7, 0x17, 0xba, 0xd1, 0xed, 0x83, 0x85, 0x6e, 0x80, 0x7e, 0xcf, 0x5d, 0xde, 0xee, 0x2c,
	0xed, 0x6e, 0x67, 0x69, 0x3f, 0x77, 0x96, 0x76, 0xb3, 0xb7, 0x5a, 0x77, 0x7b, 0xab, 0xf5, 0x7d,
	0x6f, 0xb5, 0xde, 0x3d, 0x3d, 0x78, 0xf1, 0x3f, 0xfc, 0xb9, 0x9d, 0xe1, 0xeb, 0x03, 0x93, 0x4a,
	0x16, 0x41, 0x5b, 0xbe, 0x69, 0xf6, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x73, 0xdc, 0x97, 0x00,
	0x04, 0x00, 0x00,
}

func (m *Sequencer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Sequencer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Sequencer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OptedIn {
		i--
		if m.OptedIn {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x60
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.NoticePeriodTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.NoticePeriodTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintSequencer(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x5a
	if len(m.Tokens) > 0 {
		for iNdEx := len(m.Tokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSequencer(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if m.Status != 0 {
		i = encodeVarintSequencer(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x38
	}
	if m.Proposer {
		i--
		if m.Proposer {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	{
		size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSequencer(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintSequencer(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.DymintPubKey != nil {
		{
			size, err := m.DymintPubKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSequencer(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintSequencer(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSequencer(dAtA []byte, offset int, v uint64) int {
	offset -= sovSequencer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Sequencer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSequencer(uint64(l))
	}
	if m.DymintPubKey != nil {
		l = m.DymintPubKey.Size()
		n += 1 + l + sovSequencer(uint64(l))
	}
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovSequencer(uint64(l))
	}
	l = m.Metadata.Size()
	n += 1 + l + sovSequencer(uint64(l))
	if m.Proposer {
		n += 2
	}
	if m.Status != 0 {
		n += 1 + sovSequencer(uint64(m.Status))
	}
	if len(m.Tokens) > 0 {
		for _, e := range m.Tokens {
			l = e.Size()
			n += 1 + l + sovSequencer(uint64(l))
		}
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.NoticePeriodTime)
	n += 1 + l + sovSequencer(uint64(l))
	if m.OptedIn {
		n += 2
	}
	return n
}

func sovSequencer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSequencer(x uint64) (n int) {
	return sovSequencer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Sequencer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSequencer
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
			return fmt.Errorf("proto: Sequencer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Sequencer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DymintPubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DymintPubKey == nil {
				m.DymintPubKey = &types.Any{}
			}
			if err := m.DymintPubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
			m.Proposer = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= OperatingStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tokens = append(m.Tokens, types1.Coin{})
			if err := m.Tokens[len(m.Tokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoticePeriodTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.NoticePeriodTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptedIn", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
			m.OptedIn = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSequencer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSequencer
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
func skipSequencer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSequencer
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
					return 0, ErrIntOverflowSequencer
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
					return 0, ErrIntOverflowSequencer
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
				return 0, ErrInvalidLengthSequencer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSequencer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSequencer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSequencer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSequencer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSequencer = fmt.Errorf("proto: unexpected end of group")
)
