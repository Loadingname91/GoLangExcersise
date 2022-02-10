// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: calculateSqaureRoot/calculate_proto/calculate_sqRoot.proto

package calculate_proto

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

type SquareRootRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRootRequest) Reset()         { *m = SquareRootRequest{} }
func (m *SquareRootRequest) String() string { return proto.CompactTextString(m) }
func (*SquareRootRequest) ProtoMessage()    {}
func (*SquareRootRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_986cddc1688b4d45, []int{0}
}
func (m *SquareRootRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRootRequest.Unmarshal(m, b)
}
func (m *SquareRootRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRootRequest.Marshal(b, m, deterministic)
}
func (m *SquareRootRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRootRequest.Merge(m, src)
}
func (m *SquareRootRequest) XXX_Size() int {
	return xxx_messageInfo_SquareRootRequest.Size(m)
}
func (m *SquareRootRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRootRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRootRequest proto.InternalMessageInfo

func (m *SquareRootRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type SqaureRootResponse struct {
	Result               float64  `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SqaureRootResponse) Reset()         { *m = SqaureRootResponse{} }
func (m *SqaureRootResponse) String() string { return proto.CompactTextString(m) }
func (*SqaureRootResponse) ProtoMessage()    {}
func (*SqaureRootResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_986cddc1688b4d45, []int{1}
}
func (m *SqaureRootResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SqaureRootResponse.Unmarshal(m, b)
}
func (m *SqaureRootResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SqaureRootResponse.Marshal(b, m, deterministic)
}
func (m *SqaureRootResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SqaureRootResponse.Merge(m, src)
}
func (m *SqaureRootResponse) XXX_Size() int {
	return xxx_messageInfo_SqaureRootResponse.Size(m)
}
func (m *SqaureRootResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SqaureRootResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SqaureRootResponse proto.InternalMessageInfo

func (m *SqaureRootResponse) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*SquareRootRequest)(nil), "SqaureRootPackage.SquareRootRequest")
	proto.RegisterType((*SqaureRootResponse)(nil), "SqaureRootPackage.SqaureRootResponse")
}

func init() {
	proto.RegisterFile("calculateSqaureRoot/calculate_proto/calculate_sqRoot.proto", fileDescriptor_986cddc1688b4d45)
}

var fileDescriptor_986cddc1688b4d45 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4a, 0x4e, 0xcc, 0x49,
	0x2e, 0xcd, 0x49, 0x2c, 0x49, 0x0d, 0x2e, 0x4c, 0x2c, 0x2d, 0x4a, 0x0d, 0xca, 0xcf, 0x2f, 0xd1,
	0x87, 0x8b, 0xc5, 0x17, 0x14, 0xe5, 0x97, 0xe4, 0x23, 0xf1, 0x8b, 0x0b, 0x41, 0x0a, 0xf4, 0xc0,
	0xc2, 0x42, 0x82, 0x08, 0x2d, 0x01, 0x89, 0xc9, 0xd9, 0x89, 0xe9, 0xa9, 0x4a, 0xda, 0x5c, 0x82,
	0xc1, 0x85, 0xa5, 0x89, 0x10, 0xc1, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x31, 0x2e,
	0xb6, 0xbc, 0xd2, 0xdc, 0xa4, 0xd4, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x28, 0x4f,
	0x49, 0x87, 0x4b, 0x08, 0x61, 0x42, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0x48, 0x75,
	0x51, 0x6a, 0x71, 0x69, 0x4e, 0x09, 0x58, 0x35, 0x63, 0x10, 0x94, 0x67, 0x94, 0xcb, 0xc5, 0xe7,
	0x96, 0x99, 0x97, 0x82, 0x30, 0x5e, 0x28, 0x9a, 0x8b, 0x0b, 0xa1, 0x5f, 0x48, 0x45, 0x0f, 0xc3,
	0x39, 0x7a, 0x18, 0x6e, 0x91, 0x52, 0xc5, 0xaa, 0x0a, 0xdd, 0x11, 0x4a, 0x0c, 0x4e, 0x42, 0x51,
	0x02, 0xe8, 0xe1, 0x90, 0xc4, 0x06, 0xa6, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x56, 0xb4,
	0xce, 0x02, 0x35, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FindSquareRootClient is the client API for FindSquareRoot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FindSquareRootClient interface {
	// Unary
	// Error handling
	// This RPC will throw an expection if the sent number is negative
	// The error being sent is of the type INVALID_ARGUMENT
	SqaureRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SqaureRootResponse, error)
}

type findSquareRootClient struct {
	cc *grpc.ClientConn
}

func NewFindSquareRootClient(cc *grpc.ClientConn) FindSquareRootClient {
	return &findSquareRootClient{cc}
}

func (c *findSquareRootClient) SqaureRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SqaureRootResponse, error) {
	out := new(SqaureRootResponse)
	err := c.cc.Invoke(ctx, "/SqaureRootPackage.FindSquareRoot/SqaureRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FindSquareRootServer is the server API for FindSquareRoot service.
type FindSquareRootServer interface {
	// Unary
	// Error handling
	// This RPC will throw an expection if the sent number is negative
	// The error being sent is of the type INVALID_ARGUMENT
	SqaureRoot(context.Context, *SquareRootRequest) (*SqaureRootResponse, error)
}

// UnimplementedFindSquareRootServer can be embedded to have forward compatible implementations.
type UnimplementedFindSquareRootServer struct {
}

func (*UnimplementedFindSquareRootServer) SqaureRoot(ctx context.Context, req *SquareRootRequest) (*SqaureRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SqaureRoot not implemented")
}

func RegisterFindSquareRootServer(s *grpc.Server, srv FindSquareRootServer) {
	s.RegisterService(&_FindSquareRoot_serviceDesc, srv)
}

func _FindSquareRoot_SqaureRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FindSquareRootServer).SqaureRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SqaureRootPackage.FindSquareRoot/SqaureRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FindSquareRootServer).SqaureRoot(ctx, req.(*SquareRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FindSquareRoot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SqaureRootPackage.FindSquareRoot",
	HandlerType: (*FindSquareRootServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SqaureRoot",
			Handler:    _FindSquareRoot_SqaureRoot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculateSqaureRoot/calculate_proto/calculate_sqRoot.proto",
}
