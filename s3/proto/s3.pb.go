// Code generated by protoc-gen-go. DO NOT EDIT.
// source: s3.proto

package s3

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetObjectRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetObjectRequest) Reset()         { *m = GetObjectRequest{} }
func (m *GetObjectRequest) String() string { return proto.CompactTextString(m) }
func (*GetObjectRequest) ProtoMessage()    {}
func (*GetObjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_s3_176a3dee47962900, []int{0}
}
func (m *GetObjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetObjectRequest.Unmarshal(m, b)
}
func (m *GetObjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetObjectRequest.Marshal(b, m, deterministic)
}
func (dst *GetObjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetObjectRequest.Merge(dst, src)
}
func (m *GetObjectRequest) XXX_Size() int {
	return xxx_messageInfo_GetObjectRequest.Size(m)
}
func (m *GetObjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetObjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetObjectRequest proto.InternalMessageInfo

func (m *GetObjectRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetObjectResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetObjectResponse) Reset()         { *m = GetObjectResponse{} }
func (m *GetObjectResponse) String() string { return proto.CompactTextString(m) }
func (*GetObjectResponse) ProtoMessage()    {}
func (*GetObjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_s3_176a3dee47962900, []int{1}
}
func (m *GetObjectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetObjectResponse.Unmarshal(m, b)
}
func (m *GetObjectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetObjectResponse.Marshal(b, m, deterministic)
}
func (dst *GetObjectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetObjectResponse.Merge(dst, src)
}
func (m *GetObjectResponse) XXX_Size() int {
	return xxx_messageInfo_GetObjectResponse.Size(m)
}
func (m *GetObjectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetObjectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetObjectResponse proto.InternalMessageInfo

func (m *GetObjectResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetObjectResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*GetObjectRequest)(nil), "GetObjectRequest")
	proto.RegisterType((*GetObjectResponse)(nil), "GetObjectResponse")
}

func init() { proto.RegisterFile("s3.proto", fileDescriptor_s3_176a3dee47962900) }

var fileDescriptor_s3_176a3dee47962900 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x28, 0x36, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe2, 0x12, 0x70, 0x4f, 0x2d, 0xf1, 0x4f, 0xca, 0x4a, 0x4d,
	0x2e, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x51, 0x32, 0xe7, 0x12, 0x44, 0x52, 0x53, 0x5c,
	0x90, 0x9f, 0x57, 0x9c, 0x8a, 0xae, 0x48, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82,
	0x09, 0x2c, 0x02, 0x66, 0x1b, 0x59, 0x71, 0x31, 0x05, 0x1b, 0x0b, 0x99, 0x70, 0x71, 0xc2, 0xb5,
	0x0b, 0x09, 0xea, 0xa1, 0x5b, 0x27, 0x25, 0xa4, 0x87, 0x61, 0xba, 0x12, 0x43, 0x12, 0x1b, 0xd8,
	0x7d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x50, 0x09, 0xa7, 0x5b, 0xab, 0x00, 0x00, 0x00,
}
