package mysql

import (
	"BlueBell/models"
	"database/sql"
	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select id, name, sort_idx from community order by sort_idx"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

func GetCommunityDetailByID(cid int64) (communityDetail *models.CommunityDetail, err error) {
	communityDetail = new(models.CommunityDetail)
	sqlStr := "select id, name, description, sort_idx, create_time from community where id = ?"
	err = db.Get(communityDetail, sqlStr, cid)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("not found community by cid")
			return
		}
	}
	return
}
