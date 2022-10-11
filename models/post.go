package models

import "time"

type Post struct {
	PostID      int64     `json:"id" db:"post_id"`
	UserID      int64     `json:"user_id" db:"user_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}
