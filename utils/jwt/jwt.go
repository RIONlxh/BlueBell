package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

var jwtSecret = []byte(viper.GetString("jwtSecret"))

/*
	jwt.StandardClaims 是一个结构体，里面包含了一些字段信息。
	MyClaims 就相当于在原有的基础上增加了一些自定义字段来实现用户认证的可扩展性

*/

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenToken(userID int64, username string) (string, error) {
	// 实例化jwt认证信息
	claims := MyClaims{
		userID,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(
				time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour).Unix(),
			Issuer: "Rion",
		},
	}
	// 生成一个token对象，采用hs256加密方式
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用一个密钥返回一个token
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	// new 一个结构体对象，返回一个指针
	var mc = new(MyClaims)
	// 解析这个tokenstring，并将内容传递到结构体中
	// 第三个参数的function用来传递一个签名字符串
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return viper.GetString("auth.jwt_secret"), nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("token is invalid")
}
