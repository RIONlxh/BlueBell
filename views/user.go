package views

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(c *gin.Context) {
	// 1.获取用户传递的参数

	// 2.业务处理

	// 3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}
