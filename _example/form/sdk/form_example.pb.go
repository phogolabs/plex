// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: form_example.proto

package sdk

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

type FooRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: validate:"required,gt=0" form:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" validate:"required,gt=0" form:"name"`
}

func (x *FooRequest) Reset() {
	*x = FooRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_form_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FooRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FooRequest) ProtoMessage() {}

func (x *FooRequest) ProtoReflect() protoreflect.Message {
	mi := &file_form_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FooRequest.ProtoReflect.Descriptor instead.
func (*FooRequest) Descriptor() ([]byte, []int) {
	return file_form_example_proto_rawDescGZIP(), []int{0}
}

func (x *FooRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FooResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Actual items
	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *FooResponse) Reset() {
	*x = FooResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_form_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FooResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FooResponse) ProtoMessage() {}

func (x *FooResponse) ProtoReflect() protoreflect.Message {
	mi := &file_form_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FooResponse.ProtoReflect.Descriptor instead.
func (*FooResponse) Descriptor() ([]byte, []int) {
	return file_form_example_proto_rawDescGZIP(), []int{1}
}

func (x *FooResponse) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_form_example_proto protoreflect.FileDescriptor

var file_form_example_proto_rawDesc = []byte{
	0x0a, 0x12, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x66, 0x6f, 0x6f, 0x2e, 0x62, 0x61, 0x72, 0x2e, 0x73, 0x64,
	0x6b, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x20, 0x0a, 0x0a, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x21, 0x0a, 0x0b, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x32, 0x57, 0x0a, 0x06, 0x46, 0x6f, 0x6f, 0x41, 0x50, 0x49, 0x12, 0x4d,
	0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x62, 0x61, 0x72,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x62, 0x61, 0x72, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x46, 0x6f,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0c, 0x22, 0x07, 0x2f, 0x76, 0x32, 0x2f, 0x62, 0x61, 0x72, 0x3a, 0x01, 0x2a, 0x42, 0x31, 0x5a,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x68, 0x6f, 0x67,
	0x6f, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x6c, 0x65, 0x78, 0x2f, 0x5f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2f, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x64, 0x6b, 0x3b, 0x73, 0x64, 0x6b,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_form_example_proto_rawDescOnce sync.Once
	file_form_example_proto_rawDescData = file_form_example_proto_rawDesc
)

func file_form_example_proto_rawDescGZIP() []byte {
	file_form_example_proto_rawDescOnce.Do(func() {
		file_form_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_form_example_proto_rawDescData)
	})
	return file_form_example_proto_rawDescData
}

var file_form_example_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_form_example_proto_goTypes = []interface{}{
	(*FooRequest)(nil),  // 0: foo.bar.sdk.FooRequest
	(*FooResponse)(nil), // 1: foo.bar.sdk.FooResponse
}
var file_form_example_proto_depIdxs = []int32{
	0, // 0: foo.bar.sdk.FooAPI.Post:input_type -> foo.bar.sdk.FooRequest
	1, // 1: foo.bar.sdk.FooAPI.Post:output_type -> foo.bar.sdk.FooResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_form_example_proto_init() }
func file_form_example_proto_init() {
	if File_form_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_form_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FooRequest); i {
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
		file_form_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FooResponse); i {
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
			RawDescriptor: file_form_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_form_example_proto_goTypes,
		DependencyIndexes: file_form_example_proto_depIdxs,
		MessageInfos:      file_form_example_proto_msgTypes,
	}.Build()
	File_form_example_proto = out.File
	file_form_example_proto_rawDesc = nil
	file_form_example_proto_goTypes = nil
	file_form_example_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FooAPIClient is the client API for FooAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FooAPIClient interface {
	Post(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooResponse, error)
}

type fooAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewFooAPIClient(cc grpc.ClientConnInterface) FooAPIClient {
	return &fooAPIClient{cc}
}

func (c *fooAPIClient) Post(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooResponse, error) {
	out := new(FooResponse)
	err := c.cc.Invoke(ctx, "/foo.bar.sdk.FooAPI/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FooAPIServer is the server API for FooAPI service.
type FooAPIServer interface {
	Post(context.Context, *FooRequest) (*FooResponse, error)
}

// UnimplementedFooAPIServer can be embedded to have forward compatible implementations.
type UnimplementedFooAPIServer struct {
}

func (*UnimplementedFooAPIServer) Post(context.Context, *FooRequest) (*FooResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}

func RegisterFooAPIServer(s *grpc.Server, srv FooAPIServer) {
	s.RegisterService(&_FooAPI_serviceDesc, srv)
}

func _FooAPI_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FooRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FooAPIServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foo.bar.sdk.FooAPI/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FooAPIServer).Post(ctx, req.(*FooRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FooAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "foo.bar.sdk.FooAPI",
	HandlerType: (*FooAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Post",
			Handler:    _FooAPI_Post_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "form_example.proto",
}