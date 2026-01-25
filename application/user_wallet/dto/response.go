package dto

import (
	"github.com/dromara/carbon/v2"
	"github.com/feilongjump/jigsaw-api/domain/entity"
)

type UserWalletResp struct {
	ID          uint64          `json:"id"`
	UserID      uint64          `json:"user_id"`
	Name        string          `json:"name"`
	Type        uint8           `json:"type"`
	Balance     float64         `json:"balance"`
	Liability   float64         `json:"liability"`
	Remark      string          `json:"remark"`
	Sort        int             `json:"sort"`
	IsHidden    bool            `json:"is_hidden"`
	ExtraConfig entity.JSONMap  `json:"extra_config"`
	CreatedAt   *carbon.DateTime `json:"created_at"`
}

type UserWalletListResponse struct {
	Data  []*UserWalletResp `json:"data"`
	Total int64             `json:"total"`
	Page  int               `json:"page"`
	Size  int               `json:"size"`
}
