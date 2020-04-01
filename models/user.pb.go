// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

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

type User struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	OfficeID             string   `protobuf:"bytes,2,opt,name=OfficeID,proto3" json:"OfficeID,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *User) GetOfficeID() string {
	if m != nil {
		return m.OfficeID
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateUser struct {
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUser) Reset()         { *m = CreateUser{} }
func (m *CreateUser) String() string { return proto.CompactTextString(m) }
func (*CreateUser) ProtoMessage()    {}
func (*CreateUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *CreateUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUser.Unmarshal(m, b)
}
func (m *CreateUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUser.Marshal(b, m, deterministic)
}
func (m *CreateUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUser.Merge(m, src)
}
func (m *CreateUser) XXX_Size() int {
	return xxx_messageInfo_CreateUser.Size(m)
}
func (m *CreateUser) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUser.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUser proto.InternalMessageInfo

func (m *CreateUser) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UserCreated struct {
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserCreated) Reset()         { *m = UserCreated{} }
func (m *UserCreated) String() string { return proto.CompactTextString(m) }
func (*UserCreated) ProtoMessage()    {}
func (*UserCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCreated.Unmarshal(m, b)
}
func (m *UserCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCreated.Marshal(b, m, deterministic)
}
func (m *UserCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCreated.Merge(m, src)
}
func (m *UserCreated) XXX_Size() int {
	return xxx_messageInfo_UserCreated.Size(m)
}
func (m *UserCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCreated.DiscardUnknown(m)
}

var xxx_messageInfo_UserCreated proto.InternalMessageInfo

func (m *UserCreated) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UserUpdated struct {
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserUpdated) Reset()         { *m = UserUpdated{} }
func (m *UserUpdated) String() string { return proto.CompactTextString(m) }
func (*UserUpdated) ProtoMessage()    {}
func (*UserUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UserUpdated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserUpdated.Unmarshal(m, b)
}
func (m *UserUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserUpdated.Marshal(b, m, deterministic)
}
func (m *UserUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserUpdated.Merge(m, src)
}
func (m *UserUpdated) XXX_Size() int {
	return xxx_messageInfo_UserUpdated.Size(m)
}
func (m *UserUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_UserUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_UserUpdated proto.InternalMessageInfo

func (m *UserUpdated) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateUser struct {
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUser) Reset()         { *m = UpdateUser{} }
func (m *UpdateUser) String() string { return proto.CompactTextString(m) }
func (*UpdateUser) ProtoMessage()    {}
func (*UpdateUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *UpdateUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUser.Unmarshal(m, b)
}
func (m *UpdateUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUser.Marshal(b, m, deterministic)
}
func (m *UpdateUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUser.Merge(m, src)
}
func (m *UpdateUser) XXX_Size() int {
	return xxx_messageInfo_UpdateUser.Size(m)
}
func (m *UpdateUser) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUser.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUser proto.InternalMessageInfo

func (m *UpdateUser) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type DeleteUser struct {
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUser) Reset()         { *m = DeleteUser{} }
func (m *DeleteUser) String() string { return proto.CompactTextString(m) }
func (*DeleteUser) ProtoMessage()    {}
func (*DeleteUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *DeleteUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUser.Unmarshal(m, b)
}
func (m *DeleteUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUser.Marshal(b, m, deterministic)
}
func (m *DeleteUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUser.Merge(m, src)
}
func (m *DeleteUser) XXX_Size() int {
	return xxx_messageInfo_DeleteUser.Size(m)
}
func (m *DeleteUser) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUser.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUser proto.InternalMessageInfo

func (m *DeleteUser) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UserDeleted struct {
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDeleted) Reset()         { *m = UserDeleted{} }
func (m *UserDeleted) String() string { return proto.CompactTextString(m) }
func (*UserDeleted) ProtoMessage()    {}
func (*UserDeleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *UserDeleted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDeleted.Unmarshal(m, b)
}
func (m *UserDeleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDeleted.Marshal(b, m, deterministic)
}
func (m *UserDeleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDeleted.Merge(m, src)
}
func (m *UserDeleted) XXX_Size() int {
	return xxx_messageInfo_UserDeleted.Size(m)
}
func (m *UserDeleted) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDeleted.DiscardUnknown(m)
}

var xxx_messageInfo_UserDeleted proto.InternalMessageInfo

func (m *UserDeleted) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*CreateUser)(nil), "proto.CreateUser")
	proto.RegisterType((*UserCreated)(nil), "proto.UserCreated")
	proto.RegisterType((*UserUpdated)(nil), "proto.UserUpdated")
	proto.RegisterType((*UpdateUser)(nil), "proto.UpdateUser")
	proto.RegisterType((*DeleteUser)(nil), "proto.DeleteUser")
	proto.RegisterType((*UserDeleted)(nil), "proto.UserDeleted")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x6e, 0x5c, 0x2c, 0xa1, 0xc5,
	0xa9, 0x45, 0x42, 0x7c, 0x5c, 0x4c, 0x9e, 0x2e, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x4c,
	0x9e, 0x2e, 0x42, 0x52, 0x5c, 0x1c, 0xfe, 0x69, 0x69, 0x99, 0xc9, 0xa9, 0x9e, 0x2e, 0x12, 0x4c,
	0x60, 0x51, 0x38, 0x5f, 0x48, 0x88, 0x8b, 0xc5, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x19, 0x2c, 0x0e,
	0x66, 0x2b, 0xe9, 0x72, 0x71, 0x39, 0x17, 0xa5, 0x26, 0x96, 0xa4, 0x82, 0x4d, 0x93, 0x87, 0x98,
	0x0a, 0x36, 0x8f, 0xdb, 0x88, 0x1b, 0x62, 0xa5, 0x1e, 0x48, 0x28, 0x08, 0x2c, 0xa1, 0xa4, 0xc7,
	0xc5, 0x0d, 0xa2, 0x21, 0x5a, 0x52, 0x88, 0x56, 0x1f, 0x5a, 0x90, 0x42, 0x9c, 0x7a, 0x5d, 0x2e,
	0x2e, 0x88, 0x5a, 0xe2, 0x9c, 0xa3, 0xcb, 0xc5, 0xe5, 0x92, 0x9a, 0x93, 0x4a, 0xa2, 0xeb, 0x21,
	0x5a, 0x08, 0xbb, 0x26, 0x89, 0x0d, 0x2c, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x83, 0x13,
	0xa1, 0x00, 0x80, 0x01, 0x00, 0x00,
}