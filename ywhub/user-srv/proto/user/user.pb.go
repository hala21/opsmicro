// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package ywhub_micro_v1_basic_srv_user

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
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	CreateTime           uint64   `protobuf:"varint,4,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime           uint64   `protobuf:"varint,5,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
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

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *User) GetUpdateTime() uint64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

type Request struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	UserPwd              string   `protobuf:"bytes,3,opt,name=userPwd,proto3" json:"userPwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Request) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Request) GetUserPwd() string {
	if m != nil {
		return m.UserPwd
	}
	return ""
}

type Response struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	User                 *User    `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type StreamingRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRequest.Unmarshal(m, b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
}
func (m *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(m, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingRequest.Size(m)
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{5}
}

func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (m *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(m, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{6}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{7}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "ywhub.micro.v1.basic.srv.user.user")
	proto.RegisterType((*Error)(nil), "ywhub.micro.v1.basic.srv.user.Error")
	proto.RegisterType((*Request)(nil), "ywhub.micro.v1.basic.srv.user.Request")
	proto.RegisterType((*Response)(nil), "ywhub.micro.v1.basic.srv.user.Response")
	proto.RegisterType((*StreamingRequest)(nil), "ywhub.micro.v1.basic.srv.user.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "ywhub.micro.v1.basic.srv.user.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "ywhub.micro.v1.basic.srv.user.Ping")
	proto.RegisterType((*Pong)(nil), "ywhub.micro.v1.basic.srv.user.Pong")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4f, 0xe3, 0x30,
	0x10, 0xc5, 0x37, 0x6d, 0xd2, 0x3f, 0x53, 0x69, 0x77, 0xb1, 0x00, 0x45, 0x95, 0xa8, 0xaa, 0x80,
	0x20, 0x5c, 0x8c, 0x68, 0x0f, 0x48, 0x1c, 0x11, 0x1c, 0xb8, 0xa0, 0x62, 0x40, 0x9c, 0xd3, 0x64,
	0x54, 0x2c, 0x9a, 0x38, 0xd8, 0x49, 0xab, 0x7e, 0x09, 0xee, 0x7c, 0x5b, 0x64, 0xc7, 0x81, 0x0a,
	0x01, 0xbd, 0x44, 0xfe, 0xf9, 0xbd, 0x71, 0xe6, 0x8d, 0x0d, 0x3b, 0xb9, 0x14, 0x85, 0x38, 0x29,
	0x15, 0x4a, 0xf3, 0xa1, 0x86, 0xc9, 0xde, 0x6a, 0xf9, 0x54, 0x4e, 0x69, 0xca, 0x63, 0x29, 0xe8,
	0xe2, 0x94, 0x4e, 0x23, 0xc5, 0x63, 0xaa, 0xe4, 0x82, 0x6a, 0x53, 0xf0, 0xea, 0x80, 0xab, 0x17,
	0xe4, 0x2f, 0x34, 0x78, 0xe2, 0x3b, 0x43, 0x27, 0x6c, 0xb2, 0x06, 0x4f, 0x48, 0x1f, 0x3a, 0x7a,
	0x3f, 0x8b, 0x52, 0xf4, 0x1b, 0x43, 0x27, 0xec, 0xb2, 0x0f, 0xd6, 0x5a, 0x1e, 0x29, 0xb5, 0x14,
	0x32, 0xf1, 0x9b, 0x95, 0x56, 0x33, 0x19, 0x00, 0xc4, 0x12, 0xa3, 0x02, 0xef, 0x79, 0x8a, 0xbe,
	0x3b, 0x74, 0x42, 0x97, 0xad, 0xed, 0x68, 0xbd, 0xcc, 0x93, 0x5a, 0xf7, 0x2a, 0xfd, 0x73, 0x27,
	0x18, 0x83, 0x77, 0x25, 0xa5, 0x90, 0x84, 0x80, 0x1b, 0x8b, 0x04, 0x4d, 0x4b, 0x1e, 0x33, 0x6b,
	0xb2, 0x0b, 0xad, 0x04, 0x8b, 0x88, 0xcf, 0x6d, 0x4b, 0x96, 0x82, 0x47, 0x68, 0x33, 0x7c, 0x29,
	0x51, 0x15, 0xda, 0xa2, 0xfb, 0xbc, 0xbe, 0x34, 0x85, 0x5d, 0x66, 0xa9, 0xce, 0x73, 0xf3, 0x25,
	0x8f, 0x66, 0xe2, 0x43, 0x5b, 0xaf, 0x27, 0xcb, 0x3a, 0x4e, 0x8d, 0xc1, 0x9b, 0x03, 0x1d, 0x86,
	0x2a, 0x17, 0x99, 0x32, 0x36, 0x55, 0xc6, 0x31, 0x2a, 0x65, 0xce, 0xee, 0xb0, 0x1a, 0xc9, 0x39,
	0x78, 0xa8, 0x9b, 0x36, 0x27, 0xf7, 0x46, 0x07, 0xf4, 0xd7, 0xa1, 0x53, 0x13, 0x90, 0x55, 0x25,
	0xe4, 0xac, 0xba, 0x00, 0xf3, 0xe7, 0xde, 0x68, 0x7f, 0x43, 0xa9, 0xfe, 0x30, 0x53, 0x10, 0x84,
	0xf0, 0xff, 0xae, 0x90, 0x18, 0xa5, 0x3c, 0x9b, 0xd5, 0xe9, 0xb7, 0xc1, 0x8b, 0x45, 0x99, 0x15,
	0xf6, 0x22, 0x2b, 0x08, 0x8e, 0x61, 0x6b, 0xcd, 0x69, 0xd3, 0x7c, 0x6f, 0x1d, 0x80, 0x3b, 0xe1,
	0xd9, 0x4c, 0x8f, 0x51, 0x15, 0x52, 0x3c, 0xa3, 0x95, 0x2d, 0x19, 0x5d, 0xfc, 0xac, 0x8f, 0xe6,
	0xe0, 0x3e, 0xe8, 0xe7, 0x94, 0xc0, 0xbf, 0xdb, 0x12, 0xe5, 0x4a, 0xc3, 0xc5, 0xca, 0x4c, 0xf9,
	0x70, 0x43, 0x34, 0x9b, 0xa1, 0x7f, 0xb4, 0xd1, 0x57, 0x25, 0x08, 0xfe, 0x4c, 0x5b, 0xe6, 0x8d,
	0x8f, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x71, 0xd4, 0x79, 0xfc, 0x02, 0x00, 0x00,
}
