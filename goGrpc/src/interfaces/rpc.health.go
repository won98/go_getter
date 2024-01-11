package interfaces

import (
	"context"

	"google.golang.org/grpc/health/grpc_health_v1"
)

func (r HealthRpcServer) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	retval := &grpc_health_v1.HealthCheckResponse{}
	retval.Status = grpc_health_v1.HealthCheckResponse_SERVING
	return retval, nil
}

func (r HealthRpcServer) Watch(*grpc_health_v1.HealthCheckRequest, grpc_health_v1.Health_WatchServer) error {
	return nil
}

// func (r *HealthRpcServer) Ping(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
// 	return &emptypb.Empty{}, nil
// }
