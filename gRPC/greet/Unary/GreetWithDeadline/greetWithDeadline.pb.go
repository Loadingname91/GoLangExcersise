// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greet/greetPb/greetWithDeadline.proto

package GreetWithDeadline

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
	return fileDescriptor_f1fe1ab054ece335, []int{0}
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

type GreetWithDealineRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetWithDealineRequest) Reset()         { *m = GreetWithDealineRequest{} }
func (m *GreetWithDealineRequest) String() string { return proto.CompactTextString(m) }
func (*GreetWithDealineRequest) ProtoMessage()    {}
func (*GreetWithDealineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1fe1ab054ece335, []int{1}
}
func (m *GreetWithDealineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetWithDealineRequest.Unmarshal(m, b)
}
func (m *GreetWithDealineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetWithDealineRequest.Marshal(b, m, deterministic)
}
func (m *GreetWithDealineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetWithDealineRequest.Merge(m, src)
}
func (m *GreetWithDealineRequest) XXX_Size() int {
	return xxx_messageInfo_GreetWithDealineRequest.Size(m)
}
func (m *GreetWithDealineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetWithDealineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetWithDealineRequest proto.InternalMessageInfo

func (m *GreetWithDealineRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetWithDeadlineResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetWithDeadlineResponse) Reset()         { *m = GreetWithDeadlineResponse{} }
func (m *GreetWithDeadlineResponse) String() string { return proto.CompactTextString(m) }
func (*GreetWithDeadlineResponse) ProtoMessage()    {}
func (*GreetWithDeadlineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1fe1ab054ece335, []int{2}
}
func (m *GreetWithDeadlineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetWithDeadlineResponse.Unmarshal(m, b)
}
func (m *GreetWithDeadlineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetWithDeadlineResponse.Marshal(b, m, deterministic)
}
func (m *GreetWithDeadlineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetWithDeadlineResponse.Merge(m, src)
}
func (m *GreetWithDeadlineResponse) XXX_Size() int {
	return xxx_messageInfo_GreetWithDeadlineResponse.Size(m)
}
func (m *GreetWithDeadlineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetWithDeadlineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetWithDeadlineResponse proto.InternalMessageInfo

func (m *GreetWithDeadlineResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "GreetDeadline.Greeting")
	proto.RegisterType((*GreetWithDealineRequest)(nil), "GreetDeadline.GreetWithDealineRequest")
	proto.RegisterType((*GreetWithDeadlineResponse)(nil), "GreetDeadline.GreetWithDeadlineResponse")
}

func init() {
	proto.RegisterFile("greet/greetPb/greetWithDeadline.proto", fileDescriptor_f1fe1ab054ece335)
}

var fileDescriptor_f1fe1ab054ece335 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0x2f, 0x4a, 0x4d,
	0x2d, 0xd1, 0x07, 0x93, 0x01, 0x49, 0x10, 0x3a, 0x3c, 0xb3, 0x24, 0xc3, 0x25, 0x35, 0x31, 0x25,
	0x27, 0x33, 0x2f, 0x55, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x88, 0xd7, 0x1d, 0x24, 0x01, 0x13,
	0x54, 0xf2, 0xe2, 0xe2, 0x00, 0x0b, 0x64, 0xe6, 0xa5, 0x0b, 0xc9, 0x72, 0x71, 0xa5, 0x65, 0x16,
	0x15, 0x97, 0xc4, 0xe7, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x71, 0x82,
	0x45, 0xfc, 0x12, 0x73, 0x53, 0x85, 0xe4, 0xb9, 0xb8, 0x8b, 0x53, 0x93, 0xf3, 0xf3, 0x52, 0x20,
	0xf2, 0x4c, 0x60, 0x79, 0x2e, 0x88, 0x10, 0x48, 0x81, 0x92, 0x1f, 0x97, 0xb8, 0x3b, 0x92, 0xad,
	0x20, 0xf3, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x8c, 0xb9, 0x38, 0xd2, 0xa1, 0xd6,
	0x80, 0x0d, 0xe6, 0x36, 0x12, 0xd7, 0x43, 0x71, 0x88, 0x1e, 0xcc, 0x15, 0x41, 0x70, 0x85, 0x4a,
	0xc6, 0x5c, 0x92, 0xee, 0xe8, 0xbe, 0x08, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12,
	0xe3, 0x62, 0x2b, 0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0x81, 0x3a, 0x14, 0xca, 0x33, 0x2a, 0xe3, 0xe2,
	0x01, 0x6b, 0x0a, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x4a, 0xe3, 0x12, 0x40, 0x77, 0x94,
	0x90, 0x1a, 0x36, 0xbb, 0x31, 0x5d, 0x2d, 0xa5, 0x81, 0x47, 0x1d, 0x8a, 0x6b, 0x94, 0x18, 0x9c,
	0x24, 0xa3, 0xc4, 0x43, 0xf3, 0x12, 0x8b, 0x2a, 0xf5, 0x31, 0x14, 0x25, 0xb1, 0x81, 0x43, 0xde,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x95, 0x53, 0xcc, 0xa2, 0x01, 0x00, 0x00,
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
	// Unary with Deadlines
	GreetWithDealine(ctx context.Context, in *GreetWithDealineRequest, opts ...grpc.CallOption) (*GreetWithDeadlineResponse, error)
}

type greetServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetServiceClient(cc *grpc.ClientConn) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) GreetWithDealine(ctx context.Context, in *GreetWithDealineRequest, opts ...grpc.CallOption) (*GreetWithDeadlineResponse, error) {
	out := new(GreetWithDeadlineResponse)
	err := c.cc.Invoke(ctx, "/GreetDeadline.GreetService/GreetWithDealine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	// Unary with Deadlines
	GreetWithDealine(context.Context, *GreetWithDealineRequest) (*GreetWithDeadlineResponse, error)
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) GreetWithDealine(ctx context.Context, req *GreetWithDealineRequest) (*GreetWithDeadlineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GreetWithDealine not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_GreetWithDealine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetWithDealineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).GreetWithDealine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GreetDeadline.GreetService/GreetWithDealine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).GreetWithDealine(ctx, req.(*GreetWithDealineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GreetDeadline.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GreetWithDealine",
			Handler:    _GreetService_GreetWithDealine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greet/greetPb/greetWithDeadline.proto",
}
