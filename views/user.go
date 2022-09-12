package views

import (
	"BlueBell/db/mysql"
	"BlueBell/logic"
	"BlueBell/models"
	"BlueBell/utils/snowflake"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(c *gin.Context) {
	// 1.获取用户传递的参数
	SignUpParams := new(models.SignUpParams)
	_ = c.ShouldBindJSON(&SignUpParams)
	// 用户输入参数校验
	isSuccess, msg := logic.SignUpParamCheck(SignUpParams)
	if !isSuccess {
		c.JSON(http.StatusOK, gin.H{
			"code": 600,
			"msg":  msg,
		})
		return
	}

	// 2.业务处理
	if err := mysql.CheckUserExist(SignUpParams); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 700,
			"msg":  "用户已存在",
		})
		return
	}

	user := models.User{
		UserId:   snowflake.GenID(),
		Username: SignUpParams.Username,
		Password: SignUpParams.Password,
	}
	if err := mysql.CreateUser(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 700,
			"msg":  err,
		})
		return
	}

	// 3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
