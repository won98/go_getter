package handler

import (
	"context"
	"grpc_go/src/launch"
)

type ServiceHandler struct {
	*launch.AppLauncher
	RequestContext context.Context
}

func NewHandlerContext(ctx context.Context, l *launch.AppLauncher) *ServiceHandler {
	h := &ServiceHandler{}
	h.RequestContext = ctx
	h.AppLauncher = l
	return h
}

func getIpFromCtx(ctx context.Context) string {
	return ""
}
