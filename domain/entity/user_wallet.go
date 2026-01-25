package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/dromara/carbon/v2"
	"gorm.io/gorm"
)

type JSONMap map[string]interface{}

const (
	UserWalletTypeCash        uint8 = 1
	UserWalletTypeBankCard    uint8 = 2
	UserWalletTypeWeChat      uint8 = 3
	UserWalletTypeAlipay      uint8 = 4
	UserWalletTypeCreditCard  uint8 = 5
	UserWalletTypeStoredValue uint8 = 6
	UserWalletTypeInvestment  uint8 = 7
	UserWalletTypeMargin      uint8 = 8
)

func (j *JSONMap) Scan(value interface{}) error {
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

func (j JSONMap) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

type UserWallet struct {
	ID          uint64           `gorm:"primarykey" json:"id"`
	UserID      uint64           `gorm:"index;not null" json:"user_id"`
	Name        string           `gorm:"type:varchar(100);not null" json:"name"`
	Type        uint8            `gorm:"type:tinyint;not null;default:1" json:"type"`
	Balance     float64          `gorm:"type:decimal(15,2);default:0" json:"balance"`
	Liability   float64          `gorm:"type:decimal(15,2);default:0" json:"liability"`
	ExtraConfig JSONMap          `gorm:"type:json" json:"extra_config"`
	Sort        int              `gorm:"type:int;default:0" json:"sort"`
	Remark      string           `gorm:"type:varchar(255)" json:"remark"`
	IsHidden    bool             `gorm:"type:tinyint(1);default:0" json:"is_hidden"`
	CreatedAt   *carbon.DateTime `gorm:"type:datetime" json:"created_at"`
	UpdatedAt   *carbon.DateTime `gorm:"type:datetime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt   `gorm:"index" json:"-"`
}
