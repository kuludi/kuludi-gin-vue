package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kuludi/kuludi-gin-vue/controller"
)

func main() {
	r := gin.Default()


	r.GET("/ping",controller.Ping)
  r.POST("/")

	r.Run(":8080")

}
