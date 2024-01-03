package common

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), cfg.bcryptSalt)
}

func ComparePasswordWithHash(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
