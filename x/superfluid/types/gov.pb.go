// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/superfluid/gov.proto

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

// SetSuperfluidAssetsProposal is a gov Content type to update the superfluid assets
type SetSuperfluidAssetsProposal struct {
	Title       string            `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string            `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Assets      []SuperfluidAsset `protobuf:"bytes,3,rep,name=assets,proto3" json:"assets"`
}

func (m *SetSuperfluidAssetsProposal) Reset()      { *m = SetSuperfluidAssetsProposal{} }
func (*SetSuperfluidAssetsProposal) ProtoMessage() {}
func (*SetSuperfluidAssetsProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c729e4c39568a58, []int{0}
}
func (m *SetSuperfluidAssetsProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetSuperfluidAssetsProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetSuperfluidAssetsProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetSuperfluidAssetsProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetSuperfluidAssetsProposal.Merge(m, src)
}
func (m *SetSuperfluidAssetsProposal) XXX_Size() int {
	return m.Size()
}
func (m *SetSuperfluidAssetsProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SetSuperfluidAssetsProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SetSuperfluidAssetsProposal proto.InternalMessageInfo

// AddSuperfluidAssetsProposal is a gov Content type to add the superfluid assets
type AddSuperfluidAssetsProposal struct {
	Title            string            `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description      string            `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	SuperfluidAssets []SuperfluidAsset `protobuf:"bytes,3,rep,name=superfluid_assets,json=superfluidAssets,proto3" json:"superfluid_assets"`
}

func (m *AddSuperfluidAssetsProposal) Reset()      { *m = AddSuperfluidAssetsProposal{} }
func (*AddSuperfluidAssetsProposal) ProtoMessage() {}
func (*AddSuperfluidAssetsProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c729e4c39568a58, []int{1}
}
func (m *AddSuperfluidAssetsProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddSuperfluidAssetsProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddSuperfluidAssetsProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddSuperfluidAssetsProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSuperfluidAssetsProposal.Merge(m, src)
}
func (m *AddSuperfluidAssetsProposal) XXX_Size() int {
	return m.Size()
}
func (m *AddSuperfluidAssetsProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSuperfluidAssetsProposal.DiscardUnknown(m)
}

var xxx_messageInfo_AddSuperfluidAssetsProposal proto.InternalMessageInfo

// RemoveSuperfluidAssetsProposal is a gov Content type to remove the superfluid assets by denom
type RemoveSuperfluidAssetsProposal struct {
	Title                 string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description           string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	SuperfluidAssetDenoms []string `protobuf:"bytes,3,rep,name=superfluid_asset_denoms,json=superfluidAssetDenoms,proto3" json:"superfluid_asset_denoms,omitempty"`
}

func (m *RemoveSuperfluidAssetsProposal) Reset()      { *m = RemoveSuperfluidAssetsProposal{} }
func (*RemoveSuperfluidAssetsProposal) ProtoMessage() {}
func (*RemoveSuperfluidAssetsProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c729e4c39568a58, []int{2}
}
func (m *RemoveSuperfluidAssetsProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RemoveSuperfluidAssetsProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RemoveSuperfluidAssetsProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RemoveSuperfluidAssetsProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveSuperfluidAssetsProposal.Merge(m, src)
}
func (m *RemoveSuperfluidAssetsProposal) XXX_Size() int {
	return m.Size()
}
func (m *RemoveSuperfluidAssetsProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveSuperfluidAssetsProposal.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveSuperfluidAssetsProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SetSuperfluidAssetsProposal)(nil), "osmosis.superfluid.v1beta1.SetSuperfluidAssetsProposal")
	proto.RegisterType((*AddSuperfluidAssetsProposal)(nil), "osmosis.superfluid.v1beta1.AddSuperfluidAssetsProposal")
	proto.RegisterType((*RemoveSuperfluidAssetsProposal)(nil), "osmosis.superfluid.v1beta1.RemoveSuperfluidAssetsProposal")
}

func init() { proto.RegisterFile("osmosis/superfluid/gov.proto", fileDescriptor_7c729e4c39568a58) }

