package grpc

import (
	"fmt"
	"guide_go/src/domain"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	once     sync.Once // singleton
	instance *GrpcServer
)

type GrpcServer struct {
	Auth *grpc.ClientConn
}

// NewGrpcServer는 GrpcServer의 싱글턴 인스턴스를 생성하고 반환합니다.
func NewGrpcServer(env *domain.Environment) *GrpcServer {
	once.Do(func() {
		server := &GrpcServer{}
		if env.AuthRpc.Enable {
			// GRPC 커넥션을 열고, Auth 필드에 할당합니다.
			conn, err := grpc.Dial(address(env.AuthRpc.Host, env.AuthRpc.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				fmt.Printf("Failed to dial gRPC server: %v\n", err)
				return
			}
			server.Auth = conn
		}
		instance = server
	})
	return instance
}

// address 함수는 호스트명과 포트를 결합하여 주소 문자열을 생성합니다.
func address(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

// 기존의 connectionpool 관련 코드는 제거되었습니다.
