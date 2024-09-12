// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymensionxyz/dymension/rollapp/app.proto

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

type App struct {
	// id is the unique App's id in the Rollapp
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// name is the unique App's name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// rollapp_id is the id of the Rollapp the App belongs to
	RollappId string `protobuf:"bytes,3,opt,name=rollapp_id,json=rollappId,proto3" json:"rollapp_id,omitempty"`
	// description is the description of the App
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// image_url is the URL to the App's image
	ImageUrl string `protobuf:"bytes,5,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	// url is the URL to the App's website
	Url string `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	// order is the order of the App in the Rollapp
	Order int32 `protobuf:"varint,7,opt,name=order,proto3" json:"order,omitempty"`
}

func (m *App) Reset()         { *m = App{} }
func (m *App) String() string { return proto.CompactTextString(m) }
func (*App) ProtoMessage()    {}
func (*App) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f01e4af248858b2, []int{0}
}
func (m *App) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *App) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_App.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *App) XXX_Merge(src proto.Message) {
	xxx_messageInfo_App.Merge(m, src)
}
func (m *App) XXX_Size() int {
	return m.Size()
}
func (m *App) XXX_DiscardUnknown() {
	xxx_messageInfo_App.DiscardUnknown(m)
}

var xxx_messageInfo_App proto.InternalMessageInfo

func (m *App) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *App) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *App) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *App) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *App) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *App) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *App) GetOrder() int32 {
	if m != nil {
		return m.Order
	}
	return 0
}

func init() {
	proto.RegisterType((*App)(nil), "dymensionxyz.dymension.rollapp.App")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/rollapp/app.proto", fileDescriptor_8f01e4af248858b2)
}

var fileDescriptor_8f01e4af248858b2 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xbd, 0x4a, 0xc4, 0x40,
	0x14, 0x85, 0x33, 0xf9, 0x59, 0xcd, 0x15, 0x44, 0x06, 0x8b, 0x01, 0x71, 0x08, 0x56, 0xa9, 0x32,
	0xc5, 0xfa, 0x02, 0xda, 0xd9, 0x58, 0x04, 0x6c, 0x6c, 0x96, 0xec, 0xce, 0xb0, 0x0e, 0x24, 0x99,
	0xcb, 0x24, 0x2b, 0x1b, 0x9f, 0xc2, 0x37, 0xf1, 0x35, 0x2c, 0xb7, 0xb4, 0x94, 0xe4, 0x45, 0x64,
	0xc7, 0x10, 0xd2, 0x6c, 0x77, 0xcf, 0x77, 0xce, 0x6d, 0x3e, 0x48, 0x65, 0x57, 0xa9, 0xba, 0xd1,
	0xa6, 0xde, 0x77, 0x1f, 0x62, 0x0a, 0xc2, 0x9a, 0xb2, 0x2c, 0x10, 0x45, 0x81, 0x98, 0xa1, 0x35,
	0xad, 0xa1, 0x7c, 0xbe, 0xcc, 0xa6, 0x90, 0x8d, 0xcb, 0xbb, 0x2f, 0x02, 0xc1, 0x03, 0x22, 0xbd,
	0x04, 0x5f, 0x4b, 0x46, 0x12, 0x92, 0x86, 0xb9, 0xaf, 0x25, 0xa5, 0x10, 0xd6, 0x45, 0xa5, 0x98,
	0x9f, 0x90, 0x34, 0xce, 0xdd, 0x4d, 0x6f, 0x01, 0xc6, 0xb7, 0x95, 0x96, 0x2c, 0x70, 0x4d, 0x3c,
	0x92, 0x27, 0x49, 0x13, 0xb8, 0x90, 0xaa, 0xd9, 0x58, 0x8d, 0xad, 0x36, 0x35, 0x0b, 0x5d, 0x3f,
	0x47, 0xf4, 0x06, 0x62, 0x5d, 0x15, 0x5b, 0xb5, 0xda, 0xd9, 0x92, 0x45, 0xae, 0x3f, 0x77, 0xe0,
	0xc5, 0x96, 0xf4, 0x0a, 0x82, 0x23, 0x5e, 0x38, 0x7c, 0x3c, 0xe9, 0x35, 0x44, 0xc6, 0x4a, 0x65,
	0xd9, 0x59, 0x42, 0xd2, 0x28, 0xff, 0x0f, 0x8f, 0xcf, 0xdf, 0x3d, 0x27, 0x87, 0x9e, 0x93, 0xdf,
	0x9e, 0x93, 0xcf, 0x81, 0x7b, 0x87, 0x81, 0x7b, 0x3f, 0x03, 0xf7, 0x5e, 0xef, 0xb7, 0xba, 0x7d,
	0xdb, 0xad, 0xb3, 0x8d, 0xa9, 0xc4, 0x09, 0x41, 0xef, 0x4b, 0xb1, 0x9f, 0x2c, 0xb5, 0x1d, 0xaa,
	0x66, 0xbd, 0x70, 0xa2, 0x96, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x50, 0xa1, 0x34, 0x60, 0x54,
	0x01, 0x00, 0x00,
}

func (m *App) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *App) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *App) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Order != 0 {
		i = encodeVarintApp(dAtA, i, uint64(m.Order))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintApp(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ImageUrl) > 0 {
		i -= len(m.ImageUrl)
		copy(dAtA[i:], m.ImageUrl)
		i = encodeVarintApp(dAtA, i, uint64(len(m.ImageUrl)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintApp(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintApp(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintApp(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintApp(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintApp(dAtA []byte, offset int, v uint64) int {
	offset -= sovApp(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *App) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovApp(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovApp(uint64(l))
	}
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovApp(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovApp(uint64(l))
	}
	l = len(m.ImageUrl)
	if l > 0 {
		n += 1 + l + sovApp(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovApp(uint64(l))
	}
	if m.Order != 0 {
		n += 1 + sovApp(uint64(m.Order))
	}
	return n
}

func sovApp(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApp(x uint64) (n int) {
	return sovApp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *App) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApp
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
			return fmt.Errorf("proto: App: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: App: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
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
				return ErrInvalidLengthApp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
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
				return ErrInvalidLengthApp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
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
				return ErrInvalidLengthApp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImageUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
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
				return ErrInvalidLengthApp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImageUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
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
				return ErrInvalidLengthApp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Order", wireType)
			}
			m.Order = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Order |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApp
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
func skipApp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApp
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
					return 0, ErrIntOverflowApp
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
					return 0, ErrIntOverflowApp
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
				return 0, ErrInvalidLengthApp
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApp
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApp
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApp        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApp          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApp = fmt.Errorf("proto: unexpected end of group")
)
