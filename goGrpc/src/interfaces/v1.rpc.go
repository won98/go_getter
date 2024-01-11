package interfaces

import (
	"fmt"
	"grpc_go/src/internal/authpb"
	"net"
	"os"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func (r RpcServerBase) ServeRpcV1() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", r.Env.RpcBase.Port))
	if err != nil {
		os.Exit(0)
	}
	logrus.ErrorKey = "grpc.error"
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())
	opts := []grpc.ServerOption{grpc.ChainUnaryInterceptor(
		grpc_recovery.UnaryServerInterceptor(),
		grpc_logrus.UnaryServerInterceptor(logrusEntry),
	),
		grpc.ChainStreamInterceptor(
			grpc_recovery.StreamServerInterceptor(),
			grpc_logrus.StreamServerInterceptor(logrusEntry),
		),
	}
	grpcServer := grpc.NewServer(opts...)
	grpc_health_v1.RegisterHealthServer(grpcServer, r.HealthCheck)
	authpb.RegisterAuthGrpcServiceServer(grpcServer, r.AuthService)
	fmt.Printf(" Rpc Start on %d \n", r.Env.RpcBase.Port)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			os.Exit(0)
		}
	}()
}
