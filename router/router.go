package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kuludi/kuludi-gin-vue/controller"
)

func Run() {
	r := gin.Default()

	r.GET("/ping", controller.Ping)
	r.POST("/api/user/register", controller.Register)

	r.Run(":8080")
}
