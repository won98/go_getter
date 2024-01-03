package dto

import (
	"guide_go/src/domain/entity"
	"time"
)

type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Nickname     string    `json:"nickname"`
	Profile      string    `json:"profile"`
	Platform     string    `json:"platform"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type Authentication struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Email        string `json:"email" gorm:"email"`
	Profile      string `json:"profile" gorm:"profile"`
	NickName     string `json:"nickname" gorm:"nickname"`
}

func (u *User) TransferDomain() *entity.User {
	userDomain := &entity.User{}

	userDomain.ID = u.ID
	userDomain.Email = u.Email
	userDomain.Password = u.Password
	userDomain.Profile = u.Profile
	userDomain.Platform = u.Platform
	userDomain.Nickname = u.Nickname
	userDomain.CreatedAt = time.Now()
	userDomain.UpdatedAt = time.Now()
	userDomain.RefreshToken = u.RefreshToken
	return userDomain
}

type CheckUser struct {
	U *entity.User
}

type UserInfo struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Nickname  string    `json:"nickname"`
	Profile   string    `json:"profile"`
	Platform  string    `json:"platform"`
	CreatedAt time.Time `json:"createdAt"`
}

type PasswordChange struct {
	Email       string `json:"email"`
	NewPassword string `json:"newpassword"`
	OldPassword string `json:"oldpassword"`
}
