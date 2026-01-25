package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/dromara/carbon/v2"
	"gorm.io/gorm"
)

type JSONStrings []string

func (j *JSONStrings) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, j)
}

func (j JSONStrings) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

type LedgerRecord struct {
	ID             uint64           `gorm:"primarykey" json:"id"`
	UserID         uint64           `gorm:"index;not null" json:"user_id"`
	SourceWalletID uint64           `gorm:"index" json:"source_wallet_id"`           // 支出账户 / 转出账户
	TargetWalletID uint64           `gorm:"index;default:0" json:"target_wallet_id"` // 收入账户 / 转入账户
	CategoryID     uint64           `gorm:"index;default:0" json:"category_id"`
	Type           uint8            `gorm:"type:tinyint;not null" json:"type"` // 1:支出, 2:收入, 3:转账
	Amount         float64          `gorm:"type:decimal(15,2);not null" json:"amount"`
	OccurredAt     *carbon.DateTime `gorm:"type:datetime;not null" json:"occurred_at"`
	Remark         string           `gorm:"type:varchar(255)" json:"remark"`
	Images         JSONStrings      `gorm:"type:json" json:"images"`
	CreatedAt      *carbon.DateTime `gorm:"type:datetime" json:"created_at"`
	UpdatedAt      *carbon.DateTime `gorm:"type:datetime" json:"updated_at"`
	DeletedAt      gorm.DeletedAt   `gorm:"index" json:"-"`
}
