// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Greeting struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{1}
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

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{2}
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

func (m *HelloResponse) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

func init() {
	proto.RegisterType((*Greeting)(nil), "example.Greeting")
	proto.RegisterType((*HelloRequest)(nil), "example.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "example.HelloResponse")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xe4, 0xb8,
	0x38, 0xdc, 0x8b, 0x52, 0x53, 0x4b, 0x32, 0xf3, 0xd2, 0x85, 0x84, 0xb8, 0x58, 0x4a, 0x52, 0x2b,
	0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x25, 0x2e, 0x1e, 0x8f, 0xd4,
	0x9c, 0x9c, 0xfc, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x90, 0x9a, 0xbc, 0xc4, 0xdc, 0x54,
	0x98, 0x1a, 0x10, 0x5b, 0xc9, 0x8e, 0x8b, 0x17, 0xaa, 0xa6, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55,
	0x48, 0x97, 0x8b, 0x23, 0x1d, 0x6a, 0x28, 0x58, 0x21, 0xb7, 0x91, 0xa0, 0x1e, 0xcc, 0x7e, 0x98,
	0x6d, 0x41, 0x70, 0x25, 0x46, 0x1e, 0x5c, 0x7c, 0xae, 0x10, 0xd9, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc,
	0xe4, 0x54, 0x21, 0x33, 0x2e, 0x56, 0xb0, 0x89, 0x42, 0xa2, 0x70, 0x7d, 0xc8, 0xae, 0x90, 0x12,
	0x43, 0x17, 0x86, 0x58, 0x9c, 0xc4, 0x06, 0xf6, 0x9d, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc1,
	0x98, 0xab, 0xfc, 0xee, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExampleServiceClient is the client API for ExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleServiceClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type exampleServiceClient struct {
	cc *grpc.ClientConn
}

func NewExampleServiceClient(cc *grpc.ClientConn) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/example.ExampleService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServiceServer is the server API for ExampleService service.
type ExampleServiceServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterExampleServiceServer(s *grpc.Server, srv ExampleServiceServer) {
	s.RegisterService(&_ExampleService_serviceDesc, srv)
}

func _ExampleService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.ExampleService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExampleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _ExampleService_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example.proto",
}
