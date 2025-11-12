package service

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/model"
	"bluebell_backend/pkg/snowflake"
)

func CreatePost(p *model.Post) (err error) {
	// 生成post_id
	p.ID, err = snowflake.GetID()
	if err != nil {
		err = mysql.ErrorInvalidID
		return 
	}
	return mysql.CreatePost(p)
}

func GetPostDetial(id uint64) (*model.Post, error) {
	return mysql.GetPostDetailByID(id)
}