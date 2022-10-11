package views

import (
	"BlueBell/db/mysql"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetCommunityList(c *gin.Context) {
	data, err := mysql.GetCommunityList()
	if err != nil {
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, data)
}

func GetCommunityDetail(c *gin.Context) {
	cid := c.Param("cid")
	cidInt, _ := strconv.ParseInt(cid, 10, 64)
	data, err := mysql.GetCommunityDetailByID(cidInt)
	if err != nil {
		fmt.Println(err)
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
