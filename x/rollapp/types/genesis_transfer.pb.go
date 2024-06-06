// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymension/rollapp/genesis_transfer.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Bookkeeping for the genesis transfer bridge protocol.
// Each rollapp will have one of these items corresponding to it.
type GenesisTransfers struct {
	RollappID string `protobuf:"bytes,1,opt,name=rollappID,proto3" json:"rollappID,omitempty"`
	// The total number of incoming ibc transfers to be fast tracked in the genesis transfer period
	NumTotal uint64 `protobuf:"varint,2,opt,name=numTotal,proto3" json:"numTotal,omitempty"`
	// The number of transfers already processed, when this number reaches numTotal the genesis transfer window closes.
	NumReceived uint64 `protobuf:"varint,3,opt,name=numReceived,proto3" json:"numReceived,omitempty"`
	// The indexes of genesis transfers received thus far during the genesis transfer period
	// This cannot be implemented on the rollapp type, because we index into this like a map, but the rollapp is read/writ as a whole
	Received []uint64 `protobuf:"varint,4,rep,packed,name=received,proto3" json:"received,omitempty"`
}

func (m *GenesisTransfers) Reset()         { *m = GenesisTransfers{} }
func (m *GenesisTransfers) String() string { return proto.CompactTextString(m) }
func (*GenesisTransfers) ProtoMessage()    {}
func (*GenesisTransfers) Descriptor() ([]byte, []int) {
	return fileDescriptor_abbd5969075b03fa, []int{0}
}
func (m *GenesisTransfers) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisTransfers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisTransfers.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisTransfers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisTransfers.Merge(m, src)
}
func (m *GenesisTransfers) XXX_Size() int {
	return m.Size()
}
func (m *GenesisTransfers) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisTransfers.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisTransfers proto.InternalMessageInfo

func (m *GenesisTransfers) GetRollappID() string {
	if m != nil {
		return m.RollappID
	}
	return ""
}

func (m *GenesisTransfers) GetNumTotal() uint64 {
	if m != nil {
		return m.NumTotal
	}
	return 0
}

func (m *GenesisTransfers) GetNumReceived() uint64 {
	if m != nil {
		return m.NumReceived
	}
	return 0
}

func (m *GenesisTransfers) GetReceived() []uint64 {
	if m != nil {
		return m.Received
	}
	return nil
}

// BlockHeightToTransferGenesisFinalizationQueue defines a map from block height to list of rollapps to finalize
// The queue contains rollapps whose genesis transfer window has closed, and are awaiting
// the end of their dispute period, and the opening of their bridge.
type BlockHeightToTransferGenesisFinalizations struct {
	// height is the hub height from which the transfers will be enabled
	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	// ids of rollapps for whom to finalize their transfer genesis at this height
	Rollapps []string `protobuf:"bytes,2,rep,name=rollapps,proto3" json:"rollapps,omitempty"`
}

func (m *BlockHeightToTransferGenesisFinalizations) Reset() {
	*m = BlockHeightToTransferGenesisFinalizations{}
}
func (m *BlockHeightToTransferGenesisFinalizations) String() string {
	return proto.CompactTextString(m)
}
func (*BlockHeightToTransferGenesisFinalizations) ProtoMessage() {}
func (*BlockHeightToTransferGenesisFinalizations) Descriptor() ([]byte, []int) {
	return fileDescriptor_abbd5969075b03fa, []int{1}
}
func (m *BlockHeightToTransferGenesisFinalizations) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockHeightToTransferGenesisFinalizations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockHeightToTransferGenesisFinalizations.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockHeightToTransferGenesisFinalizations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeightToTransferGenesisFinalizations.Merge(m, src)
}
func (m *BlockHeightToTransferGenesisFinalizations) XXX_Size() int {
	return m.Size()
}
func (m *BlockHeightToTransferGenesisFinalizations) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeightToTransferGenesisFinalizations.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeightToTransferGenesisFinalizations proto.InternalMessageInfo

func (m *BlockHeightToTransferGenesisFinalizations) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockHeightToTransferGenesisFinalizations) GetRollapps() []string {
	if m != nil {
		return m.Rollapps
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisTransfers)(nil), "dymensionxyz.dymension.rollapp.GenesisTransfers")
	proto.RegisterType((*BlockHeightToTransferGenesisFinalizations)(nil), "dymensionxyz.dymension.rollapp.BlockHeightToTransferGenesisFinalizations")
}

