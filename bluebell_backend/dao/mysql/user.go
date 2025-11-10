package mysql

import (
	"bluebell_backend/model"
	"bluebell_backend/pkg/errcode"
	"database/sql"
)

func CheckUserExist(username string) error {
	sqlStr := `select count(*) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errcode.ErrorUserExist
	}
	return nil
}

func InsertUser(user *model.User) error {
	sqlStr := `insert into user(user_id, username, password) values(?, ?, ?)`
	_, err := db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return err
}

func GetHashedPasswordByUsername(username string) (string, error) {
	sqlStr := `select password from user where username = ?`
	var password string
	err := db.Get(&password, sqlStr, username)
	if err == sql.ErrNoRows {
		return "", errcode.ErrorUserNotExist
	}
	if err != nil {
		return "", err
	}
	return password, nil
}
