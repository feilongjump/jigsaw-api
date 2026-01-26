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

// CreditCardConfig 信用卡扩展配置 (Type=5)
type CreditCardConfig struct {
	BillDay      uint8   `json:"bill_day"`      // 账单日 (1-31)
	RepaymentDay uint8   `json:"repayment_day"` // 还款日 (1-31)
	CreditLimit  float64 `json:"credit_limit"`  // 信用额度
}

// InvestmentConfig 投资/两融账户扩展配置 (Type=7, 8)
type InvestmentConfig struct {
	CreditLimit      float64          `json:"credit_limit,omitempty"`      // 授信额度 (仅两融 Type=8)
	MaintenanceRatio float64          `json:"maintenance_ratio,omitempty"` // 维持担保比例 (仅两融 Type=8)
	Rules            []CommissionRule `json:"rules,omitempty"`             // 费率规则列表
}

// CommissionRule 佣金费率规则
type CommissionRule struct {
	Market          string  `json:"market,omitempty"`            // 市场代码 (如: SH, SZ, BJ, HK, US)
	Type            string  `json:"type,omitempty"`              // 品种类型 (如: STOCK, ETF, BOND, FUND, OPTION)
	CommissionRate  float64 `json:"commission_rate"`             // 佣金费率 (如: 0.00025)
	MinCommission   float64 `json:"min_commission"`              // 单笔最低佣金 (如: 5.0)
	StampDutyRate   float64 `json:"stamp_duty_rate,omitempty"`   // 印花税率 (卖出时收取, 如: 0.0005)
	TransferFeeRate float64 `json:"transfer_fee_rate,omitempty"` // 过户费率 (双向收取, 如: 0.00001)
}
