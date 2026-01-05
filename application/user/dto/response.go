package dto

import "github.com/dromara/carbon/v2"

type LoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	ID        uint64           `json:"id"`
	Username  string           `json:"username"`
	CreatedAt *carbon.DateTime `json:"created_at"`
	UpdatedAt *carbon.DateTime `json:"updated_at"`
}
