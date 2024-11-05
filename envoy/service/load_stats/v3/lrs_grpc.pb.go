// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.2
// source: envoy/service/load_stats/v3/lrs.proto

package load_statsv3

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	LoadReportingService_StreamLoadStats_FullMethodName = "/envoy.service.load_stats.v3.LoadReportingService/StreamLoadStats"
)

// LoadReportingServiceClient is the client API for LoadReportingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoadReportingServiceClient interface {
	// Advanced API to allow for multi-dimensional load balancing by remote
	// server. For receiving LB assignments, the steps are:
	// 1, The management server is configured with per cluster/zone/load metric
	//
	//	capacity configuration. The capacity configuration definition is
	//	outside of the scope of this document.
	//  2. Envoy issues a standard {Stream,Fetch}Endpoints request for the clusters
	//     to balance.
	//
	// Independently, Envoy will initiate a StreamLoadStats bidi stream with a
	// management server:
	//  1. Once a connection establishes, the management server publishes a
	//     LoadStatsResponse for all clusters it is interested in learning load
	//     stats about.
	//  2. For each cluster, Envoy load balances incoming traffic to upstream hosts
	//     based on per-zone weights and/or per-instance weights (if specified)
	//     based on intra-zone LbPolicy. This information comes from the above
	//     {Stream,Fetch}Endpoints.
	//  3. When upstream hosts reply, they optionally add header <define header
	//     name> with ASCII representation of EndpointLoadMetricStats.
	//  4. Envoy aggregates load reports over the period of time given to it in
	//     LoadStatsResponse.load_reporting_interval. This includes aggregation
	//     stats Envoy maintains by itself (total_requests, rpc_errors etc.) as
	//     well as load metrics from upstream hosts.
	//  5. When the timer of load_reporting_interval expires, Envoy sends new
	//     LoadStatsRequest filled with load reports for each cluster.
	//  6. The management server uses the load reports from all reported Envoys
	//     from around the world, computes global assignment and prepares traffic
	//     assignment destined for each zone Envoys are located in. Goto 2.
	StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (LoadReportingService_StreamLoadStatsClient, error)
}

type loadReportingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoadReportingServiceClient(cc grpc.ClientConnInterface) LoadReportingServiceClient {
	return &loadReportingServiceClient{cc}
}

func (c *loadReportingServiceClient) StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (LoadReportingService_StreamLoadStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &LoadReportingService_ServiceDesc.Streams[0], LoadReportingService_StreamLoadStats_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &loadReportingServiceStreamLoadStatsClient{stream}
	return x, nil
}

type LoadReportingService_StreamLoadStatsClient interface {
	Send(*LoadStatsRequest) error
	Recv() (*LoadStatsResponse, error)
	grpc.ClientStream
}

type loadReportingServiceStreamLoadStatsClient struct {
	grpc.ClientStream
}

func (x *loadReportingServiceStreamLoadStatsClient) Send(m *LoadStatsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *loadReportingServiceStreamLoadStatsClient) Recv() (*LoadStatsResponse, error) {
	m := new(LoadStatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LoadReportingServiceServer is the server API for LoadReportingService service.
// All implementations should embed UnimplementedLoadReportingServiceServer
// for forward compatibility
type LoadReportingServiceServer interface {
	// Advanced API to allow for multi-dimensional load balancing by remote
	// server. For receiving LB assignments, the steps are:
	// 1, The management server is configured with per cluster/zone/load metric
	//
	//	capacity configuration. The capacity configuration definition is
	//	outside of the scope of this document.
	//  2. Envoy issues a standard {Stream,Fetch}Endpoints request for the clusters
	//     to balance.
	//
	// Independently, Envoy will initiate a StreamLoadStats bidi stream with a
	// management server:
	//  1. Once a connection establishes, the management server publishes a
	//     LoadStatsResponse for all clusters it is interested in learning load
	//     stats about.
	//  2. For each cluster, Envoy load balances incoming traffic to upstream hosts
	//     based on per-zone weights and/or per-instance weights (if specified)
	//     based on intra-zone LbPolicy. This information comes from the above
	//     {Stream,Fetch}Endpoints.
	//  3. When upstream hosts reply, they optionally add header <define header
	//     name> with ASCII representation of EndpointLoadMetricStats.
	//  4. Envoy aggregates load reports over the period of time given to it in
	//     LoadStatsResponse.load_reporting_interval. This includes aggregation
	//     stats Envoy maintains by itself (total_requests, rpc_errors etc.) as
	//     well as load metrics from upstream hosts.
	//  5. When the timer of load_reporting_interval expires, Envoy sends new
	//     LoadStatsRequest filled with load reports for each cluster.
	//  6. The management server uses the load reports from all reported Envoys
	//     from around the world, computes global assignment and prepares traffic
	//     assignment destined for each zone Envoys are located in. Goto 2.
	StreamLoadStats(LoadReportingService_StreamLoadStatsServer) error
}

// UnimplementedLoadReportingServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLoadReportingServiceServer struct {
}

func (UnimplementedLoadReportingServiceServer) StreamLoadStats(LoadReportingService_StreamLoadStatsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamLoadStats not implemented")
}

// UnsafeLoadReportingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoadReportingServiceServer will
// result in compilation errors.
type UnsafeLoadReportingServiceServer interface {
	mustEmbedUnimplementedLoadReportingServiceServer()
}

func RegisterLoadReportingServiceServer(s grpc.ServiceRegistrar, srv LoadReportingServiceServer) {
	s.RegisterService(&LoadReportingService_ServiceDesc, srv)
}

func _LoadReportingService_StreamLoadStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LoadReportingServiceServer).StreamLoadStats(&loadReportingServiceStreamLoadStatsServer{stream})
}

type LoadReportingService_StreamLoadStatsServer interface {
	Send(*LoadStatsResponse) error
	Recv() (*LoadStatsRequest, error)
	grpc.ServerStream
}

type loadReportingServiceStreamLoadStatsServer struct {
	grpc.ServerStream
}

func (x *loadReportingServiceStreamLoadStatsServer) Send(m *LoadStatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *loadReportingServiceStreamLoadStatsServer) Recv() (*LoadStatsRequest, error) {
	m := new(LoadStatsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LoadReportingService_ServiceDesc is the grpc.ServiceDesc for LoadReportingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoadReportingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.load_stats.v3.LoadReportingService",
	HandlerType: (*LoadReportingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamLoadStats",
			Handler:       _LoadReportingService_StreamLoadStats_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/load_stats/v3/lrs.proto",
}
