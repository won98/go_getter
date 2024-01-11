package grpc

import (
	"context"
	"fmt"
	"guide_go/src/internal/authpb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func securityAgent(conn *grpc.ClientConn) authpb.AuthGrpcServiceClient {
	return authpb.NewAuthGrpcServiceClient(conn)
}

type proxyAuthClient struct {
	*GrpcServer
}

func NewProxyAuthClient(server *GrpcServer) *proxyAuthClient {
	return &proxyAuthClient{
		GrpcServer: server,
	}
}

func (agent *proxyAuthClient) Ping() error {
	// 여기서는 커넥션 풀을 사용하지 않으므로, 바로 Auth 필드의 커넥션을 사용합니다.
	conn := agent.Auth
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	_, err := securityAgent(conn).Ping(context.Background(), &emptypb.Empty{})
	fmt.Println("err", err)
	if err != nil {
		return err
	}
	return nil
}

func (agent *proxyAuthClient) LocalT2Issuer(req *authpb.JwtSecure) (*authpb.JwtToken, error) {
	// 여기서도 커넥션 풀을 사용하지 않으므로, 바로 Auth 필드의 커넥션을 사용합니다.
	conn := agent.Auth
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	retval, err := securityAgent(conn).LocalT2Issuer(ctx, req)
	if err != nil {
		go agent.Ping()
		return nil, err
	}
	if retval.ErrorStatus {
		return nil, err
	}
	return retval, nil
}

func (agent *proxyAuthClient) VerifyJwtRefresh(rtoken string) (*authpb.DeJwtSecure, error) {
	// fmt.Println("rtoken : ", rtoken)/
	conn := agent.Auth
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	ctx, cancel := context.WithTimeout(SendRefreshToken(context.Background(), rtoken), time.Second*10)
	defer cancel()
	retval, err := securityAgent(conn).VerifyJwtRefresh(ctx, &emptypb.Empty{})
	fmt.Println("retval", retval)
	fmt.Println("retvalerr", err)
	if err != nil {
		go agent.Ping()
		return nil, err
	}
	if retval.ErrorStatus {
		return nil, err
	}
	return retval, nil
}
func (agent *proxyAuthClient) VerifyJwtAccess(token string) (*authpb.DeJwtSecure, error) {
	// fmt.Println("rtoken : ", rtoken)/
	conn := agent.Auth
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	ctx, cancel := context.WithTimeout(SendToken(context.Background(), token), time.Second*10)
	defer cancel()
	retval, err := securityAgent(conn).VerifyJwtAccess(ctx, &emptypb.Empty{})
	fmt.Println("retval", retval)
	fmt.Println("retvalerr", err)
	if err != nil {
		go agent.Ping()
		return nil, err
	}
	if retval.ErrorStatus {
		return nil, err
	}
	return retval, nil
}
