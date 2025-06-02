package utils

import "golang.org/x/crypto/bcrypt"

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	// cost 值决定了哈希计算的成本，它主要影响哈希计算的时间
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes)
}

// BcryptCheck 对比数据库的哈希值和明文密码
func BcryptCheck(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
