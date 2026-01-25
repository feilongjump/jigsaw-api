package dto

import (
	"github.com/dromara/carbon/v2"
	"github.com/feilongjump/jigsaw-api/domain/entity"
)

type LedgerRecordResp struct {
	ID               uint64             `json:"id"`
	UserID           uint64             `json:"user_id"`
	Type             uint8              `json:"type"`
	Amount           float64            `json:"amount"`
	SourceWalletID   uint64             `json:"source_wallet_id"`
	SourceWalletName string             `json:"source_wallet_name,omitempty"`
	TargetWalletID   uint64             `json:"target_wallet_id"`
	TargetWalletName string             `json:"target_wallet_name,omitempty"`
	CategoryID       uint64             `json:"category_id"`
	CategoryName     string             `json:"category_name,omitempty"`
	OccurredAt       *carbon.DateTime   `json:"occurred_at"`
	Remark           string             `json:"remark"`
	Images           entity.JSONStrings `json:"images"`
	CreatedAt        *carbon.DateTime   `json:"created_at"`
}

type LedgerRecordListResponse struct {
	Data  []*LedgerRecordResp `json:"data"`
	Total int64               `json:"total"`
	Page  int                 `json:"page"`
	Size  int                 `json:"size"`
}
