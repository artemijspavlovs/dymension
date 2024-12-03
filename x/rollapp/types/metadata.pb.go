// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymensionxyz/dymension/rollapp/metadata.proto

package types

import (
	fmt "fmt"
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

type RollappMetadata struct {
	// website is the rollapp website
	Website string `protobuf:"bytes,1,opt,name=website,proto3" json:"website,omitempty"`
	// description is the rollapp description. should be limited to 512 chars
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// logo_url is the rollapp logo url
	LogoUrl string `protobuf:"bytes,3,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	// telegram is the rollapp telegram link
	Telegram string `protobuf:"bytes,4,opt,name=telegram,proto3" json:"telegram,omitempty"`
	// x is the rollapp twitter link
	X string `protobuf:"bytes,5,opt,name=x,proto3" json:"x,omitempty"`
	// genesis_url has the genesis file
	GenesisUrl string `protobuf:"bytes,6,opt,name=genesis_url,json=genesisUrl,proto3" json:"genesis_url,omitempty"`
	// display_name is a non semantic name for displaying on gui etc. Size limited.
	DisplayName string `protobuf:"bytes,7,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// tagline is a non semantic tagline/catch-phrase. Size limited.
	Tagline string `protobuf:"bytes,8,opt,name=tagline,proto3" json:"tagline,omitempty"`
	// explorer_url is the rollapp's block explorer link
	ExplorerUrl string `protobuf:"bytes,9,opt,name=explorer_url,json=explorerUrl,proto3" json:"explorer_url,omitempty"`
	// fee_denom is the base denom for fees
	FeeDenom *DenomMetadata `protobuf:"bytes,10,opt,name=fee_denom,json=feeDenom,proto3" json:"fee_denom,omitempty"`
	// Tags is a list of tags that can be used to filter rollapps.
	Tags []string `protobuf:"bytes,11,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (m *RollappMetadata) Reset()         { *m = RollappMetadata{} }
func (m *RollappMetadata) String() string { return proto.CompactTextString(m) }
func (*RollappMetadata) ProtoMessage()    {}
func (*RollappMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_481e8ecd49e74695, []int{0}
}
func (m *RollappMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RollappMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RollappMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RollappMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RollappMetadata.Merge(m, src)
}
func (m *RollappMetadata) XXX_Size() int {
	return m.Size()
}
func (m *RollappMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_RollappMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_RollappMetadata proto.InternalMessageInfo

func (m *RollappMetadata) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *RollappMetadata) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RollappMetadata) GetLogoUrl() string {
	if m != nil {
		return m.LogoUrl
	}
	return ""
}

func (m *RollappMetadata) GetTelegram() string {
	if m != nil {
		return m.Telegram
	}
	return ""
}

func (m *RollappMetadata) GetX() string {
	if m != nil {
		return m.X
	}
	return ""
}

func (m *RollappMetadata) GetGenesisUrl() string {
	if m != nil {
		return m.GenesisUrl
	}
	return ""
}

func (m *RollappMetadata) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *RollappMetadata) GetTagline() string {
	if m != nil {
		return m.Tagline
	}
	return ""
}

func (m *RollappMetadata) GetExplorerUrl() string {
	if m != nil {
		return m.ExplorerUrl
	}
	return ""
}

func (m *RollappMetadata) GetFeeDenom() *DenomMetadata {
	if m != nil {
		return m.FeeDenom
	}
	return nil
}

func (m *RollappMetadata) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type DenomMetadata struct {
	Display  string `protobuf:"bytes,1,opt,name=display,proto3" json:"display,omitempty"`
	Base     string `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	Exponent uint32 `protobuf:"varint,3,opt,name=exponent,proto3" json:"exponent,omitempty"`
}

func (m *DenomMetadata) Reset()         { *m = DenomMetadata{} }
func (m *DenomMetadata) String() string { return proto.CompactTextString(m) }
func (*DenomMetadata) ProtoMessage()    {}
func (*DenomMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_481e8ecd49e74695, []int{1}
}
func (m *DenomMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DenomMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DenomMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DenomMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenomMetadata.Merge(m, src)
}
func (m *DenomMetadata) XXX_Size() int {
	return m.Size()
}
func (m *DenomMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DenomMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DenomMetadata proto.InternalMessageInfo

func (m *DenomMetadata) GetDisplay() string {
	if m != nil {
		return m.Display
	}
	return ""
}

func (m *DenomMetadata) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

func (m *DenomMetadata) GetExponent() uint32 {
	if m != nil {
		return m.Exponent
	}
	return 0
}

func init() {
	proto.RegisterType((*RollappMetadata)(nil), "dymensionxyz.dymension.rollapp.RollappMetadata")
	proto.RegisterType((*DenomMetadata)(nil), "dymensionxyz.dymension.rollapp.DenomMetadata")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/rollapp/metadata.proto", fileDescriptor_481e8ecd49e74695)
}

var fileDescriptor_481e8ecd49e74695 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xb1, 0x8e, 0xd4, 0x30,
	0x10, 0x86, 0xd7, 0x77, 0xc7, 0x5d, 0xe2, 0xdc, 0x09, 0xc9, 0x95, 0xa1, 0x08, 0xe1, 0xaa, 0x6d,
	0x2e, 0x91, 0x38, 0x9e, 0x00, 0x51, 0x21, 0x71, 0x45, 0x24, 0x0a, 0x68, 0x56, 0xce, 0x65, 0x36,
	0x58, 0x72, 0x6c, 0xcb, 0xf6, 0x42, 0xc2, 0x53, 0xf0, 0x58, 0x94, 0x5b, 0x52, 0xa2, 0xdd, 0x47,
	0xe0, 0x05, 0x90, 0x1d, 0x27, 0x5a, 0x0a, 0xe8, 0xe6, 0xfb, 0x3d, 0xff, 0x68, 0xfc, 0x6b, 0xf0,
	0x5d, 0x3b, 0xf6, 0x20, 0x2d, 0x57, 0x72, 0x18, 0xbf, 0x55, 0x0b, 0x54, 0x46, 0x09, 0xc1, 0xb4,
	0xae, 0x7a, 0x70, 0xac, 0x65, 0x8e, 0x95, 0xda, 0x28, 0xa7, 0x48, 0x7e, 0xda, 0x5e, 0x2e, 0x50,
	0xc6, 0xf6, 0xdb, 0xdf, 0x67, 0xf8, 0x69, 0x3d, 0xd5, 0xef, 0xa3, 0x93, 0x50, 0x7c, 0xf5, 0x15,
	0x1a, 0xcb, 0x1d, 0x50, 0x54, 0xa0, 0x75, 0x5a, 0xcf, 0x48, 0x0a, 0x9c, 0xb5, 0x60, 0x1f, 0x0d,
	0xd7, 0x8e, 0x2b, 0x49, 0xcf, 0xc2, 0xeb, 0xa9, 0x44, 0x9e, 0xe1, 0x44, 0xa8, 0x4e, 0x6d, 0x76,
	0x46, 0xd0, 0xf3, 0xc9, 0xec, 0xf9, 0x83, 0x11, 0xe4, 0x39, 0x4e, 0x1c, 0x08, 0xe8, 0x0c, 0xeb,
	0xe9, 0x45, 0x78, 0x5a, 0x98, 0x5c, 0x63, 0x34, 0xd0, 0x27, 0x41, 0x44, 0x03, 0x79, 0x81, 0xb3,
	0x0e, 0x24, 0x58, 0x6e, 0xc3, 0x9c, 0xcb, 0xa0, 0xe3, 0x28, 0xf9, 0x51, 0x2f, 0xf1, 0x75, 0xcb,
	0xad, 0x16, 0x6c, 0xdc, 0x48, 0xd6, 0x03, 0xbd, 0x8a, 0x8b, 0x4c, 0xda, 0x03, 0xeb, 0xc1, 0x7f,
	0xc2, 0xb1, 0x4e, 0x70, 0x09, 0x34, 0x99, 0xf6, 0x88, 0xe8, 0xcd, 0x30, 0x68, 0xa1, 0x0c, 0x98,
	0x30, 0x3e, 0x9d, 0xcc, 0xb3, 0xe6, 0xe7, 0xbf, 0xc3, 0xe9, 0x16, 0x60, 0xd3, 0x82, 0x54, 0x3d,
	0xc5, 0x05, 0x5a, 0x67, 0xaf, 0xee, 0xca, 0xff, 0x27, 0x59, 0xbe, 0xf5, 0xcd, 0x73, 0x86, 0x75,
	0xb2, 0x05, 0x08, 0x0a, 0x21, 0xf8, 0xc2, 0xb1, 0xce, 0xd2, 0xac, 0x38, 0x5f, 0xa7, 0x75, 0xa8,
	0x6f, 0x3f, 0xe2, 0x9b, 0xbf, 0xda, 0xfd, 0xb6, 0x71, 0xf9, 0x39, 0xf2, 0x88, 0xde, 0xde, 0x30,
	0x0b, 0x31, 0xeb, 0x50, 0xfb, 0x24, 0x61, 0xd0, 0x4a, 0x82, 0x74, 0x21, 0xe4, 0x9b, 0x7a, 0xe1,
	0x37, 0x0f, 0x3f, 0x0e, 0x39, 0xda, 0x1f, 0x72, 0xf4, 0xeb, 0x90, 0xa3, 0xef, 0xc7, 0x7c, 0xb5,
	0x3f, 0xe6, 0xab, 0x9f, 0xc7, 0x7c, 0xf5, 0xe9, 0x75, 0xc7, 0xdd, 0xe7, 0x5d, 0x53, 0x3e, 0xaa,
	0xbe, 0xfa, 0xc7, 0x11, 0x7d, 0xb9, 0xaf, 0x86, 0xe5, 0x92, 0xdc, 0xa8, 0xc1, 0x36, 0x97, 0xe1,
	0x8e, 0xee, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x3a, 0xf5, 0xd3, 0x78, 0x02, 0x00, 0x00,
}

func (m *RollappMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RollappMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RollappMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tags) > 0 {
		for iNdEx := len(m.Tags) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Tags[iNdEx])
			copy(dAtA[i:], m.Tags[iNdEx])
			i = encodeVarintMetadata(dAtA, i, uint64(len(m.Tags[iNdEx])))
			i--
			dAtA[i] = 0x5a
		}
	}
	if m.FeeDenom != nil {
		{
			size, err := m.FeeDenom.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMetadata(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if len(m.ExplorerUrl) > 0 {
		i -= len(m.ExplorerUrl)
		copy(dAtA[i:], m.ExplorerUrl)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.ExplorerUrl)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Tagline) > 0 {
		i -= len(m.Tagline)
		copy(dAtA[i:], m.Tagline)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Tagline)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.DisplayName) > 0 {
		i -= len(m.DisplayName)
		copy(dAtA[i:], m.DisplayName)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.DisplayName)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.GenesisUrl) > 0 {
		i -= len(m.GenesisUrl)
		copy(dAtA[i:], m.GenesisUrl)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.GenesisUrl)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.X) > 0 {
		i -= len(m.X)
		copy(dAtA[i:], m.X)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.X)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Telegram) > 0 {
		i -= len(m.Telegram)
		copy(dAtA[i:], m.Telegram)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Telegram)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.LogoUrl) > 0 {
		i -= len(m.LogoUrl)
		copy(dAtA[i:], m.LogoUrl)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.LogoUrl)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DenomMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DenomMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DenomMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Exponent != 0 {
		i = encodeVarintMetadata(dAtA, i, uint64(m.Exponent))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Base) > 0 {
		i -= len(m.Base)
		copy(dAtA[i:], m.Base)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Base)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Display) > 0 {
		i -= len(m.Display)
		copy(dAtA[i:], m.Display)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Display)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RollappMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.LogoUrl)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.Telegram)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.X)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.GenesisUrl)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.DisplayName)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.Tagline)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.ExplorerUrl)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	if m.FeeDenom != nil {
		l = m.FeeDenom.Size()
		n += 1 + l + sovMetadata(uint64(l))
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			l = len(s)
			n += 1 + l + sovMetadata(uint64(l))
		}
	}
	return n
}

func (m *DenomMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Display)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.Base)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	if m.Exponent != 0 {
		n += 1 + sovMetadata(uint64(m.Exponent))
	}
	return n
}

func sovMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMetadata(x uint64) (n int) {
	return sovMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RollappMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: RollappMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RollappMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogoUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogoUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Telegram", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Telegram = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.X = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenesisUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tagline", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tagline = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExplorerUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExplorerUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeDenom", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FeeDenom == nil {
				m.FeeDenom = &DenomMetadata{}
			}
			if err := m.FeeDenom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetadata
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
func (m *DenomMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: DenomMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DenomMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Display", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Display = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Base", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Base = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exponent", wireType)
			}
			m.Exponent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Exponent |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetadata
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
func skipMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetadata
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
					return 0, ErrIntOverflowMetadata
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
					return 0, ErrIntOverflowMetadata
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
				return 0, ErrInvalidLengthMetadata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMetadata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMetadata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMetadata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetadata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMetadata = fmt.Errorf("proto: unexpected end of group")
)
