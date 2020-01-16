// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/route/v3/route.proto

package envoy_config_route_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type RouteConfiguration struct {
	Name                            string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	VirtualHosts                    []*VirtualHost          `protobuf:"bytes,2,rep,name=virtual_hosts,json=virtualHosts,proto3" json:"virtual_hosts,omitempty"`
	Vhds                            *Vhds                   `protobuf:"bytes,9,opt,name=vhds,proto3" json:"vhds,omitempty"`
	InternalOnlyHeaders             []string                `protobuf:"bytes,3,rep,name=internal_only_headers,json=internalOnlyHeaders,proto3" json:"internal_only_headers,omitempty"`
	ResponseHeadersToAdd            []*v3.HeaderValueOption `protobuf:"bytes,4,rep,name=response_headers_to_add,json=responseHeadersToAdd,proto3" json:"response_headers_to_add,omitempty"`
	ResponseHeadersToRemove         []string                `protobuf:"bytes,5,rep,name=response_headers_to_remove,json=responseHeadersToRemove,proto3" json:"response_headers_to_remove,omitempty"`
	RequestHeadersToAdd             []*v3.HeaderValueOption `protobuf:"bytes,6,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	RequestHeadersToRemove          []string                `protobuf:"bytes,8,rep,name=request_headers_to_remove,json=requestHeadersToRemove,proto3" json:"request_headers_to_remove,omitempty"`
	MostSpecificHeaderMutationsWins bool                    `protobuf:"varint,10,opt,name=most_specific_header_mutations_wins,json=mostSpecificHeaderMutationsWins,proto3" json:"most_specific_header_mutations_wins,omitempty"`
	ValidateClusters                *wrappers.BoolValue     `protobuf:"bytes,7,opt,name=validate_clusters,json=validateClusters,proto3" json:"validate_clusters,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}                `json:"-"`
	XXX_unrecognized                []byte                  `json:"-"`
	XXX_sizecache                   int32                   `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdb402428dbbe6a7, []int{0}
}

func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfiguration.Unmarshal(m, b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return xxx_messageInfo_RouteConfiguration.Size(m)
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetVirtualHosts() []*VirtualHost {
	if m != nil {
		return m.VirtualHosts
	}
	return nil
}

func (m *RouteConfiguration) GetVhds() *Vhds {
	if m != nil {
		return m.Vhds
	}
	return nil
}

func (m *RouteConfiguration) GetInternalOnlyHeaders() []string {
	if m != nil {
		return m.InternalOnlyHeaders
	}
	return nil
}

func (m *RouteConfiguration) GetResponseHeadersToAdd() []*v3.HeaderValueOption {
	if m != nil {
		return m.ResponseHeadersToAdd
	}
	return nil
}

func (m *RouteConfiguration) GetResponseHeadersToRemove() []string {
	if m != nil {
		return m.ResponseHeadersToRemove
	}
	return nil
}

func (m *RouteConfiguration) GetRequestHeadersToAdd() []*v3.HeaderValueOption {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

func (m *RouteConfiguration) GetRequestHeadersToRemove() []string {
	if m != nil {
		return m.RequestHeadersToRemove
	}
	return nil
}

func (m *RouteConfiguration) GetMostSpecificHeaderMutationsWins() bool {
	if m != nil {
		return m.MostSpecificHeaderMutationsWins
	}
	return false
}

func (m *RouteConfiguration) GetValidateClusters() *wrappers.BoolValue {
	if m != nil {
		return m.ValidateClusters
	}
	return nil
}

type Vhds struct {
	ConfigSource         *v3.ConfigSource `protobuf:"bytes,1,opt,name=config_source,json=configSource,proto3" json:"config_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Vhds) Reset()         { *m = Vhds{} }
func (m *Vhds) String() string { return proto.CompactTextString(m) }
func (*Vhds) ProtoMessage()    {}
func (*Vhds) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdb402428dbbe6a7, []int{1}
}

func (m *Vhds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vhds.Unmarshal(m, b)
}
func (m *Vhds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vhds.Marshal(b, m, deterministic)
}
func (m *Vhds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vhds.Merge(m, src)
}
func (m *Vhds) XXX_Size() int {
	return xxx_messageInfo_Vhds.Size(m)
}
func (m *Vhds) XXX_DiscardUnknown() {
	xxx_messageInfo_Vhds.DiscardUnknown(m)
}

