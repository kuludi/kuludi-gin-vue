package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kuludi/kuludi-gin-vue/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", controller.Ping)
	r.POST("/api/user/register", controller.Register)
	r.POST("/api/user/login", controller.Login)

	return r
}
