package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kuludi/kuludi-gin-vue/utils"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	//获取参数

	name := c.PostForm("name")
	password := c.PostForm("password")
	phone := c.PostForm("phone")

	//数据验证
	if len(phone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "422",
			"msg":  "Phone num at least 11",
		})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "422",
			"msg":  "Password can not less than 6",
		})
		return
	}
	if len(name) == 0 {
		name = RandomString(10)
	}
	//创建用户
	log.Println(name, password, phone)
	//返回结果
	c.JSON(utils.Success("注册成功", gin.H{
		"msg": "success",
	}))
}
//生成随机字符串
func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	res := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for v, _ := range res {
		res[v] = letters[rand.Intn(len(letters))]
	}

	return string(res)

}
