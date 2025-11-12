package mysql

import "errors"

var (
	ErrorInvalidID       = errors.New("非法的ID")
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorNotLogin        = errors.New("用户未登录")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)
