package models

type User struct {
	UserId   int64  `db:user_id json:"user_id"`
	Username string `db:username json:"username"`
	Password string `db:password json:"password"`
}
