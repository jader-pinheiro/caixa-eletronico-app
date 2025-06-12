package db

import (
	"os"

	"github.com/jadermoura/caixa-eletronico-app/internal/infra/db/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(os.Getenv("GORM_DSN")), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entity.Account{})

	return db, nil
}
