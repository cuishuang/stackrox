// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.3
// source: internalapi/sensor/network_connection_iservice.proto

package sensor

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	NetworkConnectionInfoService_PushNetworkConnectionInfo_FullMethodName = "/sensor.NetworkConnectionInfoService/PushNetworkConnectionInfo"
)

// NetworkConnectionInfoServiceClient is the client API for NetworkConnectionInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A Sensor service that allows Collector to send NetworkConnectionInfo messages
type NetworkConnectionInfoServiceClient interface {
	// Note: the response is a stream due to a bug in the C++ GRPC client library. The server is not expected to
	// send anything via this stream.
	PushNetworkConnectionInfo(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[NetworkConnectionInfoMessage, NetworkFlowsControlMessage], error)
}

type networkConnectionInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkConnectionInfoServiceClient(cc grpc.ClientConnInterface) NetworkConnectionInfoServiceClient {
	return &networkConnectionInfoServiceClient{cc}
}

func (c *networkConnectionInfoServiceClient) PushNetworkConnectionInfo(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[NetworkConnectionInfoMessage, NetworkFlowsControlMessage], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &NetworkConnectionInfoService_ServiceDesc.Streams[0], NetworkConnectionInfoService_PushNetworkConnectionInfo_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[NetworkConnectionInfoMessage, NetworkFlowsControlMessage]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type NetworkConnectionInfoService_PushNetworkConnectionInfoClient = grpc.BidiStreamingClient[NetworkConnectionInfoMessage, NetworkFlowsControlMessage]

// NetworkConnectionInfoServiceServer is the server API for NetworkConnectionInfoService service.
// All implementations should embed UnimplementedNetworkConnectionInfoServiceServer
// for forward compatibility.
//
// A Sensor service that allows Collector to send NetworkConnectionInfo messages
type NetworkConnectionInfoServiceServer interface {
	// Note: the response is a stream due to a bug in the C++ GRPC client library. The server is not expected to
	// send anything via this stream.
	PushNetworkConnectionInfo(grpc.BidiStreamingServer[NetworkConnectionInfoMessage, NetworkFlowsControlMessage]) error
}

// UnimplementedNetworkConnectionInfoServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNetworkConnectionInfoServiceServer struct{}

func (UnimplementedNetworkConnectionInfoServiceServer) PushNetworkConnectionInfo(grpc.BidiStreamingServer[NetworkConnectionInfoMessage, NetworkFlowsControlMessage]) error {
	return status.Errorf(codes.Unimplemented, "method PushNetworkConnectionInfo not implemented")
}
func (UnimplementedNetworkConnectionInfoServiceServer) testEmbeddedByValue() {}

// UnsafeNetworkConnectionInfoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkConnectionInfoServiceServer will
// result in compilation errors.
type UnsafeNetworkConnectionInfoServiceServer interface {
	mustEmbedUnimplementedNetworkConnectionInfoServiceServer()
}

func RegisterNetworkConnectionInfoServiceServer(s grpc.ServiceRegistrar, srv NetworkConnectionInfoServiceServer) {
	// If the following call pancis, it indicates UnimplementedNetworkConnectionInfoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NetworkConnectionInfoService_ServiceDesc, srv)
}

func _NetworkConnectionInfoService_PushNetworkConnectionInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NetworkConnectionInfoServiceServer).PushNetworkConnectionInfo(&grpc.GenericServerStream[NetworkConnectionInfoMessage, NetworkFlowsControlMessage]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type NetworkConnectionInfoService_PushNetworkConnectionInfoServer = grpc.BidiStreamingServer[NetworkConnectionInfoMessage, NetworkFlowsControlMessage]

// NetworkConnectionInfoService_ServiceDesc is the grpc.ServiceDesc for NetworkConnectionInfoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkConnectionInfoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sensor.NetworkConnectionInfoService",
	HandlerType: (*NetworkConnectionInfoServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PushNetworkConnectionInfo",
			Handler:       _NetworkConnectionInfoService_PushNetworkConnectionInfo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "internalapi/sensor/network_connection_iservice.proto",
}
