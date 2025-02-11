// Code generated by protoc-gen-go. DO NOT EDIT.
// source: unit.proto

package unitpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateUnitReq struct {
	Unit                 *Unit    `protobuf:"bytes,1,opt,name=unit,proto3" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUnitReq) Reset()         { *m = CreateUnitReq{} }
func (m *CreateUnitReq) String() string { return proto.CompactTextString(m) }
func (*CreateUnitReq) ProtoMessage()    {}
func (*CreateUnitReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8d688cf2cb325c8, []int{0}
}

func (m *CreateUnitReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUnitReq.Unmarshal(m, b)
}
func (m *CreateUnitReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUnitReq.Marshal(b, m, deterministic)
}
func (m *CreateUnitReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUnitReq.Merge(m, src)
}
func (m *CreateUnitReq) XXX_Size() int {
	return xxx_messageInfo_CreateUnitReq.Size(m)
}
func (m *CreateUnitReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUnitReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUnitReq proto.InternalMessageInfo

func (m *CreateUnitReq) GetUnit() *Unit {
	if m != nil {
		return m.Unit
	}
	return nil
}

type CreateUnitRes struct {
	Unit                 *Unit    `protobuf:"bytes,1,opt,name=unit,proto3" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUnitRes) Reset()         { *m = CreateUnitRes{} }
func (m *CreateUnitRes) String() string { return proto.CompactTextString(m) }
func (*CreateUnitRes) ProtoMessage()    {}
func (*CreateUnitRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8d688cf2cb325c8, []int{1}
}

func (m *CreateUnitRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUnitRes.Unmarshal(m, b)
}
func (m *CreateUnitRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUnitRes.Marshal(b, m, deterministic)
}
func (m *CreateUnitRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUnitRes.Merge(m, src)
}
func (m *CreateUnitRes) XXX_Size() int {
	return xxx_messageInfo_CreateUnitRes.Size(m)
}
func (m *CreateUnitRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUnitRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUnitRes proto.InternalMessageInfo

func (m *CreateUnitRes) GetUnit() *Unit {
	if m != nil {
		return m.Unit
	}
	return nil
}

type ReadUnitReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadUnitReq) Reset()         { *m = ReadUnitReq{} }
func (m *ReadUnitReq) String() string { return proto.CompactTextString(m) }
func (*ReadUnitReq) ProtoMessage()    {}
func (*ReadUnitReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8d688cf2cb325c8, []int{2}
}

