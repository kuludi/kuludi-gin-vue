package res

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, httpStatus int, code int, msg string, data gin.H) {
	c.JSON(httpStatus, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}


