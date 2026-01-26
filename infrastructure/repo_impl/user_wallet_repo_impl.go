package repo_impl

import (
	"context"
	"errors"

	"github.com/feilongjump/jigsaw-api/domain/entity"
	"github.com/feilongjump/jigsaw-api/domain/repo"
	"github.com/feilongjump/jigsaw-api/infrastructure/db"
	"github.com/feilongjump/jigsaw-api/pkg/err_code"
	"gorm.io/gorm"
)

type userWalletRepoImpl struct {
	db *gorm.DB
}

func NewUserWalletRepo() repo.UserWalletRepo {
	return &userWalletRepoImpl{db: db.Get()}
}

func (r *userWalletRepoImpl) Create(ctx context.Context, wallet *entity.UserWallet) error {
	return db.GetDB(ctx, r.db).Create(wallet).Error
}

func (r *userWalletRepoImpl) Update(ctx context.Context, id, userID uint64, wallet *entity.UserWallet) error {
	return db.GetDB(ctx, r.db).
		Where("user_id = ?", userID).
		Model(&entity.UserWallet{ID: id}).
		Updates(wallet).Error
}

func (r *userWalletRepoImpl) Delete(ctx context.Context, id, userID uint64) (error, int64) {
	result := db.GetDB(ctx, r.db).
		Where("user_id = ?", userID).
		Delete(&entity.UserWallet{ID: id})

	return result.Error, result.RowsAffected
}

func (r *userWalletRepoImpl) GetUserWallet(id, userID uint64) (*entity.UserWallet, error) {
	var wallet entity.UserWallet
	err := r.db.
		Where("id = ? AND user_id = ?", id, userID).
		First(&wallet).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err_code.UserWalletNotFound
		}
		return nil, err
	}
	return &wallet, nil
}

func (r *userWalletRepoImpl) FindUserWallets(userID uint64) ([]*entity.UserWallet, error) {
	var wallets []*entity.UserWallet
	err := r.db.
		Where("user_id = ?", userID).
		Order("sort desc, id asc").
		Find(&wallets).Error
	return wallets, err
}

func (r *userWalletRepoImpl) UpdateBalance(ctx context.Context, id, userID uint64, amount float64) error {
	result := db.GetDB(ctx, r.db).Model(&entity.UserWallet{}).
		Where("id = ? AND user_id = ?", id, userID).
		UpdateColumn("balance", gorm.Expr("balance + ?", amount))

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return err_code.UserWalletNotFound
	}
	return nil
}

func (r *userWalletRepoImpl) UpdateLiability(ctx context.Context, id, userID uint64, amount float64) error {
	result := db.GetDB(ctx, r.db).Model(&entity.UserWallet{}).
		Where("id = ? AND user_id = ?", id, userID).
		UpdateColumn("liability", gorm.Expr("liability + ?", amount))

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return err_code.UserWalletNotFound
	}
	return nil
}
