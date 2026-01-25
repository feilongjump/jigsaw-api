package dto

import "github.com/dromara/carbon/v2"

type LedgerCategoryResp struct {
	ID        uint64           `json:"id"`
	UserID    uint64           `json:"user_id"`
	ParentID  uint64           `json:"parent_id"`
	Type      uint8            `json:"type"`
	Name      string           `json:"name"`
	Icon      string           `json:"icon"`
	Sort      int              `json:"sort"`
	Path      string           `json:"path"`
	CreatedAt *carbon.DateTime `json:"created_at"`
	Children  []*LedgerCategoryResp `json:"children,omitempty"`
}
