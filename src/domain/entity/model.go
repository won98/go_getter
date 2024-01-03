package entity

import "time"

type BaseModel struct {
	ID string `json:"id"  gorm:"id"`
	CommonModel
}

type CommonModel struct {
	CreatedAt time.Time `json:"created_at"  gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at"  gorm:"updated_at"`
}

func (c CommonModel) Local() CommonModel {
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	return c
}
