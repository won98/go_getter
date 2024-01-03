package service

import (
	"fmt"
	"guide_go/src/common"
	"guide_go/src/domain/repositories"
	"guide_go/src/infrastructure"
	"guide_go/src/transport/api/dto"

	"github.com/teris-io/shortid"
)

type UserService struct {
	*baseService
	userRepo repositories.UserRepository
}

func (svc *UserService) SignUp(user *dto.User) (*dto.Authentication, error) {
	userUniId, err := shortid.Generate()
	if err != nil {
		return nil, err
	}
	user.ID = userUniId
	hash, err := common.GeneratePasswordHash(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = string(hash)

	// token, err := infrastructure.CreateToken(userUniId)
	// if err != nil {
	// 	return nil, err
	// }
	// rtoken, err := infrastructure.CreateRefreshToken(userUniId)
	// if err != nil {
	// 	return nil, err
	// }
	token, rtoken, err := infrastructure.CreateAllToken(userUniId)
	if err != nil {
		return nil, err
	}
	user.RefreshToken = rtoken
	userobj := user.TransferDomain()
	// err = svc.userRepo.UpdateToken(rtoken)
	if err := svc.userRepo.CreateUser(userobj); err != nil {
		return nil, err
	}

	return &dto.Authentication{
		AccessToken:  token,
		RefreshToken: rtoken,
		Email:        user.Email,
		Profile:      user.Profile,
		NickName:     user.Nickname,
	}, nil
}

func (svc *UserService) SignIn(user *dto.User) (*dto.Authentication, error) {
	fmt.Println(user)
	u, err := svc.userRepo.CheckUserById(user)
	fmt.Println("u", u)
	if err != nil {
		return nil, err
	}
	err = common.ComparePasswordWithHash(u.U.Password, user.Password)
	if err != nil {
		return nil, err
	}
	token, rtoken, err := infrastructure.CreateAllToken(u.U.ID)
	if err != nil {
		return nil, err
	}

	return &dto.Authentication{
		AccessToken:  token,
		RefreshToken: rtoken,
		Email:        u.U.Email,
		Profile:      u.U.Profile,
		NickName:     u.U.Nickname,
	}, nil
}

func (svc *UserService) UpdateMypage(Id string, UserInfo *dto.UserInfo) error {

	return svc.userRepo.UpdateMypage(Id, UserInfo)
}

func (svc *UserService) ChangePassword(Id string, Password *dto.PasswordChange) error {
	hash, err := common.GeneratePasswordHash(Password.NewPassword)
	if err != nil {
		return err
	}
	Password.NewPassword = string(hash)
	return svc.userRepo.ChangePassword(Id, Password)
}
