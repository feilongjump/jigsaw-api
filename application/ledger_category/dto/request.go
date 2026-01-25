package dto

type CreateLedgerCategoryReq struct {
	ParentID uint64 `json:"parent_id"`                         // 0表示一级分类
	Type     uint8  `json:"type" binding:"required,oneof=1 2"` // 1:支出, 2:收入
	Name     string `json:"name" binding:"required,max=50"`
	Icon     string `json:"icon" binding:"max=255"`
	Sort     int    `json:"sort"`
}

type UpdateLedgerCategoryReq struct {
	ParentID *uint64 `json:"parent_id"`
	Name     *string `json:"name" binding:"omitempty,max=50"`
	Icon     *string `json:"icon" binding:"omitempty,max=255"`
	Sort     *int    `json:"sort"`
}

type LedgerCategoryURIReq struct {
	ID uint64 `uri:"id" binding:"required,min=1"`
}
