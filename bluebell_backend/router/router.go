package router

import (
	"bluebell_backend/controller"
	"bluebell_backend/logger"

	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)

	r.NoRoute(func(c *gin.Context) {
		controller.ResponseErrorWithMsg(c, controller.CodeNotFound, "404")
	})
	return r
}
