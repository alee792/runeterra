// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/runeterra.proto

package runeterra

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
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

type Core struct {
	Keywords             []*Keyword `protobuf:"bytes,1,rep,name=keywords,proto3" json:"keywords,omitempty"`
	Regions              []*Region  `protobuf:"bytes,2,rep,name=regions,proto3" json:"regions,omitempty"`
	SpellSpeeds          []*ID      `protobuf:"bytes,3,rep,name=spell_speeds,json=spellSpeeds,proto3" json:"spellSpeeds"`
	Rarities             []*ID      `protobuf:"bytes,4,rep,name=rarities,proto3" json:"rarities"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Core) Reset()         { *m = Core{} }
func (m *Core) String() string { return proto.CompactTextString(m) }
func (*Core) ProtoMessage()    {}
func (*Core) Descriptor() ([]byte, []int) {
	return fileDescriptor_88b91daf5f3c2b1a, []int{0}
}
func (m *Core) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Core.Unmarshal(m, b)
}
func (m *Core) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Core.Marshal(b, m, deterministic)
}
func (m *Core) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Core.Merge(m, src)
}
func (m *Core) XXX_Size() int {
	return xxx_messageInfo_Core.Size(m)
}
func (m *Core) XXX_DiscardUnknown() {
	xxx_messageInfo_Core.DiscardUnknown(m)
}

var xxx_messageInfo_Core proto.InternalMessageInfo

func (m *Core) GetKeywords() []*Keyword {
	if m != nil {
		return m.Keywords
	}
	return nil
}

func (m *Core) GetRegions() []*Region {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *Core) GetSpellSpeeds() []*ID {
	if m != nil {
		return m.SpellSpeeds
	}
	return nil
}

func (m *Core) GetRarities() []*ID {
	if m != nil {
		return m.Rarities
	}
	return nil
}

type ID struct {
	NameRef              string   `protobuf:"bytes,1,opt,name=name_ref,json=nameRef,proto3" json:"nameRef"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_88b91daf5f3c2b1a, []int{1}
}
func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (m *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(m, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetNameRef() string {
	if m != nil {
		return m.NameRef
	}
	return ""
}

func (m *ID) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Keyword struct {
	*ID                  `protobuf:"bytes,1,opt,name=id,proto3,embedded=id" json:""`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Keyword) Reset()         { *m = Keyword{} }
func (m *Keyword) String() string { return proto.CompactTextString(m) }
func (*Keyword) ProtoMessage()    {}
func (*Keyword) Descriptor() ([]byte, []int) {
	return fileDescriptor_88b91daf5f3c2b1a, []int{2}
}
func (m *Keyword) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Keyword.Unmarshal(m, b)
}
func (m *Keyword) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Keyword.Marshal(b, m, deterministic)
}
func (m *Keyword) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keyword.Merge(m, src)
}
func (m *Keyword) XXX_Size() int {
	return xxx_messageInfo_Keyword.Size(m)
}
func (m *Keyword) XXX_DiscardUnknown() {
	xxx_messageInfo_Keyword.DiscardUnknown(m)
}

var xxx_messageInfo_Keyword proto.InternalMessageInfo

