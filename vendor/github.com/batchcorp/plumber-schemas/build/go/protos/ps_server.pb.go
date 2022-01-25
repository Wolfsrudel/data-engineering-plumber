// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_server.proto

package protos

import (
	fmt "fmt"
	common "github.com/batchcorp/plumber-schemas/build/go/protos/common"
	opts "github.com/batchcorp/plumber-schemas/build/go/protos/opts"
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

type GetServerOptionsRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetServerOptionsRequest) Reset()         { *m = GetServerOptionsRequest{} }
func (m *GetServerOptionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetServerOptionsRequest) ProtoMessage()    {}
func (*GetServerOptionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_99cda3bcaff1a15a, []int{0}
}

func (m *GetServerOptionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServerOptionsRequest.Unmarshal(m, b)
}
func (m *GetServerOptionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServerOptionsRequest.Marshal(b, m, deterministic)
}
func (m *GetServerOptionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServerOptionsRequest.Merge(m, src)
}
func (m *GetServerOptionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetServerOptionsRequest.Size(m)
}
func (m *GetServerOptionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServerOptionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetServerOptionsRequest proto.InternalMessageInfo

func (m *GetServerOptionsRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

type GetServerOptionsResponse struct {
	ServerOptions        *opts.ServerOptions `protobuf:"bytes,1,opt,name=server_options,json=serverOptions,proto3" json:"server_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetServerOptionsResponse) Reset()         { *m = GetServerOptionsResponse{} }
func (m *GetServerOptionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetServerOptionsResponse) ProtoMessage()    {}
func (*GetServerOptionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_99cda3bcaff1a15a, []int{1}
}

func (m *GetServerOptionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServerOptionsResponse.Unmarshal(m, b)
}
func (m *GetServerOptionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServerOptionsResponse.Marshal(b, m, deterministic)
}
func (m *GetServerOptionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServerOptionsResponse.Merge(m, src)
}
func (m *GetServerOptionsResponse) XXX_Size() int {
	return xxx_messageInfo_GetServerOptionsResponse.Size(m)
}
func (m *GetServerOptionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServerOptionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetServerOptionsResponse proto.InternalMessageInfo

func (m *GetServerOptionsResponse) GetServerOptions() *opts.ServerOptions {
	if m != nil {
		return m.ServerOptions
	}
	return nil
}

type SetServerOptionsRequest struct {
	// Every gRPC request must have a valid auth config
	Auth *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	// Token returned from VC Service authorization process. This is used to authenticate against
	// the VC Service gRPC API
	VcserviceToken       string   `protobuf:"bytes,1,opt,name=vcservice_token,json=vcserviceToken,proto3" json:"vcservice_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetServerOptionsRequest) Reset()         { *m = SetServerOptionsRequest{} }
func (m *SetServerOptionsRequest) String() string { return proto.CompactTextString(m) }
func (*SetServerOptionsRequest) ProtoMessage()    {}
func (*SetServerOptionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_99cda3bcaff1a15a, []int{2}
}

func (m *SetServerOptionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetServerOptionsRequest.Unmarshal(m, b)
}
func (m *SetServerOptionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetServerOptionsRequest.Marshal(b, m, deterministic)
}
func (m *SetServerOptionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetServerOptionsRequest.Merge(m, src)
}
func (m *SetServerOptionsRequest) XXX_Size() int {
	return xxx_messageInfo_SetServerOptionsRequest.Size(m)
}
func (m *SetServerOptionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetServerOptionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetServerOptionsRequest proto.InternalMessageInfo

func (m *SetServerOptionsRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *SetServerOptionsRequest) GetVcserviceToken() string {
	if m != nil {
		return m.VcserviceToken
	}
	return ""
}

type SetServerOptionsResponse struct {
	// Returns updated copy of server options
	ServerOptions        *opts.ServerOptions `protobuf:"bytes,1,opt,name=server_options,json=serverOptions,proto3" json:"server_options,omitempty"`
	Status               *common.Status      `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SetServerOptionsResponse) Reset()         { *m = SetServerOptionsResponse{} }
func (m *SetServerOptionsResponse) String() string { return proto.CompactTextString(m) }
func (*SetServerOptionsResponse) ProtoMessage()    {}
func (*SetServerOptionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_99cda3bcaff1a15a, []int{3}
}

func (m *SetServerOptionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetServerOptionsResponse.Unmarshal(m, b)
}
func (m *SetServerOptionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetServerOptionsResponse.Marshal(b, m, deterministic)
}
func (m *SetServerOptionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetServerOptionsResponse.Merge(m, src)
}
func (m *SetServerOptionsResponse) XXX_Size() int {
	return xxx_messageInfo_SetServerOptionsResponse.Size(m)
}
func (m *SetServerOptionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetServerOptionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetServerOptionsResponse proto.InternalMessageInfo

func (m *SetServerOptionsResponse) GetServerOptions() *opts.ServerOptions {
	if m != nil {
		return m.ServerOptions
	}
	return nil
}

func (m *SetServerOptionsResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*GetServerOptionsRequest)(nil), "protos.GetServerOptionsRequest")
	proto.RegisterType((*GetServerOptionsResponse)(nil), "protos.GetServerOptionsResponse")
	proto.RegisterType((*SetServerOptionsRequest)(nil), "protos.SetServerOptionsRequest")
	proto.RegisterType((*SetServerOptionsResponse)(nil), "protos.SetServerOptionsResponse")
}

func init() { proto.RegisterFile("ps_server.proto", fileDescriptor_99cda3bcaff1a15a) }

var fileDescriptor_99cda3bcaff1a15a = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x09, 0x48, 0xc5, 0x15, 0x5b, 0x88, 0x48, 0x63, 0x44, 0x90, 0x5c, 0xec, 0xc5, 0x2c,
	0xa8, 0x78, 0xaf, 0x1e, 0xbc, 0x29, 0x24, 0x9e, 0x04, 0x09, 0xc9, 0x3a, 0x34, 0xc1, 0x24, 0xb3,
	0x66, 0x66, 0xfb, 0x16, 0xe2, 0x6b, 0xfa, 0x18, 0xb2, 0xbb, 0x55, 0xac, 0x3d, 0x8a, 0xa7, 0x49,
	0xe6, 0x9f, 0xf9, 0xfe, 0x7f, 0xd9, 0x15, 0x13, 0x4d, 0x05, 0xc1, 0xb0, 0x84, 0x21, 0xd5, 0x03,
	0x32, 0x86, 0x23, 0x57, 0x28, 0x3e, 0x52, 0xd8, 0x75, 0xd8, 0x4b, 0x4d, 0x85, 0xff, 0x2a, 0x4a,
	0xc3, 0xb5, 0x1f, 0x8a, 0x8f, 0x37, 0x44, 0xe2, 0x92, 0x0d, 0xad, 0xe4, 0x43, 0xd4, 0x4c, 0x56,
	0xb4, 0x75, 0x0d, 0x9f, 0xdc, 0x88, 0xe9, 0x2d, 0x70, 0xee, 0x5a, 0xf7, 0x9a, 0x1b, 0xec, 0x29,
	0x83, 0x57, 0x03, 0xc4, 0xe1, 0x4c, 0x6c, 0x59, 0x8b, 0xe8, 0xfd, 0xee, 0x24, 0x98, 0xed, 0x9e,
	0xef, 0xfb, 0x0d, 0x4a, 0xbd, 0x43, 0x3a, 0x37, 0x5c, 0x67, 0x6e, 0x22, 0x79, 0x12, 0xd1, 0x26,
	0x84, 0x34, 0xf6, 0x04, 0xe1, 0x5c, 0x8c, 0xbd, 0xa1, 0x35, 0xb7, 0x4a, 0x14, 0x38, 0x5c, 0xfc,
	0x85, 0xb3, 0x99, 0xd2, 0xf5, 0xdd, 0x3d, 0xfa, 0xf9, 0x9b, 0xb4, 0x62, 0x9a, 0xff, 0x35, 0x63,
	0x78, 0x2a, 0x26, 0x4b, 0x65, 0xb9, 0x8d, 0x82, 0x82, 0xf1, 0x05, 0x7a, 0x17, 0x64, 0x27, 0x1b,
	0x7f, 0xb7, 0x1f, 0x6c, 0x37, 0x79, 0x0b, 0x44, 0x94, 0xff, 0xdf, 0x69, 0xc2, 0x54, 0x8c, 0xfc,
	0xe5, 0x44, 0x1f, 0xdb, 0x6e, 0xf7, 0xe0, 0x57, 0xe8, 0xdc, 0xa9, 0xd9, 0x6a, 0xea, 0xfa, 0xea,
	0xf1, 0x72, 0xd1, 0x70, 0x6d, 0x2a, 0xab, 0xcb, 0xaa, 0x64, 0x55, 0x2b, 0x1c, 0xb4, 0xd4, 0xad,
	0xe9, 0x2a, 0x18, 0xce, 0x48, 0xd5, 0xd0, 0x95, 0x24, 0x2b, 0xd3, 0xb4, 0xcf, 0x72, 0x81, 0xd2,
	0xd3, 0x2a, 0xff, 0x70, 0x2e, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x32, 0x47, 0xa6, 0x5d, 0x52,
	0x02, 0x00, 0x00,
}
