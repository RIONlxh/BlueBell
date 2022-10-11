package models

import "time"

type Community struct {
	ID       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Sort_Idx string `json:"sort_idx,omitempty" db:"sort_idx"`
}

type CommunityDetail struct {
	ID          int64     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Sort_Idx    string    `json:"sort_idx,omitempty" db:"sort_idx"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}
