// Code generated by protoc-gen-go. DO NOT EDIT.
// source: office.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type OfficeService int32

const (
	OfficeService_OFFICES OfficeService = 0
)

var OfficeService_name = map[int32]string{
	0: "OFFICES",
}

var OfficeService_value = map[string]int32{
	"OFFICES": 0,
}

func (x OfficeService) String() string {
	return proto.EnumName(OfficeService_name, int32(x))
}

func (OfficeService) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_535fdf6b9dfa72fe, []int{0}
}

type OfficeEvents int32

const (
	OfficeEvents_CREATE_OFFICE  OfficeEvents = 0
	OfficeEvents_OFFICE_CREATED OfficeEvents = 1
	OfficeEvents_UPDATE_OFFICE  OfficeEvents = 2
	OfficeEvents_OFFICE_UPDATED OfficeEvents = 3
	OfficeEvents_DELETE_OFFICE  OfficeEvents = 4
	OfficeEvents_OFFICE_DELETED OfficeEvents = 5
)

var OfficeEvents_name = map[int32]string{
	0: "CREATE_OFFICE",
	1: "OFFICE_CREATED",
	2: "UPDATE_OFFICE",
	3: "OFFICE_UPDATED",
	4: "DELETE_OFFICE",
	5: "OFFICE_DELETED",
}

var OfficeEvents_value = map[string]int32{
	"CREATE_OFFICE":  0,
	"OFFICE_CREATED": 1,
	"UPDATE_OFFICE":  2,
	"OFFICE_UPDATED": 3,
	"DELETE_OFFICE":  4,
	"OFFICE_DELETED": 5,
}

func (x OfficeEvents) String() string {
	return proto.EnumName(OfficeEvents_name, int32(x))
}

func (OfficeEvents) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_535fdf6b9dfa72fe, []int{1}
}

type Office struct {
	ID                   string         `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Address              *Address       `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	AddressID            string         `protobuf:"bytes,4,opt,name=AddressID,proto3" json:"AddressID,omitempty"`
	ContactInfo          []*ContactInfo `protobuf:"bytes,5,rep,name=ContactInfo,proto3" json:"ContactInfo,omitempty"`
	UserIDs              []string       `protobuf:"bytes,6,rep,name=UserIDs,proto3" json:"UserIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Office) Reset()         { *m = Office{} }
func (m *Office) String() string { return proto.CompactTextString(m) }
func (*Office) ProtoMessage()    {}
func (*Office) Descriptor() ([]byte, []int) {
	return fileDescriptor_535fdf6b9dfa72fe, []int{0}
}

func (m *Office) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Office.Unmarshal(m, b)
}
func (m *Office) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Office.Marshal(b, m, deterministic)
}
func (m *Office) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Office.Merge(m, src)
}
func (m *Office) XXX_Size() int {
	return xxx_messageInfo_Office.Size(m)
}
func (m *Office) XXX_DiscardUnknown() {
	xxx_messageInfo_Office.DiscardUnknown(m)
}

var xxx_messageInfo_Office proto.InternalMessageInfo

func (m *Office) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Office) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Office) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Office) GetAddressID() string {
	if m != nil {
		return m.AddressID
	}
	return ""
}

func (m *Office) GetContactInfo() []*ContactInfo {
	if m != nil {
		return m.ContactInfo
	}
	return nil
}

func (m *Office) GetUserIDs() []string {
	if m != nil {
		return m.UserIDs
	}
	return nil
}

type Address struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	RoadName             string   `protobuf:"bytes,2,opt,name=RoadName,proto3" json:"RoadName,omitempty"`
	Number               int32    `protobuf:"varint,3,opt,name=Number,proto3" json:"Number,omitempty"`
	ZipCode              int32    `protobuf:"varint,4,opt,name=ZipCode,proto3" json:"ZipCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_535fdf6b9dfa72fe, []int{1}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Address) GetRoadName() string {
	if m != nil {
		return m.RoadName
	}
	return ""
}

func (m *Address) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Address) GetZipCode() int32 {
	if m != nil {
		return m.ZipCode
	}
	return 0
}

type ContactInfo struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	PhoneNumber          int32    `protobuf:"varint,3,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Mail                 string   `protobuf:"bytes,4,opt,name=Mail,proto3" json:"Mail,omitempty"`
	CompanyID            string   `protobuf:"bytes,5,opt,name=CompanyID,proto3" json:"CompanyID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactInfo) Reset()         { *m = ContactInfo{} }
func (m *ContactInfo) String() string { return proto.CompactTextString(m) }
func (*ContactInfo) ProtoMessage()    {}
func (*ContactInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_535fdf6b9dfa72fe, []int{2}
}

