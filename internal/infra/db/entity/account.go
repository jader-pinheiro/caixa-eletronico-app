package entity

import "gorm.io/gorm"

type Account struct {
	gorm.Model

	Number  int     `gorm:"column:nn_number_aco;type:varchar(64)" json:"number_account"`
	Key     int     `gorm:"column:nn_key_aco;type:varchar(64)" json:"key_account"`
	Balance float64 `gorm:"column:nn_balance_aco;type:varchar(64)" json:"balance_account"`
}
