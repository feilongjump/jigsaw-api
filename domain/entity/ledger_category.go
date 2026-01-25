package entity

import (
	"github.com/dromara/carbon/v2"
	"gorm.io/gorm"
)

const (
	LedgerCategoryTypeExpense uint8 = 1
	LedgerCategoryTypeIncome  uint8 = 2
)

type LedgerCategory struct {
	ID        uint64           `gorm:"primarykey" json:"id"`
	UserID    uint64           `gorm:"index;default:0" json:"user_id"` // 0 for system default
	ParentID  uint64           `gorm:"index;default:0" json:"parent_id"`
	Type      uint8            `gorm:"type:tinyint;not null" json:"type"` // 1: expense, 2: income
	Name      string           `gorm:"type:varchar(50);not null" json:"name"`
	Icon      string           `gorm:"type:varchar(255)" json:"icon"`
	Path      string           `gorm:"type:varchar(255);index" json:"path"` // e.g. "0-1-5"
	Sort      int              `gorm:"type:int;default:0" json:"sort"`
	CreatedAt *carbon.DateTime `gorm:"type:datetime" json:"created_at"`
	UpdatedAt *carbon.DateTime `gorm:"type:datetime" json:"updated_at"`
	DeletedAt gorm.DeletedAt   `gorm:"index" json:"-"`
}
