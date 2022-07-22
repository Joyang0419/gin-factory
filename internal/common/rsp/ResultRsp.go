package rsp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResultRsp struct {
	code int         `json:"code"`
	msg  string      `json:"success"`
	data interface{} `json:"data"`
}

func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
		"data": data,
	})
}

func Error(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest,
		gin.H{
			"error": msg,
		})
}
