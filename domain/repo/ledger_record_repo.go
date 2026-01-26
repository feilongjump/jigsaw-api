package repo

import (
	"context"

	"github.com/feilongjump/jigsaw-api/domain/entity"
)

type LedgerRecordRepo interface {
	Create(ctx context.Context, record *entity.LedgerRecord) error
	Update(ctx context.Context, id, userID uint64, record *entity.LedgerRecord) error
	Delete(ctx context.Context, id, userID uint64) (error, int64)
	GetLedgerRecord(id, userID uint64) (*entity.LedgerRecord, error)
	FindLedgerRecords(userID uint64, page, pageSize int, filter map[string]interface{}) ([]*entity.LedgerRecord, int64, error)
	CountByWalletID(userID, walletID uint64) (int64, error)
}
