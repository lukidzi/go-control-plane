// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/core/v4alpha/event_service_config.proto

package envoy_config_core_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type EventServiceConfig struct {
	// Types that are valid to be assigned to ConfigSourceSpecifier:
	//	*EventServiceConfig_GrpcService
	ConfigSourceSpecifier isEventServiceConfig_ConfigSourceSpecifier `protobuf_oneof:"config_source_specifier"`
	XXX_NoUnkeyedLiteral  struct{}                                   `json:"-"`
	XXX_unrecognized      []byte                                     `json:"-"`
	XXX_sizecache         int32                                      `json:"-"`
}

func (m *EventServiceConfig) Reset()         { *m = EventServiceConfig{} }
func (m *EventServiceConfig) String() string { return proto.CompactTextString(m) }
func (*EventServiceConfig) ProtoMessage()    {}
func (*EventServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ed78bbaa9d45819, []int{0}
}

func (m *EventServiceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventServiceConfig.Unmarshal(m, b)
}
func (m *EventServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventServiceConfig.Marshal(b, m, deterministic)
}
func (m *EventServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventServiceConfig.Merge(m, src)
}
func (m *EventServiceConfig) XXX_Size() int {
	return xxx_messageInfo_EventServiceConfig.Size(m)
}
func (m *EventServiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_EventServiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_EventServiceConfig proto.InternalMessageInfo

type isEventServiceConfig_ConfigSourceSpecifier interface {
	isEventServiceConfig_ConfigSourceSpecifier()
}

type EventServiceConfig_GrpcService struct {
	GrpcService *GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3,oneof"`
}

func (*EventServiceConfig_GrpcService) isEventServiceConfig_ConfigSourceSpecifier() {}

func (m *EventServiceConfig) GetConfigSourceSpecifier() isEventServiceConfig_ConfigSourceSpecifier {
	if m != nil {
		return m.ConfigSourceSpecifier
	}
	return nil
}

func (m *EventServiceConfig) GetGrpcService() *GrpcService {
	if x, ok := m.GetConfigSourceSpecifier().(*EventServiceConfig_GrpcService); ok {
		return x.GrpcService
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*EventServiceConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*EventServiceConfig_GrpcService)(nil),
	}
}

func init() {
	proto.RegisterType((*EventServiceConfig)(nil), "envoy.config.core.v4alpha.EventServiceConfig")
}

func init() {
	proto.RegisterFile("envoy/config/core/v4alpha/event_service_config.proto", fileDescriptor_3ed78bbaa9d45819)
}

var fileDescriptor_3ed78bbaa9d45819 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x49, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x33,
	0x49, 0xcc, 0x29, 0xc8, 0x48, 0xd4, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0x89, 0x2f, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x8d, 0x87, 0xa8, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x04, 0xeb,
	0xd2, 0x83, 0x8a, 0x81, 0x74, 0xe9, 0x41, 0x75, 0x49, 0xe9, 0xe0, 0x36, 0x30, 0xbd, 0xa8, 0x20,
	0x19, 0x66, 0x1e, 0xc4, 0x20, 0x29, 0xd9, 0xd2, 0x94, 0x82, 0x44, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc,
	0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0xfd, 0xe2, 0x92, 0xc4, 0x92, 0xd2, 0x62, 0xa8, 0xb4,
	0x22, 0x86, 0x74, 0x59, 0x6a, 0x51, 0x71, 0x66, 0x7e, 0x5e, 0x66, 0x1e, 0xd4, 0x29, 0x52, 0xe2,
	0x65, 0x89, 0x39, 0x99, 0x29, 0x89, 0x25, 0xa9, 0xfa, 0x30, 0x06, 0x44, 0x42, 0x69, 0x23, 0x23,
	0x97, 0x90, 0x2b, 0xc8, 0x0b, 0xc1, 0x10, 0x1b, 0x9d, 0xc1, 0x2e, 0x12, 0xf2, 0xe6, 0xe2, 0x41,
	0x76, 0x87, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x9a, 0x1e, 0x4e, 0x1f, 0xe9, 0xb9, 0x17,
	0x15, 0x24, 0x43, 0xcd, 0xf0, 0x60, 0x08, 0xe2, 0x4e, 0x47, 0x70, 0xad, 0xf4, 0x66, 0x1d, 0xed,
	0x90, 0xd3, 0xe4, 0x52, 0xc7, 0xa2, 0xd9, 0x58, 0x0f, 0xd3, 0x72, 0x27, 0x39, 0x2e, 0x71, 0x88,
	0xa2, 0xf8, 0xe2, 0xfc, 0xd2, 0xa2, 0xe4, 0xd4, 0xf8, 0xe2, 0x82, 0xd4, 0xe4, 0xcc, 0xb4, 0xcc,
	0xd4, 0x22, 0x21, 0xe6, 0x1f, 0x4e, 0x8c, 0x4e, 0x3e, 0xbb, 0x1a, 0x4e, 0x5c, 0x64, 0x63, 0x12,
	0x60, 0xe6, 0x52, 0xcf, 0xcc, 0x87, 0x38, 0xa9, 0xa0, 0x28, 0xbf, 0xa2, 0x12, 0xb7, 0xeb, 0x9c,
	0xc4, 0x31, 0xad, 0x09, 0x00, 0xf9, 0x3f, 0x80, 0x31, 0x89, 0x0d, 0x1c, 0x10, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x6e, 0x09, 0xb9, 0x59, 0xe4, 0x01, 0x00, 0x00,
}