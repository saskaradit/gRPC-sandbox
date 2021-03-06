// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: maximum/maxpb/max.proto

package maxpb

import (
	context "context"
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

type MaximumNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *MaximumNumberRequest) Reset() {
	*x = MaximumNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maximum_maxpb_max_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaximumNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaximumNumberRequest) ProtoMessage() {}

func (x *MaximumNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_maximum_maxpb_max_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaximumNumberRequest.ProtoReflect.Descriptor instead.
func (*MaximumNumberRequest) Descriptor() ([]byte, []int) {
	return file_maximum_maxpb_max_proto_rawDescGZIP(), []int{0}
}

func (x *MaximumNumberRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type MaximumNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []int32 `protobuf:"varint,1,rep,packed,name=result,proto3" json:"result,omitempty"`
}

func (x *MaximumNumberResponse) Reset() {
	*x = MaximumNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maximum_maxpb_max_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaximumNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaximumNumberResponse) ProtoMessage() {}

func (x *MaximumNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_maximum_maxpb_max_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaximumNumberResponse.ProtoReflect.Descriptor instead.
func (*MaximumNumberResponse) Descriptor() ([]byte, []int) {
	return file_maximum_maxpb_max_proto_rawDescGZIP(), []int{1}
}

func (x *MaximumNumberResponse) GetResult() []int32 {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_maximum_maxpb_max_proto protoreflect.FileDescriptor

var file_maximum_maxpb_max_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x2f, 0x6d, 0x61, 0x78, 0x70, 0x62, 0x2f,
	0x6d, 0x61, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d,
	0x75, 0x6d, 0x22, 0x2e, 0x0a, 0x14, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x2f, 0x0a, 0x15, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x32, 0x66, 0x0a, 0x0e, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d,
	0x2e, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x2e,
	0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x11, 0x5a, 0x0f, 0x2e,
	0x2f, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x2f, 0x6d, 0x61, 0x78, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_maximum_maxpb_max_proto_rawDescOnce sync.Once
	file_maximum_maxpb_max_proto_rawDescData = file_maximum_maxpb_max_proto_rawDesc
)

func file_maximum_maxpb_max_proto_rawDescGZIP() []byte {
	file_maximum_maxpb_max_proto_rawDescOnce.Do(func() {
		file_maximum_maxpb_max_proto_rawDescData = protoimpl.X.CompressGZIP(file_maximum_maxpb_max_proto_rawDescData)
	})
	return file_maximum_maxpb_max_proto_rawDescData
}

var file_maximum_maxpb_max_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_maximum_maxpb_max_proto_goTypes = []interface{}{
	(*MaximumNumberRequest)(nil),  // 0: maximum.MaximumNumberRequest
	(*MaximumNumberResponse)(nil), // 1: maximum.MaximumNumberResponse
}
var file_maximum_maxpb_max_proto_depIdxs = []int32{
	0, // 0: maximum.MaximumService.MaximumNumber:input_type -> maximum.MaximumNumberRequest
	1, // 1: maximum.MaximumService.MaximumNumber:output_type -> maximum.MaximumNumberResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_maximum_maxpb_max_proto_init() }
func file_maximum_maxpb_max_proto_init() {
	if File_maximum_maxpb_max_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_maximum_maxpb_max_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaximumNumberRequest); i {
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
		file_maximum_maxpb_max_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaximumNumberResponse); i {
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
			RawDescriptor: file_maximum_maxpb_max_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_maximum_maxpb_max_proto_goTypes,
		DependencyIndexes: file_maximum_maxpb_max_proto_depIdxs,
		MessageInfos:      file_maximum_maxpb_max_proto_msgTypes,
	}.Build()
	File_maximum_maxpb_max_proto = out.File
	file_maximum_maxpb_max_proto_rawDesc = nil
	file_maximum_maxpb_max_proto_goTypes = nil
	file_maximum_maxpb_max_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MaximumServiceClient is the client API for MaximumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MaximumServiceClient interface {
	// Bi-Directional
	MaximumNumber(ctx context.Context, opts ...grpc.CallOption) (MaximumService_MaximumNumberClient, error)
}

type maximumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMaximumServiceClient(cc grpc.ClientConnInterface) MaximumServiceClient {
	return &maximumServiceClient{cc}
}

func (c *maximumServiceClient) MaximumNumber(ctx context.Context, opts ...grpc.CallOption) (MaximumService_MaximumNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MaximumService_serviceDesc.Streams[0], "/maximum.MaximumService/MaximumNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &maximumServiceMaximumNumberClient{stream}
	return x, nil
}

type MaximumService_MaximumNumberClient interface {
	Send(*MaximumNumberRequest) error
	Recv() (*MaximumNumberResponse, error)
	grpc.ClientStream
}

type maximumServiceMaximumNumberClient struct {
	grpc.ClientStream
}

func (x *maximumServiceMaximumNumberClient) Send(m *MaximumNumberRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *maximumServiceMaximumNumberClient) Recv() (*MaximumNumberResponse, error) {
	m := new(MaximumNumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MaximumServiceServer is the server API for MaximumService service.
type MaximumServiceServer interface {
	// Bi-Directional
	MaximumNumber(MaximumService_MaximumNumberServer) error
}

// UnimplementedMaximumServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMaximumServiceServer struct {
}

func (*UnimplementedMaximumServiceServer) MaximumNumber(MaximumService_MaximumNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method MaximumNumber not implemented")
}

func RegisterMaximumServiceServer(s *grpc.Server, srv MaximumServiceServer) {
	s.RegisterService(&_MaximumService_serviceDesc, srv)
}

func _MaximumService_MaximumNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MaximumServiceServer).MaximumNumber(&maximumServiceMaximumNumberServer{stream})
}

type MaximumService_MaximumNumberServer interface {
	Send(*MaximumNumberResponse) error
	Recv() (*MaximumNumberRequest, error)
	grpc.ServerStream
}

type maximumServiceMaximumNumberServer struct {
	grpc.ServerStream
}

func (x *maximumServiceMaximumNumberServer) Send(m *MaximumNumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *maximumServiceMaximumNumberServer) Recv() (*MaximumNumberRequest, error) {
	m := new(MaximumNumberRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MaximumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "maximum.MaximumService",
	HandlerType: (*MaximumServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MaximumNumber",
			Handler:       _MaximumService_MaximumNumber_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "maximum/maxpb/max.proto",
}
