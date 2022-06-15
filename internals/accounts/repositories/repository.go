package repositories

import (
	"github.com/paper-plane/internals/accounts"
	"github.com/paper-plane/internals/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type accountsRepository struct {
	db *sqlx.DB
}

func NewAccountsRepository(db *sqlx.DB) accounts.Repository {
	return &accountsRepository{db: db}
}

func (r *accountsRepository) CreateAccount(userId string, balance float64) error {
	query := `
INSERT INTO accounts(
	id,
	balance,
	user_id
)
VALUES (
	?,
	?,
	(SELECT id FROM users WHERE id = ? LIMIT 1)
)
`

	_, err := r.db.Exec(query, uuid.New(), balance, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *accountsRepository) GetAccountInfo(accountId string, userId string) (*models.AccountInfo, error) {
	query := `
SELECT
a.id,
a.balance,
u.id,
u.username,
a.created_at,
a.updated_at
FROM accounts a
LEFT JOIN users u
ON u.id = a.user_id
WHERE a.id = ?
AND u.id = ?
`

	rows, err := r.db.Query(query, accountId, userId)
	if err != nil {
		return nil, err
	}

	var account models.AccountInfo
	for rows.Next() {
		if err := rows.Scan(
			&account.Id,
			&account.Balance,
			&account.User.Id,
			&account.User.Username,
			&account.CreatedAt,
			&account.UpdatedAt,
		); err != nil {
			return nil, err
		}
		break
	}
	rows.Close()

	return &account, nil
}

func (r *accountsRepository) UpdateBalance(accountId string, balance float64, userId string) error {
	query := `
UPDATE accounts
SET
balance = ?
WHERE id = ?
LIMIT 1
`

	_, err := r.db.Exec(query, balance, accountId)
	if err != nil {
		return err
	}
	return nil
}

func (r *accountsRepository) GetBalance(accountId string) (*models.Balance, error) {
	query := `
SELECT
balance
FROM accounts
WHERE id = ?
LIMIT 1
`

	var balance models.Balance
	if err := r.db.Get(&balance, query, accountId); err != nil {
		return nil, err
	}
	return &balance, nil
}
