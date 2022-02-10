// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greet/greetPb/BiDirectionalStreaming.proto

package BiDirectionalAPI

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

type GreetingBiDirectional struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	SecondName           string   `protobuf:"bytes,2,opt,name=second_name,json=secondName,proto3" json:"second_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetingBiDirectional) Reset()         { *m = GreetingBiDirectional{} }
func (m *GreetingBiDirectional) String() string { return proto.CompactTextString(m) }
func (*GreetingBiDirectional) ProtoMessage()    {}
func (*GreetingBiDirectional) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e3b9e1b73407e93, []int{0}
}
func (m *GreetingBiDirectional) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetingBiDirectional.Unmarshal(m, b)
}
func (m *GreetingBiDirectional) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetingBiDirectional.Marshal(b, m, deterministic)
}
func (m *GreetingBiDirectional) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetingBiDirectional.Merge(m, src)
}
func (m *GreetingBiDirectional) XXX_Size() int {
	return xxx_messageInfo_GreetingBiDirectional.Size(m)
}
func (m *GreetingBiDirectional) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetingBiDirectional.DiscardUnknown(m)
}

var xxx_messageInfo_GreetingBiDirectional proto.InternalMessageInfo

func (m *GreetingBiDirectional) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *GreetingBiDirectional) GetSecondName() string {
	if m != nil {
		return m.SecondName
	}
	return ""
}

type GreetEveryoneRequest struct {
	Greeting             *GreetingBiDirectional `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GreetEveryoneRequest) Reset()         { *m = GreetEveryoneRequest{} }
func (m *GreetEveryoneRequest) String() string { return proto.CompactTextString(m) }
func (*GreetEveryoneRequest) ProtoMessage()    {}
func (*GreetEveryoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e3b9e1b73407e93, []int{1}
}
func (m *GreetEveryoneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetEveryoneRequest.Unmarshal(m, b)
}
func (m *GreetEveryoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetEveryoneRequest.Marshal(b, m, deterministic)
}
func (m *GreetEveryoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetEveryoneRequest.Merge(m, src)
}
func (m *GreetEveryoneRequest) XXX_Size() int {
	return xxx_messageInfo_GreetEveryoneRequest.Size(m)
}
func (m *GreetEveryoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetEveryoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetEveryoneRequest proto.InternalMessageInfo

func (m *GreetEveryoneRequest) GetGreeting() *GreetingBiDirectional {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetEveryoneResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetEveryoneResponse) Reset()         { *m = GreetEveryoneResponse{} }
func (m *GreetEveryoneResponse) String() string { return proto.CompactTextString(m) }
func (*GreetEveryoneResponse) ProtoMessage()    {}
func (*GreetEveryoneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e3b9e1b73407e93, []int{2}
}
func (m *GreetEveryoneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetEveryoneResponse.Unmarshal(m, b)
}
func (m *GreetEveryoneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetEveryoneResponse.Marshal(b, m, deterministic)
}
func (m *GreetEveryoneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetEveryoneResponse.Merge(m, src)
}
func (m *GreetEveryoneResponse) XXX_Size() int {
	return xxx_messageInfo_GreetEveryoneResponse.Size(m)
}
func (m *GreetEveryoneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetEveryoneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetEveryoneResponse proto.InternalMessageInfo

func (m *GreetEveryoneResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*GreetingBiDirectional)(nil), "BiDirectionalAPI.GreetingBiDirectional")
	proto.RegisterType((*GreetEveryoneRequest)(nil), "BiDirectionalAPI.GreetEveryoneRequest")
	proto.RegisterType((*GreetEveryoneResponse)(nil), "BiDirectionalAPI.GreetEveryoneResponse")
}

func init() {
	proto.RegisterFile("greet/greetPb/BiDirectionalStreaming.proto", fileDescriptor_2e3b9e1b73407e93)
}

