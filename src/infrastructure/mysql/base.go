package mysql

import (
	"guide_go/src/domain/repositories"

	"gorm.io/gorm"
)

type BaseRepository struct {
	ORMMysql *gorm.DB
}

type Repository struct {
	*BaseRepository
	UserRepositoryImpl *UserRepositoryImpl
}

type UserRepositoryImpl struct {
	*BaseRepository
}

var _ repositories.UserRepository = (*UserRepositoryImpl)(nil)
