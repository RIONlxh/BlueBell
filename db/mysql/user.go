package mysql

import (
	"BlueBell/models"
	"BlueBell/utils"
	"BlueBell/utils/jwt"
	"database/sql"
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
	u.Password = utils.Md5Encrypt(u.Password)
	_, err = db.Exec(sql, u.UserId, u.Username, u.Password)
	if err != nil {
		return err
	}
	return
}

func UserLogin(u *models.User) (token string, err error) {
	md5Password := utils.Md5Encrypt(u.Password)
	sqlStr := `select user_id,username,password from user where username = ?`
	err = db.Get(u, sqlStr, u.Username)

	// 用户查询校验
	if err == sql.ErrNoRows {
		return "", errors.New("用户名不存在")
	}
	if err != nil {
		return "", errors.New("查询失败")
	}
	if md5Password != u.Password {
		return "", errors.New("用户名或密码错误")
	}

	// 生成用户 token
	return jwt.GenToken(u.UserId, u.Username)

}
