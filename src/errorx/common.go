package errorx

import "errors"

var (
	Env_ERROR_SALT error = errors.New("env Vars \"SALT\" Load Error")
)
