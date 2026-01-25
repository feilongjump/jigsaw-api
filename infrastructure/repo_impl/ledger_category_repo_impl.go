package repo_impl

import (
	"errors"

	"github.com/feilongjump/jigsaw-api/domain/entity"
	"github.com/feilongjump/jigsaw-api/domain/repo"
	"github.com/feilongjump/jigsaw-api/infrastructure/db"
	"github.com/feilongjump/jigsaw-api/pkg/err_code"
	"gorm.io/gorm"
)

type ledgerCategoryRepoImpl struct {
	db *gorm.DB
}

func NewLedgerCategoryRepo() repo.LedgerCategoryRepo {
	return &ledgerCategoryRepoImpl{db: db.Get()}
}

func (r *ledgerCategoryRepoImpl) Create(category *entity.LedgerCategory) error {
	return r.db.Create(category).Error
}

func (r *ledgerCategoryRepoImpl) Update(id, userID uint64, category *entity.LedgerCategory) error {
	return r.db.
		Where("user_id = ?", userID).
		Model(&entity.LedgerCategory{ID: id}).
		Select("Name", "Icon", "Sort", "ParentID", "Type", "Path").
		Updates(category).
		Error
}

func (r *ledgerCategoryRepoImpl) Delete(id, userID uint64) (error, int64) {
	result := r.db.
		Where("user_id = ?", userID).
		Delete(&entity.LedgerCategory{ID: id})

	return result.Error, result.RowsAffected
}

func (r *ledgerCategoryRepoImpl) GetLedgerCategory(id, userID uint64) (*entity.LedgerCategory, error) {
	var category entity.LedgerCategory
	// 允许查询属于该用户的分类 或 系统公共分类(user_id=0)
	err := r.db.
		Where("id = ?", id).
		Where("user_id = ? OR user_id = ?", userID, 0).
		First(&category).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err_code.LedgerCategoryNotFound
		}
		return nil, err
	}
	return &category, nil
}

func (r *ledgerCategoryRepoImpl) FindLedgerCategories(userID uint64) ([]*entity.LedgerCategory, error) {
	var categories []*entity.LedgerCategory
	err := r.db.
		Where("user_id = ? OR user_id = ?", userID, 0).
		Order("sort desc, id asc").
		Find(&categories).Error
	return categories, err
}

func (r *ledgerCategoryRepoImpl) FindChildren(parentID uint64) ([]*entity.LedgerCategory, error) {
	var categories []*entity.LedgerCategory
	err := r.db.
		Where("parent_id = ?", parentID).
		Order("sort desc, id asc").
		Find(&categories).Error
	return categories, err
}