func init() {
	proto.RegisterFile("dymension/rollapp/genesis_transfer.proto", fileDescriptor_abbd5969075b03fa)
}

var fileDescriptor_abbd5969075b03fa = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0xbf, 0x6a, 0x02, 0x31,
	0x1c, 0xbe, 0xe8, 0x21, 0x9a, 0x2e, 0xe5, 0x28, 0xe5, 0x90, 0x12, 0x83, 0xd3, 0x75, 0xb9, 0x0c,
	0xf6, 0x09, 0xa4, 0xf4, 0xcf, 0xd2, 0xe1, 0x70, 0xea, 0x22, 0x51, 0xe3, 0x19, 0x7a, 0x97, 0x1c,
	0x49, 0x14, 0xf5, 0x29, 0x7c, 0x2c, 0x47, 0xc7, 0x4e, 0xa5, 0xe8, 0x8b, 0x94, 0xc6, 0x18, 0x05,
	0xa7, 0xe4, 0xfb, 0xf3, 0xcb, 0xf7, 0xf1, 0x0b, 0x4c, 0x26, 0xab, 0x92, 0x09, 0xcd, 0xa5, 0x20,
	0x4a, 0x16, 0x05, 0xad, 0x2a, 0x92, 0x33, 0xc1, 0x34, 0xd7, 0x43, 0xa3, 0xa8, 0xd0, 0x53, 0xa6,
	0xd2, 0x4a, 0x49, 0x23, 0x23, 0xe4, 0x9d, 0xcb, 0xd5, 0x3a, 0xf5, 0x20, 0x75, 0x63, 0xed, 0xbb,
	0x5c, 0xe6, 0xd2, 0x5a, 0xc9, 0xff, 0xed, 0x38, 0xd5, 0x46, 0xd7, 0xef, 0x57, 0x54, 0xd1, 0x52,
	0x3b, 0xbd, 0x73, 0xad, 0xbb, 0xd3, 0x19, 0xba, 0xd7, 0x06, 0x6d, 0xa8, 0x61, 0x43, 0x2e, 0xa6,
	0x2e, 0xa4, 0xbb, 0x01, 0xf0, 0xf6, 0xf5, 0xd8, 0x7a, 0xe0, 0x4a, 0xeb, 0xe8, 0x01, 0xb6, 0xdc,
	0xc0, 0xfb, 0x73, 0x0c, 0x30, 0x48, 0x5a, 0xd9, 0x99, 0x88, 0xda, 0xb0, 0x29, 0xe6, 0xe5, 0x40,
	0x1a, 0x5a, 0xc4, 0x35, 0x0c, 0x92, 0x30, 0xf3, 0x38, 0xc2, 0xf0, 0x46, 0xcc, 0xcb, 0x8c, 0x8d,
	0x19, 0x5f, 0xb0, 0x49, 0x5c, 0xb7, 0xf2, 0x25, 0x15, 0x61, 0xd8, 0x54, 0x27, 0x39, 0xc4, 0xf5,
	0x24, 0xec, 0x87, 0xdb, 0x9f, 0x4e, 0x90, 0x79, 0xb6, 0xcb, 0xe0, 0x63, 0xbf, 0x90, 0xe3, 0xaf,
	0x37, 0xc6, 0xf3, 0x99, 0x19, 0xc8, 0x53, 0x2f, 0x57, 0xf3, 0x85, 0x0b, 0x5a, 0xf0, 0x35, 0x35,
	0x5c, 0x0a, 0x1d, 0xdd, 0xc3, 0xc6, 0xcc, 0xfa, 0x6c, 0xcf, 0x30, 0x73, 0xc8, 0xc6, 0x1c, 0x1b,
	0xeb, 0xb8, 0x86, 0xeb, 0x49, 0xcb, 0xc7, 0x38, 0xb6, 0xff, 0xb1, 0xdd, 0x23, 0xb0, 0xdb, 0x23,
	0xf0, 0xbb, 0x47, 0x60, 0x73, 0x40, 0xc1, 0xee, 0x80, 0x82, 0xef, 0x03, 0x0a, 0x3e, 0x9f, 0x72,
	0x6e, 0x66, 0xf3, 0x51, 0x3a, 0x96, 0x25, 0xb9, 0xfc, 0xb9, 0x33, 0x20, 0x8b, 0x1e, 0x59, 0xfa,
	0xa5, 0x9a, 0x55, 0xc5, 0xf4, 0xa8, 0x61, 0x17, 0xda, 0xfb, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xa0,
	0xd1, 0x43, 0xba, 0x17, 0x02, 0x00, 0x00,
}

