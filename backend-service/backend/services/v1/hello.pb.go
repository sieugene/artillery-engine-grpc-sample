// Code generated by protoc-gen-go. DO NOT EDIT.
// source: backend/services/v1/hello.proto

package backend_services_v1

import (
	v1 "github.com/kenju/artillery-engine-grpc/sample/backend-service/backend/resources/v1"
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

type HelloRequest struct {
	Id                   int32       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Platform             v1.Platform `protobuf:"varint,3,opt,name=platform,proto3,enum=backend.resources.v1.Platform" json:"platform,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cb2e5bedbde1b42, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloRequest) GetPlatform() v1.Platform {
	if m != nil {
		return m.Platform
	}
	return v1.Platform_Web
}

type HelloResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	User                 *v1.User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cb2e5bedbde1b42, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *HelloResponse) GetUser() *v1.User {
	if m != nil {
		return m.User
	}
	return nil
}

type ByeRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ByeRequest) Reset()         { *m = ByeRequest{} }
func (m *ByeRequest) String() string { return proto.CompactTextString(m) }
func (*ByeRequest) ProtoMessage()    {}
func (*ByeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cb2e5bedbde1b42, []int{2}
}

func (m *ByeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ByeRequest.Unmarshal(m, b)
}
func (m *ByeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ByeRequest.Marshal(b, m, deterministic)
}
func (m *ByeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ByeRequest.Merge(m, src)
}
func (m *ByeRequest) XXX_Size() int {
	return xxx_messageInfo_ByeRequest.Size(m)
}
func (m *ByeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ByeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ByeRequest proto.InternalMessageInfo

func (m *ByeRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ByeResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ByeResponse) Reset()         { *m = ByeResponse{} }
func (m *ByeResponse) String() string { return proto.CompactTextString(m) }
func (*ByeResponse) ProtoMessage()    {}
func (*ByeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cb2e5bedbde1b42, []int{3}
}

func (m *ByeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ByeResponse.Unmarshal(m, b)
}
func (m *ByeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ByeResponse.Marshal(b, m, deterministic)
}
func (m *ByeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ByeResponse.Merge(m, src)
}
func (m *ByeResponse) XXX_Size() int {
	return xxx_messageInfo_ByeResponse.Size(m)
}
func (m *ByeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ByeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ByeResponse proto.InternalMessageInfo

func (m *ByeResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "backend.services.v1.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "backend.services.v1.HelloResponse")
	proto.RegisterType((*ByeRequest)(nil), "backend.services.v1.ByeRequest")
	proto.RegisterType((*ByeResponse)(nil), "backend.services.v1.ByeResponse")
}

func init() { proto.RegisterFile("backend/services/v1/hello.proto", fileDescriptor_9cb2e5bedbde1b42) }

var fileDescriptor_9cb2e5bedbde1b42 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4f, 0xb3, 0x40,
	0x10, 0x86, 0x03, 0x6d, 0xbf, 0xcf, 0x4e, 0xb5, 0x87, 0xf5, 0x42, 0x88, 0x11, 0xc4, 0x83, 0x9c,
	0x96, 0x80, 0x37, 0x8f, 0x3d, 0xf5, 0x64, 0xcc, 0x1a, 0x0f, 0x1e, 0x69, 0x19, 0x95, 0x08, 0x2c,
	0xee, 0x00, 0x49, 0xff, 0x91, 0x3f, 0xd3, 0x74, 0x59, 0xaa, 0x07, 0xaa, 0x37, 0x26, 0x3c, 0xb3,
	0xcf, 0xcc, 0x3b, 0xe0, 0x6d, 0xd2, 0xed, 0x3b, 0x56, 0x59, 0x44, 0xa8, 0xba, 0x7c, 0x8b, 0x14,
	0x75, 0x71, 0xf4, 0x86, 0x45, 0x21, 0x79, 0xad, 0x64, 0x23, 0xd9, 0xb9, 0x01, 0xf8, 0x00, 0xf0,
	0x2e, 0x76, 0xaf, 0x87, 0x2e, 0x85, 0x24, 0x5b, 0x65, 0xda, 0xea, 0x22, 0x6d, 0x5e, 0xa4, 0x2a,
	0xfb, 0x4e, 0xd7, 0x1b, 0x85, 0x5a, 0x42, 0xd5, 0x03, 0x41, 0x05, 0xa7, 0xeb, 0xbd, 0x49, 0xe0,
	0x47, 0x8b, 0xd4, 0xb0, 0x25, 0xd8, 0x79, 0xe6, 0x58, 0xbe, 0x15, 0xce, 0x84, 0x9d, 0x67, 0x8c,
	0xc1, 0xb4, 0x4a, 0x4b, 0x74, 0x6c, 0xdf, 0x0a, 0xe7, 0x42, 0x7f, 0xb3, 0x3b, 0x38, 0x19, 0x34,
	0xce, 0xc4, 0xb7, 0xc2, 0x65, 0x72, 0xc9, 0x87, 0x09, 0x0f, 0x1e, 0xde, 0xc5, 0xfc, 0xc1, 0x50,
	0xe2, 0xc0, 0x07, 0xcf, 0x70, 0x66, 0x7c, 0x54, 0xcb, 0x8a, 0x90, 0x39, 0xf0, 0xbf, 0x44, 0xa2,
	0xf4, 0x15, 0xb5, 0x75, 0x2e, 0x86, 0x92, 0x71, 0x98, 0xee, 0x07, 0xd5, 0xea, 0x45, 0xe2, 0x8e,
	0x2b, 0x9e, 0x08, 0x95, 0xd0, 0x5c, 0x70, 0x01, 0xb0, 0xda, 0xe1, 0x91, 0x45, 0x82, 0x1b, 0x58,
	0xe8, 0xbf, 0x7f, 0x69, 0x93, 0x4f, 0xcb, 0x44, 0xf2, 0xd8, 0x87, 0xcd, 0xee, 0x61, 0xa6, 0x6b,
	0x76, 0xc5, 0x47, 0xee, 0xc0, 0x7f, 0xc6, 0xe7, 0x06, 0xbf, 0x21, 0x46, 0xbd, 0x86, 0xc9, 0x6a,
	0x87, 0xcc, 0x1b, 0x45, 0xbf, 0x37, 0x70, 0xfd, 0xe3, 0x40, 0xff, 0xd2, 0xe6, 0x9f, 0xbe, 0xe1,
	0xed, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x71, 0x82, 0xf8, 0x41, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	Bye(ctx context.Context, in *ByeRequest, opts ...grpc.CallOption) (*ByeResponse, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/backend.services.v1.HelloService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) Bye(ctx context.Context, in *ByeRequest, opts ...grpc.CallOption) (*ByeResponse, error) {
	out := new(ByeResponse)
	err := c.cc.Invoke(ctx, "/backend.services.v1.HelloService/Bye", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	Bye(context.Context, *ByeRequest) (*ByeResponse, error)
}

// UnimplementedHelloServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (*UnimplementedHelloServiceServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (*UnimplementedHelloServiceServer) Bye(ctx context.Context, req *ByeRequest) (*ByeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bye not implemented")
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend.services.v1.HelloService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_Bye_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Bye(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend.services.v1.HelloService/Bye",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Bye(ctx, req.(*ByeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "backend.services.v1.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _HelloService_Hello_Handler,
		},
		{
			MethodName: "Bye",
			Handler:    _HelloService_Bye_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend/services/v1/hello.proto",
}