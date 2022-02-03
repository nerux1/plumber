// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_dynamic.proto

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

type GetAllDynamicRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetAllDynamicRequest) Reset()         { *m = GetAllDynamicRequest{} }
func (m *GetAllDynamicRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllDynamicRequest) ProtoMessage()    {}
func (*GetAllDynamicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d091d947a15a54f7, []int{0}
}

func (m *GetAllDynamicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllDynamicRequest.Unmarshal(m, b)
}
func (m *GetAllDynamicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllDynamicRequest.Marshal(b, m, deterministic)
}
func (m *GetAllDynamicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllDynamicRequest.Merge(m, src)
}
func (m *GetAllDynamicRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllDynamicRequest.Size(m)
}
func (m *GetAllDynamicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllDynamicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllDynamicRequest proto.InternalMessageInfo

func (m *GetAllDynamicRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

type GetAllDynamicResponse struct {
	Status *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	// Will be set as empty []Dynamic if no dynamic destinations are configured
	Opts                 []*opts.DynamicOptions `protobuf:"bytes,1,rep,name=opts,proto3" json:"opts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GetAllDynamicResponse) Reset()         { *m = GetAllDynamicResponse{} }
func (m *GetAllDynamicResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllDynamicResponse) ProtoMessage()    {}
func (*GetAllDynamicResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d091d947a15a54f7, []int{1}
}

func (m *GetAllDynamicResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllDynamicResponse.Unmarshal(m, b)
}
func (m *GetAllDynamicResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllDynamicResponse.Marshal(b, m, deterministic)
}
func (m *GetAllDynamicResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllDynamicResponse.Merge(m, src)
}
func (m *GetAllDynamicResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllDynamicResponse.Size(m)
}
func (m *GetAllDynamicResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllDynamicResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllDynamicResponse proto.InternalMessageInfo

func (m *GetAllDynamicResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetAllDynamicResponse) GetOpts() []*opts.DynamicOptions {
	if m != nil {
		return m.Opts
	}
	return nil
}

type GetDynamicRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	DynamicId            string       `protobuf:"bytes,1,opt,name=dynamic_id,json=dynamicId,proto3" json:"dynamic_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetDynamicRequest) Reset()         { *m = GetDynamicRequest{} }
func (m *GetDynamicRequest) String() string { return proto.CompactTextString(m) }
func (*GetDynamicRequest) ProtoMessage()    {}
func (*GetDynamicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d091d947a15a54f7, []int{2}
}

func (m *GetDynamicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDynamicRequest.Unmarshal(m, b)
}
func (m *GetDynamicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDynamicRequest.Marshal(b, m, deterministic)
}
func (m *GetDynamicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDynamicRequest.Merge(m, src)
}
func (m *GetDynamicRequest) XXX_Size() int {
	return xxx_messageInfo_GetDynamicRequest.Size(m)
}
func (m *GetDynamicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDynamicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDynamicRequest proto.InternalMessageInfo

func (m *GetDynamicRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *GetDynamicRequest) GetDynamicId() string {
	if m != nil {
		return m.DynamicId
	}
	return ""
}

type GetDynamicResponse struct {
	Status *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	// Set only if status is OK
	Opts                 *opts.DynamicOptions `protobuf:"bytes,1,opt,name=opts,proto3" json:"opts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetDynamicResponse) Reset()         { *m = GetDynamicResponse{} }
func (m *GetDynamicResponse) String() string { return proto.CompactTextString(m) }
func (*GetDynamicResponse) ProtoMessage()    {}
func (*GetDynamicResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d091d947a15a54f7, []int{3}
}

func (m *GetDynamicResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDynamicResponse.Unmarshal(m, b)
}
func (m *GetDynamicResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDynamicResponse.Marshal(b, m, deterministic)
}
func (m *GetDynamicResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDynamicResponse.Merge(m, src)
}
func (m *GetDynamicResponse) XXX_Size() int {
	return xxx_messageInfo_GetDynamicResponse.Size(m)
}
func (m *GetDynamicResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDynamicResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDynamicResponse proto.InternalMessageInfo

func (m *GetDynamicResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetDynamicResponse) GetOpts() *opts.DynamicOptions {
	if m != nil {
		return m.Opts
	}
	return nil
}

type CreateDynamicRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth         `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	Opts                 *opts.DynamicOptions `protobuf:"bytes,1,opt,name=opts,proto3" json:"opts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateDynamicRequest) Reset()         { *m = CreateDynamicRequest{} }
func (m *CreateDynamicRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDynamicRequest) ProtoMessage()    {}
func (*CreateDynamicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d091d947a15a54f7, []int{4}
}

func (m *CreateDynamicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDynamicRequest.Unmarshal(m, b)
}
func (m *CreateDynamicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDynamicRequest.Marshal(b, m, deterministic)
}
func (m *CreateDynamicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDynamicRequest.Merge(m, src)
}
func (m *CreateDynamicRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDynamicRequest.Size(m)
}
func (m *CreateDynamicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDynamicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDynamicRequest proto.InternalMessageInfo

func (m *CreateDynamicRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *CreateDynamicRequest) GetOpts() *opts.DynamicOptions {
	if m != nil {
		return m.Opts
	}
	return nil
}

type CreateDynamicResponse struct {
	Status *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	// ID of the created dynamic destination entry
	DynamicId            string   `protobuf:"bytes,1,opt,name=dynamic_id,json=dynamicId,proto3" json:"dynamic_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDynamicResponse) Reset()         { *m = CreateDynamicResponse{} }
func (m *CreateDynamicResponse) String() string { return proto.CompactTextString(m) }
func (*CreateDynamicResponse) ProtoMessage()    {}
func (*CreateDynamicResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d091d947a15a54f7, []int{5}
}

func (m *CreateDynamicResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDynamicResponse.Unmarshal(m, b)
}
func (m *CreateDynamicResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDynamicResponse.Marshal(b, m, deterministic)
}
func (m *CreateDynamicResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDynamicResponse.Merge(m, src)
}
func (m *CreateDynamicResponse) XXX_Size() int {
	return xxx_messageInfo_CreateDynamicResponse.Size(m)
}
func (m *CreateDynamicResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDynamicResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDynamicResponse proto.InternalMessageInfo

func (m *CreateDynamicResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CreateDynamicResponse) GetDynamicId() string {
	if m != nil {
		return m.DynamicId
	}
	return ""
}

// WARNING: Updating a destination that is in-use can result in missing data during replay
type UpdateDynamicRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth         `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	DynamicId            string               `protobuf:"bytes,1,opt,name=dynamic_id,json=dynamicId,proto3" json:"dynamic_id,omitempty"`
	Opts                 *opts.DynamicOptions `protobuf:"bytes,2,opt,name=opts,proto3" json:"opts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UpdateDynamicRequest) Reset()         { *m = UpdateDynamicRequest{} }
func (m *UpdateDynamicRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateDynamicRequest) ProtoMessage()    {}
func (*UpdateDynamicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d091d947a15a54f7, []int{6}
}

func (m *UpdateDynamicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDynamicRequest.Unmarshal(m, b)
}
func (m *UpdateDynamicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDynamicRequest.Marshal(b, m, deterministic)
}
func (m *UpdateDynamicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDynamicRequest.Merge(m, src)
}
func (m *UpdateDynamicRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateDynamicRequest.Size(m)
}
func (m *UpdateDynamicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDynamicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDynamicRequest proto.InternalMessageInfo

func (m *UpdateDynamicRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *UpdateDynamicRequest) GetDynamicId() string {
	if m != nil {
		return m.DynamicId
	}
	return ""
}

func (m *UpdateDynamicRequest) GetOpts() *opts.DynamicOptions {
	if m != nil {
		return m.Opts
	}
	return nil
}

type UpdateDynamicResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateDynamicResponse) Reset()         { *m = UpdateDynamicResponse{} }
func (m *UpdateDynamicResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateDynamicResponse) ProtoMessage()    {}
func (*UpdateDynamicResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d091d947a15a54f7, []int{7}
}

func (m *UpdateDynamicResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDynamicResponse.Unmarshal(m, b)
}
func (m *UpdateDynamicResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDynamicResponse.Marshal(b, m, deterministic)
}
func (m *UpdateDynamicResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDynamicResponse.Merge(m, src)
}
func (m *UpdateDynamicResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateDynamicResponse.Size(m)
}
func (m *UpdateDynamicResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDynamicResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDynamicResponse proto.InternalMessageInfo

func (m *UpdateDynamicResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// Resume a paused relay
type ResumeDynamicRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	DynamicId            string       `protobuf:"bytes,1,opt,name=dynamic_id,json=dynamicId,proto3" json:"dynamic_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ResumeDynamicRequest) Reset()         { *m = ResumeDynamicRequest{} }
func (m *ResumeDynamicRequest) String() string { return proto.CompactTextString(m) }
func (*ResumeDynamicRequest) ProtoMessage()    {}
func (*ResumeDynamicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d091d947a15a54f7, []int{8}
}

func (m *ResumeDynamicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResumeDynamicRequest.Unmarshal(m, b)
}
func (m *ResumeDynamicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResumeDynamicRequest.Marshal(b, m, deterministic)
}
func (m *ResumeDynamicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResumeDynamicRequest.Merge(m, src)
}
func (m *ResumeDynamicRequest) XXX_Size() int {
	return xxx_messageInfo_ResumeDynamicRequest.Size(m)
}
func (m *ResumeDynamicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResumeDynamicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResumeDynamicRequest proto.InternalMessageInfo

func (m *ResumeDynamicRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *ResumeDynamicRequest) GetDynamicId() string {
	if m != nil {
		return m.DynamicId
	}
	return ""
}

type ResumeDynamicResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ResumeDynamicResponse) Reset()         { *m = ResumeDynamicResponse{} }
func (m *ResumeDynamicResponse) String() string { return proto.CompactTextString(m) }
func (*ResumeDynamicResponse) ProtoMessage()    {}
func (*ResumeDynamicResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d091d947a15a54f7, []int{9}
}

func (m *ResumeDynamicResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResumeDynamicResponse.Unmarshal(m, b)
}
func (m *ResumeDynamicResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResumeDynamicResponse.Marshal(b, m, deterministic)
}
func (m *ResumeDynamicResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResumeDynamicResponse.Merge(m, src)
}
func (m *ResumeDynamicResponse) XXX_Size() int {
	return xxx_messageInfo_ResumeDynamicResponse.Size(m)
}
func (m *ResumeDynamicResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResumeDynamicResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResumeDynamicResponse proto.InternalMessageInfo

func (m *ResumeDynamicResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// Temporarily stop/pause a relay
type StopDynamicRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	DynamicId            string       `protobuf:"bytes,1,opt,name=dynamic_id,json=dynamicId,proto3" json:"dynamic_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StopDynamicRequest) Reset()         { *m = StopDynamicRequest{} }
func (m *StopDynamicRequest) String() string { return proto.CompactTextString(m) }
func (*StopDynamicRequest) ProtoMessage()    {}
func (*StopDynamicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d091d947a15a54f7, []int{10}
}

func (m *StopDynamicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopDynamicRequest.Unmarshal(m, b)
}
func (m *StopDynamicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopDynamicRequest.Marshal(b, m, deterministic)
}
func (m *StopDynamicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopDynamicRequest.Merge(m, src)
}
func (m *StopDynamicRequest) XXX_Size() int {
	return xxx_messageInfo_StopDynamicRequest.Size(m)
}
func (m *StopDynamicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopDynamicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopDynamicRequest proto.InternalMessageInfo

func (m *StopDynamicRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *StopDynamicRequest) GetDynamicId() string {
	if m != nil {
		return m.DynamicId
	}
	return ""
}

type StopDynamicResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StopDynamicResponse) Reset()         { *m = StopDynamicResponse{} }
func (m *StopDynamicResponse) String() string { return proto.CompactTextString(m) }
func (*StopDynamicResponse) ProtoMessage()    {}
func (*StopDynamicResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d091d947a15a54f7, []int{11}
}

func (m *StopDynamicResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopDynamicResponse.Unmarshal(m, b)
}
func (m *StopDynamicResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopDynamicResponse.Marshal(b, m, deterministic)
}
func (m *StopDynamicResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopDynamicResponse.Merge(m, src)
}
func (m *StopDynamicResponse) XXX_Size() int {
	return xxx_messageInfo_StopDynamicResponse.Size(m)
}
func (m *StopDynamicResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StopDynamicResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StopDynamicResponse proto.InternalMessageInfo

func (m *StopDynamicResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type DeleteDynamicRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	DynamicId            string       `protobuf:"bytes,1,opt,name=dynamic_id,json=dynamicId,proto3" json:"dynamic_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DeleteDynamicRequest) Reset()         { *m = DeleteDynamicRequest{} }
func (m *DeleteDynamicRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDynamicRequest) ProtoMessage()    {}
func (*DeleteDynamicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d091d947a15a54f7, []int{12}
}

func (m *DeleteDynamicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDynamicRequest.Unmarshal(m, b)
}
func (m *DeleteDynamicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDynamicRequest.Marshal(b, m, deterministic)
}
func (m *DeleteDynamicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDynamicRequest.Merge(m, src)
}
func (m *DeleteDynamicRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDynamicRequest.Size(m)
}
func (m *DeleteDynamicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDynamicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDynamicRequest proto.InternalMessageInfo

func (m *DeleteDynamicRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *DeleteDynamicRequest) GetDynamicId() string {
	if m != nil {
		return m.DynamicId
	}
	return ""
}

type DeleteDynamicResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeleteDynamicResponse) Reset()         { *m = DeleteDynamicResponse{} }
func (m *DeleteDynamicResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteDynamicResponse) ProtoMessage()    {}
func (*DeleteDynamicResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d091d947a15a54f7, []int{13}
}

func (m *DeleteDynamicResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDynamicResponse.Unmarshal(m, b)
}
func (m *DeleteDynamicResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDynamicResponse.Marshal(b, m, deterministic)
}
func (m *DeleteDynamicResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDynamicResponse.Merge(m, src)
}
func (m *DeleteDynamicResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteDynamicResponse.Size(m)
}
func (m *DeleteDynamicResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDynamicResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDynamicResponse proto.InternalMessageInfo

func (m *DeleteDynamicResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*GetAllDynamicRequest)(nil), "protos.GetAllDynamicRequest")
	proto.RegisterType((*GetAllDynamicResponse)(nil), "protos.GetAllDynamicResponse")
	proto.RegisterType((*GetDynamicRequest)(nil), "protos.GetDynamicRequest")
	proto.RegisterType((*GetDynamicResponse)(nil), "protos.GetDynamicResponse")
	proto.RegisterType((*CreateDynamicRequest)(nil), "protos.CreateDynamicRequest")
	proto.RegisterType((*CreateDynamicResponse)(nil), "protos.CreateDynamicResponse")
	proto.RegisterType((*UpdateDynamicRequest)(nil), "protos.UpdateDynamicRequest")
	proto.RegisterType((*UpdateDynamicResponse)(nil), "protos.UpdateDynamicResponse")
	proto.RegisterType((*ResumeDynamicRequest)(nil), "protos.ResumeDynamicRequest")
	proto.RegisterType((*ResumeDynamicResponse)(nil), "protos.ResumeDynamicResponse")
	proto.RegisterType((*StopDynamicRequest)(nil), "protos.StopDynamicRequest")
	proto.RegisterType((*StopDynamicResponse)(nil), "protos.StopDynamicResponse")
	proto.RegisterType((*DeleteDynamicRequest)(nil), "protos.DeleteDynamicRequest")
	proto.RegisterType((*DeleteDynamicResponse)(nil), "protos.DeleteDynamicResponse")
}

func init() { proto.RegisterFile("ps_dynamic.proto", fileDescriptor_d091d947a15a54f7) }

var fileDescriptor_d091d947a15a54f7 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x4f, 0x4b, 0xeb, 0x40,
	0x14, 0xc5, 0xc9, 0x7b, 0xa5, 0x8f, 0xde, 0x6e, 0x9e, 0x69, 0x02, 0xa1, 0xa5, 0x50, 0xb2, 0xca,
	0xc6, 0x04, 0xaa, 0xb8, 0xb6, 0x5a, 0x29, 0x6e, 0x14, 0x52, 0xdc, 0x88, 0x12, 0xf2, 0x67, 0x6c,
	0x02, 0x49, 0x66, 0x9a, 0xb9, 0x03, 0xfa, 0x29, 0xf4, 0x63, 0xfa, 0x31, 0x24, 0x99, 0x28, 0x36,
	0x82, 0xd6, 0x94, 0xae, 0x26, 0xdc, 0x7b, 0x38, 0xe7, 0xfc, 0xee, 0x22, 0xf0, 0x9f, 0x71, 0x2f,
	0x7a, 0xca, 0xfd, 0x2c, 0x09, 0x6d, 0x56, 0x50, 0xa4, 0x6a, 0xb7, 0x7a, 0xf8, 0x70, 0x14, 0xd2,
	0x2c, 0xa3, 0xb9, 0xc3, 0xb8, 0x27, 0xbf, 0x3c, 0x5f, 0x60, 0x2c, 0x45, 0xc3, 0xf1, 0x97, 0x25,
	0x47, 0x1f, 0x05, 0xaf, 0xd7, 0x43, 0xca, 0x90, 0x97, 0xcb, 0xf2, 0xdd, 0xf4, 0x37, 0x4f, 0x41,
	0x5b, 0x10, 0x9c, 0xa5, 0xe9, 0x5c, 0x8e, 0x5d, 0xb2, 0x16, 0x84, 0xa3, 0x6a, 0x41, 0xa7, 0x0c,
	0x30, 0x9e, 0xaf, 0x26, 0x8a, 0xd5, 0x9f, 0x0e, 0xa4, 0x9c, 0xdb, 0xd2, 0xdf, 0x9e, 0x09, 0x8c,
	0xdd, 0x4a, 0x61, 0x3e, 0x82, 0xde, 0x70, 0xe0, 0x8c, 0xe6, 0x9c, 0xa8, 0x36, 0x74, 0x65, 0x0d,
	0xe3, 0xf5, 0x5f, 0x65, 0xa2, 0x37, 0x4c, 0x96, 0xd5, 0xd6, 0xad, 0x55, 0xaa, 0x03, 0x9d, 0xb2,
	0xa0, 0xa1, 0x4c, 0xfe, 0x5a, 0xfd, 0xe9, 0xe8, 0x5d, 0x5c, 0xce, 0xec, 0xda, 0xfb, 0x9a, 0x61,
	0x42, 0x73, 0xee, 0x56, 0x42, 0xf3, 0x0e, 0x0e, 0x16, 0x04, 0xdb, 0x16, 0x57, 0xc7, 0x00, 0xf5,
	0x2d, 0xbc, 0x24, 0x32, 0x94, 0x89, 0x62, 0xf5, 0xdc, 0x5e, 0x3d, 0xb9, 0x8c, 0x4c, 0x01, 0xea,
	0x67, 0xf7, 0x9d, 0xa1, 0x94, 0xed, 0xa0, 0xd6, 0xa0, 0x9d, 0x17, 0xc4, 0x47, 0xd2, 0x9a, 0xeb,
	0xd7, 0x91, 0x0f, 0xa0, 0x37, 0x22, 0x5b, 0xc2, 0xfe, 0x70, 0xd1, 0x17, 0x05, 0xb4, 0x1b, 0x16,
	0xed, 0xc2, 0xf6, 0x7d, 0xc2, 0x07, 0xfa, 0x9f, 0x6d, 0xd1, 0x17, 0xa0, 0x37, 0x1a, 0xb5, 0x43,
	0x37, 0x3d, 0xd0, 0x5c, 0xc2, 0x45, 0xb6, 0x2f, 0xb4, 0xb2, 0x69, 0x23, 0xa0, 0x65, 0xd3, 0x7b,
	0x50, 0x97, 0x48, 0xd9, 0xbe, 0x7a, 0x5e, 0xc0, 0x60, 0xc3, 0xbe, 0xfd, 0x3d, 0xe7, 0x24, 0x25,
	0xb8, 0xcf, 0x7b, 0x36, 0x02, 0xda, 0x35, 0x3d, 0x3b, 0xb9, 0x3d, 0x5e, 0x25, 0x18, 0x8b, 0xa0,
	0xdc, 0x3b, 0x81, 0x8f, 0x61, 0x1c, 0xd2, 0x82, 0x39, 0x2c, 0x15, 0x59, 0x40, 0x8a, 0x43, 0x1e,
	0xc6, 0x24, 0xf3, 0xb9, 0x13, 0x88, 0x24, 0x8d, 0x9c, 0x15, 0x75, 0xa4, 0x5b, 0x20, 0xff, 0xec,
	0x47, 0x6f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x59, 0x03, 0xb9, 0xf4, 0x05, 0x00, 0x00,
}