func (m *GenesisTransfers) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisTransfers) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisTransfers) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Received) > 0 {
		dAtA2 := make([]byte, len(m.Received)*10)
		var j1 int
		for _, num := range m.Received {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintGenesisTransfer(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x22
	}
	if m.NumReceived != 0 {
		i = encodeVarintGenesisTransfer(dAtA, i, uint64(m.NumReceived))
		i--
		dAtA[i] = 0x18
	}
	if m.NumTotal != 0 {
		i = encodeVarintGenesisTransfer(dAtA, i, uint64(m.NumTotal))
		i--
		dAtA[i] = 0x10
	}
	if len(m.RollappID) > 0 {
		i -= len(m.RollappID)
		copy(dAtA[i:], m.RollappID)
		i = encodeVarintGenesisTransfer(dAtA, i, uint64(len(m.RollappID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BlockHeightToTransferGenesisFinalizations) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHeightToTransferGenesisFinalizations) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockHeightToTransferGenesisFinalizations) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rollapps) > 0 {
		for iNdEx := len(m.Rollapps) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Rollapps[iNdEx])
			copy(dAtA[i:], m.Rollapps[iNdEx])
			i = encodeVarintGenesisTransfer(dAtA, i, uint64(len(m.Rollapps[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Height != 0 {
		i = encodeVarintGenesisTransfer(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesisTransfer(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesisTransfer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisTransfers) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RollappID)
	if l > 0 {
		n += 1 + l + sovGenesisTransfer(uint64(l))
	}
	if m.NumTotal != 0 {
		n += 1 + sovGenesisTransfer(uint64(m.NumTotal))
	}
	if m.NumReceived != 0 {
		n += 1 + sovGenesisTransfer(uint64(m.NumReceived))
	}
	if len(m.Received) > 0 {
		l = 0
		for _, e := range m.Received {
			l += sovGenesisTransfer(uint64(e))
		}
		n += 1 + sovGenesisTransfer(uint64(l)) + l
	}
	return n
}

func (m *BlockHeightToTransferGenesisFinalizations) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovGenesisTransfer(uint64(m.Height))
	}
	if len(m.Rollapps) > 0 {
		for _, s := range m.Rollapps {
			l = len(s)
			n += 1 + l + sovGenesisTransfer(uint64(l))
		}
	}
	return n
}

func sovGenesisTransfer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesisTransfer(x uint64) (n int) {
	return sovGenesisTransfer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisTransfers) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesisTransfer
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
			return fmt.Errorf("proto: GenesisTransfers: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisTransfers: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisTransfer
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
				return ErrInvalidLengthGenesisTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumTotal", wireType)
			}
			m.NumTotal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumTotal |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumReceived", wireType)
			}
			m.NumReceived = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumReceived |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesisTransfer
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Received = append(m.Received, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesisTransfer
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGenesisTransfer
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGenesisTransfer
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Received) == 0 {
					m.Received = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesisTransfer
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Received = append(m.Received, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Received", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesisTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesisTransfer
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
func (m *BlockHeightToTransferGenesisFinalizations) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesisTransfer
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
			return fmt.Errorf("proto: BlockHeightToTransferGenesisFinalizations: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHeightToTransferGenesisFinalizations: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisTransfer
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rollapps", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisTransfer
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
				return ErrInvalidLengthGenesisTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rollapps = append(m.Rollapps, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesisTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesisTransfer
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
func skipGenesisTransfer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesisTransfer
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
					return 0, ErrIntOverflowGenesisTransfer
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
					return 0, ErrIntOverflowGenesisTransfer
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
				return 0, ErrInvalidLengthGenesisTransfer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesisTransfer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesisTransfer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesisTransfer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesisTransfer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesisTransfer = fmt.Errorf("proto: unexpected end of group")
)
