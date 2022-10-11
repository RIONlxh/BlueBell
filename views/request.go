package views

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var RequestUserIDKey = "userID"

// getCurrentUserID 获取当前登录的用户ID
func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(RequestUserIDKey)
	if !ok {
		err = errors.New("用户未登录")
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = errors.New("用户未登录")
		return
	}
	return
}
