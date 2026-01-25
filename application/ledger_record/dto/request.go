package dto

import "github.com/dromara/carbon/v2"

type CreateLedgerRecordReq struct {
	Type           uint8            `json:"type" binding:"required,oneof=1 2 3"` // 1:支出, 2:收入, 3:转账
	Amount         float64          `json:"amount" binding:"required,gt=0"`
	SourceWalletID uint64           `json:"source_wallet_id"` // 支出账户 / 转出账户
	TargetWalletID uint64           `json:"target_wallet_id"` // 收入账户 / 转入账户
	CategoryID     uint64           `json:"category_id"`      // 分类ID
	OccurredAt     *carbon.DateTime `json:"occurred_at"`      // 记账时间
	Remark         string           `json:"remark" binding:"max=255"`
	Images         []string         `json:"images"`
}

type UpdateLedgerRecordReq struct {
	Type           uint8            `json:"type" binding:"required,oneof=1 2 3"`
	Amount         float64          `json:"amount" binding:"required,gt=0"`
	SourceWalletID uint64           `json:"source_wallet_id"`
	TargetWalletID uint64           `json:"target_wallet_id"`
	CategoryID     uint64           `json:"category_id"`
	OccurredAt     *carbon.DateTime `json:"occurred_at"`
	Remark         string           `json:"remark" binding:"max=255"`
	Images         []string         `json:"images"`
}

type ListLedgerRecordReq struct {
	Page       int    `form:"page,default=1"`
	PageSize   int    `form:"page_size,default=20"`
	Type       uint8  `form:"type"`
	WalletID   uint64 `form:"wallet_id"`
	CategoryID uint64 `form:"category_id"`
	// StartTime, EndTime...
}

type LedgerRecordURIReq struct {
	ID uint64 `uri:"id" binding:"required,min=1"`
}
