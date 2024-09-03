// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.26.1
// source: envoy/extensions/filters/http/basic_auth/v3/basic_auth.proto

package basic_authv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Basic HTTP authentication.
//
// Example:
//
// .. code-block:: yaml
//
//	users:
//	  inline_string: |-
//	    user1:{SHA}hashed_user1_password
//	    user2:{SHA}hashed_user2_password
type BasicAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Username-password pairs used to verify user credentials in the "Authorization" header.
	// The value needs to be the htpasswd format.
	// Reference to https://httpd.apache.org/docs/2.4/programs/htpasswd.html
	Users *v3.DataSource `protobuf:"bytes,1,opt,name=users,proto3" json:"users,omitempty"`
	// This field specifies the header name to forward a successfully authenticated user to
	// the backend. The header will be added to the request with the username as the value.
	//
	// If it is not specified, the username will not be forwarded.
	ForwardUsernameHeader string `protobuf:"bytes,2,opt,name=forward_username_header,json=forwardUsernameHeader,proto3" json:"forward_username_header,omitempty"`
}

func (x *BasicAuth) Reset() {
	*x = BasicAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicAuth) ProtoMessage() {}

func (x *BasicAuth) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicAuth.ProtoReflect.Descriptor instead.
func (*BasicAuth) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_rawDescGZIP(), []int{0}
}

func (x *BasicAuth) GetUsers() *v3.DataSource {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *BasicAuth) GetForwardUsernameHeader() string {
	if x != nil {
		return x.ForwardUsernameHeader
	}
	return ""
}

// Extra settings that may be added to per-route configuration for
// a virtual host or a cluster.
type BasicAuthPerRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Username-password pairs for this route.
	Users *v3.DataSource `protobuf:"bytes,1,opt,name=users,proto3" json:"users,omitempty"`
}

func (x *BasicAuthPerRoute) Reset() {
	*x = BasicAuthPerRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicAuthPerRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicAuthPerRoute) ProtoMessage() {}

func (x *BasicAuthPerRoute) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicAuthPerRoute.ProtoReflect.Descriptor instead.
func (*BasicAuthPerRoute) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_rawDescGZIP(), []int{1}
}

func (x *BasicAuthPerRoute) GetUsers() *v3.DataSource {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x3e, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x06, 0xb8, 0xb7, 0x8b, 0xa4, 0x02, 0x01, 0x52, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x43, 0x0a, 0x17, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xc8, 0x01, 0x00, 0xc0,
	0x01, 0x01, 0x52, 0x15, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x5b, 0x0a, 0x11, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x50, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x46,
	0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42,
	0x0e, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0xb8, 0xb7, 0x8b, 0xa4, 0x02, 0x01, 0x52,
	0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x42, 0xb6, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10,
	0x02, 0x0a, 0x39, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x33, 0x42, 0x0e, 0x42, 0x61,
	0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x76, 0x33, 0x3b, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x76, 0x33, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_rawDescData = file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_rawDesc
)

func file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_rawDescData
}

var file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_goTypes = []interface{}{
	(*BasicAuth)(nil),         // 0: envoy.extensions.filters.http.basic_auth.v3.BasicAuth
	(*BasicAuthPerRoute)(nil), // 1: envoy.extensions.filters.http.basic_auth.v3.BasicAuthPerRoute
	(*v3.DataSource)(nil),     // 2: envoy.config.core.v3.DataSource
}
var file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.filters.http.basic_auth.v3.BasicAuth.users:type_name -> envoy.config.core.v3.DataSource
	2, // 1: envoy.extensions.filters.http.basic_auth.v3.BasicAuthPerRoute.users:type_name -> envoy.config.core.v3.DataSource
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_init() }
func file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_init() {
	if File_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicAuth); i {
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
		file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicAuthPerRoute); i {
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
			RawDescriptor: file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto = out.File
	file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_rawDesc = nil
	file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_goTypes = nil
	file_envoy_extensions_filters_http_basic_auth_v3_basic_auth_proto_depIdxs = nil
}
