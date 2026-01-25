package dto

import "github.com/feilongjump/jigsaw-api/domain/entity"

type CreateUserWalletReq struct {
	Name        string         `json:"name" binding:"required,max=100"`
	Type        uint8          `json:"type" binding:"required,wallet_type"`
	Balance     float64        `json:"balance"`
	Liability   float64        `json:"liability"`
	Remark      string         `json:"remark" binding:"max=255"`
	Sort        int            `json:"sort"`
	ExtraConfig entity.JSONMap `json:"extra_config"` // 直接透传 JSON
}

type UpdateUserWalletReq struct {
	Name        *string         `json:"name" binding:"omitempty,max=100"`
	Remark      *string         `json:"remark" binding:"omitempty,max=255"`
	Sort        *int            `json:"sort"`
	IsHidden    *bool           `json:"is_hidden"`
	ExtraConfig *entity.JSONMap `json:"extra_config"`
}

type UserWalletURIReq struct {
	ID uint64 `uri:"id" binding:"required,min=1"`
}
