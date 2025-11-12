package service

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/model"
)

func GetCommunityList() ([]*model.Community, error) {
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id uint64) (*model.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}