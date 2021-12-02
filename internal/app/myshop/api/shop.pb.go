// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shop.proto

package api

import (
	context "context"
	fmt "fmt"
	gogoproto "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

var E_GoprotoEnumPrefix = gogoproto.E_GoprotoEnumPrefix

var E_GoprotoEnumStringer = gogoproto.E_GoprotoEnumStringer

var E_EnumStringer = gogoproto.E_EnumStringer

var E_EnumCustomname = gogoproto.E_EnumCustomname

var E_Enumdecl = gogoproto.E_Enumdecl

var E_EnumvalueCustomname = gogoproto.E_EnumvalueCustomname

var E_GoprotoGettersAll = gogoproto.E_GoprotoGettersAll

var E_GoprotoEnumPrefixAll = gogoproto.E_GoprotoEnumPrefixAll

var E_GoprotoStringerAll = gogoproto.E_GoprotoStringerAll

var E_VerboseEqualAll = gogoproto.E_VerboseEqualAll

var E_FaceAll = gogoproto.E_FaceAll

var E_GostringAll = gogoproto.E_GostringAll

var E_PopulateAll = gogoproto.E_PopulateAll

var E_StringerAll = gogoproto.E_StringerAll

var E_OnlyoneAll = gogoproto.E_OnlyoneAll

var E_EqualAll = gogoproto.E_EqualAll

var E_DescriptionAll = gogoproto.E_DescriptionAll

var E_TestgenAll = gogoproto.E_TestgenAll

var E_BenchgenAll = gogoproto.E_BenchgenAll

var E_MarshalerAll = gogoproto.E_MarshalerAll

var E_UnmarshalerAll = gogoproto.E_UnmarshalerAll

var E_StableMarshalerAll = gogoproto.E_StableMarshalerAll

var E_SizerAll = gogoproto.E_SizerAll

var E_GoprotoEnumStringerAll = gogoproto.E_GoprotoEnumStringerAll

var E_EnumStringerAll = gogoproto.E_EnumStringerAll

var E_UnsafeMarshalerAll = gogoproto.E_UnsafeMarshalerAll

var E_UnsafeUnmarshalerAll = gogoproto.E_UnsafeUnmarshalerAll

var E_GoprotoExtensionsMapAll = gogoproto.E_GoprotoExtensionsMapAll

var E_GoprotoUnrecognizedAll = gogoproto.E_GoprotoUnrecognizedAll

var E_GogoprotoImport = gogoproto.E_GogoprotoImport

var E_ProtosizerAll = gogoproto.E_ProtosizerAll

var E_CompareAll = gogoproto.E_CompareAll

var E_TypedeclAll = gogoproto.E_TypedeclAll

var E_EnumdeclAll = gogoproto.E_EnumdeclAll

var E_GoprotoRegistration = gogoproto.E_GoprotoRegistration

var E_MessagenameAll = gogoproto.E_MessagenameAll

var E_GoprotoSizecacheAll = gogoproto.E_GoprotoSizecacheAll

var E_GoprotoUnkeyedAll = gogoproto.E_GoprotoUnkeyedAll

var E_GoprotoGetters = gogoproto.E_GoprotoGetters

var E_GoprotoStringer = gogoproto.E_GoprotoStringer

var E_VerboseEqual = gogoproto.E_VerboseEqual

var E_Face = gogoproto.E_Face

var E_Gostring = gogoproto.E_Gostring

var E_Populate = gogoproto.E_Populate

var E_Stringer = gogoproto.E_Stringer

var E_Onlyone = gogoproto.E_Onlyone

var E_Equal = gogoproto.E_Equal

var E_Description = gogoproto.E_Description

var E_Testgen = gogoproto.E_Testgen

var E_Benchgen = gogoproto.E_Benchgen

var E_Marshaler = gogoproto.E_Marshaler

var E_Unmarshaler = gogoproto.E_Unmarshaler

var E_StableMarshaler = gogoproto.E_StableMarshaler

var E_Sizer = gogoproto.E_Sizer

var E_UnsafeMarshaler = gogoproto.E_UnsafeMarshaler

var E_UnsafeUnmarshaler = gogoproto.E_UnsafeUnmarshaler

var E_GoprotoExtensionsMap = gogoproto.E_GoprotoExtensionsMap

var E_GoprotoUnrecognized = gogoproto.E_GoprotoUnrecognized

var E_Protosizer = gogoproto.E_Protosizer

var E_Compare = gogoproto.E_Compare

var E_Typedecl = gogoproto.E_Typedecl

var E_Messagename = gogoproto.E_Messagename

var E_GoprotoSizecache = gogoproto.E_GoprotoSizecache

var E_GoprotoUnkeyed = gogoproto.E_GoprotoUnkeyed

var E_Nullable = gogoproto.E_Nullable

var E_Embed = gogoproto.E_Embed

var E_Customtype = gogoproto.E_Customtype

var E_Customname = gogoproto.E_Customname

var E_Jsontag = gogoproto.E_Jsontag

var E_Moretags = gogoproto.E_Moretags

var E_Casttype = gogoproto.E_Casttype

var E_Castkey = gogoproto.E_Castkey

var E_Castvalue = gogoproto.E_Castvalue

var E_Stdtime = gogoproto.E_Stdtime

var E_Stdduration = gogoproto.E_Stdduration

var E_Wktpointer = gogoproto.E_Wktpointer

type ShopResponse struct {
	Code                 int64    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShopResponse) Reset()         { *m = ShopResponse{} }
func (m *ShopResponse) String() string { return proto.CompactTextString(m) }
func (*ShopResponse) ProtoMessage()    {}
func (*ShopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{0}
}
func (m *ShopResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ShopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ShopResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ShopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShopResponse.Merge(m, src)
}
func (m *ShopResponse) XXX_Size() int {
	return m.Size()
}
func (m *ShopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShopResponse proto.InternalMessageInfo

func (m *ShopResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ShopResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ShopRequest struct {
	Tid                  string   `protobuf:"bytes,1,opt,name=Tid,proto3" json:"Tid,omitempty"`
	RefundId             int64    `protobuf:"varint,2,opt,name=RefundId,proto3" json:"RefundId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShopRequest) Reset()         { *m = ShopRequest{} }
func (m *ShopRequest) String() string { return proto.CompactTextString(m) }
func (*ShopRequest) ProtoMessage()    {}
func (*ShopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{1}
}
func (m *ShopRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ShopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ShopRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ShopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShopRequest.Merge(m, src)
}
func (m *ShopRequest) XXX_Size() int {
	return m.Size()
}
func (m *ShopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShopRequest proto.InternalMessageInfo

func (m *ShopRequest) GetTid() string {
	if m != nil {
		return m.Tid
	}
	return ""
}

func (m *ShopRequest) GetRefundId() int64 {
	if m != nil {
		return m.RefundId
	}
	return 0
}

type ShopOrderResponse struct {
	Code                 int64    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Order                *Order   `protobuf:"bytes,3,opt,name=Order,proto3" json:"Order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShopOrderResponse) Reset()         { *m = ShopOrderResponse{} }
func (m *ShopOrderResponse) String() string { return proto.CompactTextString(m) }
func (*ShopOrderResponse) ProtoMessage()    {}
func (*ShopOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{2}
}
func (m *ShopOrderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ShopOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ShopOrderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ShopOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShopOrderResponse.Merge(m, src)
}
func (m *ShopOrderResponse) XXX_Size() int {
	return m.Size()
}
func (m *ShopOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShopOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShopOrderResponse proto.InternalMessageInfo

func (m *ShopOrderResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ShopOrderResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ShopOrderResponse) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type Order struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ShopId               int64    `protobuf:"varint,2,opt,name=ShopId,proto3" json:"ShopId,omitempty"`
	Tid                  string   `protobuf:"bytes,3,opt,name=Tid,proto3" json:"Tid,omitempty"`
	RefundId             int64    `protobuf:"varint,4,opt,name=RefundId,proto3" json:"RefundId,omitempty"`
	Type                 int64    `protobuf:"varint,5,opt,name=Type,proto3" json:"Type,omitempty"`
	CurrentVersion       int64    `protobuf:"varint,6,opt,name=CurrentVersion,proto3" json:"CurrentVersion,omitempty"`
	Content              string   `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{3}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Order.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return m.Size()
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Order) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *Order) GetTid() string {
	if m != nil {
		return m.Tid
	}
	return ""
}

func (m *Order) GetRefundId() int64 {
	if m != nil {
		return m.RefundId
	}
	return 0
}

func (m *Order) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Order) GetCurrentVersion() int64 {
	if m != nil {
		return m.CurrentVersion
	}
	return 0
}

func (m *Order) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*ShopResponse)(nil), "main.ShopResponse")
	proto.RegisterType((*ShopRequest)(nil), "main.ShopRequest")
	proto.RegisterType((*ShopOrderResponse)(nil), "main.ShopOrderResponse")
	proto.RegisterType((*Order)(nil), "main.Order")
}

func init() { proto.RegisterFile("shop.proto", fileDescriptor_0f3030369b20fd61) }

var fileDescriptor_0f3030369b20fd61 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4e, 0xc2, 0x30,
	0x18, 0xc7, 0x29, 0x1b, 0x43, 0x3e, 0x0c, 0x48, 0x0f, 0xda, 0xec, 0xb0, 0xe0, 0x0e, 0x86, 0x8b,
	0x23, 0x41, 0x6f, 0xc4, 0x83, 0x72, 0x30, 0x24, 0x1a, 0xcd, 0x24, 0x1e, 0x8c, 0x17, 0xa0, 0x65,
	0xec, 0x40, 0x3b, 0xb7, 0xee, 0xc0, 0x9b, 0xf8, 0x1e, 0xbe, 0x84, 0x47, 0x1f, 0xc1, 0xe0, 0x8b,
	0x98, 0xb6, 0x80, 0xcc, 0x78, 0xf1, 0xf6, 0xff, 0xf7, 0xeb, 0xef, 0xfb, 0xfe, 0xfd, 0x52, 0x80,
	0x6c, 0x2e, 0x92, 0x20, 0x49, 0x85, 0x14, 0xd8, 0x5e, 0x8c, 0x63, 0xee, 0x9e, 0x46, 0xb1, 0x9c,
	0xe7, 0x93, 0x60, 0x2a, 0x16, 0xdd, 0x48, 0x44, 0xa2, 0xab, 0x8b, 0x93, 0x7c, 0xa6, 0x9d, 0x36,
	0x5a, 0x19, 0xc8, 0x3f, 0x87, 0xfd, 0x87, 0xb9, 0x48, 0x42, 0x96, 0x25, 0x82, 0x67, 0x0c, 0x63,
	0xb0, 0x07, 0x82, 0x32, 0x82, 0xda, 0xa8, 0x63, 0x85, 0x5a, 0xe3, 0x03, 0xb0, 0x6e, 0xb3, 0x88,
	0x94, 0xdb, 0xa8, 0x53, 0x0b, 0x95, 0xf4, 0xfb, 0x50, 0x37, 0xd4, 0x4b, 0xce, 0x32, 0xa9, 0x2e,
	0x8c, 0x62, 0xaa, 0x99, 0x5a, 0xa8, 0x24, 0x76, 0x61, 0x2f, 0x64, 0xb3, 0x9c, 0xd3, 0x21, 0xd5,
	0x9c, 0x15, 0x6e, 0xbd, 0xff, 0x0c, 0x2d, 0x05, 0xdf, 0xa5, 0x94, 0xa5, 0xff, 0x9b, 0x8b, 0x8f,
	0xa1, 0xa2, 0x31, 0x62, 0xb5, 0x51, 0xa7, 0xde, 0xab, 0x07, 0xea, 0xc9, 0x81, 0xe9, 0x64, 0x2a,
	0xfe, 0x1b, 0x5a, 0xdf, 0xc1, 0x0d, 0x28, 0x0f, 0xe9, 0xba, 0x61, 0x79, 0x48, 0xf1, 0x21, 0x38,
	0x6a, 0xee, 0x36, 0xd1, 0xda, 0x6d, 0xd2, 0x5b, 0x7f, 0xa7, 0xb7, 0x8b, 0xe9, 0x55, 0xd0, 0xd1,
	0x32, 0x61, 0xa4, 0x62, 0x82, 0x2a, 0x8d, 0x4f, 0xa0, 0x31, 0xc8, 0xd3, 0x94, 0x71, 0xf9, 0xc8,
	0xd2, 0x2c, 0x16, 0x9c, 0x38, 0xba, 0xfa, 0xeb, 0x14, 0x13, 0xa8, 0x4e, 0x05, 0x97, 0x8c, 0x4b,
	0x52, 0xd5, 0xd3, 0x36, 0xb6, 0xb7, 0x04, 0x5b, 0xa5, 0xc1, 0x17, 0xd0, 0xbc, 0x66, 0x72, 0xbb,
	0x9e, 0x1b, 0x11, 0xe1, 0x96, 0x79, 0xe4, 0xce, 0xbe, 0xdd, 0xa3, 0x9f, 0xa3, 0xe2, 0x16, 0x7b,
	0xd0, 0xbc, 0xa4, 0xb4, 0x80, 0xef, 0xee, 0xc8, 0xc5, 0xbb, 0xbd, 0x0c, 0x73, 0x45, 0xde, 0x57,
	0x1e, 0xfa, 0x58, 0x79, 0xe8, 0x73, 0xe5, 0xa1, 0xd7, 0x2f, 0xaf, 0xf4, 0xe4, 0x04, 0xdd, 0xfe,
	0x38, 0x89, 0xef, 0x4b, 0x13, 0x47, 0x7f, 0x92, 0xb3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6c,
	0xf6, 0xad, 0x71, 0x67, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShopClient is the client API for Shop service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShopClient interface {
	GetShopOrderLog(ctx context.Context, in *ShopRequest, opts ...grpc.CallOption) (*ShopOrderResponse, error)
	AddShopOrderLog(ctx context.Context, in *Order, opts ...grpc.CallOption) (*ShopResponse, error)
}

type shopClient struct {
	cc *grpc.ClientConn
}

func NewShopClient(cc *grpc.ClientConn) ShopClient {
	return &shopClient{cc}
}

func (c *shopClient) GetShopOrderLog(ctx context.Context, in *ShopRequest, opts ...grpc.CallOption) (*ShopOrderResponse, error) {
	out := new(ShopOrderResponse)
	err := c.cc.Invoke(ctx, "/main.Shop/GetShopOrderLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) AddShopOrderLog(ctx context.Context, in *Order, opts ...grpc.CallOption) (*ShopResponse, error) {
	out := new(ShopResponse)
	err := c.cc.Invoke(ctx, "/main.Shop/AddShopOrderLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServer is the server API for Shop service.
type ShopServer interface {
	GetShopOrderLog(context.Context, *ShopRequest) (*ShopOrderResponse, error)
	AddShopOrderLog(context.Context, *Order) (*ShopResponse, error)
}

// UnimplementedShopServer can be embedded to have forward compatible implementations.
type UnimplementedShopServer struct {
}

func (*UnimplementedShopServer) GetShopOrderLog(ctx context.Context, req *ShopRequest) (*ShopOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopOrderLog not implemented")
}
func (*UnimplementedShopServer) AddShopOrderLog(ctx context.Context, req *Order) (*ShopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddShopOrderLog not implemented")
}

func RegisterShopServer(s *grpc.Server, srv ShopServer) {
	s.RegisterService(&_Shop_serviceDesc, srv)
}

func _Shop_GetShopOrderLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).GetShopOrderLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Shop/GetShopOrderLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).GetShopOrderLog(ctx, req.(*ShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shop_AddShopOrderLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).AddShopOrderLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Shop/AddShopOrderLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).AddShopOrderLog(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

var _Shop_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.Shop",
	HandlerType: (*ShopServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShopOrderLog",
			Handler:    _Shop_GetShopOrderLog_Handler,
		},
		{
			MethodName: "AddShopOrderLog",
			Handler:    _Shop_AddShopOrderLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop.proto",
}

func (m *ShopResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShopResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ShopResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintShop(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintShop(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ShopRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShopRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ShopRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.RefundId != 0 {
		i = encodeVarintShop(dAtA, i, uint64(m.RefundId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Tid) > 0 {
		i -= len(m.Tid)
		copy(dAtA[i:], m.Tid)
		i = encodeVarintShop(dAtA, i, uint64(len(m.Tid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ShopOrderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShopOrderResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ShopOrderResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Order != nil {
		{
			size, err := m.Order.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintShop(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintShop(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintShop(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Order) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Order) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Order) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintShop(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x3a
	}
	if m.CurrentVersion != 0 {
		i = encodeVarintShop(dAtA, i, uint64(m.CurrentVersion))
		i--
		dAtA[i] = 0x30
	}
	if m.Type != 0 {
		i = encodeVarintShop(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x28
	}
	if m.RefundId != 0 {
		i = encodeVarintShop(dAtA, i, uint64(m.RefundId))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Tid) > 0 {
		i -= len(m.Tid)
		copy(dAtA[i:], m.Tid)
		i = encodeVarintShop(dAtA, i, uint64(len(m.Tid)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ShopId != 0 {
		i = encodeVarintShop(dAtA, i, uint64(m.ShopId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintShop(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintShop(dAtA []byte, offset int, v uint64) int {
	offset -= sovShop(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ShopResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovShop(uint64(m.Code))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovShop(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ShopRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Tid)
	if l > 0 {
		n += 1 + l + sovShop(uint64(l))
	}
	if m.RefundId != 0 {
		n += 1 + sovShop(uint64(m.RefundId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ShopOrderResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovShop(uint64(m.Code))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovShop(uint64(l))
	}
	if m.Order != nil {
		l = m.Order.Size()
		n += 1 + l + sovShop(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Order) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovShop(uint64(m.Id))
	}
	if m.ShopId != 0 {
		n += 1 + sovShop(uint64(m.ShopId))
	}
	l = len(m.Tid)
	if l > 0 {
		n += 1 + l + sovShop(uint64(l))
	}
	if m.RefundId != 0 {
		n += 1 + sovShop(uint64(m.RefundId))
	}
	if m.Type != 0 {
		n += 1 + sovShop(uint64(m.Type))
	}
	if m.CurrentVersion != 0 {
		n += 1 + sovShop(uint64(m.CurrentVersion))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovShop(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovShop(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozShop(x uint64) (n int) {
	return sovShop(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ShopResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShop
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ShopResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShopResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShop
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ShopRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShop
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ShopRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShopRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefundId", wireType)
			}
			m.RefundId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RefundId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShop
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ShopOrderResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShop
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ShopOrderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShopOrderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Order", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthShop
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Order == nil {
				m.Order = &Order{}
			}
			if err := m.Order.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShop
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Order) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShop
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Order: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Order: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShopId", wireType)
			}
			m.ShopId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShopId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefundId", wireType)
			}
			m.RefundId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RefundId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentVersion", wireType)
			}
			m.CurrentVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentVersion |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShop
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipShop(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowShop
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowShop
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowShop
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthShop
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupShop
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthShop
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthShop        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowShop          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupShop = fmt.Errorf("proto: unexpected end of group")
)
