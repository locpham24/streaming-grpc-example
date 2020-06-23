// Code generated by protoc-gen-go. DO NOT EDIT.
// source: number.proto

package proto

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

type InputNumber struct {
	Num                  int64    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InputNumber) Reset()         { *m = InputNumber{} }
func (m *InputNumber) String() string { return proto.CompactTextString(m) }
func (*InputNumber) ProtoMessage()    {}
func (*InputNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6c6c2f2b9341206, []int{0}
}

func (m *InputNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputNumber.Unmarshal(m, b)
}
func (m *InputNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputNumber.Marshal(b, m, deterministic)
}
func (m *InputNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputNumber.Merge(m, src)
}
func (m *InputNumber) XXX_Size() int {
	return xxx_messageInfo_InputNumber.Size(m)
}
func (m *InputNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_InputNumber.DiscardUnknown(m)
}

var xxx_messageInfo_InputNumber proto.InternalMessageInfo

func (m *InputNumber) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

type OutputNumber struct {
	Num                  int64    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutputNumber) Reset()         { *m = OutputNumber{} }
func (m *OutputNumber) String() string { return proto.CompactTextString(m) }
func (*OutputNumber) ProtoMessage()    {}
func (*OutputNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6c6c2f2b9341206, []int{1}
}

func (m *OutputNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputNumber.Unmarshal(m, b)
}
func (m *OutputNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputNumber.Marshal(b, m, deterministic)
}
func (m *OutputNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputNumber.Merge(m, src)
}
func (m *OutputNumber) XXX_Size() int {
	return xxx_messageInfo_OutputNumber.Size(m)
}
func (m *OutputNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputNumber.DiscardUnknown(m)
}

var xxx_messageInfo_OutputNumber proto.InternalMessageInfo

func (m *OutputNumber) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

