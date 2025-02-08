package pkg

import (
	"context"

	"gorm.io/gorm"
)

const TXKey = "tx"

func SetTx(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, TXKey, db)
}

func GetTx(ctx context.Context) (*gorm.DB, bool) {
	tx, ok := ctx.Value(TXKey).(*gorm.DB)
	return tx, ok
}
