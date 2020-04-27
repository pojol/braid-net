// Code generated by protoc-gen-go. DO NOT EDIT.
// source: route.proto

package brpc

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

type RouteReq struct {
	Nod                  string   `protobuf:"bytes,1,opt,name=nod,proto3" json:"nod,omitempty"`
	Service              string   `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	ReqBody              []byte   `protobuf:"bytes,3,opt,name=reqBody,proto3" json:"reqBody,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteReq) Reset()         { *m = RouteReq{} }
func (m *RouteReq) String() string { return proto.CompactTextString(m) }
func (*RouteReq) ProtoMessage()    {}
func (*RouteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0984d49a362b6b9f, []int{0}
}

func (m *RouteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteReq.Unmarshal(m, b)
}
func (m *RouteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteReq.Marshal(b, m, deterministic)
}
func (m *RouteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteReq.Merge(m, src)
}
func (m *RouteReq) XXX_Size() int {
	return xxx_messageInfo_RouteReq.Size(m)
}
func (m *RouteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteReq.DiscardUnknown(m)
}

var xxx_messageInfo_RouteReq proto.InternalMessageInfo

func (m *RouteReq) GetNod() string {
	if m != nil {
		return m.Nod
	}
	return ""
}

func (m *RouteReq) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *RouteReq) GetReqBody() []byte {
	if m != nil {
		return m.ReqBody
	}
	return nil
}

type Header struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val                  string   `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_0984d49a362b6b9f, []int{1}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Header) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

type RouteRes struct {
	ResBody              []byte    `protobuf:"bytes,1,opt,name=resBody,proto3" json:"resBody,omitempty"`
	Headers              []*Header `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RouteRes) Reset()         { *m = RouteRes{} }
func (m *RouteRes) String() string { return proto.CompactTextString(m) }
func (*RouteRes) ProtoMessage()    {}
func (*RouteRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_0984d49a362b6b9f, []int{2}
}

func (m *RouteRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteRes.Unmarshal(m, b)
}
func (m *RouteRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteRes.Marshal(b, m, deterministic)
}
func (m *RouteRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteRes.Merge(m, src)
}
func (m *RouteRes) XXX_Size() int {
	return xxx_messageInfo_RouteRes.Size(m)
}
func (m *RouteRes) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteRes.DiscardUnknown(m)
}

var xxx_messageInfo_RouteRes proto.InternalMessageInfo

func (m *RouteRes) GetResBody() []byte {
	if m != nil {
		return m.ResBody
	}
	return nil
}

func (m *RouteRes) GetHeaders() []*Header {
	if m != nil {
		return m.Headers
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteReq)(nil), "brpc.routeReq")
	proto.RegisterType((*Header)(nil), "brpc.Header")
	proto.RegisterType((*RouteRes)(nil), "brpc.routeRes")
}

func init() { proto.RegisterFile("route.proto", fileDescriptor_0984d49a362b6b9f) }

var fileDescriptor_0984d49a362b6b9f = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xca, 0x2f, 0x2d,
	0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x2a, 0x2a, 0x48, 0x56, 0x0a, 0xe0,
	0xe2, 0x00, 0x0b, 0x06, 0xa5, 0x16, 0x0a, 0x09, 0x70, 0x31, 0xe7, 0xe5, 0xa7, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x42, 0x12, 0x5c, 0xec, 0xc5, 0xa9, 0x45, 0x65, 0x99, 0xc9,
	0xa9, 0x12, 0x4c, 0x60, 0x51, 0x18, 0x17, 0x24, 0x53, 0x94, 0x5a, 0xe8, 0x94, 0x9f, 0x52, 0x29,
	0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x13, 0x04, 0xe3, 0x2a, 0xe9, 0x70, 0xb1, 0x79, 0xa4, 0x26, 0xa6,
	0xa4, 0x16, 0x81, 0xcc, 0xcb, 0x4e, 0xad, 0x84, 0x99, 0x97, 0x9d, 0x5a, 0x09, 0x12, 0x29, 0x4b,
	0xcc, 0x81, 0x9a, 0x05, 0x62, 0x2a, 0xf9, 0xc0, 0xed, 0x2f, 0x86, 0x98, 0x59, 0x0c, 0x36, 0x93,
	0x11, 0x66, 0x26, 0x98, 0x2b, 0xa4, 0xc6, 0xc5, 0x9e, 0x01, 0x36, 0xb3, 0x58, 0x82, 0x49, 0x81,
	0x59, 0x83, 0xdb, 0x88, 0x47, 0x0f, 0xe4, 0x7a, 0x3d, 0x88, 0x45, 0x41, 0x30, 0x49, 0x23, 0x33,
	0x2e, 0xf6, 0xf4, 0xc4, 0x92, 0xd4, 0xf2, 0xc4, 0x4a, 0x21, 0x6d, 0x2e, 0x76, 0x90, 0xc1, 0x99,
	0x79, 0xe9, 0x42, 0x7c, 0x10, 0xc5, 0x30, 0x7f, 0x4a, 0xa1, 0xf2, 0x8b, 0x95, 0x18, 0x92, 0xd8,
	0xc0, 0x41, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x09, 0x11, 0x94, 0x18, 0x21, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayClient interface {
	Routing(ctx context.Context, in *RouteReq, opts ...grpc.CallOption) (*RouteRes, error)
}

type gatewayClient struct {
	cc *grpc.ClientConn
}

func NewGatewayClient(cc *grpc.ClientConn) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) Routing(ctx context.Context, in *RouteReq, opts ...grpc.CallOption) (*RouteRes, error) {
	out := new(RouteRes)
	err := c.cc.Invoke(ctx, "/brpc.gateway/routing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
type GatewayServer interface {
	Routing(context.Context, *RouteReq) (*RouteRes, error)
}

// UnimplementedGatewayServer can be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (*UnimplementedGatewayServer) Routing(ctx context.Context, req *RouteReq) (*RouteRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Routing not implemented")
}

func RegisterGatewayServer(s *grpc.Server, srv GatewayServer) {
	s.RegisterService(&_Gateway_serviceDesc, srv)
}

func _Gateway_Routing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Routing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brpc.gateway/Routing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Routing(ctx, req.(*RouteReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "brpc.gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "routing",
			Handler:    _Gateway_Routing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "route.proto",
}
