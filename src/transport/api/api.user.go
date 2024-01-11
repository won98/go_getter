package api

import (
	"fmt"
	"guide_go/src/infrastructure"
	"guide_go/src/infrastructure/grpc"
	"guide_go/src/internal/authpb"
	"guide_go/src/transport/api/dto"
	"log"
)

func (l *APIStandardLauncher) SignUp(c ApiContext) error {
	fmt.Println("api들어옴")
	var (
		user  dto.User
		err   error
		reply = make(map[string]interface{})
	)

	err = c.bind(&user)
	// fmt.Println(err)
	if err != nil {
		log.Printf("Error binding")
		reply["error"] = err.Error()
		return c.Reply(reply)
	}
	fmt.Println("123", 123)

	services := l.ServicePool.Get()
	fmt.Println("services", services)
	defer l.ServicePool.Return(services)

	retval, err := services.UserService.SignUp(&user)
	if err != nil {
		return err
	}

	return c.Reply(retval)
}

func (l *APIStandardLauncher) SignIn(c ApiContext) error {
	var (
		user  dto.User
		err   error
		reply = make(map[string]interface{})
	)
	err = c.bind(&user)
	if err != nil {
		reply["error"] = err.Error()
		return c.Reply(reply)
	}
	services := l.ServicePool.Get()
	defer l.ServicePool.Return(services)
	retval, err := services.UserService.SignIn(&user)
	if err != nil {
		reply["error"] = err.Error()
		return c.Reply(reply)
	}

	return c.Reply(retval)
}

func (l *APIStandardLauncher) EmailCheck(c ApiContext) error {
	var (
		user  dto.User
		err   error
		reply = make(map[string]interface{})
	)
	err = c.bind(&user)
	if err != nil {
		reply["error"] = err.Error()
		return c.Reply(reply)
	}
	cnt, err := l.Mysql.UserRepositoryImpl.EmailCheck(user.Email)
	if err != nil {
		reply["error"] = err.Error()
		return c.Reply(reply)
	}

	return c.Reply(cnt)
}

func (l *APIStandardLauncher) AutoLogin(c ApiContext) error {

	var (
		err       error
		reply     = make(map[string]interface{})
		user      dto.User
		rtoken    string
		token     string
		grpcToken *authpb.JwtToken
	)

	fmt.Println("123")
	refreshToken, err := infrastructure.VerifyRefresh(c.GetReTokenString())
	fmt.Println("c.GetGrpcReTokenString() : ", c.GetGrpcReTokenString())
	//grpc token server
	grpcServer := grpc.NewGrpcServer(l.Env)
	proxyClient := grpc.NewProxyAuthClient(grpcServer)
	grpcrefreshToken, rpcerr := proxyClient.VerifyJwtRefresh(c.GetGrpcReTokenString())
	if err != nil {
		return err
	}
	if rpcerr != nil {
		return rpcerr
	}
	fmt.Println(refreshToken)
	fmt.Println("grpcrefreshToken", grpcrefreshToken.Id)
	fmt.Println("grpcrefreshTokenrpcerr", rpcerr)
	if refreshToken == "NULL" {
		reply["error"] = "Invalid refresh token"
		return c.Reply(reply)
	}
	if grpcrefreshToken == nil {
		reply["error"] = "Invalid refresh token"
		return c.Reply(reply)
	}
	err = c.bind(&user)
	if err != nil {
		reply["error"] = err.Error()
		return c.Reply(reply)
	}
	userPk, err := l.Mysql.UserRepositoryImpl.CheckUserById(&user)

	if err != nil {
		reply["error"] = err.Error()
		return c.Reply(reply)
	}
	if refreshToken == userPk.U.ID {
		token, rtoken, err = infrastructure.CreateAllToken(userPk.U.ID)
		if err != nil {
			reply["error"] = err.Error()
			return c.Reply(reply)
		}

	}
	if grpcrefreshToken.Id == userPk.U.ID {
		grpcToken, err = proxyClient.LocalT2Issuer(&authpb.JwtSecure{
			Id:                userPk.U.ID,
			JwtIssuanceStatus: authpb.JwtIssuanceStatus_BOTH_ISSUANCE,
		})
		if err != nil {
			reply["error"] = err.Error()
			return c.Reply(reply)
		}

	}
	err = l.Mysql.UserRepositoryImpl.UpdateRefresh(rtoken, userPk.U.Email)
	if err != nil {
		reply["error"] = "Update failed: " + err.Error()
		return c.Reply(reply)
	}
	return c.Reply(dto.Authentication{
		Email:            userPk.U.Email,
		Profile:          userPk.U.Profile,
		NickName:         userPk.U.Nickname,
		AccessToken:      token,
		RefreshToken:     rtoken,
		GrpcAccessToken:  grpcToken.Authorization,
		GrpcRefreshToken: grpcToken.RefreshAuthorization,
	})
}

