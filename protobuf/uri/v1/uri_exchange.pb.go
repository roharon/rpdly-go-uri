// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.9.1
// source: protobuf/uri/v1/uri_exchange.proto

package v1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_uri_v1_uri_exchange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_uri_v1_uri_exchange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_protobuf_uri_v1_uri_exchange_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_uri_v1_uri_exchange_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_uri_v1_uri_exchange_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_protobuf_uri_v1_uri_exchange_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

var File_protobuf_uri_v1_uri_exchange_proto protoreflect.FileDescriptor

var file_protobuf_uri_v1_uri_exchange_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x75, 0x72, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x72, 0x69, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x69, 0x22, 0x1c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x69, 0x32, 0x82, 0x01, 0x0a, 0x0b, 0x55, 0x72, 0x69, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x3a, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x55, 0x72, 0x69, 0x12, 0x0b, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x72, 0x69, 0x2f, 0x7b, 0x75, 0x72, 0x69, 0x7d, 0x12, 0x37, 0x0a,
	0x06, 0x53, 0x65, 0x74, 0x55, 0x72, 0x69, 0x12, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x22, 0x07, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x72, 0x69, 0x3a, 0x01, 0x2a, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x75, 0x72, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protobuf_uri_v1_uri_exchange_proto_rawDescOnce sync.Once
	file_protobuf_uri_v1_uri_exchange_proto_rawDescData = file_protobuf_uri_v1_uri_exchange_proto_rawDesc
)

func file_protobuf_uri_v1_uri_exchange_proto_rawDescGZIP() []byte {
	file_protobuf_uri_v1_uri_exchange_proto_rawDescOnce.Do(func() {
		file_protobuf_uri_v1_uri_exchange_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_uri_v1_uri_exchange_proto_rawDescData)
	})
	return file_protobuf_uri_v1_uri_exchange_proto_rawDescData
}

var file_protobuf_uri_v1_uri_exchange_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protobuf_uri_v1_uri_exchange_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: v1.Request
	(*Response)(nil), // 1: v1.Response
}
var file_protobuf_uri_v1_uri_exchange_proto_depIdxs = []int32{
	0, // 0: v1.UriExchange.GetUri:input_type -> v1.Request
	0, // 1: v1.UriExchange.SetUri:input_type -> v1.Request
	1, // 2: v1.UriExchange.GetUri:output_type -> v1.Response
	1, // 3: v1.UriExchange.SetUri:output_type -> v1.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_uri_v1_uri_exchange_proto_init() }
func file_protobuf_uri_v1_uri_exchange_proto_init() {
	if File_protobuf_uri_v1_uri_exchange_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_uri_v1_uri_exchange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_uri_v1_uri_exchange_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_uri_v1_uri_exchange_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_uri_v1_uri_exchange_proto_goTypes,
		DependencyIndexes: file_protobuf_uri_v1_uri_exchange_proto_depIdxs,
		MessageInfos:      file_protobuf_uri_v1_uri_exchange_proto_msgTypes,
	}.Build()
	File_protobuf_uri_v1_uri_exchange_proto = out.File
	file_protobuf_uri_v1_uri_exchange_proto_rawDesc = nil
	file_protobuf_uri_v1_uri_exchange_proto_goTypes = nil
	file_protobuf_uri_v1_uri_exchange_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UriExchangeClient is the client API for UriExchange service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UriExchangeClient interface {
	GetUri(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	SetUri(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type uriExchangeClient struct {
	cc grpc.ClientConnInterface
}

func NewUriExchangeClient(cc grpc.ClientConnInterface) UriExchangeClient {
	return &uriExchangeClient{cc}
}

func (c *uriExchangeClient) GetUri(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v1.UriExchange/GetUri", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uriExchangeClient) SetUri(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v1.UriExchange/SetUri", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UriExchangeServer is the server API for UriExchange service.
type UriExchangeServer interface {
	GetUri(context.Context, *Request) (*Response, error)
	SetUri(context.Context, *Request) (*Response, error)
}

// UnimplementedUriExchangeServer can be embedded to have forward compatible implementations.
type UnimplementedUriExchangeServer struct {
}

func (*UnimplementedUriExchangeServer) GetUri(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUri not implemented")
}
func (*UnimplementedUriExchangeServer) SetUri(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUri not implemented")
}

func RegisterUriExchangeServer(s *grpc.Server, srv UriExchangeServer) {
	s.RegisterService(&_UriExchange_serviceDesc, srv)
}

func _UriExchange_GetUri_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UriExchangeServer).GetUri(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UriExchange/GetUri",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UriExchangeServer).GetUri(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _UriExchange_SetUri_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UriExchangeServer).SetUri(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UriExchange/SetUri",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UriExchangeServer).SetUri(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _UriExchange_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.UriExchange",
	HandlerType: (*UriExchangeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUri",
			Handler:    _UriExchange_GetUri_Handler,
		},
		{
			MethodName: "SetUri",
			Handler:    _UriExchange_SetUri_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/uri/v1/uri_exchange.proto",
}
