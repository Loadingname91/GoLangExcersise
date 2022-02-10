// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greet/greetPb/UnaryGreet.proto

package GreetAPI

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Greeting struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	SecondName           string   `protobuf:"bytes,2,opt,name=second_name,json=secondName,proto3" json:"second_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d82eb17080e41fe, []int{0}
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

func (m *Greeting) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Greeting) GetSecondName() string {
	if m != nil {
		return m.SecondName
	}
	return ""
}

type GreetRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d82eb17080e41fe, []int{1}
}
func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetResponse) Reset()         { *m = GreetResponse{} }
func (m *GreetResponse) String() string { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()    {}
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d82eb17080e41fe, []int{2}
}
func (m *GreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetResponse.Unmarshal(m, b)
}
func (m *GreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetResponse.Marshal(b, m, deterministic)
}
func (m *GreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResponse.Merge(m, src)
}
func (m *GreetResponse) XXX_Size() int {
	return xxx_messageInfo_GreetResponse.Size(m)
}
func (m *GreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResponse proto.InternalMessageInfo

func (m *GreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "UnaryGreet.Greeting")
	proto.RegisterType((*GreetRequest)(nil), "UnaryGreet.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "UnaryGreet.GreetResponse")
}

func init() { proto.RegisterFile("greet/greetPb/UnaryGreet.proto", fileDescriptor_2d82eb17080e41fe) }

var fileDescriptor_2d82eb17080e41fe = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0x2f, 0x4a, 0x4d,
	0x2d, 0xd1, 0x07, 0x93, 0x01, 0x49, 0xfa, 0xa1, 0x79, 0x89, 0x45, 0x95, 0xee, 0x20, 0x8e, 0x5e,
	0x41, 0x51, 0x7e, 0x49, 0xbe, 0x10, 0x17, 0x42, 0x44, 0xc9, 0x8b, 0x8b, 0x03, 0xcc, 0xc8, 0xcc,
	0x4b, 0x17, 0x92, 0xe5, 0xe2, 0x4a, 0xcb, 0x2c, 0x2a, 0x2e, 0x89, 0xcf, 0x4b, 0xcc, 0x4d, 0x95,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x04, 0x8b, 0xf8, 0x25, 0xe6, 0xa6, 0x0a, 0xc9, 0x73,
	0x71, 0x17, 0xa7, 0x26, 0xe7, 0xe7, 0xa5, 0x40, 0xe4, 0x99, 0xc0, 0xf2, 0x5c, 0x10, 0x21, 0x90,
	0x02, 0x25, 0x07, 0x2e, 0x1e, 0xb0, 0x59, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x06,
	0x5c, 0x1c, 0xe9, 0x50, 0xb3, 0xc1, 0xa6, 0x71, 0x1b, 0x89, 0xe8, 0x21, 0x39, 0x06, 0x66, 0x6f,
	0x10, 0x5c, 0x95, 0x92, 0x3a, 0x17, 0x2f, 0xd4, 0x84, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21,
	0x31, 0x2e, 0xb6, 0xa2, 0xd4, 0xe2, 0xd2, 0x9c, 0x12, 0xa8, 0x73, 0xa0, 0x3c, 0x23, 0x3f, 0xa8,
	0x55, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x76, 0x5c, 0xac, 0x60, 0xbe, 0x90, 0x04,
	0x86, 0x0d, 0x50, 0xd7, 0x48, 0x49, 0x62, 0x91, 0x81, 0xd8, 0xa2, 0xc4, 0xe0, 0x24, 0x10, 0xc5,
	0x07, 0x96, 0xd5, 0x07, 0x4b, 0x38, 0x06, 0x78, 0x26, 0xb1, 0x81, 0xc3, 0xca, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x08, 0x7c, 0x69, 0xf7, 0x4d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	// Unary
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
}

type greetServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetServiceClient(cc *grpc.ClientConn) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/UnaryGreet.GreetService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	// Unary
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) Greet(ctx context.Context, req *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UnaryGreet.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UnaryGreet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greet/greetPb/UnaryGreet.proto",
}
