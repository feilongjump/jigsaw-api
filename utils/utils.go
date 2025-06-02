package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateRandomString 生成随机加密字符串
func GenerateRandomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
