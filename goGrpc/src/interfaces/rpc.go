package interfaces

import (
	"grpc_go/src/internal/authpb"
	"grpc_go/src/launch"
	"os"
)

type RpcServerBase struct {
	*launch.AppLauncher
	AuthService *AuthRpcService
	HealthCheck *HealthRpcServer
}

type HealthRpcServer struct {
	RpcServerBase
}

type AuthRpcService struct {
	authpb.UnimplementedAuthGrpcServiceServer
	RpcServerBase
}

func ServeRpc(launcher *launch.AppLauncher) {
	defer func() {
		if r := recover(); r != nil {
			os.Exit(0)
		}
	}()
	r := RpcServerBase{AppLauncher: launcher}
	r.HealthCheck = &HealthRpcServer{RpcServerBase: r}
	r.AuthService = &AuthRpcService{RpcServerBase: r}
	r.ServeRpcV1()
}
