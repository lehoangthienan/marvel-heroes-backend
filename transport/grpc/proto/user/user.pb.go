// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transport/grpc/proto/user/user.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	transport/grpc/proto/user/user.proto

It has these top-level messages:
	SignInUserRequest
	SignInUserResponse
*/
package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SignInUserRequest struct {
	Username string `protobuf:"bytes,1,opt,name=Username,json=username" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,json=password" json:"Password,omitempty"`
}

func (m *SignInUserRequest) Reset()                    { *m = SignInUserRequest{} }
func (m *SignInUserRequest) String() string            { return proto.CompactTextString(m) }
func (*SignInUserRequest) ProtoMessage()               {}
func (*SignInUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SignInUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SignInUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SignInUserResponse struct {
	Status bool   `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *SignInUserResponse) Reset()                    { *m = SignInUserResponse{} }
func (m *SignInUserResponse) String() string            { return proto.CompactTextString(m) }
func (*SignInUserResponse) ProtoMessage()               {}
func (*SignInUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SignInUserResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *SignInUserResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*SignInUserRequest)(nil), "user.SignInUserRequest")
	proto.RegisterType((*SignInUserResponse)(nil), "user.SignInUserResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ClinicService service

type ClinicServiceClient interface {
	SignInUser(ctx context.Context, in *SignInUserRequest, opts ...grpc.CallOption) (*SignInUserResponse, error)
}

type clinicServiceClient struct {
	cc *grpc.ClientConn
}

func NewClinicServiceClient(cc *grpc.ClientConn) ClinicServiceClient {
	return &clinicServiceClient{cc}
}

func (c *clinicServiceClient) SignInUser(ctx context.Context, in *SignInUserRequest, opts ...grpc.CallOption) (*SignInUserResponse, error) {
	out := new(SignInUserResponse)
	err := grpc.Invoke(ctx, "/user.ClinicService/SignInUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ClinicService service

type ClinicServiceServer interface {
	SignInUser(context.Context, *SignInUserRequest) (*SignInUserResponse, error)
}

func RegisterClinicServiceServer(s *grpc.Server, srv ClinicServiceServer) {
	s.RegisterService(&_ClinicService_serviceDesc, srv)
}

func _ClinicService_SignInUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClinicServiceServer).SignInUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ClinicService/SignInUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClinicServiceServer).SignInUser(ctx, req.(*SignInUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClinicService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.ClinicService",
	HandlerType: (*ClinicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignInUser",
			Handler:    _ClinicService_SignInUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transport/grpc/proto/user/user.proto",
}

func init() { proto.RegisterFile("transport/grpc/proto/user/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xc1, 0x6a, 0xc3, 0x30,
	0x0c, 0x86, 0x97, 0xb1, 0x85, 0x4c, 0xb0, 0xc3, 0xcc, 0xd8, 0x42, 0x4e, 0x23, 0xec, 0xb0, 0x53,
	0x02, 0xdb, 0x13, 0x6c, 0x3b, 0x95, 0x5e, 0x8a, 0x43, 0x1f, 0xc0, 0x4d, 0x45, 0x30, 0x6d, 0x6d,
	0x57, 0xb2, 0xdb, 0xd7, 0x2f, 0x71, 0x0c, 0x2d, 0xe4, 0x22, 0xf4, 0xf1, 0xc3, 0x27, 0xfd, 0xf0,
	0xe9, 0x49, 0x19, 0x76, 0x96, 0x7c, 0x3b, 0x90, 0xeb, 0x5b, 0x47, 0xd6, 0xdb, 0x36, 0x30, 0x52,
	0x1c, 0x4d, 0x64, 0xf1, 0x30, 0xee, 0xf5, 0x12, 0x5e, 0x3a, 0x3d, 0x98, 0x85, 0x59, 0x33, 0x92,
	0xc4, 0x63, 0x40, 0xf6, 0xa2, 0x82, 0x62, 0x44, 0xa3, 0x0e, 0x58, 0x66, 0x1f, 0xd9, 0xd7, 0x93,
	0x2c, 0x42, 0xe2, 0x31, 0x5b, 0x29, 0xe6, 0xb3, 0xa5, 0x6d, 0x79, 0x3f, 0x65, 0x2e, 0x71, 0xfd,
	0x07, 0xe2, 0x56, 0xc6, 0xce, 0x1a, 0x46, 0xf1, 0x06, 0x39, 0x7b, 0xe5, 0x03, 0x47, 0x57, 0x21,
	0x13, 0x89, 0x57, 0x78, 0xf4, 0x76, 0x87, 0x26, 0x69, 0x26, 0xf8, 0x96, 0xf0, 0xfc, 0xbf, 0xd7,
	0x46, 0xf7, 0x1d, 0xd2, 0x49, 0xf7, 0x28, 0x7e, 0x01, 0xae, 0x52, 0xf1, 0xde, 0xc4, 0x0a, 0xb3,
	0x9f, 0xab, 0x72, 0x1e, 0x4c, 0xf7, 0xeb, 0xbb, 0x4d, 0x1e, 0x1b, 0xff, 0x5c, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x6e, 0x7a, 0xac, 0x58, 0x19, 0x01, 0x00, 0x00,
}