type HelloOutput struct {
	Greeting             string   `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloOutput) Reset()         { *m = HelloOutput{} }
func (m *HelloOutput) String() string { return proto.CompactTextString(m) }
func (*HelloOutput) ProtoMessage()    {}
func (*HelloOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6c6c2f2b9341206, []int{2}
}

func (m *HelloOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloOutput.Unmarshal(m, b)
}
func (m *HelloOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloOutput.Marshal(b, m, deterministic)
}
func (m *HelloOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloOutput.Merge(m, src)
}
func (m *HelloOutput) XXX_Size() int {
	return xxx_messageInfo_HelloOutput.Size(m)
}
func (m *HelloOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloOutput.DiscardUnknown(m)
}

var xxx_messageInfo_HelloOutput proto.InternalMessageInfo

func (m *HelloOutput) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*InputNumber)(nil), "proto.InputNumber")
	proto.RegisterType((*OutputNumber)(nil), "proto.OutputNumber")
	proto.RegisterType((*HelloOutput)(nil), "proto.HelloOutput")
}

func init() {
	proto.RegisterFile("number.proto", fileDescriptor_e6c6c2f2b9341206)
}

var fileDescriptor_e6c6c2f2b9341206 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x2b, 0xcd, 0x4d,
	0x4a, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xf2, 0x5c, 0xdc,
	0x9e, 0x79, 0x05, 0xa5, 0x25, 0x7e, 0x60, 0x39, 0x21, 0x01, 0x2e, 0xe6, 0xbc, 0xd2, 0x5c, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x10, 0x53, 0x49, 0x81, 0x8b, 0xc7, 0xbf, 0xb4, 0x04, 0x9f,
	0x0a, 0x4d, 0x2e, 0x6e, 0x8f, 0xd4, 0x9c, 0x9c, 0x7c, 0x88, 0x32, 0x21, 0x29, 0x2e, 0x8e, 0xf4,
	0xa2, 0xd4, 0xd4, 0x92, 0xcc, 0xbc, 0x74, 0xb0, 0x2a, 0xce, 0x20, 0x38, 0xdf, 0xa8, 0x9e, 0x8b,
	0x0d, 0x6a, 0x8c, 0x39, 0x54, 0x13, 0x94, 0x2b, 0x04, 0x71, 0x95, 0x1e, 0x92, 0x5b, 0xa4, 0x60,
	0x62, 0x48, 0x86, 0x2b, 0x31, 0x08, 0xd9, 0x70, 0xf1, 0xfa, 0x96, 0xe6, 0x94, 0x64, 0x16, 0xe4,
	0xa4, 0xfa, 0xa7, 0x85, 0x94, 0xe7, 0x63, 0xd5, 0x2a, 0x0c, 0x15, 0x43, 0x76, 0xb9, 0x12, 0x83,
	0x01, 0x63, 0x12, 0x1b, 0x58, 0xdc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xb5, 0xa3, 0xff, 0x15,
	0x0c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NumberClient is the client API for Number service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NumberClient interface {
	HelloNumber(ctx context.Context, in *InputNumber, opts ...grpc.CallOption) (*HelloOutput, error)
	MultipleOfTwo(ctx context.Context, in *InputNumber, opts ...grpc.CallOption) (Number_MultipleOfTwoClient, error)
}

type numberClient struct {
	cc grpc.ClientConnInterface
}

func NewNumberClient(cc grpc.ClientConnInterface) NumberClient {
	return &numberClient{cc}
}

func (c *numberClient) HelloNumber(ctx context.Context, in *InputNumber, opts ...grpc.CallOption) (*HelloOutput, error) {
	out := new(HelloOutput)
	err := c.cc.Invoke(ctx, "/proto.Number/HelloNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *numberClient) MultipleOfTwo(ctx context.Context, in *InputNumber, opts ...grpc.CallOption) (Number_MultipleOfTwoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Number_serviceDesc.Streams[0], "/proto.Number/MultipleOfTwo", opts...)
	if err != nil {
		return nil, err
	}
	x := &numberMultipleOfTwoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Number_MultipleOfTwoClient interface {
	Recv() (*OutputNumber, error)
	grpc.ClientStream
}

type numberMultipleOfTwoClient struct {
	grpc.ClientStream
}

func (x *numberMultipleOfTwoClient) Recv() (*OutputNumber, error) {
	m := new(OutputNumber)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NumberServer is the server API for Number service.
type NumberServer interface {
	HelloNumber(context.Context, *InputNumber) (*HelloOutput, error)
	MultipleOfTwo(*InputNumber, Number_MultipleOfTwoServer) error
}

// UnimplementedNumberServer can be embedded to have forward compatible implementations.
type UnimplementedNumberServer struct {
}

func (*UnimplementedNumberServer) HelloNumber(ctx context.Context, req *InputNumber) (*HelloOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloNumber not implemented")
}
func (*UnimplementedNumberServer) MultipleOfTwo(req *InputNumber, srv Number_MultipleOfTwoServer) error {
	return status.Errorf(codes.Unimplemented, "method MultipleOfTwo not implemented")
}

func RegisterNumberServer(s *grpc.Server, srv NumberServer) {
	s.RegisterService(&_Number_serviceDesc, srv)
}

func _Number_HelloNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputNumber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NumberServer).HelloNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Number/HelloNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NumberServer).HelloNumber(ctx, req.(*InputNumber))
	}
	return interceptor(ctx, in, info, handler)
}

func _Number_MultipleOfTwo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InputNumber)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NumberServer).MultipleOfTwo(m, &numberMultipleOfTwoServer{stream})
}

type Number_MultipleOfTwoServer interface {
	Send(*OutputNumber) error
	grpc.ServerStream
}

type numberMultipleOfTwoServer struct {
	grpc.ServerStream
}

func (x *numberMultipleOfTwoServer) Send(m *OutputNumber) error {
	return x.ServerStream.SendMsg(m)
}

var _Number_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Number",
	HandlerType: (*NumberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloNumber",
			Handler:    _Number_HelloNumber_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MultipleOfTwo",
			Handler:       _Number_MultipleOfTwo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "number.proto",
}
