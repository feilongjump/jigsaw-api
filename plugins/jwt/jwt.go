package jwt

import (
	"errors"
	"github.com/feilongjump/jigsaw-api/config"
	"github.com/feilongjump/jigsaw-api/utils"
	jwtpkg "github.com/golang-jwt/jwt/v5"
	"time"
)

// JWTCustomClaims 自定义载荷
type JWTCustomClaims struct {
	UserID uint64 `json:"user_id"`

	jwtpkg.RegisteredClaims
}

// GenerateToken 生成 jwt 令牌
func GenerateToken(UserID uint64) (string, error) {

	appConfig := config.GetAppConfig()
	jwtConfig := config.GetJWTConfig()

	expiresAt := time.Now().Add(time.Second * time.Duration(jwtConfig.Expires))

	claims := JWTCustomClaims{
		UserID,
		jwtpkg.RegisteredClaims{
			Issuer:    appConfig.Name,                                          // 签发者
			Subject:   "API Token",                                             // 签发主题
			Audience:  jwtpkg.ClaimStrings{appConfig.Name},                     // 签发受众
			ExpiresAt: jwtpkg.NewNumericDate(expiresAt),                        // 过期时间
			NotBefore: jwtpkg.NewNumericDate(time.Now().Add(time.Microsecond)), // 最早使用时间
			IssuedAt:  jwtpkg.NewNumericDate(time.Now()),                       // 签发时间
			ID:        utils.GenerateRandomString(12),                          // wt ID, 类似于盐值
		},
	}

	// 使用特定的加密方式进行加密
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)

	// 使用指定的 secret 进行签名加密
	return token.SignedString([]byte(jwtConfig.Secret))
}

// ParseToken 解析 jwt 令牌
func ParseToken(tokenString string) (*JWTCustomClaims, error) {

	jwtConfig := config.GetJWTConfig()

	token, err := jwtpkg.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(t *jwtpkg.Token) (interface{}, error) {
		return []byte(jwtConfig.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// GenerateSecret 生成 jwt secret
func GenerateSecret() string {
	return utils.GenerateRandomString(32)
}
