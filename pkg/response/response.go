package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code   int                 `json:"code"`
	Msg    string              `json:"msg"`
	Data   interface{}         `json:"data"`
	Errors map[string][]string `json:"errors,omitempty"`
}

type UploadFileResponse struct {
	Path string `json:"path"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 200,
		Msg:  "成功",
		Data: data,
	})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{ // Usually return 200 OK with error code in body, or match HTTP status
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// ValidationError handles parameter validation failures (HTTP 422)
func ValidationError(c *gin.Context, errors map[string][]string) {
	c.JSON(http.StatusUnprocessableEntity, Response{
		Code:   422,
		Msg:    "参数校验失败",
		Data:   nil,
		Errors: errors,
	})
}

// ErrorWithStatus returns error with specific HTTP status code
func ErrorWithStatus(c *gin.Context, httpStatus int, code int, msg string) {
	c.JSON(httpStatus, Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
