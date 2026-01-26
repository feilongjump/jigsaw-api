package db

import (
	"context"

	"github.com/feilongjump/jigsaw-api/domain/repo"
	"gorm.io/gorm"
)

type contextKey string

const txKey contextKey = "tx"

type transactionManager struct {
	db *gorm.DB
}

func NewTransactionManager() repo.TransactionManager {
	return &transactionManager{db: Get()}
}

func (tm *transactionManager) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return tm.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		c := context.WithValue(ctx, txKey, tx)
		return fn(c)
	})
}

func GetDB(ctx context.Context, defaultDB *gorm.DB) *gorm.DB {
	tx, ok := ctx.Value(txKey).(*gorm.DB)
	if ok {
		return tx
	}
	return defaultDB.WithContext(ctx)
}
