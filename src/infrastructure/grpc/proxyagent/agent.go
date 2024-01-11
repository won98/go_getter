package proxyagent

import "guide_go/src/infrastructure/grpc"

var V1 *grpc.GrpcServer

func Register(proxystub *grpc.GrpcServer) {
	V1 = &grpc.GrpcServer{}
	V1 = proxystub
}
