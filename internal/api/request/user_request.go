package request

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20" label:"用户名"`
	Password string `json:"password" binding:"required,min=6" label:"密码"`
	Email    string `json:"email" binding:"required,email" label:"邮箱"`
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
	Avatar string `json:"avatar" label:"头像"`
	Email  string `json:"email" binding:"omitempty,email" label:"邮箱"`
}

// RegisterRequest 用户注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20" label:"用户名"`
	Password string `json:"password" binding:"required,min=6" label:"密码"`
	Email    string `json:"email" binding:"required,email" label:"邮箱"`
}

// LoginRequest 用户登录请求
type LoginRequest struct {
	Account  string `json:"account" binding:"required" label:"账号"`
	Password string `json:"password" binding:"required" label:"密码"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required" label:"旧密码"`
	NewPassword string `json:"new_password" binding:"required,min=6" label:"新密码"`
}
