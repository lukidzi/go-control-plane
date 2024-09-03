// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.26.1
// source: envoy/extensions/tracers/opentelemetry/samplers/v3/dynatrace_sampler.proto

package samplersv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
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

// Configuration for the Dynatrace Sampler extension.
// [#extension: envoy.tracers.opentelemetry.samplers.dynatrace]
type DynatraceSamplerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Dynatrace tenant.
	//
	// The value can be obtained from the Envoy deployment page in Dynatrace.
	Tenant string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	// The id of the Dynatrace cluster id.
	//
	// The value can be obtained from the Envoy deployment page in Dynatrace.
	ClusterId int32 `protobuf:"varint,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// The HTTP service to fetch the sampler configuration from the Dynatrace API (root spans per minute). For example:
	//
	// .. code-block:: yaml
	//
	//	http_service:
	//	  http_uri:
	//	    cluster: dynatrace
	//	    uri: <tenant>.dev.dynatracelabs.com/api/v2/samplingConfiguration
	//	    timeout: 10s
	//	  request_headers_to_add:
	//	  - header:
	//	      key : "authorization"
	//	      value: "Api-Token dt..."
	HttpService *v3.HttpService `protobuf:"bytes,3,opt,name=http_service,json=httpService,proto3" json:"http_service,omitempty"`
	// Default number of root spans per minute, used when the value can't be obtained from the Dynatrace API.
	//
	// A default value of “1000“ is used when:
	//
	// - “root_spans_per_minute“ is unset
	// - “root_spans_per_minute“ is set to 0
	RootSpansPerMinute uint32 `protobuf:"varint,4,opt,name=root_spans_per_minute,json=rootSpansPerMinute,proto3" json:"root_spans_per_minute,omitempty"`
}

func (x *DynatraceSamplerConfig) Reset() {
	*x = DynatraceSamplerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynatraceSamplerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynatraceSamplerConfig) ProtoMessage() {}

func (x *DynatraceSamplerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynatraceSamplerConfig.ProtoReflect.Descriptor instead.
func (*DynatraceSamplerConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_rawDescGZIP(), []int{0}
}

func (x *DynatraceSamplerConfig) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *DynatraceSamplerConfig) GetClusterId() int32 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

func (x *DynatraceSamplerConfig) GetHttpService() *v3.HttpService {
	if x != nil {
		return x.HttpService
	}
	return nil
}

func (x *DynatraceSamplerConfig) GetRootSpansPerMinute() uint32 {
	if x != nil {
		return x.RootSpansPerMinute
	}
	return 0
}

var File_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto protoreflect.FileDescriptor

var file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x74,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72,
	0x73, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x32, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x72, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x33,
	0x1a, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01, 0x0a, 0x16, 0x44, 0x79, 0x6e,
	0x61, 0x74, 0x72, 0x61, 0x63, 0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x0c, 0x68, 0x74,
	0x74, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x31, 0x0a, 0x15, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x5f, 0x70,
	0x65, 0x72, 0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x12, 0x72, 0x6f, 0x6f, 0x74, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x50, 0x65, 0x72, 0x4d, 0x69, 0x6e,
	0x75, 0x74, 0x65, 0x42, 0xc9, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x40,
	0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x72, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x33,
	0x42, 0x15, 0x44, 0x79, 0x6e, 0x61, 0x74, 0x72, 0x61, 0x63, 0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x64, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x74,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72,
	0x73, 0x2f, 0x76, 0x33, 0x3b, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x73, 0x76, 0x33, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_rawDescOnce sync.Once
	file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_rawDescData = file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_rawDesc
)

func file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_rawDescGZIP() []byte {
	file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_rawDescData)
	})
	return file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_rawDescData
}

var file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_goTypes = []interface{}{
	(*DynatraceSamplerConfig)(nil), // 0: envoy.extensions.tracers.opentelemetry.samplers.v3.DynatraceSamplerConfig
	(*v3.HttpService)(nil),         // 1: envoy.config.core.v3.HttpService
}
var file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.tracers.opentelemetry.samplers.v3.DynatraceSamplerConfig.http_service:type_name -> envoy.config.core.v3.HttpService
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_init() }
func file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_init() {
	if File_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DynatraceSamplerConfig); i {
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
			RawDescriptor: file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_msgTypes,
	}.Build()
	File_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto = out.File
	file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_rawDesc = nil
	file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_goTypes = nil
	file_envoy_extensions_tracers_opentelemetry_samplers_v3_dynatrace_sampler_proto_depIdxs = nil
}
