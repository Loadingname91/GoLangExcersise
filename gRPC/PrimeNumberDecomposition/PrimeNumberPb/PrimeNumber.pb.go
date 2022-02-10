// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: PrimeNUmberDecomposition/PrimeNumber_proto/PrimeNumber.proto

package PrimeNumberPb

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

type PrimeNumber struct {
	GivenNumber          int32    `protobuf:"varint,1,opt,name=given_number,json=givenNumber,proto3" json:"given_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumber) Reset()         { *m = PrimeNumber{} }
func (m *PrimeNumber) String() string { return proto.CompactTextString(m) }
func (*PrimeNumber) ProtoMessage()    {}
func (*PrimeNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_26dc1d888eecfa3f, []int{0}
}
func (m *PrimeNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumber.Unmarshal(m, b)
}
func (m *PrimeNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumber.Marshal(b, m, deterministic)
}
func (m *PrimeNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumber.Merge(m, src)
}
func (m *PrimeNumber) XXX_Size() int {
	return xxx_messageInfo_PrimeNumber.Size(m)
}
func (m *PrimeNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumber.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumber proto.InternalMessageInfo

func (m *PrimeNumber) GetGivenNumber() int32 {
	if m != nil {
		return m.GivenNumber
	}
	return 0
}

type PrimeNumberRequest struct {
	NumberGiven          *PrimeNumber `protobuf:"bytes,1,opt,name=number_given,json=numberGiven,proto3" json:"number_given,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PrimeNumberRequest) Reset()         { *m = PrimeNumberRequest{} }
func (m *PrimeNumberRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberRequest) ProtoMessage()    {}
func (*PrimeNumberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_26dc1d888eecfa3f, []int{1}
}
func (m *PrimeNumberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberRequest.Unmarshal(m, b)
}
func (m *PrimeNumberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberRequest.Marshal(b, m, deterministic)
}
func (m *PrimeNumberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberRequest.Merge(m, src)
}
func (m *PrimeNumberRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberRequest.Size(m)
}
func (m *PrimeNumberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberRequest proto.InternalMessageInfo

func (m *PrimeNumberRequest) GetNumberGiven() *PrimeNumber {
	if m != nil {
		return m.NumberGiven
	}
	return nil
}

type PrimeNumberResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberResponse) Reset()         { *m = PrimeNumberResponse{} }
func (m *PrimeNumberResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberResponse) ProtoMessage()    {}
func (*PrimeNumberResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_26dc1d888eecfa3f, []int{2}
}
func (m *PrimeNumberResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberResponse.Unmarshal(m, b)
}
func (m *PrimeNumberResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberResponse.Marshal(b, m, deterministic)
}
func (m *PrimeNumberResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberResponse.Merge(m, src)
}
func (m *PrimeNumberResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberResponse.Size(m)
}
func (m *PrimeNumberResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberResponse proto.InternalMessageInfo

func (m *PrimeNumberResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*PrimeNumber)(nil), "PrimeNumberProto.PrimeNumber")
	proto.RegisterType((*PrimeNumberRequest)(nil), "PrimeNumberProto.PrimeNumberRequest")
	proto.RegisterType((*PrimeNumberResponse)(nil), "PrimeNumberProto.PrimeNumberResponse")
}

func init() {
	proto.RegisterFile("PrimeNUmberDecomposition/PrimeNumber_proto/PrimeNumber.proto", fileDescriptor_26dc1d888eecfa3f)
}

