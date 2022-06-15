package usecases

import (
	"errors"

	"github.com/paper-plane/internals/accounts"
	"github.com/paper-plane/internals/models"
	"github.com/paper-plane/internals/transactions"
	"github.com/paper-plane/internals/users"
	"github.com/paper-plane/pkg/utils"

	"github.com/google/uuid"
)

type transactiosnUsecase struct {
	transactionsRepo transactions.Repository
	accountsRepo     accounts.Repository
	usersRepo        users.Repository
}

func NewTransactionsUsecase(transactionsRepo transactions.Repository, accountsRepo accounts.Repository, usersRepo users.Repository) transactions.Usecase {
	return &transactiosnUsecase{
		transactionsRepo: transactionsRepo,
		accountsRepo:     accountsRepo,
		usersRepo:        usersRepo,
	}
}

func (u *transactiosnUsecase) AddTransaction(req *models.CreateReqTransaction, userId string) (*models.Transaction, error) {
	req.Id = uuid.New().String()

	if req.Income != 0 {
		req.Expense = 0
	} else {
		req.Income = 0
	}

	if err := u.transactionsRepo.AddTransaction(req, userId); err != nil {
		return nil, err
	}

	balance, err := u.accountsRepo.GetBalance(req.AccountId)
	if err != nil {
		if err := u.transactionsRepo.RemoveTransaction(req.Id, userId); err != nil {
			return nil, err
		}
		return nil, err
	}

	if req.Income != 0 {
		req.Balance = balance.Balance + req.Income
	} else if req.Expense != 0 {
		req.Balance = balance.Balance - req.Expense
	}

	updateErrChan := make(chan error)

	transactionChan := make(chan *models.Transaction)
	transactionErrChan := make(chan error)

	go func() {
		err := u.accountsRepo.UpdateBalance(req.AccountId, req.Balance, userId)
		updateErrChan <- err
		close(updateErrChan)
	}()

	go func() {
		transaction, err := u.transactionsRepo.GetTransactionById(req.Id, userId)
		transactionErrChan <- err
		transactionChan <- transaction
		close(transactionErrChan)
		close(transactionChan)
	}()

	updateErr := <-updateErrChan
	transactionErr := <-transactionErrChan

	if updateErr != nil {
		return nil, updateErr
	}
	if transactionErr != nil {
		return nil, transactionErr
	}

	transaction := <-transactionChan

	return transaction, nil
}

func (u *transactiosnUsecase) GetTransactionLists(accountId string, reqQuery *models.ReqTransactionQuery, userId string) (*models.TransactionPagination, error) {
	transactions, totalItem, err := u.transactionsRepo.GetTransactionLists(accountId, reqQuery, userId)
	if err != nil {
		return nil, err
	}

	res := new(models.TransactionPagination)
	res.Page = reqQuery.Page
	res.ItemPerPage = reqQuery.ItemPerPage
	res.TotalItem = totalItem
	res.Data = transactions

	return res, nil
}

func (u *transactiosnUsecase) GetTransactionById(id string, userId string) (*models.Transaction, error) {
	transaction, err := u.transactionsRepo.GetTransactionById(id, userId)
	if err != nil {
		return nil, errors.New("error, transaction not found.")
	}
	return transaction, nil
}

func (u *transactiosnUsecase) EditTransaction(id string, req *models.EditReqTransaction, userId string) (*models.Transaction, error) {
	accountId, err := u.transactionsRepo.EditTransaction(id, req, userId)
	if err != nil {
		return nil, err
	}

	if req.Expense != 0 {
		req.Income = 0
	} else {
		req.Expense = 0
	}

	balance, err := u.accountsRepo.GetBalance(*accountId)
	if err != nil {
		return nil, err
	}

	if req.Expense != 0 {
		balance.Balance -= req.Expense
	} else {
		balance.Balance += req.Income
	}

	if err := u.accountsRepo.UpdateBalance(*accountId, balance.Balance, userId); err != nil {
		return nil, err
	}

	res, err := u.transactionsRepo.GetTransactionById(id, userId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *transactiosnUsecase) RemoveTransaction(req *models.TransactionRemoveReq, userId string) error {
	userPassword, err := u.usersRepo.GetUserPassword(req.UserId)
	if err != nil {
		return err
	}

	if !utils.ComparePasswordHash(req.Password, *userPassword) {
		return errors.New("error, password is incorrect.")
	}

	transaction, err := u.transactionsRepo.GetTransactionById(req.Id, userId)
	if err != nil {
		return err
	}

	balance, err := u.accountsRepo.GetBalance(transaction.AccountId)
	if err != nil {
		return err
	}

	updateBalanceErrChan := make(chan error)
	removeTransErrChan := make(chan error)

	go func() {
		if transaction.Income != 0 {
			balance.Balance -= transaction.Income
		} else {
			balance.Balance += transaction.Expense
		}
		err := u.accountsRepo.UpdateBalance(transaction.AccountId, balance.Balance, userId)
		updateBalanceErrChan <- err
		close(updateBalanceErrChan)
	}()

	go func() {
		err := u.transactionsRepo.RemoveTransaction(req.Id, userId)
		removeTransErrChan <- err
		close(removeTransErrChan)
	}()

	removeErrTrans := <-removeTransErrChan
	updateBalanceErr := <-updateBalanceErrChan

	if updateBalanceErr != nil {
		return updateBalanceErr
	}
	if removeErrTrans != nil {
		return removeErrTrans
	}
	return nil
}
