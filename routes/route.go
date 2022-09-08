package routes

import (
	"BlueBell/logger"
	"BlueBell/views"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1/")

	// 注册路由
	v1.GET("sign_up", views.SignUp)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "Page Not Found",
		})
	})
	return r
}
