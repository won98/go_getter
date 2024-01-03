package mysql

import (
	"fmt"
	"guide_go/src/domain/entity"
	"guide_go/src/transport/api/dto"
	"time"
)

func (my *UserRepositoryImpl) CreateUser(user *entity.User) error {

	if err := my.ORMMysql.Create(user); err != nil {
		return err.Error
	}
	return nil
}
func (my *UserRepositoryImpl) CheckUserById(user *dto.User) (*dto.CheckUser, error) {
	dto := &dto.CheckUser{
		U: &entity.User{},
	}
	err := my.ORMMysql.Model(&entity.User{}).
		Select("*").Where("email = ?", user.Email).First(dto.U).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(err)
	return dto, nil
}

func (my *UserRepositoryImpl) EmailCheck(email string) (int64, error) {
	var cnt int64
	if err := my.ORMMysql.Model(&entity.User{}).Where("email = ?", email).Count(&cnt).Error; err != nil {
		return 0, err
	}
	return cnt, nil
}

func (my *UserRepositoryImpl) MyPage(id string) (*dto.UserInfo, error) {
	req := &dto.UserInfo{}
	if err := my.ORMMysql.Model(&entity.User{}).Where("id =?", id).First(&req).Error; err != nil {
		return nil, err
	}
	return req, nil
}

func (my *UserRepositoryImpl) UpdateMypage(id string, UserInfo *dto.UserInfo) error {

	updateInfo := map[string]interface{}{
		"email":      UserInfo.Email,
		"nickname":   UserInfo.Nickname,
		"profile":    UserInfo.Profile,
		"platform":   UserInfo.Platform,
		"updated_at": time.Now(),
	}
	if err := my.ORMMysql.Model(&entity.User{}).Where("id =?", id).Updates(updateInfo).Error; err != nil {
		return err
	}
	return nil
}

func (my *UserRepositoryImpl) ChangePassword(id string, Password *dto.PasswordChange) error {

	updateInfo := map[string]interface{}{
		"email":      Password.Email,
		"password":   Password.NewPassword,
		"updated_at": time.Now(),
	}
	if err := my.ORMMysql.Model(&entity.User{}).Where("id =?", id).Updates(updateInfo).Error; err != nil {
		return err
	}
	return nil
}