var fileDescriptor_7c729e4c39568a58 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc9, 0x2f, 0xce, 0xcd,
	0x2f, 0xce, 0x2c, 0xd6, 0x2f, 0x2e, 0x2d, 0x48, 0x2d, 0x4a, 0xcb, 0x29, 0xcd, 0x4c, 0xd1, 0x4f,
	0xcf, 0x2f, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x82, 0xca, 0xea, 0x21, 0x64, 0xf5,
	0xca, 0x0c, 0x93, 0x52, 0x4b, 0x12, 0x0d, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xca, 0xf4,
	0x41, 0x2c, 0x88, 0x0e, 0x29, 0x65, 0x2c, 0xe6, 0x21, 0x69, 0x06, 0x2b, 0x52, 0x5a, 0xc6, 0xc8,
	0x25, 0x1d, 0x9c, 0x5a, 0x12, 0x0c, 0x17, 0x77, 0x2c, 0x2e, 0x4e, 0x2d, 0x29, 0x0e, 0x28, 0xca,
	0x2f, 0xc8, 0x2f, 0x4e, 0xcc, 0x11, 0x12, 0xe1, 0x62, 0x2d, 0xc9, 0x2c, 0xc9, 0x49, 0x95, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0x14, 0xb8, 0xb8, 0x53, 0x52, 0x8b, 0x93, 0x8b,
	0x32, 0x0b, 0x4a, 0x32, 0xf3, 0xf3, 0x24, 0x98, 0xc0, 0x72, 0xc8, 0x42, 0x42, 0x8e, 0x5c, 0x6c,
	0x89, 0x60, 0x93, 0x24, 0x98, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x94, 0xf5, 0xb0, 0xb8, 0x1f, 0xcd,
	0x56, 0x27, 0x96, 0x13, 0xf7, 0xe4, 0x19, 0x82, 0xa0, 0x1a, 0xad, 0x78, 0x3a, 0x16, 0xc8, 0x33,
	0xcc, 0x58, 0x20, 0xcf, 0xf0, 0x62, 0x81, 0x3c, 0xa3, 0xd2, 0x6e, 0x46, 0x2e, 0x69, 0xc7, 0x94,
	0x14, 0xaa, 0x3b, 0x34, 0x8c, 0x4b, 0x10, 0xe1, 0xa2, 0x78, 0x72, 0xdd, 0x2c, 0x50, 0x8c, 0xe6,
	0x2e, 0x34, 0xd7, 0xcf, 0x63, 0xe4, 0x92, 0x0b, 0x4a, 0xcd, 0xcd, 0x2f, 0x4b, 0xa5, 0xba, 0x07,
	0xcc, 0xb8, 0xc4, 0xd1, 0x3d, 0x10, 0x9f, 0x92, 0x9a, 0x97, 0x9f, 0x0b, 0xf1, 0x06, 0x67, 0x90,
	0x28, 0x9a, 0xdb, 0x5c, 0xc0, 0x92, 0xa8, 0x0e, 0x74, 0xf2, 0x39, 0xf1, 0x48, 0x8e, 0xf1, 0xc2,
	0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1,
	0xc6, 0x63, 0x39, 0x86, 0x28, 0xa3, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c,
	0x7d, 0x68, 0x78, 0xe8, 0xe6, 0x24, 0x26, 0x15, 0xc3, 0x38, 0xfa, 0x15, 0xc8, 0x09, 0xac, 0xa4,
	0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x9c, 0xb8, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x01,
	0xe5, 0x15, 0xee, 0xd3, 0x02, 0x00, 0x00,
}

func (this *SetSuperfluidAssetsProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SetSuperfluidAssetsProposal)
	if !ok {
		that2, ok := that.(SetSuperfluidAssetsProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.Assets) != len(that1.Assets) {
		return false
	}
	for i := range this.Assets {
		if !this.Assets[i].Equal(&that1.Assets[i]) {
			return false
		}
	}
	return true
}
func (this *AddSuperfluidAssetsProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AddSuperfluidAssetsProposal)
	if !ok {
		that2, ok := that.(AddSuperfluidAssetsProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.SuperfluidAssets) != len(that1.SuperfluidAssets) {
		return false
	}
	for i := range this.SuperfluidAssets {
		if !this.SuperfluidAssets[i].Equal(&that1.SuperfluidAssets[i]) {
			return false
		}
	}
	return true
}
func (this *RemoveSuperfluidAssetsProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RemoveSuperfluidAssetsProposal)
	if !ok {
		that2, ok := that.(RemoveSuperfluidAssetsProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.SuperfluidAssetDenoms) != len(that1.SuperfluidAssetDenoms) {
		return false
	}
	for i := range this.SuperfluidAssetDenoms {
		if this.SuperfluidAssetDenoms[i] != that1.SuperfluidAssetDenoms[i] {
			return false
		}
	}
	return true
}
func (m *SetSuperfluidAssetsProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetSuperfluidAssetsProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetSuperfluidAssetsProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Assets) > 0 {
		for iNdEx := len(m.Assets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Assets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *AddSuperfluidAssetsProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddSuperfluidAssetsProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddSuperfluidAssetsProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SuperfluidAssets) > 0 {
		for iNdEx := len(m.SuperfluidAssets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SuperfluidAssets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *RemoveSuperfluidAssetsProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoveSuperfluidAssetsProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RemoveSuperfluidAssetsProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SuperfluidAssetDenoms) > 0 {
		for iNdEx := len(m.SuperfluidAssetDenoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SuperfluidAssetDenoms[iNdEx])
			copy(dAtA[i:], m.SuperfluidAssetDenoms[iNdEx])
			i = encodeVarintGov(dAtA, i, uint64(len(m.SuperfluidAssetDenoms[iNdEx])))
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
func (m *SetSuperfluidAssetsProposal) Size() (n int) {
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
	if len(m.Assets) > 0 {
		for _, e := range m.Assets {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	return n
}

func (m *AddSuperfluidAssetsProposal) Size() (n int) {
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
	if len(m.SuperfluidAssets) > 0 {
		for _, e := range m.SuperfluidAssets {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	return n
}

func (m *RemoveSuperfluidAssetsProposal) Size() (n int) {
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
	if len(m.SuperfluidAssetDenoms) > 0 {
		for _, s := range m.SuperfluidAssetDenoms {
			l = len(s)
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
func (m *SetSuperfluidAssetsProposal) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: SetSuperfluidAssetsProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetSuperfluidAssetsProposal: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Assets", wireType)
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
			m.Assets = append(m.Assets, SuperfluidAsset{})
			if err := m.Assets[len(m.Assets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGov
			}
			if (iNdEx + skippy) < 0 {
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
func (m *AddSuperfluidAssetsProposal) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: AddSuperfluidAssetsProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddSuperfluidAssetsProposal: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field SuperfluidAssets", wireType)
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
			m.SuperfluidAssets = append(m.SuperfluidAssets, SuperfluidAsset{})
			if err := m.SuperfluidAssets[len(m.SuperfluidAssets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGov
			}
			if (iNdEx + skippy) < 0 {
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
func (m *RemoveSuperfluidAssetsProposal) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: RemoveSuperfluidAssetsProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoveSuperfluidAssetsProposal: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field SuperfluidAssetDenoms", wireType)
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
			m.SuperfluidAssetDenoms = append(m.SuperfluidAssetDenoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGov
			}
			if (iNdEx + skippy) < 0 {
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
