// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.29.3
// source: envoy/extensions/filters/http/dynamic_modules/v3/dynamic_modules.proto

package dynamic_modulesv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/dynamic_modules/v3"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Configuration of the HTTP filter for dynamic modules. This filter allows loading shared object files
// that can be loaded via dlopen by the HTTP filter.
//
// A module can be loaded by multiple HTTP filters, hence the program can be structured in a way that
// the module is loaded only once and shared across multiple filters providing multiple functionalities.
//
// Currently, the implementation is work in progress and not usable.
type DynamicModuleFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the shared-object level configuration.
	DynamicModuleConfig *v3.DynamicModuleConfig `protobuf:"bytes,1,opt,name=dynamic_module_config,json=dynamicModuleConfig,proto3" json:"dynamic_module_config,omitempty"`
	// The name for this filter configuration. This can be used to distinguish between different filter implementations
	// inside a dynamic module. For example, a module can have completely different filter implementations.
	// When Envoy receives this configuration, it passes the filter_name to the dynamic module's HTTP filter config init function
	// together with the filter_config.
	// That way a module can decide which in-module filter implementation to use based on the name at load time.
	FilterName string `protobuf:"bytes,2,opt,name=filter_name,json=filterName,proto3" json:"filter_name,omitempty"`
	// The configuration for the filter chosen by filter_name. This is passed to the module's HTTP filter initialization function.
	// Together with the filter_name, the module can decide which in-module filter implementation to use and
	// fine-tune the behavior of the filter.
	//
	// For example, if a module has two filter implementations, one for logging and one for header manipulation,
	// filter_name is used to choose either logging or header manipulation. The filter_config can be used to
	// configure the logging level or the header manipulation behavior.
	//
	// “google.protobuf.Struct“ is serialized as JSON before
	// passing it to the plugin. “google.protobuf.BytesValue“ and
	// “google.protobuf.StringValue“ are passed directly without the wrapper.
	//
	// .. code-block:: yaml
	//
	//	# Passing in a string
	//	filter_config:
	//	  "@type": "type.googleapis.com/google.protobuf.StringValue"
	//	  value: hello
	//
	//	# Passing in raw bytes
	//	filter_config:
	//	  "@type": "type.googleapis.com/google.protobuf.BytesValue"
	//	  value: aGVsbG8= # echo -n "hello" | base64
	FilterConfig *anypb.Any `protobuf:"bytes,3,opt,name=filter_config,json=filterConfig,proto3" json:"filter_config,omitempty"`
}

func (x *DynamicModuleFilter) Reset() {
	*x = DynamicModuleFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynamicModuleFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamicModuleFilter) ProtoMessage() {}

func (x *DynamicModuleFilter) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamicModuleFilter.ProtoReflect.Descriptor instead.
func (*DynamicModuleFilter) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_rawDescGZIP(), []int{0}
}

func (x *DynamicModuleFilter) GetDynamicModuleConfig() *v3.DynamicModuleConfig {
	if x != nil {
		return x.DynamicModuleConfig
	}
	return nil
}

func (x *DynamicModuleFilter) GetFilterName() string {
	if x != nil {
		return x.FilterName
	}
	return ""
}

func (x *DynamicModuleFilter) GetFilterConfig() *anypb.Any {
	if x != nil {
		return x.FilterConfig
	}
	return nil
}

var File_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_rawDesc = []byte{
	0x0a, 0x46, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f,
	0x76, 0x33, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x39, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x79, 0x6e,
	0x61, 0x6d, 0x69, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x33, 0x2f,
	0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xdf, 0x01, 0x0a, 0x13, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x6c, 0x0a, 0x15, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x13, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x42, 0xd2, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0xd2, 0xc6, 0xa4,
	0xe1, 0x06, 0x02, 0x08, 0x01, 0x0a, 0x3e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x2e, 0x76, 0x33, 0x42, 0x13, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x69, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74,
	0x74, 0x70, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x2f, 0x76, 0x33, 0x3b, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_rawDescData = file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_rawDesc
)

func file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_rawDescData
}

var file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_goTypes = []interface{}{
	(*DynamicModuleFilter)(nil),    // 0: envoy.extensions.filters.http.dynamic_modules.v3.DynamicModuleFilter
	(*v3.DynamicModuleConfig)(nil), // 1: envoy.extensions.dynamic_modules.v3.DynamicModuleConfig
	(*anypb.Any)(nil),              // 2: google.protobuf.Any
}
var file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.filters.http.dynamic_modules.v3.DynamicModuleFilter.dynamic_module_config:type_name -> envoy.extensions.dynamic_modules.v3.DynamicModuleConfig
	2, // 1: envoy.extensions.filters.http.dynamic_modules.v3.DynamicModuleFilter.filter_config:type_name -> google.protobuf.Any
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_init() }
func file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_init() {
	if File_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DynamicModuleFilter); i {
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
			RawDescriptor: file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto = out.File
	file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_rawDesc = nil
	file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_goTypes = nil
	file_envoy_extensions_filters_http_dynamic_modules_v3_dynamic_modules_proto_depIdxs = nil
}
