package common

import (
	"guide_go/src/errorx"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	bcryptSalt int
}

var cfg Config

func InitConfig() error {
	godotenv.Load()

	salt, err := strconv.Atoi(os.Getenv("SALT"))
	if err != nil {
		return errorx.Env_ERROR_SALT
	}
	cfg.bcryptSalt = salt

	return nil
}
