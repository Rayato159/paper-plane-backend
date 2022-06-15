package transactions

import "github.com/paper-plane/internals/models"

type Usecase interface {
	AddTransaction(req *models.CreateReqTransaction, userId string) (*models.Transaction, error)
	GetTransactionLists(accountId string, reqQuery *models.ReqTransactionQuery, userId string) (*models.TransactionPagination, error)
	GetTransactionById(id string, userId string) (*models.Transaction, error)
	EditTransaction(id string, req *models.EditReqTransaction, userId string) (*models.Transaction, error)
	RemoveTransaction(req *models.TransactionRemoveReq, userId string) error
}
