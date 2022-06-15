package usecases

import (
	"github.com/paper-plane/internals/accounts"
	"github.com/paper-plane/internals/models"
)

type accountsUsecase struct {
	accountsRepo accounts.Repository
}

func NewAccountsUsecase(accountsRepo accounts.Repository) accounts.Usecase {
	return &accountsUsecase{accountsRepo: accountsRepo}
}

func (u *accountsUsecase) GetAccountInfo(accountId string, userId string) (*models.AccountInfo, error) {
	res, err := u.accountsRepo.GetAccountInfo(accountId, userId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *accountsUsecase) UpdateBalance(accountId string, balance float64, userId string) (*models.AccountInfo, error) {
	if err := u.accountsRepo.UpdateBalance(accountId, balance, userId); err != nil {
		return nil, err
	}

	res, err := u.accountsRepo.GetAccountInfo(accountId, userId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
