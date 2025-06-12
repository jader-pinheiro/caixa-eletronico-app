package data

import (
	"context"
	"fmt"

	"github.com/jadermoura/caixa-eletronico-app/internal/infra/db/entity"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, ctx context.Context, account entity.Account) (uint, error) {
	err := db.WithContext(ctx).Create(&account).Error
	if err != nil {
		return uint(0), fmt.Errorf("failed to insert account in database error: %v", err)
	}

	return account.ID, nil

}
