package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

func HandleError(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
	})
}
