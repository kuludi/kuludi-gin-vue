package utils

import "net/http"

const (
	succ = http.StatusOK
)

type Reply struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
	Data interface{} `json:"data"`
}

func NewReply(code int, msg string, data ...interface{}) (int, Reply) {
	if len(data) > 0 {
		return 200, Reply{
			Code: code,
			Msg:  msg,
			Data: data[0],
		}
	}
	return 200, Reply{
		Code: code,
		Msg:  msg,
	}
}

func Success(msg string, data ...interface{}) (int, Reply) {
	return NewReply(succ, msg, data...)
}
