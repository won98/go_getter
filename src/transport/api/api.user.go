package api

import (
	"fmt"
	"guide_go/src/infrastructure"
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
		err    error
		reply  = make(map[string]interface{})
		user   dto.User
		rtoken string
		token  string
	)
	refreshToken, err := infrastructure.VerifyRefresh(c.GetReTokenString())
	if err != nil {
		return err
	}
	fmt.Println(refreshToken)
	if refreshToken == "NULL" {
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
	return c.Reply(dto.Authentication{
		Email:        userPk.U.Email,
		Profile:      userPk.U.Profile,
		NickName:     userPk.U.Nickname,
		AccessToken:  token,
		RefreshToken: rtoken,
	})
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
