package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kuludi/kuludi-gin-vue/utils"
)

func Register(c *gin.Context) {

	c.JSON(utils.Success("注册成功", gin.H{
		"msg": "success",
	}))
}
