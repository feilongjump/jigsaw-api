package repo

import (
	"context"

	"github.com/feilongjump/jigsaw-api/domain/entity"
)

type UserWalletRepo interface {
	Create(ctx context.Context, wallet *entity.UserWallet) error
	Update(ctx context.Context, id, userID uint64, wallet *entity.UserWallet) error
	Delete(ctx context.Context, id, userID uint64) (error, int64)
	GetUserWallet(id, userID uint64) (*entity.UserWallet, error)
	FindUserWallets(userID uint64) ([]*entity.UserWallet, error)
	// UpdateBalance 更新余额 (原子操作)
	UpdateBalance(ctx context.Context, id, userID uint64, amount float64) error
	// UpdateLiability 更新负债 (原子操作)
	UpdateLiability(ctx context.Context, id, userID uint64, amount float64) error
}
