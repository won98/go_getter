package interfaces

import (
	"context"
	"fmt"
	"grpc_go/src/application"
	"grpc_go/src/handler"
	"grpc_go/src/internal/authpb"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (r *AuthRpcService) Ping(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (r *AuthRpcService) LocalT1Issuer(ctx context.Context, req *authpb.JwtSecure) (*authpb.JwtToken, error) {
	retval := &authpb.JwtToken{}
	if req.JwtIssuanceStatus == 1 {
		accessToken, err := application.CreateToken(req.Id)
		if err != nil {
			retval.ErrorMessage = err.Error()
		}
		retval.Authorization = accessToken
		return retval, nil
	} else if req.JwtIssuanceStatus == 2 {
		refreshToken, err := application.CreateRefreshToken(req.Id)
		if err != nil {
			retval.ErrorMessage = err.Error()
		}
		retval.RefreshAuthorization = refreshToken
		return retval, nil
	} else {
		retval.ErrorMessage = "에러입니다"
		return retval, nil
	}
}

func (r *AuthRpcService) LocalT2Issuer(ctx context.Context, req *authpb.JwtSecure) (*authpb.JwtToken, error) {
	retval := &authpb.JwtToken{}
	if req.JwtIssuanceStatus == 3 {
		accessToken, err := application.CreateToken(req.Id)
		if err != nil {
			retval.ErrorMessage = err.Error()
			return retval, nil
		}

		refreshToken, err := application.CreateRefreshToken(req.Id)
		if err != nil {
			retval.ErrorMessage = err.Error()
			return retval, nil
		}

		retval.Authorization = accessToken
		retval.RefreshAuthorization = refreshToken
		return retval, nil
	}
	accessToken, err := application.CreateToken(req.Id)
	if err != nil {
		retval.ErrorMessage = err.Error()
		return retval, nil
	}
	refreshToken, err := application.CreateRefreshToken(req.Id)
	if err != nil {
		retval.ErrorMessage = err.Error()
		return retval, nil
	}
	retval.Authorization = accessToken
	retval.RefreshAuthorization = refreshToken
	return retval, nil
}

func (r *AuthRpcService) VerifyJwtAccess(ctx context.Context, req *emptypb.Empty) (*authpb.DeJwtSecure, error) {
	retval := &authpb.DeJwtSecure{}
	h := handler.NewHandlerContext(context.Background(), r.AppLauncher)
	payload, err := h.LocalAFromMetaIncomingContext(ctx)
	if err != nil {
		if err.Error() == NOT_EXISTS_METADATA {
			retval.ErrorMessage = NOT_EXISTS_METADATA
			retval.ErrorStatus = true
			return retval, nil
		}
		if err.Error() == INVALID_TOKEN {
			retval.ErrorMessage = INVALID_TOKEN
			retval.ErrorStatus = true
			return retval, nil
		}
		if err.Error() == EXPIRED_SESSION {
			retval.ErrorMessage = EXPIRED_SESSION
			retval.ErrorStatus = true
			return retval, nil
		}
		retval.ErrorMessage = NOT_DEFINED_ERROR
		retval.ErrorStatus = true
		return retval, nil
	}
	return payload.SendRpc(), nil
}

func (r *AuthRpcService) VerifyJwtRefresh(ctx context.Context, req *emptypb.Empty) (*authpb.DeJwtSecure, error) {
	fmt.Println("sdfsdfsdfsdfsdfsdsdf")
	retval := &authpb.DeJwtSecure{}
	fmt.Println("retval : ", retval)
	h := handler.NewHandlerContext(context.Background(), r.AppLauncher)
	payload, err := h.LocalRFromMetaIncomingContext(ctx)
	fmt.Println("error : ", err)
	if err != nil {
		if err.Error() == NOT_EXISTS_METADATA {
			retval.ErrorMessage = NOT_EXISTS_METADATA
			retval.ErrorStatus = true
			return retval, nil
		}
		if err.Error() == INVALID_TOKEN {
			retval.ErrorMessage = INVALID_TOKEN
			retval.ErrorStatus = true
			return retval, nil
		}
		if err.Error() == EXPIRED_SESSION {
			retval.ErrorMessage = EXPIRED_SESSION
			retval.ErrorStatus = true
			return retval, nil
		}
		retval.ErrorMessage = NOT_DEFINED_ERROR
		retval.ErrorStatus = true
		return retval, nil
	}
	return payload.SendRpc(), nil
}
