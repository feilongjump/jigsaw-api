package repo

import (
	"github.com/feilongjump/jigsaw-api/domain/entity"
)

type UserWalletRepo interface {
	Create(wallet *entity.UserWallet) error
	Update(id, userID uint64, wallet *entity.UserWallet) error
	Delete(id, userID uint64) (error, int64)
	GetUserWallet(id, userID uint64) (*entity.UserWallet, error)
	FindUserWallets(userID uint64) ([]*entity.UserWallet, error)
	// UpdateBalance 更新余额 (原子操作)
	UpdateBalance(id, userID uint64, amount float64) error
	// UpdateLiability 更新负债 (原子操作)
	UpdateLiability(id, userID uint64, amount float64) error
}
