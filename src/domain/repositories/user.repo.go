package repositories

import (
	"guide_go/src/domain/entity"
	"guide_go/src/transport/api/dto"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	CheckUserById(user *dto.User) (*dto.CheckUser, error)
	UpdateMypage(Id string, UserInfo *dto.UserInfo) error
	ChangePassword(Id string, PasswordChange *dto.PasswordChange) error
}
