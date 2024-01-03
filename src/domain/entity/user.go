package entity

type User struct {
	BaseModel
	Email        string `json:"email" gorm:"email"`
	Password     string `json:"password" gorm:"password"`
	Nickname     string `json:"nickname" gorm:"nickname"`
	Profile      string `json:"profile" gorm:"profile"`
	Platform     string `json:"platform" gorm:"platform"`
	RefreshToken string `json:"refresh_token" gorm:"refresh_token"`
}

func (u *User) TableName() string {
	return "tbl_user"
}
