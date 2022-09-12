package mysql

import (
	"BlueBell/models"
	"errors"
)

// 检测用户是否存在
func CheckUserExist(su *models.SignUpParams) (err error) {
	var count int
	sql := "select count(user_id) from user where username = ?"
	if err := db.Get(&count, sql, su.Username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

// 创建一个用户
func CreateUser(u *models.User) (err error) {
	sql := "INSERT INTO user(user_id,username,password) VALUES (?,?,?)"
	userId := u.UserId
	username := u.Username
	password := u.Password
	_, err = db.Exec(sql, userId, username, password)
	if err != nil {
		return err
	}
	return err
}
