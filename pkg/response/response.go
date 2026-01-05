package response

import (
	"errors"
	"net/http"

	"github.com/feilongjump/jigsaw-api/pkg/err_code"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Data   any    `json:"data,omitempty"`
	Errors any    `json:"errors,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code: err_code.Success.Code,
		Msg:  err_code.Success.Msg,
		Data: data,
	})
}

// Fail 失败响应
func Fail(c *gin.Context, err error) {
	var e err_code.ErrCode
	if errors.As(err, &e) {
		// 自定义错误信息
		c.JSON(http.StatusOK, Response{
			Code: e.Code,
			Msg:  e.Msg,
		})
		return
	}

	// 默认作为服务器内部错误处理
	c.JSON(http.StatusOK, Response{
		Code: err_code.SystemException.Code,
		Msg:  err_code.SystemException.Msg,
	})
}

// ValidateFail 验证失败响应
func ValidateFail(c *gin.Context, errors any) {
	c.JSON(http.StatusOK, Response{
		Code:   err_code.ValidationFailed.Code,
		Msg:    err_code.ValidationFailed.Msg,
		Errors: errors,
	})
}

// FailWithDetail 自定义失败响应信息
func FailWithDetail(c *gin.Context, err err_code.ErrCode, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: err.Code,
		Msg:  msg,
	})
}