var xxx_messageInfo_Vhds proto.InternalMessageInfo

func (m *Vhds) GetConfigSource() *v3.ConfigSource {
	if m != nil {
		return m.ConfigSource
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.config.route.v3.RouteConfiguration")
	proto.RegisterType((*Vhds)(nil), "envoy.config.route.v3.Vhds")
}

func init() { proto.RegisterFile("envoy/config/route/v3/route.proto", fileDescriptor_bdb402428dbbe6a7) }

var fileDescriptor_bdb402428dbbe6a7 = []byte{
	// 585 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcf, 0x6e, 0xd3, 0x3e,
	0x1c, 0x57, 0xd6, 0xee, 0x4f, 0xbd, 0x4d, 0xda, 0xdc, 0xdf, 0xd6, 0xfc, 0x8a, 0xb4, 0xb5, 0x9d,
	0x04, 0x39, 0xa0, 0x44, 0x6a, 0x4f, 0x8c, 0x13, 0xd9, 0x61, 0x3b, 0x80, 0x36, 0x32, 0x34, 0x8e,
	0x91, 0x9b, 0xb8, 0xad, 0xa5, 0xd4, 0xdf, 0x60, 0x3b, 0x19, 0x7d, 0x03, 0xc4, 0x91, 0x23, 0xef,
	0xc3, 0x83, 0x70, 0xe3, 0x19, 0x76, 0x42, 0xb1, 0x1d, 0x41, 0xd7, 0xc2, 0x81, 0x9b, 0x9d, 0xef,
	0xe7, 0x9f, 0x9d, 0x8f, 0x51, 0x9f, 0xf2, 0x12, 0x16, 0x41, 0x02, 0x7c, 0xc2, 0xa6, 0x81, 0x80,
	0x42, 0xd1, 0xa0, 0x1c, 0x99, 0x85, 0x9f, 0x0b, 0x50, 0x80, 0x8f, 0x34, 0xc4, 0x37, 0x10, 0xdf,
	0x4c, 0xca, 0x51, 0xf7, 0x74, 0x89, 0x99, 0x80, 0xd0, 0xc4, 0x31, 0x91, 0x96, 0xd7, 0xf5, 0xd6,
	0x02, 0xcc, 0x36, 0x96, 0x50, 0x88, 0xa4, 0x46, 0x3e, 0xff, 0x4b, 0x88, 0x38, 0x81, 0x79, 0x0e,
	0x9c, 0x72, 0x25, 0x2d, 0xfa, 0x64, 0x0a, 0x30, 0xcd, 0x68, 0xa0, 0x77, 0xe3, 0x62, 0x12, 0xdc,
	0x0b, 0x92, 0xe7, 0x54, 0xd4, 0xf3, 0x7e, 0x91, 0xe6, 0x24, 0x20, 0x9c, 0x83, 0x22, 0x8a, 0x01,
	0x97, 0x41, 0x49, 0x85, 0x64, 0xc0, 0x19, 0x9f, 0x5a, 0x48, 0xa7, 0x24, 0x19, 0x4b, 0x49, 0x65,
	0x63, 0x17, 0x66, 0x30, 0xf8, 0xbe, 0x89, 0x70, 0x54, 0xd9, 0x5e, 0xe8, 0x2c, 0x85, 0xd0, 0x0a,
	0x18, 0xa3, 0x26, 0x27, 0x73, 0xea, 0x3a, 0x3d, 0xc7, 0x6b, 0x45, 0x7a, 0x8d, 0x2f, 0xd1, 0x7e,
	0xc9, 0x84, 0x2a, 0x48, 0x16, 0xcf, 0x40, 0x2a, 0xe9, 0x6e, 0xf4, 0x1a, 0xde, 0xee, 0x70, 0xe0,
	0xaf, 0xbd, 0x2e, 0xff, 0xce, 0x60, 0xaf, 0x40, 0xaa, 0x68, 0xaf, 0xfc, 0xb5, 0x91, 0x38, 0x40,
	0xcd, 0x72, 0x96, 0x4a, 0xb7, 0xd5, 0x73, 0xbc, 0xdd, 0xe1, 0x93, 0x3f, 0xf1, 0x67, 0xa9, 0x8c,
	0x34, 0x10, 0x0f, 0xd1, 0x11, 0xe3, 0x8a, 0x0a, 0x4e, 0xb2, 0x18, 0x78, 0xb6, 0x88, 0x67, 0x94,
	0xa4, 0x54, 0x48, 0xb7, 0xd1, 0x6b, 0x78, 0xad, 0xa8, 0x5d, 0x0f, 0xaf, 0x79, 0xb6, 0xb8, 0x32,
	0x23, 0xcc, 0x50, 0x47, 0x50, 0x99, 0x03, 0x97, 0xb4, 0x86, 0xc7, 0x0a, 0x62, 0x92, 0xa6, 0x6e,
	0x53, 0xe7, 0x7e, 0xb6, 0xec, 0x5b, 0xfd, 0xae, 0xca, 0xd6, 0xf0, 0xef, 0x48, 0x56, 0xd0, 0xeb,
	0xbc, 0xba, 0x8b, 0xb0, 0xf5, 0x10, 0x6e, 0x7d, 0x71, 0x1a, 0x07, 0x3f, 0xb6, 0xa3, 0xff, 0x6a,
	0x49, 0xeb, 0xf2, 0x0e, 0x5e, 0xa5, 0x29, 0x7e, 0x89, 0xba, 0xeb, 0xac, 0x04, 0x9d, 0x43, 0x49,
	0xdd, 0x4d, 0x9d, 0xb1, 0xb3, 0xc2, 0x8c, 0xf4, 0x18, 0x4f, 0xd1, 0xb1, 0xa0, 0x1f, 0x0a, 0x2a,
	0xd5, 0xe3, 0x98, 0x5b, 0xff, 0x1c, 0xb3, 0x6d, 0x15, 0x97, 0x52, 0xbe, 0x40, 0xff, 0xaf, 0x31,
	0xb2, 0x21, 0x77, 0x74, 0xc8, 0xe3, 0xc7, 0x3c, 0x9b, 0xf1, 0x35, 0x3a, 0x9b, 0x83, 0x54, 0xb1,
	0xcc, 0x69, 0xc2, 0x26, 0x2c, 0xb1, 0x02, 0xf1, 0xbc, 0xb0, 0x85, 0x8b, 0xef, 0x19, 0x97, 0x2e,
	0xea, 0x39, 0xde, 0x4e, 0x74, 0x5a, 0x41, 0x6f, 0x2d, 0xd2, 0x28, 0xbd, 0xa9, 0x71, 0xef, 0x19,
	0x97, 0xf8, 0x12, 0x1d, 0xd6, 0x25, 0x8c, 0x93, 0xac, 0x90, 0xaa, 0xfa, 0x93, 0xdb, 0xba, 0x0b,
	0x5d, 0xdf, 0x54, 0xdd, 0xaf, 0xab, 0xee, 0x87, 0x00, 0x99, 0x3e, 0x65, 0x74, 0x50, 0x93, 0x2e,
	0x2c, 0xe7, 0xfc, 0xe9, 0xd7, 0x6f, 0x9f, 0x4e, 0xfa, 0xc8, 0xbc, 0x4b, 0x9f, 0xe4, 0xcc, 0x2f,
	0x87, 0xfe, 0x6a, 0x99, 0x07, 0x12, 0x35, 0xab, 0x32, 0xe1, 0xb7, 0x68, 0x7f, 0xe9, 0x31, 0xea,
	0x76, 0xaf, 0x14, 0xb8, 0xbe, 0x61, 0xa3, 0x71, 0xab, 0x91, 0xe1, 0xce, 0x43, 0xb8, 0xf9, 0xd9,
	0xd9, 0x38, 0x70, 0xa2, 0xbd, 0xe4, 0xb7, 0xef, 0xe7, 0x6e, 0x15, 0xa1, 0x8d, 0x0e, 0x97, 0x22,
	0x54, 0x66, 0xe1, 0x08, 0x9d, 0x31, 0x30, 0xca, 0xb9, 0x80, 0x8f, 0x8b, 0xf5, 0x2d, 0x0f, 0x91,
	0xce, 0x7b, 0x53, 0x1d, 0xf7, 0xc6, 0x19, 0x6f, 0xe9, 0x73, 0x8f, 0x7e, 0x06, 0x00, 0x00, 0xff,
	0xff, 0xc3, 0xa9, 0xdf, 0x81, 0xa5, 0x04, 0x00, 0x00,
}