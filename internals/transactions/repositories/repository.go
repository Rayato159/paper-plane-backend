package repositories

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/paper-plane/internals/models"
	"github.com/paper-plane/internals/transactions"
)

type transactionsRepository struct {
	db *sqlx.DB
}

func NewTransactionsRepository(db *sqlx.DB) transactions.Repository {
	return &transactionsRepository{db: db}
}

func (r *transactionsRepository) AddTransaction(req *models.CreateReqTransaction, userId string) error {
	query := `
INSERT INTO transactions(
	id,
	income,
	expense,
	account_id
)	
VALUES (
	?,
	?,
	?,
	(SELECT a.id FROM accounts a LEFT JOIN users u ON u.id = a.user_id WHERE a.id = ? AND u.id = ? LIMIT 1)
)
`

	_, err := r.db.Exec(query, req.Id, req.Income, req.Expense, req.AccountId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *transactionsRepository) GetTransactionById(id string, userId string) (*models.Transaction, error) {
	query := `
SELECT
t.id,
t.income,
t.expense,
t.created_at,
t.updated_at,
t.account_id
FROM transactions t
LEFT JOIN accounts a ON a.id = t.account_id
LEFT JOIN users u ON u.id = a.user_id
WHERE t.id = ?
AND u.id = ?
`

	var transaction models.Transaction
	if err := r.db.Get(&transaction, query, id, userId); err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *transactionsRepository) RemoveTransaction(id string, userId string) error {

	queryUser := `
SELECT
u.id
FROM transactions t
LEFT JOIN accounts a ON a.id = t.account_id
LEFT JOIN users u ON u.id = a.user_id
WHERE t.id = ?
AND u.id = ?
`

	var userCheck string
	if err := r.db.Get(&userCheck, queryUser, id, userId); err != nil {
		return errors.New("error, user is invalid.")
	}

	query := `
DELETE FROM transactions 
WHERE id = ?
`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *transactionsRepository) GetTransactionLists(accountId string, reqQuery *models.ReqTransactionQuery, userId string) ([]*models.Transaction, int, error) {
	query := `
SELECT
t.id,
t.income,
t.expense,
t.created_at,
t.updated_at,
t.account_id
FROM transactions t
LEFT JOIN accounts a ON a.id = t.account_id
LEFT JOIN users u ON u.id = a.user_id
WHERE t.account_id = ?
AND u.id = ?
`

	queryCount := `
SELECT
COUNT(*)
FROM transactions t
LEFT JOIN accounts a ON a.id = t.account_id
LEFT JOIN users u ON u.id = a.user_id
WHERE t.account_id = ?
AND u.id = ?
`
	if reqQuery.StartDate != "" && reqQuery.EndDate != "" {
		query += fmt.Sprintf("AND DATE(created_at) BETWEEN '%v' AND '%v'\n", reqQuery.StartDate, reqQuery.EndDate)
		queryCount += fmt.Sprintf("AND DATE(created_at) BETWEEN '%v' AND '%v'\n", reqQuery.StartDate, reqQuery.EndDate)
	}
	pageNew := reqQuery.ItemPerPage * (reqQuery.Page - 1)
	query += fmt.Sprintf("ORDER BY %s %s LIMIT %d OFFSET %d\n", reqQuery.OrderBy, reqQuery.Sort, reqQuery.ItemPerPage, pageNew)

	transactionsChan := make(chan []*models.Transaction)
	transactionsErrChan := make(chan error)

	totalItemChan := make(chan int)
	totalItemErrChan := make(chan error)

	go func() {
		var transactions []*models.Transaction
		err := r.db.Select(&transactions, query, accountId, userId)
		transactionsErrChan <- err
		transactionsChan <- transactions
		close(transactionsErrChan)
		close(transactionsChan)
	}()

	go func() {
		var totalItem int
		err := r.db.Get(&totalItem, queryCount, accountId, userId)
		totalItemErrChan <- err
		totalItemChan <- totalItem
		close(totalItemErrChan)
		close(totalItemChan)
	}()

	totalItemErr := <-totalItemErrChan
	transactionsErr := <-transactionsErrChan

	if transactionsErr != nil {
		return nil, 0, transactionsErr
	}
	if totalItemErr != nil {
		return nil, 0, totalItemErr
	}

	transactions := <-transactionsChan
	totalItem := <-totalItemChan

	return transactions, totalItem, nil
}

func (r *transactionsRepository) EditTransaction(id string, req *models.EditReqTransaction, userId string) (*string, error) {
	query := `
UPDATE transactions t
LEFT JOIN accounts a ON a.id = t.account_id
LEFT JOIN users u ON u.id = a.user_id
SET	
t.income = ?,
t.expense = ?
WHERE t.id = ?
AND u.id = ?
`

	queryAccount := `
SELECT
t.account_id
FROM transactions t
LEFT JOIN accounts a ON a.id = t.account_id
LEFT JOIN users u ON u.id = a.user_id
WHERE t.id = ?
AND u.id = ?
LIMIT 1
`

	updateErrChan := make(chan error)

	accountIdChan := make(chan *string)
	accountErrChan := make(chan error)

	go func() {
		_, err := r.db.Exec(query, req.Income, req.Expense, id, userId)
		updateErrChan <- err
		close(updateErrChan)
	}()

	go func() {
		var accountId string
		err := r.db.Get(&accountId, queryAccount, id, userId)
		accountErrChan <- err
		accountIdChan <- &accountId
		close(accountErrChan)
		close(accountIdChan)
	}()

	updateErr := <-updateErrChan
	accountErr := <-accountErrChan

	if updateErr != nil {
		return nil, updateErr
	}
	if accountErr != nil {
		return nil, accountErr
	}

	accountId := <-accountIdChan

	return accountId, nil
}