var fileDescriptor_26dc1d888eecfa3f = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x09, 0x28, 0xca, 0xcc,
	0x4d, 0xf5, 0x0b, 0xcd, 0x4d, 0x4a, 0x2d, 0x72, 0x49, 0x4d, 0xce, 0xcf, 0x2d, 0xc8, 0x2f, 0xce,
	0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x87, 0x48, 0x94, 0x82, 0x24, 0xe2, 0x0b, 0x8a, 0xf2, 0x4b, 0xf2,
	0x91, 0x45, 0xf4, 0xc0, 0x22, 0x42, 0x02, 0x48, 0x42, 0x01, 0x20, 0x11, 0x25, 0x03, 0x2e, 0x6e,
	0x24, 0x31, 0x21, 0x45, 0x2e, 0x9e, 0xf4, 0xcc, 0xb2, 0xd4, 0xbc, 0xf8, 0x3c, 0x30, 0x5f, 0x82,
	0x51, 0x81, 0x51, 0x83, 0x35, 0x88, 0x1b, 0x2c, 0x06, 0x51, 0xa2, 0x14, 0xc6, 0x25, 0x84, 0xa4,
	0x23, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0xc8, 0x81, 0x8b, 0x07, 0xa2, 0x25, 0x1e, 0xac,
	0x16, 0xac, 0x91, 0xdb, 0x48, 0x56, 0x0f, 0xdd, 0x42, 0x64, 0x81, 0x20, 0x6e, 0x88, 0x16, 0x77,
	0x90, 0x0e, 0x25, 0x5d, 0x2e, 0x61, 0x14, 0x73, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xc4,
	0xb8, 0xd8, 0x8a, 0x52, 0x8b, 0x4b, 0x73, 0x4a, 0xa0, 0x6e, 0x81, 0xf2, 0x8c, 0x1a, 0x19, 0x51,
	0xdc, 0x11, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x94, 0xcd, 0x25, 0x81, 0x24, 0x8a, 0x12,
	0x42, 0x42, 0x2a, 0xf8, 0x5d, 0x03, 0xf1, 0x89, 0x94, 0x2a, 0x01, 0x55, 0x10, 0x77, 0x29, 0x31,
	0x18, 0x30, 0x3a, 0x09, 0x44, 0xf1, 0x21, 0x07, 0x72, 0x40, 0x52, 0x12, 0x1b, 0x38, 0x9c, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x4f, 0xb7, 0x1f, 0xa7, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PrimeNumberServiceClient is the client API for PrimeNumberService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PrimeNumberServiceClient interface {
	// Server Streaming
	PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberRequest, opts ...grpc.CallOption) (PrimeNumberService_PrimeNumberDecompositionClient, error)
}

type primeNumberServiceClient struct {
	cc *grpc.ClientConn
}

func NewPrimeNumberServiceClient(cc *grpc.ClientConn) PrimeNumberServiceClient {
	return &primeNumberServiceClient{cc}
}

func (c *primeNumberServiceClient) PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberRequest, opts ...grpc.CallOption) (PrimeNumberService_PrimeNumberDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PrimeNumberService_serviceDesc.Streams[0], "/PrimeNumberProto.PrimeNumberService/PrimeNumberDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &primeNumberServicePrimeNumberDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PrimeNumberService_PrimeNumberDecompositionClient interface {
	Recv() (*PrimeNumberResponse, error)
	grpc.ClientStream
}

type primeNumberServicePrimeNumberDecompositionClient struct {
	grpc.ClientStream
}

func (x *primeNumberServicePrimeNumberDecompositionClient) Recv() (*PrimeNumberResponse, error) {
	m := new(PrimeNumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PrimeNumberServiceServer is the server API for PrimeNumberService service.
type PrimeNumberServiceServer interface {
	// Server Streaming
	PrimeNumberDecomposition(*PrimeNumberRequest, PrimeNumberService_PrimeNumberDecompositionServer) error
}

// UnimplementedPrimeNumberServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPrimeNumberServiceServer struct {
}

func (*UnimplementedPrimeNumberServiceServer) PrimeNumberDecomposition(req *PrimeNumberRequest, srv PrimeNumberService_PrimeNumberDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNumberDecomposition not implemented")
}

func RegisterPrimeNumberServiceServer(s *grpc.Server, srv PrimeNumberServiceServer) {
	s.RegisterService(&_PrimeNumberService_serviceDesc, srv)
}

func _PrimeNumberService_PrimeNumberDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PrimeNumberServiceServer).PrimeNumberDecomposition(m, &primeNumberServicePrimeNumberDecompositionServer{stream})
}

type PrimeNumberService_PrimeNumberDecompositionServer interface {
	Send(*PrimeNumberResponse) error
	grpc.ServerStream
}

type primeNumberServicePrimeNumberDecompositionServer struct {
	grpc.ServerStream
}

func (x *primeNumberServicePrimeNumberDecompositionServer) Send(m *PrimeNumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _PrimeNumberService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PrimeNumberProto.PrimeNumberService",
	HandlerType: (*PrimeNumberServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDecomposition",
			Handler:       _PrimeNumberService_PrimeNumberDecomposition_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "PrimeNUmberDecomposition/PrimeNumber_proto/PrimeNumber.proto",
}