func (l *APIStandardLauncher) Refresh(c ApiContext) error {
	var (
		err   error
		reply = make(map[string]interface{})
		// user      dto.User
		rtoken    string
		token     string
		grpcToken *authpb.JwtToken
	)
	refreshToken, err := infrastructure.VerifyRefresh(c.GetReTokenString())
	grpcServer := grpc.NewGrpcServer(l.Env)
	proxyClient := grpc.NewProxyAuthClient(grpcServer)
	grpcrefreshToken, rpcerr := proxyClient.VerifyJwtRefresh(c.GetGrpcReTokenString())
	if err != nil {
		return err
	}
	if rpcerr != nil {
		return rpcerr
	}
	if refreshToken == "NULL" {
		reply["error"] = "Invalid refresh token"
		return c.Reply(reply)
	}
	if grpcrefreshToken == nil {
		reply["error"] = "Invalid refresh token"
		return c.Reply(reply)
	}
	userPk, err := l.Mysql.UserRepositoryImpl.CheckUserByIdAndRefresh(refreshToken, c.GetReTokenString())
	if err != nil {
		reply["error"] = "record not found"
		return c.Reply(reply)
	}
	token, rtoken, err = infrastructure.CreateAllToken(userPk.U.ID)
	if err != nil {
		reply["error"] = "token generation failed"
		return c.Reply(reply)
	}
	grpcToken, err = proxyClient.LocalT2Issuer(&authpb.JwtSecure{
		Id:                userPk.U.ID,
		JwtIssuanceStatus: authpb.JwtIssuanceStatus_BOTH_ISSUANCE,
	})
	if err != nil {
		reply["error"] = "Grpc_token generation failed"
		return c.Reply(reply)
	}
	err = l.Mysql.UserRepositoryImpl.UpdateRefresh(rtoken, userPk.U.Email)
	if err != nil {
		reply["error"] = "Update failed: " + err.Error()
		return c.Reply(reply)
	}
	retval := &dto.Authentication{
		AccessToken:      token,
		RefreshToken:     rtoken,
		GrpcAccessToken:  grpcToken.Authorization,
		GrpcRefreshToken: grpcToken.RefreshAuthorization,
	}
	reply["result"] = retval
	return c.Reply(reply)
}

func (l *APIStandardLauncher) Mypage(c ApiContext) error {
	var (
		err   error
		reply = make(map[string]interface{})
	)
	token, err := infrastructure.VerifyToken(c.GetTokenString())

	if err != nil {
		reply["error"] = err.Error()
		return c.Reply(reply)
	}
	//grpc token server
	grpcServer := grpc.NewGrpcServer(l.Env)
	proxyClient := grpc.NewProxyAuthClient(grpcServer)
	grpcreaccessToken, rpcerr := proxyClient.VerifyJwtAccess(c.GetGrpcTokenString())
	if rpcerr != nil {
		return rpcerr
	}
	fmt.Println("grpcreaccessToken : ", grpcreaccessToken)
	row, err := l.Mysql.UserRepositoryImpl.MyPage(token)
	if err != nil {
		reply["error"] = err.Error()
		return c.Reply(reply)
	}
	return c.Reply(row)
}

func (l *APIStandardLauncher) UpdateMypage(c ApiContext) error {
	var (
		err      error
		reply    = make(map[string]interface{})
		userInfo dto.UserInfo
	)
	token, err := infrastructure.VerifyToken(c.GetTokenString())
	if err != nil {
		reply["error"] = err.Error()
		return c.Reply(reply)
	}
	grpcServer := grpc.NewGrpcServer(l.Env)
	proxyClient := grpc.NewProxyAuthClient(grpcServer)
	grpcreaccessToken, rpcerr := proxyClient.VerifyJwtAccess(c.GetGrpcTokenString())
	fmt.Println("grpcreaccessToken : ", grpcreaccessToken)
	if rpcerr != nil {
		return rpcerr
	}
	err = c.bind(&userInfo)
	if err != nil {
		reply["error"] = err.Error()
		return c.Reply(reply)
	}
	services := l.ServicePool.Get()
	defer l.ServicePool.Return(services)
	err = services.UserService.UpdateMypage(token, &userInfo)
	if err != nil {
		reply["error"] = err.Error()
		return c.Reply(reply)
	}
	reply["result"] = true
	return c.Reply(reply)
}

func (l *APIStandardLauncher) ChangePassword(c ApiContext) error {
	var (
		err         error
		reply       = make(map[string]interface{})
		newPassword dto.PasswordChange
	)
	token, err := infrastructure.VerifyToken(c.GetTokenString())
	if err != nil {
		reply["error"] = err.Error()
		return c.Reply(reply)
	}
	//grpc token server
	grpcServer := grpc.NewGrpcServer(l.Env)
	proxyClient := grpc.NewProxyAuthClient(grpcServer)
	grpcreaccessToken, rpcerr := proxyClient.VerifyJwtAccess(c.GetGrpcTokenString())
	fmt.Println("grpcreaccessToken : ", grpcreaccessToken)
	if rpcerr != nil {
		return rpcerr
	}
	err = c.bind(&newPassword)
	if err != nil {
		reply["error"] = err.Error()
		return c.Reply(reply)
	}
	services := l.ServicePool.Get()
	defer l.ServicePool.Return(services)
	err = services.UserService.ChangePassword(token, &newPassword)
	if err != nil {
		reply["error"] = err.Error()
		return c.Reply(reply)
	}
	reply["result"] = true
	return c.Reply(reply)
}
