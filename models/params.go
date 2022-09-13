package models

type SignUpParams struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type LoginParams struct {
	Username string `db:"user_id"`
	Password string `db:"password"`
	VCode    string
}
