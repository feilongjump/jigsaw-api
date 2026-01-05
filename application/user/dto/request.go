package dto

type RegisterRequest struct {
	Username string `json:"username" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required,min=6" label:"密码"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required,min=6" label:"密码"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required,min=6" label:"旧密码"`
	NewPassword string `json:"new_password" binding:"required,min=6" label:"新密码"`
}
