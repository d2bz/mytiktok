package utils

import "github.com/gin-gonic/gin"

type reJson struct {
	Code int
	Msg  string
	Data any
}

func Response(c *gin.Context, code int, msg string, data any) {
	Json := reJson{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.JSON(code, Json)
}
