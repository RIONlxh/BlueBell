package views

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseResponse struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"Data"`
}

func ResponseSuccess(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &BaseResponse{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &BaseResponse{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &BaseResponse{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
