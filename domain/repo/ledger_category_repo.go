package repo

import (
	"github.com/feilongjump/jigsaw-api/domain/entity"
)

type LedgerCategoryRepo interface {
	Create(category *entity.LedgerCategory) error
	Update(id, userID uint64, category *entity.LedgerCategory) error
	Delete(id, userID uint64) (error, int64)
	GetLedgerCategory(id, userID uint64) (*entity.LedgerCategory, error)
	// FindLedgerCategories 获取用户所有分类（包括系统公共分类）
	FindLedgerCategories(userID uint64) ([]*entity.LedgerCategory, error)
	// FindChildren 获取直接子分类
	FindChildren(parentID uint64) ([]*entity.LedgerCategory, error)
}
