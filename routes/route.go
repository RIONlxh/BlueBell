package routes

import (
	"BlueBell/logger"
	"BlueBell/middleware"
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

	v1.POST("sign_up", views.SignUp)
	v1.POST("login", views.Login)
	v1.Use(middleware.AuthLogin())
	{
		v1.GET("admin", views.AdminIndex)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "Page Not Found",
		})
	})
	return r
}
