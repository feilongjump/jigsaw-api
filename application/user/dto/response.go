package dto

import "github.com/dromara/carbon/v2"

type LoginResponse struct {
	Token string `json:"token"`
}

type MeResponse struct {
	Username  string           `json:"username"`
	Avatar    string           `json:"avatar"`
	CreatedAt *carbon.DateTime `json:"created_at"`
}
