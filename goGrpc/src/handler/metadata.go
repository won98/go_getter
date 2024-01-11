package handler

import (
	"context"
	"errors"
	"fmt"
	"grpc_go/src/application"

	"google.golang.org/grpc/metadata"
)

func (h *ServiceHandler) LocalAFromMetaIncomingContext(ctx context.Context) (*application.AccessPayload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, nil
	}
	token := md["authorization"]
	if len(token) == 0 {
		return nil, errors.New("authorization token not provided")
	}
	userNoStr, err := application.VerifyToken(token[0])
	if err != nil {
		return nil, err
	}

	ap := &application.AccessPayload{
		Id: userNoStr,
	}

	return ap, nil
}

func (h *ServiceHandler) LocalRFromMetaIncomingContext(ctx context.Context) (*application.RefreshPayload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Println("md : ", md)
	if !ok {
		return nil, nil
	}
	token := md["refreshauthorization"]
	fmt.Println("token : ", token)
	if len(token) == 0 {
		return nil, errors.New("RefreshAuthorization token not provided")
	}
	userNoStr, err := application.VerifyRefreshToken(token[0])
	if err != nil {
		return nil, err
	}

	ap := &application.RefreshPayload{
		Id: userNoStr,
	}

	return ap, nil
}
