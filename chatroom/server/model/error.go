package model
import (
	"errors"
)

var (
	ERROR_USER_NOT_EXIT = errors.New("用户不存在")
	ERROR_USERPWD = errors.New("密码错误")
	ERROR_USER_EXIT = errors.New("用户已经存在")
)