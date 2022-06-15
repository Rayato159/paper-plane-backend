package transactions

import "github.com/paper-plane/internals/models"

type Repository interface {
	AddTransaction(req *models.CreateReqTransaction, userId string) error
	GetTransactionById(id string, userId string) (*models.Transaction, error)
	GetTransactionLists(accountId string, reqQuery *models.ReqTransactionQuery, userId string) ([]*models.Transaction, int, error)
	EditTransaction(id string, req *models.EditReqTransaction, userId string) (*string, error)
	RemoveTransaction(id string, userId string) error
}
