package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jadermoura/caixa-eletronico-app/internal/infra/db/entity"
	"gorm.io/gorm"
)

const (
	MsgNotFoundRecord   string = "not found record for key: %s"
	MsgErrGetSettlement string = "failed to get settlement in database error: %v"
	DefaultBaseDecimal  int    = 10
	MsgDefaultGetOffice string = "failed to get office in database error: %v"
)

func Create(db *gorm.DB, ctx context.Context, account entity.Account) (uint, error) {
	err := db.WithContext(ctx).Create(&account).Error
	if err != nil {
		return uint(0), fmt.Errorf("failed to insert account in database error: %v", err)
	}

	return account.ID, nil

}

func GetBalance(db *gorm.DB, ctx context.Context, n int) ([]entity.Account, error) {
	var ac []entity.Account
	err := db.WithContext(ctx).
		Where("nn_number_aco = ?", n).Find(&ac).Error

	fmt.Println("***************OFFICE************************************")
	ret, _ := json.Marshal(ac)
	fmt.Println(string(ret))

	if len(ac) == 0 {
		return nil, fmt.Errorf(MsgNotFoundRecord, "teste2")
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf(MsgNotFoundRecord, "teste")
		}
		return nil, fmt.Errorf("failed to get balance from database error: %v", err)
	}

	return ac, nil
}
