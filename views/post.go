package views

import (
	"BlueBell/db/mysql"
	"BlueBell/models"
	"BlueBell/utils/snowflake"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func CreateOnePost(c *gin.Context) {
	// 获取参数
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("GetOnePost Params error")
		ResponseError(c, CodeServerBusy)
		return
	}
	// 写入数据库
	p.PostID = snowflake.GenID()
	p.UserID, _ = GetCurrentUserID(c)
	if err := mysql.CreateOnePost(p); err != nil {
		fmt.Println(err)
		zap.L().Error("Create Post failed! ")
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, "success")
}

func DeleteOnePost(c *gin.Context) {
	// 获取postid
	postIDStr := c.Param("postid")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		zap.L().Error(err.Error())
		ResponseError(c, CodeInvalidParam)
		return
	}
	//数据查询
	err = mysql.DeleteOnePost(postID)
	if err != nil {
		zap.L().Error(err.Error())
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, "success")
}
