package accounts

import "github.com/paper-plane/internals/models"

type Usecase interface {
	GetAccountInfo(accountId string, userId string) (*models.AccountInfo, error)
	UpdateBalance(accountId string, balance float64, userId string) (*models.AccountInfo, error)
}
