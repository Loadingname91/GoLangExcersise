// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: FindMaxium/FindMaximum_proto/FindMaximum.proto

package FindMaximumPb

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

type FindMaximumRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaximumRequest) Reset()         { *m = FindMaximumRequest{} }
func (m *FindMaximumRequest) String() string { return proto.CompactTextString(m) }
func (*FindMaximumRequest) ProtoMessage()    {}
func (*FindMaximumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_49c82b0727b06490, []int{0}
}
func (m *FindMaximumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaximumRequest.Unmarshal(m, b)
}
func (m *FindMaximumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaximumRequest.Marshal(b, m, deterministic)
}
func (m *FindMaximumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaximumRequest.Merge(m, src)
}
func (m *FindMaximumRequest) XXX_Size() int {
	return xxx_messageInfo_FindMaximumRequest.Size(m)
}
func (m *FindMaximumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaximumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaximumRequest proto.InternalMessageInfo

func (m *FindMaximumRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type FindMaximumResponse struct {
	Maximum              int32    `protobuf:"varint,1,opt,name=maximum,proto3" json:"maximum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaximumResponse) Reset()         { *m = FindMaximumResponse{} }
func (m *FindMaximumResponse) String() string { return proto.CompactTextString(m) }
func (*FindMaximumResponse) ProtoMessage()    {}
func (*FindMaximumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_49c82b0727b06490, []int{1}
}
func (m *FindMaximumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaximumResponse.Unmarshal(m, b)
}
func (m *FindMaximumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaximumResponse.Marshal(b, m, deterministic)
}
func (m *FindMaximumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaximumResponse.Merge(m, src)
}
func (m *FindMaximumResponse) XXX_Size() int {
	return xxx_messageInfo_FindMaximumResponse.Size(m)
}
func (m *FindMaximumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaximumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaximumResponse proto.InternalMessageInfo

func (m *FindMaximumResponse) GetMaximum() int32 {
	if m != nil {
		return m.Maximum
	}
	return 0
}

func init() {
	proto.RegisterType((*FindMaximumRequest)(nil), "StreamClient.FindMaximumRequest")
	proto.RegisterType((*FindMaximumResponse)(nil), "StreamClient.FindMaximumResponse")
}

func init() {
	proto.RegisterFile("FindMaxium/FindMaximum_proto/FindMaximum.proto", fileDescriptor_49c82b0727b06490)
}

var fileDescriptor_49c82b0727b06490 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x73, 0xcb, 0xcc, 0x4b,
	0xf1, 0x4d, 0xac, 0xc8, 0x2c, 0xcd, 0xd5, 0x87, 0x31, 0x73, 0x4b, 0x73, 0xe3, 0x0b, 0x8a, 0xf2,
	0x4b, 0xf2, 0x91, 0x45, 0xf4, 0xc0, 0x22, 0x42, 0x3c, 0xc1, 0x25, 0x45, 0xa9, 0x89, 0xb9, 0xce,
	0x39, 0x99, 0xa9, 0x79, 0x25, 0x4a, 0x3a, 0x5c, 0x42, 0x48, 0x4a, 0x82, 0x52, 0x0b, 0x4b, 0x53,
	0x8b, 0x4b, 0x84, 0xc4, 0xb8, 0xd8, 0xf2, 0x4a, 0x73, 0x93, 0x52, 0x8b, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x58, 0x83, 0xa0, 0x3c, 0x25, 0x7d, 0x2e, 0x61, 0x14, 0xd5, 0xc5, 0x05, 0xf9, 0x79, 0xc5,
	0xa9, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0x10, 0x21, 0xa8, 0x7a, 0x18, 0xd7, 0x28, 0x9b, 0x8b, 0xdf,
	0x39, 0x3f, 0xb7, 0xa0, 0xb4, 0x24, 0x35, 0xc5, 0xb1, 0x2c, 0xb5, 0x28, 0x31, 0x3d, 0x55, 0x28,
	0x82, 0x8b, 0x1b, 0xc9, 0x0c, 0x21, 0x05, 0x3d, 0x64, 0xf7, 0xe8, 0x61, 0x3a, 0x46, 0x4a, 0x11,
	0x8f, 0x0a, 0x88, 0x03, 0x94, 0x18, 0x34, 0x18, 0x0d, 0x18, 0x9d, 0x04, 0xa2, 0xf8, 0x90, 0xfd,
	0x1b, 0x90, 0x94, 0xc4, 0x06, 0xf6, 0xb2, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x12, 0x9b,
	0x4c, 0x24, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ComputedAverageClient is the client API for ComputedAverage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ComputedAverageClient interface {
	// Bidrectional streamming
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (ComputedAverage_FindMaximumClient, error)
}

type computedAverageClient struct {
	cc *grpc.ClientConn
}

func NewComputedAverageClient(cc *grpc.ClientConn) ComputedAverageClient {
	return &computedAverageClient{cc}
}

func (c *computedAverageClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (ComputedAverage_FindMaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ComputedAverage_serviceDesc.Streams[0], "/StreamClient.ComputedAverage/FindMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &computedAverageFindMaximumClient{stream}
	return x, nil
}

type ComputedAverage_FindMaximumClient interface {
	Send(*FindMaximumRequest) error
	Recv() (*FindMaximumResponse, error)
	grpc.ClientStream
}

type computedAverageFindMaximumClient struct {
	grpc.ClientStream
}

func (x *computedAverageFindMaximumClient) Send(m *FindMaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *computedAverageFindMaximumClient) Recv() (*FindMaximumResponse, error) {
	m := new(FindMaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ComputedAverageServer is the server API for ComputedAverage service.
type ComputedAverageServer interface {
	// Bidrectional streamming
	FindMaximum(ComputedAverage_FindMaximumServer) error
}

// UnimplementedComputedAverageServer can be embedded to have forward compatible implementations.
type UnimplementedComputedAverageServer struct {
}

func (*UnimplementedComputedAverageServer) FindMaximum(srv ComputedAverage_FindMaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaximum not implemented")
}

func RegisterComputedAverageServer(s *grpc.Server, srv ComputedAverageServer) {
	s.RegisterService(&_ComputedAverage_serviceDesc, srv)
}

func _ComputedAverage_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ComputedAverageServer).FindMaximum(&computedAverageFindMaximumServer{stream})
}

type ComputedAverage_FindMaximumServer interface {
	Send(*FindMaximumResponse) error
	Recv() (*FindMaximumRequest, error)
	grpc.ServerStream
}

type computedAverageFindMaximumServer struct {
	grpc.ServerStream
}

func (x *computedAverageFindMaximumServer) Send(m *FindMaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *computedAverageFindMaximumServer) Recv() (*FindMaximumRequest, error) {
	m := new(FindMaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ComputedAverage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "StreamClient.ComputedAverage",
	HandlerType: (*ComputedAverageServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindMaximum",
			Handler:       _ComputedAverage_FindMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "FindMaxium/FindMaximum_proto/FindMaximum.proto",
}
