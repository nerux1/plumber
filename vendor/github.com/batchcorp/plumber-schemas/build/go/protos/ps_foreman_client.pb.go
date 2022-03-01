// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_foreman_client.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type RegisterRequest struct {
	// API token obtained from https://console.batch.sh/account/security
	ApiToken string `protobuf:"bytes,1,opt,name=api_token,json=apiToken,proto3" json:"api_token,omitempty"`
	// Plumber cluster ID
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Token used to authenticate calls to this plumber instance
	// This is filled out by plumber, and then stored in memory while it is connected
	// All requests back to plumber need to provide this token
	PlumberToken         string   `protobuf:"bytes,3,opt,name=plumber_token,json=plumberToken,proto3" json:"plumber_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8b3f3eaace7db4c, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetApiToken() string {
	if m != nil {
		return m.ApiToken
	}
	return ""
}

func (m *RegisterRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *RegisterRequest) GetPlumberToken() string {
	if m != nil {
		return m.PlumberToken
	}
	return ""
}

type RegisterResponse struct {
	// Authorized or not
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// Error message if any
	// Will be empty on success
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8b3f3eaace7db4c, []int{1}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RegisterResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "protos.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "protos.RegisterResponse")
}

func init() { proto.RegisterFile("ps_foreman_client.proto", fileDescriptor_e8b3f3eaace7db4c) }

var fileDescriptor_e8b3f3eaace7db4c = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x50, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0x65, 0x15, 0x74, 0x3b, 0xb8, 0x28, 0xb9, 0x6c, 0x51, 0x04, 0x59, 0x2f, 0x5e, 0x6c, 0x41,
	0xc5, 0x9b, 0x17, 0x85, 0x05, 0x2f, 0x1e, 0x8a, 0x27, 0x2f, 0x25, 0x49, 0xc7, 0x36, 0xd8, 0x34,
	0x31, 0x93, 0xfc, 0x7f, 0x69, 0x92, 0x22, 0xe8, 0x29, 0xbc, 0x79, 0x93, 0xf7, 0x31, 0xb0, 0xb5,
	0xd4, 0x7e, 0x1a, 0x87, 0x9a, 0x4f, 0xad, 0x1c, 0x15, 0x4e, 0xbe, 0xb2, 0xce, 0x78, 0xc3, 0x8e,
	0xe2, 0x43, 0x3b, 0x07, 0xa7, 0x0d, 0xf6, 0x8a, 0x3c, 0xba, 0x06, 0xbf, 0x03, 0x92, 0x67, 0x17,
	0x50, 0x70, 0xab, 0x5a, 0x6f, 0xbe, 0x70, 0x2a, 0x57, 0x57, 0xab, 0x9b, 0xa2, 0x59, 0x73, 0xab,
	0xde, 0x67, 0xcc, 0x2e, 0x01, 0xe4, 0x18, 0xe6, 0xf5, 0x56, 0x75, 0xe5, 0x41, 0x64, 0x8b, 0x3c,
	0x79, 0xed, 0xd8, 0x35, 0x6c, 0xec, 0x18, 0xb4, 0x40, 0x97, 0xff, 0x1f, 0xc6, 0x8d, 0x93, 0x3c,
	0x8c, 0x1a, 0xbb, 0x3d, 0x9c, 0xfd, 0x7a, 0x92, 0x35, 0x13, 0x21, 0x2b, 0xe1, 0x98, 0x82, 0x94,
	0x48, 0x14, 0x2d, 0xd7, 0xcd, 0x02, 0x67, 0x46, 0x23, 0x11, 0xef, 0x31, 0xdb, 0x2d, 0xf0, 0xee,
	0x0d, 0x36, 0xfb, 0xd4, 0xed, 0x25, 0x56, 0x63, 0x4f, 0xb0, 0x5e, 0x84, 0xd9, 0x36, 0x15, 0xa5,
	0xea, 0x4f, 0xbd, 0xf3, 0xf2, 0x3f, 0x91, 0x32, 0x3c, 0x3f, 0x7e, 0x3c, 0xf4, 0xca, 0x0f, 0x41,
	0x54, 0xd2, 0xe8, 0x5a, 0x70, 0x2f, 0x07, 0x69, 0x9c, 0xad, 0x73, 0xf8, 0x5b, 0x92, 0x03, 0x6a,
	0x4e, 0xb5, 0x08, 0x6a, 0xec, 0xea, 0xde, 0xd4, 0x49, 0x48, 0xa4, 0x5b, 0xde, 0xff, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xaf, 0xef, 0x5f, 0x20, 0x6d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ForemanClientClient is the client API for ForemanClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ForemanClientClient interface {
	// Used by plumber to connect to Foreman and announce its presence.
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type foremanClientClient struct {
	cc *grpc.ClientConn
}

func NewForemanClientClient(cc *grpc.ClientConn) ForemanClientClient {
	return &foremanClientClient{cc}
}

func (c *foremanClientClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/protos.ForemanClient/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ForemanClientServer is the server API for ForemanClient service.
type ForemanClientServer interface {
	// Used by plumber to connect to Foreman and announce its presence.
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
}

// UnimplementedForemanClientServer can be embedded to have forward compatible implementations.
type UnimplementedForemanClientServer struct {
}

func (*UnimplementedForemanClientServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterForemanClientServer(s *grpc.Server, srv ForemanClientServer) {
	s.RegisterService(&_ForemanClient_serviceDesc, srv)
}

func _ForemanClient_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForemanClientServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ForemanClient/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForemanClientServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ForemanClient_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ForemanClient",
	HandlerType: (*ForemanClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _ForemanClient_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ps_foreman_client.proto",
}