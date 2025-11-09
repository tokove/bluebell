package service

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/model"
	"bluebell_backend/pkg/snowflake"
	"bluebell_backend/pkg/utils"
)

func SignUp(p *model.ParamSignUp) (err error) {
	// 查询用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 生成UID
	var userID uint64
	userID, err = snowflake.GetID()
	if err != nil {
		return err
	}
	// 加密
	user := &model.User{
		UserID:   userID,
		Username: p.Username,
	}
	user.Password, err = utils.HashPassword(p.Password)
	if err != nil {
		return err
	}
	// 保存进数据库
	return mysql.InsertUser(user)
}
