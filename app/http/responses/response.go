package responses

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Success 请求成功
func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}

// SuccessNoContent 请求成功，没有返回实体
func SuccessNoContent(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, nil)
}

// Unauthorized 未授权
func Unauthorized(ctx *gin.Context, msg string) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": msg,
		"data":    nil,
	})
}

// NotFound 数据不存在
func NotFound(ctx *gin.Context, msg string) {
	ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": msg,
		"data":    nil,
	})
}

// ValidatorError 请求参数错误
//
//	{
//	    "errors": {
//	        "email": "邮箱格式错误"
//	    },
//	    "message": "请求参数错误"
//	}
func ValidatorError(ctx *gin.Context, errors map[string]string) {
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": "请求参数错误",
		"errors":  errors,
		"data":    nil,
	})
}

// Abort500 服务器错误
func Abort500(ctx *gin.Context, msg string) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": msg,
		"errors":  nil,
		"data":    nil,
	})
}
