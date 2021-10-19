// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: options.proto

package options

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MethodHandler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authorization      *string `protobuf:"bytes,1,opt,name=authorization,proto3,oneof" json:"authorization,omitempty"`                                     // sso
	AuthorizationProxy *string `protobuf:"bytes,2,opt,name=authorization_proxy,json=authorizationProxy,proto3,oneof" json:"authorization_proxy,omitempty"` // signature
	Whitelisting       *string `protobuf:"bytes,3,opt,name=whitelisting,proto3,oneof" json:"whitelisting,omitempty"`                                       // ip
	Journal            *bool   `protobuf:"varint,4,opt,name=journal,proto3,oneof" json:"journal,omitempty"`
	MetricsAlias       *string `protobuf:"bytes,5,opt,name=metrics_alias,json=metricsAlias,proto3,oneof" json:"metrics_alias,omitempty"`
}

func (x *MethodHandler) Reset() {
	*x = MethodHandler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodHandler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodHandler) ProtoMessage() {}

func (x *MethodHandler) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodHandler.ProtoReflect.Descriptor instead.
func (*MethodHandler) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{0}
}

func (x *MethodHandler) GetAuthorization() string {
	if x != nil && x.Authorization != nil {
		return *x.Authorization
	}
	return ""
}

func (x *MethodHandler) GetAuthorizationProxy() string {
	if x != nil && x.AuthorizationProxy != nil {
		return *x.AuthorizationProxy
	}
	return ""
}

func (x *MethodHandler) GetWhitelisting() string {
	if x != nil && x.Whitelisting != nil {
		return *x.Whitelisting
	}
	return ""
}

func (x *MethodHandler) GetJournal() bool {
	if x != nil && x.Journal != nil {
		return *x.Journal
	}
	return false
}

func (x *MethodHandler) GetMetricsAlias() string {
	if x != nil && x.MetricsAlias != nil {
		return *x.MetricsAlias
	}
	return ""
}

type ServiceHandler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authorization      *string `protobuf:"bytes,1,opt,name=authorization,proto3,oneof" json:"authorization,omitempty"`                                     // sso
	AuthorizationProxy *string `protobuf:"bytes,2,opt,name=authorization_proxy,json=authorizationProxy,proto3,oneof" json:"authorization_proxy,omitempty"` // signature
	Whitelisting       *string `protobuf:"bytes,3,opt,name=whitelisting,proto3,oneof" json:"whitelisting,omitempty"`                                       // ip
}

func (x *ServiceHandler) Reset() {
	*x = ServiceHandler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceHandler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceHandler) ProtoMessage() {}

func (x *ServiceHandler) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceHandler.ProtoReflect.Descriptor instead.
func (*ServiceHandler) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceHandler) GetAuthorization() string {
	if x != nil && x.Authorization != nil {
		return *x.Authorization
	}
	return ""
}

func (x *ServiceHandler) GetAuthorizationProxy() string {
	if x != nil && x.AuthorizationProxy != nil {
		return *x.AuthorizationProxy
	}
	return ""
}

func (x *ServiceHandler) GetWhitelisting() string {
	if x != nil && x.Whitelisting != nil {
		return *x.Whitelisting
	}
	return ""
}

var file_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*MethodHandler)(nil),
		Field:         74500,
		Name:          "interceptor.method_handler",
		Tag:           "bytes,74500,opt,name=method_handler",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*ServiceHandler)(nil),
		Field:         74501,
		Name:          "interceptor.service_handler",
		Tag:           "bytes,74501,opt,name=service_handler",
		Filename:      "options.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional interceptor.MethodHandler method_handler = 74500;
	E_MethodHandler = &file_options_proto_extTypes[0]
)

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional interceptor.ServiceHandler service_handler = 74501;
	E_ServiceHandler = &file_options_proto_extTypes[1]
)

var File_options_proto protoreflect.FileDescriptor

var file_options_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb,
	0x02, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x12, 0x29, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x13, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x12, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x88, 0x01,
	0x01, 0x12, 0x27, 0x0a, 0x0c, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0c, 0x77, 0x68, 0x69, 0x74, 0x65,
	0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6a, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x07, 0x6a,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x5f, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x04, 0x52, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x41, 0x6c, 0x69, 0x61, 0x73,
	0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x42, 0x0f, 0x0a,
	0x0d, 0x5f, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x22, 0xd5, 0x01, 0x0a,
	0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12,
	0x29, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x13, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x12, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x88, 0x01, 0x01,
	0x12, 0x27, 0x0a, 0x0c, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0c, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6c,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x16, 0x0a, 0x14, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x3a, 0x66, 0x0a, 0x0e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x84, 0xc6, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x0d, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x88, 0x01, 0x01, 0x3a, 0x6a, 0x0a, 0x0f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12,
	0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x85, 0xc6, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_options_proto_rawDescOnce sync.Once
	file_options_proto_rawDescData = file_options_proto_rawDesc
)

func file_options_proto_rawDescGZIP() []byte {
	file_options_proto_rawDescOnce.Do(func() {
		file_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_options_proto_rawDescData)
	})
	return file_options_proto_rawDescData
}

var file_options_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_options_proto_goTypes = []interface{}{
	(*MethodHandler)(nil),               // 0: interceptor.MethodHandler
	(*ServiceHandler)(nil),              // 1: interceptor.ServiceHandler
	(*descriptorpb.MethodOptions)(nil),  // 2: google.protobuf.MethodOptions
	(*descriptorpb.ServiceOptions)(nil), // 3: google.protobuf.ServiceOptions
}
var file_options_proto_depIdxs = []int32{
	2, // 0: interceptor.method_handler:extendee -> google.protobuf.MethodOptions
	3, // 1: interceptor.service_handler:extendee -> google.protobuf.ServiceOptions
	0, // 2: interceptor.method_handler:type_name -> interceptor.MethodHandler
	1, // 3: interceptor.service_handler:type_name -> interceptor.ServiceHandler
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	2, // [2:4] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_options_proto_init() }
func file_options_proto_init() {
	if File_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodHandler); i {
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
		file_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceHandler); i {
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
	file_options_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_options_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_options_proto_goTypes,
		DependencyIndexes: file_options_proto_depIdxs,
		MessageInfos:      file_options_proto_msgTypes,
		ExtensionInfos:    file_options_proto_extTypes,
	}.Build()
	File_options_proto = out.File
	file_options_proto_rawDesc = nil
	file_options_proto_goTypes = nil
	file_options_proto_depIdxs = nil
}
