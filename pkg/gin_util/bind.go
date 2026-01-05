package gin_util

import (
	"github.com/feilongjump/jigsaw-api/pkg/err_code"
	"github.com/feilongjump/jigsaw-api/pkg/response"
	"github.com/feilongjump/jigsaw-api/pkg/validator"
	"github.com/gin-gonic/gin"
)

// BindJSON 绑定 JSON 请求体到结构体
func BindJSON(c *gin.Context, obj any) bool {
	err := c.ShouldBindJSON(obj)
	if err == nil {
		return true
	}

	handlerBindError(c, err, obj)
	return false
}

// BindQuery 绑定查询参数到结构体
func BindQuery(c *gin.Context, obj any) bool {
	err := c.ShouldBindQuery(obj)
	if err == nil {
		return true
	}

	handlerBindError(c, err, obj)
	return false
}

// BindURI 绑定 URI 参数到结构体
func BindURI(c *gin.Context, obj any) bool {
	err := c.ShouldBindUri(obj)
	if err == nil {
		return true
	}

	handlerBindError(c, err, obj)
	return false
}

// handlerBindError 处理绑定错误
func handlerBindError(c *gin.Context, err error, obj any) {
	// 翻译错误
	errTrans := validator.Translate(err, obj)

	// 参数校验失败
	if len(errTrans) > 0 {
		response.ValidateFail(c, errTrans)
		return
	}

	// 其他错误默认返回 MalformedRequest，但不将详细的错误信息返回给客户端
	response.Fail(c, err_code.MalformedRequest)
}
