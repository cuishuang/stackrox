// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api/v2/compliance_results_stats_service.proto

package v2

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
	ComplianceResultsStatsService_GetComplianceProfileStats_FullMethodName        = "/v2.ComplianceResultsStatsService/GetComplianceProfileStats"
	ComplianceResultsStatsService_GetComplianceProfilesStats_FullMethodName       = "/v2.ComplianceResultsStatsService/GetComplianceProfilesStats"
	ComplianceResultsStatsService_GetComplianceProfileCheckStats_FullMethodName   = "/v2.ComplianceResultsStatsService/GetComplianceProfileCheckStats"
	ComplianceResultsStatsService_GetComplianceClusterScanStats_FullMethodName    = "/v2.ComplianceResultsStatsService/GetComplianceClusterScanStats"
	ComplianceResultsStatsService_GetComplianceOverallClusterStats_FullMethodName = "/v2.ComplianceResultsStatsService/GetComplianceOverallClusterStats"
	ComplianceResultsStatsService_GetComplianceClusterStats_FullMethodName        = "/v2.ComplianceResultsStatsService/GetComplianceClusterStats"
)

// ComplianceResultsStatsServiceClient is the client API for ComplianceResultsStatsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComplianceResultsStatsServiceClient interface {
	// GetComplianceProfileStats lists current scan stats grouped by profile
	// Optional RawQuery query fields can be combined.
	// Commonly used ones include but are not limited to
	// - scan: id(s) of the compliance scan
	// - cluster: id(s) of the cluster
	// - profile: id(s) of the profile
	GetComplianceProfileStats(ctx context.Context, in *ComplianceProfileResultsRequest, opts ...grpc.CallOption) (*ListComplianceProfileScanStatsResponse, error)
	// GetComplianceProfileScanStats lists current scan stats grouped by profile
	// Optional RawQuery query fields can be combined.
	// Commonly used ones include but are not limited to
	// - scan: id(s) of the compliance scan
	// - cluster: id(s) of the cluster
	// - profile: id(s) of the profile
	GetComplianceProfilesStats(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*ListComplianceProfileScanStatsResponse, error)
	// GetComplianceProfileCheckStats lists current stats for a specific cluster check
	GetComplianceProfileCheckStats(ctx context.Context, in *ComplianceProfileCheckRequest, opts ...grpc.CallOption) (*ListComplianceProfileResults, error)
	// GetComplianceClusterScanStats lists current scan stats grouped by cluster
	// Optional RawQuery query fields can be combined.
	// Commonly used ones include but are not limited to
	// - scan: id(s) of the compliance scan
	// - cluster: id(s) of the cluster
	// - profile: id(s) of the profile
	GetComplianceClusterScanStats(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*ListComplianceClusterScanStatsResponse, error)
	// Deprecated in favor of GetComplianceClusterStats
	GetComplianceOverallClusterStats(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*ListComplianceClusterOverallStatsResponse, error)
	GetComplianceClusterStats(ctx context.Context, in *ComplianceProfileResultsRequest, opts ...grpc.CallOption) (*ListComplianceClusterOverallStatsResponse, error)
}

type complianceResultsStatsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComplianceResultsStatsServiceClient(cc grpc.ClientConnInterface) ComplianceResultsStatsServiceClient {
	return &complianceResultsStatsServiceClient{cc}
}

