// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/poolmanager/v1beta1/gov.proto

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

// DenomPairTakerFeeProposal is a type for adding/removing a custom taker fee(s)
// for one or more denom pairs.
type DenomPairTakerFeeProposal struct {
	Title             string              `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description       string              `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	DenomPairTakerFee []DenomPairTakerFee `protobuf:"bytes,3,rep,name=denom_pair_taker_fee,json=denomPairTakerFee,proto3" json:"denom_pair_taker_fee"`
}

func (m *DenomPairTakerFeeProposal) Reset()      { *m = DenomPairTakerFeeProposal{} }
func (*DenomPairTakerFeeProposal) ProtoMessage() {}
func (*DenomPairTakerFeeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_c95b3c1cda2a8632, []int{0}
}
func (m *DenomPairTakerFeeProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DenomPairTakerFeeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DenomPairTakerFeeProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DenomPairTakerFeeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenomPairTakerFeeProposal.Merge(m, src)
}
func (m *DenomPairTakerFeeProposal) XXX_Size() int {
	return m.Size()
}
func (m *DenomPairTakerFeeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_DenomPairTakerFeeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_DenomPairTakerFeeProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DenomPairTakerFeeProposal)(nil), "osmosis.poolmanager.v1beta1.DenomPairTakerFeeProposal")
}

func init() {
	proto.RegisterFile("osmosis/poolmanager/v1beta1/gov.proto", fileDescriptor_c95b3c1cda2a8632)
}

var fileDescriptor_c95b3c1cda2a8632 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcd, 0x2f, 0xce, 0xcd,
	0x2f, 0xce, 0x2c, 0xd6, 0x2f, 0xc8, 0xcf, 0xcf, 0xc9, 0x4d, 0xcc, 0x4b, 0x4c, 0x4f, 0x2d, 0xd2,
	0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x4f, 0xcf, 0x2f, 0xd3, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0x86, 0x2a, 0xd3, 0x43, 0x52, 0xa6, 0x07, 0x55, 0x26, 0x25, 0x92, 0x9e, 0x9f,
	0x9e, 0x0f, 0x56, 0xa7, 0x0f, 0x62, 0x41, 0xb4, 0x48, 0xa9, 0xe0, 0x33, 0xb9, 0xa4, 0x02, 0xa2,
	0x4a, 0xe9, 0x08, 0x23, 0x97, 0xa4, 0x4b, 0x6a, 0x5e, 0x7e, 0x6e, 0x40, 0x62, 0x66, 0x51, 0x48,
	0x62, 0x76, 0x6a, 0x91, 0x5b, 0x6a, 0x6a, 0x40, 0x51, 0x7e, 0x41, 0x7e, 0x71, 0x62, 0x8e, 0x90,
	0x08, 0x17, 0x6b, 0x49, 0x66, 0x49, 0x4e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x84,
	0x23, 0xa4, 0xc0, 0xc5, 0x9d, 0x92, 0x5a, 0x9c, 0x5c, 0x94, 0x59, 0x50, 0x92, 0x99, 0x9f, 0x27,
	0xc1, 0x04, 0x96, 0x43, 0x16, 0x12, 0x4a, 0xe5, 0x12, 0x49, 0x01, 0x19, 0x1a, 0x5f, 0x90, 0x98,
	0x59, 0x14, 0x5f, 0x02, 0x32, 0x36, 0x3e, 0x2d, 0x35, 0x55, 0x82, 0x59, 0x81, 0x59, 0x83, 0xdb,
	0x48, 0x4f, 0x0f, 0x8f, 0x6f, 0xf4, 0x30, 0x5c, 0xe3, 0xc4, 0x72, 0xe2, 0x9e, 0x3c, 0x43, 0x90,
	0x60, 0x0a, 0xba, 0x84, 0x15, 0x47, 0xc7, 0x02, 0x79, 0x86, 0x19, 0x0b, 0xe4, 0x19, 0x9c, 0x02,
	0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5,
	0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x3c, 0x3d, 0xb3, 0x24, 0xa3,
	0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0x6a, 0xad, 0x6e, 0x4e, 0x62, 0x52, 0x31, 0x8c, 0xa3,
	0x5f, 0x66, 0x6c, 0xa0, 0x5f, 0x81, 0x12, 0x48, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0,
	0x00, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x28, 0xd2, 0x60, 0xa2, 0x01, 0x00, 0x00,
}

func (m *DenomPairTakerFeeProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DenomPairTakerFeeProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DenomPairTakerFeeProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DenomPairTakerFee) > 0 {
		for iNdEx := len(m.DenomPairTakerFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DenomPairTakerFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DenomPairTakerFeeProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.DenomPairTakerFee) > 0 {
		for _, e := range m.DenomPairTakerFee {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DenomPairTakerFeeProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: DenomPairTakerFeeProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DenomPairTakerFeeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomPairTakerFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomPairTakerFee = append(m.DenomPairTakerFee, DenomPairTakerFee{})
			if err := m.DenomPairTakerFee[len(m.DenomPairTakerFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