var fileDescriptor_2e3b9e1b73407e93 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0x5d, 0x0f, 0xc5, 0x4e, 0x15, 0x24, 0xf8, 0x1f, 0x44, 0xc9, 0xc1, 0x16, 0x0f, 0x1b,
	0xa9, 0x9f, 0xc0, 0xaa, 0x88, 0x17, 0x29, 0xdb, 0x83, 0xa0, 0x07, 0x49, 0xd7, 0x69, 0x08, 0x74,
	0x93, 0x3a, 0x49, 0x17, 0xfc, 0xf6, 0xb2, 0x49, 0x50, 0x76, 0x51, 0xbc, 0x04, 0x32, 0x6f, 0xde,
	0x9b, 0x1f, 0x3c, 0xb8, 0x54, 0x84, 0xe8, 0x45, 0x78, 0xa7, 0x73, 0x31, 0xd1, 0x77, 0x9a, 0xb0,
	0xf4, 0xda, 0x1a, 0xb9, 0x9c, 0x79, 0x42, 0x59, 0x69, 0xa3, 0xf2, 0x15, 0x59, 0x6f, 0xd9, 0x6e,
	0x4b, 0xbd, 0x99, 0x3e, 0xf2, 0x67, 0xd8, 0x7f, 0x68, 0x9c, 0xda, 0xa8, 0x96, 0xc6, 0x4e, 0x01,
	0x16, 0x9a, 0x9c, 0x7f, 0x33, 0xb2, 0xc2, 0xa3, 0xec, 0x3c, 0x1b, 0xf5, 0x8b, 0x7e, 0x98, 0x3c,
	0xc9, 0x0a, 0xd9, 0x19, 0x0c, 0x1c, 0x96, 0xd6, 0xbc, 0x47, 0x7d, 0x33, 0xe8, 0x10, 0x47, 0xcd,
	0x02, 0x7f, 0x85, 0xbd, 0x10, 0x7c, 0x5f, 0x23, 0x7d, 0x5a, 0x83, 0x05, 0x7e, 0xac, 0xd1, 0x79,
	0x76, 0x0b, 0x5b, 0x2a, 0x1d, 0x0c, 0xa9, 0x83, 0xf1, 0x30, 0xef, 0x52, 0xe5, 0xbf, 0x22, 0x15,
	0xdf, 0x46, 0x2e, 0x12, 0xf5, 0x4f, 0xb8, 0x5b, 0x59, 0xe3, 0x90, 0x1d, 0x40, 0x8f, 0xd0, 0xad,
	0x97, 0x3e, 0x11, 0xa7, 0xdf, 0xb8, 0x86, 0xed, 0x60, 0x98, 0x21, 0xd5, 0xba, 0x44, 0xb6, 0x80,
	0x9d, 0x56, 0x00, 0xbb, 0xf8, 0x03, 0xa2, 0x83, 0x7f, 0x32, 0xfc, 0x77, 0x2f, 0x92, 0xf0, 0x8d,
	0x51, 0x76, 0x95, 0x4d, 0x8e, 0x5f, 0x0e, 0x63, 0x07, 0xa2, 0x6b, 0x9b, 0xf7, 0x42, 0x25, 0xd7,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xe2, 0x3a, 0x48, 0xc0, 0x01, 0x00, 0x00,
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
	// Bi directional streamming
	GreetEveryone(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetEveryoneClient, error)
}

type greetServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetServiceClient(cc *grpc.ClientConn) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) GreetEveryone(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetEveryoneClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[0], "/BiDirectionalAPI.GreetService/GreetEveryone", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetEveryoneClient{stream}
	return x, nil
}

type GreetService_GreetEveryoneClient interface {
	Send(*GreetEveryoneRequest) error
	Recv() (*GreetEveryoneResponse, error)
	grpc.ClientStream
}

type greetServiceGreetEveryoneClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetEveryoneClient) Send(m *GreetEveryoneRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceGreetEveryoneClient) Recv() (*GreetEveryoneResponse, error) {
	m := new(GreetEveryoneResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	// Bi directional streamming
	GreetEveryone(GreetService_GreetEveryoneServer) error
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) GreetEveryone(srv GreetService_GreetEveryoneServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetEveryone not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_GreetEveryone_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).GreetEveryone(&greetServiceGreetEveryoneServer{stream})
}

type GreetService_GreetEveryoneServer interface {
	Send(*GreetEveryoneResponse) error
	Recv() (*GreetEveryoneRequest, error)
	grpc.ServerStream
}

type greetServiceGreetEveryoneServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetEveryoneServer) Send(m *GreetEveryoneResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceGreetEveryoneServer) Recv() (*GreetEveryoneRequest, error) {
	m := new(GreetEveryoneRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "BiDirectionalAPI.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetEveryone",
			Handler:       _GreetService_GreetEveryone_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "greet/greetPb/BiDirectionalStreaming.proto",
}
