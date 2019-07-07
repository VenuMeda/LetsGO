// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package gRPC

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

//The user type message that we will use to create the user in [CreateUser] function+
//and the response with User information in LoginUser
type UserMessage struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserMessage) Reset()         { *m = UserMessage{} }
func (m *UserMessage) String() string { return proto.CompactTextString(m) }
func (*UserMessage) ProtoMessage()    {}
func (*UserMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

func (m *UserMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserMessage.Unmarshal(m, b)
}
func (m *UserMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserMessage.Marshal(b, m, deterministic)
}
func (m *UserMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserMessage.Merge(m, src)
}
func (m *UserMessage) XXX_Size() int {
	return xxx_messageInfo_UserMessage.Size(m)
}
func (m *UserMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_UserMessage.DiscardUnknown(m)
}

var xxx_messageInfo_UserMessage proto.InternalMessageInfo

func (m *UserMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*UserMessage)(nil), "gRPC.UserMessage")
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x0f, 0x0a, 0x70, 0x56, 0x52,
	0xe7, 0xe2, 0x0e, 0x2d, 0x4e, 0x2d, 0xf2, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x92, 0xe0,
	0x62, 0xcf, 0x85, 0x30, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x23, 0x5f, 0x2e,
	0x41, 0xa8, 0x22, 0xdf, 0xc4, 0xbc, 0xc4, 0xf4, 0xd4, 0xdc, 0xd4, 0xbc, 0x12, 0x21, 0x0b, 0x2e,
	0xbe, 0x80, 0xa2, 0xfc, 0xe4, 0xd4, 0xe2, 0x62, 0x98, 0x01, 0x82, 0x7a, 0x20, 0x63, 0xf5, 0x90,
	0xcc, 0x94, 0xc2, 0x14, 0x52, 0x62, 0x48, 0x62, 0x03, 0x3b, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0xdf, 0x59, 0x00, 0x0f, 0x94, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessageManagementClient is the client API for MessageManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageManagementClient interface {
	ProcessMessage(ctx context.Context, in *UserMessage, opts ...grpc.CallOption) (*UserMessage, error)
}

type messageManagementClient struct {
	cc *grpc.ClientConn
}

func NewMessageManagementClient(cc *grpc.ClientConn) MessageManagementClient {
	return &messageManagementClient{cc}
}

func (c *messageManagementClient) ProcessMessage(ctx context.Context, in *UserMessage, opts ...grpc.CallOption) (*UserMessage, error) {
	out := new(UserMessage)
	err := c.cc.Invoke(ctx, "/gRPC.MessageManagement/ProcessMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageManagementServer is the server API for MessageManagement service.
type MessageManagementServer interface {
	ProcessMessage(context.Context, *UserMessage) (*UserMessage, error)
}

func RegisterMessageManagementServer(s *grpc.Server, srv MessageManagementServer) {
	s.RegisterService(&_MessageManagement_serviceDesc, srv)
}

func _MessageManagement_ProcessMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageManagementServer).ProcessMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gRPC.MessageManagement/ProcessMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageManagementServer).ProcessMessage(ctx, req.(*UserMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageManagement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gRPC.MessageManagement",
	HandlerType: (*MessageManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessMessage",
			Handler:    _MessageManagement_ProcessMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}