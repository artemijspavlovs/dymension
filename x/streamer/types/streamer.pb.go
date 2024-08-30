// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymensionxyz/dymension/streamer/streamer.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

// EpochPointer is a special object used for the streamer pagination. It helps iterate over
// streams with the specified epoch identifier within one epoch. Additionally, holds coins
// that must be distributed in this epoch.
type EpochPointer struct {
	// StreamID is the ID of a stream.
	StreamId uint64 `protobuf:"varint,1,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
	// GaugeID is the ID of a gauge.
	GaugeId uint64 `protobuf:"varint,2,opt,name=gauge_id,json=gaugeId,proto3" json:"gauge_id,omitempty"`
	// EpochIdentifier is a unique reference to this particular timer.
	EpochIdentifier string `protobuf:"bytes,3,opt,name=epoch_identifier,json=epochIdentifier,proto3" json:"epoch_identifier,omitempty"`
	// EpochDuration is the time in between epoch ticks. It is stored in order to have
	// an ability to sort the EpochPointer slice.
	EpochDuration time.Duration `protobuf:"bytes,4,opt,name=epoch_duration,json=epochDuration,proto3,stdduration" json:"epoch_duration"`
}

func (m *EpochPointer) Reset()         { *m = EpochPointer{} }
func (m *EpochPointer) String() string { return proto.CompactTextString(m) }
func (*EpochPointer) ProtoMessage()    {}
func (*EpochPointer) Descriptor() ([]byte, []int) {
	return fileDescriptor_5216d7d357ab09b8, []int{0}
}
func (m *EpochPointer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpochPointer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpochPointer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpochPointer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochPointer.Merge(m, src)
}
func (m *EpochPointer) XXX_Size() int {
	return m.Size()
}
func (m *EpochPointer) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochPointer.DiscardUnknown(m)
}

var xxx_messageInfo_EpochPointer proto.InternalMessageInfo

func (m *EpochPointer) GetStreamId() uint64 {
	if m != nil {
		return m.StreamId
	}
	return 0
}

func (m *EpochPointer) GetGaugeId() uint64 {
	if m != nil {
		return m.GaugeId
	}
	return 0
}

func (m *EpochPointer) GetEpochIdentifier() string {
	if m != nil {
		return m.EpochIdentifier
	}
	return ""
}

func (m *EpochPointer) GetEpochDuration() time.Duration {
	if m != nil {
		return m.EpochDuration
	}
	return 0
}

func init() {
	proto.RegisterType((*EpochPointer)(nil), "dymensionxyz.dymension.streamer.EpochPointer")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/streamer/streamer.proto", fileDescriptor_5216d7d357ab09b8)
}

var fileDescriptor_5216d7d357ab09b8 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xa9, 0xcc, 0x4d,
	0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0xab, 0xa8, 0xac, 0xd2, 0x87, 0x73, 0xf4, 0x8b, 0x4b, 0x8a, 0x52,
	0x13, 0x73, 0x53, 0x8b, 0xe0, 0x0c, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x79, 0x64, 0xf5,
	0x08, 0xcd, 0x7a, 0x30, 0x65, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0xb5, 0xfa, 0x20, 0x16,
	0x44, 0x9b, 0x94, 0x5c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x98, 0x97, 0x54, 0x9a, 0xa6,
	0x9f, 0x52, 0x5a, 0x94, 0x58, 0x02, 0xd2, 0x08, 0x16, 0x51, 0xda, 0xcb, 0xc8, 0xc5, 0xe3, 0x5a,
	0x90, 0x9f, 0x9c, 0x11, 0x90, 0x9f, 0x99, 0x57, 0x92, 0x5a, 0x24, 0x24, 0xcd, 0xc5, 0x09, 0x31,
	0x32, 0x3e, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x25, 0x88, 0x03, 0x22, 0xe0, 0x99, 0x22,
	0x24, 0xc9, 0xc5, 0x91, 0x9e, 0x58, 0x9a, 0x9e, 0x0a, 0x92, 0x63, 0x02, 0xcb, 0xb1, 0x83, 0xf9,
	0x9e, 0x29, 0x42, 0x9a, 0x5c, 0x02, 0xa9, 0x20, 0x73, 0xe2, 0x33, 0x53, 0x52, 0xf3, 0x4a, 0x32,
	0xd3, 0x32, 0x53, 0x8b, 0x24, 0x98, 0x15, 0x18, 0x35, 0x38, 0x83, 0xf8, 0xc1, 0xe2, 0x9e, 0x70,
	0x61, 0x21, 0x2f, 0x2e, 0x3e, 0x88, 0x52, 0x98, 0x5b, 0x24, 0x58, 0x14, 0x18, 0x35, 0xb8, 0x8d,
	0x24, 0xf5, 0x20, 0x8e, 0xd5, 0x83, 0x39, 0x56, 0xcf, 0x05, 0xaa, 0xc0, 0x89, 0xe3, 0xc4, 0x3d,
	0x79, 0x86, 0x19, 0xf7, 0xe5, 0x19, 0x83, 0x78, 0xc1, 0x5a, 0xe1, 0x12, 0xfe, 0x27, 0x1e, 0xc9,
	0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e,
	0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x9a, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97,
	0x9c, 0x9f, 0xab, 0x8f, 0x23, 0xac, 0xcb, 0x8c, 0xf5, 0x2b, 0x10, 0x01, 0x5e, 0x52, 0x59, 0x90,
	0x5a, 0x9c, 0xc4, 0x06, 0xb6, 0xdc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x96, 0x00, 0x50, 0x36,
	0xa0, 0x01, 0x00, 0x00,
}

func (m *EpochPointer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpochPointer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpochPointer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.EpochDuration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.EpochDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintStreamer(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if len(m.EpochIdentifier) > 0 {
		i -= len(m.EpochIdentifier)
		copy(dAtA[i:], m.EpochIdentifier)
		i = encodeVarintStreamer(dAtA, i, uint64(len(m.EpochIdentifier)))
		i--
		dAtA[i] = 0x1a
	}
	if m.GaugeId != 0 {
		i = encodeVarintStreamer(dAtA, i, uint64(m.GaugeId))
		i--
		dAtA[i] = 0x10
	}
	if m.StreamId != 0 {
		i = encodeVarintStreamer(dAtA, i, uint64(m.StreamId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStreamer(dAtA []byte, offset int, v uint64) int {
	offset -= sovStreamer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EpochPointer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StreamId != 0 {
		n += 1 + sovStreamer(uint64(m.StreamId))
	}
	if m.GaugeId != 0 {
		n += 1 + sovStreamer(uint64(m.GaugeId))
	}
	l = len(m.EpochIdentifier)
	if l > 0 {
		n += 1 + l + sovStreamer(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.EpochDuration)
	n += 1 + l + sovStreamer(uint64(l))
	return n
}

func sovStreamer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStreamer(x uint64) (n int) {
	return sovStreamer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EpochPointer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStreamer
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
			return fmt.Errorf("proto: EpochPointer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpochPointer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamId", wireType)
			}
			m.StreamId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StreamId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GaugeId", wireType)
			}
			m.GaugeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GaugeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamer
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
				return ErrInvalidLengthStreamer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStreamer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamer
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
				return ErrInvalidLengthStreamer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStreamer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.EpochDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStreamer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStreamer
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
func skipStreamer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStreamer
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
					return 0, ErrIntOverflowStreamer
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
					return 0, ErrIntOverflowStreamer
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
				return 0, ErrInvalidLengthStreamer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStreamer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStreamer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStreamer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStreamer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStreamer = fmt.Errorf("proto: unexpected end of group")
)
