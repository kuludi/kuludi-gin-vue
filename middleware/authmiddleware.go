package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kuludi/kuludi-gin-vue/common"
	"github.com/kuludi/kuludi-gin-vue/db"
	"github.com/kuludi/kuludi-gin-vue/model"
	"net/http"
	"strings"
)

const Sailt = "Bearer"

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusUnauthorized, "msg": "token 为空"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if !(len(parts) == 2 && parts[0] == Sailt) {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusUnauthorized, "msg": "token 格式错误"})
			c.Abort()
			return
		}

		token, claims, err := common.ParseToken(parts[1])

		if err != nil || !token.Valid {

			c.JSON(http.StatusOK, gin.H{"code": http.StatusUnauthorized, "msg": "token 无效"})
			c.Abort()
			return
		}
		user := &model.User{}
		db.DB.Where("id = ?", claims.UserId).First(&user)
		c.Set("user", user)
		c.Next()

	}
}