func (m *Keyword) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Region struct {
	*ID                  `protobuf:"bytes,1,opt,name=id,proto3,embedded=id" json:""`
	Abbreviation         string   `protobuf:"bytes,2,opt,name=abbreviation,proto3" json:"abbreviation,omitempty"`
	IconAbsolutePath     string   `protobuf:"bytes,3,opt,name=icon_absolute_path,json=iconAbsolutePath,proto3" json:"icon_absolute_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Region) Reset()         { *m = Region{} }
func (m *Region) String() string { return proto.CompactTextString(m) }
func (*Region) ProtoMessage()    {}
func (*Region) Descriptor() ([]byte, []int) {
	return fileDescriptor_88b91daf5f3c2b1a, []int{3}
}
func (m *Region) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Region.Unmarshal(m, b)
}
func (m *Region) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Region.Marshal(b, m, deterministic)
}
func (m *Region) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Region.Merge(m, src)
}
func (m *Region) XXX_Size() int {
	return xxx_messageInfo_Region.Size(m)
}
func (m *Region) XXX_DiscardUnknown() {
	xxx_messageInfo_Region.DiscardUnknown(m)
}

var xxx_messageInfo_Region proto.InternalMessageInfo

func (m *Region) GetAbbreviation() string {
	if m != nil {
		return m.Abbreviation
	}
	return ""
}

func (m *Region) GetIconAbsolutePath() string {
	if m != nil {
		return m.IconAbsolutePath
	}
	return ""
}

type SpellSpeed struct {
	*ID                  `protobuf:"bytes,1,opt,name=id,proto3,embedded=id" json:""`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpellSpeed) Reset()         { *m = SpellSpeed{} }
func (m *SpellSpeed) String() string { return proto.CompactTextString(m) }
func (*SpellSpeed) ProtoMessage()    {}
func (*SpellSpeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_88b91daf5f3c2b1a, []int{4}
}
func (m *SpellSpeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpellSpeed.Unmarshal(m, b)
}
func (m *SpellSpeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpellSpeed.Marshal(b, m, deterministic)
}
func (m *SpellSpeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpellSpeed.Merge(m, src)
}
func (m *SpellSpeed) XXX_Size() int {
	return xxx_messageInfo_SpellSpeed.Size(m)
}
func (m *SpellSpeed) XXX_DiscardUnknown() {
	xxx_messageInfo_SpellSpeed.DiscardUnknown(m)
}

var xxx_messageInfo_SpellSpeed proto.InternalMessageInfo

type Rarity struct {
	*ID                  `protobuf:"bytes,1,opt,name=id,proto3,embedded=id" json:""`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rarity) Reset()         { *m = Rarity{} }
func (m *Rarity) String() string { return proto.CompactTextString(m) }
func (*Rarity) ProtoMessage()    {}
func (*Rarity) Descriptor() ([]byte, []int) {
	return fileDescriptor_88b91daf5f3c2b1a, []int{5}
}
func (m *Rarity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rarity.Unmarshal(m, b)
}
func (m *Rarity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rarity.Marshal(b, m, deterministic)
}
func (m *Rarity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rarity.Merge(m, src)
}
func (m *Rarity) XXX_Size() int {
	return xxx_messageInfo_Rarity.Size(m)
}
func (m *Rarity) XXX_DiscardUnknown() {
	xxx_messageInfo_Rarity.DiscardUnknown(m)
}

var xxx_messageInfo_Rarity proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Core)(nil), "runeterra.Core")
	proto.RegisterType((*ID)(nil), "runeterra.ID")
	proto.RegisterType((*Keyword)(nil), "runeterra.Keyword")
	proto.RegisterType((*Region)(nil), "runeterra.Region")
	proto.RegisterType((*SpellSpeed)(nil), "runeterra.SpellSpeed")
	proto.RegisterType((*Rarity)(nil), "runeterra.Rarity")
}

func init() { proto.RegisterFile("proto/runeterra.proto", fileDescriptor_88b91daf5f3c2b1a) }

var fileDescriptor_88b91daf5f3c2b1a = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x6f, 0x0b, 0x81, 0x72, 0xca, 0xcd, 0xbd, 0x77, 0x92, 0x9b, 0x34, 0x6e, 0x4a, 0xba,
	0x50, 0x12, 0xb5, 0x44, 0x8d, 0x71, 0x2b, 0xc8, 0x86, 0xb8, 0x31, 0x83, 0xfb, 0x66, 0x4a, 0x87,
	0x32, 0x11, 0x3a, 0xcd, 0xcc, 0x54, 0xc3, 0xda, 0x17, 0xf4, 0x09, 0xea, 0xbe, 0x4f, 0x61, 0x3a,
	0x85, 0x5a, 0xa3, 0x0b, 0x76, 0xff, 0x39, 0xe7, 0xfb, 0xff, 0xd3, 0x9e, 0x0c, 0xfc, 0x4f, 0x05,
	0x57, 0x7c, 0x24, 0xb2, 0x84, 0x2a, 0x2a, 0x04, 0xf1, 0x75, 0x8d, 0x7a, 0x75, 0xe3, 0xe8, 0x3c,
	0x66, 0x6a, 0x95, 0x85, 0xfe, 0x82, 0x6f, 0x46, 0x31, 0x8f, 0xf9, 0x48, 0x13, 0x61, 0xb6, 0xd4,
	0x55, 0x65, 0x2f, 0x55, 0xe5, 0xf4, 0xde, 0x0d, 0x68, 0xdf, 0x71, 0x41, 0x91, 0x0f, 0xd6, 0x13,
	0xdd, 0xbe, 0x70, 0x11, 0x49, 0xc7, 0x18, 0xb4, 0x86, 0xf6, 0x25, 0xf2, 0x3f, 0xd7, 0xdc, 0x57,
	0x23, 0x5c, 0x33, 0xe8, 0x14, 0xba, 0x82, 0xc6, 0x8c, 0x27, 0xd2, 0x31, 0x35, 0xfe, 0xaf, 0x81,
	0x63, 0x3d, 0xc1, 0x7b, 0x02, 0x8d, 0xa1, 0x2f, 0x53, 0xba, 0x5e, 0x07, 0x32, 0xa5, 0x34, 0x92,
	0x4e, 0x4b, 0x3b, 0x7e, 0x37, 0x1c, 0xb3, 0xe9, 0xe4, 0x4f, 0x91, 0xbb, 0xb6, 0xc6, 0xe6, 0x9a,
	0xc2, 0xcd, 0x02, 0xdd, 0x80, 0x25, 0x88, 0x60, 0x8a, 0x51, 0xe9, 0xb4, 0x7f, 0xb2, 0xf7, 0x8b,
	0xdc, 0xad, 0x11, 0x5c, 0x2b, 0xef, 0x16, 0xcc, 0xd9, 0x14, 0x1d, 0x83, 0x95, 0x90, 0x0d, 0x0d,
	0x04, 0x5d, 0x3a, 0xc6, 0xc0, 0x18, 0xf6, 0x26, 0x76, 0x91, 0xbb, 0xdd, 0xb2, 0x87, 0xe9, 0x12,
	0xef, 0x05, 0x42, 0xd0, 0x2e, 0xa5, 0x63, 0x96, 0x0c, 0xd6, 0xda, 0x7b, 0x84, 0xee, 0xee, 0xff,
	0xd1, 0x09, 0x98, 0x2c, 0xd2, 0x01, 0xdf, 0xf6, 0x5b, 0x6f, 0xb9, 0x6b, 0x14, 0xb9, 0xfb, 0x0b,
	0x9b, 0x2c, 0x42, 0x03, 0xb0, 0x23, 0x2a, 0x17, 0x82, 0xa5, 0x8a, 0xf1, 0x64, 0x17, 0xd7, 0x6c,
	0x79, 0xaf, 0x06, 0x74, 0xaa, 0x3b, 0x1d, 0x9e, 0xea, 0x41, 0x9f, 0x84, 0xa1, 0xa0, 0xcf, 0x8c,
	0x34, 0x62, 0xbf, 0xf4, 0xd0, 0x19, 0x20, 0xb6, 0xe0, 0x49, 0x40, 0x42, 0xc9, 0xd7, 0x99, 0xa2,
	0x41, 0x4a, 0xd4, 0xca, 0x69, 0x69, 0xf2, 0x6f, 0x39, 0x19, 0xef, 0x06, 0x0f, 0x44, 0xad, 0xbc,
	0x6b, 0x80, 0x79, 0x7d, 0xe5, 0x83, 0x3f, 0xc4, 0xbb, 0x80, 0x0e, 0x2e, 0x0f, 0xbc, 0x3d, 0xd8,
	0x12, 0x76, 0xf4, 0x83, 0xbb, 0xfa, 0x08, 0x00, 0x00, 0xff, 0xff, 0x21, 0x48, 0x25, 0x25, 0xc3,
	0x02, 0x00, 0x00,
}
