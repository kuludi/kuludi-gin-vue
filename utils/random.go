package utils

import (
	"math/rand"
	"time"
)

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