func (m *ContactInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactInfo.Unmarshal(m, b)
}
func (m *ContactInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactInfo.Marshal(b, m, deterministic)
}
func (m *ContactInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactInfo.Merge(m, src)
}
func (m *ContactInfo) XXX_Size() int {
	return xxx_messageInfo_ContactInfo.Size(m)
}
func (m *ContactInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ContactInfo proto.InternalMessageInfo

func (m *ContactInfo) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ContactInfo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ContactInfo) GetPhoneNumber() int32 {
	if m != nil {
		return m.PhoneNumber
	}
	return 0
}

func (m *ContactInfo) GetMail() string {
	if m != nil {
		return m.Mail
	}
	return ""
}

func (m *ContactInfo) GetCompanyID() string {
	if m != nil {
		return m.CompanyID
	}
	return ""
}

type CreateOffice struct {
	Office               *Office  `protobuf:"bytes,1,opt,name=Office,proto3" json:"Office,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOffice) Reset()         { *m = CreateOffice{} }
func (m *CreateOffice) String() string { return proto.CompactTextString(m) }
func (*CreateOffice) ProtoMessage()    {}
func (*CreateOffice) Descriptor() ([]byte, []int) {
	return fileDescriptor_535fdf6b9dfa72fe, []int{3}
}

func (m *CreateOffice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOffice.Unmarshal(m, b)
}
func (m *CreateOffice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOffice.Marshal(b, m, deterministic)
}
func (m *CreateOffice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOffice.Merge(m, src)
}
func (m *CreateOffice) XXX_Size() int {
	return xxx_messageInfo_CreateOffice.Size(m)
}
func (m *CreateOffice) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOffice.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOffice proto.InternalMessageInfo

func (m *CreateOffice) GetOffice() *Office {
	if m != nil {
		return m.Office
	}
	return nil
}

type OfficeCreated struct {
	Office               *Office  `protobuf:"bytes,1,opt,name=Office,proto3" json:"Office,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OfficeCreated) Reset()         { *m = OfficeCreated{} }
func (m *OfficeCreated) String() string { return proto.CompactTextString(m) }
func (*OfficeCreated) ProtoMessage()    {}
func (*OfficeCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_535fdf6b9dfa72fe, []int{4}
}

func (m *OfficeCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OfficeCreated.Unmarshal(m, b)
}
func (m *OfficeCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OfficeCreated.Marshal(b, m, deterministic)
}
func (m *OfficeCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OfficeCreated.Merge(m, src)
}
func (m *OfficeCreated) XXX_Size() int {
	return xxx_messageInfo_OfficeCreated.Size(m)
}
func (m *OfficeCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_OfficeCreated.DiscardUnknown(m)
}

var xxx_messageInfo_OfficeCreated proto.InternalMessageInfo

func (m *OfficeCreated) GetOffice() *Office {
	if m != nil {
		return m.Office
	}
	return nil
}

type OfficeUpdated struct {
	Office               *Office  `protobuf:"bytes,1,opt,name=Office,proto3" json:"Office,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OfficeUpdated) Reset()         { *m = OfficeUpdated{} }
func (m *OfficeUpdated) String() string { return proto.CompactTextString(m) }
func (*OfficeUpdated) ProtoMessage()    {}
func (*OfficeUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_535fdf6b9dfa72fe, []int{5}
}

func (m *OfficeUpdated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OfficeUpdated.Unmarshal(m, b)
}
func (m *OfficeUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OfficeUpdated.Marshal(b, m, deterministic)
}
func (m *OfficeUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OfficeUpdated.Merge(m, src)
}
func (m *OfficeUpdated) XXX_Size() int {
	return xxx_messageInfo_OfficeUpdated.Size(m)
}
func (m *OfficeUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_OfficeUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_OfficeUpdated proto.InternalMessageInfo

func (m *OfficeUpdated) GetOffice() *Office {
	if m != nil {
		return m.Office
	}
	return nil
}

type UpdateOffice struct {
	Office               *Office  `protobuf:"bytes,1,opt,name=Office,proto3" json:"Office,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateOffice) Reset()         { *m = UpdateOffice{} }
func (m *UpdateOffice) String() string { return proto.CompactTextString(m) }
func (*UpdateOffice) ProtoMessage()    {}
func (*UpdateOffice) Descriptor() ([]byte, []int) {
	return fileDescriptor_535fdf6b9dfa72fe, []int{6}
}

func (m *UpdateOffice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateOffice.Unmarshal(m, b)
}
func (m *UpdateOffice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateOffice.Marshal(b, m, deterministic)
}
func (m *UpdateOffice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateOffice.Merge(m, src)
}
func (m *UpdateOffice) XXX_Size() int {
	return xxx_messageInfo_UpdateOffice.Size(m)
}
func (m *UpdateOffice) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateOffice.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateOffice proto.InternalMessageInfo

func (m *UpdateOffice) GetOffice() *Office {
	if m != nil {
		return m.Office
	}
	return nil
}

type DeleteOffice struct {
	Office               *Office  `protobuf:"bytes,1,opt,name=Office,proto3" json:"Office,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteOffice) Reset()         { *m = DeleteOffice{} }
func (m *DeleteOffice) String() string { return proto.CompactTextString(m) }
func (*DeleteOffice) ProtoMessage()    {}
func (*DeleteOffice) Descriptor() ([]byte, []int) {
	return fileDescriptor_535fdf6b9dfa72fe, []int{7}
}

func (m *DeleteOffice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteOffice.Unmarshal(m, b)
}
func (m *DeleteOffice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteOffice.Marshal(b, m, deterministic)
}
func (m *DeleteOffice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteOffice.Merge(m, src)
}
func (m *DeleteOffice) XXX_Size() int {
	return xxx_messageInfo_DeleteOffice.Size(m)
}
func (m *DeleteOffice) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteOffice.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteOffice proto.InternalMessageInfo

func (m *DeleteOffice) GetOffice() *Office {
	if m != nil {
		return m.Office
	}
	return nil
}

type OfficeDeleted struct {
	Office               *Office  `protobuf:"bytes,1,opt,name=Office,proto3" json:"Office,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OfficeDeleted) Reset()         { *m = OfficeDeleted{} }
func (m *OfficeDeleted) String() string { return proto.CompactTextString(m) }
func (*OfficeDeleted) ProtoMessage()    {}
func (*OfficeDeleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_535fdf6b9dfa72fe, []int{8}
}

func (m *OfficeDeleted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OfficeDeleted.Unmarshal(m, b)
}
func (m *OfficeDeleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OfficeDeleted.Marshal(b, m, deterministic)
}
func (m *OfficeDeleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OfficeDeleted.Merge(m, src)
}
func (m *OfficeDeleted) XXX_Size() int {
	return xxx_messageInfo_OfficeDeleted.Size(m)
}
func (m *OfficeDeleted) XXX_DiscardUnknown() {
	xxx_messageInfo_OfficeDeleted.DiscardUnknown(m)
}

var xxx_messageInfo_OfficeDeleted proto.InternalMessageInfo

func (m *OfficeDeleted) GetOffice() *Office {
	if m != nil {
		return m.Office
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.OfficeService", OfficeService_name, OfficeService_value)
	proto.RegisterEnum("proto.OfficeEvents", OfficeEvents_name, OfficeEvents_value)
	proto.RegisterType((*Office)(nil), "proto.Office")
	proto.RegisterType((*Address)(nil), "proto.Address")
	proto.RegisterType((*ContactInfo)(nil), "proto.ContactInfo")
	proto.RegisterType((*CreateOffice)(nil), "proto.CreateOffice")
	proto.RegisterType((*OfficeCreated)(nil), "proto.OfficeCreated")
	proto.RegisterType((*OfficeUpdated)(nil), "proto.OfficeUpdated")
	proto.RegisterType((*UpdateOffice)(nil), "proto.UpdateOffice")
	proto.RegisterType((*DeleteOffice)(nil), "proto.DeleteOffice")
	proto.RegisterType((*OfficeDeleted)(nil), "proto.OfficeDeleted")
}

func init() { proto.RegisterFile("office.proto", fileDescriptor_535fdf6b9dfa72fe) }

var fileDescriptor_535fdf6b9dfa72fe = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x51, 0x4f, 0xaf, 0x9a, 0x40,
	0x10, 0x7f, 0x80, 0x60, 0x1d, 0xd0, 0xd0, 0x4d, 0xd3, 0x6c, 0x9a, 0x77, 0x20, 0x24, 0x4d, 0xc8,
	0x3b, 0xbc, 0x83, 0xad, 0xbd, 0xbf, 0xb0, 0x98, 0x90, 0xb4, 0x6a, 0x56, 0xbd, 0xf4, 0x62, 0x50,
	0xd6, 0x96, 0x44, 0x59, 0x02, 0xd4, 0xa4, 0xe7, 0x1e, 0xfa, 0xc5, 0xfa, 0xc1, 0x1a, 0x76, 0x17,
	0xc5, 0x78, 0xa9, 0x3d, 0x31, 0xf3, 0xfb, 0xb3, 0xf3, 0x9b, 0x01, 0x1c, 0xbe, 0xdf, 0x67, 0x3b,
	0xf6, 0x5c, 0x94, 0xbc, 0xe6, 0xc8, 0x14, 0x1f, 0xff, 0x8f, 0x06, 0xd6, 0x5c, 0xe0, 0x68, 0x04,
	0x7a, 0x4c, 0xb0, 0xe6, 0x69, 0xc1, 0x80, 0xea, 0x31, 0x41, 0x08, 0x7a, 0xb3, 0xe4, 0xc8, 0xb0,
	0x2e, 0x10, 0x51, 0xa3, 0x00, 0xfa, 0x2f, 0x69, 0x5a, 0xb2, 0xaa, 0xc2, 0x86, 0xa7, 0x05, 0xf6,
	0x78, 0x24, 0x9f, 0x7b, 0x56, 0x28, 0x6d, 0x69, 0xf4, 0x08, 0x03, 0x55, 0xc6, 0x04, 0xf7, 0xc4,
	0x13, 0x17, 0x00, 0x7d, 0x04, 0x3b, 0xe4, 0x79, 0x9d, 0xec, 0xea, 0x38, 0xdf, 0x73, 0x6c, 0x7a,
	0x46, 0x60, 0x8f, 0x91, 0x7a, 0xab, 0xc3, 0xd0, 0xae, 0x0c, 0x61, 0xe8, 0xaf, 0x2b, 0x56, 0xc6,
	0xa4, 0xc2, 0x96, 0x67, 0x04, 0x03, 0xda, 0xb6, 0xfe, 0xb7, 0x73, 0xae, 0x9b, 0x35, 0xde, 0xc1,
	0x2b, 0xca, 0x93, 0xb4, 0xb3, 0xca, 0xb9, 0x47, 0x6f, 0xc1, 0x9a, 0xfd, 0x38, 0x6e, 0x59, 0x29,
	0xb6, 0x31, 0xa9, 0xea, 0x9a, 0x41, 0x5f, 0xb3, 0x22, 0xe4, 0x29, 0x13, 0xd1, 0x4d, 0xda, 0xb6,
	0xfe, 0x6f, 0xed, 0x2a, 0xf9, 0xcd, 0xb4, 0x37, 0x60, 0xae, 0xb2, 0xfa, 0xd0, 0x8e, 0x92, 0x0d,
	0xf2, 0xc0, 0x5e, 0x7c, 0xe7, 0x39, 0xbb, 0x1a, 0xd6, 0x85, 0x9a, 0x63, 0x7f, 0x49, 0xb2, 0x83,
	0xba, 0x94, 0xa8, 0x9b, 0x13, 0x86, 0xfc, 0x58, 0x24, 0xf9, 0xcf, 0x98, 0x60, 0x53, 0x9e, 0xf0,
	0x0c, 0xf8, 0x13, 0x70, 0xc2, 0x92, 0x25, 0x35, 0x53, 0xbf, 0xef, 0x7d, 0xfb, 0x23, 0x45, 0x1a,
	0x7b, 0x3c, 0x54, 0xd7, 0x94, 0x20, 0x55, 0xa4, 0xff, 0x09, 0x86, 0xb2, 0x92, 0xe6, 0xf4, 0x6e,
	0xdf, 0xba, 0x48, 0xef, 0xf1, 0x4d, 0xc0, 0x91, 0x8e, 0xfb, 0x62, 0x4e, 0xc0, 0x21, 0xec, 0xc0,
	0xfe, 0x7b, 0x3b, 0x69, 0xfe, 0xd7, 0x94, 0x4f, 0x8f, 0xad, 0x6f, 0xc9, 0xca, 0x53, 0x33, 0xcf,
	0x86, 0xfe, 0x7c, 0x3a, 0x8d, 0xc3, 0x68, 0xe9, 0x3e, 0x3c, 0xfd, 0xd2, 0xc0, 0x91, 0x74, 0x74,
	0x62, 0x79, 0x5d, 0xa1, 0xd7, 0x30, 0x0c, 0x69, 0xf4, 0xb2, 0x8a, 0x36, 0x52, 0xe4, 0x3e, 0x20,
	0x04, 0x23, 0x59, 0x6f, 0x24, 0x43, 0x5c, 0xad, 0x91, 0xad, 0x17, 0xa4, 0x23, 0xd3, 0x3b, 0x32,
	0xc9, 0x10, 0xd7, 0x68, 0x64, 0x24, 0xfa, 0x1c, 0x5d, 0x64, 0xbd, 0x8e, 0x4c, 0x32, 0xc4, 0x35,
	0xb7, 0x96, 0x48, 0xfe, 0xe1, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x48, 0x81, 0x62, 0xff, 0xc8,
	0x03, 0x00, 0x00,
}
