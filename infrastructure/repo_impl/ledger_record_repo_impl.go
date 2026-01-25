package repo_impl

import (
	"errors"

	"github.com/feilongjump/jigsaw-api/domain/entity"
	"github.com/feilongjump/jigsaw-api/domain/repo"
	"github.com/feilongjump/jigsaw-api/infrastructure/db"
	"github.com/feilongjump/jigsaw-api/pkg/err_code"
	"gorm.io/gorm"
)

type ledgerRecordRepoImpl struct {
	db *gorm.DB
}

func NewLedgerRecordRepo() repo.LedgerRecordRepo {
	return &ledgerRecordRepoImpl{db: db.Get()}
}

func (r *ledgerRecordRepoImpl) Create(record *entity.LedgerRecord) error {
	return r.db.Create(record).Error
}

func (r *ledgerRecordRepoImpl) Update(id, userID uint64, record *entity.LedgerRecord) error {
	return r.db.
		Where("user_id = ?", userID).
		Model(&entity.LedgerRecord{ID: id}).
		Select("Type", "Amount", "SourceWalletID", "TargetWalletID", "CategoryID", "OccurredAt", "Remark", "Images").
		Updates(record).
		Error
}

func (r *ledgerRecordRepoImpl) Delete(id, userID uint64) (error, int64) {
	result := r.db.
		Where("user_id = ?", userID).
		Delete(&entity.LedgerRecord{ID: id})

	return result.Error, result.RowsAffected
}

func (r *ledgerRecordRepoImpl) GetLedgerRecord(id, userID uint64) (*entity.LedgerRecord, error) {
	var record entity.LedgerRecord
	err := r.db.
		Where("id = ? AND user_id = ?", id, userID).
		First(&record).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err_code.LedgerRecordNotFound
		}
		return nil, err
	}
	return &record, nil
}

func (r *ledgerRecordRepoImpl) FindLedgerRecords(userID uint64, page, pageSize int, filter map[string]interface{}) ([]*entity.LedgerRecord, int64, error) {
	var records []*entity.LedgerRecord
	var total int64

	db := r.db.Model(&entity.LedgerRecord{}).Where("user_id = ?", userID)

	if v, ok := filter["type"]; ok {
		db = db.Where("type = ?", v)
	}
	if v, ok := filter["wallet_id"]; ok {
		db = db.Where("source_wallet_id = ? OR target_wallet_id = ?", v, v)
	}
	if v, ok := filter["category_id"]; ok {
		db = db.Where("category_id = ?", v)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := db.
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order("created_at desc").
		Find(&records).
		Error
	if err != nil {
		return nil, 0, err
	}
	return records, total, nil
}

func (r *ledgerRecordRepoImpl) CountByWalletID(userID, walletID uint64) (int64, error) {
	var total int64
	err := r.db.
		Model(&entity.LedgerRecord{}).
		Where("user_id = ?", userID).
		Where("source_wallet_id = ? OR target_wallet_id = ?", walletID, walletID).
		Count(&total).
		Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
