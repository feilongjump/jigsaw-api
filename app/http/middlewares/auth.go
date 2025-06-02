package middlewares

import (
	"github.com/feilongjump/jigsaw-api/app/http/responses"
	"github.com/feilongjump/jigsaw-api/plugins/jwt"
	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt/v5"
	"strings"
)

type AuthUser struct {
	UserID uint64
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authUser := AuthUser{}

		checkAuth(c, &authUser)

		c.Set("user_id", authUser.UserID)
	}
}

// 检查用户是否登录
func checkAuth(ctx *gin.Context, authUser *AuthUser) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		responses.Unauthorized(ctx, "请先登录")
		return
	}

	tokenSplit := strings.Fields(token)
	// 令牌类型是否正确
	if tokenSplit[0] != "Bearer" {
		responses.Unauthorized(ctx, "请先登录账号！")
		return
	}

	// 令牌可否正常解析
	claims, err := jwt.ParseToken(tokenSplit[1])
	if err != nil {
		if strings.Contains(err.Error(), jwtpkg.ErrTokenExpired.Error()) {
			responses.Unauthorized(ctx, "登录已过期，请重新登录")
			return
		}

		responses.Unauthorized(ctx, err.Error())
		return
	}

	authUser.UserID = claims.UserID
}
