package db

import (
	"fmt"
	"os"

	"github.com/jadermoura/caixa-eletronico-app/internal/infra/db/entity"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
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

// package main

// import (
//     "gorm.io/driver/postgres"
//     "gorm.io/gorm"
//     "os"
//     "your_project/entity" // Altere para o caminho correto da sua entidade
// )

func ConnectPOstgress() (*gorm.DB, error) {
	// Defina a string de conexão com PostgreSQL
	dsn := os.Getenv("GORM_DSN")

	// Se a string de conexão estiver vazia, podemos lançar um erro
	if dsn == "" {
		return nil, fmt.Errorf("GORM_DSN environment variable not set")
	}

	// Conectando ao banco de dados PostgreSQL com GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	// Realizando migração da tabela Account
	err = db.AutoMigrate(&entity.Account{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
