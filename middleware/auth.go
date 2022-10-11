package middleware

import (
	"BlueBell/utils/jwt"
	"BlueBell/views"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			views.ResponseError(c, views.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按照空格分割，当匹配到一个空格时，后续的内容作为一个整体。组合成一个含2个元素的切片
		splitContent := strings.SplitN(authHeader, " ", 2)
		if !(len(splitContent) == 2 && splitContent[0] == "Bearer") {
			views.ResponseError(c, views.CodeInvalidToken)
			c.Abort() // 打断后续的中间件
			return
		}

		// 解析token
		mc, err := jwt.ParseToken(splitContent[1])
		fmt.Println(mc, err)
		if err != nil {
			views.ResponseError(c, views.CodeInvalidToken)
			c.Abort()
			return
		}
		// 设置全局的用户信息
		c.Set(views.RequestUserIDKey, mc.UserID)
		c.Next() //执行后续的中间件（函数）
	}
}
