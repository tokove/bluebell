package mysql

import (
	"bluebell_backend/model"
	"database/sql"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*model.Community, err error) {
	sqlStr := `select community_id, community_name from community`
	if err = db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

func GetCommunityDetailByID(id uint64) (communityDetail *model.CommunityDetail, err error) {
	communityDetail = new(model.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time from community where community_id = ?`

	if err = db.Get(communityDetail, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = ErrorInvalidID
		}
	}
	return
}