func (m *ReadUnitReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadUnitReq.Unmarshal(m, b)
}
func (m *ReadUnitReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadUnitReq.Marshal(b, m, deterministic)
}
func (m *ReadUnitReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadUnitReq.Merge(m, src)
}
func (m *ReadUnitReq) XXX_Size() int {
	return xxx_messageInfo_ReadUnitReq.Size(m)
}
func (m *ReadUnitReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadUnitReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReadUnitReq proto.InternalMessageInfo

func (m *ReadUnitReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadUnitRes struct {
	Unit                 *Unit    `protobuf:"bytes,1,opt,name=unit,proto3" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadUnitRes) Reset()         { *m = ReadUnitRes{} }
func (m *ReadUnitRes) String() string { return proto.CompactTextString(m) }
func (*ReadUnitRes) ProtoMessage()    {}
func (*ReadUnitRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8d688cf2cb325c8, []int{3}
}

func (m *ReadUnitRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadUnitRes.Unmarshal(m, b)
}
func (m *ReadUnitRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadUnitRes.Marshal(b, m, deterministic)
}
func (m *ReadUnitRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadUnitRes.Merge(m, src)
}
func (m *ReadUnitRes) XXX_Size() int {
	return xxx_messageInfo_ReadUnitRes.Size(m)
}
func (m *ReadUnitRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadUnitRes.DiscardUnknown(m)
}

var xxx_messageInfo_ReadUnitRes proto.InternalMessageInfo

func (m *ReadUnitRes) GetUnit() *Unit {
	if m != nil {
		return m.Unit
	}
	return nil
}

type UpdateUnitReq struct {
	Unit                 *Unit    `protobuf:"bytes,1,opt,name=unit,proto3" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUnitReq) Reset()         { *m = UpdateUnitReq{} }
func (m *UpdateUnitReq) String() string { return proto.CompactTextString(m) }
func (*UpdateUnitReq) ProtoMessage()    {}
func (*UpdateUnitReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8d688cf2cb325c8, []int{4}
}

func (m *UpdateUnitReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUnitReq.Unmarshal(m, b)
}
func (m *UpdateUnitReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUnitReq.Marshal(b, m, deterministic)
}
func (m *UpdateUnitReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUnitReq.Merge(m, src)
}
func (m *UpdateUnitReq) XXX_Size() int {
	return xxx_messageInfo_UpdateUnitReq.Size(m)
}
func (m *UpdateUnitReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUnitReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUnitReq proto.InternalMessageInfo

func (m *UpdateUnitReq) GetUnit() *Unit {
	if m != nil {
		return m.Unit
	}
	return nil
}

type UpdateUnitRes struct {
	Unit                 *Unit    `protobuf:"bytes,1,opt,name=unit,proto3" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUnitRes) Reset()         { *m = UpdateUnitRes{} }
func (m *UpdateUnitRes) String() string { return proto.CompactTextString(m) }
func (*UpdateUnitRes) ProtoMessage()    {}
func (*UpdateUnitRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8d688cf2cb325c8, []int{5}
}

func (m *UpdateUnitRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUnitRes.Unmarshal(m, b)
}
func (m *UpdateUnitRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUnitRes.Marshal(b, m, deterministic)
}
func (m *UpdateUnitRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUnitRes.Merge(m, src)
}
func (m *UpdateUnitRes) XXX_Size() int {
	return xxx_messageInfo_UpdateUnitRes.Size(m)
}
func (m *UpdateUnitRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUnitRes.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUnitRes proto.InternalMessageInfo

func (m *UpdateUnitRes) GetUnit() *Unit {
	if m != nil {
		return m.Unit
	}
	return nil
}

type DeleteUnitReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUnitReq) Reset()         { *m = DeleteUnitReq{} }
func (m *DeleteUnitReq) String() string { return proto.CompactTextString(m) }
func (*DeleteUnitReq) ProtoMessage()    {}
func (*DeleteUnitReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8d688cf2cb325c8, []int{6}
}

func (m *DeleteUnitReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUnitReq.Unmarshal(m, b)
}
func (m *DeleteUnitReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUnitReq.Marshal(b, m, deterministic)
}
func (m *DeleteUnitReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUnitReq.Merge(m, src)
}
func (m *DeleteUnitReq) XXX_Size() int {
	return xxx_messageInfo_DeleteUnitReq.Size(m)
}
func (m *DeleteUnitReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUnitReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUnitReq proto.InternalMessageInfo

func (m *DeleteUnitReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteUnitRes struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUnitRes) Reset()         { *m = DeleteUnitRes{} }
func (m *DeleteUnitRes) String() string { return proto.CompactTextString(m) }
func (*DeleteUnitRes) ProtoMessage()    {}
func (*DeleteUnitRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8d688cf2cb325c8, []int{7}
}

func (m *DeleteUnitRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUnitRes.Unmarshal(m, b)
}
func (m *DeleteUnitRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUnitRes.Marshal(b, m, deterministic)
}
func (m *DeleteUnitRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUnitRes.Merge(m, src)
}
func (m *DeleteUnitRes) XXX_Size() int {
	return xxx_messageInfo_DeleteUnitRes.Size(m)
}
func (m *DeleteUnitRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUnitRes.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUnitRes proto.InternalMessageInfo

func (m *DeleteUnitRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ListUnitReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUnitReq) Reset()         { *m = ListUnitReq{} }
func (m *ListUnitReq) String() string { return proto.CompactTextString(m) }
func (*ListUnitReq) ProtoMessage()    {}
func (*ListUnitReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8d688cf2cb325c8, []int{8}
}

func (m *ListUnitReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUnitReq.Unmarshal(m, b)
}
func (m *ListUnitReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUnitReq.Marshal(b, m, deterministic)
}
func (m *ListUnitReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUnitReq.Merge(m, src)
}
func (m *ListUnitReq) XXX_Size() int {
	return xxx_messageInfo_ListUnitReq.Size(m)
}
func (m *ListUnitReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUnitReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListUnitReq proto.InternalMessageInfo

type ListUnitRes struct {
	Unit                 *Unit    `protobuf:"bytes,1,opt,name=unit,proto3" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUnitRes) Reset()         { *m = ListUnitRes{} }
func (m *ListUnitRes) String() string { return proto.CompactTextString(m) }
func (*ListUnitRes) ProtoMessage()    {}
func (*ListUnitRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8d688cf2cb325c8, []int{9}
}

func (m *ListUnitRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUnitRes.Unmarshal(m, b)
}
func (m *ListUnitRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUnitRes.Marshal(b, m, deterministic)
}
func (m *ListUnitRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUnitRes.Merge(m, src)
}
func (m *ListUnitRes) XXX_Size() int {
	return xxx_messageInfo_ListUnitRes.Size(m)
}
func (m *ListUnitRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUnitRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListUnitRes proto.InternalMessageInfo

func (m *ListUnitRes) GetUnit() *Unit {
	if m != nil {
		return m.Unit
	}
	return nil
}

type Unit struct {
	UId                  string   `protobuf:"bytes,1,opt,name=u_id,json=uId,proto3" json:"u_id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Unit) Reset()         { *m = Unit{} }
func (m *Unit) String() string { return proto.CompactTextString(m) }
func (*Unit) ProtoMessage()    {}
func (*Unit) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8d688cf2cb325c8, []int{10}
}

func (m *Unit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Unit.Unmarshal(m, b)
}
func (m *Unit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Unit.Marshal(b, m, deterministic)
}
func (m *Unit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unit.Merge(m, src)
}
func (m *Unit) XXX_Size() int {
	return xxx_messageInfo_Unit.Size(m)
}
func (m *Unit) XXX_DiscardUnknown() {
	xxx_messageInfo_Unit.DiscardUnknown(m)
}

var xxx_messageInfo_Unit proto.InternalMessageInfo

func (m *Unit) GetUId() string {
	if m != nil {
		return m.UId
	}
	return ""
}

func (m *Unit) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateUnitReq)(nil), "service.CreateUnitReq")
	proto.RegisterType((*CreateUnitRes)(nil), "service.CreateUnitRes")
	proto.RegisterType((*ReadUnitReq)(nil), "service.ReadUnitReq")
	proto.RegisterType((*ReadUnitRes)(nil), "service.ReadUnitRes")
	proto.RegisterType((*UpdateUnitReq)(nil), "service.UpdateUnitReq")
	proto.RegisterType((*UpdateUnitRes)(nil), "service.UpdateUnitRes")
	proto.RegisterType((*DeleteUnitReq)(nil), "service.DeleteUnitReq")
	proto.RegisterType((*DeleteUnitRes)(nil), "service.DeleteUnitRes")
	proto.RegisterType((*ListUnitReq)(nil), "service.ListUnitReq")
	proto.RegisterType((*ListUnitRes)(nil), "service.ListUnitRes")
	proto.RegisterType((*Unit)(nil), "service.Unit")
}

func init() { proto.RegisterFile("unit.proto", fileDescriptor_e8d688cf2cb325c8) }

var fileDescriptor_e8d688cf2cb325c8 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0x49, 0xda, 0x7f, 0x9b, 0x4e, 0x48, 0xa1, 0xfb, 0x2f, 0x52, 0x82, 0xa2, 0xee, 0x49,
	0x7b, 0x68, 0x4a, 0xbd, 0x79, 0x54, 0x2f, 0x45, 0x41, 0x8c, 0xf4, 0xe2, 0x45, 0xd2, 0x66, 0x29,
	0x0b, 0x25, 0x89, 0x9d, 0x8d, 0x17, 0xf1, 0xe2, 0x2b, 0x78, 0xf3, 0xb5, 0x7c, 0x05, 0x1f, 0x44,
	0x76, 0xd3, 0x24, 0x9b, 0x92, 0x43, 0xbc, 0x75, 0xbe, 0x99, 0xf9, 0xe6, 0xb7, 0x7c, 0x0d, 0x40,
	0x1a, 0x71, 0x31, 0x49, 0xb6, 0xb1, 0x88, 0x49, 0x17, 0xd9, 0xf6, 0x95, 0xaf, 0x98, 0x7b, 0xb8,
	0x8e, 0xe3, 0xf5, 0x86, 0x79, 0x41, 0xc2, 0xbd, 0x20, 0x8a, 0x62, 0x11, 0x08, 0x1e, 0x47, 0x98,
	0x8d, 0xd1, 0x19, 0x38, 0xd7, 0x5b, 0x16, 0x08, 0xb6, 0x88, 0xb8, 0xf0, 0xd9, 0x0b, 0x39, 0x85,
	0xb6, 0x74, 0x19, 0x19, 0x27, 0xc6, 0x99, 0x3d, 0x73, 0x26, 0x3b, 0x9b, 0x89, 0xea, 0xab, 0xd6,
	0xfe, 0x0e, 0x36, 0xd9, 0x39, 0x02, 0xdb, 0x67, 0x41, 0x98, 0x5f, 0xe9, 0x83, 0xc9, 0x43, 0x35,
	0xdf, 0xf3, 0x4d, 0x1e, 0xd2, 0xa9, 0xde, 0xc6, 0x86, 0x10, 0x8b, 0x24, 0xfc, 0x33, 0xb8, 0xbe,
	0xd3, 0xe8, 0xce, 0x31, 0x38, 0x37, 0x6c, 0xc3, 0xca, 0x3b, 0xfb, 0xe8, 0xe7, 0xd5, 0x01, 0x24,
	0x23, 0xe8, 0x62, 0xba, 0x5a, 0x31, 0x44, 0x35, 0x65, 0xf9, 0x79, 0x49, 0x1d, 0xb0, 0xef, 0x38,
	0x8a, 0x9d, 0x93, 0x7c, 0x74, 0x59, 0x36, 0x82, 0xf1, 0xa0, 0x2d, 0x2b, 0x32, 0x80, 0x76, 0xfa,
	0x5c, 0x50, 0xb4, 0xd2, 0x79, 0x48, 0x86, 0xf0, 0x4f, 0x70, 0xb1, 0x61, 0x23, 0x53, 0x69, 0x59,
	0x31, 0xfb, 0x6a, 0x81, 0x2d, 0x37, 0x1e, 0x33, 0x2f, 0x72, 0x0f, 0x50, 0x46, 0x47, 0x0e, 0x8a,
	0x1b, 0x95, 0xff, 0x80, 0x5b, 0xaf, 0x23, 0x1d, 0x7c, 0x7c, 0xff, 0x7c, 0x9a, 0x36, 0xed, 0x78,
	0x92, 0x06, 0x2f, 0x8d, 0x31, 0xb9, 0x05, 0x2b, 0x0f, 0x8e, 0x0c, 0x8b, 0x35, 0x2d, 0x6a, 0xb7,
	0x4e, 0x45, 0xfa, 0x5f, 0x59, 0x39, 0xc4, 0xce, 0xac, 0xbc, 0x37, 0x1e, 0xbe, 0x4b, 0xba, 0x32,
	0x1f, 0x8d, 0xae, 0x12, 0xb4, 0x5b, 0xaf, 0x17, 0x74, 0xae, 0x46, 0xf7, 0x00, 0x50, 0x66, 0xa3,
	0x19, 0x56, 0x12, 0x75, 0xeb, 0xf5, 0x82, 0x71, 0x5c, 0x61, 0x9c, 0x43, 0x2f, 0x0f, 0x0d, 0xb5,
	0x17, 0x6b, 0xb9, 0xba, 0x75, 0x2a, 0xd2, 0xbe, 0x72, 0xb3, 0xc8, 0x0e, 0x6f, 0x6a, 0x5c, 0x59,
	0x4f, 0x1d, 0xf9, 0x33, 0x59, 0x2e, 0x3b, 0xea, 0x63, 0xbc, 0xf8, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0xaa, 0x95, 0x06, 0xd7, 0xc1, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UnitServiceClient is the client API for UnitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UnitServiceClient interface {
	CreateUnit(ctx context.Context, in *CreateUnitReq, opts ...grpc.CallOption) (*CreateUnitRes, error)
	ReadUnit(ctx context.Context, in *ReadUnitReq, opts ...grpc.CallOption) (*ReadUnitRes, error)
	UpdateUnit(ctx context.Context, in *UpdateUnitReq, opts ...grpc.CallOption) (*UpdateUnitRes, error)
	DeleteUnit(ctx context.Context, in *DeleteUnitReq, opts ...grpc.CallOption) (*DeleteUnitRes, error)
	ListUnits(ctx context.Context, in *ListUnitReq, opts ...grpc.CallOption) (UnitService_ListUnitsClient, error)
}

type unitServiceClient struct {
	cc *grpc.ClientConn
}

func NewUnitServiceClient(cc *grpc.ClientConn) UnitServiceClient {
	return &unitServiceClient{cc}
}

func (c *unitServiceClient) CreateUnit(ctx context.Context, in *CreateUnitReq, opts ...grpc.CallOption) (*CreateUnitRes, error) {
	out := new(CreateUnitRes)
	err := c.cc.Invoke(ctx, "/service.UnitService/CreateUnit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unitServiceClient) ReadUnit(ctx context.Context, in *ReadUnitReq, opts ...grpc.CallOption) (*ReadUnitRes, error) {
	out := new(ReadUnitRes)
	err := c.cc.Invoke(ctx, "/service.UnitService/ReadUnit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unitServiceClient) UpdateUnit(ctx context.Context, in *UpdateUnitReq, opts ...grpc.CallOption) (*UpdateUnitRes, error) {
	out := new(UpdateUnitRes)
	err := c.cc.Invoke(ctx, "/service.UnitService/UpdateUnit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unitServiceClient) DeleteUnit(ctx context.Context, in *DeleteUnitReq, opts ...grpc.CallOption) (*DeleteUnitRes, error) {
	out := new(DeleteUnitRes)
	err := c.cc.Invoke(ctx, "/service.UnitService/DeleteUnit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unitServiceClient) ListUnits(ctx context.Context, in *ListUnitReq, opts ...grpc.CallOption) (UnitService_ListUnitsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UnitService_serviceDesc.Streams[0], "/service.UnitService/ListUnits", opts...)
	if err != nil {
		return nil, err
	}
	x := &unitServiceListUnitsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UnitService_ListUnitsClient interface {
	Recv() (*ListUnitRes, error)
	grpc.ClientStream
}

type unitServiceListUnitsClient struct {
	grpc.ClientStream
}

func (x *unitServiceListUnitsClient) Recv() (*ListUnitRes, error) {
	m := new(ListUnitRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UnitServiceServer is the server API for UnitService service.
type UnitServiceServer interface {
	CreateUnit(context.Context, *CreateUnitReq) (*CreateUnitRes, error)
	ReadUnit(context.Context, *ReadUnitReq) (*ReadUnitRes, error)
	UpdateUnit(context.Context, *UpdateUnitReq) (*UpdateUnitRes, error)
	DeleteUnit(context.Context, *DeleteUnitReq) (*DeleteUnitRes, error)
	ListUnits(*ListUnitReq, UnitService_ListUnitsServer) error
}

// UnimplementedUnitServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUnitServiceServer struct {
}

func (*UnimplementedUnitServiceServer) CreateUnit(ctx context.Context, req *CreateUnitReq) (*CreateUnitRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUnit not implemented")
}
func (*UnimplementedUnitServiceServer) ReadUnit(ctx context.Context, req *ReadUnitReq) (*ReadUnitRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadUnit not implemented")
}
func (*UnimplementedUnitServiceServer) UpdateUnit(ctx context.Context, req *UpdateUnitReq) (*UpdateUnitRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUnit not implemented")
}
func (*UnimplementedUnitServiceServer) DeleteUnit(ctx context.Context, req *DeleteUnitReq) (*DeleteUnitRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUnit not implemented")
}
func (*UnimplementedUnitServiceServer) ListUnits(req *ListUnitReq, srv UnitService_ListUnitsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListUnits not implemented")
}

func RegisterUnitServiceServer(s *grpc.Server, srv UnitServiceServer) {
	s.RegisterService(&_UnitService_serviceDesc, srv)
}

func _UnitService_CreateUnit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUnitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnitServiceServer).CreateUnit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UnitService/CreateUnit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnitServiceServer).CreateUnit(ctx, req.(*CreateUnitReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UnitService_ReadUnit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadUnitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnitServiceServer).ReadUnit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UnitService/ReadUnit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnitServiceServer).ReadUnit(ctx, req.(*ReadUnitReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UnitService_UpdateUnit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUnitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnitServiceServer).UpdateUnit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UnitService/UpdateUnit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnitServiceServer).UpdateUnit(ctx, req.(*UpdateUnitReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UnitService_DeleteUnit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUnitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnitServiceServer).DeleteUnit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UnitService/DeleteUnit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnitServiceServer).DeleteUnit(ctx, req.(*DeleteUnitReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UnitService_ListUnits_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListUnitReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UnitServiceServer).ListUnits(m, &unitServiceListUnitsServer{stream})
}

type UnitService_ListUnitsServer interface {
	Send(*ListUnitRes) error
	grpc.ServerStream
}

type unitServiceListUnitsServer struct {
	grpc.ServerStream
}

func (x *unitServiceListUnitsServer) Send(m *ListUnitRes) error {
	return x.ServerStream.SendMsg(m)
}

var _UnitService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.UnitService",
	HandlerType: (*UnitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUnit",
			Handler:    _UnitService_CreateUnit_Handler,
		},
		{
			MethodName: "ReadUnit",
			Handler:    _UnitService_ReadUnit_Handler,
		},
		{
			MethodName: "UpdateUnit",
			Handler:    _UnitService_UpdateUnit_Handler,
		},
		{
			MethodName: "DeleteUnit",
			Handler:    _UnitService_DeleteUnit_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListUnits",
			Handler:       _UnitService_ListUnits_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "unit.proto",
}
