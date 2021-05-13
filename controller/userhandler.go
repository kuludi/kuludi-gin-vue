package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kuludi/kuludi-gin-vue/common"
	"github.com/kuludi/kuludi-gin-vue/dao"
	"github.com/kuludi/kuludi-gin-vue/dto"
	"github.com/kuludi/kuludi-gin-vue/model"
	res "github.com/kuludi/kuludi-gin-vue/response"
	"github.com/kuludi/kuludi-gin-vue/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
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
		name = utils.RandomString(10)
	}
	//创建用户
	if dao.IsExist(phone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "422",
			"msg":  "the user is exist",
		})
		return
	}

	//加密密码
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": "500",
			"msg":  "The password hash error",
		})
		return
	}

	user := &model.User{
		Name:     name,
		Password: string(hashPassword),
		Phone:    phone,
	}
	err = dao.Register(user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "422",
			"msg":  err,
		})

	} else {
		c.JSON(utils.Success("注册成功", gin.H{
			"msg": "success",
		}))

	}

}

func Login(c *gin.Context) {
	phone := c.PostForm("phone")
	password := c.PostForm("password")

	if phone == "" || len(phone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "phone is not null or length is not equal 11",
		})
		return
	}
	if password == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "password is not null",
		})
		return
	}

	if dao.IsExist(phone) == false {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "The user ie not register",
		})
		return
	}
	user, _ := dao.GetUserByPhone(phone)

	//比对密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "phone or password is not match",
		})
		return
	}
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "token error",
		})

		log.Println("token 生成错误: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Login success",
		"data": gin.H{
			"token": token,
		},
	})
}

func Info(c *gin.Context) {
	user ,_:= c.Get("user")
	//gin.H{"user":dto.UserToDto(user.(*model.User))}
	res.Response(c,http.StatusOK,200,"success",gin.H{"user":dto.UserToDto(user.(*model.User))})
}
