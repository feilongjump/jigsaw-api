package requests

type LoginRequest struct {
	Username string `json:"username" binding:"required,min=5,max=16"`
	Password string `json:"password" binding:"required,min=6"`
}

func (lr *LoginRequest) GetErrMessage(str string) string {
	errMsg := map[string]string{
		"Username.required": "用户名不能为空",
		"Username.min":      "用户名长度不能小于 5",
		"Username.max":      "用户名长度不能大于 16",
		"Password.required": "密码不能为空",
		"Password.min":      "密码长度不能小于 6",
	}

	return errMsg[str]
}

type SignUpRequest struct {
	Username string `json:"username" binding:"required,min=5,max=16,c_unique_db=users:name"`
	Email    string `json:"email" binding:"required,email,c_unique_db=users:email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (sur *SignUpRequest) GetErrMessage(str string) string {
	errMsg := map[string]string{
		"Username.required":    "用户名不能为空",
		"Username.min":         "用户名长度不能小于 5",
		"Username.max":         "用户名长度不能大于 16",
		"Username.c_unique_db": "用户名已存在",
		"Email.required":       "邮箱不能为空",
		"Email.email":          "邮箱格式不正确",
		"Email.c_unique_db":    "邮箱已存在",
		"Password.required":    "密码不能为空",
		"Password.min":         "密码长度不能小于 6",
	}

	return errMsg[str]
}
