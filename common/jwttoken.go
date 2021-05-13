package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kuludi/kuludi-gin-vue/model"
	"time"
)

//设置过期时间
const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("kuludi")

type MyClaims struct {
	UserId uint
	jwt.StandardClaims
}

//生成token
func ReleaseToken(user *model.User) (string, error) {

	claims := &MyClaims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "red", //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(MySecret)

}

//解析token
func ParseToken(tokenString string) (*jwt.Token, *MyClaims, error) {
	claims := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})

	return token, claims, err

}
