// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.2
// source: example.proto

package sdk

import (
	_ "github.com/amsokol/protoc-gen-gotagger/proto/tagger"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// CreateUserRequest creates an account for given email and password
type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Represents the user's email address.
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty" validate:"required,gt=0,lte=256,email" form:"email"`
	// Represents the user's password.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" validate:"required,gte=8" form:"password"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// CreateUserResponse is the payload returned when a new user is created
type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Represents the publication's unique identifier.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validate:"required,gt=0,uuid"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_example_proto protoreflect.FileDescriptor

var file_example_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1a, 0x70, 0x68, 0x6f, 0x67, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x70, 0x6c, 0x65, 0x78, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x64, 0x6b, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74,
	0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x02, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0xf3, 0x01, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0xdc, 0x01, 0x92, 0x41, 0xa0, 0x01, 0x40, 0x01, 0x78, 0x80, 0x02, 0x8a, 0x01, 0x97,
	0x01, 0x5e, 0x28, 0x28, 0x5b, 0x5e, 0x3c, 0x3e, 0x28, 0x29, 0x5c, 0x5b, 0x5c, 0x5d, 0x5c, 0x5c,
	0x2e, 0x2c, 0x3b, 0x3a, 0x5c, 0x73, 0x40, 0x22, 0x5d, 0x2b, 0x28, 0x5c, 0x2e, 0x5b, 0x5e, 0x3c,
	0x3e, 0x28, 0x29, 0x5c, 0x5b, 0x5c, 0x5d, 0x5c, 0x5c, 0x2e, 0x2c, 0x3b, 0x3a, 0x5c, 0x73, 0x40,
	0x22, 0x5d, 0x2b, 0x29, 0x2a, 0x29, 0x7c, 0x28, 0x22, 0x2e, 0x2b, 0x22, 0x29, 0x29, 0x40, 0x28,
	0x28, 0x5c, 0x5b, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x31, 0x2c, 0x33, 0x7d, 0x5c, 0x2e, 0x5b,
	0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x31, 0x2c, 0x33, 0x7d, 0x5c, 0x2e, 0x5b, 0x30, 0x2d, 0x39, 0x5d,
	0x7b, 0x31, 0x2c, 0x33, 0x7d, 0x5c, 0x2e, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x31, 0x2c, 0x33,
	0x7d, 0x5c, 0x5d, 0x29, 0x7c, 0x28, 0x28, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x5c, 0x2d,
	0x30, 0x2d, 0x39, 0x5d, 0x2b, 0x5c, 0x2e, 0x29, 0x2b, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a,
	0x5d, 0x7b, 0x32, 0x2c, 0x7d, 0x29, 0x29, 0x24, 0x9a, 0x84, 0x9e, 0x03, 0x33, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2c,
	0x67, 0x74, 0x3d, 0x30, 0x2c, 0x6c, 0x74, 0x65, 0x3d, 0x32, 0x35, 0x36, 0x2c, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x52, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0x92, 0x41, 0x05, 0x40, 0x01,
	0x80, 0x01, 0x08, 0x9a, 0x84, 0x9e, 0x03, 0x29, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2c, 0x67, 0x74, 0x65, 0x3d, 0x38,
	0x22, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x3a, 0x18, 0x92, 0x41, 0x15,
	0x0a, 0x13, 0xd2, 0x01, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0xd2, 0x01, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x9f, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x88, 0x01, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x78, 0x92, 0x41, 0x53, 0x40, 0x01,
	0x80, 0x01, 0x01, 0x8a, 0x01, 0x4b, 0x5b, 0x61, 0x2d, 0x66, 0x41, 0x2d, 0x46, 0x30, 0x2d, 0x39,
	0x5d, 0x7b, 0x38, 0x7d, 0x2d, 0x5b, 0x61, 0x2d, 0x66, 0x41, 0x2d, 0x46, 0x30, 0x2d, 0x39, 0x5d,
	0x7b, 0x34, 0x7d, 0x2d, 0x5b, 0x61, 0x2d, 0x66, 0x41, 0x2d, 0x46, 0x30, 0x2d, 0x39, 0x5d, 0x7b,
	0x34, 0x7d, 0x2d, 0x5b, 0x61, 0x2d, 0x66, 0x41, 0x2d, 0x46, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x34,
	0x7d, 0x2d, 0x5b, 0x61, 0x2d, 0x66, 0x41, 0x2d, 0x46, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x31, 0x32,
	0x7d, 0x9a, 0x84, 0x9e, 0x03, 0x1d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2c, 0x67, 0x74, 0x3d, 0x30, 0x2c, 0x75, 0x75,
	0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x32, 0x8d, 0x01, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x50, 0x49, 0x12, 0x81, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x2d, 0x2e, 0x70, 0x68, 0x6f, 0x67, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x70,
	0x6c, 0x65, 0x78, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x70, 0x68, 0x6f, 0x67, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x70, 0x6c,
	0x65, 0x78, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x9c, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x68, 0x6f, 0x67, 0x6f, 0x6c, 0x61, 0x62, 0x73,
	0x2f, 0x70, 0x6c, 0x65, 0x78, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x64,
	0x6b, 0x3b, 0x73, 0x64, 0x6b, 0x92, 0x41, 0x6e, 0x12, 0x58, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x20, 0x41, 0x50, 0x49, 0x22, 0x47, 0x0a, 0x0a, 0x50, 0x68, 0x6f, 0x67, 0x6f, 0x20, 0x4c, 0x61,
	0x62, 0x73, 0x12, 0x21, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x68, 0x6f, 0x67, 0x6f, 0x6c, 0x61, 0x62, 0x73,
	0x2f, 0x70, 0x6c, 0x65, 0x78, 0x1a, 0x16, 0x6e, 0x6f, 0x2d, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x40,
	0x70, 0x68, 0x6f, 0x67, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31,
	0x2e, 0x30, 0x1a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x38, 0x30,
	0x38, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_proto_rawDescOnce sync.Once
	file_example_proto_rawDescData = file_example_proto_rawDesc
)

func file_example_proto_rawDescGZIP() []byte {
	file_example_proto_rawDescOnce.Do(func() {
		file_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_proto_rawDescData)
	})
	return file_example_proto_rawDescData
}

var file_example_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_example_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),  // 0: phogolabs.plex.example.sdk.CreateUserRequest
	(*CreateUserResponse)(nil), // 1: phogolabs.plex.example.sdk.CreateUserResponse
}
var file_example_proto_depIdxs = []int32{
	0, // 0: phogolabs.plex.example.sdk.UserAPI.CreateUser:input_type -> phogolabs.plex.example.sdk.CreateUserRequest
	1, // 1: phogolabs.plex.example.sdk.UserAPI.CreateUser:output_type -> phogolabs.plex.example.sdk.CreateUserResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_example_proto_init() }
func file_example_proto_init() {
	if File_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
			RawDescriptor: file_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_proto_goTypes,
		DependencyIndexes: file_example_proto_depIdxs,
		MessageInfos:      file_example_proto_msgTypes,
	}.Build()
	File_example_proto = out.File
	file_example_proto_rawDesc = nil
	file_example_proto_goTypes = nil
	file_example_proto_depIdxs = nil
}
