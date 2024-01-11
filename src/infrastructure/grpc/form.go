package grpc

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func SendToken(ctx context.Context, token string) context.Context {
	return SendData(ctx, "Authorization", token)
}

func SendRefreshToken(ctx context.Context, token string) context.Context {
	return SendData(ctx, "RefreshAuthorization", token)
}

func SendData(ctx context.Context, key, value string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, key, value)
}
