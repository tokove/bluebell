package mysql

import (
	"bluebell_backend/model"
	"errors"
)

func CheckUserExist(username string) error {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return nil
}

func InsertUser(user *model.User) error {
	sqlStr := `insert into user(user_id, username, password) values(?, ?, ?)`
	_, err := db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return err
}
