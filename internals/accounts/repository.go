package accounts

import "github.com/paper-plane/internals/models"

type Repository interface {
	CreateAccount(userId string, balance float64) error
	GetAccountInfo(accountId string, userId string) (*models.AccountInfo, error)
	UpdateBalance(accountId string, balance float64, userId string) error
	GetBalance(accountId string) (*models.Balance, error)
}