func (c *complianceResultsStatsServiceClient) GetComplianceProfileStats(ctx context.Context, in *ComplianceProfileResultsRequest, opts ...grpc.CallOption) (*ListComplianceProfileScanStatsResponse, error) {
	out := new(ListComplianceProfileScanStatsResponse)
	err := c.cc.Invoke(ctx, ComplianceResultsStatsService_GetComplianceProfileStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complianceResultsStatsServiceClient) GetComplianceProfilesStats(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*ListComplianceProfileScanStatsResponse, error) {
	out := new(ListComplianceProfileScanStatsResponse)
	err := c.cc.Invoke(ctx, ComplianceResultsStatsService_GetComplianceProfilesStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complianceResultsStatsServiceClient) GetComplianceProfileCheckStats(ctx context.Context, in *ComplianceProfileCheckRequest, opts ...grpc.CallOption) (*ListComplianceProfileResults, error) {
	out := new(ListComplianceProfileResults)
	err := c.cc.Invoke(ctx, ComplianceResultsStatsService_GetComplianceProfileCheckStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complianceResultsStatsServiceClient) GetComplianceClusterScanStats(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*ListComplianceClusterScanStatsResponse, error) {
	out := new(ListComplianceClusterScanStatsResponse)
	err := c.cc.Invoke(ctx, ComplianceResultsStatsService_GetComplianceClusterScanStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complianceResultsStatsServiceClient) GetComplianceOverallClusterStats(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*ListComplianceClusterOverallStatsResponse, error) {
	out := new(ListComplianceClusterOverallStatsResponse)
	err := c.cc.Invoke(ctx, ComplianceResultsStatsService_GetComplianceOverallClusterStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complianceResultsStatsServiceClient) GetComplianceClusterStats(ctx context.Context, in *ComplianceProfileResultsRequest, opts ...grpc.CallOption) (*ListComplianceClusterOverallStatsResponse, error) {
	out := new(ListComplianceClusterOverallStatsResponse)
	err := c.cc.Invoke(ctx, ComplianceResultsStatsService_GetComplianceClusterStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComplianceResultsStatsServiceServer is the server API for ComplianceResultsStatsService service.
// All implementations should embed UnimplementedComplianceResultsStatsServiceServer
// for forward compatibility
type ComplianceResultsStatsServiceServer interface {
	// GetComplianceProfileStats lists current scan stats grouped by profile
	// Optional RawQuery query fields can be combined.
	// Commonly used ones include but are not limited to
	// - scan: id(s) of the compliance scan
	// - cluster: id(s) of the cluster
	// - profile: id(s) of the profile
	GetComplianceProfileStats(context.Context, *ComplianceProfileResultsRequest) (*ListComplianceProfileScanStatsResponse, error)
	// GetComplianceProfileScanStats lists current scan stats grouped by profile
	// Optional RawQuery query fields can be combined.
	// Commonly used ones include but are not limited to
	// - scan: id(s) of the compliance scan
	// - cluster: id(s) of the cluster
	// - profile: id(s) of the profile
	GetComplianceProfilesStats(context.Context, *RawQuery) (*ListComplianceProfileScanStatsResponse, error)
	// GetComplianceProfileCheckStats lists current stats for a specific cluster check
	GetComplianceProfileCheckStats(context.Context, *ComplianceProfileCheckRequest) (*ListComplianceProfileResults, error)
	// GetComplianceClusterScanStats lists current scan stats grouped by cluster
	// Optional RawQuery query fields can be combined.
	// Commonly used ones include but are not limited to
	// - scan: id(s) of the compliance scan
	// - cluster: id(s) of the cluster
	// - profile: id(s) of the profile
	GetComplianceClusterScanStats(context.Context, *RawQuery) (*ListComplianceClusterScanStatsResponse, error)
	// Deprecated in favor of GetComplianceClusterStats
	GetComplianceOverallClusterStats(context.Context, *RawQuery) (*ListComplianceClusterOverallStatsResponse, error)
	GetComplianceClusterStats(context.Context, *ComplianceProfileResultsRequest) (*ListComplianceClusterOverallStatsResponse, error)
}

// UnimplementedComplianceResultsStatsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedComplianceResultsStatsServiceServer struct {
}

func (UnimplementedComplianceResultsStatsServiceServer) GetComplianceProfileStats(context.Context, *ComplianceProfileResultsRequest) (*ListComplianceProfileScanStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComplianceProfileStats not implemented")
}
func (UnimplementedComplianceResultsStatsServiceServer) GetComplianceProfilesStats(context.Context, *RawQuery) (*ListComplianceProfileScanStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComplianceProfilesStats not implemented")
}
func (UnimplementedComplianceResultsStatsServiceServer) GetComplianceProfileCheckStats(context.Context, *ComplianceProfileCheckRequest) (*ListComplianceProfileResults, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComplianceProfileCheckStats not implemented")
}
func (UnimplementedComplianceResultsStatsServiceServer) GetComplianceClusterScanStats(context.Context, *RawQuery) (*ListComplianceClusterScanStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComplianceClusterScanStats not implemented")
}
func (UnimplementedComplianceResultsStatsServiceServer) GetComplianceOverallClusterStats(context.Context, *RawQuery) (*ListComplianceClusterOverallStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComplianceOverallClusterStats not implemented")
}
func (UnimplementedComplianceResultsStatsServiceServer) GetComplianceClusterStats(context.Context, *ComplianceProfileResultsRequest) (*ListComplianceClusterOverallStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComplianceClusterStats not implemented")
}

// UnsafeComplianceResultsStatsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComplianceResultsStatsServiceServer will
// result in compilation errors.
type UnsafeComplianceResultsStatsServiceServer interface {
	mustEmbedUnimplementedComplianceResultsStatsServiceServer()
}

func RegisterComplianceResultsStatsServiceServer(s grpc.ServiceRegistrar, srv ComplianceResultsStatsServiceServer) {
	s.RegisterService(&ComplianceResultsStatsService_ServiceDesc, srv)
}

func _ComplianceResultsStatsService_GetComplianceProfileStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComplianceProfileResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceResultsStatsServiceServer).GetComplianceProfileStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComplianceResultsStatsService_GetComplianceProfileStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceResultsStatsServiceServer).GetComplianceProfileStats(ctx, req.(*ComplianceProfileResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComplianceResultsStatsService_GetComplianceProfilesStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceResultsStatsServiceServer).GetComplianceProfilesStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComplianceResultsStatsService_GetComplianceProfilesStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceResultsStatsServiceServer).GetComplianceProfilesStats(ctx, req.(*RawQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComplianceResultsStatsService_GetComplianceProfileCheckStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComplianceProfileCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceResultsStatsServiceServer).GetComplianceProfileCheckStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComplianceResultsStatsService_GetComplianceProfileCheckStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceResultsStatsServiceServer).GetComplianceProfileCheckStats(ctx, req.(*ComplianceProfileCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComplianceResultsStatsService_GetComplianceClusterScanStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceResultsStatsServiceServer).GetComplianceClusterScanStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComplianceResultsStatsService_GetComplianceClusterScanStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceResultsStatsServiceServer).GetComplianceClusterScanStats(ctx, req.(*RawQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComplianceResultsStatsService_GetComplianceOverallClusterStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceResultsStatsServiceServer).GetComplianceOverallClusterStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComplianceResultsStatsService_GetComplianceOverallClusterStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceResultsStatsServiceServer).GetComplianceOverallClusterStats(ctx, req.(*RawQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComplianceResultsStatsService_GetComplianceClusterStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComplianceProfileResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceResultsStatsServiceServer).GetComplianceClusterStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComplianceResultsStatsService_GetComplianceClusterStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceResultsStatsServiceServer).GetComplianceClusterStats(ctx, req.(*ComplianceProfileResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ComplianceResultsStatsService_ServiceDesc is the grpc.ServiceDesc for ComplianceResultsStatsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComplianceResultsStatsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v2.ComplianceResultsStatsService",
	HandlerType: (*ComplianceResultsStatsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetComplianceProfileStats",
			Handler:    _ComplianceResultsStatsService_GetComplianceProfileStats_Handler,
		},
		{
			MethodName: "GetComplianceProfilesStats",
			Handler:    _ComplianceResultsStatsService_GetComplianceProfilesStats_Handler,
		},
		{
			MethodName: "GetComplianceProfileCheckStats",
			Handler:    _ComplianceResultsStatsService_GetComplianceProfileCheckStats_Handler,
		},
		{
			MethodName: "GetComplianceClusterScanStats",
			Handler:    _ComplianceResultsStatsService_GetComplianceClusterScanStats_Handler,
		},
		{
			MethodName: "GetComplianceOverallClusterStats",
			Handler:    _ComplianceResultsStatsService_GetComplianceOverallClusterStats_Handler,
		},
		{
			MethodName: "GetComplianceClusterStats",
			Handler:    _ComplianceResultsStatsService_GetComplianceClusterStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v2/compliance_results_stats_service.proto",
}
