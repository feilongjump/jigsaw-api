package middleware

import (
	"net/http"
	"strings"

	"jigsaw-api/pkg/response"
	"jigsaw-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.ErrorWithStatus(c, http.StatusUnauthorized, 401, "缺少 Authorization 请求头")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && strings.EqualFold(parts[0], "Bearer")) {
			response.ErrorWithStatus(c, http.StatusUnauthorized, 401, "Authorization 格式必须为 Bearer {token}")
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			response.ErrorWithStatus(c, http.StatusUnauthorized, 401, "无效或已过期的 Token")
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
